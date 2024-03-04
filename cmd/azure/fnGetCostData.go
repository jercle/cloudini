package azure

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

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
