package m365

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func GetM365LicenseCountsForTenant(currentDate *time.Time, token *lib.AzureMultiAuthToken) (m365LicenseCounts M365LicenseCounts) {
	if currentDate != nil {
		m365LicenseCounts.DateDataFetched = *currentDate
	} else {
		m365LicenseCounts.DateDataFetched = time.Now().Local()
	}
	m365LicenseCounts.TenantId = token.TenantId
	m365LicenseCounts.TenantName = token.TenantName
	countsBySku := make(M365LicenseCountBySkuPartNumber)
	skuPartNumbersBySkuId := make(map[string]string)

	users := GetM365UsersAssignedLicenses(token)
	licenseSkus := GetM365LicenseSkus(token)

	for _, sku := range licenseSkus {
		var currSku M365LicenseCount
		currSku.ConsumedUnits = sku.ConsumedUnits
		currSku.SkuId = sku.SkuID
		countsBySku[sku.SkuPartNumber] = currSku
		skuPartNumbersBySkuId[sku.SkuID] = sku.SkuPartNumber
	}

	for _, user := range users {
		currUser := user
		currUser.AssignedLicenses = []string{}
		for _, licenseSkuId := range user.AssignedLicenses {
			skuPartNum := skuPartNumbersBySkuId[licenseSkuId]
			currSku := countsBySku[skuPartNum]
			currUser.AssignedLicenses = append(currUser.AssignedLicenses, skuPartNum)
			currSku.ConsumedUnitsChecked++
			if user.IsActiveUser {
				currSku.UsersActiveCount++
			}
			if !user.AccountEnabled {
				currSku.UsersDisabledCount++
			}
			if user.NoLogonLast45Days {
				currSku.UsersNoLogonLast45DaysCount++
			}
			if user.NoLogonLast90Days {
				currSku.UsersNoLogonLast90DaysCount++
			}

			if user.OnPremisesSamAccountName == "" {
				if strings.HasPrefix(strings.ToLower(user.UserPrincipalName), "drew") {
					upnSplit := strings.Split(user.UserPrincipalName, "@")
					currSku.Users = append(currSku.Users, "drew@"+upnSplit[1])
				} else {
					currSku.Users = append(currSku.Users, user.UserPrincipalName)
				}
			} else {
				currSku.Users = append(currSku.Users, user.OnPremisesSamAccountName)
			}
			countsBySku[skuPartNum] = currSku
		}
		m365LicenseCounts.Users = append(m365LicenseCounts.Users, user)
		m365LicenseCounts.UsersCount++
		if user.NoLogonLast45Days {
			m365LicenseCounts.UsersNoLogonLast45DaysCount++
		} else if user.NoLogonLast90Days {
			m365LicenseCounts.UsersNoLogonLast90DaysCount++
		} else if user.IsActiveUser {
			m365LicenseCounts.UsersActiveCount++
		}
		if !user.AccountEnabled {
			m365LicenseCounts.UsersDisabledCount++
		}
	}

	m365LicenseCounts.UsersCountChecked = m365LicenseCounts.UsersNoLogonLast45DaysCount + m365LicenseCounts.UsersNoLogonLast90DaysCount + m365LicenseCounts.UsersActiveCount
	m365LicenseCounts.LicenseCountsBySkuPartNumber = countsBySku

	return
}

//
//

func GetM365LicenseCountsForAllConfiguredTenants(options *lib.GetAllM365LicenseCountsForAllConfiguredTenantsOptions) (m365LicenseCounts map[string]M365LicenseCounts) {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	var opts lib.GetAllM365LicenseCountsForAllConfiguredTenantsOptions
	if options != nil {
		opts = *options
	}

	cldConfig := lib.GetCldConfig(nil)
	azConfigs := cldConfig.Azure.MultiTenantAuth.Tenants

	m365LicenseCounts = make(map[string]M365LicenseCounts)
	dateDataFetched := time.Now().Local()

	tokens, err := azure.GetAllTenantSPTokens(lib.AzureMultiAuthTokenRequestOptions{
		Scope:         "graph",
		GetWriteToken: false,
	}, nil)
	lib.CheckFatalError(err)

	for _, token := range tokens {
		if azConfigs[token.TenantName].IsB2C {
			continue
		}
		wg.Go(func() {
			if !opts.SuppressSteps {
				fmt.Println(token.TenantName + ": Fetching license counts")
			}
			// allResources[token.TenantName] = make(map[string]SubscriptionResourceList)
			tenantLicenseCounts := GetM365LicenseCountsForTenant(&dateDataFetched, &token)
			if !opts.SuppressSteps {
				fmt.Println(token.TenantName + ": Fetch complete")
			}

			mutex.Lock()
			m365LicenseCounts[token.TenantName] = tenantLicenseCounts
			mutex.Unlock()
			if !opts.SuppressSteps {
				fmt.Println(token.TenantName + ": Processing complete")
			}
		})
	}
	wg.Wait()

	outputFilePath := opts.OutputFilePath
	if outputFilePath != "" {
		jsonStr, _ := json.Marshal(m365LicenseCounts, jsontext.WithIndent("  "))
		currentDate := time.Now().Format("20060102")
		fileName := outputFilePath + "/m365LicenseCounts-" + currentDate + ".json"
		err := os.WriteFile(fileName, jsonStr, 0644)
		lib.CheckFatalError(err)
		fmt.Println("Saved to " + fileName)
	}

	return
}

//
//

func GetM365UsersAssignedLicenses(token *lib.AzureMultiAuthToken) (users []LicenseReportUser) {
	currentDateTime := time.Now().Local()
	dateTime45DaysAgo := currentDateTime.AddDate(0, 0, -45)
	dateTime90DaysAgo := currentDateTime.AddDate(0, 0, -90)

	selects := []string{
		"accountEnabled",
		"assignedLicenses",
		"userPrincipalName",
		"createdDateTime",
		"signInActivity",
		"onPremisesDistinguishedName",
		"onPremisesSamAccountName",
	}
	allUsers := azure.GetAllEntraUsersForTenant(token, &selects, nil)

	for _, user := range allUsers {
		currUser := LicenseReportUser{
			AccountEnabled:              user.AccountEnabled,
			CreatedDateTime:             user.CreatedDateTime.Local(),
			ID:                          user.ID,
			OnPremisesDistinguishedName: user.OnPremisesDistinguishedName,
			OnPremisesSamAccountName:    user.OnPremisesSamAccountName,
			UserPrincipalName:           user.UserPrincipalName,
		}

		for _, license := range user.AssignedLicenses {
			currUser.AssignedLicenses = append(currUser.AssignedLicenses, license.SkuID)
		}
		var lastSignInDateTimeLocal time.Time
		if user.SignInActivity != nil && user.SignInActivity.LastSignInDateTime != nil {
			lastSignInDateTimeLocal = user.SignInActivity.LastSignInDateTime.Local()
			if lastSignInDateTimeLocal.Before(dateTime90DaysAgo) {
				currUser.NoLogonLast90Days = true
			} else if lastSignInDateTimeLocal.Before(dateTime45DaysAgo) {
				currUser.NoLogonLast45Days = true
			} else {
				currUser.IsActiveUser = true
			}
		} else {
			if currUser.CreatedDateTime.Before(dateTime90DaysAgo) {
				currUser.NoLogonLast90Days = true
			} else if currUser.CreatedDateTime.Before(dateTime45DaysAgo) {
				currUser.NoLogonLast45Days = true
			} else {
				currUser.IsActiveUser = true
			}
		}
		currUser.LastSignInDateTime = &lastSignInDateTimeLocal
		users = append(users, currUser)
	}
	return
}

//
//

func GetM365LicenseSkus(token *lib.AzureMultiAuthToken) (licenseSkus []M365LicenseSku) {
	var (
		resData GetSubscribedSkusByIdResponse
	)

	baseGraphUrl := "https://graph.microsoft.com/"
	endpoint := "/subscribedSkus?$select=consumedUnits,skuId,skuPartNumber"
	apiVersion := "v1.0"
	urlString := baseGraphUrl + apiVersion + endpoint

	res, resErr := azure.HttpGet(urlString, *token)
	lib.CheckHttpGetError(resErr)
	err := json.Unmarshal(res, &resData)
	lib.CheckFatalError(err)

	for _, sku := range resData.Value {
		licenseSkus = append(licenseSkus, M365LicenseSku{
			SkuID:         sku.SkuID,
			SkuPartNumber: sku.SkuPartNumber,
			ConsumedUnits: sku.ConsumedUnits,
		})
	}

	return
}
