package m365

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func GetMailboxStorageUsed(token lib.AzureMultiAuthToken) (mbDetails []MailboxUsageDetail, err error) {
	urlBase := "https://graph.microsoft.com/beta/"
	urlString := urlBase + "/reports/getMailboxUsageDetail(period='D7')"
	res, err := azure.HttpGet(urlString, token)
	lib.CheckFatalError(err)

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

	tokenReq, err := azure.GetAllTenantSPTokens(lib.AzureMultiAuthTokenRequestOptions{
		Scope:         "graph",
		GetWriteToken: false,
	}, nil)
	lib.CheckFatalError(err)

	for _, token := range tokenReq {
		tenantConfig := config.Azure.MultiTenantAuth.Tenants[token.TenantName]
		if !tenantConfig.CheckExchange {
			continue
		}
		data, err := GetMailboxStorageUsed(token)
		lib.CheckFatalError(err)
		mbDetails = append(mbDetails, data...)
	}
	return
}
