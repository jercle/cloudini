package m365

import (
	"bytes"
	_ "embed"
	"encoding/json/v2"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

//go:embed getMailboxes.ps1
var psGetMailboxes string

//go:embed getPermissions.ps1
var psGetPerms string

func GetMailboxStorageUsed(token lib.AzureMultiAuthToken) (mbDetails []MailboxUsageDetail, err error) {
	urlBase := "https://graph.microsoft.com/v1.0/"
	// urlString := urlBase + url.QueryEscape("reports/microsoft.graph.getMailboxUsageDetail(period='D7')")
	urlString := urlBase + "/reports/getMailboxUsageDetail(period='D7')"

	// https://graph.microsoft.com/v1.0/reports/getMailboxUsageDetail(period=
	// https://graph.microsoft.com/v1.0/reports/getMailboxUsageDetail(period='D7')

	res, err := azure.HttpGet(urlString, token)
	if err != nil {
		// fmt.Println(token.TenantName)
		// fmt.Println(urlString)
		lib.CheckFatalError(err)
	}

	bytesReader := bytes.NewReader(res)
	var csvData []MailboxUsageDetail
	err = gocsv.Unmarshal(bytesReader, &csvData)
	if err != nil {
		_, _, cachePath := lib.InitConfig(nil)

		os.WriteFile(cachePath+"/GetMailboxStorageUsed-error.csv", res, 0644)
		fmt.Println("Saved " + cachePath + "/GetMailboxStorageUsed-error.csv")
		fmt.Println("tenant: " + token.TenantName)
		lib.CheckFatalError(err)
	}

	for _, mb := range csvData {
		curr := mb
		curr.TenantName = token.TenantName
		curr.LastAzureSync = time.Now()
		mbDetails = append(mbDetails, curr)
	}

	// lib.JsonMarshalAndPrint(csvData)

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

			curr.Permissions = FormatMailboxPermissions(mb.Permissions, token)

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

func FormatMailboxPermissions(permissions []EXOMailboxUserPermissionRaw, token *lib.AzureMultiAuthToken) (processedPerms []EXOMailboxUserPermission) {
	var (
		wg  sync.WaitGroup
		mut sync.Mutex
	)

	currPerms := make(map[string]EXOMailboxUserPermission)
	for _, perm := range permissions {
		wg.Go(func() {
			var currUser string
			if strings.HasSuffix(perm.User, token.TenantDomain) {
				currUser = perm.User
			} else {
				group, err := azure.GetEntraGroupByObjectId(perm.User, token, &[]string{"displayName"}, nil)
				lib.CheckFatalError(err)
				if group != nil {
					currUser = group.DisplayName
				} else {
					currUser = perm.User
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
			mut.Unlock()
		})
	}
	wg.Wait()

	for _, perm := range currPerms {
		processedPerms = append(processedPerms, perm)
	}

	return
}
