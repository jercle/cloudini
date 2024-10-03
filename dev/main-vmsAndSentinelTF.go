package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func main() {
	startTime := time.Now()
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

	// subscriptionId := """
	// resourceGroupName := ""
	// workspaceName := ""
	// _ = subscriptionId
	// _ = resourceGroupName
	// _ = workspaceName
	// fmt.Println(string(res))

	// urlString := "https://management.azure.com/subscriptions/" +
	// subscriptionId +
	// "/providers/Microsoft.Compute/virtualMachines?api-version=2024-07-01"

	// res, err := azure.HttpGet(urlString, *token)
	// lib.CheckFatalError(err)

	// fmt.Println(string(res))
	// TerraformGenerateAllEnabledSentinelAlertRules(subscriptionId, resourceGroupName, workspaceName, token, "./imports.tf")
	// rules := GetAllEnabledSentinelAlertRules(subscriptionId, resourceGroupName, workspaceName, token)

	// jsonStr, _ := json.MarshalIndent(rules, "", "  ")

	allVirtualMachines := make(map[string][]VirtualMachine)

	// config.AddAzureTenant("e9f4bce2-7308-461a-91ce-3213f50f54f1", "FakeTest")
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{}, nil)
	// fmt.Println(tokens)
	// lib.CheckFatalError(err)
	// ListAllVirtualMachinesInSubscription(subscriptionId, token)
	for _, token := range tokenReq {
		subs, err := azure.ListSubscriptions(token)
		// fmt.Println(token)
		lib.CheckFatalError(err)
		for _, sub := range subs {
			// lib.PrintSrcLoc("getting kvs")
			// GetAllSubscriptionKeyvaults(tenant, sub.ID, sub.DisplayName)
			subVirtualMachines := ListAllVirtualMachinesInSubscription(sub.SubscriptionID, &token, true)
			// fmt.Println(subVirtualMachines)
			allVirtualMachines[sub.DisplayName] = subVirtualMachines
			jsonStr, _ := json.MarshalIndent(allVirtualMachines[sub.DisplayName], "", "  ")
			os.WriteFile("./outputs/virtualMachines/"+sub.TenantName+"/"+sub.DisplayName+".json", jsonStr, 0644)
		}
	}

	// fmt.Println(string(jsonStr))
	// virtualMachines := ListAllVirtualMachinesInSubscription(subscriptionId, token)

	jsonStr, _ := json.MarshalIndent(allVirtualMachines, "", "  ")
	fmt.Println(string(jsonStr))
	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}

func ListAllVirtualMachinesInSubscription(subscriptionId string, token *lib.MultiAuthToken, filterCitrixVMs bool) []VirtualMachine {
	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/providers/Microsoft.Compute/virtualMachines?api-version=2024-07-01"
	res, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	var listVirtualMachinesResponse ListAllVirtualMachinesInSubscriptionResponse

	json.Unmarshal(res, &listVirtualMachinesResponse)

	if filterCitrixVMs {
		var filteredList []VirtualMachine

		for _, vm := range listVirtualMachinesResponse.Value {

			if _, ok := vm.Tags["CitrixResource"]; !ok {
				filteredList = append(filteredList, vm)
			}
		}
		return filteredList
	} else {
		return listVirtualMachinesResponse.Value
	}

	// fmt.Println(listVirtualMachinesResponse)

}

func TerraformGenerateAllEnabledSentinelAlertRules(subscriptionId string, resourceGroupName string, workspaceName string, token *lib.MultiAuthToken, filePath string) {
	enabledAlertRules := GetAllEnabledSentinelAlertRules(subscriptionId, resourceGroupName, workspaceName, token)
	importBlocks := ""
	var resourceNames []string
	// fmt.Println(len(enabledAlertRules))
	for _, rule := range enabledAlertRules {
		processedName := ProcessNamesToTerraformReference(rule.Properties.DisplayName)
		checkedName := CheckResourceNameIsUnique(&resourceNames, processedName)
		resourceNames = append(resourceNames, checkedName)
		entry := "\nimport {\n" +
			"  id = \"" +
			rule.ID +
			"\"\n  to = azurerm_sentinel_alert_rule_scheduled." +
			checkedName +
			"\n}\n\n"

		importBlocks += entry
	}

	importBlocks = strings.TrimSpace(importBlocks)
	SaveImportBlocks(importBlocks, filePath)
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
	byteData := []byte(fileData)
	err := os.WriteFile(filePath, byteData, 0644)
	lib.CheckFatalError(err)
}

func GetAllEnabledSentinelAlertRules(subscriptionID string, resourceGroupName string, workspaceName string, token *lib.MultiAuthToken) []SentinelAlertRule {
	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionID +
		"/resourceGroups/" +
		resourceGroupName +
		"/providers/Microsoft.OperationalInsights/workspaces/" +
		workspaceName +
		"/providers/Microsoft.SecurityInsights/alertRules?api-version=2024-03-01"
	res, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	var rsp ListSentinelAlertRulesResponse

	json.Unmarshal(res, &rsp)

	var enabledAlertRules []SentinelAlertRule

	for _, rule := range rsp.Value {
		if rule.Properties.Enabled && rule.Kind != "MLBehaviorAnalytics" && rule.Kind != "Fusion" {
			enabledAlertRules = append(enabledAlertRules, rule)
		}
	}

	return enabledAlertRules
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
	processedName = strings.ReplaceAll(processedName, ".", "")
	processedName = strings.ReplaceAll(processedName, "__", "_")
	processedName = strings.ReplaceAll(processedName, "__", "_")
	processedName = strings.ToLower(processedName)
	return processedName
}

type ListSentinelAlertRulesResponse struct {
	Value []SentinelAlertRule `json:"value"`
}

type ListAllVirtualMachinesInSubscriptionResponse struct {
	Value []VirtualMachine `json:"value"`
}

type VirtualMachine struct {
	Etag     string `json:"etag"`
	ID       string `json:"id"`
	Identity struct {
		PrincipalID string `json:"principalId"`
		TenantID    string `json:"tenantId"`
		Type        string `json:"type"`
	} `json:"identity"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		DiagnosticsProfile struct {
			BootDiagnostics struct {
				Enabled bool `json:"enabled"`
			} `json:"bootDiagnostics"`
		} `json:"diagnosticsProfile"`
		HardwareProfile struct {
			VmSize string `json:"vmSize"`
		} `json:"hardwareProfile"`
		LicenseType    string `json:"licenseType"`
		NetworkProfile struct {
			NetworkInterfaces []struct {
				ID         string `json:"id"`
				Properties struct {
					Primary bool `json:"primary"`
				} `json:"properties"`
			} `json:"networkInterfaces"`
		} `json:"networkProfile"`
		OSProfile struct {
			WindowsConfiguration struct {
				PatchSettings struct {
					AssessmentMode              string `json:"assessmentMode"`
					AutomaticByPlatformSettings struct {
						BypassPlatformSafetyChecksOnUserSchedule bool `json:"bypassPlatformSafetyChecksOnUserSchedule"`
					} `json:"automaticByPlatformSettings"`
					PatchMode string `json:"patchMode"`
				} `json:"patchSettings"`
			} `json:"windowsConfiguration"`
		} `json:"osProfile"`
		ProvisioningState string `json:"provisioningState"`
		StorageProfile    struct {
			DataDisks []struct {
				Caching      string  `json:"caching"`
				CreateOption string  `json:"createOption"`
				DeleteOption string  `json:"deleteOption"`
				DiskSizeGb   float64 `json:"diskSizeGB"`
				Lun          float64 `json:"lun"`
				ManagedDisk  struct {
					ID                 string `json:"id"`
					StorageAccountType string `json:"storageAccountType"`
				} `json:"managedDisk"`
				Name         string `json:"name"`
				ToBeDetached bool   `json:"toBeDetached"`
			} `json:"dataDisks"`
			DiskControllerType string `json:"diskControllerType"`
			OSDisk             struct {
				Caching      string  `json:"caching"`
				CreateOption string  `json:"createOption"`
				DeleteOption string  `json:"deleteOption"`
				DiskSizeGb   float64 `json:"diskSizeGB"`
				ManagedDisk  struct {
					ID                 string `json:"id"`
					StorageAccountType string `json:"storageAccountType"`
				} `json:"managedDisk"`
				Name   string `json:"name"`
				OSType string `json:"osType"`
			} `json:"osDisk"`
		} `json:"storageProfile"`
		TimeCreated time.Time `json:"timeCreated"`
		VmID        string    `json:"vmId"`
	} `json:"properties"`
	Resources []struct {
		ID string `json:"id"`
	} `json:"resources"`
	Tags  map[string]string `json:"tags"`
	Type  string            `json:"type"`
	Zones []string          `json:"zones"`
}

type SentinelAlertRule struct {
	Etag       string `json:"etag"`
	ID         string `json:"id"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Properties struct {
		AlertDetailsOverride *struct {
			AlertDescriptionFormat string `json:"alertDescriptionFormat,omitempty"`
			AlertDisplayNameFormat string `json:"alertDisplayNameFormat,omitempty"`
			AlertDynamicProperties []any  `json:"alertDynamicProperties"`
		} `json:"alertDetailsOverride,omitempty"`
		AlertRuleTemplateName *string `json:"alertRuleTemplateName"`
		CustomDetails         *struct {
			DnsQueries        string `json:"DNSQueries,omitempty"`
			DnsQueryCount     string `json:"DNSQueryCount,omitempty"`
			DnsQueryThreshold string `json:"DNSQueryThreshold,omitempty"`
			DnsQuerythreshold string `json:"DNSQuerythreshold,omitempty"`
			NxdomainCount     string `json:"NXDOMAINCount,omitempty"`
			NxdomaiNthreshold string `json:"NXDOMAINthreshold,omitempty"`
			ProcessName       string `json:"ProcessName,omitempty"`
			SubjectUserName   string `json:"SubjectUserName,omitempty"`
			TimeEnabled       string `json:"TimeEnabled,omitempty"`
		} `json:"customDetails,omitempty"`
		Description               string `json:"description"`
		DisplayName               string `json:"displayName"`
		DisplayNamesExcludeFilter any    `json:"displayNamesExcludeFilter,omitempty"`
		DisplayNamesFilter        any    `json:"displayNamesFilter,omitempty"`
		Enabled                   bool   `json:"enabled"`
		EntityMappings            []struct {
			EntityType    string `json:"entityType"`
			FieldMappings []struct {
				ColumnName string `json:"columnName"`
				Identifier string `json:"identifier"`
			} `json:"fieldMappings"`
		} `json:"entityMappings,omitempty"`
		EventGroupingSettings *struct {
			AggregationKind string `json:"aggregationKind"`
		} `json:"eventGroupingSettings,omitempty"`
		IncidentConfiguration *struct {
			CreateIncident        bool `json:"createIncident"`
			GroupingConfiguration struct {
				Enabled              bool     `json:"enabled"`
				GroupByAlertDetails  []string `json:"groupByAlertDetails"`
				GroupByCustomDetails []any    `json:"groupByCustomDetails"`
				GroupByEntities      []string `json:"groupByEntities"`
				LookbackDuration     string   `json:"lookbackDuration"`
				MatchingMethod       string   `json:"matchingMethod"`
				ReopenClosedIncident bool     `json:"reopenClosedIncident"`
			} `json:"groupingConfiguration"`
		} `json:"incidentConfiguration,omitempty"`
		LastModifiedUtc     time.Time `json:"lastModifiedUtc"`
		ProductFilter       string    `json:"productFilter,omitempty"`
		Query               string    `json:"query,omitempty"`
		QueryFrequency      string    `json:"queryFrequency,omitempty"`
		QueryPeriod         string    `json:"queryPeriod,omitempty"`
		SeveritiesFilter    []string  `json:"severitiesFilter,omitempty"`
		Severity            string    `json:"severity,omitempty"`
		SuppressionDuration string    `json:"suppressionDuration,omitempty"`
		SuppressionEnabled  bool      `json:"suppressionEnabled"`
		Tactics             []string  `json:"tactics"`
		Techniques          []string  `json:"techniques"`
		TemplateVersion     string    `json:"templateVersion,omitempty"`
		TriggerOperator     string    `json:"triggerOperator,omitempty"`
		TriggerThreshold    float64   `json:"triggerThreshold"`
	} `json:"properties"`
	Type string `json:"type"`
}
