package forgerock

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jercle/cloudini/cmd/web"
	"github.com/jercle/cloudini/lib"
)

func GetAllTenantADGroups(tenantName string) ([]LDAPConnectorGroup, error) {
	var groups []LDAPConnectorGroup
	token, err := GetTokenForConfiguredTenant(tenantName)
	if err != nil {
		return nil, err
	}

	config := lib.GetCldConfig(nil)
	frConfig := config.Forgerock.Domains[tenantName]
	urlString := frConfig.UrlBase + "/openidm/system/" + frConfig.LDAPConnector + "/group?_queryFilter=true"
	res := web.SimpleGetRequestWithToken(urlString, token.AccessToken)

	// os.WriteFile("fr-groups.json", res, 0644)

	var resData LDAPConectorGrpResponse
	err = json.Unmarshal(res, &resData)
	if err != nil {
		return nil, err
	}

	for _, grp := range resData.Result {
		// jsonStr, _ := json.Marshal(grpInterface)
		// var grp LDAPConnectorGroup
		// err := json.Unmarshal(jsonStr, &grp)
		// if err != nil {
		// return nil, err
		// }

		groups = append(groups, grp)
	}

	return groups, nil
}

//
//

// func GetAllTenantADGroups(tenantName string) ([]LDAPConnectorUser, error) {
func GetAllTenantADUsers(tenantName string) {
	var users []LDAPConnectorUser
	token, err := GetTokenForConfiguredTenant(tenantName)
	// if err != nil {
	// 	return nil, err
	// }
	lib.CheckFatalError(err)

	config := lib.GetCldConfig(nil)
	frConfig := config.Forgerock.Domains[tenantName]
	urlString := frConfig.UrlBase + "/openidm/system/" + frConfig.LDAPConnector + "/account?_queryFilter=true"
	res := web.SimpleGetRequestWithToken(urlString, token.AccessToken)

	// fmt.Println(string(res))
	var resData LDAPConectorResponse
	err = json.Unmarshal(res, &resData)
	// if err != nil {
	// 	return nil, err
	// }
	lib.CheckFatalError(err)

	// lib.JsonMarshalAndPrint(resData)
	// os.Exit(0)

	for _, userInterface := range resData.Result {
		jsonStr, _ := json.Marshal(userInterface)
		var user LDAPConnectorUser
		err := json.Unmarshal(jsonStr, &user)
		// if err != nil {
		// 	return nil, err
		// }
		lib.CheckFatalError(err)

		users = append(users, user)
	}

	for _, user := range users {
		if strings.Contains(user.SAmAccountName, "svc.") {
			fmt.Println(user.SAmAccountName)
			// lib.JsonMarshalAndPrint(user)
			// os.Exit(0)
		}
	}

	// lib.JsonMarshalAndPrint(users)

	// return groups, nil
}

//
//

func GetADUser(tenantName string, commonName string) LDAPConnectorUser {
	// var users []LDAPConnectorUser
	token, err := GetTokenForConfiguredTenant(tenantName)
	// if err != nil {
	// 	return nil, err
	// }
	lib.CheckFatalError(err)

	// urlEncoded := url.QueryEscape(sAMAccountName)
	// fmt.Println(urlEncoded)

	config := lib.GetCldConfig(nil)
	frConfig := config.Forgerock.Domains[tenantName]
	urlString := frConfig.UrlBase + "/openidm/system/" + frConfig.LDAPConnector + "/account?_queryFilter=cn+eq+'" + commonName + "'"
	// urlString := frConfig.UrlBase + "/openidm/system/" + frConfig.LDAPConnector + "/account?_queryFilter=true"
	res := web.SimpleGetRequestWithToken(urlString, token.AccessToken)

	// fmt.Println(string(res))
	// os.Exit(0)
	var resData LDAPConectorResponse
	err = json.Unmarshal(res, &resData)
	// if err != nil {
	// 	return nil, err
	// }
	lib.CheckFatalError(err)

	// lib.JsonMarshalAndPrint(resData)
	// os.Exit(0)

	if len(resData.Result) == 0 {
		return LDAPConnectorUser{}
	}

	// for _, userInterface := range resData.Result {
	jsonStr, _ := json.Marshal(resData.Result[0])
	var user LDAPConnectorUser
	err = json.Unmarshal(jsonStr, &user)
	// if err != nil {
	// 	return nil, err
	// }
	lib.CheckFatalError(err)
	return user

}

//
//

func GetAdGroupsInConfiguredOUs() (groupList []LDAPConnectorGroup, err error) {
	config := lib.GetCldConfig(nil)

	groupOUs := config.Forgerock.GroupOUs
	tenantName := config.Forgerock.GroupOUsTenant
	groupOUsPrefix := config.Forgerock.GroupOUsPrefix

	groups, err := GetAllTenantADGroups(tenantName)
	if err != nil {
		return groupList, err
	}

	for _, grp := range groups {
		for _, ou := range groupOUs {
			if groupOUsPrefix != nil {
				if strings.Contains(grp.Dn, ou) && strings.HasPrefix(grp.Dn, *groupOUsPrefix) {
					fmt.Println(grp.Dn)
					groupList = append(groupList, grp)
					break
				}
			} else {
				if strings.Contains(grp.Dn, ou) {
					groupList = append(groupList, grp)
					break
				}
			}

		}
	}

	return groupList, nil
}

//
//
