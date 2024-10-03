package azure

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/jercle/cloudini/lib"
	"github.com/xuri/excelize/v2"
)

func getCostData(cred *azidentity.DefaultAzureCredential, subscriptionId string, resourceGroup string) {
	ctx := context.Background()
	tokenRequestOptions := policy.TokenRequestOptions{
		Scopes: []string{
			"https://management.core.windows.net/.default",
		},
	}

	token, err := cred.GetToken(ctx, tokenRequestOptions)
	if err != nil {
		log.Fatal(err)
	}

	urlString := "https://management.azure.com/subscriptions/" + subscriptionId + "/resourceGroups/" + resourceGroup + "/providers/Microsoft.CostManagement/query?api-version=2023-11-01"

	body := []byte(`{
			"type": "ActualCost",
			"dataSet": {
					"granularity": "Daily",
					"aggregation": {
							"totalCost": {
									"name": "Cost",
									"function": "Sum"
							}
					},
					"sorting": [
							{
									"direction": "ascending",
									"name": "UsageDate"
							}
					],
					"grouping": [
							{
									"type": "Dimension",
									"name": "ResourceId"
							},
							{
									"type": "Dimension",
									"name": "ChargeType"
							},
							{
									"type": "Dimension",
									"name": "PublisherType"
							}
					]
			},
			"timeframe": "Custom",
			"timePeriod": {
					"from": "2024-01-16T00:00:00+00:00",
					"to": "2024-01-21T23:59:59+00:00"
			}
	}`)
	// r, err := http.Post(urlString, "application/json", bytes.NewBuffer(body))
	// http.Post does not allow for custom headers
	// if err != nil {
	// 	panic(err)
	// }

	r, err := http.NewRequest("POST", urlString, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+token.Token)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	// post := &Post{}
	response := &CostQueryResponse{}
	derr := json.NewDecoder(res.Body).Decode(response)
	if derr != nil {
		panic(derr)
	}

	fmt.Println(res.StatusCode)
	// if res.StatusCode != http.StatusOK {
	// panic(derr)
	// }

	fmt.Println(response)
}

func CostDataToExcel(data TransformedCostItemsByTenantMap, outFileName string) {
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

func CombineCostExportCSVData(dataPath string) CostExportData {
	var (
		wg             sync.WaitGroup
		costExportData CostExportData
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
			go func() {
				defer wg.Done()
				// fmt.Println(file)
				data, err := GetCostExportCSVFileData(file)
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

func CombineCostExportJSONData(dataPath string) CostExportData {
	var (
		wg             sync.WaitGroup
		costExportData CostExportData
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

func TransformCostDataNew(data CostExportData) AggregatedCostData {
	// func TransformCostDataNew(data CostExportData) AggregatedCostData {
	// fmt.Println(len(*data))

	var (
	// transformedTenantData transformedTenantData
	// allDataNEW AggregatedCostData
	// allSubscriptions []string
	)
	// allData := TransformedCostItemsByTenantMap{}
	allDataNEW := AggregatedCostData{}
	// allData
	// allSubscriptions = lib.UniqueNonEmptyElementsOf(allSubscriptions)
	// fmt.Println(allSubscriptions)

	// os.Exit(0)

	counter := 1

	for _, costData := range data {
		// jsonStr, _ := json.MarshalIndent(data, "", "  ")
		// fmt.Println(string(jsonStr))
		// fmt.Println(costData.UsageDateTime)
		// os.Exit(0)
		// fmt.Println(costData.Datafile)
		// fmt.Println(costData.PreTaxCost)
		var (
			tenantName         string
			tagData            map[string]string
			additionalInfoData interface{}
			// additionalInfoData map[string]string
		)

		// let lcSubName = SubscriptionName.toLowerCase()
		sn := strings.ToLower(costData.SubscriptionName)
		instanceId := strings.ToLower(costData.InstanceId)

		switch {
		case sn == "pud":
			tenantName = "GREEN"
		case sn == "puddtq":
			tenantName = "GREENDTQ"
		case strings.ToLower(costData.Datafile) == "yellow":
			tenantName = "YELLOW"
		case strings.Contains(sn, "devdtq") && costData.Datafile != "BLUE" && costData.Datafile != "BLUEDTQ":
			tenantName = "PURPLEDTQ"
		case strings.Contains(sn, "dev") && costData.Datafile != "YELLOW":
			tenantName = "PURPLE"
		case strings.Contains(sn, "apcdtq"):
			tenantName = "REDDTQ"
		case strings.Contains(sn, "apc") && costData.Datafile == "REDDTQ":
			tenantName = "REDDTQ"
		case strings.Contains(sn, "apc") && costData.Datafile != "REDDTQ":
			tenantName = "RED"
		case costData.Datafile == "BLUEDTQ":
			tenantName = "BLUEDTQ"
		case costData.Datafile == "BLUE":
			tenantName = "BLUE"
		case costData.Datafile == "PUD":
			tenantName = "PUD"
		default:
			tenantName = strings.ToUpper(costData.Datafile)
		}

		resourceNameSplit := strings.Split(instanceId, "/")
		resourceName := resourceNameSplit[len(resourceNameSplit)-1]

		resourceMeterIdentifier := resourceName +
			"_" +
			strings.ToLower(costData.MeterCategory) +
			"_" +
			strings.ToLower(costData.MeterSubcategory) +
			"_" +
			strings.ToLower(costData.MeterName)

		// json.
		tagStr := []byte("{" + costData.Tags + "}")
		json.Unmarshal(tagStr, &tagData)
		json.Unmarshal([]byte(costData.AdditionalInfo), &additionalInfoData)

		// lib.CheckFatalError(err)
		// jsonStr, _ := json.MarshalIndent(tagData, "", "  ")
		// fmt.Println(string(jsonStr))

		tci := TransformedCostItem{
			// ResourceGroup:    costData.ResourceGroup,
			// PreTaxCost:       costData.PreTaxCost,
			// SubscriptionName: costData.SubscriptionName,
			SubscriptionName: strings.ToLower(costData.SubscriptionName),
			ResourceGroup:    strings.ToLower(costData.ResourceGroup),
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
			Tenant:                  strings.ToLower(tenantName),
			Datafile:                costData.Datafile,
		}
		// _ = tci

		// allDataNEW.AppendTenantData(tci)
		// fmt.Println(tci.UsageDateTime)
		allDataNEW = AggregateCostData(allDataNEW, tci)
		counter++

		// if counter > 200 {
		// 	// fmt.Println(allDataNEW)
		// 	// fmt.Println()
		// 	jsonStr, _ := json.MarshalIndent(allDataNEW, "", "  ")
		// 	fmt.Println(string(jsonStr))
		// 	os.Exit(0)
		// }
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
		// allData.AppendTenantData(tci)

		// allData
		// transformedTenantData.ResGroups = append(transformedTenantData.ResGroups, tci)
		// fmt.Println(tci)
	}
	return allDataNEW
}

func AggregateCostData(data AggregatedCostData, tci TransformedCostItem) AggregatedCostData {
	aggCostData := data

	tciCostGroup, tciCostGroupExists := tci.Tags["cost_group"]

	tenant, tenantExists := aggCostData[tci.Tenant]
	if !tenantExists {
		tenant.CostPerDay = make(map[string]float64)
		tenant.Subscriptions = make(map[string]AggregatedCostSubscription)
		if tciCostGroupExists {
			tenant.CostGroups[tciCostGroup] = make(map[string]float64)
		}
		aggCostData[tci.Tenant] = tenant

	}

	sub, subExists := aggCostData[tci.Tenant].Subscriptions[tci.SubscriptionName]
	if !subExists {
		if aggCostData[tci.Tenant].Subscriptions == nil {
			t := aggCostData[tci.Tenant]
			t.Subscriptions = make(map[string]AggregatedCostSubscription)
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
		s.ResourceGroups = make(map[string]AggregatedCostResourceGroup)

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
		rg.Resources = make(map[string]AggregatedCostResource)

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
		res.MeterData = make(map[string]AggregatedCostItem)

		aggCostData[tci.Tenant].
			Subscriptions[tci.SubscriptionName].
			ResourceGroups[tci.ResourceGroup].
			Resources[tci.ResourceName] = res
	}

	var aci AggregatedCostItem
	tciStr, err := json.Marshal(tci)
	lib.CheckFatalError(err)
	json.Unmarshal(tciStr, &aci)

	if !mdExists {
		aci.CostPerDay = make(map[string]float64)
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

func TransformCostData(data CostExportData) TransformedCostItemsByTenantMap {
	// fmt.Println(len(*data))

	var (
		// transformedTenantData transformedTenantData
		// allData
		allSubscriptions []string
	)
	allData := TransformedCostItemsByTenantMap{}
	allSubscriptions = lib.UniqueNonEmptyElementsOf(allSubscriptions)
	// type transformedCostItem struct {
	// 	ResourceGroup    string
	// 	PreTaxCost       float64
	// 	SubscriptionName string
	// 	Tenant           string
	// }
	for _, costData := range data {
		// fmt.Println(costData.Datafile)
		// fmt.Println(costData.PreTaxCost)
		var (
			tenantName         string
			tagData            map[string]string
			additionalInfoData interface{}
			// additionalInfoData map[string]string
		)

		// let lcSubName = SubscriptionName.toLowerCase()
		sn := strings.ToLower(costData.SubscriptionName)

		switch {
		case sn == "pud":
			tenantName = "GREEN"
		case sn == "puddtq":
			tenantName = "GREENDTQ"
		case strings.ToLower(costData.Datafile) == "yellow":
			tenantName = "YELLOW"
		case strings.Contains(sn, "devdtq") && costData.Datafile != "BLUE" && costData.Datafile != "BLUEDTQ":
			tenantName = "PURPLEDTQ"
		case strings.Contains(sn, "dev") && costData.Datafile != "YELLOW":
			tenantName = "PURPLE"
		case strings.Contains(sn, "apcdtq"):
			tenantName = "REDDTQ"
		case strings.Contains(sn, "apc") && costData.Datafile == "REDDTQ":
			tenantName = "REDDTQ"
		case strings.Contains(sn, "apc") && costData.Datafile != "REDDTQ":
			tenantName = "RED"
		case costData.Datafile == "BLUEDTQ":
			tenantName = "BLUEDTQ"
		case costData.Datafile == "BLUE":
			tenantName = "BLUE"
		case costData.Datafile == "PUD":
			tenantName = "PUD"
		default:
			tenantName = strings.ToUpper(costData.Datafile)
		}

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

		tci := TransformedCostItem{
			// ResourceGroup:    costData.ResourceGroup,
			// PreTaxCost:       costData.PreTaxCost,
			// SubscriptionName: costData.SubscriptionName,
			SubscriptionName: costData.SubscriptionName,
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

func (tenants *TransformedCostItemsByTenant) AddPreTaxCost(tci TransformedCostItem) {
	switch tci.Datafile {
	case "Blue":
		tenants.Blue.PreTaxCost += tci.PreTaxCost
	case "BlueDTQ":
		tenants.BlueDTQ.PreTaxCost += tci.PreTaxCost
	case "Red":
		tenants.Red.PreTaxCost += tci.PreTaxCost
	case "RedDTQ":
		tenants.RedDTQ.PreTaxCost += tci.PreTaxCost
	case "Yellow":
		tenants.Yellow.PreTaxCost += tci.PreTaxCost
	case "PUD":
		tenants.PUD.PreTaxCost += tci.PreTaxCost
	case "PUDDTQ":
		tenants.PUDDTQ.PreTaxCost += tci.PreTaxCost
	case "Purple":
		tenants.Purple.PreTaxCost += tci.PreTaxCost
	case "PurpleDTQ":
		tenants.PurpleDTQ.PreTaxCost += tci.PreTaxCost
	}
}

func (tenants *TransformedCostItemsByTenant) AppendTenantData(tci TransformedCostItem) {
	switch tci.Tenant {
	case "BLUE":
		tenants.Blue.ResGroups = append(tenants.Blue.ResGroups, tci)
		tenants.Blue.PreTaxCost += tci.PreTaxCost
	case "BLUEDTQ":
		tenants.BlueDTQ.ResGroups = append(tenants.BlueDTQ.ResGroups, tci)
		tenants.BlueDTQ.PreTaxCost += tci.PreTaxCost
	case "RED":
		tenants.Red.ResGroups = append(tenants.Red.ResGroups, tci)
		tenants.Red.PreTaxCost += tci.PreTaxCost
	case "REDDTQ":
		tenants.RedDTQ.ResGroups = append(tenants.RedDTQ.ResGroups, tci)
		tenants.RedDTQ.PreTaxCost += tci.PreTaxCost
	case "YELLOW":
		tenants.Yellow.ResGroups = append(tenants.Yellow.ResGroups, tci)
		tenants.Yellow.PreTaxCost += tci.PreTaxCost
	case "PUD":
		tenants.PUD.ResGroups = append(tenants.PUD.ResGroups, tci)
		tenants.PUD.PreTaxCost += tci.PreTaxCost
	case "PUDDTQ":
		tenants.PUDDTQ.ResGroups = append(tenants.PUDDTQ.ResGroups, tci)
		tenants.PUDDTQ.PreTaxCost += tci.PreTaxCost
	case "PURPLE":
		tenants.Purple.ResGroups = append(tenants.Purple.ResGroups, tci)
		tenants.Purple.PreTaxCost += tci.PreTaxCost
	case "PURPLEDTQ":
		tenants.PurpleDTQ.ResGroups = append(tenants.PurpleDTQ.ResGroups, tci)
		tenants.PurpleDTQ.PreTaxCost += tci.PreTaxCost
	}
}

func GetCostExportCSVFileData(fileName string) (CostExportData, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	_, err = reader.Read()

	var rowData RowData
	var costExport CostExportData
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

func GetCostExportJSONFileData(fileName string) (CostExportData, error) {
	var costExport CostExportData
	jsonFile, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(jsonFile, &costExport)
	return costExport, nil
}

func UnmarshalCostExportCSV(reader *csv.Reader, v *RowData, fileName string) error {

	record, err := reader.Read()
	if err != nil {
		return err
	}

	tenant := strings.Split(fileName, "__")[1]
	tenant = strings.Split(tenant, ".")[0]

	s := reflect.ValueOf(v).Elem()
	if s.NumField() != len(record)+1 {
		return &FieldMismatch{s.NumField(), len(record)}
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
			return &UnsupportedType{f.Type().String()}
		}
	}

	return nil
}

func (e *UnsupportedType) Error() string {
	return "Unsupported type: " + e.Type
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

func (tenants *TransformedCostItemsByTenantMap) AddPreTaxCost(tci TransformedCostItem) {
	t := *tenants

	thing := t[tci.Datafile]
	thing.PreTaxCost += tci.PreTaxCost

	fmt.Println(t)
}

func (t *TransformedCostItemsByTenantMap) AppendTenantData(tci TransformedCostItem) {
	tenants := *t
	entry, exists := tenants[tci.Tenant]
	if !exists {
		tenant := TransformedTenantData{}
		tenant.PreTaxCost += tci.PreTaxCost
		tenant.ResGroups = append(tenant.ResGroups, tci)
		tenants[tci.Tenant] = tenant
	} else {
		tenant := entry
		tenant.PreTaxCost += tci.PreTaxCost
		tenant.ResGroups = append(tenant.ResGroups, tci)
		tenants[tci.Tenant] = tenant
	}
	*t = tenants
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

func (t *TransformedCostItemsByTenantMap) SumCosts() {
	tenants := *t

	for key, val := range tenants {
		_ = key
		_ = val
		// fmt.Println(val)

		for _, res := range val.ResGroups {
			jsonStr, _ := json.MarshalIndent(res, "", "  ")
			fmt.Println(string(jsonStr))
			os.Exit(0)
		}
	}
}

func (e *FieldMismatch) Error() string {
	return "CSV line fields mismatch. Expected " + strconv.Itoa(e.expected) + " found " + strconv.Itoa(e.found)
}

func DownloadAllConfiguredTenantLastMonthCostExports(opts DownloadAllConfiguredTenantLastMonthCostExportsOptions, cldConfOpts *lib.CldConfigOptions) {
	var wg sync.WaitGroup

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

			options := StorageAccountRequestOptions{
				StorageAccountName:   strings.Split(split[2], ".")[0],
				ContainerName:        split[3],
				ConfiguredTenantName: tenant.TenantName,
			}

			blobList := ListStorageContainerBlobs(options, cldConfOpts)
			blobList.Filter(BlobListFilterOptions{FilterPrefix: opts.BlobPrefix})
			blobList.SortByCreateDate("descending")

			cred := GetTenantAzCred(tenant.TenantName, false, cldConfOpts)
			fileName := opts.OutfilePath + "__" + tenant.TenantName + ".csv"
			fmt.Println(tenant.TenantName, len(blobList))
			if len(blobList) > 0 {
				blobList[0].Download(cred, fileName)
				fmt.Println("Downloaded", tenant.TenantName, "to", fileName)
			} else {
				fmt.Println("No blobs found for ", tenant.TenantName)
			}

		}()
	}
	wg.Wait()
}

func ListStorageContainerBlobs(options StorageAccountRequestOptions, cldConfigOpts *lib.CldConfigOptions) BlobList {
	var (
		cred     *azidentity.ClientSecretCredential
		err      error
		ctx      = context.Background()
		BlobList BlobList
	)

	config := lib.GetCldConfig(cldConfigOpts)
	tenant := config.Azure.MultiTenantAuth.Tenants[options.ConfiguredTenantName]
	lib.PrintSrcLoc(tenant.TenantName)

	if options.GetWriteToken {
		cred, err = azidentity.NewClientSecretCredential(tenant.TenantID, tenant.Writer.ClientID, tenant.Writer.ClientSecret, nil)
		lib.CheckFatalError(err)
	} else {
		cred, err = azidentity.NewClientSecretCredential(tenant.TenantID, tenant.Reader.ClientID, tenant.Reader.ClientSecret, nil)
		lib.CheckFatalError(err)
	}

	serviceURL := "https://" + options.StorageAccountName + ".blob.core.windows.net"
	client, err := azblob.NewClient(serviceURL, cred, nil)

	pager := client.NewListBlobsFlatPager(options.ContainerName, &azblob.ListBlobsFlatOptions{
		Include: container.ListBlobsInclude{Deleted: false, Versions: false},
	})

	for pager.More() {
		resp, err := pager.NextPage(ctx)
		lib.CheckFatalError(err)
		for _, blob := range resp.Segment.BlobItems {
			var blobItem BlobItem
			jsonBytes, _ := json.MarshalIndent(blob, "", "  ")
			blobItem.TenantName = tenant.TenantName
			blobItem.StorageAccountName = options.StorageAccountName
			blobItem.ContainerName = options.ContainerName
			json.Unmarshal(jsonBytes, &blobItem)
			// fmt.Println(blobItem)
			BlobList = append(BlobList, blobItem)
		}
	}

	return BlobList
}

func (blob *BlobItem) Download(cred *azidentity.ClientSecretCredential, fileName string) {
	var (
		ctx = context.Background()
	)
	serviceURL := "https://" + blob.StorageAccountName + ".blob.core.windows.net"
	client, err := azblob.NewClient(serviceURL, cred, nil)
	lib.CheckFatalError(err)
	file, err := os.Create(fileName)
	lib.CheckFatalError(err)
	defer file.Close()
	_, err = client.DownloadFile(ctx, blob.ContainerName, blob.Name, file, nil)
}

func (bl *BlobList) Filter(opts BlobListFilterOptions) {
	var filteredBlobs []BlobItem
	for _, blob := range *bl {
		if strings.HasPrefix(blob.Name, opts.FilterPrefix) {
			filteredBlobs = append(filteredBlobs, blob)
		}
	}
	*bl = filteredBlobs
}

func (blobList *BlobList) SortByCreateDate(sortOrder string) {
	bl := *blobList
	if sortOrder == "ascending" {
		sort.Slice(bl, func(i, j int) bool {
			return bl[i].Properties.CreationTime.Before(bl[j].Properties.CreationTime)
		})
	} else if sortOrder == "descending" {
		sort.Slice(bl, func(i, j int) bool {
			return bl[j].Properties.CreationTime.Before(bl[i].Properties.CreationTime)
		})
	} else {
		fmt.Println("Sort order must be 'ascending' or 'descending'")
	}

	*blobList = bl
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

	DownloadAllConfiguredTenantLastMonthCostExports(DownloadAllConfiguredTenantLastMonthCostExportsOptions{
		BlobPrefix:  blobPrefix + "/" + costExportMonth,
		OutfilePath: outputDirectory + "cost-export",
	}, nil)

	combinedCostData := CombineCostExportCSVData(outputDirectory)
	transformedData := TransformCostData(combinedCostData)

	CostDataToExcel(transformedData, outputFileName)
}
