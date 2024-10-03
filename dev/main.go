package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/jercle/cloudini/cmd/azure"
)

type envConfig struct {
	ClientId          string
	TenantId          string
	ServiceConnection string
}

func main() {
	lastMonth := time.Now().AddDate(0, -1, 0)
	month := lastMonth.Format("01")
	year := lastMonth.Format("2006")
	costExportMonth := year + month

	// agentTempDir := os.Getenv("AGENT_TEMPDIRECTORY")
	// fileDirectory := os.Getenv("AGENT_TEMPDIRECTORY")
	fileDirectory := "."
	dataPath := fileDirectory + "/cost-exports/" + costExportMonth + "/"
	// outputFileName := dataPath + "MonthlyCostReport-" + costExportMonth + ".xlsx"
	// _ = outputFileName
	outputJsonName := dataPath + "MonthlyCostReport-" + costExportMonth + ".json"
	_ = outputJsonName

	// os.RemoveAll(dataPath)

	// azure.GenerateMonthlyCostReport(dataPath, outputFileName, "monthly-cost-exports", costExportMonth)

	// func GenerateMonthlyCostReport(outputDirectory string, outputFileName string, blobPrefix string, costExportMonth string) {
	// if !lib.CheckDirExists(dataPath) {
	// 	os.MkdirAll(dataPath, os.ModePerm)
	// }

	// cldConfOpts := &lib.CldConfigOptions{
	// 	ConfigFilePath: configFilePath,
	// }

	// DownloadAllConfiguredTenantLastMonthCostExports(DownloadAllConfiguredTenantLastMonthCostExportsOptions{
	// 	BlobPrefix:  blobPrefix + "/" + costExportMonth,
	// 	OutfilePath: outputDirectory + "cost-export",
	// }, nil)

	combinedCostData := azure.CombineCostExportCSVData(dataPath)
	transformedData := azure.TransformCostDataNew(combinedCostData)
	_ = transformedData
	jsonStr, _ := json.MarshalIndent(transformedData, "", "  ")
	os.WriteFile(outputJsonName, jsonStr, os.ModePerm)

	// transformedData.SumCosts()

	// CostDataToExcel(transformedData, outputFileName)
	// }

	// siteId := ""
	// driveId := """

	// fileName := "MonthlyCostReport-" + costExportMonth + ".xlsx"
	// folderPath := ""

	// tokenReq, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{
	// 	Scope:         "graph",
	// 	GetWriteToken: true,
	// }, nil)
	// lib.CheckFatalError(err)

	// token, err := tokenReq.SelectTenant("")
	// lib.CheckFatalError(err)

	// m365.UploadFileToSharepoint(dataPath, siteId, driveId, folderPath, fileName, token.TokenData.Token)
}
