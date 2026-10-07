package m365

import (
	"bytes"
	_ "embed"
	"encoding/json/v2"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

//go:embed getMailboxes.ps1
var psGetMailboxes string

//go:embed getPermissions.ps1
var psGetPerms string

func GetMailboxStorageUsed(token lib.AzureMultiAuthToken) (mbDetails []MailboxUsageDetail, err error) {

	// urlBase := "https://graph.microsoft.com/v1.0/"
	urlBase := "https://graph.microsoft.com/beta/"

	// urlString := urlBase + "/reports/getMailboxUsageDetail(period='D7')"
	urlString := urlBase + "/reports/getMailboxUsageDetail(period='D7')?$format=application/json"

	res, err := azure.HttpGet(urlString, token)
	lib.CheckFatalError(err)

	var resData AzureGraphResponse[MailboxUsageDetail]

	err = json.Unmarshal(res, &resData)
	lib.CheckFatalError(err)

	for _, mb := range resData.Value {
		curr := mb
		curr.TenantName = token.TenantName
		curr.LastAzureSync = time.Now()
		mbDetails = append(mbDetails, curr)
	}

	nextLink := resData.NextLink

	for nextLink != nil {
		var currentSet AzureGraphResponse[MailboxUsageDetail]

		currRes, _ := azure.HttpGet(*nextLink, token)
		lib.CheckFatalError(err)
		json.Unmarshal(currRes, &currentSet)
		nextLink = currentSet.NextLink

		for _, mb := range currentSet.Value {
			curr := mb
			curr.TenantName = token.TenantName
			curr.LastAzureSync = time.Now()
			mbDetails = append(mbDetails, curr)
		}
	}

	return
}

func GetMailboxStorageUsedAllConfiguredTenants() (mbDetails []MailboxUsageDetail) {
	config := lib.GetCldConfig(nil)

	// tokenReq, err := azure.GetAllTenantSPTokens(lib.AzureMultiAuthTokenRequestOptions{
	// 	Scope:         "graph",
	// 	GetWriteToken: false,
	// }, nil)
	// lib.CheckFatalError(err)

	azTenants := config.Azure.MultiTenantAuth.Tenants

	for tName, tData := range azTenants {
		if !tData.CheckExchange {
			continue
		}
		token, err := azure.GetTenantSPToken(lib.AzureMultiAuthTokenRequestOptions{
			TenantName: tName,
			Scope:      "graph",
		}, nil)
		lib.CheckFatalError(err)

		data, err := GetMailboxStorageUsed(*token)
		lib.CheckFatalError(err)

		mbDetails = append(mbDetails, data...)
	}

	return
}

//
//

func GetEXOMailboxesWithPermissionsAllConfiguredTenants() (mailboxes []EXOMailBox) {
	var (
		wg  sync.WaitGroup
		mut sync.Mutex
	)

	config := lib.GetCldConfig(nil)

	azTenants := config.Azure.MultiTenantAuth.Tenants

	for tName, tData := range azTenants {
		if !tData.CheckExchange {
			continue
		}
		wg.Go(func() {
			token, err := azure.GetTenantSPToken(lib.AzureMultiAuthTokenRequestOptions{
				TenantName: tName,
				Scope:      "m365",
			}, nil)
			lib.CheckFatalError(err)

			tokenGraph, err := azure.GetTenantSPToken(lib.AzureMultiAuthTokenRequestOptions{
				TenantName: tName,
				Scope:      "graph",
			}, nil)
			lib.CheckFatalError(err)

			tenantMailboxes, err := GetEXOMailboxesWithPermissionsForTenant(token, tokenGraph)
			lib.CheckFatalError(err)

			mut.Lock()
			mailboxes = append(mailboxes, tenantMailboxes...)
			mut.Unlock()
		})
	}

	wg.Wait()
	return
}

//
//

func GetEXOMailboxesWithPermissionsForTenant(token *lib.AzureMultiAuthToken, tokenGraph *lib.AzureMultiAuthToken) ([]EXOMailBox, error) {
	tokenStr := token.TokenData.Token
	organization := token.TenantDomain
	appId := token.ClientId

	fmt.Println(token.TenantName + ": Getting mailboxes")

	scriptArgs := []string{tokenStr, organization, appId}
	getMailboxesArgs := append([]string{"-CommandWithArgs", psGetMailboxes}, scriptArgs...)

	getMailboxes := exec.Command("pwsh", getMailboxesArgs...)

	var (
		getMailboxesStdOut bytes.Buffer
		getMailboxesStdErr bytes.Buffer
		getPermsStdOut     bytes.Buffer
		getPermsStdErr     bytes.Buffer
		mailboxes          []EXOMailBox
		mailboxesWithPerms []EXOMailBoxRaw
		mailboxesProcessed []EXOMailBox
	)
	getMailboxes.Stdout = &getMailboxesStdOut
	getMailboxes.Stderr = &getMailboxesStdErr

	err := getMailboxes.Run()
	if err != nil {
		return nil, lib.FormatErrorString(getMailboxesStdErr.String())
	}

	if getMailboxesStdErr.String() != "" {
		return nil, lib.FormatErrorString(getMailboxesStdErr.String())
	}

	err = json.Unmarshal(getMailboxesStdOut.Bytes(), &mailboxes)
	if err != nil {
		return nil, lib.FormatError(err)
	}

	fmt.Println(token.TenantName + ": Fetched " + strconv.Itoa(len(mailboxes)) + " Mailboxes")

	getPermsArgs := append([]string{"-CommandWithArgs", psGetPerms}, scriptArgs...)
	getPerms := exec.Command("pwsh", getPermsArgs...)

	getPerms.Stdout = &getPermsStdOut
	getPerms.Stderr = &getPermsStdErr

	if getPermsStdErr.String() != "" {
		return nil, lib.FormatError(fmt.Errorf(getPermsStdOut.String()))
	}

	stdinPipe, err := getPerms.StdinPipe()
	if err != nil {
		return nil, lib.FormatError(err)
	}

	if err := getPerms.Start(); err != nil {
		return nil, lib.FormatErrorString(getPermsStdErr.String())
	}

	_, err = stdinPipe.Write(getMailboxesStdOut.Bytes())
	if err != nil {
		return nil, lib.FormatError(err)
	}
	stdinPipe.Close()

	if err := getPerms.Wait(); err != nil {
		return nil, lib.FormatError(err)
	}

	if getPermsStdErr.String() != "" {
		return nil, lib.FormatErrorString(getPermsStdErr.String())
	}

	err = json.Unmarshal(getPermsStdOut.Bytes(), &mailboxesWithPerms)
	if err != nil {
		fmt.Println(getPermsStdOut.String())
		fmt.Println(getPermsStdErr.String())
		return nil, lib.FormatError(err)
	}

	mailboxesProcessed = FormatMailboxes(mailboxesWithPerms, tokenGraph)

	fmt.Println(token.TenantName + ": Completed")

	return mailboxesProcessed, nil
}

//
//

func FormatMailboxes(mailboxes []EXOMailBoxRaw, token *lib.AzureMultiAuthToken) (mailboxesProcessed []EXOMailBox) {
	var (
		wg  sync.WaitGroup
		mut sync.Mutex
	)

	for _, mb := range mailboxes {
		wg.Go(func() {
			var curr EXOMailBox

			curr.TenantName = token.TenantName
			curr.Alias = mb.Alias
			curr.DisplayName = mb.DisplayName
			curr.UserPrincipalName = mb.UserPrincipalName
			curr.EmailAddresses = mb.EmailAddresses
			curr.RecipientTypeDetails = mb.RecipientTypeDetails

			perms, permsMap, usersByPerm := FormatMailboxPermissions(mb.Permissions, token)
			curr.Permissions = perms
			curr.PermissionsMap = permsMap
			curr.UsersByPermission = usersByPerm

			mut.Lock()
			mailboxesProcessed = append(mailboxesProcessed, curr)
			mut.Unlock()
		})
	}

	wg.Wait()

	return
}

//
//

func FormatMailboxPermissions(permissions []EXOMailboxUserPermissionRaw, token *lib.AzureMultiAuthToken) ([]EXOMailboxUserPermission, map[string][]string, map[string][]string) {
	var (
		wg             sync.WaitGroup
		mut            sync.Mutex
		processedPerms []EXOMailboxUserPermission
	)

	processedPermsMap := make(map[string][]string)
	usersByPerm := make(map[string][]string)

	currPerms := make(map[string]EXOMailboxUserPermission)
	for _, perm := range permissions {
		wg.Go(func() {
			var currUser string
			if strings.HasSuffix(perm.User, token.TenantDomain) {
				// currUser = strings.ReplaceAll(perm.User, token.TenantDomain, "")
				currUser = perm.User
			} else {
				if perm.Permission == "SendOnBehalf" {
					user, _ := azure.GetEntraUserByObjectId(perm.User, token, nil, nil)
					if user != nil {
						currUser = user.UserPrincipalName
					} else {
						group, _ := azure.GetEntraGroupByObjectId(perm.User, token, &[]string{"displayName"}, nil)
						if group != nil {
							currUser = group.DisplayName
						} else {
							currUser = perm.User
						}
					}
				} else {
					group, err := azure.GetEntraGroupByObjectId(perm.User, token, &[]string{"displayName"}, nil)
					lib.CheckFatalError(err)
					if group != nil {
						currUser = group.DisplayName
					} else {
						currUser = perm.User
					}
				}
			}

			mut.Lock()
			if _, ok := currPerms[currUser]; ok {
				currPerm := currPerms[currUser]
				currPerm.AccessRights = append(currPerm.AccessRights, perm.Permission)
				if !currPerm.IsOwner {
					currPerm.IsOwner = perm.IsOwner
				}
				currPerms[currUser] = currPerm
			} else {
				var currPerm EXOMailboxUserPermission
				currPerm.AccessRights = append(currPerm.AccessRights, perm.Permission)
				currPerm.IsOwner = perm.IsOwner
				currPerm.User = currUser
				currPerms[currUser] = currPerm
			}
			if _, ok := usersByPerm[perm.Permission]; !ok {
				usersByPerm[perm.Permission] = []string{}
			}

			usersByPerm[perm.Permission] = append(usersByPerm[perm.Permission], currUser)

			mut.Unlock()
		})
	}
	wg.Wait()

	for _, perm := range currPerms {
		processedPerms = append(processedPerms, perm)
		processedPermsMap[perm.User] = perm.AccessRights
	}

	return processedPerms, processedPermsMap, usersByPerm
}
