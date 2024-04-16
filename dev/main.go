// main-nsgFlowLogs.go
package main

import (
	"fmt"

	"github.com/jercle/azg/cmd/azure"
)

func main() {

	// var filePaths = getFullFilePaths("./fakedata/nsgLogs")
	// var dataPath = "../../azgo/dev/nsgLogs"
	var dataPath = "./outputs/dan"

	combinedData := azure.CombineLogFileRecords(dataPath)
	combinedData.FilterIp("172.233.228.93,144.172.79.92,66.235.168.222", "source")
	combinedData.FilterIp("172.233.228.93,144.172.79.92,66.235.168.222", "dest")
	// combinedData.filterIp("192.168.0.1,192.168.0.2,200.197.39.223", "dest")
	uniqueIps := azure.GetUniqueIpAddresses(combinedData.NsgFlowLogs)

	// fmt.Println()
	fmt.Println("Files processed: ", combinedData.FileCount)
	// uniqueIps.filterSourceIp("192.168.0.1,192.168.0.2,130.111.165.184")
	// fmt.Println(len(combinedData.nsgFlowLogs))
	// uniqueIps.PrintCount()
	uniqueIps.PrintTable()

	// for _, r := range combinedData.nsgFlowLogs {
	// 	fmt.Println(record)
	// 	r.printJSON()
	// 	os.Exit(0)
	// }
}
