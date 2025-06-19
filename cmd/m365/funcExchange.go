package m365

import (
	"bytes"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func GetMailboxStorageUsed(token lib.AzureMultiAuthToken) (mbDetails []MailboxUsageDetail, err error) {
	urlBase := "https://graph.microsoft.com/beta/"
	urlString := urlBase + "/reports/getMailboxUsageDetail(period='D7')"
	res, err := azure.HttpGet(urlString, token)

	bytesReader := bytes.NewReader(res)
	var csvData []MailboxUsageDetail
	gocsv.Unmarshal(bytesReader, &csvData)

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
	tokenReq, err := azure.GetAllTenantSPTokens(lib.AzureMultiAuthTokenRequestOptions{
		Scope:         "graph",
		GetWriteToken: false,
	}, nil)
	lib.CheckFatalError(err)

	for _, token := range tokenReq {
		data, err := GetMailboxStorageUsed(token)
		lib.CheckFatalError(err)
		mbDetails = append(mbDetails, data...)
	}
	return
}
