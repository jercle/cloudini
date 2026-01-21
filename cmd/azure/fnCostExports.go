package azure

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/briandowns/spinner"
	"github.com/jercle/cloudini/lib"
	"github.com/xuri/excelize/v2"
)

func CostDataToExcel(data lib.TransformedCostItemsByTenant, outFileName string) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	for tenant, tData := range data {
		fmt.Println("Processing '" + tenant + "' data")
		// fmt.Println(tData)
		_, err := f.NewSheet(tenant)
		err = f.SetSheetRow(tenant, "A1", &[]string{
			// "Subscription",
			// "Resource Group",
			// "PreTaxCost",
			"SubscriptionName",
			"ResourceGroup",
			"PreTaxCost",
			"UsageDateTime",
			"ProductName",
			"MeterCategory",
			"MeterSubcategory",
			"MeterName",
			"UnitOfMeasure",
			"UsageQuantity",
			"ResourceRate",
			"CostCenter",
			"ResourceType",
			"ConsumedService",
			"ResourceName",
			"InstanceId",
			"Tags",
			"AdditionalInfo",
		})

		row := 2
		for _, rowData := range tData.ResGroups {
			var cell string
			cell, _ = excelize.CoordinatesToCellName(1, row)
			f.SetCellValue(tenant, cell, rowData.SubscriptionName)
			cell, _ = excelize.CoordinatesToCellName(2, row)
			f.SetCellValue(tenant, cell, rowData.ResourceGroup)
			cell, _ = excelize.CoordinatesToCellName(3, row)
			f.SetCellValue(tenant, cell, rowData.PreTaxCost)
			cell, _ = excelize.CoordinatesToCellName(4, row)
			f.SetCellValue(tenant, cell, rowData.UsageDateTime)
			cell, _ = excelize.CoordinatesToCellName(5, row)
			f.SetCellValue(tenant, cell, rowData.ProductName)
			cell, _ = excelize.CoordinatesToCellName(6, row)
			f.SetCellValue(tenant, cell, rowData.MeterCategory)
			cell, _ = excelize.CoordinatesToCellName(7, row)
			f.SetCellValue(tenant, cell, rowData.MeterSubcategory)
			cell, _ = excelize.CoordinatesToCellName(8, row)
			f.SetCellValue(tenant, cell, rowData.MeterName)
			cell, _ = excelize.CoordinatesToCellName(9, row)
			f.SetCellValue(tenant, cell, rowData.UnitOfMeasure)
			cell, _ = excelize.CoordinatesToCellName(10, row)
			f.SetCellValue(tenant, cell, rowData.UsageQuantity)
			cell, _ = excelize.CoordinatesToCellName(11, row)
			f.SetCellValue(tenant, cell, rowData.ResourceRate)
			// f.SetCellValue(tenant, cell, rowData.CostCenter)
			cell, _ = excelize.CoordinatesToCellName(12, row)
			f.SetCellValue(tenant, cell, rowData.ResourceType)
			cell, _ = excelize.CoordinatesToCellName(13, row)
			f.SetCellValue(tenant, cell, rowData.ConsumedService)
			cell, _ = excelize.CoordinatesToCellName(14, row)
			f.SetCellValue(tenant, cell, rowData.ResourceName)
			cell, _ = excelize.CoordinatesToCellName(15, row)
			f.SetCellValue(tenant, cell, rowData.Tags)
			cell, _ = excelize.CoordinatesToCellName(16, row)
			f.SetCellValue(tenant, cell, rowData.AdditionalInfo)
			cell, _ = excelize.CoordinatesToCellName(17, row)
			// cell, _ = excelize.CoordinatesToCellName(18, row)
			f.SetCellValue(tenant, cell, rowData.InstanceId)

			// err = f.SetCellValue(tenant, cell, rowData.ResourceGroup)
			lib.CheckFatalError(err)
			// if err != nil {
			// 	time.Sleep(10 * time.Minute)
			// }
			row++
		}

		// TODO: Add pivot table
		if err := f.AddPivotTable(&excelize.PivotTableOptions{
			DataRange:       tenant + "!A1:C" + strconv.Itoa(row),
			PivotTableRange: tenant + "!G2:M34",
			Rows: []excelize.PivotTableField{
				{Data: "Subscription"}},
			// Columns: []excelize.PivotTableField{
			// 	{Data: "Subscription", DefaultSubtotal: true}},
			Data: []excelize.PivotTableField{
				{Data: "PreTaxCost", Name: "Sum of PreTaxcost", Subtotal: "Sum"}},
			RowGrandTotals: true,
			ColGrandTotals: true,
			ShowDrill:      true,
			ShowRowHeaders: true,
			ShowColHeaders: true,
			ShowLastColumn: true,
		}); err != nil {
			fmt.Println(err)
			return
		}

		// Autofit all columns according to their text content
		cols, err := f.GetCols(tenant)
		if err != nil {
			fmt.Println(err)
		}
		for idx, col := range cols {
			largestWidth := 0
			for _, rowCell := range col {
				cellWidth := utf8.RuneCountInString(rowCell) + 2 // + 2 for margin
				if cellWidth > largestWidth {
					largestWidth = cellWidth
				}
			}
			name, err := excelize.ColumnNumberToName(idx + 1)
			if err != nil {
				fmt.Println(err)
			}
			f.SetColWidth(tenant, name, name, float64(largestWidth))
		}
		fmt.Println("Finished processing '" + tenant + "' data")
	}

	err := f.DeleteSheet("Sheet1")
	lib.CheckFatalError(err)
	err = f.SaveAs(outFileName)
	fmt.Println("Saved to " + outFileName)
	lib.CheckFatalError(err)
}

func CombineCostExportCSVData(dataPath string) lib.CostExportData {
	var (
		wg             sync.WaitGroup
		costExportData lib.CostExportData
		mutex          sync.Mutex
		filePaths      = lib.GetFullFilePaths(dataPath)
		// productNames       []string
		// meterCategories    []string
		// meterSubcategories []string
		// meterIds           []string
		// meterNames         []string
		// meterRegions       []string
		// unitsOfMeasure     []string
		// consumedServices   []string
		// resourceTypes      []string
		// offerIds           []string
	)
	for _, file := range filePaths {
		// if true {
		if strings.HasSuffix(file, ".csv") {
			wg.Add(1)
			// fmt.Println(file)
			go func() {
				defer wg.Done()
				data, err := GetCostExportCSVFileData(file)
				if err != nil {
					fmt.Println(file)
					panic(err)
				}

				// jsonStr, _ := json.MarshalIndent(data, "", "  ")
				// fmt.Println(string(jsonStr))
				// os.Exit(0)
				mutex.Lock()
				// productNames = append(productNames)
				costExportData = append(costExportData, data...)
				mutex.Unlock()
			}()
		}
	}
	wg.Wait()

	// jsonStr, _ := json.MarshalIndent(costExportData, "", "  ")
	// fmt.Println(string(jsonStr))
	// os.Exit(0)
	return costExportData
}

func CombineCostExportJSONData(dataPath string) lib.CostExportData {
	var (
		wg             sync.WaitGroup
		costExportData lib.CostExportData
		mutex          sync.Mutex
		filePaths      = lib.GetFullFilePaths(dataPath)
		// productNames       []string
		// meterCategories    []string
		// meterSubcategories []string
		// meterIds           []string
		// meterNames         []string
		// meterRegions       []string
		// unitsOfMeasure     []string
		// consumedServices   []string
		// resourceTypes      []string
		// offerIds           []string
	)
	for _, file := range filePaths {
		if true {
			// if strings.HasSuffix(file, ".csv") {
			wg.Add(1)
			go func() {
				defer wg.Done()
				// fmt.Println(file)
				data, err := GetCostExportJSONFileData(file)
				if err != nil {
					panic(err)
				}
				mutex.Lock()
				// productNames = append(productNames)
				costExportData = append(costExportData, data...)
				mutex.Unlock()
			}()
		}
	}
	wg.Wait()
	return costExportData
}

func TransformCostDataNew(data lib.CostExportData, progBarNum int, progBarTotal int) lib.AggregatedCostData {
	cfg := lib.GetCldConfig(nil)

	allDataNEW := make(lib.AggregatedCostData)

	bar := lib.ProgressBar(len(data), "resource", progBarNum, progBarTotal, "Transforming cost data...")

	for _, costData := range data {
		var (
			tagData            map[string]string
			additionalInfoData interface{}
		)

		// sn := strings.ToLower(costData.SubscriptionName)
		instanceId := strings.ToLower(costData.InstanceId)

		tenantName := ""
		custTntName := lib.MapAzureSubscriptionToCustomTenantName(costData.SubscriptionGuid, *cfg.Azure)
		if custTntName != "" {
			tenantName = custTntName
		} else {
			tenantName = strings.ToUpper(costData.Datafile)
		}

		resourceNameSplit := strings.Split(instanceId, "/")
		resourceName := resourceNameSplit[len(resourceNameSplit)-1]
		identifierResNameSplit := strings.Split(instanceId, "providers/")
		identifierResName := "providers/" + strings.Join(identifierResNameSplit[1:], "providers/")

		resourceMeterIdentifier := ""
		resourceMeterIdentifier = costData.SubscriptionGuid +
			"/"
		if costData.ResourceGroup != "" {
			resourceMeterIdentifier += strings.ToLower(costData.ResourceGroup) +
				"/"
		}

		resourceMeterIdentifier += strings.ToLower(identifierResName) +
			"_" +
			strings.ToLower(costData.MeterCategory) +
			"_" +
			strings.ToLower(costData.MeterSubcategory) +
			"_" +
			strings.ToLower(costData.MeterName)

		tagStr := []byte("{" + costData.Tags + "}")
		json.Unmarshal(tagStr, &tagData)
		json.Unmarshal([]byte(costData.AdditionalInfo), &additionalInfoData)

		resGrp := ""

		if costData.ResourceGroup == "" {
			resGrp = "nil_res_group"
		} else {
			resGrp = strings.ToLower(costData.ResourceGroup)
		}

		tci := lib.TransformedCostItem{
			// ResourceGroup:    costData.ResourceGroup,
			// PreTaxCost:       costData.PreTaxCost,
			// SubscriptionName: costData.SubscriptionName,
			SubscriptionName: strings.ToLower(costData.SubscriptionName),
			SubscriptionId:   costData.SubscriptionGuid,
			ResourceGroup:    resGrp,
			UsageDateTime:    strings.ToLower(costData.UsageDateTime),
			ProductName:      strings.ToLower(costData.ProductName),
			MeterCategory:    strings.ToLower(costData.MeterCategory),
			MeterSubcategory: strings.ToLower(costData.MeterSubcategory),
			MeterName:        strings.ToLower(costData.MeterName),
			UnitOfMeasure:    costData.UnitOfMeasure,
			UsageQuantity:    costData.UsageQuantity,
			ResourceRate:     costData.ResourceRate,
			PreTaxCost:       costData.PreTaxCost,
			// CostCenter:       costData.CostCenter,
			ConsumedService: strings.ToLower(costData.ConsumedService),
			ResourceType:    strings.ToLower(costData.ResourceType),
			InstanceId:      strings.ToLower(instanceId),
			Tags:            tagData,
			// AdditionalInfo:  costData.AdditionalInfo,
			AdditionalInfo: additionalInfoData,

			ResourceMeterIdentifier: resourceMeterIdentifier,
			ResourceName:            strings.ToLower(resourceName),
			Tenant:                  strings.ToUpper(tenantName),
			Datafile:                costData.Datafile,
		}

		// jsonStr, _ := json.MarshalIndent(allDataNEW, "", "  ")
		// os.WriteFile("main_allDataNEW.json", jsonStr, 0644)
		// os.Exit(0)

		allDataNEW = AggregateCostData(allDataNEW, tci)
		bar.Add(1)
	}
	fmt.Println("")
	return allDataNEW
}

func FlattenAggregatedCostData(data lib.AggregatedCostData, resources []lib.AzureResourceDetails) []lib.AzureResourceDetails {

	byInstanceId := make(map[string][]lib.AggregatedCostItem)
	var processedResources []lib.AzureResourceDetails

	for _, tenantData := range data {
		for _, subData := range tenantData.Subscriptions {
			for _, resGrpData := range subData.ResourceGroups {
				for resName, resData := range resGrpData.Resources {
					_ = resName
					for _, meterData := range resData.MeterData {
						byInstanceId[strings.ToLower(meterData.InstanceId)] = append(byInstanceId[strings.ToLower(meterData.InstanceId)], meterData)
					}
				}
			}
		}
	}
	for _, res := range resources {
		var currRes lib.AzureResourceDetails

		jsonStr, _ := json.MarshalIndent(res, "", "  ")
		err := json.Unmarshal(jsonStr, &currRes)
		lib.CheckFatalError(err)

		// currRes.MeterData = byInstanceId[strings.ToLower(res.ID)]
		processedResources = append(processedResources, currRes)
	}

	// jsonStr, _ := json.MarshalIndent(processedResources, "", "  ")

	// os.WriteFile("cost-exports/byResource.json", jsonStr, 0644)
	return processedResources
}

func CostDataByInstanceId(data lib.AggregatedCostData) map[string][]lib.AggregatedCostItem {
	byInstanceId := make(map[string][]lib.AggregatedCostItem)

	for _, tenantData := range data {
		for _, subData := range tenantData.Subscriptions {
			for _, resGrpData := range subData.ResourceGroups {
				for resName, resData := range resGrpData.Resources {
					_ = resName
					for _, meterData := range resData.MeterData {
						byInstanceId[strings.ToLower(meterData.InstanceId)] = append(byInstanceId[strings.ToLower(meterData.InstanceId)], meterData)
					}
				}
			}
		}
	}

	return byInstanceId
}

func CombineAllCostMeters(costDataPath string) (allCostMetersByMeterIdentifier map[string]lib.MongoDbCostItem, allCostMetersSlice []lib.MongoDbCostItem) {
	allCostMetersByMeterIdentifier = make(map[string]lib.MongoDbCostItem)

	paths := lib.GetFullFilePaths(costDataPath)

	for _, path := range paths {
		// fmt.Println(path)
		// continue
		if !strings.Contains(path, "MonthlyCostReport-") {
			continue
		}
		file, err := os.ReadFile(path)
		fileBomRemoved := lib.RemoveJsonByteOrderMark(file)
		lib.CheckFatalError(err)
		var fileData map[string]lib.AggregatedCostItem
		err = json.Unmarshal(fileBomRemoved, &fileData)
		lib.CheckFatalError(err)

		for id, meterData := range fileData {
			jsonStr, _ := json.Marshal(meterData)
			var currData lib.MongoDbCostItem
			// fmt.Println(string(jsonStr))
			// os.Exit(0)
			err := json.Unmarshal(jsonStr, &currData)
			lib.CheckFatalError(err)
			// currData := meterData
			// currData.CostPerDay = nil
			// currData.MonthTotalCost = 0
			allCostMetersByMeterIdentifier[id] = currData
		}
	}

	for _, meterData := range allCostMetersByMeterIdentifier {
		allCostMetersSlice = append(allCostMetersSlice, meterData)
	}

	return allCostMetersByMeterIdentifier, allCostMetersSlice
}

func ProcessCostData(data lib.AggregatedCostData) (costDataByMeterIdentifer map[string]lib.AggregatedCostItem, costDataSlice []lib.AggregatedCostItem) {
	initialCheck := make(map[string][]lib.AggregatedCostItem)
	costDataByMeterIdentifer = make(map[string]lib.AggregatedCostItem)
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)

	fmt.Println("Processing cost data...")
	s.Start()
	for tenantName, tenantData := range data {
		for subName, subData := range tenantData.Subscriptions {
			for _, resGrpData := range subData.ResourceGroups {
				for resName, resData := range resGrpData.Resources {
					_ = resName
					for _, meterData := range resData.MeterData {
						currData := meterData
						for _, v := range currData.CostPerDay {
							currData.MonthTotalCost += v
						}
						currData.TenantName = tenantName
						currData.SubscriptionName = subName
						initialCheck[strings.ToLower(meterData.ResourceMeterIdentifier)] = append(initialCheck[strings.ToLower(meterData.ResourceMeterIdentifier)], currData)
					}
				}
			}
		}
	}

	for _, tenantData := range data {
		for _, subData := range tenantData.Subscriptions {
			for _, resGrpData := range subData.ResourceGroups {
				for resName, resData := range resGrpData.Resources {
					_ = resName
					for id, meterData := range resData.MeterData {
						if len(initialCheck[id]) > 1 {
							err := fmt.Errorf("Incorrect length of meterdata, check data of " + meterData.ResourceMeterIdentifier)
							jsonStr, _ := json.MarshalIndent(initialCheck[id], "", "  ")
							fmt.Println(string(jsonStr))
							lib.CheckFatalError(err)
							os.Exit(1)
						}
						costDataByMeterIdentifer[strings.ToLower(meterData.ResourceMeterIdentifier)] = initialCheck[id][0]
					}
				}
			}
		}
	}
	s.Stop()

	fmt.Println("Creating array from processed data...")
	s.Start()
	for _, meterData := range costDataByMeterIdentifer {
		costDataSlice = append(costDataSlice, meterData)
	}
	s.Stop()

	return costDataByMeterIdentifer, costDataSlice
}

type AzureCostData map[string]AzureTenantCostData

type AzureTenantCostData map[string]AzureCostDataMeter

type AzureCostDataMeter map[string]lib.AggregatedCostItem

func GatherRelatedResourcesAndCostMeters(costData []lib.AggregatedCostItem, resources []lib.AzureResourceDetails, progBarNum int, progBarTotal int) (map[string][]lib.AzureResourceDetails, []lib.AzureResourceDetails) {

	bar := lib.ProgressBar(len(resources), "resource", progBarNum, progBarTotal, "Processing all resources and cost meters to create relations...")

	processedResources := make(map[string][]lib.AzureResourceDetails)

	resById := make(map[string]lib.AzureResourceDetails)
	// vmById := make(map[string]lib.AzureResourceDetails)

	for _, res := range resources {
		resById[res.ID] = res
		// if strings.EqualFold(res.Type, "microsoft.compute/virtualmachines") {
		// 	vmById[res.ID] = res
		// }
	}

	for _, md := range costData {
		meterId := md.ResourceMeterIdentifier
		res := resById[md.InstanceId]
		if !slices.Contains(res.RelatedCostMeters, meterId) {
			res.RelatedCostMeters = append(res.RelatedCostMeters, meterId)
			resById[md.InstanceId] = res
		}
		// if strings.Contains(meterId, res.SubscriptionID) &&
		// 	strings.Contains(meterId, resName) &&
		// 	strings.Contains(meterId, strings.ToLower(res.ResourceGroup)) {
		// 	if !slices.Contains(currRes.RelatedCostMeters, meterId) {
		// 		currRes.RelatedCostMeters = append(currRes.RelatedCostMeters, meterId)
		// 	}
		// }
	}

	// lib.JsonMarshalAndPrint(resById)
	_, _, cachePath := lib.InitConfig(nil)
	jsonStr, err := json.MarshalIndent(resById, "", "  ")
	lib.CheckFatalError(err)
	os.WriteFile(cachePath+"/resById.json", jsonStr, 0644)

	// os.Exit(0)

	for _, res := range resources {
		bar.Add(1)
		// bar.Describe("Processing resource " + strconv.Itoa(i) + " of " + strconv.Itoa(len(resources)))
		resName := strings.ToLower(res.Name)
		// fmt.Println("Processing " + strconv.Itoa(i) + " of " + strconv.Itoa(len(resources)) + " resources")
		var currRes lib.AzureResourceDetails
		currRes.LastAzureSync = time.Now()
		jsonStr, _ := json.Marshal(res)
		err := json.Unmarshal(jsonStr, &currRes)
		lib.CheckFatalError(err)

		currRes.TenantName = strings.ToLower(res.TenantName)

		for _, meterData := range costData {
			meterId := meterData.ResourceMeterIdentifier
			if strings.Contains(meterId, res.SubscriptionID) &&
				strings.Contains(meterId, resName) &&
				strings.Contains(meterId, strings.ToLower(res.ResourceGroup)) {
				if !slices.Contains(currRes.RelatedCostMeters, meterId) {
					currRes.RelatedCostMeters = append(currRes.RelatedCostMeters, meterId)
				}
			}
		}
		for _, comparisonResource := range resources {
			if comparisonResource.ID == res.ID {
				continue
			}
			if strings.ToLower(res.Type) == "microsoft.compute/virtualmachines/extensions" {
				if strings.Contains(strings.ToLower(res.ID), strings.ToLower(comparisonResource.ID)) {
					if !slices.Contains(currRes.RelatedResources, comparisonResource.ID) {
						currRes.RelatedCostMeters = append(currRes.RelatedCostMeters, comparisonResource.ID)
					}
					currRes.RelatedResources = append(currRes.RelatedResources, comparisonResource.ID)
				}
			} else if strings.ToLower(res.Type) == "microsoft.network/networkinterfaces" {
				check := comparisonResource.Properties.VirtualMachine
				if check != nil {
					if strings.EqualFold(res.ID, comparisonResource.Properties.VirtualMachine.ID) {
						if !slices.Contains(currRes.RelatedResources, comparisonResource.ID) {
							currRes.RelatedCostMeters = append(currRes.RelatedCostMeters, comparisonResource.ID)
						}
					}
				}
			} else if strings.ToLower(res.Type) == "microsoft.compute/restorepointcollections" {
				// var rpSource lib.AzureRestorePointCollectionSource

				jsonStr, _ := json.Marshal(res.Properties.Source)
				var rpSource lib.AzureRestorePointCollectionSource
				err := json.Unmarshal(jsonStr, &rpSource)
				if err != nil {
					// lib.JsonMarshalAndPrint(res.Properties)
					rpSource = lib.AzureRestorePointCollectionSource{}
					// lib.CheckFatalError(err)
				}
				// rpSource := res.Properties.Source.(lib.AzureRestorePointCollectionSource)
				// err = json.Unmarshal([]byte(res.Properties.Source), &rpSource)
				// rpSourceProc, ok := rpSource.(lib.AzureRestorePointCollectionSource)
				// rpSourceId := ""

				// if ok {
				// 	rpSourceId = rpSourceProc.ID
				// }

				// if strings.EqualFold(res.ID, comparisonResource.Properties.VirtualMachine.ID) {
				if !slices.Contains(currRes.RelatedResources, rpSource.ID) {
					currRes.RelatedResources = append(currRes.RelatedResources, rpSource.ID)
				}

				// lib.JsonMarshalAndPrint(currRes)

				// os.Exit(0)
				// }
				// resNameSplit := strings.Split(res.Name, "_")
				// if len(resNameSplit) > 1 {
				// 	resName = strings.ToLower(resNameSplit[1])
				// } else {
				// 	fmt.Println("res.Name")
				// 	fmt.Println(res.Name)

				// 	lib.JsonMarshalAndPrint(res)
				// 	os.Exit(0)
				// }
				// if len(resName) != 3 {
				// 	jsonName, _ := json.MarshalIndent(resName, "", "  ")
				// 	fmt.Println(string(jsonName))
				// 	os.Exit(0)
				// }
			} else {
				if strings.Contains(strings.ToLower(comparisonResource.SubscriptionID), strings.ToLower(res.SubscriptionID)) &&
					strings.Contains(strings.ToLower(comparisonResource.Name), resName) &&
					strings.Contains(strings.ToLower(comparisonResource.ResourceGroup), strings.ToLower(res.ResourceGroup)) {
					if !slices.Contains(currRes.RelatedResources, comparisonResource.ID) {
						currRes.RelatedCostMeters = append(currRes.RelatedCostMeters, comparisonResource.ID)
					}
				}
			}
		}

		processedResources[res.TenantName] = append(processedResources[res.TenantName], currRes)
	}

	// jsonStr, _ := json.MarshalIndent(processedResources, "", "  ")
	// os.WriteFile("cost-exports/testRelated-ResourceMeterIdentifier.json", jsonStr, 0644)
	var processedResourcesSlice []lib.AzureResourceDetails
	for _, resSlice := range processedResources {
		processedResourcesSlice = append(processedResourcesSlice, resSlice...)
	}

	return processedResources, processedResourcesSlice
}

func AggregateCostData(data lib.AggregatedCostData, tci lib.TransformedCostItem) lib.AggregatedCostData {

	aggCostData := data

	// dataStr, _ := json.MarshalIndent(data, "", "  ")
	// os.WriteFile("main_data.json", dataStr, 0644)
	// os.Exit(0)

	// tciCostGroup, tciCostGroupExists := tci.Tags["cost_group"]

	tenant, tenantExists := aggCostData[tci.Tenant]

	// fmt.Println(tci.Tenant)
	if !tenantExists {
		tenant.CostPerDay = make(map[string]float64)
		tenant.Subscriptions = make(map[string]lib.AggregatedCostSubscription)
		// tenant.CostGroups = make(map[string]string)
		// jsonStrTenant, _ := json.MarshalIndent(tenant, "", "  ")
		// os.WriteFile("main_jsonStrTenant.json", jsonStrTenant, 0644)
		// jsonStrAggCostData, _ := json.MarshalIndent(aggCostData, "", "  ")
		// os.WriteFile("main_jsonStrAggCostData.json", jsonStrAggCostData, 0644)
		// os.Exit(0)
		aggCostData[tci.Tenant] = tenant
	}

	sub, subExists := aggCostData[tci.Tenant].Subscriptions[tci.SubscriptionName]
	if !subExists {
		if aggCostData[tci.Tenant].Subscriptions == nil {
			t := aggCostData[tci.Tenant]
			t.Subscriptions = make(map[string]lib.AggregatedCostSubscription)
			aggCostData[tci.Tenant] = t
		}
		// s := sub
		sub.CostPerDay = make(map[string]float64)
		aggCostData[tci.Tenant].Subscriptions[tci.SubscriptionName] = sub
	}

	rg, rgExists := aggCostData[tci.Tenant].Subscriptions[tci.SubscriptionName].ResourceGroups[tci.ResourceGroup]
	if aggCostData[tci.Tenant].Subscriptions[tci.SubscriptionName].ResourceGroups == nil {
		s := aggCostData[tci.Tenant].
			Subscriptions[tci.SubscriptionName]

		s.CostPerDay = make(map[string]float64)
		s.ResourceGroups = make(map[string]lib.AggregatedCostResourceGroup)

		aggCostData[tci.Tenant].Subscriptions[tci.SubscriptionName] = s
	}
	if !rgExists {
		aggCostData[tci.Tenant].Subscriptions[tci.SubscriptionName].ResourceGroups[tci.ResourceGroup] = rg
	}

	res, resExists := aggCostData[tci.Tenant].
		Subscriptions[tci.SubscriptionName].
		ResourceGroups[tci.ResourceGroup].
		Resources[tci.ResourceName]

	if aggCostData[tci.Tenant].
		Subscriptions[tci.SubscriptionName].
		ResourceGroups[tci.ResourceGroup].
		Resources == nil {
		rg = aggCostData[tci.Tenant].
			Subscriptions[tci.SubscriptionName].
			ResourceGroups[tci.ResourceGroup]

		rg.CostPerDay = make(map[string]float64)
		rg.Resources = make(map[string]lib.AggregatedCostResource)

		aggCostData[tci.Tenant].
			Subscriptions[tci.SubscriptionName].
			ResourceGroups[tci.ResourceGroup] = rg
	}

	if !resExists {
		aggCostData[tci.Tenant].
			Subscriptions[tci.SubscriptionName].
			ResourceGroups[tci.ResourceGroup].
			Resources[tci.ResourceName] = res
	}

	mdData, mdExists := aggCostData[tci.Tenant].
		Subscriptions[tci.SubscriptionName].
		ResourceGroups[tci.ResourceGroup].
		Resources[tci.ResourceName].
		MeterData[tci.ResourceMeterIdentifier]

	if aggCostData[tci.Tenant].
		Subscriptions[tci.SubscriptionName].
		ResourceGroups[tci.ResourceGroup].
		Resources[tci.ResourceName].
		MeterData == nil {
		res = aggCostData[tci.Tenant].
			Subscriptions[tci.SubscriptionName].
			ResourceGroups[tci.ResourceGroup].
			Resources[tci.ResourceName]

		res.CostPerDay = make(map[string]float64)
		res.MeterData = make(map[string]lib.AggregatedCostItem)

		aggCostData[tci.Tenant].
			Subscriptions[tci.SubscriptionName].
			ResourceGroups[tci.ResourceGroup].
			Resources[tci.ResourceName] = res
	}

	var aci lib.AggregatedCostItem
	tciStr, err := json.Marshal(tci)
	lib.CheckFatalError(err)
	json.Unmarshal(tciStr, &aci)

	if !mdExists {
		aci.CostPerDay = make(map[string]float64)
		aci.UsageQuantityPerDay = make(map[string]float64)
		aggCostData[tci.Tenant].
			Subscriptions[tci.SubscriptionName].
			ResourceGroups[tci.ResourceGroup].
			Resources[tci.ResourceName].
			MeterData[tci.ResourceMeterIdentifier] = aci
	}

	// } else {

	// 	// mdData = aggCostData[tci.Tenant].
	// 	// 	Subscriptions[tci.SubscriptionName].
	// 	// 	ResourceGroups[tci.ResourceGroup].
	// 	// 	Resources[tci.ResourceName].
	// 	// 	MeterData[tci.ResourceMeterIdentifier]

	// 	// mdData.CostPerDay[tci.UsageDateTime] = tci.PreTaxCost
	// 	// mdData.month
	// 	// jsonStr, _ := json.MarshalIndent(mdData, "", "  ")
	// 	// fmt.Println(string(jsonStr))
	// 	// os.Exit(0)
	// }

	tenantData := aggCostData[tci.Tenant]
	tenantData.CostPerDay[tci.UsageDateTime] += tci.PreTaxCost
	tenantData.MonthTotalCost += tci.PreTaxCost

	subData := tenantData.Subscriptions[tci.SubscriptionName]
	subData.CostPerDay[tci.UsageDateTime] += tci.PreTaxCost
	subData.MonthTotalCost += tci.PreTaxCost

	rgData := subData.ResourceGroups[tci.ResourceGroup]
	rgData.CostPerDay[tci.UsageDateTime] += tci.PreTaxCost
	rgData.MonthTotalCost += tci.PreTaxCost

	resData := rgData.Resources[tci.ResourceName]
	resData.CostPerDay[tci.UsageDateTime] += tci.PreTaxCost
	resData.MonthTotalCost += tci.PreTaxCost

	mdData = resData.MeterData[tci.ResourceMeterIdentifier]
	mdData.CostPerDay[tci.UsageDateTime] += tci.PreTaxCost
	mdData.UsageQuantityPerDay[tci.UsageDateTime] += tci.UsageQuantity
	mdData.MonthTotalCost += tci.PreTaxCost

	resData.MeterData[tci.ResourceMeterIdentifier] = mdData
	rgData.Resources[tci.ResourceName] = resData
	subData.ResourceGroups[tci.ResourceGroup] = rgData
	tenantData.Subscriptions[tci.SubscriptionName] = subData

	aggCostData[tci.Tenant] = tenantData

	// jsonStr, _ := json.MarshalIndent(aggCostData, "", "  ")
	// fmt.Println(string(jsonStr))
	// os.Exit(0)

	// fmt.Println(tci.UsageDateTime)

	return aggCostData
}

func TransformCostData(data lib.CostExportData) lib.TransformedCostItemsByTenant {

	cfg := lib.GetCldConfig(nil)
	allData := make(lib.TransformedCostItemsByTenant)

	for _, costData := range data {

		// jsonStr, _ := json.MarshalIndent(costData, "", "  ")
		// fmt.Println(string(jsonStr))

		// os.Exit(0)

		var (
			tagData            map[string]string
			additionalInfoData interface{}
		)

		tenantName := ""
		custTntName := lib.MapAzureSubscriptionToCustomTenantName(costData.SubscriptionGuid, *cfg.Azure)
		if custTntName != "" {
			tenantName = custTntName
		} else {
			tenantName = strings.ToUpper(costData.Datafile)
		}

		// fmt.Println(tenantName)
		// os.Exit(0)

		resourceNameSplit := strings.Split(costData.InstanceId, "/")
		resourceName := resourceNameSplit[len(resourceNameSplit)-1]

		resourceMeterIdentifier := resourceName +
			"_" +
			costData.MeterCategory +
			"_" +
			costData.MeterSubcategory +
			"_" +
			costData.MeterName

		// json.
		tagStr := []byte("{" + costData.Tags + "}")
		json.Unmarshal(tagStr, &tagData)
		json.Unmarshal([]byte(costData.AdditionalInfo), &additionalInfoData)

		// lib.CheckFatalError(err)
		// jsonStr, _ := json.MarshalIndent(tagData, "", "  ")
		// fmt.Println(string(jsonStr))

		tci := lib.TransformedCostItem{
			// ResourceGroup:    costData.ResourceGroup,
			// PreTaxCost:       costData.PreTaxCost,
			// SubscriptionName: costData.SubscriptionName,
			SubscriptionName: costData.SubscriptionName,
			SubscriptionId:   costData.SubscriptionGuid,
			ResourceGroup:    costData.ResourceGroup,
			UsageDateTime:    costData.UsageDateTime,
			ProductName:      costData.ProductName,
			MeterCategory:    costData.MeterCategory,
			MeterSubcategory: costData.MeterSubcategory,
			MeterName:        costData.MeterName,
			UnitOfMeasure:    costData.UnitOfMeasure,
			UsageQuantity:    costData.UsageQuantity,
			ResourceRate:     costData.ResourceRate,
			PreTaxCost:       costData.PreTaxCost,
			// CostCenter:       costData.CostCenter,
			ConsumedService: costData.ConsumedService,
			ResourceType:    costData.ResourceType,
			InstanceId:      costData.InstanceId,
			Tags:            tagData,
			// AdditionalInfo:  costData.AdditionalInfo,
			AdditionalInfo: additionalInfoData,

			ResourceMeterIdentifier: resourceMeterIdentifier,
			ResourceName:            resourceName,
			Tenant:                  tenantName,
			Datafile:                costData.Datafile,
		}

		// jsonStr, _ := json.MarshalIndent(tci, "", "  ")
		// fmt.Println(string(jsonStr))
		// os.Exit(0)

		// _, ok := additionalInfoData["VCPUs"]
		// // If the key exists
		// if ok {
		// 	jsonStr, _ := json.MarshalIndent(tci.AdditionalInfo["VCPUs"], "", "  ")
		// 	fmt.Println(string(jsonStr))
		// }

		// if additionalInfoData["VCPUs"] == "" {
		// 	jsonStr, _ := json.MarshalIndent(tci, "", "  ")
		// 	fmt.Println(string(jsonStr))
		// }

		// allData[costData.Datafile].ResGroups = append(allData[costData.Datafile].ResGroups, tci)
		// transformedTenantData.PreTaxCost += costData.PreTaxCost
		// allData.addPreTaxCost(tci)
		allData.AppendTenantData(tci)

		// allData
		// transformedTenantData.ResGroups = append(transformedTenantData.ResGroups, tci)
		// fmt.Println(tci)
	}
	return allData
}

func GetCostExportCSVFileData(fileName string) (lib.CostExportData, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	_, err = reader.Read()

	var rowData lib.RowData
	var costExport lib.CostExportData
	for {
		err := UnmarshalCostExportCSV(reader, &rowData, fileName)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if rowData.ResourceGroup != "ResourceGroup" {
			costExport = append(costExport, rowData)
		}
	}
	return costExport, nil
}

func GetCostExportJSONFileData(fileName string) (lib.CostExportData, error) {
	var costExport lib.CostExportData
	jsonFile, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(jsonFile, &costExport)
	return costExport, nil
}

func UnmarshalCostExportCSV(reader *csv.Reader, v *lib.RowData, fileName string) error {

	record, err := reader.Read()
	if err != nil {
		return err
	}

	tenant := strings.Split(fileName, "__")[1]
	tenant = strings.Split(tenant, ".")[0]

	s := reflect.ValueOf(v).Elem()
	if s.NumField() != len(record)+1 {
		return &lib.FieldMismatch{s.NumField(), len(record)}
	}

	for i := range s.NumField() {
		f := s.Field(i)
		switch f.Type().String() {
		case "float64":
			fval, err := strconv.ParseFloat(record[i], 64)
			if err != nil {
				return err
			}
			f.SetFloat(fval)
		case "string":
			if i == len(record) {
				f.SetString(tenant)
			} else {
				f.SetString(record[i])
			}
		case "int":
			ival, err := strconv.ParseInt(record[i], 10, 0)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		default:
			fmt.Println(f)
			return &lib.UnsupportedType{f.Type().String()}
		}
	}

	return nil
}

func ConvertCsvFileToExcel(sheetName string, inFileName string, outFileName string) {
	// now := time.Now()
	// .Format("200601")
	// year := strconv.Itoa(now.Year())
	// month := fmt.Sprintf("%02d", int(now.Month()))
	// fileName := "MonthlyReport-" + year + month + ".xlsx"

	csvFile, err := os.Open(inFileName)
	lib.CheckFatalError(err)
	reader := csv.NewReader(csvFile)
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	f.NewSheet(sheetName)

	err = f.DeleteSheet("Sheet1")
	lib.CheckFatalError(err)

	row := 1

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		lib.CheckFatalError(err)

		cell, err := excelize.CoordinatesToCellName(1, row)
		lib.CheckFatalError(err)

		err = f.SetSheetRow(sheetName, cell, &record)
		lib.CheckFatalError(err)

		row++
	}

	err = f.SaveAs(outFileName)
	lib.CheckFatalError(err)
}

func ConvertCsvDataToExcel(sheetName string, inFileName string, outFileName string) {

	csvFile, err := os.Open(inFileName)
	lib.CheckFatalError(err)
	reader := csv.NewReader(csvFile)
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	f.NewSheet(sheetName)

	err = f.DeleteSheet("Sheet1")
	lib.CheckFatalError(err)

	row := 1

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		lib.CheckFatalError(err)

		cell, err := excelize.CoordinatesToCellName(1, row)
		lib.CheckFatalError(err)

		err = f.SetSheetRow(sheetName, cell, &record)
		lib.CheckFatalError(err)

		row++
	}

	err = f.SaveAs(outFileName)
	lib.CheckFatalError(err)
}

// func (data *AggregatedCostData) SumCosts() {
// 	costData := *data

// 	for key, val := range costData {
// 		_ = key
// 		_ = val
// 		// fmt.Println(val)

// 		for mdkey, mdval := range val.MeterData {
// 			jsonStr, _ := json.MarshalIndent(mdval, "", "  ")
// 			fmt.Println(string(jsonStr))
// 			os.Exit(0)
// 		}
// 	}
// }

func DownloadAllConfiguredTenantCostExportsForMonth(opts lib.DownloadAllConfiguredTenantCostExportsForMonthOptions, cldConfOpts *lib.CldConfigOptions) {
	var wg sync.WaitGroup

	if _, err := os.Stat(opts.OutfilePath); err != nil {
		os.MkdirAll(opts.OutfilePath, os.ModePerm)
	}

	config := lib.GetCldConfig(cldConfOpts)

	for _, tenant := range config.Azure.MultiTenantAuth.Tenants {
		if tenant.CostExportsLocation == "" {
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			split := strings.Split(tenant.CostExportsLocation, "/")
			if len(split) != 4 {
				lib.CheckFatalError(fmt.Errorf("CostExportsLocation does not seem to be correctly formatted"))
			}

			options := lib.StorageAccountRequestOptions{
				StorageAccountName:   strings.Split(split[2], ".")[0],
				ContainerName:        split[3],
				ConfiguredTenantName: tenant.TenantName,
			}

			blobList := ListStorageContainerBlobs(options, cldConfOpts)
			blobList.Filter(lib.BlobListFilterOptions{FilterPrefix: opts.BlobPrefix})
			blobList.SortByCreateDate("descending")

			cred, err := GetTenantAzCred(tenant.TenantName, false, cldConfOpts)
			lib.CheckFatalError(err)
			fileName := opts.OutfilePath + "/" + opts.OutfileNamePrefix + "__" + tenant.TenantName + ".csv"

			if !opts.SuppressSteps {
				fmt.Println(tenant.TenantName, len(blobList))
			}
			if len(blobList) > 0 {
				blobList[0].Download(cred, fileName)
				if !opts.SuppressSteps {
					fmt.Println("Downloaded", tenant.TenantName, "to", fileName)
				}
			} else {
				if !opts.SuppressSteps {
					fmt.Println("No blobs found for ", tenant.TenantName)
				}
			}

		}()
	}
	wg.Wait()
}

// Generates monthly cost report from Cost Data stored in Azure Blob Storages, exported from Azure
//
// outputDirectory = Base directory to store downloaded data and generated Excel workbook.
//
// Full save directory would be outputDirectory + "/cost-exports/" + lastMonth.Format("2006") + lastMonth.Format("01") + "/"
//
// blobPrefix is the directory within the Blob Container to find files.
//
// Full blob prefix becomes blobPrefix + "/" + lastMonth.Format("2006") + lastMonth.Format("01")
func GenerateMonthlyCostReport(outputDirectory string, outputFileName string, blobPrefix string, costExportMonth string) {
	if !lib.CheckDirExists(outputDirectory) {
		os.MkdirAll(outputDirectory, os.ModePerm)
	}

	// cldConfOpts := &lib.CldConfigOptions{
	// 	ConfigFilePath: configFilePath,
	// }

	DownloadAllConfiguredTenantCostExportsForMonth(lib.DownloadAllConfiguredTenantCostExportsForMonthOptions{
		BlobPrefix:        blobPrefix + "/" + costExportMonth,
		OutfilePath:       outputDirectory + "/" + costExportMonth,
		OutfileNamePrefix: "cost-export",
		CostExportMonth:   costExportMonth,
		SuppressSteps:     false,
	}, nil)

	combinedCostData := CombineCostExportCSVData(outputDirectory + "/" + costExportMonth)

	// fmt.Println(combinedCostData)
	transformedData := TransformCostData(combinedCostData)
	// fmt.Println(transformedData)
	// os.Exit(0)

	CostDataToExcel(transformedData, outputFileName)
}

func SortCostPerDay(costPerDay lib.CostPerDay) lib.CostPerDay {
	keys := make([]string, 0, len(costPerDay))

	sorted := make(map[string]float64)

	for k := range costPerDay {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		// fmt.Println(k, costPerDay[k])
		sorted[k] = costPerDay[k]
	}

	return sorted
}

//
//

func CombineCostDataSlices(basePath string, saveOutput bool) (combinedSlice []lib.MongoDbCostItem) {
	byResourceMeterIdentifier := make(map[string]lib.MongoDbCostItem)

	paths := lib.GetFullFilePaths(basePath)

	for _, path := range paths {
		if strings.Contains(path, "costData-CombinedSlice") {
			continue
		}
		var currentFileData []lib.MongoDbCostItem
		file, err := os.ReadFile(path)
		lib.CheckFatalError(err)
		err = json.Unmarshal(file, &currentFileData)
		lib.CheckFatalError(err)
		for _, cdi := range currentFileData {
			byResourceMeterIdentifier[cdi.ResourceMeterIdentifier] = cdi
		}
	}

	for _, cdi := range byResourceMeterIdentifier {
		combinedSlice = append(combinedSlice, cdi)
	}

	if saveOutput {
		jsonStr, err := json.MarshalIndent(combinedSlice, "", "  ")
		lib.CheckFatalError(err)

		err = os.WriteFile(basePath+"costData-CombinedSlice.json", jsonStr, 0644)
		lib.CheckFatalError(err)
		fmt.Println("Combined data saved to: " + basePath + "costData-CombinedSlice.json")
	}

	// fmt.Println("len(combinedSlice)")
	// fmt.Println(len(combinedSlice))

	return
}

//
//
