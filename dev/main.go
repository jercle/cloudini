package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/jercle/azg/lib"
)

type PriceSheetItem struct {
	BasePrice          any     `json:"BasePrice"`
	BillingAccountID   string  `json:"BillingAccountId"`
	BillingAccountName string  `json:"BillingAccountName"`
	BillingCurrency    string  `json:"BillingCurrency"`
	BillingProfileID   string  `json:"BillingProfileId"`
	BillingProfileName string  `json:"BillingProfileName"`
	Currency           string  `json:"Currency"`
	EffectiveEndDate   string  `json:"EffectiveEndDate"`
	EffectiveStartDate string  `json:"EffectiveStartDate"`
	MarketPrice        float64 `json:"MarketPrice"`
	MeterCategory      string  `json:"MeterCategory"`
	MeterID            string  `json:"MeterId"`
	MeterName          string  `json:"MeterName"`
	MeterRegion        string  `json:"MeterRegion"`
	MeterSubCategory   string  `json:"MeterSubCategory"`
	MeterType          string  `json:"MeterType"`
	PriceType          string  `json:"PriceType"`
	Product            string  `json:"Product"`
	ProductID          string  `json:"ProductId"`
	ServiceFamily      string  `json:"ServiceFamily"`
	SkuID              string  `json:"SkuId"`
	Term               string  `json:"Term"`
	TierMinimumUnits   float64 `json:"TierMinimumUnits"`
	UnitOfMeasure      string  `json:"UnitOfMeasure"`
	UnitPrice          float64 `json:"UnitPrice"`
}

func main() {

}

func combinePriceSheets(path string) ([]PriceSheetItem, error) {
	var (
		wg             sync.WaitGroup
		mutex          sync.Mutex
		fullPriceSheet []PriceSheetItem
	)
	filePaths := lib.GetFullFilePaths(path)

	for _, file := range filePaths {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data, err := getPricingSheetData(file)
			if err != nil {
				panic(err)
			}
			mutex.Lock()
			fullPriceSheet = append(fullPriceSheet, data...)
			mutex.Unlock()
		}()
	}
	wg.Wait()
	return fullPriceSheet, nil
}

func getPricingSheetData(path string) ([]PriceSheetItem, error) {
	var allPricingSheetItems []PriceSheetItem

	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		return nil, err
	}

	r := bufio.NewReader(f)
	s, err := ReadPricingSheetLine(r)
	allPricingSheetItems = append(allPricingSheetItems, s)
	if err != nil {
		return nil, err
	}
	for err == nil {
		s, err = ReadPricingSheetLine(r)
		allPricingSheetItems = append(allPricingSheetItems, s)
	}

	return allPricingSheetItems, nil
}

func ReadPricingSheetLine(r *bufio.Reader) (PriceSheetItem, error) {
	var (
		item     PriceSheetItem
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}

	json.Unmarshal(ln, &item)
	return item, err
}

func ChunkPricingSheet(path string) {
	jsonFile, err := os.Open("fullFakePricingSheet.json")
	lib.CheckFatalError(err)
	defer jsonFile.Close()
}
