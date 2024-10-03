package main

import (
	"os"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func main() {
	costExportMonth := "202407"

	dataPath := "./cost-exports/" + costExportMonth + "/"

	if !lib.CheckDirExists(dataPath) {
		os.MkdirAll(dataPath, os.ModePerm)
	}

	azure.DownloadAllConfiguredTenantLastMonthCostExports(azure.DownloadAllConfiguredTenantLastMonthCostExportsOptions{
		BlobPrefix:  "monthly-cost-exports/" + costExportMonth,
		OutfilePath: dataPath + "cost-export",
	})

	combinedCostData := azure.CombineCostExportCSVData(dataPath)
	transformedData := azure.TransformCostData(combinedCostData)

	azure.CostDataToExcel(transformedData, dataPath+"MonthlyCostReport-"+costExportMonth+".xlsx")
}

// https://SITE.sharepoint.com/
// https://SITE.sharepoint.com/:f:/r/sites/TEAM/PATH
// /sites/TEAM/drive/items/{item-id}/content
// /sites/TEAM/drive/items/{item-id}/children
