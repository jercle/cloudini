package ado

import (
	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

// Downloads default to $HOME/.config/cld/cache/aib-logs
// if nil passed into function for downloadPathParam
func DownloadPackerHostLogs(downloadPathParam *string) (downloadPath string, dlFileCount int) {
	_, _, cachePath := lib.InitConfig(nil)
	cfg := lib.GetCldConfig(nil)
	if downloadPathParam == nil {
		downloadPath = cachePath + "/aib-logs"
	} else {
		downloadPath = *downloadPathParam
	}

	options := lib.StorageAccountRequestOptions{
		StorageAccountName:   cfg.AzureDevOps.Packer.Logs.StorageAcct,
		ContainerName:        cfg.AzureDevOps.Packer.Logs.BlobContainer.Hosts,
		ConfiguredTenantName: cfg.AzureDevOps.Packer.Logs.TenantName,
		OverwriteExisting:    false,
		ShowDownloadedCount:  false,
		DownloadPath:         downloadPath,
	}

	// jsonStr, _ := json.MarshalIndent(options, "", "  ")
	// fmt.Println(string(jsonStr))
	// os.Exit(0)

	dlFileCount = azure.DownloadAllBlobsInContainer(options)

	return
}
