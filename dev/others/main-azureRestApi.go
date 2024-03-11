package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type providerListResult struct {
	Value []struct {
		Namespace     string `json:"namespace"`
		ResourceTypes []struct {
			APIProfiles []struct {
				APIVersion     string `json:"apiVersion"`
				ProfileVersion string `json:"profileVersion"`
			} `json:"apiProfiles"`
			APIVersions       []string `json:"apiVersions"`
			Capabilities      string   `json:"capabilities"`
			DefaultAPIVersion string   `json:"defaultApiVersion,omitempty"`
			Locations         []string `json:"locations"`
			ResourceType      string   `json:"resourceType"`
			ZoneMappings      []any    `json:"zoneMappings"`
		} `json:"resourceTypes"`
	} `json:"value"`
}

type azureResourceTypes struct {
	ProviderNamespaces []string
	ResourceTypes      []string
}

func main() {

	// token := azure.GetAzCliToken()
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(token)

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
	// urlString := "https://management.azure.com/providers?api-version=2021-04-01"
	// urlString := "https://management.azure.com/subscriptions/bae338c7-6098-4d52-b173-e2147e107dfa/providers/Microsoft.HDInsight/locations/australiaeast/billingSpecs?api-version=2021-06-01"
	// urlString := "https://management.azure.com/subscriptions/bae338c7-6098-4d52-b173-e2147e107dfa/providers/Microsoft.Consumption/pricesheets/default?api-version=2019-10-01"
	urlString := "https://prices.azure.com/api/retail/prices?currencyCode='AUD'&armRegionName='australiaeast'"
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		log.Fatal(err)
	}

	// GET https://management.azure.com/subscriptions/bae338c7-6098-4d52-b173-e2147e107dfa/providers/Microsoft.Web/billingMeters?api-version=2022-03-01
	// POST https://management.azure.com/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/providers/Microsoft.CostManagement/pricesheets/default/download?api-version=2023-11-01
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {

		log.Fatal("Error fetching LA Workspace Tables")
	}

	responseBody, err := io.ReadAll(res.Body)
	if res.StatusCode == 400 {
		log.Fatal("Error fetching LA Workspace Tables: ", string(responseBody))
	}

	fmt.Println(string(responseBody))

	// var resBody providerListResult
	//
	// json.Unmarshal(responseBody, &resBody)

	// fmt.Println(resBody)

	// var azureResourceTypes azureResourceTypes

	// for _, val := range resBody.Value {
	// 	azureResourceTypes.ProviderNamespaces = append(azureResourceTypes.ProviderNamespaces, val.Namespace)

	// 	for _, resType := range val.ResourceTypes {
	// 		azureResourceTypes.ResourceTypes = append(azureResourceTypes.ResourceTypes, val.Namespace+"/"+resType.ResourceType)
	// 	}
	// }

	// jsonData, _ := json.MarshalIndent(resBody.Value, "", "  ")

	// fmt.Println(string(jsonData))
}
