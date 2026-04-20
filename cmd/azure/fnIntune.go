package azure

import (
	"encoding/json/v2"
	"fmt"
	"strconv"
	"sync"
	"time"

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

//
//

func ListManagedDeviceIds(token *lib.AzureMultiAuthToken) (deviceIds []string) {
	urlString := "https://graph.microsoft.com/beta/deviceManagement/managedDevices/?$select=id"
	res, err := HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	var rsp ListManagedDevicesResponse
	json.Unmarshal(res, &rsp)

	for _, d := range rsp.Value {
		deviceIds = append(deviceIds, d.ID)
	}

	return
}

//
//

func GetDeviceInformation(deviceId string, token *lib.AzureMultiAuthToken) (managedDevice ManagedDevice) {
	url := "https://graph.microsoft.com/beta/deviceManagement/managedDevices/" +
		deviceId +
		"?$select=hardwareInformation,id,deviceName,enrolledDateTime,lastSyncDateTime,osVersion,userPrincipalName,azureADDeviceId,model,imei,serialNumber,wiFiMacAddress,managedDeviceName,managementCertificateExpirationDate,ethernetMacAddress,enrollmentProfileName,usersLoggedOn"
	devRes, err := HttpGet(url, *token)
	lib.CheckFatalError(err)

	json.Unmarshal(devRes, &managedDevice)

	managedDevice.TenantId = token.TenantId
	managedDevice.TenantName = token.TenantName
	managedDevice.LastDatabaseSync = time.Now().Local()

	userDataSelects := []string{
		// "mail",
		"userPrincipalName",
		// "surname",
	}
	var loggedInUsersProcessed []ManagedDeviceUserLoggedOn
	for _, user := range managedDevice.UsersLoggedOn {
		userDetails := GetEntraUserByObjectId(user.UserID, token, &userDataSelects, nil)
		// userJson, _ := json.Marshal(userDetails)
		// var loggedInUser ManagedDeviceUserLoggedOn
		// err := json.Unmarshal(userJson, &loggedInUser)
		// lib.CheckFatalError(err)
		loggedInUser := ManagedDeviceUserLoggedOn{
			UserID:            user.UserID,
			UserPrincipalName: userDetails.UserPrincipalName,
			LastLogOnDateTime: user.LastLogOnDateTime,
		}
		loggedInUsersProcessed = append(loggedInUsersProcessed, loggedInUser)

	}
	managedDevice.UsersLoggedOn = loggedInUsersProcessed

	return
}

//
//

func GetMultiDeviceInformation(deviceIds []string, token *lib.AzureMultiAuthToken) (managedDevices []ManagedDevice) {
	var (
		wg  sync.WaitGroup
		mut sync.Mutex
	)

	for _, did := range deviceIds {
		wg.Go(func() {
			d := GetDeviceInformation(did, token)
			mut.Lock()
			managedDevices = append(managedDevices, d)
			mut.Unlock()
		})
	}

	wg.Wait()

	return
}

//
//

func GetIntuneManagedDevicesForConfiguredTenant(tenantName string) (managedDevices []ManagedDevice) {

	token, err := GetTenantSPToken(lib.AzureMultiAuthTokenRequestOptions{
		Scope:         "graph",
		GetWriteToken: true,
		TenantName:    tenantName,
	}, nil)
	lib.CheckFatalError(err)

	deviceIds := ListManagedDeviceIds(token)

	managedDevices = GetMultiDeviceInformation(deviceIds, token)

	return
}

//
//

func GetIntuneManagedDevicesForAllConfiguredTenants() (managedDevices []ManagedDevice) {
	config := lib.GetCldConfig(nil)
	azConfTenants := config.Azure.MultiTenantAuth.Tenants

	var (
		wg  sync.WaitGroup
		mut sync.Mutex
	)

	for tenantName, tData := range azConfTenants {
		if tData.FetchIntuneDevices {
			wg.Go(func() {
				devices := GetIntuneManagedDevicesForConfiguredTenant(tenantName)
				mut.Lock()
				managedDevices = append(managedDevices, devices...)
				mut.Unlock()
			})

		}
	}
	wg.Wait()

	return
}
