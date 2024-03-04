package azure

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/jercle/azg/lib"
	"github.com/rodaine/table"
)

type nsgFlowLogRecord struct {
	Category      string `json:"category"`
	MacAddress    string `json:"macAddress"`
	OperationName string `json:"operationName"`
	Properties    struct {
		Version int `json:"Version"`
		Flows   []struct {
			Flows []struct {
				FlowTuples []string `json:"flowTuples"`
				Mac        string   `json:"mac"`
			} `json:"flows"`
			Rule string `json:"rule"`
		} `json:"flows"`
	} `json:"properties"`
	ResourceID string    `json:"resourceId"`
	SystemID   string    `json:"systemId"`
	Time       time.Time `json:"time"`
}

type nsgFlowLog struct {
	Records []nsgFlowLogRecord
}

type combinedFlowLogs struct {
	nsgFlowLogs []nsgFlowLogRecord
	FileCount   int
}

type ipList struct {
	SourceIps []string
	DestIps   []string
}

func (m *ipList) printCount() {
	fmt.Println("Source IPs:      ", len(m.SourceIps))
	fmt.Println("Destination IPs: ", len(m.DestIps))
}

func getUniqueIpAddresses(dataset []nsgFlowLogRecord) ipList {
	var ipList ipList

	for _, record := range dataset {
		for _, outerFlow := range record.Properties.Flows {
			for _, innerFlow := range outerFlow.Flows {
				for _, tuple := range innerFlow.FlowTuples {
					split := strings.Split(tuple, ",")
					ipList.SourceIps = append(ipList.SourceIps, split[1])
					ipList.DestIps = append(ipList.DestIps, split[2])
				}
			}
		}
	}
	ipList.SourceIps = lib.UniqueNonEmptyElementsOf(ipList.SourceIps)
	ipList.DestIps = lib.UniqueNonEmptyElementsOf(ipList.DestIps)
	return ipList
}

func getFlowLogData(path string) nsgFlowLog {
	var flowLogData nsgFlowLog

	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &flowLogData)

	return flowLogData
}

func combineLogFileRecords(dataPath string) combinedFlowLogs {
	var (
		wg             sync.WaitGroup
		allFlowLogData combinedFlowLogs
		mutex          sync.Mutex
	)

	filePaths := lib.GetFullFilePaths(dataPath)
	allFlowLogData.FileCount = len(filePaths)

	for _, file := range filePaths {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := getFlowLogData(file).Records
			mutex.Lock()
			allFlowLogData.nsgFlowLogs = append(allFlowLogData.nsgFlowLogs, data...)
			mutex.Unlock()
		}()
	}
	wg.Wait()

	return allFlowLogData
}

func (m *ipList) printTable() {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("IP Address", "Source/Dest")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, ipAddress := range *&m.DestIps {
		tbl.AddRow(ipAddress, "Destination")
	}
	for _, ipAddress := range *&m.SourceIps {
		tbl.AddRow(ipAddress, "Source")
	}
	tbl.Print()
}

func (r *combinedFlowLogs) filterIp(filter string, filterDirection string) {
	var filteredLogs combinedFlowLogs
	filteredLogs.FileCount = r.FileCount
	filterSlice := strings.Split(filter, ",")

	for _, record := range r.nsgFlowLogs {
	recordLoop:
		for _, outerFlow := range record.Properties.Flows {
			for _, innerFlow := range outerFlow.Flows {
				for _, tuple := range innerFlow.FlowTuples {
					split := strings.Split(tuple, ",")
					if strings.ToLower(filterDirection) == "source" {
						if slices.Contains(filterSlice, split[1]) {
							filteredLogs.nsgFlowLogs = append(filteredLogs.nsgFlowLogs, record)
							break recordLoop
						}
					}
					if strings.ToLower(filterDirection) == "dest" {
						if slices.Contains(filterSlice, split[2]) {
							filteredLogs.nsgFlowLogs = append(filteredLogs.nsgFlowLogs, record)
							break recordLoop
						}
					}
				}
			}
		}
	}
	*r = filteredLogs
}

func (m *ipList) filterSourceIp(filter string) {
	var filteredTables ipList
	filterSlice := strings.Split(filter, ",")

	for _, ipAddress := range *&m.SourceIps {
		// if strings.Contains(ipAddress, filter) {
		if slices.Contains(filterSlice, ipAddress) {
			filteredTables.SourceIps = append(filteredTables.SourceIps, ipAddress)
		}
	}

	*m = filteredTables
}

func (r *nsgFlowLogRecord) printJSON() {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}
