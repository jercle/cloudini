package m365

import (
	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func GetLicencesCurrentAssignedE5() {
	urlBase := "https://graph.microsoft.com/beta/"
	urlString := urlBase + "/reports/getMailboxUsageDetail(period='D7')"
	res, err := azure.HttpGet(urlString, token)
	lib.CheckFatalError(err)

}

// https://graph.microsoft.com/v1.0/users?$filter=assignedLicenses/any(s:s/skuId eq 06ebc4ee-1bb5-47dd-8120-11324bc54e06)

// func GetMailboxStorageUsed(token lib.AzureMultiAuthToken) (mbDetails []MailboxUsageDetail, err error) {
// 	urlBase := "https://graph.microsoft.com/beta/"
// 	urlString := urlBase + "/reports/getMailboxUsageDetail(period='D7')"
// 	res, err := azure.HttpGet(urlString, token)
// 	lib.CheckFatalError(err)

// 	bytesReader := bytes.NewReader(res)
// 	var csvData []MailboxUsageDetail
// 	err = gocsv.Unmarshal(bytesReader, &csvData)
// 	if err != nil {
// 		_, _, cachePath := lib.InitConfig(nil)

// 		os.WriteFile(cachePath+"/GetMailboxStorageUsed-error.csv", res, 0644)
// 		fmt.Println("Saved " + cachePath + "/GetMailboxStorageUsed-error.csv")
// 		fmt.Println("tenant: " + token.TenantName)
// 		lib.CheckFatalError(err)
// 	}

// 	for _, mb := range csvData {
// 		curr := mb
// 		curr.TenantName = token.TenantName
// 		curr.LastAzureSync = time.Now()
// 		mbDetails = append(mbDetails, curr)
// 	}

// 	// lib.JsonMarshalAndPrint(csvData)

// 	return
// }
