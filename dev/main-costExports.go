package main

import (
	"os"

	"github.com/gocarina/gocsv"
)

type costExport []struct {
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
}

func main() {

	file, err := os.Open("cost-exports/monthly-cost-exports_BLUEDTQ-single-noempty.csv")
	if err != nil {
		panic(err)
	}
	// reader := gocsv.LazyCSVReader(file)
	// csv.NewReader()
	// reader := csv.NewReader(file)
	// reader.LazyQuotes = true
	// reader.TrimLeadingSpace = true

	// reader := csv.NewReader(file)
	// if err != nil {
	// 	panic(err)
	// }
	// data, err := reader.ReadAll()

	// fmt.Println(data)

	// // data, err := reader.ReadAll()

	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// Read the CSV data
	// reader := csv.NewReader(file)
	// reader := gocsv.CSVReader(file)
	// reader.FieldsPerRecord = 0 // Allow variable number of fields
	// data, err := reader.ReadAll()
	// if err != nil {
	// 	panic(err)
	// }

	// var allData *costExport
	var allData interface{}

	if err := gocsv.UnmarshalFile(file, allData); err != nil {
		panic(err)
	}

}
