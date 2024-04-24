// main-nsgFlowLogs.go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/jercle/cloudini/cmd/azure"
)

func main() {

	// var filePaths = getFullFilePaths("./fakedata/nsgLogs")
	// var dataPath = "../../azgo/dev/nsgLogs"
	var dataPath = "./outputs/nsgflows"

	combinedData := azure.CombineLogFileRecords(dataPath)
	// combinedData.FilterIp("172.233.228.93,144.172.79.92,66.235.168.222", "source")
	combinedData.FilterIp("172.26.0.4", "dest")
	// combinedData.filterIp("192.168.0.1,192.168.0.2,200.197.39.223", "dest")
	// uniqueIps := azure.GetUniqueIpAddresses(combinedData.NsgFlowLogs)

	// fmt.Println()
	// fmt.Println("Files processed: ", combinedData.FileCount)
	// uniqueIps.filterSourceIp("192.168.0.1,192.168.0.2,130.111.165.184")
	// fmt.Println(combinedData.NsgFlowLogs)
	// fmt.Println(len(combinedData.NsgFlowLogs))
	// uniqueIps.PrintCount()
	// uniqueIps.PrintTable()

	jsonBytes, _ := json.MarshalIndent(combinedData.NsgFlowLogs, "", "  ")
	fmt.Println(string(jsonBytes))

	// for _, r := range combinedData.NsgFlowLogs {
	// 	fmt.Println(r)
	// 	r.PrintJSON()
	// 	os.Exit(0)
	// }
}
