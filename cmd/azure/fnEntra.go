package azure

import (
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jercle/cloudini/lib"
)

func GetAppRegDataForAllConfiguredTenants(outputPath string) (allAppRegistrations []EntraApplication, expiringCreds []EntraExpiringCredential) {
	config := lib.GetCldConfig(nil)
	azConfTenants := config.Azure.MultiTenantAuth.Tenants
	tokenReq, err := GetAllTenantSPTokens(lib.AzureMultiAuthTokenRequestOptions{
		Scope:         "graph",
		GetWriteToken: true,
	}, nil)
	lib.CheckFatalError(err)

	for tenantName, tData := range azConfTenants {
		token, err := tokenReq.SelectTenant(tenantName)
		lib.CheckFatalError(err)

		entraApps := ListEntraAppRegistrations(token)
		var entraAppsProcessed []EntraApplication
		for _, app := range entraApps {
			currApp := app
			currApp.TenantId = tData.TenantID
			entraAppsProcessed = append(entraAppsProcessed, currApp)
		}
		allAppRegistrations = append(allAppRegistrations, entraAppsProcessed...)
		tenantExpiringCreds := GetAppRegExpiringCredData(entraAppsProcessed, 30)
		expiringCreds = append(expiringCreds, tenantExpiringCreds...)
	}

	if outputPath != "" {
		if _, err := os.Stat(outputPath); err != nil {
			os.MkdirAll(outputPath, os.ModePerm)
		}
		expiringCredsStr, _ := json.Marshal(expiringCreds, jsontext.WithIndent("  "))
		os.WriteFile(outputPath+"/entraApps-ALL-expiringCredentials.json", expiringCredsStr, 0644)

		allAppRegistrationsStr, _ := json.Marshal(allAppRegistrations, jsontext.WithIndent("  "))
		os.WriteFile(outputPath+"/entraApps-ALL-allAppRegistrations.json", allAppRegistrationsStr, 0644)
	}

	return
}

//
//

func ListEntraAppRegistrations(token *lib.AzureMultiAuthToken) []EntraApplication {
	var (
		resData   EntraListApplicationsResponse
		entraApps []EntraApplication
		nextLink  string
	)

	urlString := "https://graph.microsoft.com/beta/applications"

	res, err := HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	json.Unmarshal(res, &resData)

	for _, app := range resData.Value {
		currApp := app
		currApp.LastAzureSync = time.Now()
		currApp.TenantName = token.TenantName
		entraApps = append(entraApps, currApp)
	}

	// entraApps = append(entraApps, resData.Value...)

	nextLink = resData.NextLink
	// count := 1
	// fmt.Println(string(res))
	// os.WriteFile("outputs/entraApps-"+strconv.Itoa(count)+".json", res, 0644)

	for nextLink != "" {
		// count++
		var currentSet EntraListApplicationsResponse
		response, _ := HttpGet(nextLink, *token)
		// os.WriteFile("outputs/entraApps-"+strconv.Itoa(count)+".json", response, 0644)
		json.Unmarshal(response, &currentSet)
		nextLink = currentSet.NextLink
		// entraApps = append(entraApps, currentSet.Value...)
		for _, app := range currentSet.Value {
			currApp := app
			currApp.LastAzureSync = time.Now()
			currApp.TenantName = token.TenantName
			// currApp.PortalUrl = "https://portal.azure.com/#view/Microsoft_AAD_RegisteredApps/ApplicationMenuBlade/~/Overview/appId/" + app.AppID + "/isMSAApp~/false"
			entraApps = append(entraApps, currApp)
		}
	}

	return entraApps
}

//
//

func ListEntraRoleDefinitions(mat lib.AzureMultiAuthToken) ([]EntraRoleDefinition, error) {
	var (
		unmarshResponse ListEntraRoleDefinitionsResponse
		roleDefs        []EntraRoleDefinition
	)

	urlString := "https://graph.microsoft.com/v1.0/roleManagement/directory/roleDefinitions"

	response, err := HttpGet(urlString, mat)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &unmarshResponse)
	if err != nil {
		return nil, err
	}

	jsonData, _ := json.Marshal(unmarshResponse.Value)

	err = json.Unmarshal(jsonData, &roleDefs)
	if err != nil {
		return nil, err
	}

	return roleDefs, nil
}

//
//

func GetAllB2CTenantUsers() (users []B2CUserMinimal) {

	tokenReq, err := GetAllTenantSPTokens(lib.AzureMultiAuthTokenRequestOptions{
		Scope:         "graph",
		GetWriteToken: true,
	}, nil)
	lib.CheckFatalError(err)

	config := lib.GetCldConfig(nil)
	tenants := config.Azure.MultiTenantAuth.Tenants

	for tName, tConfig := range tenants {
		if tConfig.IsB2C {
			var nextLink string
			token, err := tokenReq.SelectTenant(tName)
			lib.CheckFatalError(err)
			urlString := "https://graph.microsoft.com/beta/users"
			res, err := HttpGet(urlString, *token)
			lib.CheckFatalError(err)

			var resData GetAllB2CTenantUsersResponse
			err = json.Unmarshal(res, &resData)
			lib.CheckFatalError(err)
			nextLink = resData.NextLink

			for _, user := range resData.Value {
				var userData B2CUserMinimal
				jsonStr, _ := json.Marshal(user)
				err := json.Unmarshal(jsonStr, &userData)
				lib.CheckFatalError(err)
				upn := userData.UserPrincipalName
				domain := strings.Split(upn, "@")
				tenant := strings.Split(domain[1], ".")[0]
				userData.B2CTenant = tenant

				var unknownFields map[string]time.Time
				err = json.Unmarshal(user.UnknownFields, &unknownFields)
				for k, v := range unknownFields {
					if strings.HasPrefix(k, "extension") && strings.HasSuffix(k, "lastLogonTime") {
						userData.ExtensionLastLogonTime = v
					}
					if strings.HasPrefix(k, "extension") && strings.HasSuffix(k, "passwordResetOn") {
						userData.ExtensionPasswordResetOn = v
					}
				}

				users = append(users, userData)
			}

			for nextLink != "" {
				// count++
				var currentSet GetAllB2CTenantUsersResponse
				response, _ := HttpGet(nextLink, *token)
				// os.WriteFile("outputs/entraApps-"+strconv.Itoa(count)+".json", response, 0644)
				json.Unmarshal(response, &currentSet)
				nextLink = currentSet.NextLink
				// entraApps = append(entraApps, currentSet.Value...)
				for _, user := range currentSet.Value {
					jsonStr, _ := json.Marshal(user)
					var currUser B2CUserMinimal
					err := json.Unmarshal(jsonStr, &currUser)
					lib.CheckFatalError(err)
					upn := currUser.UserPrincipalName
					domain := strings.Split(upn, "@")
					tenant := strings.Split(domain[1], ".")[0]
					currUser.B2CTenant = tenant
					// currUser.TenantName = tName
					users = append(users, currUser)
				}
			}
		}
	}

	// jsonStr, _ := json.Marshal(users)

	// os.WriteFile("/home/jercle/git/cld/b2cusers.json", jsonStr, 0644)
	return
}

//
//

func GetB2CUserByUPN(upn string, token *lib.AzureMultiAuthToken) (user interface{}) {
	urlString := "https://graph.microsoft.com/beta/users/" + upn
	res, err := HttpGet(urlString, *token)
	lib.CheckFatalError(err)
	err = json.Unmarshal(res, &user)
	lib.CheckFatalError(err)
	return
}

//
//

func GetB2CUserByObjectId(objectId string, token *lib.AzureMultiAuthToken) (user interface{}) {
	urlString := "https://graph.microsoft.com/beta/users/" + objectId
	res, err := HttpGet(urlString, *token)
	lib.CheckFatalError(err)
	err = json.Unmarshal(res, &user)
	lib.CheckFatalError(err)
	return
}

//
//

func GetAppRegExpiringCredData(apps []EntraApplication, daysUntilExpired int) (expiringCredentials []EntraExpiringCredential) {
	for _, app := range apps {

		currApp := app
		if app.PasswordCredentials != nil {
			for _, cred := range *currApp.PasswordCredentials {
				daysUntilExpiry := time.Until(cred.EndDateTime).Hours() / 24
				if int(daysUntilExpiry) < daysUntilExpired {
					var curr EntraExpiringCredential
					curr.AppRegAppID = app.AppID
					curr.AppRegCreatedDateTime = app.CreatedDateTime
					curr.AppRegDescription = app.Description
					curr.AppRegDisplayName = app.DisplayName
					curr.AppRegObjectID = app.ID
					curr.CredDisplayName = cred.DisplayName
					curr.CredEndDateTime = cred.EndDateTime
					curr.CredID = cred.KeyID
					curr.CredCustomKeyIdentifier = cred.CustomKeyIdentifier
					curr.CredStartDateTime = cred.StartDateTime
					curr.TenantName = app.TenantName
					curr.TenantId = app.TenantId
					curr.LastAzureSync = app.LastAzureSync

					curr.CredType = "pwd"
					curr.MongoDbId = app.TenantName + "_" + curr.CredType + "_" + curr.CredID + "_" + app.AppID

					expiringCredentials = append(expiringCredentials, curr)
				}
			}
		}

		if app.KeyCredentials != nil {
			for _, cred := range *currApp.KeyCredentials {
				daysUntilExpiry := time.Until(cred.EndDateTime).Hours() / 24
				if int(daysUntilExpiry) < daysUntilExpired {
					var curr EntraExpiringCredential
					curr.AppRegAppID = app.AppID
					curr.AppRegCreatedDateTime = app.CreatedDateTime
					curr.AppRegDescription = app.Description
					curr.AppRegDisplayName = app.DisplayName
					curr.AppRegObjectID = app.ID
					curr.CredDisplayName = cred.DisplayName
					curr.CredEndDateTime = cred.EndDateTime
					curr.CredID = cred.KeyID
					curr.CredCustomKeyIdentifier = cred.CustomKeyIdentifier
					curr.CredStartDateTime = cred.StartDateTime
					curr.CredKeyType = cred.Type
					curr.CredKeyUsage = cred.Usage
					curr.TenantName = app.TenantName
					curr.LastAzureSync = app.LastAzureSync

					curr.CredType = "key"
					curr.MongoDbId = app.TenantName + "_" + curr.CredType + "_" + curr.CredID + "_" + app.AppID

					expiringCredentials = append(expiringCredentials, curr)
				}
			}
		}
	}
	return
}

//
//

func ListAllTenantPIMScheduleInstances(token *lib.AzureMultiAuthToken) (assignmentsSlice []RoleAssignmentScheduleInstance, eligibilitiesSlice []RoleEligibilityScheduleInstance) {
	var (
		aMutex sync.Mutex
		eMutex sync.Mutex
		mgWg   sync.WaitGroup
		sWg    sync.WaitGroup
		rgWg   sync.WaitGroup
		// aMutex        sync.Mutex
	)
	config := lib.GetCldConfig(nil)
	_ = config
	// var allTenants []lib.CldConfigTenantAuth
	assignments := make(map[string]RoleAssignmentScheduleInstance)
	eligibilities := make(map[string]RoleEligibilityScheduleInstance)

	mgmtGroups, err := ListManagementGroups(token)
	_ = mgmtGroups
	lib.CheckFatalError(err)
	// if err != nil {
	// 	return nil, err
	// }

	subs, err := ListSubscriptions(*token)
	_ = subs
	lib.CheckFatalError(err)

	resGrps := GetAllTenantResourceGroups("", token)

	for _, rg := range resGrps {
		rgWg.Add(1)
		go func() {
			defer rgWg.Done()
			// fmt.Println(sub.DisplayName, ": Sub ListRoleAssignmentScheduleInstances")
			a, err := ListRoleAssignmentScheduleInstances(rg.ID, token)
			lib.CheckFatalError(err)

			// fmt.Println(rg.Name, ": RG ListRoleAssignmentScheduleInstances - ", strconv.Itoa(len(a)))
			for _, inst := range a {
				curr := inst
				curr.TenantName = token.TenantName
				curr.LastAzureSync = time.Now()
				aMutex.Lock()
				assignments[inst.ID] = curr
				aMutex.Unlock()
			}
		}()
		rgWg.Add(1)
		go func() {
			defer rgWg.Done()
			e, err := ListRoleEligibilityScheduleInstances(rg.ID, token)
			lib.CheckFatalError(err)
			// fmt.Println(rg.Name, ": RG ListRoleEligibilityScheduleInstances - ", strconv.Itoa(len(e)))
			for _, inst := range e {
				curr := inst
				curr.TenantName = token.TenantName
				curr.LastAzureSync = time.Now()
				eMutex.Lock()
				eligibilities[inst.ID] = curr
				eMutex.Unlock()
			}
		}()
	}
	rgWg.Wait()

	for _, sub := range subs {
		sWg.Add(1)
		go func() {
			defer sWg.Done()
			// fmt.Println(sub.DisplayName, ": Sub ListRoleAssignmentScheduleInstances")
			a, err := ListRoleAssignmentScheduleInstances(sub.ID, token)
			lib.CheckFatalError(err)
			// fmt.Println(sub.DisplayName, ": Sub ListRoleAssignmentScheduleInstances - ", strconv.Itoa(len(a)))
			for _, inst := range a {
				curr := inst
				curr.TenantName = token.TenantName
				curr.LastAzureSync = time.Now()
				aMutex.Lock()
				assignments[inst.ID] = curr
				aMutex.Unlock()
			}
		}()
		sWg.Add(1)
		go func() {
			defer sWg.Done()
			e, err := ListRoleEligibilityScheduleInstances(sub.ID, token)
			lib.CheckFatalError(err)
			// fmt.Println(sub.DisplayName, ": Sub ListRoleEligibilityScheduleInstances - ", strconv.Itoa(len(e)))
			for _, inst := range e {
				curr := inst
				curr.TenantName = token.TenantName
				curr.LastAzureSync = time.Now()
				eMutex.Lock()
				eligibilities[inst.ID] = curr
				eMutex.Unlock()
			}
		}()
	}
	sWg.Wait()

	for _, mg := range mgmtGroups {
		mgWg.Add(1)
		go func() {
			defer mgWg.Done()
			// fmt.Println(mg.Name, ": Mgmt Grp ListRoleAssignmentScheduleInstances")
			a, err := ListRoleAssignmentScheduleInstances(mg.ID, token)
			lib.CheckFatalError(err)
			// fmt.Println(mg.Name, ": Mgmt Grp ListRoleAssignmentScheduleInstances - ", strconv.Itoa(len(a)))
			for _, inst := range a {
				curr := inst
				curr.TenantName = token.TenantName
				curr.LastAzureSync = time.Now()
				aMutex.Lock()
				assignments[inst.ID] = curr
				aMutex.Unlock()
			}
		}()
		mgWg.Add(1)
		go func() {
			defer mgWg.Done()
			e, err := ListRoleEligibilityScheduleInstances(mg.ID, token)
			lib.CheckFatalError(err)
			// fmt.Println(mg.Name, ": Mgmt Grp ListRoleEligibilityScheduleInstances - ", strconv.Itoa(len(e)))

			for _, inst := range e {
				curr := inst
				curr.TenantName = token.TenantName
				curr.LastAzureSync = time.Now()
				eMutex.Lock()
				eligibilities[inst.ID] = curr
				eMutex.Unlock()
			}
		}()
	}
	mgWg.Wait()

	for _, item := range assignments {
		assignmentsSlice = append(assignmentsSlice, item)
	}

	for _, item := range eligibilities {
		eligibilitiesSlice = append(eligibilitiesSlice, item)
	}

	return
}

//
//

func ListAllTenantPIMScheduleInstancesForAllTenants() (processedAssignments []RoleAssignmentScheduleInstance, processedEligibilities []RoleEligibilityScheduleInstance) {
	var (
		assignments   []RoleAssignmentScheduleInstance
		eligibilities []RoleEligibilityScheduleInstance
		// aMutex        sync.Mutex
	)
	config := lib.GetCldConfig(nil)
	_ = config

	// os.Exit(0)

	tokenReq, err := GetAllTenantSPTokens(lib.AzureMultiAuthTokenRequestOptions{
		// Scope:         "graph",
		GetWriteToken: true,
	}, nil)
	lib.CheckFatalError(err)

	for _, token := range tokenReq {
		tenantAssignments, tenantEligibilities := ListAllTenantPIMScheduleInstances(&token)
		assignments = append(assignments, tenantAssignments...)
		eligibilities = append(eligibilities, tenantEligibilities...)
	}

	asgnMap := make(map[string]RoleAssignmentScheduleInstance)
	for _, item := range assignments {
		asgnMap[item.ID] = item
	}

	eligMap := make(map[string]RoleEligibilityScheduleInstance)
	for _, item := range eligibilities {
		eligMap[item.ID] = item
	}

	for _, item := range asgnMap {
		processedAssignments = append(processedAssignments, item)
	}

	for _, item := range eligMap {
		processedEligibilities = append(processedEligibilities, item)
	}

	return
}

//
//

func ListRoleEligibilityScheduleInstances(scope string, token *lib.AzureMultiAuthToken) ([]RoleEligibilityScheduleInstance, error) {
	var (
		response ListRoleEligibilityScheduleInstancesResponse
	)
	urlString := "https://management.azure.com/" +
		scope +
		"/providers/Microsoft.Authorization/roleEligibilityScheduleInstances?api-version=2020-10-01"
	res, err := HttpGet(urlString, *token)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(res, &response)
	// fmt.Println(string(res))

	return response.Value, nil
}

//
//

func ListRoleAssignmentScheduleInstances(scope string, token *lib.AzureMultiAuthToken) ([]RoleAssignmentScheduleInstance, error) {
	var (
		// response ListRoleAssignmentScheduleInstancesResponse
		response ListRoleAssignmentScheduleInstancesResponse
	)
	urlString := "https://management.azure.com/" +
		scope +
		"/providers/Microsoft.Authorization/roleAssignmentScheduleInstances?api-version=2020-10-01"
	res, err := HttpGet(urlString, *token)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(res, &response)
	// fmt.Println(string(res))

	return response.Value, nil
}

//
//

func GetAllEntraUsersForTenant(token *lib.AzureMultiAuthToken, selects *[]string, apiVersion *string) (users []EntraUser) {
	var (
		response GetAllEntraUsersForTenantResponse
		nextLink *string
	)

	baseGraphUrl := "https://graph.microsoft.com/"
	apiVers := "v1.0"
	if apiVersion != nil {
		apiVers = *apiVersion
	}
	endpoint := "/users"
	if selects != nil {
		endpoint += "?$select=" + strings.Join(*selects, ",")
	}
	urlString := baseGraphUrl + apiVers + endpoint

	res, resErr := HttpGet(urlString, *token)
	lib.CheckHttpGetError(resErr)
	err := json.Unmarshal(res, &response)
	lib.CheckFatalError(err)

	users = append(users, response.Value...)

	nextLink = response.NextLink

	for nextLink != nil {
		var currentSet GetAllEntraUsersForTenantResponse
		res, _ := HttpGet(*nextLink, *token)
		json.Unmarshal(res, &currentSet)
		nextLink = currentSet.NextLink
		users = append(users, currentSet.Value...)
	}

	return
}
