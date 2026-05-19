package citrix

import (
	"encoding/json/v2"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/jercle/cloudini/lib"
)

func GetAllMachineMetrics(creds lib.CitrixCloudAccountConfig, tenantName string, tokenData lib.CitrixTokenData) []MachineMetric {
	// fmt.Println(time.Now().Add(-11 * time.Hour).Format("2006-01-02T15:04:05.000Z"))
	// fmt.Println(strings.Split(time.Now().Add(-11*time.Hour).Format(time.RFC3339), "+")[0])
	// os.Exit(0)

	// formattedDate := time.Now().Add(-24 * time.Hour).Format("2006-01-02T15:04:05.000Z")
	// formattedDate := time.Now().Add(-11 * time.Hour).Format("2006-01-02T15:04:05.000Z")
	// date, err := time.Parse("2006-01-02T15:04:05.000Z", "2026-05-01T00:04:05.000Z")
	// lib.CheckFatalError(err)
	// formattedDate := date.Format("2006-01-02T15:04:05.000Z")
	// formattedDate := time.Now().Add(-11 * time.Hour).Format("2006-01-02T15:04:05.000Z")

	formattedDate := time.Now().Add(-730 * time.Minute).Format("2006-01-02T15:04:05.000Z")
	filterString := "$filter=CollectedDate gt " + formattedDate + "&$orderby=CollectedDate desc"
	// filterString := "$orderby=CollectedDate desc"
	urlString := "https://api.cloud.com/monitorodata/MachineMetric?" + url.PathEscape(filterString)

	res, err := HttpGet(urlString, creds.CustomerId, creds.SiteId, tokenData)
	lib.CheckFatalError(err)

	// fmt.Println(string(res))

	var resData MachineMetricResponse
	err = json.Unmarshal(res, &resData)
	lib.CheckFatalError(err)

	var metrics []MachineMetric

	for _, m := range resData.Value {
		curr := m
		curr.TenantName = tenantName
		metrics = append(metrics, curr)
	}

	nextLink := resData.NextLink
	iteration := 1
	for nextLink != "" {
		var curr MachineMetricResponse
		iteration++
		fmt.Println(tenantName+" - Metrics:", strconv.Itoa(iteration))

		resNew, _ := HttpGet(nextLink, creds.CustomerId, creds.SiteId, tokenData)
		json.Unmarshal(resNew, &curr)
		// fmt.Println(string(res))

		nextLink = curr.NextLink

		for _, m := range curr.Value {
			currLI := m
			currLI.TenantName = tenantName
			metrics = append(metrics, currLI)
		}

	}

	// jsonStr, _ := json.Marshal(metrics, jsontext.WithIndent("  "))
	// os.WriteFile("main-citrix-machineMetrics-"+tenantName+".json", jsonStr, 0644)
	return metrics
}

//
//

func GetAllMachineResourceUtilisation(creds lib.CitrixCloudAccountConfig, tenantName string, tokenData lib.CitrixTokenData) []MachineResourceUtilisation {
	// formattedDate := time.Now().Add(-24 * time.Hour).Format("2006-01-02T15:04:05.000Z")
	// formattedDate := time.Now().Add(-11 * time.Hour).Format("2006-01-02T15:04:05.000Z")
	// date, err := time.Parse("2006-01-02T15:04:05.000Z", "2026-05-01T00:04:05.000Z")
	// lib.CheckFatalError(err)
	// formattedDate := date.Format("2006-01-02T15:04:05.000Z")

	formattedDate := time.Now().Add(-630 * time.Minute).Format("2006-01-02T15:04:05.000Z")
	filterString := "$filter=CollectedDate gt " + formattedDate + "&$orderby=CollectedDate desc"
	// filterString := "$orderby=CollectedDate desc"
	urlString := "https://api.cloud.com/monitorodata/ResourceUtilization?" + url.PathEscape(filterString)

	res, err := HttpGet(urlString, creds.CustomerId, creds.SiteId, tokenData)
	lib.CheckFatalError(err)

	// fmt.Println(string(res))

	var resData MachineResourceUtilisationResponse
	err = json.Unmarshal(res, &resData)
	lib.CheckFatalError(err)

	var resUtl []MachineResourceUtilisation

	for _, ru := range resData.Value {
		curr := ru
		curr.TenantName = tenantName
		resUtl = append(resUtl, curr)
	}

	nextLink := resData.NextLink
	iteration := 1
	for nextLink != "" {
		var curr MachineResourceUtilisationResponse
		iteration++
		fmt.Println(tenantName+" - ResUtil:", strconv.Itoa(iteration))

		resNew, _ := HttpGet(nextLink, creds.CustomerId, creds.SiteId, tokenData)
		json.Unmarshal(resNew, &curr)
		// fmt.Println(string(res))

		nextLink = curr.NextLink

		for _, ru := range curr.Value {
			currRU := ru
			currRU.TenantName = tenantName
			resUtl = append(resUtl, currRU)
		}

	}

	// jsonStr, _ := json.Marshal(resUtl, jsontext.WithIndent("  "))
	// os.WriteFile("main-citrix-resourceUtilisation-"+tenantName+".json", jsonStr, 0644)
	return resUtl
}

//
//

func GetAllMachineLoadIndexes(creds lib.CitrixCloudAccountConfig, tenantName string, tokenData lib.CitrixTokenData) []MachineLoadIndex {
	// formattedDate := time.Now().Add(-10 * time.Hour).Add(-30 * time.Minute).Format("2006-01-02T15:04:05.000Z")
	// formattedDate := time.Now().Add(-12 * time.Hour).Format("2006-01-02T15:04:05.000Z")

	// date, err := time.Parse("2006-01-02T15:04:05.000Z", "2026-05-01T00:04:05.000Z")
	// lib.CheckFatalError(err)
	// formattedDate := date.Format("2006-01-02T15:04:05.000Z")
	// formattedDate := time.Now().Add(-11 * time.Hour).Format("2006-01-02T15:04:05.000Z")
	// filterString := "$filter=CreatedDate gt 2026-04-01:04:05.000Z&$orderby=CreatedDate desc"

	formattedDate := time.Now().Add(-630 * time.Minute).Format("2006-01-02T15:04:05.000Z")
	filterString := "$filter=CreatedDate gt " + formattedDate + "&$orderby=CreatedDate desc"
	// filterString := "$orderby=CreatedDate desc"
	// filterString := ""
	urlString := "https://api.cloud.com/monitorodata/LoadIndexes?" + url.PathEscape(filterString)

	res, err := HttpGet(urlString, creds.CustomerId, creds.SiteId, tokenData)
	lib.CheckFatalError(err)

	// fmt.Println(string(res))
	// os.Exit(0)

	var resData MachineLoadIndexesResponse
	err = json.Unmarshal(res, &resData)
	lib.CheckFatalError(err)

	var loadIndexes []MachineLoadIndex

	for _, li := range resData.Value {
		curr := li
		curr.TenantName = tenantName
		loadIndexes = append(loadIndexes, curr)
	}

	nextLink := resData.NextLink

	iteration := 1
	for nextLink != "" {
		iteration++
		fmt.Println(tenantName+" - LoadIndexes:", strconv.Itoa(iteration))
		var curr MachineLoadIndexesResponse
		// fmt.Println(nextLink)

		resNew, _ := HttpGet(nextLink, creds.CustomerId, creds.SiteId, tokenData)
		json.Unmarshal(resNew, &curr)
		// fmt.Println(string(res))

		nextLink = curr.NextLink
		for _, li := range curr.Value {
			currLI := li
			currLI.TenantName = tenantName
			loadIndexes = append(loadIndexes, currLI)
		}

	}

	// jsonStr, _ := json.Marshal(loadIndexes, jsontext.WithIndent("  "))
	// os.WriteFile("main-citrix-loadIndexes-"+tenantName+".json", jsonStr, 0644)
	return loadIndexes
}

//
//

func GetAllMonitorMachines(creds lib.CitrixCloudAccountConfig, tenantName string, tokenData lib.CitrixTokenData) []MonitorMachine {
	// formattedDate := time.Now().Add(-10 * time.Hour).Add(-30 * time.Minute).Format("2006-01-02T15:04:05.000Z")
	// formattedDate := time.Now().Add(-12 * time.Hour).Format("2006-01-02T15:04:05.000Z")
	// filterString := "$filter=CreatedDate gt " + formattedDate + "&$orderby=CreatedDate desc"
	// filterString := "$filter=CurrentPowerState eq 3 and LifecycleState eq 0&$expand=CurrentLoadIndex&$select=CurrentPowerState,AgentVersion,IPAddress,DnsName,Id,IsInMaintenanceMode,AssociatedUserNames,CurrentSessionCount,FaultState,CurrentRegistrationState,Name,CurrentLoadIndexID"
	// filterString := "$expand=CurrentLoadIndex&$select=CurrentPowerState,LifecycleState,AgentVersion,IPAddress,DnsName,Id,IsInMaintenanceMode,FaultState,AssociatedUserNames,CurrentSessionCount,FaultState,CurrentRegistrationState,Name,CurrentLoadIndexID"

	filterString := "$select=CurrentPowerState,LifecycleState,AgentVersion,IPAddress,DnsName,Id,IsInMaintenanceMode,FaultState,AssociatedUserNames,CurrentSessionCount,FaultState,CurrentRegistrationState,Name,CurrentLoadIndexID"
	urlString := "https://api.cloud.com/monitorodata/Machines?" + url.PathEscape(filterString)

	res, err := HttpGet(urlString, creds.CustomerId, creds.SiteId, tokenData)
	lib.CheckFatalError(err)

	// fmt.Println(string(res))
	// os.Exit(0)

	var resData MonitorMachinesResponse
	err = json.Unmarshal(res, &resData)
	// lib.CheckFatalError(err)
	if err != nil {
		lib.JsonMarshalAndPrint(resData)
		lib.CheckFatalError(err)
	}

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
	// os.WriteFile("main-citrix-monitorMachines-"+tenantName+".json", jsonStr, 0644)
	return machines
}
