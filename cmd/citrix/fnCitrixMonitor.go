package citrix

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/jercle/cloudini/lib"
)

func GetAllMachineMetrics(creds lib.CitrixCloudAccountConfig, tokenData lib.CitrixTokenData) []MachineMetric {
	// fmt.Println(time.Now().Add(-11 * time.Hour).Format("2006-01-02T15:04:05.000Z"))
	// fmt.Println(strings.Split(time.Now().Add(-11*time.Hour).Format(time.RFC3339), "+")[0])
	// os.Exit(0)

	formattedDate := time.Now().Add(-11 * time.Hour).Format("2006-01-02T15:04:05.000Z")
	// formattedDate := time.Now().Add(-11 * time.Hour).Format("2006-01-02T15:04:05.000Z")
	filterString := "$filter=CollectedDate gt " + formattedDate + "&$orderby=CollectedDate desc"
	urlString := "https://api.cloud.com/monitorodata/MachineMetric?" + url.PathEscape(filterString)

	res, err := HttpGet(urlString, creds.CustomerId, creds.SiteId, tokenData)
	lib.CheckFatalError(err)

	// fmt.Println(string(res))

	var resData MachineMetricResponse
	err = json.Unmarshal(res, &resData)
	lib.CheckFatalError(err)

	metrics := resData.Value

	nextLink := resData.NextLink

	for nextLink != "" {
		var curr MachineMetricResponse
		// fmt.Println(nextLink)

		resNew, _ := HttpGet(nextLink, creds.CustomerId, creds.SiteId, tokenData)
		json.Unmarshal(resNew, &curr)
		// fmt.Println(string(res))

		nextLink = curr.NextLink
		metrics = append(metrics, curr.Value...)

	}

	jsonStr, _ := json.Marshal(metrics, jsontext.WithIndent("  "))
	os.WriteFile("main-citrix-machineMetrics.json", jsonStr, 0644)
	return metrics
}

func GetAllMachineResourceUtilisation(creds lib.CitrixCloudAccountConfig, tokenData lib.CitrixTokenData) []MachineResourceUtilisation {
	formattedDate := time.Now().Add(-1 * time.Hour).Format("2006-01-02T15:04:05.000Z")
	filterString := "$filter=CollectedDate gt " + formattedDate + "&$orderby=CollectedDate desc"
	urlString := "https://api.cloud.com/monitorodata/ResourceUtilization?" + url.PathEscape(filterString)

	res, err := HttpGet(urlString, creds.CustomerId, creds.SiteId, tokenData)
	lib.CheckFatalError(err)

	// fmt.Println(string(res))

	var resData MachineResourceUtilisationResponse
	err = json.Unmarshal(res, &resData)
	lib.CheckFatalError(err)

	resUtl := resData.Value

	nextLink := resData.NextLink

	for nextLink != "" {
		var curr MachineResourceUtilisationResponse
		// fmt.Println(nextLink)

		resNew, _ := HttpGet(nextLink, creds.CustomerId, creds.SiteId, tokenData)
		json.Unmarshal(resNew, &curr)
		// fmt.Println(string(res))

		nextLink = curr.NextLink
		resUtl = append(resUtl, curr.Value...)

	}

	// jsonStr, _ := json.Marshal(resUtl, jsontext.WithIndent("  "))
	// os.WriteFile("main-citrix-resourceUtilisation.json", jsonStr, 0644)
	return resUtl
}

//
//

func GetAllMachineLoadIndexes(creds lib.CitrixCloudAccountConfig, tokenData lib.CitrixTokenData) []MachineLoadIndex {
	formattedDate := time.Now().Add(-10 * time.Hour).Add(-30 * time.Minute).Format("2006-01-02T15:04:05.000Z")
	// formattedDate := time.Now().Add(-12 * time.Hour).Format("2006-01-02T15:04:05.000Z")
	filterString := "$filter=CreatedDate gt " + formattedDate + "&$orderby=CreatedDate desc"
	// filterString := ""
	urlString := "https://api.cloud.com/monitorodata/LoadIndexes?" + url.PathEscape(filterString)

	res, err := HttpGet(urlString, creds.CustomerId, creds.SiteId, tokenData)
	lib.CheckFatalError(err)

	// fmt.Println(string(res))
	// os.Exit(0)

	var resData MachineLoadIndexesResponse
	err = json.Unmarshal(res, &resData)
	lib.CheckFatalError(err)

	metrics := resData.Value

	nextLink := resData.NextLink

	for nextLink != "" {
		var curr MachineLoadIndexesResponse
		fmt.Println(nextLink)

		resNew, _ := HttpGet(nextLink, creds.CustomerId, creds.SiteId, tokenData)
		json.Unmarshal(resNew, &curr)
		// fmt.Println(string(res))

		nextLink = curr.NextLink
		metrics = append(metrics, curr.Value...)

	}

	// jsonStr, _ := json.Marshal(metrics, jsontext.WithIndent("  "))
	// os.WriteFile("main-citrix-loadIndexes.json", jsonStr, 0644)
	return metrics
}

//
//

func GetAllMonitorMachines(creds lib.CitrixCloudAccountConfig, tenantName string, tokenData lib.CitrixTokenData) []MonitorMachine {
	// formattedDate := time.Now().Add(-10 * time.Hour).Add(-30 * time.Minute).Format("2006-01-02T15:04:05.000Z")
	// formattedDate := time.Now().Add(-12 * time.Hour).Format("2006-01-02T15:04:05.000Z")
	// filterString := "$filter=CreatedDate gt " + formattedDate + "&$orderby=CreatedDate desc"

	filterString := "$filter=CurrentPowerState eq 3 and LifecycleState eq 0&$expand=CurrentLoadIndex&$select=CurrentPowerState,AgentVersion,IPAddress,DnsName,Id,IsInMaintenanceMode,AssociatedUserNames,CurrentSessionCount,FaultState,CurrentRegistrationState,Name,CurrentLoadIndexID"
	urlString := "https://api.cloud.com/monitorodata/Machines?" + url.PathEscape(filterString)

	res, err := HttpGet(urlString, creds.CustomerId, creds.SiteId, tokenData)
	lib.CheckFatalError(err)

	// fmt.Println(string(res))
	// os.Exit(0)

	var resData MonitorMachinesResponse
	err = json.Unmarshal(res, &resData)
	lib.CheckFatalError(err)

	// machines := resData.Value

	var machines []MonitorMachine

	for _, m := range resData.Value {
		curr := m
		curr.TenantName = tenantName
		machines = append(machines, curr)
	}

	nextLink := resData.NextLink

	for nextLink != "" {
		// fmt.Println(nextLink)

		resNew, _ := HttpGet(nextLink, creds.CustomerId, creds.SiteId, tokenData)
		var currRes MonitorMachinesResponse
		err = json.Unmarshal(resNew, &currRes)
		// fmt.Println(string(res))

		nextLink = currRes.NextLink

		for _, m := range currRes.Value {
			curr := m
			curr.TenantName = tenantName
			machines = append(machines, curr)
		}

	}

	// jsonStr, _ := json.Marshal(machines, jsontext.WithIndent("  "))
	// os.WriteFile("main-citrix-monitorMachines.json", jsonStr, 0644)
	return machines
}
