package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/jercle/cloudini/lib"
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
	// ChunkPriceSheet()
	CombinePriceSheet()
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

func ChunkPriceSheet() {
	var (
		// priceItems []PriceSheetItem
		priceItemsChunk  []PriceSheetItem
		pricingSheetPath = "./fakedata/pricing-sheet-chunks"
	)
	start := time.Now()

	f, err := os.Open("./outputs/fullFakePricingSheet.json")
	lib.CheckFatalError(err)
	defer f.Close()

	if !lib.CheckDirExists(pricingSheetPath) {
		os.MkdirAll(pricingSheetPath, os.ModePerm)
	}

	r := bufio.NewReader(f)
	d := json.NewDecoder(r)
	i := 0
	fileCount := 0

	d.Token()
	for d.More() {
		item := &PriceSheetItem{}
		d.Decode(item)
		priceItemsChunk = append(priceItemsChunk, *item)

		i++

		if math.Mod(float64(i), 1000) == 0 {
			// fmt.Println(i, " chunk of 100")
			chunkBytes, err := json.Marshal(priceItemsChunk)
			lib.CheckFatalError(err)
			err = os.WriteFile(pricingSheetPath+"/fakePricingSheet-"+strconv.Itoa(i)+".json", chunkBytes, os.ModePerm)
			priceItemsChunk = nil
			fileCount++
		}
	}
	d.Token()

	chunkBytes, err := json.Marshal(priceItemsChunk)
	lib.CheckFatalError(err)
	fileCount++

	err = os.WriteFile(pricingSheetPath+"/fakePricingSheet-"+strconv.Itoa(i)+".json", chunkBytes, os.ModePerm)

	elapsed := time.Since(start)
	fmt.Printf("Total of [%v] objects created in [%v] files.\n", i, fileCount)
	fmt.Printf("To parse the file took [%v]\n", elapsed)
}

func CombinePriceSheet() {
	var (
		fullPriceSheet []PriceSheetItem
		// pricingSheetPath = "./fakedata/combinedData"
		pricingSheetChunkPath = "./fakedata/pricing-sheet-chunks"
		mut                   sync.Mutex
		wg                    sync.WaitGroup
	)
	start := time.Now()

	// ReadChunkedPricingSheetFile()
	filesParsed := 0
	filenames := lib.GetFullFilePaths(pricingSheetChunkPath)

	for _, fn := range filenames {
		wg.Add(1)
		// fmt.Println(fn)
		go func() {
			chunkData, err := ReadChunkedPricingSheetFile(fn)
			lib.CheckFatalError(err)
			mut.Lock()
			fullPriceSheet = append(fullPriceSheet, chunkData...)
			filesParsed++
			mut.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	byteValue, err := json.MarshalIndent(fullPriceSheet, "", "  ")
	lib.CheckFatalError(err)

	if !lib.CheckDirExists("./outputs") {
		os.MkdirAll("./outputs", os.ModePerm)
	}

	os.WriteFile("./outputs/combinedPriceSheetIndent.json", byteValue, os.ModePerm)

	elapsed := time.Since(start)
	fmt.Printf("Total of [%v] objects parsed from [%v] files.\n", len(fullPriceSheet), filesParsed)
	fmt.Printf("To parse the file took [%v]\n", elapsed)
}

func ReadChunkedPricingSheetFile(filename string) ([]PriceSheetItem, error) {
	var (
		priceItemSheet []PriceSheetItem
	)
	file, err := os.Open(filename)
	lib.CheckFatalError(err)

	byteValue, err := io.ReadAll(file)
	lib.CheckFatalError(err)

	err = json.Unmarshal(byteValue, &priceItemSheet)
	lib.CheckFatalError(err)

	return priceItemSheet, nil
}
