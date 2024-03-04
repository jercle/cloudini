package azure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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

func getCostData(subscriptionId string, resourceGroup string) {
	// 857d431a-06d9-4fe2-951e-cf697aa64376
	// rg-devdtqdesktop-w11dev
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
	r.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IjVCM25SeHRRN2ppOGVORGMzRnkwNUtmOTdaRSIsImtpZCI6IjVCM25SeHRRN2ppOGVORGMzRnkwNUtmOTdaRSJ9.eyJhdWQiOiJodHRwczovL21hbmFnZW1lbnQuY29yZS53aW5kb3dzLm5ldCIsImlzcyI6Imh0dHBzOi8vc3RzLndpbmRvd3MubmV0L2VhNTA4YmM3LWI0M2MtNGI5Ni04NDcwLTQ4OTc1NmU1OWExNC8iLCJpYXQiOjE3MDU4OTQ2NTUsIm5iZiI6MTcwNTg5NDY1NSwiZXhwIjoxNzA1OTAwMjMyLCJhY3IiOiIxIiwiYWlvIjoiQVZRQXEvOFZBQUFBUUNRVDN5TjZSTVk0cXdnK0FyVm5US3lGVDJGNFZVaG5IcmxvdWVERktjNUZsVGY5UXVuQkgyNHJkcVVqVTkxNFJGUEFZZzJXY0M4S2dmb0daWHg3TTZqb1hiM1M3WEM3TWliVnpLRWJHOWM9IiwiYW1yIjpbInB3ZCIsIm1mYSJdLCJhcHBpZCI6IjE4ZmJjYTE2LTIyMjQtNDVmNi04NWIwLWY3YmYyYjM5YjNmMyIsImFwcGlkYWNyIjoiMCIsImZhbWlseV9uYW1lIjoiYWRtLjExNDQwZSIsImdpdmVuX25hbWUiOiJFIiwiZ3JvdXBzIjpbIjc1NzE0Y2YzLWE0ZDMtNDYxMi1iY2I0LTdkMDA2N2M0Yzc3YiIsIjhiN2M4ZjA2LWIyMGMtNDg2My1hZjVlLTUwMmQzN2UwMWFjMSIsIjFkNjBkZjA4LWE4NDMtNDMzYy05YWM5LTVlODRhNTVhMjY5MiIsIjk2MTUwYTBiLTMzM2ItNDg5Zi1hNzZlLTc0ODgwOGQzOTU2YiIsIjJlNzdlMjJjLTRlMzMtNGEzMy1hMzRhLTk2YTNkYzNmNzFhNCIsIjc3MGNiYjQwLWE5YzgtNDhiMy1hOTcxLWMyM2VmYzk1Mzk2YSIsImQzZTJmMDRhLTNlZDItNGJhYi1hNmExLWQyYTE1Mjk2ZTdhMiIsIjg4ZTFiNjRlLWI1YWItNGNjNy05OGRhLTAwMTQyMmIwNDZiYiIsIjMxZjQ0ZDUxLWQ5MDMtNDc5My1hZDA4LTUwMTA5ZGIyNTZhYyIsIjI2ZjFlYzU1LWE3ZmMtNDlkNi1iNTI4LWRkZmU1ZjlmNjliMCIsIjNhM2JiZTY5LTYxZGYtNDVhYy1iMjFjLTEzZGMzZjhlZmYyNyIsIjhjNzMwZDZiLWYwN2MtNDIzNi1hNTA2LTkzNDZhYjQwZjlmZSIsImM0MDA0YzZmLTQ3ZjQtNDg1NC05OTg3LTc0ZDJlNDdkMzEwZSIsImU0MDVlNTdmLTBlOTQtNDc1Ni05ZDUyLTIyYzNjNWVhYTBmNSIsImFiZDVjYTkzLWRkNzktNDI3MS1hZWZiLTkzN2JmZDQxMmQ5NSIsImUwMDQ1Yjk1LTA2NDUtNGIxYS1iMmM2LWQ0YmVkZDAwYWJlNyIsImJmZGYyOTlkLTc1ZjEtNDNjNy05N2ZmLTdmOTQ0ZGMzY2E0MyIsImQxNjI4ZGExLWU4Y2MtNDNmOS04ZTliLWY5NTQ4NTIxYWM2NyIsImIxMTgzNmE0LTAxNGUtNDM1Ni04ZWFjLWE2MjI2Zjc0ZDljYiIsIjJhZTZhM2I1LTVhMWEtNGVmNC05NmQ3LWU2YmMxNGJmNTMzYiIsIjNmNDZmMGJjLWU4ZDAtNDg1Zi1hZDllLTE2NmE5NTg3Y2FlZiIsIjMzYTY2OGM2LTJlZWQtNDRkMy1iZjQ5LWMyNTU2M2YyYTFlZCIsIjYwZWE2MmM4LWQ0YTgtNGVmMy04MjFkLTFiYTFmZDgzMGUxNiIsImU3YTU2YmQ2LWU0YWMtNGRkMy04YmIzLTM2OTE4OTMzZTk0YyIsImIyMjY5NWUzLTQ0YzMtNDlkNC04ZjkxLTFlOWQ0NjkzODZhMyIsImQ0MmJkM2U2LWVhMjUtNGFkZS04OGVkLWU0OGE5M2Q1OWFkOCIsIjI1NzBkYWZkLTk0MzMtNDRiZC04M2M4LTMyYjE0MmFmNTcyNyJdLCJpZHR5cCI6InVzZXIiLCJpcGFkZHIiOiIxNDQuMTQwLjE1MC41IiwibmFtZSI6ImFkbS4xMTQ0MGUiLCJvaWQiOiIxYTdiYjc2Mi1jMmYyLTRmOGItOGZlZC03M2IyMjQyNjFkNDUiLCJvbnByZW1fc2lkIjoiUy0xLTUtMjEtMjI5OTM3NjUxMC0yNzQyNDA3NDE2LTIxNTk1NzcxMDEtMTQ4MCIsInB1aWQiOiIxMDAzMjAwMzNDMjE4MTQ5IiwicmgiOiIwLkFVRUF4NHRRNmp5MGxrdUVjRWlYVnVXYUZFWklmM2tBdXRkUHVrUGF3ZmoyTUJOQkFBNC4iLCJzY3AiOiJ1c2VyX2ltcGVyc29uYXRpb24iLCJzdWIiOiJJd0N4VWdUR01BMEZlbnBGMmlRNzB4ZkM5SWNSQWRhSVdkUG5aNVdUY0tBIiwidGlkIjoiZWE1MDhiYzctYjQzYy00Yjk2LTg0NzAtNDg5NzU2ZTU5YTE0IiwidW5pcXVlX25hbWUiOiJhZG0uMTE0NDBlQGR0cS5hc2lvLmdvdi5hdSIsInVwbiI6ImFkbS4xMTQ0MGVAZHRxLmFzaW8uZ292LmF1IiwidXRpIjoicElnNENWbTNqVVdWRVRzWmpqUTBBQSIsInZlciI6IjEuMCIsIndpZHMiOlsiZjJlZjk5MmMtM2FmYi00NmI5LWI3Y2YtYTEyNmVlNzRjNDUxIiwiYjc5ZmJmNGQtM2VmOS00Njg5LTgxNDMtNzZiMTk0ZTg1NTA5Il0sInhtc19jYWUiOiIxIiwieG1zX3RjZHQiOjE2MDcwNTAwMDB9.cthr8v1g-H7CzhTET3X54jVofvib3zqDWgBCqNGSP1diIp-_IbwQV9KUtIWEcTCiQBSapkfkThlOyRmXt4EUoV7rLXrV61pn628N0MJAj3f7aQvvOISEAUOHx5WIIvlreChpPD2kora90W_yyMsK6N-Isq-gVLCcmmAtfNPzevG6f6LPUlQkPw-c8UtLj6UNXL9TvkxChE9vcDQJOmDDoUyty6iskmMhEELo-k6I-xX4mPHfDxlM0mx1bgzohwYK8UrhbYIFH_YmAiLVAFKUgDFRoS0s_BurHPUtk5XNpsc83b9-O4iPPrCKcn9ia9uv3tQ0RejCHu_eekkNbXGEyA")

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
