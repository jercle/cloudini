package main

import (
	"time"

	"github.com/jercle/cloudini/cmd/azure"
)

func main() {

	startTime := time.Now()

	// jsonBytes, _ := json.MarshalIndent(blobList, "", "  ")
	// fmt.Println(string(jsonBytes))

	// config := lib.GetCldConfig(nil)

	// tenant := config.Azure.MultiTenantAuth.Tenants["REDDTQ"]
	// split := strings.Split(tenant.CostExportsLocation, "/")

	// options := StorageAccountRequestOptions{
	// 	StorageAccountName:   strings.Split(split[2], ".")[0],
	// 	ContainerName:        split[3],
	// 	ConfiguredTenantName: tenant.TenantName,
	// }
	// blobs := ListStorageContainerBlobs(options)
	// fmt.Println(len(blobs))
	azure.DownloadAllConfiguredTenantLastMonthCostExports(azure.DownloadAllConfiguredTenantLastMonthCostExportsOptions{
		BlobPrefix:  "monthly-cost-exports/202406",
		OutfilePath: "cost-exports/to-process/cost-export",
	})
	elapsed := time.Since(startTime)
	_ = elapsed
}
