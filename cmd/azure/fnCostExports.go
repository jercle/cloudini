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
	"strconv"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/jercle/azg/lib"
	"github.com/xuri/excelize/v2"
)

type rowData struct {
	DepartmentName   string  `csv:"-"`
	AccountName      string  `csv:"-"`
	AccountOwnerId   string  `csv:"-"`
	SubscriptionGuid string  `csv:"SubscriptionGuid" json:"SubscriptionGuid"`
	SubscriptionName string  `csv:"SubscriptionName" json:"SubscriptionName"`
	ResourceGroup    string  `csv:"ResourceGroup" json:"ResourceGroup"`
	ResourceLocation string  `csv:"-"`
	AvailabilityZone string  `csv:"-"`
	UsageDateTime    string  `csv:"UsageDateTime" json:"UsageDateTime"`
	ProductName      string  `csv:"ProductName" json:"ProductName"`
	MeterCategory    string  `csv:"-"`
	MeterSubcategory string  `csv:"-"`
	MeterId          string  `csv:"-"`
	MeterName        string  `csv:"MeterName" json:"MeterName"`
	MeterRegion      string  `csv:"-"`
	UnitOfMeasure    string  `csv:"UnitOfMeasure" json:"UnitOfMeasure"`
	UsageQuantity    float64 `csv:"UsageQuantity" json:"UsageQuantity"`
	ResourceRate     float64 `csv:"ResourceRate" json:"ResourceRate"`
	PreTaxCost       float64 `csv:"PreTaxCost" json:"PreTaxCost"`
	CostCenter       string  `csv:"-"`
	ConsumedService  string  `csv:"ConsumedService" json:"ConsumedService"`
	ResourceType     string  `csv:"ResourceType" json:"ResourceType"`
	InstanceId       string  `csv:"InstanceId" json:"InstanceId"`
	Tags             string  `csv:"-"`
	OfferId          string  `csv:"-"`
	AdditionalInfo   string  `csv:"-"`
	ServiceInfo1     string  `csv:"-"`
	ServiceInfo2     string  `csv:"-"`
	Currency         string  `csv:"Currency" json:"Currency"`
	Datafile         string
}

type costExportData []rowData

type FieldMismatch struct {
	expected, found int
}

type UnsupportedType struct {
	Type string
}

type TransformedCostItem struct {
	ResourceGroup    string
	PreTaxCost       float64
	SubscriptionName string
	Tenant           string
	Datafile         string
}

type TransformedTenantData struct {
	PreTaxCost float64
	ResGroups  []TransformedCostItem
}

type TransformedCostItemsByTenantMap map[string]TransformedTenantData

// type TransformedCostItemsByTenant map[string]TransformedTenantData
type TransformedCostItemsByTenant struct {
	Blue      TransformedTenantData
	BlueDTQ   TransformedTenantData
	Red       TransformedTenantData
	RedDTQ    TransformedTenantData
	Yellow    TransformedTenantData
	PUD       TransformedTenantData
	PUDDTQ    TransformedTenantData
	Purple    TransformedTenantData
	PurpleDTQ TransformedTenantData
}

type CostQueryResponse struct {
	ID         string `json:"id"`
	Properties struct {
		NextLink string `json:"nextLink"`
		Columns  []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"columns"`
		Rows [][]interface{} `json:"rows"`
	} `json:"properties"`
}

type Post struct {
	DataSet struct {
		Aggregation struct {
			TotalCost struct {
				Function string
				Name     string
			}
		}
		Granularity string
		Grouping    []struct {
			Name string
			Type string
		}
		Sorting []struct {
			Direction string
			Name      string
		}
	}
	TimePeriod struct {
		From string
		To   string
	}
	Timeframe string
	Type      string
}

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
		// fmt.Println(tenant)
		// fmt.Println(tData)
		_, err := f.NewSheet(tenant)
		err = f.SetSheetRow(tenant, "A1", &[]string{
			"Subscription",
			"Resource Group",
			"PreTaxCost",
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

			// err = f.SetCellValue(tenant, cell, rowData.ResourceGroup)
			lib.CheckFatalError(err)
			row++
		}

		// TODO: Add pivot table
		// if err := f.AddPivotTable(&excelize.PivotTableOptions{
		// 	DataRange:       tenant + "!A1:C" + strconv.Itoa(row),
		// 	PivotTableRange: tenant + "!G2:M34",
		// 	// Rows: []excelize.PivotTableField{
		// 	// 	{Data: "Subscription", Subtotal: "Sum"}, {Data: "PreTaxCost"}},
		// 	// Columns: []excelize.PivotTableField{
		// 	// 	{Data: "Subscription", DefaultSubtotal: true}},
		// 	// Data: []excelize.PivotTableField{
		// 	// 	{Data: "PreTaxCost", Name: "Summarize", Subtotal: "Sum"}},
		// 	RowGrandTotals: true,
		// 	ColGrandTotals: true,
		// 	ShowDrill:      true,
		// 	ShowRowHeaders: true,
		// 	ShowColHeaders: true,
		// 	ShowLastColumn: true,
		// }); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

	}

	err := f.DeleteSheet("Sheet1")
	lib.CheckFatalError(err)
	err = f.SaveAs(outFileName)
	lib.CheckFatalError(err)
}

func CombineCostExportCSVData(dataPath string) costExportData {
	var (
		wg             sync.WaitGroup
		costExportData costExportData
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

func CombineCostExportJSONData(dataPath string) costExportData {
	var (
		wg             sync.WaitGroup
		costExportData costExportData
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

func TransformCostData(data costExportData) TransformedCostItemsByTenantMap {
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
		var tenantName string

		// let lcSubName = SubscriptionName.toLowerCase()
		sn := strings.ToLower(costData.SubscriptionName)

		switch {
		case sn == "pud":
			tenantName = "PUD"
		case sn == "puddtq":
			tenantName = "PUDDTQ"
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

		tci := TransformedCostItem{
			ResourceGroup:    costData.ResourceGroup,
			PreTaxCost:       costData.PreTaxCost,
			SubscriptionName: costData.SubscriptionName,
			Tenant:           tenantName,
			Datafile:         costData.Datafile,
		}

		// allData[costData.Datafile].ResGroups = append(allData[costData.Datafile].ResGroups, tci)
		// transformedTenantData.PreTaxCost += costData.PreTaxCost
		// allData.addPreTaxCost(tci)
		allData.appendTenantData(tci)
		// allData
		// transformedTenantData.ResGroups = append(transformedTenantData.ResGroups, tci)
		// fmt.Println(tci)
		// os.Exit(0)
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

func GetCostExportCSVFileData(fileName string) (costExportData, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	_, err = reader.Read()

	var rowData rowData
	var costExport costExportData
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

func GetCostExportJSONFileData(fileName string) (costExportData, error) {
	var costExport costExportData
	jsonFile, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(jsonFile, &costExport)
	return costExport, nil
}

func UnmarshalCostExportCSV(reader *csv.Reader, v *rowData, fileName string) error {

	record, err := reader.Read()
	if err != nil {
		return err
	}

	tenant := strings.Split(fileName, "_")[1]
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

func (tenants *TransformedCostItemsByTenantMap) addPreTaxCost(tci TransformedCostItem) {
	t := *tenants

	thing := t[tci.Datafile]
	thing.PreTaxCost += tci.PreTaxCost

	fmt.Println(t)
}

func (t *TransformedCostItemsByTenantMap) appendTenantData(tci TransformedCostItem) {
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

func (e *FieldMismatch) Error() string {
	return "CSV line fields mismatch. Expected " + strconv.Itoa(e.expected) + " found " + strconv.Itoa(e.found)
}
