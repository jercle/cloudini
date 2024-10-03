package main

import (
	"encoding/json"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func main() {
	config := lib.GetCldConfig(nil)
	_ = config
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	// token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{})
	tokenReq, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{}, &lib.CldConfigOptions{
		ConfigFilePath: "/home/jercle/.config/cld/cldConf.json",
	})
	lib.CheckFatalError(err)
	token, err := tokenReq.SelectTenant("RED")
	lib.CheckFatalError(err)
	// _ = tokens
	_ = token

	subscriptionId := ""
	filePath := "./imports"
	TerraformGenerateAllPrivateZonesAndRecordSets(subscriptionId, token, filePath)

	// jsonStr, _ := json.MarshalIndent(allZoneRecordSets, "", "  ")
	// fmt.Println(string(jsonStr))
}

func TerraformGenerateAllPrivateZonesAndRecordSets(subscriptionId string, token *lib.MultiAuthToken, filePath string) {
	var allZoneRecordSets []PrivateZoneRecordSet
	privateZones := ListAllPrivateZonesInSubscription(subscriptionId, token)

	for _, zone := range privateZones {
		// fmt.Println(zone.Name)
		zoneRecordSets := ListAllPrivateZoneRecordSets(zone.ID, zone.Name, token)
		allZoneRecordSets = append(allZoneRecordSets, zoneRecordSets...)
	}

	// importBlocks := ""
	importBlocks := make(map[string]string)
	var resourceNames []string

	for _, zone := range privateZones {
		processedName := ProcessNamesToTerraformReference(zone.Name)
		checkedName := CheckResourceNameIsUnique(&resourceNames, processedName)
		resourceNames = append(resourceNames, checkedName)
		entry := "\nimport {\n" +
			"  id = \"" +
			zone.ID +
			"\"\n  to = azurerm_private_dns_zone." +
			checkedName +
			"\n}\n\n"

		importBlocks["azurerm_private_dns_zone."] += entry
	}

	for _, rs := range allZoneRecordSets {
		// fmt.Println(rs.ZoneName)
		recordType := strings.ToLower(strings.Split(rs.Type, "privateDnsZones/")[1])
		if recordType != "soa" {
			processedRecordSetName := ProcessNamesToTerraformReference(rs.Name)
			processedZoneName := ProcessNamesToTerraformReference(rs.ZoneName)
			processedName := processedZoneName + "__" + processedRecordSetName

			terraformResourceType := "azurerm_private_dns_" +
				recordType +
				"_record."

			fullRecordName := terraformResourceType +
				processedName

			checkedName := CheckResourceNameIsUnique(&resourceNames, fullRecordName)
			resourceNames = append(resourceNames, checkedName)

			// fmt.Println(recordType)

			entry := "\nimport {\n" +
				"  id = \"" +
				rs.ID +
				"\"\n  to = " +
				checkedName +
				"\n}\n\n"

			importBlocks[terraformResourceType] += entry
		}
	}

	// importBlocks = strings.TrimSpace(importBlocks)
	for key, val := range importBlocks {
		savePath := filePath + "-" + key + "tf"
		SaveImportBlocks(val, savePath)
	}
}

// func ListAllPrivateZoneRecordSets(subscriptionId string, resourceGroupName string, privateZoneName string, token *lib.MultiAuthToken) {
func ListAllPrivateZoneRecordSets(privateZoneResourceId string, zoneName string, token *lib.MultiAuthToken) []PrivateZoneRecordSet {
	var response ListAllPrivateZoneRecordSetsResponse
	urlString := "https://management.azure.com/" +
		privateZoneResourceId +
		"/ALL?api-version=2018-09-01"

	res, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)
	json.Unmarshal(res, &response)

	// fmt.Println(zoneName)

	var processedRecordSets []PrivateZoneRecordSet
	for _, rs := range response.Value {
		recordSet := rs
		recordSet.ZoneName = zoneName
		// fmt.Println(recordSet.ZoneName)
		processedRecordSets = append(processedRecordSets, recordSet)
	}

	return processedRecordSets
}

func ListAllPrivateZonesInSubscription(subscriptionId string, token *lib.MultiAuthToken) []PrivateZone {
	var (
		listPrivateZonesResponse ListAllPrivateZonesInSubscriptionResponse
	)

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/providers/Microsoft.Network/privateDnsZones?api-version=2018-09-01"
	res, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	json.Unmarshal(res, &listPrivateZonesResponse)

	return listPrivateZonesResponse.Value
}

type ListAllPrivateZoneRecordSetsResponse struct {
	Value []PrivateZoneRecordSet `json:"value"`
}

type PrivateZoneRecordSet struct {
	Etag       string `json:"etag"`
	ID         string `json:"id"`
	Name       string `json:"name"`
	Properties struct {
		ARecords []struct {
			Ipv4Address string `json:"ipv4Address"`
		} `json:"aRecords,omitempty"`
		CnameRecord *struct {
			Cname string `json:"cname"`
		} `json:"cnameRecord,omitempty"`
		Fqdn             string `json:"fqdn"`
		IsAutoRegistered bool   `json:"isAutoRegistered"`
		SoaRecord        *struct {
			Email        string  `json:"email"`
			ExpireTime   float64 `json:"expireTime"`
			Host         string  `json:"host"`
			MinimumTtl   float64 `json:"minimumTtl"`
			RefreshTime  float64 `json:"refreshTime"`
			RetryTime    float64 `json:"retryTime"`
			SerialNumber float64 `json:"serialNumber"`
		} `json:"soaRecord,omitempty"`
		Ttl float64 `json:"ttl"`
	} `json:"properties"`
	Type     string `json:"type"`
	ZoneName string `json:"zoneName"`
}

type ListAllPrivateZonesInSubscriptionResponse struct {
	Value []PrivateZone `json:"value"`
}

type PrivateZone struct {
	Etag       string `json:"etag"`
	ID         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		MaxNumberOfRecordSets                          float64 `json:"maxNumberOfRecordSets"`
		MaxNumberOfVirtualNetworkLinks                 float64 `json:"maxNumberOfVirtualNetworkLinks"`
		MaxNumberOfVirtualNetworkLinksWithRegistration float64 `json:"maxNumberOfVirtualNetworkLinksWithRegistration"`
		NumberOfRecordSets                             float64 `json:"numberOfRecordSets"`
		NumberOfVirtualNetworkLinks                    float64 `json:"numberOfVirtualNetworkLinks"`
		NumberOfVirtualNetworkLinksWithRegistration    float64 `json:"numberOfVirtualNetworkLinksWithRegistration"`
		ProvisioningState                              string  `json:"provisioningState"`
	} `json:"properties"`
	Tags struct{} `json:"tags"`
	Type string   `json:"type"`
}

func RemoveLastRune(s string) string {
	r := []rune(s)
	return string(r[:len(r)-1])
}

func CheckResourceNameIsUnique(resourceNames *[]string, resourceName string) string {
	updatedName := resourceName
	inc := 1
	var checkUnniqueAndIncrement func()
	checkUnniqueAndIncrement = func() {
		if slices.Contains(*resourceNames, updatedName) {
			lastChar := updatedName[len(updatedName)-1:]
			if _, err := strconv.Atoi(lastChar); err == nil {
				updatedName = RemoveLastRune(updatedName)
				inc, err = strconv.Atoi(lastChar)
				lib.CheckFatalError(err)
				inc++
				updatedName = updatedName + strconv.Itoa(inc)
				inc++
			} else {
				updatedName = updatedName + strconv.Itoa(inc)
				inc++
			}
			checkUnniqueAndIncrement()
		}
	}
	checkUnniqueAndIncrement()

	return updatedName
}

// func CheckResourceNameIsUnique(resourceNames *[]string, resourceName string) string {
// 	if slices.Contains(*resourceNames, resourceName) {
// 		return resourceName + "ADD"
// 	} else {
// 		return resourceName
// 	}
// }

func SaveImportBlocks(fileData string, filePath string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(fileData); err != nil {
		panic(err)
	}
	// byteData := []byte(fileData)
	// err := os.WriteFile(filePath, byteData, 0644)
	// lib.CheckFatalError(err)
}

func ProcessNamesToTerraformReference(name string) string {
	processedName := strings.ReplaceAll(name, " ", "_")
	processedName = strings.ReplaceAll(processedName, "/", "")
	processedName = strings.ReplaceAll(processedName, "\"", "")
	processedName = strings.ReplaceAll(processedName, ",", "")
	processedName = strings.ReplaceAll(processedName, "+", "")
	processedName = strings.ReplaceAll(processedName, "(", "")
	processedName = strings.ReplaceAll(processedName, ")", "")
	processedName = strings.ReplaceAll(processedName, "[", "")
	processedName = strings.ReplaceAll(processedName, "]", "")
	processedName = strings.ReplaceAll(processedName, "'", "")
	processedName = strings.ReplaceAll(processedName, "-", "_")
	processedName = strings.ReplaceAll(processedName, "*", "root")
	processedName = strings.ReplaceAll(processedName, "@", "root")
	processedName = strings.ReplaceAll(processedName, ".", "")
	processedName = strings.ReplaceAll(processedName, "__", "_")
	processedName = strings.ReplaceAll(processedName, "__", "_")
	processedName = strings.ToLower(processedName)

	if processedName[0] >= '0' && processedName[0] <= '9' {
		processedName = "_" + processedName
	}
	return processedName
}
