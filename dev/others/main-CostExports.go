package main

import (
	"github.com/jercle/azg/cmd/azure"
)

func main() {
	// now := time.Now()
	// // .Format("200601")
	// year := strconv.Itoa(now.Year())
	// month := fmt.Sprintf("%02d", int(now.Month()))
	// fileName := "MonthlyReport-" + year + month + ".xlsx"

	var dataPath = "./fakedata/cost-exports"
	// var combinedCostData costExportData

	// jsonFile, err := os.Open("fakedata/cost-data/blue.json")
	// lib.CheckFatalError(err)
	// defer jsonFile.Close()

	// byteValue, _ := io.ReadAll(jsonFile)
	// json.Unmarshal(byteValue, &combinedCostData)

	combinedCostData := azure.CombineCostExportJSONData(dataPath)

	// ccdJson, _ := json.MarshalIndent(combinedCostData, "", "  ")
	// fmt.Println(string(ccdJson))

	// fmt.Println(len(combinedCostData))

	transformedData := azure.TransformCostData(combinedCostData)

	azure.CostDataToExcel(transformedData, "outputs/combined.xlsx")

	// costData, err := getCostExportCSVFileData("cost-exports/monthly-cost-exports_BLUEDTQ.csv")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(combinedCostData)
	// jsonData, _ := json.MarshalIndent(transformedData, "", "  ")
	// fmt.Println(string(jsonData))
	// _ = jsonData
	// fmt.Println(len(costData))

}
