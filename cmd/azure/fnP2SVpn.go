package azure

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jercle/cloudini/lib"
)

func GenerateP2SVpnConnectionHealthDetailed(p2sVpnGatewayResourceId string, tenantName string, storageAccountName string, containerName string) (outputBlobSasUrl string) {

	tenantConfig := lib.GetCldConfig(nil).Azure.MultiTenantAuth.Tenants[tenantName]
	// os.Exit(0)

	token, err := GetTenantSPToken(lib.AzureMultiAuthTokenRequestOptions{
		TenantName:    tenantName,
		GetWriteToken: true,
	}, nil)
	lib.CheckFatalError(err)

	urlString := "https://management.azure.com" + p2sVpnGatewayResourceId + "/getP2sVpnConnectionHealthDetailed?api-version=2025-03-01"

	// fmt.Println(urlString)
	// os.Exit(0)

	t := time.Now().Local()
	// formattedDatetime := fmt.Sprintf("%d%02d%02d-%02d.%02d.%02d",
	// 	t.Year(), t.Month(), t.Day(),
	// 	t.Hour(), t.Minute(), t.Second())
	// filename := formattedDatetime + ".json"

	filename := strconv.Itoa(int(t.Unix())) + ".json"
	fmt.Println(filename)
	fmt.Println(t.String())

	os.Exit(0)

	encodedSAS := GetBlobSAS(storageAccountName, containerName, tenantConfig.TenantID, tenantConfig.Writer.ClientID, tenantConfig.Writer.ClientSecret)

	outputBlobSasUrl = "https://" + storageAccountName + ".blob.core.windows.net/" + containerName + "/" + filename + "?" + encodedSAS
	jsonBody := `{"outputBlobSasUrl":"` + outputBlobSasUrl + `"}`

	// fmt.Println(jsonBody)
	// os.Exit(0)
	_, _, err = HttpPost(urlString, jsonBody, *token)
	lib.CheckFatalError(err)

	return
}

//
//

func GetP2SVpnConnectionDetailsFromBlobSAS(blobSAS string) (connections []AzureP2SConnectionHealth) {
	file, err := StorageBlobHttpGetFromSAS(blobSAS)
	lib.CheckFatalError(err)

	var connectionDetailsRaw []AzureP2SConnectionDetails
	json.Unmarshal(file, &connectionDetailsRaw)

	for _, conn := range connectionDetailsRaw {
		for _, c := range conn.UserNameVpnConnectionHealths {
			for _, detail := range c.VpnConnectionHealths {
				connections = append(connections, detail)
			}
		}
	}

	return
}
