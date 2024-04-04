package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/jercle/azg/lib"
)

type rowData struct {
	DepartmentName   string  `csv:"-"`
	AccountName      string  `csv:"-"`
	AccountOwnerId   string  `csv:"-"`
	SubscriptionGuid string  `csv:"SubscriptionGuid"`
	SubscriptionName string  `csv:"SubscriptionName"`
	ResourceGroup    string  `csv:"ResourceGroup"`
	ResourceLocation string  `csv:"-"`
	AvailabilityZone string  `csv:"-"`
	UsageDateTime    string  `csv:"UsageDateTime"`
	ProductName      string  `csv:"ProductName"`
	MeterCategory    string  `csv:"-"`
	MeterSubcategory string  `csv:"-"`
	MeterId          string  `csv:"-"`
	MeterName        string  `csv:"MeterName"`
	MeterRegion      string  `csv:"-"`
	UnitOfMeasure    string  `csv:"UnitOfMeasure"`
	UsageQuantity    float64 `csv:"UsageQuantity"`
	ResourceRate     float64 `csv:"ResourceRate"`
	PreTaxCost       float64 `csv:"PreTaxCost"`
	CostCenter       string  `csv:"-"`
	ConsumedService  string  `csv:"ConsumedService"`
	ResourceType     string  `csv:"ResourceType"`
	InstanceId       string  `csv:"InstanceId"`
	Tags             string  `csv:"-"`
	OfferId          string  `csv:"-"`
	AdditionalInfo   string  `csv:"-"`
	ServiceInfo1     string  `csv:"-"`
	ServiceInfo2     string  `csv:"-"`
	Currency         string  `csv:"Currency"`
	Datafile         string
}

type costExportData []rowData

type FieldMismatch struct {
	expected, found int
}

func (e *FieldMismatch) Error() string {
	return "CSV line fields mismatch. Expected " + strconv.Itoa(e.expected) + " found " + strconv.Itoa(e.found)
}

type UnsupportedType struct {
	Type string
}

type transformedCostItem struct {
	ResourceGroup    string
	PreTaxCost       float64
	SubscriptionName string
	Tenant           string
	Datafile         string
}

type transformedTenantData struct {
	PreTaxCost float64
	ResGroups  []transformedCostItem
}

type transformedCostItemsByTenant struct {
	Blue      transformedTenantData
	BlueDTQ   transformedTenantData
	Red       transformedTenantData
	RedDTQ    transformedTenantData
	Yellow    transformedTenantData
	PUD       transformedTenantData
	PUDDTQ    transformedTenantData
	Purple    transformedTenantData
	PurpleDTQ transformedTenantData
}

func main() {
	var dataPath = "./testdata/cost-exports"

	combinedCostData := combineCostExportData(dataPath)

	ccdJson, _ := json.MarshalIndent(combinedCostData, "", "  ")
	fmt.Println(string(ccdJson))

	// fmt.Println(len(combinedCostData))

	// transformedData := transformCostData(combinedCostData)

	// costData, err := getCostExportFileData("cost-exports/monthly-cost-exports_BLUEDTQ.csv")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(combinedCostData)
	// jsonData, _ := json.MarshalIndent(transformedData, "", "  ")
	// fmt.Println(string(jsonData))
	// _ = jsonData
	// fmt.Println(len(costData))
}

func combineCostExportData(dataPath string) costExportData {
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
		wg.Add(1)
		go func() {
			defer wg.Done()
			// fmt.Println(file)
			data, err := getCostExportFileData(file)
			if err != nil {
				panic(err)
			}
			mutex.Lock()
			// productNames = append(productNames)
			costExportData = append(costExportData, data...)
			mutex.Unlock()
		}()
	}
	wg.Wait()
	return costExportData
}

func transformCostData(data costExportData) transformedCostItemsByTenant {
	// fmt.Println(len(*data))

	var (
		// transformedTenantData transformedTenantData
		allData          transformedCostItemsByTenant
		allSubscriptions []string
	)
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
		case strings.Contains(sn, "hapdtq"):
			tenantName = "BLUEDTQ"
		default:
			tenantName = strings.ToUpper(costData.Datafile)
		}

		tci := transformedCostItem{
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

func (tenants *transformedCostItemsByTenant) addPreTaxCost(tci transformedCostItem) {

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

func (tenants *transformedCostItemsByTenant) appendTenantData(tci transformedCostItem) {
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

func getCostExportFileData(fileName string) (costExportData, error) {
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
