package azure

import (
	"encoding/json/v2"
	"fmt"
	"strconv"

	"github.com/jercle/cloudini/lib"
)

func GetUserFromDeviceSerial(tenantName string, deviceName string) (string, ManagedDeviceMinimal, error) {
	token, err := GetTenantSPToken(lib.AzureMultiAuthTokenRequestOptions{
		TenantName:    tenantName,
		GetWriteToken: true,
		Scope:         "graph",
	}, nil)

	urlString := "https://graph.microsoft.com/v1.0/deviceManagement/managedDevices?$filter=serialNumber+eq+'" + deviceName + "'&$select=deviceName,userPrincipalName,lastSyncDateTime,serialNumber,azureADDeviceId,deviceName,id"
	res, err := HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	var unmarshRes GetManagedDevicesResponse

	json.Unmarshal(res, &unmarshRes)

	if len(unmarshRes.Value) > 1 || len(unmarshRes.Value) == 0 {
		return "", ManagedDeviceMinimal{}, fmt.Errorf("Devices found:", strconv.Itoa(len(unmarshRes.Value)))
	} else {
		return unmarshRes.Value[0].UserPrincipalName, unmarshRes.Value[0], nil
	}
}
