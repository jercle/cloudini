package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type CostQueryResponse struct {
	ID         string `json:"id"`
	Properties struct {
		NextLink string `json:"nextLink"`
		Columns  []struct {
			Name string `json:"name"`
		} `json:"columns"`
		Rows [][]interface{} `json:"rows"`
	} `json:"properties"`
}

type CostData struct {
	Cost          float64
	UsageDate     uint
	ResourceId    string
	ChargeType    string
	PublisherType string
	Currency      string
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

// func (c *CostData) MarshalJSON() ([]byte, error) {
// 	type Alias CostData
// 	return json.Marshal(&struct {
// 	 FooStatus string `json:"fooStatus"`
// 	 *Alias
// 	}{
// 	 FooStatus: c.FooStatus.String(),
// 	 Alias:     (*Alias)(c),
// 	})
//  }

func main2() {
	// file, err := os.Open("/home/jercle/git/azg/testdata/costData.json") // wsl
	file, err := os.Open("/Users/jercle/git/azg/testdata/costData.json") // mac
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var result CostQueryResponse

	json.Unmarshal([]byte(byteValue), &result)

	dataRows := result.Properties.Rows
	var data []CostData
	// data := make([]CostData, 5)

	for _, row := range dataRows {
		// var rowData CostData
		rowData := CostData{
			Cost:          row[0].(float64),
			UsageDate:     uint(row[1].(float64)),
			ResourceId:    row[2].(string),
			ChargeType:    row[3].(string),
			PublisherType: row[4].(string),
			Currency:      row[5].(string),
		}
		data = append(data, rowData)

	}
	fmt.Println(data)

}

// func main() {
// 	subscriptionId :=
// 	resourceGroup :=
// 	urlString := "https://management.azure.com/subscriptions/" + subscriptionId + "/resourceGroups/" + resourceGroup + "/providers/Microsoft.CostManagement/query?api-version=2023-11-01"

// 	body := []byte(`{
// 			"type": "ActualCost",
// 			"dataSet": {
// 					"granularity": "Daily",
// 					"aggregation": {
// 							"totalCost": {
// 									"name": "Cost",
// 									"function": "Sum"
// 							}
// 					},
// 					"sorting": [
// 							{
// 									"direction": "ascending",
// 									"name": "UsageDate"
// 							}
// 					],
// 					"grouping": [
// 							{
// 									"type": "Dimension",
// 									"name": "ResourceId"
// 							},
// 							{
// 									"type": "Dimension",
// 									"name": "ChargeType"
// 							},
// 							{
// 									"type": "Dimension",
// 									"name": "PublisherType"
// 							}
// 					]
// 			},
// 			"timeframe": "Custom",
// 			"timePeriod": {
// 					"from": "2024-01-16T00:00:00+00:00",
// 					"to": "2024-01-21T23:59:59+00:00"
// 			}
// 	}`)
// 	// r, err := http.Post(urlString, "application/json", bytes.NewBuffer(body))
// 	// http.Post does not allow for custom headers
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	r, err := http.NewRequest("POST", urlString, bytes.NewBuffer(body))
// 	if err != nil {
// 		panic(err)
// 	}
// 	r.Header.Add("Content-Type", "application/json")
// 	r.Header.Add("Authorization", "Bearer "+token.token)

// 	client := &http.Client{}
// 	res, err := client.Do(r)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer res.Body.Close()

// 	// post := &Post{}
// 	response := &CostQueryResponse{}
// 	derr := json.NewDecoder(res.Body).Decode(response)
// 	if derr != nil {
// 		panic(derr)
// 	}

// 	fmt.Println(res.StatusCode)
// 	// if res.StatusCode != http.StatusOK {
// 	// panic(derr)
// 	// }

// 	jsonString, err := json.MarshalIndent(response, "", "  ")

// 	// topRow := response.Properties.Columns
// 	// dataRows := response.Properties.Rows
// 	// // var constructedMap []CostData

// 	// if err != nil {
// 	// 	log.Fatal(nil)
// 	// }

// 	// for _, row := range dataRows {
// 	// 	var costData CostData
// 	// 	for _, col := range topRow {
// 	// 		costData = row
// 	// 		// fmt.Println(col)
// 	// 		// fmt.Println(i)
// 	// 		// fmt.Println(row)
// 	// 	}
// 	// }

// 	fmt.Println(string(jsonString))
// }
