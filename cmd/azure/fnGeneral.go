package azure

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jercle/cloudini/lib"
)

func CheckStorageAccountTlsVersions(outputFile string, getAll bool, token *lib.AzureMultiAuthToken) (allResources []StorageAccountTlsVersion) {
	cldConfig := lib.GetCldConfig(nil)

	urlString := "https://management.azure.com/providers/Microsoft.ResourceGraph/resources?api-version=2022-10-01"

	graphQuery := `Resources
    | where type == 'microsoft.storage/storageaccounts'
    | where properties['minimumTlsVersion'] != 'TLS1_2'
    | project name, resourceGroup, minimumTlsVersion = properties.minimumTlsVersion, tenantId, subscriptionId, id`

	if getAll {
		graphQuery = `Resources
    | where type == 'microsoft.storage/storageaccounts'
    | project name, resourceGroup, minimumTlsVersion = properties.minimumTlsVersion, tenantId, subscriptionId, id`
	}

	jsonBody := `{
	"query": "` + graphQuery + `"
}`

	res, _, err := HttpPost(urlString, jsonBody, *token)
	// if err != nil {
	// 	fmt.Println(token.TenantName)
	// }
	lib.CheckFatalError(err)

	var response CheckStorageAccountTlsVersionsResponse
	err = json.Unmarshal(res, &response)
	lib.CheckFatalError(err)

	subs, err := ListSubscriptions(*token)
	lib.CheckFatalError(err)

	// lib.JsonMarshalAndPrint(subs)
	// os.Exit(0)

	for _, res := range response.Data {
		currRes := res

		for _, sub := range subs {
			if sub.SubscriptionID == res.SubscriptionID {
				currRes.SubscriptionName = sub.DisplayName
			}
		}

		currRes.ID = strings.ToLower(res.ID)
		currRes.LastAzureSync = time.Now()
		currRes.TenantName = lib.MapTenantIdToConfiguredTenantName(res.TenantID, *cldConfig.Azure)

		allResources = append(allResources, currRes)
	}

	hasSkipToken := false
	skipToken := ""

	if response.SkipToken != "" {
		hasSkipToken = true
		skipToken = response.SkipToken
	}

	for hasSkipToken {
		var whileRes CheckStorageAccountTlsVersionsResponse
		jsonBody := `{
			"query": "` + graphQuery + `",
			"options": {
				"$skipToken": "` + skipToken + `"
			}
		}`

		res, _, err := HttpPost(urlString, jsonBody, *token)
		lib.CheckFatalError(err)
		err = json.Unmarshal(res, &whileRes)
		lib.CheckFatalError(err)

		for _, res := range whileRes.Data {
			currRes := res
			currRes.ID = strings.ToLower(res.ID)
			currRes.LastAzureSync = time.Now()
			currRes.TenantName = lib.MapTenantIdToConfiguredTenantName(res.TenantID, *cldConfig.Azure)
			allResources = append(allResources, currRes)
		}

		if whileRes.SkipToken != "" {
			hasSkipToken = true
			skipToken = whileRes.SkipToken
		} else {
			hasSkipToken = false
			skipToken = ""
		}
	}

	return
}

//
//

func CheckStorageAccountTlsVersionsForAllConfiguredTenants(opts *lib.GetAllResourcesForAllConfiguredTenantsOptions, tokens lib.AllTenantTokens) (allResources []StorageAccountTlsVersion) {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)
	azConfig := lib.GetCldConfig(nil)
	azConfigs := azConfig.Azure.MultiTenantAuth.Tenants

	for _, token := range tokens {
		if azConfigs[token.TenantName].IsB2C {
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !opts.SuppressSteps {
				fmt.Println(token.TenantName + ": Fetching resources")
			}
			// allResources[token.TenantName] = make(map[string]SubscriptionResourceList)
			storageAccounts := CheckStorageAccountTlsVersions("", opts.GetAllStorageAccountsInTlsCheck, &token)

			mutex.Lock()
			allResources = append(allResources, storageAccounts...)
			mutex.Unlock()
			if !opts.SuppressSteps {
				fmt.Println(token.TenantName + ": Fetch complete")
			}

		}()
	}

	wg.Wait()

	options := *opts

	outputFilePath := options.OutputFilePath

	if outputFilePath != "" {
		jsonStr, _ := json.MarshalIndent(allResources, "", "  ")

		currentDate := time.Now().Format("20060102")

		fileName := outputFilePath + "/allRes-Storage-CheckTls-" + currentDate + ".json"

		err := os.WriteFile(fileName, jsonStr, 0644)
		lib.CheckFatalError(err)

	}

	// fmt.Println(len(allResourcesSlice))

	return allResources
}
