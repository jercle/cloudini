package azure

import (
	json "encoding/json/v2"
	"fmt"
	"os"
	"strings"

	"github.com/jercle/cloudini/lib"
)

func GetResourceRoleAssignments(scope string, token *lib.AzureMultiAuthToken) []RoleAssignment {
	urlBase := "https://management.azure.com/" +
		scope +
		"/providers/Microsoft.Authorization/roleAssignments?api-version=2022-04-01"
		// "/providers/Microsoft.Authorization/roleAssignments?api-version=2022-04-01&$filter={filter}"
	urlString := urlBase + ""
	res, err := HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	var resData RoleAssignmentListResponse
	err = json.Unmarshal(res, &resData)

	return resData.Value
}

//
//

func GetAllResourceRoleAssignmentsInTenant(configuredTenantName string) (roleAssignments []RoleAssignment) {
	token, err := GetTenantSPToken(lib.AzureMultiAuthTokenRequestOptions{
		TenantName: configuredTenantName,
	}, nil)
	lib.CheckFatalError(err)

	fmt.Println("Fetching all resources")
	resources := GetAllTenantResources("", token)

	fmt.Println("Fetching all users")

	roleAssignmentsMap := make(map[string]RoleAssignment)

	fetchCount := 1

	var allResources []lib.AzureResourceDetails

	for _, s := range resources.Subscriptions {
		for _, r := range s.Resources {
			allResources = append(allResources, r)
		}
	}

	bar := lib.ProgressBar(len(allResources), "resource", 1, 1, "Fetching role assignments for all resources")

	for _, r := range allResources {
		// fmt.Println("Fetch " + strconv.Itoa(fetchCount) + ": " + r.ID)
		// bar.Describe("Fetching: " + r.Name + "\n")
		resRoleAssignments := GetResourceRoleAssignments(r.ID, token)
		fetchCount++

		for _, ra := range resRoleAssignments {
			curr := ra
			// fmt.Println("Iterating...", curr.Properties.Scope)

			scopeSplit := strings.Split(curr.Properties.Scope, "/")
			// lib.PrintSliceStringsWithIndexes(scopeSplit)
			// os.Exit(0)
			if strings.Contains(strings.ToLower(curr.Properties.Scope), "providers/microsoft.management/managementgroups") || curr.Properties.Scope == "/" {
				curr.ScopeType = "Management Group"
			} else if strings.Contains(strings.ToLower(curr.Properties.Scope), "/resourcegroups/") {
				curr.ScopeType = "Resource Group"
			} else if strings.Contains(strings.ToLower(curr.Properties.Scope), "subscriptions/") && len(scopeSplit) == 3 {
				curr.ScopeType = "Subscription"
			} else {
				curr.ScopeType = "Unknown"
			}
			// fmt.Println(curr.ScopeType)
			// fmt.Println(strings.ToLower(curr.Properties.Scope))
			if curr.ScopeType == "Unknown" {
				// fmt.Println(strings.ToLower(curr.Properties.Scope))
				lib.JsonMarshalAndPrint(curr)
				fmt.Println("Scope type unknown")
				os.Exit(1)
			}

			roleAssignmentsMap[r.ID] = curr

			// if fetchCount > 9 {
			// 	lib.JsonMarshalAndPrint(roleAssignmentsMap)
			// 	os.Exit(0)
			// }
		}

		bar.Add(1)
	}

	for _, ra := range roleAssignmentsMap {
		roleAssignments = append(roleAssignments, ra)
	}

	return
}
