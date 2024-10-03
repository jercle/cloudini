package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func main() {
	startTime := time.Now()
	config := lib.GetCldConfig(nil)
	_ = config
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	// token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{})
	tokenReq, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{
		Scope:         "graph",
		GetWriteToken: true,
	}, nil)
	lib.CheckFatalError(err)
	token, err := tokenReq.SelectTenant("RED")
	lib.CheckFatalError(err)
	// _ = tokens
	_ = token

	// urlString := "https://graph.microsoft.com/v1.0/devices"
	// urlString := "https://graph.microsoft.com/v1.0/deviceManagement/managedDevices/DEVICEID"
	// urlString := "https://graph.microsoft.com/v1.0/devices/DEVICEID/checkMemberObjects"
	// urlString := "https://graph.microsoft.com/beta/deviceManagement/deviceEnrollmentConfigurations"

	// urlString := "https://graph.microsoft.com/beta/deviceManagement/windowsAutopilotDeviceIdentities"
	// urlString := "https://graph.microsoft.com/beta/deviceManagement/windowsAutopilotDeviceIdentities/DEVICE_AUTOPILOT_ID"
	// jsonBody := []byte(`{"ids": ["DEVICEID"]}`)
	// bodyReader := bytes.NewReader(jsonBody)
	// res, _, err := azure.HttpPost(urlString, bodyReader, *token)
	// res, err := azure.HttpGet(urlString, *token)
	// res, _ := azure.HttpGet(urlString, *token)

	var (
		autopilotIdentities            WindowsAutopilotDeviceIdentityList
		deviceConfigurations           DeviceConfigurationList
		entraDevices                   EntraDeviceList
		deviceEnrollmentConfigurations DeviceEnrollmentConfigurationList
	)
	file, err := os.ReadFile("files/autopilotIdentities.json")
	json.Unmarshal(file, &autopilotIdentities)
	file, err = os.ReadFile("files/deviceConfigurations.json")
	json.Unmarshal(file, &deviceConfigurations)
	file, err = os.ReadFile("files/entraDevices.json")
	json.Unmarshal(file, &entraDevices)
	file, err = os.ReadFile("files/deviceEnrollmentConfigurations.json")
	json.Unmarshal(file, &deviceEnrollmentConfigurations)

	// autopilotIdentities := ListWindowsAutopilotDeviceIdentities(token)
	// autopilotIdentitiesStr, _ := json.MarshalIndent(autopilotIdentities, "", "  ")
	// os.WriteFile("files/autopilotIdentities.json", autopilotIdentitiesStr, os.ModePerm)
	// _ = autopilotIdentities
	// fmt.Println("autopilotIdentities", len(autopilotIdentities))

	// deviceConfigurations := ListDeviceConfigurations(token)
	// deviceConfigurationsStr, _ := json.MarshalIndent(deviceConfigurations, "", "  ")
	// os.WriteFile("files/deviceConfigurations.json", deviceConfigurationsStr, os.ModePerm)
	// _ = deviceConfigurations
	// fmt.Println("deviceConfigurations", len(deviceConfigurations))

	// entraDevices := ListAllEntraDevices(token)
	// entraDevicesStr, _ := json.MarshalIndent(entraDevices, "", "  ")
	// os.WriteFile("files/entraDevices.json", entraDevicesStr, os.ModePerm)
	// _ = entraDevices
	// fmt.Println("entraDevices", len(entraDevices))

	// deviceEnrollmentConfigurations := ListDeviceEnrollmentConfigurations(token)
	// deviceEnrollmentConfigurationsStr, _ := json.MarshalIndent(deviceEnrollmentConfigurations, "", "  ")
	// os.WriteFile("files/deviceEnrollmentConfigurations.json", deviceEnrollmentConfigurationsStr, os.ModePerm)
	// _ = deviceEnrollmentConfigurations
	// fmt.Println("deviceEnrollmentConfigurations", len(deviceEnrollmentConfigurations))

	// fmt.Println(len(allDevices))
	var processedDevices []ProcessedDevice
	_ = processedDevices
	// managementTypes := make(map[string][]string)

	for _, device := range autopilotIdentities {
		// deviceStr, err := json.Marshal(device)
		// lib.CheckFatalError(err)
		// var deviceUnmarsh interface{}
		// json.Unmarshal(deviceStr, &deviceUnmarsh)
		currentDevice := make(map[string]string)
		_ = currentDevice

		// jsonStr, _ := json.MarshalIndent(device, "", "  ")
		// fmt.Println(string(jsonStr))
		// os.Exit(0)
		// currentDevice.AutopilotDetails = device
		v := reflect.ValueOf(device)
		typeOfS := v.Type()
		for i := 0; i < v.NumField(); i++ {
			// fmt.Printf("Fild: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
			key := typeOfS.Field(i).Name
			val := v.Field(i).String()
			currentDevice["ap-"+key] = val
		}

		// for key, value := range v {
		// 	currentDevice["ap-"+key] = value
		// }
		// var currentDeviceEntraDetails EntraDevice

	CheckEntraDevices:
		for _, ed := range entraDevices {
			if strings.Contains(ed.DisplayName, device.SerialNumber) {
				// currentDeviceEntraDetails = ed
				v := reflect.ValueOf(ed)
				typeOfS := v.Type()
				// fmt.Println(typeOfS)

				for i := 0; i < v.NumField(); i++ {
					fmt.Printf("Fild: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
					fmt.Println(v.Field(i).Type())
					// if v.Field(i).Type().Kind() == reflect.Slice {
					// fmt.Println(v.Field(i).Type())
					// fmt.Println(v.Field(i))
					// }
					key := typeOfS.Field(i).Name
					val := v.Field(i).String()
					currentDevice["e-"+key] = val
				}
				currentDevice["ContainsEntraDetails"] = "true"
				break CheckEntraDevices
			} else {
				currentDevice["ContainsEntraDetails"] = "false"
			}
		}

		// jsonStr, _ := json.MarshalIndent(currentDevice, "", "  ")
		// fmt.Println(string(jsonStr))
		os.Exit(0)

		// 	processedDevices = append(processedDevices, currentDevice)
		// 	// managementTypes[device.ManagementType] = append(managementTypes[device.ManagementType], device.DisplayName)

	}

	// processedDevicesStr, _ := json.MarshalIndent(processedDevices, "", "  ")
	// // jsonStr, _ := json.MarshalIndent(managementTypes, "", "  ")
	// fmt.Println(string(processedDevicesStr))

	// for _, pd := range processedDevices {
	// 	if !pd.ContainsEntraDetails {
	// 		fmt.Println(pd)
	// 	}
	// }

	// for key, val := range managementTypes {
	// 	var keyname string
	// 	if key == "" {
	// 		keyname = "none"
	// 	} else {
	// 		keyname = key
	// 	}

	// 	fmt.Println(keyname, len(val))
	// }
	// os.Exit(0)

	// jsonStr, _ := json.MarshalIndent(allDevices, "", "  ")
	// fmt.Println(string(jsonStr))
	// fmt.Println(string(res))
	// UploadFileToSharepoint(urlString, dataPath+fileName, token.TokenData.Token)

	elapsed := time.Since(startTime)
	_ = elapsed
}

type ProcessedDevice struct {
	EntraDetails         EntraDevice
	AutopilotDetails     WindowsAutopilotDeviceIdentity
	ContainsEntraDetails bool
}

func ListDeviceEnrollmentConfigurations(token *lib.MultiAuthToken) DeviceEnrollmentConfigurationList {
	var response ListDeviceEnrollmentConfigurationResponse
	urlString := "https://graph.microsoft.com/beta/deviceManagement/deviceEnrollmentConfigurations"

	res, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	// fmt.Println(string(res))

	err = json.Unmarshal(res, &response)
	lib.CheckFatalError(err)

	return response.Value
}

type ListDeviceEnrollmentConfigurationResponse struct {
	Odata_Context string                          `json:"@odata.context,omitempty"`
	Value         []DeviceEnrollmentConfiguration `json:"value,omitempty"`
}

type DeviceEnrollmentConfiguration struct {
	_Odata_Type                        string `json:"@odata.type,omitempty"`
	AllowDeviceResetOnInstallFailure   bool   `json:"allowDeviceResetOnInstallFailure,omitempty"`
	AllowDeviceUseOnInstallFailure     bool   `json:"allowDeviceUseOnInstallFailure,omitempty"`
	AllowLogCollectionOnInstallFailure bool   `json:"allowLogCollectionOnInstallFailure,omitempty"`
	AllowNonBlockingAppInstallation    bool   `json:"allowNonBlockingAppInstallation,omitempty"`
	AndroidForWorkRestriction          *struct {
		BlockedManufacturers            []any  `json:"blockedManufacturers,omitempty"`
		BlockedSkus                     []any  `json:"blockedSkus,omitempty"`
		OSMaximumVersion                string `json:"osMaximumVersion,omitempty"`
		OSMinimumVersion                string `json:"osMinimumVersion,omitempty"`
		PersonalDeviceEnrollmentBlocked bool   `json:"personalDeviceEnrollmentBlocked,omitempty"`
		PlatformBlocked                 bool   `json:"platformBlocked,omitempty"`
	} `json:"androidForWorkRestriction,omitempty"`
	AndroidRestriction *struct {
		BlockedManufacturers            []any  `json:"blockedManufacturers,omitempty"`
		BlockedSkus                     []any  `json:"blockedSkus,omitempty"`
		OSMaximumVersion                string `json:"osMaximumVersion,omitempty"`
		OSMinimumVersion                string `json:"osMinimumVersion,omitempty"`
		PersonalDeviceEnrollmentBlocked bool   `json:"personalDeviceEnrollmentBlocked,omitempty"`
		PlatformBlocked                 bool   `json:"platformBlocked,omitempty"`
	} `json:"androidRestriction,omitempty"`
	BlockDeviceSetupRetryByUser             bool      `json:"blockDeviceSetupRetryByUser,omitempty"`
	CreatedDateTime                         time.Time `json:"createdDateTime,omitempty"`
	CustomErrorMessage                      string    `json:"customErrorMessage,omitempty"`
	Description                             string    `json:"description,omitempty"`
	DeviceEnrollmentConfigurationType       string    `json:"deviceEnrollmentConfigurationType,omitempty"`
	DisableUserStatusTrackingAfterFirstUser bool      `json:"disableUserStatusTrackingAfterFirstUser,omitempty"`
	DisplayName                             string    `json:"displayName,omitempty"`
	EnhancedBiometricsState                 string    `json:"enhancedBiometricsState,omitempty"`
	EnhancedSignInSecurity                  float64   `json:"enhancedSignInSecurity,omitempty"`
	ID                                      string    `json:"id,omitempty"`
	InstallProgressTimeoutInMinutes         float64   `json:"installProgressTimeoutInMinutes,omitempty"`
	InstallQualityUpdates                   bool      `json:"installQualityUpdates,omitempty"`
	IosRestriction                          *struct {
		BlockedManufacturers            []any  `json:"blockedManufacturers,omitempty"`
		BlockedSkus                     []any  `json:"blockedSkus,omitempty"`
		OSMaximumVersion                string `json:"osMaximumVersion,omitempty"`
		OSMinimumVersion                string `json:"osMinimumVersion,omitempty"`
		PersonalDeviceEnrollmentBlocked bool   `json:"personalDeviceEnrollmentBlocked,omitempty"`
		PlatformBlocked                 bool   `json:"platformBlocked,omitempty"`
	} `json:"iosRestriction,omitempty"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty"`
	Limit                float64   `json:"limit,omitempty"`
	MacOSRestriction     *struct {
		BlockedManufacturers            []any `json:"blockedManufacturers,omitempty"`
		BlockedSkus                     []any `json:"blockedSkus,omitempty"`
		OSMaximumVersion                any   `json:"osMaximumVersion,omitempty"`
		OSMinimumVersion                any   `json:"osMinimumVersion,omitempty"`
		PersonalDeviceEnrollmentBlocked bool  `json:"personalDeviceEnrollmentBlocked,omitempty"`
		PlatformBlocked                 bool  `json:"platformBlocked,omitempty"`
	} `json:"macOSRestriction,omitempty"`
	MacRestriction *struct {
		BlockedManufacturers            []any `json:"blockedManufacturers,omitempty"`
		BlockedSkus                     []any `json:"blockedSkus,omitempty"`
		OSMaximumVersion                any   `json:"osMaximumVersion,omitempty"`
		OSMinimumVersion                any   `json:"osMinimumVersion,omitempty"`
		PersonalDeviceEnrollmentBlocked bool  `json:"personalDeviceEnrollmentBlocked,omitempty"`
		PlatformBlocked                 bool  `json:"platformBlocked,omitempty"`
	} `json:"macRestriction,omitempty"`
	PinExpirationInDays         float64 `json:"pinExpirationInDays,omitempty"`
	PinLowercaseCharactersUsage string  `json:"pinLowercaseCharactersUsage,omitempty"`
	PinMaximumLength            float64 `json:"pinMaximumLength,omitempty"`
	PinMinimumLength            float64 `json:"pinMinimumLength,omitempty"`
	PinPreviousBlockCount       float64 `json:"pinPreviousBlockCount,omitempty"`
	PinSpecialCharactersUsage   string  `json:"pinSpecialCharactersUsage,omitempty"`
	PinUppercaseCharactersUsage string  `json:"pinUppercaseCharactersUsage,omitempty"`
	PlatformRestriction         *struct {
		BlockedManufacturers            []any `json:"blockedManufacturers,omitempty"`
		BlockedSkus                     []any `json:"blockedSkus,omitempty"`
		OSMaximumVersion                any   `json:"osMaximumVersion,omitempty"`
		OSMinimumVersion                any   `json:"osMinimumVersion,omitempty"`
		PersonalDeviceEnrollmentBlocked bool  `json:"personalDeviceEnrollmentBlocked,omitempty"`
		PlatformBlocked                 bool  `json:"platformBlocked,omitempty"`
	} `json:"platformRestriction,omitempty"`
	PlatformType                         string   `json:"platformType,omitempty"`
	Priority                             float64  `json:"priority,omitempty"`
	RemotePassportEnabled                bool     `json:"remotePassportEnabled,omitempty"`
	RoleScopeTagIds                      []string `json:"roleScopeTagIds,omitempty"`
	SecurityDeviceRequired               bool     `json:"securityDeviceRequired,omitempty"`
	SecurityKeyForSignIn                 string   `json:"securityKeyForSignIn,omitempty"`
	SelectedMobileAppIds                 []any    `json:"selectedMobileAppIds,omitempty"`
	ShowInstallationProgress             bool     `json:"showInstallationProgress,omitempty"`
	State                                string   `json:"state,omitempty"`
	TrackInstallProgressForAutopilotOnly bool     `json:"trackInstallProgressForAutopilotOnly,omitempty"`
	UnlockWithBiometricsEnabled          bool     `json:"unlockWithBiometricsEnabled,omitempty"`
	Version                              float64  `json:"version,omitempty"`
	WindowsHomeSkuRestriction            *struct {
		BlockedManufacturers            []any `json:"blockedManufacturers,omitempty"`
		BlockedSkus                     []any `json:"blockedSkus,omitempty"`
		OSMaximumVersion                any   `json:"osMaximumVersion,omitempty"`
		OSMinimumVersion                any   `json:"osMinimumVersion,omitempty"`
		PersonalDeviceEnrollmentBlocked bool  `json:"personalDeviceEnrollmentBlocked,omitempty"`
		PlatformBlocked                 bool  `json:"platformBlocked,omitempty"`
	} `json:"windowsHomeSkuRestriction,omitempty"`
	WindowsMobileRestriction *struct {
		BlockedManufacturers            []any  `json:"blockedManufacturers,omitempty"`
		BlockedSkus                     []any  `json:"blockedSkus,omitempty"`
		OSMaximumVersion                string `json:"osMaximumVersion,omitempty"`
		OSMinimumVersion                string `json:"osMinimumVersion,omitempty"`
		PersonalDeviceEnrollmentBlocked bool   `json:"personalDeviceEnrollmentBlocked,omitempty"`
		PlatformBlocked                 bool   `json:"platformBlocked,omitempty"`
	} `json:"windowsMobileRestriction,omitempty"`
	WindowsRestriction *struct {
		BlockedManufacturers            []any  `json:"blockedManufacturers,omitempty"`
		BlockedSkus                     []any  `json:"blockedSkus,omitempty"`
		OSMaximumVersion                string `json:"osMaximumVersion,omitempty"`
		OSMinimumVersion                string `json:"osMinimumVersion,omitempty"`
		PersonalDeviceEnrollmentBlocked bool   `json:"personalDeviceEnrollmentBlocked,omitempty"`
		PlatformBlocked                 bool   `json:"platformBlocked,omitempty"`
	} `json:"windowsRestriction,omitempty"`
}
type DeviceEnrollmentConfigurationList []DeviceEnrollmentConfiguration

func ListDeviceConfigurations(token *lib.MultiAuthToken) DeviceConfigurationList {
	var (
		response ListDeviceConfigurationsResponse
	)
	urlString := "https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations/"
	res, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	err = json.Unmarshal(res, &response)
	lib.CheckFatalError(err)

	return response.Value
}

func ListAllEntraDevices(token *lib.MultiAuthToken) EntraDeviceList {
	var unmarshRes ListEntraDevicesResponse
	var allDevices EntraDeviceList

	urlString := "https://graph.microsoft.com/v1.0/devices"

	res, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)
	json.Unmarshal(res, &unmarshRes)
	// fmt.Println(unmarshRes)
	// fmt.Println(string(res))
	var nextLink string
	nextLink = unmarshRes.Odata_NextLink

	allDevices = append(allDevices, unmarshRes.Value...)
	// fmt.Println(len(allDevices))
	// fmt.Println(nextLink)

	for nextLink != "" {
		var currentSet ListEntraDevicesResponse
		// fmt.Println("Getting next set")
		res, _ := azure.HttpGet(nextLink, *token)
		// fmt.Println(string(res))
		json.Unmarshal(res, &currentSet)
		nextLink = currentSet.Odata_NextLink
		// fmt.Println(nextLink)
		allDevices = append(allDevices, currentSet.Value...)
	}

	// fmt.Println(string(res))
	// fmt.Println(len(allDevices))
	// jsonStr, _ := json.MarshalIndent(allDevices, "", "  ")
	// fmt.Println(string(jsonStr))
	return allDevices
}

// func (tokens AllTenantTokens) SelectTenant(tenantName string) (*MultiAuthToken, error) {
// func (devices EntraDeviceList) Count() {
// 	fmt.Println(len(devices))
// }

func ListWindowsAutopilotDeviceIdentities(token *lib.MultiAuthToken) WindowsAutopilotDeviceIdentityList {
	var (
		response ListWindowsAutopilotDeviceIdentitiesResponse
	)

	urlString := "https://graph.microsoft.com/beta/deviceManagement/windowsAutopilotDeviceIdentities"
	res, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	err = json.Unmarshal(res, &response)
	lib.CheckFatalError(err)

	return response.Value
}

type ListWindowsAutopilotDeviceIdentitiesResponse struct {
	Odata_Context string                             `json:"@odata.context,omitempty"`
	Odata_Count   float64                            `json:"@odata.count,omitempty"`
	Value         WindowsAutopilotDeviceIdentityList `json:"value,omitempty"`
}

type WindowsAutopilotDeviceIdentityList []WindowsAutopilotDeviceIdentity

type WindowsAutopilotDeviceIdentity struct {
	AddressableUserName                       string    `json:"addressableUserName,omitempty"`
	AzureActiveDirectoryDeviceID              string    `json:"azureActiveDirectoryDeviceId,omitempty"`
	AzureAdDeviceID                           string    `json:"azureAdDeviceId,omitempty"`
	DeploymentProfileAssignedDateTime         time.Time `json:"deploymentProfileAssignedDateTime,omitempty"`
	DeploymentProfileAssignmentDetailedStatus string    `json:"deploymentProfileAssignmentDetailedStatus,omitempty"`
	DeploymentProfileAssignmentStatus         string    `json:"deploymentProfileAssignmentStatus,omitempty"`
	DeviceAccountPassword                     any       `json:"deviceAccountPassword,omitempty"`
	DeviceAccountUpn                          string    `json:"deviceAccountUpn,omitempty"`
	DeviceFriendlyName                        any       `json:"deviceFriendlyName,omitempty"`
	DisplayName                               string    `json:"displayName,omitempty"`
	EnrollmentState                           string    `json:"enrollmentState,omitempty"`
	GroupTag                                  string    `json:"groupTag,omitempty"`
	ID                                        string    `json:"id,omitempty"`
	LastContactedDateTime                     time.Time `json:"lastContactedDateTime,omitempty"`
	ManagedDeviceID                           string    `json:"managedDeviceId,omitempty"`
	Manufacturer                              string    `json:"manufacturer,omitempty"`
	Model                                     string    `json:"model,omitempty"`
	ProductKey                                string    `json:"productKey,omitempty"`
	PurchaseOrderIdentifier                   string    `json:"purchaseOrderIdentifier,omitempty"`
	RemediationState                          string    `json:"remediationState,omitempty"`
	RemediationStateLastModifiedDateTime      time.Time `json:"remediationStateLastModifiedDateTime,omitempty"`
	ResourceName                              string    `json:"resourceName,omitempty"`
	SerialNumber                              string    `json:"serialNumber,omitempty"`
	SkuNumber                                 string    `json:"skuNumber,omitempty"`
	SystemFamily                              string    `json:"systemFamily,omitempty"`
	UserPrincipalName                         string    `json:"userPrincipalName,omitempty"`
	UserlessEnrollmentStatus                  string    `json:"userlessEnrollmentStatus,omitempty"`
}

type EntraDevice struct {
	AccountEnabled         bool `json:"accountEnabled,omitempty"`
	AlternativeSecurityIds []struct {
		IdentityProvider any     `json:"identityProvider,omitempty"`
		Key              string  `json:"key,omitempty"`
		Type             float64 `json:"type,omitempty"`
	} `json:"alternativeSecurityIds,omitempty"`
	ApproximateLastSignInDateTime        time.Time `json:"approximateLastSignInDateTime,omitempty"`
	ComplianceExpirationDateTime         any       `json:"complianceExpirationDateTime,omitempty"`
	CompliantApplicationsManagementAppID string    `json:"compliantApplicationsManagementAppId,omitempty"`
	CreatedDateTime                      time.Time `json:"createdDateTime,omitempty"`
	DeletedDateTime                      any       `json:"deletedDateTime,omitempty"`
	DeviceCategory                       any       `json:"deviceCategory,omitempty"`
	DeviceID                             string    `json:"deviceId,omitempty"`
	DeviceMetadata                       any       `json:"deviceMetadata,omitempty"`
	DeviceOwnership                      string    `json:"deviceOwnership,omitempty"`
	DeviceSystemMetadata                 []struct {
		Key   string `json:"key,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"deviceSystemMetadata,omitempty"`
	DeviceVersion         float64  `json:"deviceVersion,omitempty"`
	DisplayName           string   `json:"displayName,omitempty"`
	DomainName            any      `json:"domainName,omitempty"`
	EnrollmentProfileName string   `json:"enrollmentProfileName,omitempty"`
	EnrollmentType        string   `json:"enrollmentType,omitempty"`
	ExchangeActiveSyncIds []string `json:"exchangeActiveSyncIds,omitempty"`
	ExtensionAttributes   struct {
		ExtensionAttribute1  any `json:"extensionAttribute1,omitempty"`
		ExtensionAttribute10 any `json:"extensionAttribute10,omitempty"`
		ExtensionAttribute11 any `json:"extensionAttribute11,omitempty"`
		ExtensionAttribute12 any `json:"extensionAttribute12,omitempty"`
		ExtensionAttribute13 any `json:"extensionAttribute13,omitempty"`
		ExtensionAttribute14 any `json:"extensionAttribute14,omitempty"`
		ExtensionAttribute15 any `json:"extensionAttribute15,omitempty"`
		ExtensionAttribute2  any `json:"extensionAttribute2,omitempty"`
		ExtensionAttribute3  any `json:"extensionAttribute3,omitempty"`
		ExtensionAttribute4  any `json:"extensionAttribute4,omitempty"`
		ExtensionAttribute5  any `json:"extensionAttribute5,omitempty"`
		ExtensionAttribute6  any `json:"extensionAttribute6,omitempty"`
		ExtensionAttribute7  any `json:"extensionAttribute7,omitempty"`
		ExtensionAttribute8  any `json:"extensionAttribute8,omitempty"`
		ExtensionAttribute9  any `json:"extensionAttribute9,omitempty"`
	} `json:"extensionAttributes,omitempty"`
	ExternalSourceName         any       `json:"externalSourceName,omitempty"`
	ID                         string    `json:"id,omitempty"`
	IsCompliant                bool      `json:"isCompliant,omitempty"`
	IsManaged                  bool      `json:"isManaged,omitempty"`
	IsRooted                   bool      `json:"isRooted,omitempty"`
	ManagementType             string    `json:"managementType,omitempty"`
	Manufacturer               string    `json:"manufacturer,omitempty"`
	MdmAppID                   string    `json:"mdmAppId,omitempty"`
	Model                      string    `json:"model,omitempty"`
	OnPremisesLastSyncDateTime time.Time `json:"onPremisesLastSyncDateTime,omitempty"`
	OnPremisesSyncEnabled      bool      `json:"onPremisesSyncEnabled,omitempty"`
	OperatingSystem            string    `json:"operatingSystem,omitempty"`
	OperatingSystemVersion     string    `json:"operatingSystemVersion,omitempty"`
	PhysicalIds                []string  `json:"physicalIds,omitempty"`
	ProfileType                string    `json:"profileType,omitempty"`
	RegistrationDateTime       time.Time `json:"registrationDateTime,omitempty"`
	SourceType                 any       `json:"sourceType,omitempty"`
	SystemLabels               []any     `json:"systemLabels,omitempty"`
	TrustType                  string    `json:"trustType,omitempty"`
	UserCertificate            string    `json:"userCertificate,omitempty"`
}

type EntraDeviceList []EntraDevice

type ListEntraDevicesResponse struct {
	Odata_Context  string        `json:"@odata.context,omitempty"`
	Odata_NextLink string        `json:"@odata.nextLink,omitempty"`
	Value          []EntraDevice `json:"value,omitempty"`
}

type ListDeviceConfigurationsResponse struct {
	Odata_Context string                `json:"@odata.context,omitempty"`
	Value         []DeviceConfiguration `json:"value,omitempty"`
}

type DeviceConfigurationList []DeviceConfiguration
type DeviceConfiguration struct {
	Odata_Type                                          string `json:"@odata.type,omitempty"`
	AccountBlockModification                            bool   `json:"accountBlockModification,omitempty"`
	AccountsBlockAddingNonMicrosoftAccountEmail         bool   `json:"accountsBlockAddingNonMicrosoftAccountEmail,omitempty"`
	ActivateAppsWithVoice                               string `json:"activateAppsWithVoice,omitempty"`
	ActivationLockAllowWhenSupervised                   bool   `json:"activationLockAllowWhenSupervised,omitempty"`
	ActiveHoursEnd                                      string `json:"activeHoursEnd,omitempty"`
	ActiveHoursStart                                    string `json:"activeHoursStart,omitempty"`
	AirDropBlocked                                      bool   `json:"airDropBlocked,omitempty"`
	AirDropForceUnmanagedDropTarget                     bool   `json:"airDropForceUnmanagedDropTarget,omitempty"`
	AirPlayForcePairingPasswordForOutgoingRequests      bool   `json:"airPlayForcePairingPasswordForOutgoingRequests,omitempty"`
	AirPrintBlockCredentialsStorage                     bool   `json:"airPrintBlockCredentialsStorage,omitempty"`
	AirPrintBlocked                                     bool   `json:"airPrintBlocked,omitempty"`
	AirPrintBlockiBeaconDiscovery                       bool   `json:"airPrintBlockiBeaconDiscovery,omitempty"`
	AirPrintDestinations                                []any  `json:"airPrintDestinations,omitempty"`
	AirPrintForceTrustedTls                             bool   `json:"airPrintForceTrustedTLS,omitempty"`
	AllowDeviceHealthMonitoring                         string `json:"allowDeviceHealthMonitoring,omitempty"`
	AllowWindows11Upgrade                               bool   `json:"allowWindows11Upgrade,omitempty"`
	AntiTheftModeBlocked                                bool   `json:"antiTheftModeBlocked,omitempty"`
	AppClipsBlocked                                     bool   `json:"appClipsBlocked,omitempty"`
	AppLockerApplicationControl                         string `json:"appLockerApplicationControl,omitempty"`
	AppManagementMsiAllowUserControlOverInstall         bool   `json:"appManagementMSIAllowUserControlOverInstall,omitempty"`
	AppManagementMsiAlwaysInstallWithElevatedPrivileges bool   `json:"appManagementMSIAlwaysInstallWithElevatedPrivileges,omitempty"`
	AppManagementPackageFamilyNamesToLaunchAfterLogOn   []any  `json:"appManagementPackageFamilyNamesToLaunchAfterLogOn,omitempty"`
	AppRemovalBlocked                                   bool   `json:"appRemovalBlocked,omitempty"`
	AppStoreBlockAutomaticDownloads                     bool   `json:"appStoreBlockAutomaticDownloads,omitempty"`
	AppStoreBlockInAppPurchases                         bool   `json:"appStoreBlockInAppPurchases,omitempty"`
	AppStoreBlockUiAppInstallation                      bool   `json:"appStoreBlockUIAppInstallation,omitempty"`
	AppStoreBlocked                                     bool   `json:"appStoreBlocked,omitempty"`
	AppStoreRequirePassword                             bool   `json:"appStoreRequirePassword,omitempty"`
	AppleNewsBlocked                                    bool   `json:"appleNewsBlocked,omitempty"`
	ApplePersonalizedAdsBlocked                         bool   `json:"applePersonalizedAdsBlocked,omitempty"`
	AppleWatchBlockPairing                              bool   `json:"appleWatchBlockPairing,omitempty"`
	AppleWatchForceWristDetection                       bool   `json:"appleWatchForceWristDetection,omitempty"`
	ApplicationGuardAllowCameraMicrophoneRedirection    any    `json:"applicationGuardAllowCameraMicrophoneRedirection,omitempty"`
	ApplicationGuardAllowFileSaveOnHost                 bool   `json:"applicationGuardAllowFileSaveOnHost,omitempty"`
	ApplicationGuardAllowPersistence                    bool   `json:"applicationGuardAllowPersistence,omitempty"`
	ApplicationGuardAllowPrintToLocalPrinters           bool   `json:"applicationGuardAllowPrintToLocalPrinters,omitempty"`
	ApplicationGuardAllowPrintToNetworkPrinters         bool   `json:"applicationGuardAllowPrintToNetworkPrinters,omitempty"`
	ApplicationGuardAllowPrintToPdf                     bool   `json:"applicationGuardAllowPrintToPDF,omitempty"`
	ApplicationGuardAllowPrintToXps                     bool   `json:"applicationGuardAllowPrintToXPS,omitempty"`
	ApplicationGuardAllowVirtualGpu                     bool   `json:"applicationGuardAllowVirtualGPU,omitempty"`
	ApplicationGuardBlockClipboardSharing               string `json:"applicationGuardBlockClipboardSharing,omitempty"`
	ApplicationGuardBlockFileTransfer                   string `json:"applicationGuardBlockFileTransfer,omitempty"`
	ApplicationGuardBlockNonEnterpriseContent           bool   `json:"applicationGuardBlockNonEnterpriseContent,omitempty"`
	ApplicationGuardCertificateThumbprints              []any  `json:"applicationGuardCertificateThumbprints,omitempty"`
	ApplicationGuardEnabled                             bool   `json:"applicationGuardEnabled,omitempty"`
	ApplicationGuardEnabledOptions                      string `json:"applicationGuardEnabledOptions,omitempty"`
	ApplicationGuardForceAuditing                       bool   `json:"applicationGuardForceAuditing,omitempty"`
	AppsAllowTrustedAppsSideloading                     string `json:"appsAllowTrustedAppsSideloading,omitempty"`
	AppsBlockWindowsStoreOriginatedApps                 bool   `json:"appsBlockWindowsStoreOriginatedApps,omitempty"`
	AppsSingleAppModeList                               []any  `json:"appsSingleAppModeList,omitempty"`
	AppsVisibilityList                                  []struct {
		AppID       string `json:"appId,omitempty"`
		AppStoreURL string `json:"appStoreUrl,omitempty"`
		Name        string `json:"name,omitempty"`
		Publisher   string `json:"publisher,omitempty"`
	} `json:"appsVisibilityList,omitempty"`
	AppsVisibilityListType                         string  `json:"appsVisibilityListType,omitempty"`
	AssetTagTemplate                               any     `json:"assetTagTemplate,omitempty"`
	AssociatedDomains                              []any   `json:"associatedDomains,omitempty"`
	AuthenticationAllowSecondaryDevice             bool    `json:"authenticationAllowSecondaryDevice,omitempty"`
	AuthenticationMethod                           string  `json:"authenticationMethod,omitempty"`
	AuthenticationPreferredAzureAdTenantDomainName any     `json:"authenticationPreferredAzureADTenantDomainName,omitempty"`
	AuthenticationWebSignIn                        string  `json:"authenticationWebSignIn,omitempty"`
	AutoFillForceAuthentication                    bool    `json:"autoFillForceAuthentication,omitempty"`
	AutoRestartNotificationDismissal               string  `json:"autoRestartNotificationDismissal,omitempty"`
	AutoUnlockBlocked                              bool    `json:"autoUnlockBlocked,omitempty"`
	AutomaticUpdateMode                            string  `json:"automaticUpdateMode,omitempty"`
	BackgroundDownloadFromHTTPDelayInSeconds       float64 `json:"backgroundDownloadFromHttpDelayInSeconds,omitempty"`
	BandwidthMode                                  *struct {
		Odata_Type                           string  `json:"@odata.type,omitempty"`
		MaximumBackgroundBandwidthPercentage float64 `json:"maximumBackgroundBandwidthPercentage,omitempty"`
		MaximumForegroundBandwidthPercentage float64 `json:"maximumForegroundBandwidthPercentage,omitempty"`
	} `json:"bandwidthMode,omitempty"`
	BitLockerAllowStandardUserEncryption          bool `json:"bitLockerAllowStandardUserEncryption,omitempty"`
	BitLockerDisableWarningForOtherDiskEncryption bool `json:"bitLockerDisableWarningForOtherDiskEncryption,omitempty"`
	BitLockerEnableStorageCardEncryptionOnMobile  bool `json:"bitLockerEnableStorageCardEncryptionOnMobile,omitempty"`
	BitLockerEncryptDevice                        bool `json:"bitLockerEncryptDevice,omitempty"`
	BitLockerFixedDrivePolicy                     *struct {
		EncryptionMethod                any  `json:"encryptionMethod,omitempty"`
		RecoveryOptions                 any  `json:"recoveryOptions,omitempty"`
		RequireEncryptionForWriteAccess bool `json:"requireEncryptionForWriteAccess,omitempty"`
	} `json:"bitLockerFixedDrivePolicy,omitempty"`
	BitLockerRecoveryPasswordRotation string `json:"bitLockerRecoveryPasswordRotation,omitempty"`
	BitLockerRemovableDrivePolicy     *struct {
		BlockCrossOrganizationWriteAccess bool `json:"blockCrossOrganizationWriteAccess,omitempty"`
		EncryptionMethod                  any  `json:"encryptionMethod,omitempty"`
		RequireEncryptionForWriteAccess   bool `json:"requireEncryptionForWriteAccess,omitempty"`
	} `json:"bitLockerRemovableDrivePolicy,omitempty"`
	BitLockerSystemDrivePolicy *struct {
		EncryptionMethod                         any    `json:"encryptionMethod,omitempty"`
		MinimumPinLength                         any    `json:"minimumPinLength,omitempty"`
		PrebootRecoveryEnableMessageAndURL       bool   `json:"prebootRecoveryEnableMessageAndUrl,omitempty"`
		PrebootRecoveryMessage                   any    `json:"prebootRecoveryMessage,omitempty"`
		PrebootRecoveryURL                       any    `json:"prebootRecoveryUrl,omitempty"`
		RecoveryOptions                          any    `json:"recoveryOptions,omitempty"`
		StartupAuthenticationBlockWithoutTpmChip bool   `json:"startupAuthenticationBlockWithoutTpmChip,omitempty"`
		StartupAuthenticationRequired            bool   `json:"startupAuthenticationRequired,omitempty"`
		StartupAuthenticationTpmKeyUsage         string `json:"startupAuthenticationTpmKeyUsage,omitempty"`
		StartupAuthenticationTpmPinAndKeyUsage   string `json:"startupAuthenticationTpmPinAndKeyUsage,omitempty"`
		StartupAuthenticationTpmPinUsage         string `json:"startupAuthenticationTpmPinUsage,omitempty"`
		StartupAuthenticationTpmUsage            string `json:"startupAuthenticationTpmUsage,omitempty"`
	} `json:"bitLockerSystemDrivePolicy,omitempty"`
	BlockSystemAppRemoval                                     bool    `json:"blockSystemAppRemoval,omitempty"`
	BluetoothAllowedServices                                  []any   `json:"bluetoothAllowedServices,omitempty"`
	BluetoothBlockAdvertising                                 bool    `json:"bluetoothBlockAdvertising,omitempty"`
	BluetoothBlockDiscoverableMode                            bool    `json:"bluetoothBlockDiscoverableMode,omitempty"`
	BluetoothBlockModification                                bool    `json:"bluetoothBlockModification,omitempty"`
	BluetoothBlockPrePairing                                  bool    `json:"bluetoothBlockPrePairing,omitempty"`
	BluetoothBlockPromptedProximalConnections                 bool    `json:"bluetoothBlockPromptedProximalConnections,omitempty"`
	BluetoothBlocked                                          bool    `json:"bluetoothBlocked,omitempty"`
	BusinessReadyUpdatesOnly                                  string  `json:"businessReadyUpdatesOnly,omitempty"`
	CacheServerBackgroundDownloadFallbackToHTTPDelayInSeconds float64 `json:"cacheServerBackgroundDownloadFallbackToHttpDelayInSeconds,omitempty"`
	CacheServerForegroundDownloadFallbackToHTTPDelayInSeconds float64 `json:"cacheServerForegroundDownloadFallbackToHttpDelayInSeconds,omitempty"`
	CacheServerHostNames                                      []any   `json:"cacheServerHostNames,omitempty"`
	CameraBlocked                                             bool    `json:"cameraBlocked,omitempty"`
	CellularBlockDataRoaming                                  bool    `json:"cellularBlockDataRoaming,omitempty"`
	CellularBlockDataWhenRoaming                              bool    `json:"cellularBlockDataWhenRoaming,omitempty"`
	CellularBlockGlobalBackgroundFetchWhileRoaming            bool    `json:"cellularBlockGlobalBackgroundFetchWhileRoaming,omitempty"`
	CellularBlockPerAppDataModification                       bool    `json:"cellularBlockPerAppDataModification,omitempty"`
	CellularBlockPersonalHotspot                              bool    `json:"cellularBlockPersonalHotspot,omitempty"`
	CellularBlockPersonalHotspotModification                  bool    `json:"cellularBlockPersonalHotspotModification,omitempty"`
	CellularBlockPlanModification                             bool    `json:"cellularBlockPlanModification,omitempty"`
	CellularBlockVoiceRoaming                                 bool    `json:"cellularBlockVoiceRoaming,omitempty"`
	CellularBlockVpn                                          bool    `json:"cellularBlockVpn,omitempty"`
	CellularBlockVpnWhenRoaming                               bool    `json:"cellularBlockVpnWhenRoaming,omitempty"`
	CellularData                                              string  `json:"cellularData,omitempty"`
	CertFileName                                              string  `json:"certFileName,omitempty"`
	CertificateStore                                          string  `json:"certificateStore,omitempty"`
	CertificateValidityPeriodScale                            string  `json:"certificateValidityPeriodScale,omitempty"`
	CertificateValidityPeriodValue                            float64 `json:"certificateValidityPeriodValue,omitempty"`
	CertificatesBlockManualRootCertificateInstallation        bool    `json:"certificatesBlockManualRootCertificateInstallation,omitempty"`
	CertificatesBlockUntrustedTlsCertificates                 bool    `json:"certificatesBlockUntrustedTlsCertificates,omitempty"`
	ClassroomAppBlockRemoteScreenObservation                  bool    `json:"classroomAppBlockRemoteScreenObservation,omitempty"`
	ClassroomAppForceUnpromptedScreenObservation              bool    `json:"classroomAppForceUnpromptedScreenObservation,omitempty"`
	ClassroomForceAutomaticallyJoinClasses                    bool    `json:"classroomForceAutomaticallyJoinClasses,omitempty"`
	ClassroomForceRequestPermissionToLeaveClasses             bool    `json:"classroomForceRequestPermissionToLeaveClasses,omitempty"`
	ClassroomForceUnpromptedAppAndDeviceLock                  bool    `json:"classroomForceUnpromptedAppAndDeviceLock,omitempty"`
	CloudName                                                 any     `json:"cloudName,omitempty"`
	CompliantAppListType                                      string  `json:"compliantAppListType,omitempty"`
	CompliantAppsList                                         []struct {
		AppID       string `json:"appId,omitempty"`
		AppStoreURL string `json:"appStoreUrl,omitempty"`
		Name        string `json:"name,omitempty"`
		Publisher   string `json:"publisher,omitempty"`
	} `json:"compliantAppsList,omitempty"`
	ConfigDeviceHealthMonitoringCustomScope any       `json:"configDeviceHealthMonitoringCustomScope,omitempty"`
	ConfigDeviceHealthMonitoringScope       string    `json:"configDeviceHealthMonitoringScope,omitempty"`
	ConfigurationProfileBlockChanges        bool      `json:"configurationProfileBlockChanges,omitempty"`
	ConfigureTimeZone                       any       `json:"configureTimeZone,omitempty"`
	ConnectedDevicesServiceBlocked          bool      `json:"connectedDevicesServiceBlocked,omitempty"`
	ConnectionName                          string    `json:"connectionName,omitempty"`
	ConnectionType                          string    `json:"connectionType,omitempty"`
	ContactsAllowManagedToUnmanagedWrite    bool      `json:"contactsAllowManagedToUnmanagedWrite,omitempty"`
	ContactsAllowUnmanagedToManagedRead     bool      `json:"contactsAllowUnmanagedToManagedRead,omitempty"`
	ContentFilterSettings                   any       `json:"contentFilterSettings,omitempty"`
	ContinuousPathKeyboardBlocked           bool      `json:"continuousPathKeyboardBlocked,omitempty"`
	CopyPasteBlocked                        bool      `json:"copyPasteBlocked,omitempty"`
	CortanaBlocked                          bool      `json:"cortanaBlocked,omitempty"`
	CreatedDateTime                         time.Time `json:"createdDateTime,omitempty"`
	CryptographyAllowFipsAlgorithmPolicy    bool      `json:"cryptographyAllowFipsAlgorithmPolicy,omitempty"`
	CustomData                              []struct {
		Key   string `json:"key,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"customData,omitempty"`
	CustomKeyValueData []struct {
		Name  string `json:"name,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"customKeyValueData,omitempty"`
	CustomSubjectAlternativeNames                              []any   `json:"customSubjectAlternativeNames,omitempty"`
	CustomUpdateTimeWindows                                    []any   `json:"customUpdateTimeWindows,omitempty"`
	DataProtectionBlockDirectMemoryAccess                      bool    `json:"dataProtectionBlockDirectMemoryAccess,omitempty"`
	DateAndTimeForceSetAutomatically                           bool    `json:"dateAndTimeForceSetAutomatically,omitempty"`
	DeadlineForFeatureUpdatesInDays                            float64 `json:"deadlineForFeatureUpdatesInDays,omitempty"`
	DeadlineForQualityUpdatesInDays                            float64 `json:"deadlineForQualityUpdatesInDays,omitempty"`
	DeadlineGracePeriodInDays                                  float64 `json:"deadlineGracePeriodInDays,omitempty"`
	DefenderAdditionalGuardedFolders                           []any   `json:"defenderAdditionalGuardedFolders,omitempty"`
	DefenderAdobeReaderLaunchChildProcess                      string  `json:"defenderAdobeReaderLaunchChildProcess,omitempty"`
	DefenderAdvancedRansomewareProtectionType                  string  `json:"defenderAdvancedRansomewareProtectionType,omitempty"`
	DefenderAllowBehaviorMonitoring                            any     `json:"defenderAllowBehaviorMonitoring,omitempty"`
	DefenderAllowCloudProtection                               any     `json:"defenderAllowCloudProtection,omitempty"`
	DefenderAllowEndUserAccess                                 any     `json:"defenderAllowEndUserAccess,omitempty"`
	DefenderAllowIntrusionPreventionSystem                     any     `json:"defenderAllowIntrusionPreventionSystem,omitempty"`
	DefenderAllowOnAccessProtection                            any     `json:"defenderAllowOnAccessProtection,omitempty"`
	DefenderAllowRealTimeMonitoring                            any     `json:"defenderAllowRealTimeMonitoring,omitempty"`
	DefenderAllowScanArchiveFiles                              any     `json:"defenderAllowScanArchiveFiles,omitempty"`
	DefenderAllowScanDownloads                                 any     `json:"defenderAllowScanDownloads,omitempty"`
	DefenderAllowScanNetworkFiles                              any     `json:"defenderAllowScanNetworkFiles,omitempty"`
	DefenderAllowScanRemovableDrivesDuringFullScan             any     `json:"defenderAllowScanRemovableDrivesDuringFullScan,omitempty"`
	DefenderAllowScanScriptsLoadedInInternetExplorer           any     `json:"defenderAllowScanScriptsLoadedInInternetExplorer,omitempty"`
	DefenderAttackSurfaceReductionExcludedPaths                []any   `json:"defenderAttackSurfaceReductionExcludedPaths,omitempty"`
	DefenderBlockEndUserAccess                                 bool    `json:"defenderBlockEndUserAccess,omitempty"`
	DefenderBlockOnAccessProtection                            bool    `json:"defenderBlockOnAccessProtection,omitempty"`
	DefenderBlockPersistenceThroughWmiType                     string  `json:"defenderBlockPersistenceThroughWmiType,omitempty"`
	DefenderCheckForSignaturesBeforeRunningScan                any     `json:"defenderCheckForSignaturesBeforeRunningScan,omitempty"`
	DefenderCloudBlockLevel                                    string  `json:"defenderCloudBlockLevel,omitempty"`
	DefenderCloudExtendedTimeout                               any     `json:"defenderCloudExtendedTimeout,omitempty"`
	DefenderCloudExtendedTimeoutInSeconds                      any     `json:"defenderCloudExtendedTimeoutInSeconds,omitempty"`
	DefenderDaysBeforeDeletingQuarantinedMalware               any     `json:"defenderDaysBeforeDeletingQuarantinedMalware,omitempty"`
	DefenderDetectedMalwareActions                             any     `json:"defenderDetectedMalwareActions,omitempty"`
	DefenderDisableBehaviorMonitoring                          any     `json:"defenderDisableBehaviorMonitoring,omitempty"`
	DefenderDisableCatchupFullScan                             bool    `json:"defenderDisableCatchupFullScan,omitempty"`
	DefenderDisableCatchupQuickScan                            bool    `json:"defenderDisableCatchupQuickScan,omitempty"`
	DefenderDisableCloudProtection                             any     `json:"defenderDisableCloudProtection,omitempty"`
	DefenderDisableIntrusionPreventionSystem                   any     `json:"defenderDisableIntrusionPreventionSystem,omitempty"`
	DefenderDisableOnAccessProtection                          any     `json:"defenderDisableOnAccessProtection,omitempty"`
	DefenderDisableRealTimeMonitoring                          any     `json:"defenderDisableRealTimeMonitoring,omitempty"`
	DefenderDisableScanArchiveFiles                            any     `json:"defenderDisableScanArchiveFiles,omitempty"`
	DefenderDisableScanDownloads                               any     `json:"defenderDisableScanDownloads,omitempty"`
	DefenderDisableScanNetworkFiles                            any     `json:"defenderDisableScanNetworkFiles,omitempty"`
	DefenderDisableScanRemovableDrivesDuringFullScan           any     `json:"defenderDisableScanRemovableDrivesDuringFullScan,omitempty"`
	DefenderDisableScanScriptsLoadedInInternetExplorer         any     `json:"defenderDisableScanScriptsLoadedInInternetExplorer,omitempty"`
	DefenderEmailContentExecution                              string  `json:"defenderEmailContentExecution,omitempty"`
	DefenderEmailContentExecutionType                          string  `json:"defenderEmailContentExecutionType,omitempty"`
	DefenderEnableLowCpuPriority                               any     `json:"defenderEnableLowCpuPriority,omitempty"`
	DefenderEnableScanIncomingMail                             any     `json:"defenderEnableScanIncomingMail,omitempty"`
	DefenderEnableScanMappedNetworkDrivesDuringFullScan        any     `json:"defenderEnableScanMappedNetworkDrivesDuringFullScan,omitempty"`
	DefenderExploitProtectionXML                               string  `json:"defenderExploitProtectionXml,omitempty"`
	DefenderExploitProtectionXMLFileName                       any     `json:"defenderExploitProtectionXmlFileName,omitempty"`
	DefenderFileExtensionsToExclude                            []any   `json:"defenderFileExtensionsToExclude,omitempty"`
	DefenderFilesAndFoldersToExclude                           []any   `json:"defenderFilesAndFoldersToExclude,omitempty"`
	DefenderGuardMyFoldersType                                 string  `json:"defenderGuardMyFoldersType,omitempty"`
	DefenderGuardedFoldersAllowedAppPaths                      []any   `json:"defenderGuardedFoldersAllowedAppPaths,omitempty"`
	DefenderMonitorFileActivity                                string  `json:"defenderMonitorFileActivity,omitempty"`
	DefenderNetworkProtectionType                              string  `json:"defenderNetworkProtectionType,omitempty"`
	DefenderOfficeAppsExecutableContentCreationOrLaunch        string  `json:"defenderOfficeAppsExecutableContentCreationOrLaunch,omitempty"`
	DefenderOfficeAppsExecutableContentCreationOrLaunchType    string  `json:"defenderOfficeAppsExecutableContentCreationOrLaunchType,omitempty"`
	DefenderOfficeAppsLaunchChildProcess                       string  `json:"defenderOfficeAppsLaunchChildProcess,omitempty"`
	DefenderOfficeAppsLaunchChildProcessType                   string  `json:"defenderOfficeAppsLaunchChildProcessType,omitempty"`
	DefenderOfficeAppsOtherProcessInjection                    string  `json:"defenderOfficeAppsOtherProcessInjection,omitempty"`
	DefenderOfficeAppsOtherProcessInjectionType                string  `json:"defenderOfficeAppsOtherProcessInjectionType,omitempty"`
	DefenderOfficeCommunicationAppsLaunchChildProcess          string  `json:"defenderOfficeCommunicationAppsLaunchChildProcess,omitempty"`
	DefenderOfficeMacroCodeAllowWin32Imports                   string  `json:"defenderOfficeMacroCodeAllowWin32Imports,omitempty"`
	DefenderOfficeMacroCodeAllowWin32ImportsType               string  `json:"defenderOfficeMacroCodeAllowWin32ImportsType,omitempty"`
	DefenderPotentiallyUnwantedAppAction                       any     `json:"defenderPotentiallyUnwantedAppAction,omitempty"`
	DefenderPotentiallyUnwantedAppActionSetting                string  `json:"defenderPotentiallyUnwantedAppActionSetting,omitempty"`
	DefenderPreventCredentialStealingType                      string  `json:"defenderPreventCredentialStealingType,omitempty"`
	DefenderProcessCreation                                    string  `json:"defenderProcessCreation,omitempty"`
	DefenderProcessCreationType                                string  `json:"defenderProcessCreationType,omitempty"`
	DefenderProcessesToExclude                                 []any   `json:"defenderProcessesToExclude,omitempty"`
	DefenderPromptForSampleSubmission                          string  `json:"defenderPromptForSampleSubmission,omitempty"`
	DefenderRequireBehaviorMonitoring                          bool    `json:"defenderRequireBehaviorMonitoring,omitempty"`
	DefenderRequireCloudProtection                             bool    `json:"defenderRequireCloudProtection,omitempty"`
	DefenderRequireNetworkInspectionSystem                     bool    `json:"defenderRequireNetworkInspectionSystem,omitempty"`
	DefenderRequireRealTimeMonitoring                          bool    `json:"defenderRequireRealTimeMonitoring,omitempty"`
	DefenderScanArchiveFiles                                   bool    `json:"defenderScanArchiveFiles,omitempty"`
	DefenderScanDirection                                      any     `json:"defenderScanDirection,omitempty"`
	DefenderScanDownloads                                      bool    `json:"defenderScanDownloads,omitempty"`
	DefenderScanIncomingMail                                   bool    `json:"defenderScanIncomingMail,omitempty"`
	DefenderScanMappedNetworkDrivesDuringFullScan              bool    `json:"defenderScanMappedNetworkDrivesDuringFullScan,omitempty"`
	DefenderScanMaxCpu                                         any     `json:"defenderScanMaxCpu,omitempty"`
	DefenderScanMaxCpuPercentage                               any     `json:"defenderScanMaxCpuPercentage,omitempty"`
	DefenderScanNetworkFiles                                   bool    `json:"defenderScanNetworkFiles,omitempty"`
	DefenderScanRemovableDrivesDuringFullScan                  bool    `json:"defenderScanRemovableDrivesDuringFullScan,omitempty"`
	DefenderScanScriptsLoadedInInternetExplorer                bool    `json:"defenderScanScriptsLoadedInInternetExplorer,omitempty"`
	DefenderScanType                                           string  `json:"defenderScanType,omitempty"`
	DefenderScheduleScanEnableLowCpuPriority                   bool    `json:"defenderScheduleScanEnableLowCpuPriority,omitempty"`
	DefenderScheduledQuickScanTime                             any     `json:"defenderScheduledQuickScanTime,omitempty"`
	DefenderScheduledScanDay                                   any     `json:"defenderScheduledScanDay,omitempty"`
	DefenderScheduledScanTime                                  any     `json:"defenderScheduledScanTime,omitempty"`
	DefenderScriptDownloadedPayloadExecution                   string  `json:"defenderScriptDownloadedPayloadExecution,omitempty"`
	DefenderScriptDownloadedPayloadExecutionType               string  `json:"defenderScriptDownloadedPayloadExecutionType,omitempty"`
	DefenderScriptObfuscatedMacroCode                          string  `json:"defenderScriptObfuscatedMacroCode,omitempty"`
	DefenderScriptObfuscatedMacroCodeType                      string  `json:"defenderScriptObfuscatedMacroCodeType,omitempty"`
	DefenderSecurityCenterBlockExploitProtectionOverride       bool    `json:"defenderSecurityCenterBlockExploitProtectionOverride,omitempty"`
	DefenderSecurityCenterDisableAccountUi                     any     `json:"defenderSecurityCenterDisableAccountUI,omitempty"`
	DefenderSecurityCenterDisableAppBrowserUi                  any     `json:"defenderSecurityCenterDisableAppBrowserUI,omitempty"`
	DefenderSecurityCenterDisableClearTpmUi                    any     `json:"defenderSecurityCenterDisableClearTpmUI,omitempty"`
	DefenderSecurityCenterDisableFamilyUi                      any     `json:"defenderSecurityCenterDisableFamilyUI,omitempty"`
	DefenderSecurityCenterDisableHardwareUi                    bool    `json:"defenderSecurityCenterDisableHardwareUI,omitempty"`
	DefenderSecurityCenterDisableHealthUi                      any     `json:"defenderSecurityCenterDisableHealthUI,omitempty"`
	DefenderSecurityCenterDisableNetworkUi                     any     `json:"defenderSecurityCenterDisableNetworkUI,omitempty"`
	DefenderSecurityCenterDisableNotificationAreaUi            any     `json:"defenderSecurityCenterDisableNotificationAreaUI,omitempty"`
	DefenderSecurityCenterDisableRansomwareUi                  any     `json:"defenderSecurityCenterDisableRansomwareUI,omitempty"`
	DefenderSecurityCenterDisableSecureBootUi                  bool    `json:"defenderSecurityCenterDisableSecureBootUI,omitempty"`
	DefenderSecurityCenterDisableTroubleshootingUi             bool    `json:"defenderSecurityCenterDisableTroubleshootingUI,omitempty"`
	DefenderSecurityCenterDisableVirusUi                       any     `json:"defenderSecurityCenterDisableVirusUI,omitempty"`
	DefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUi any     `json:"defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI,omitempty"`
	DefenderSecurityCenterHelpEmail                            any     `json:"defenderSecurityCenterHelpEmail,omitempty"`
	DefenderSecurityCenterHelpPhone                            string  `json:"defenderSecurityCenterHelpPhone,omitempty"`
	DefenderSecurityCenterHelpURL                              any     `json:"defenderSecurityCenterHelpURL,omitempty"`
	DefenderSecurityCenterItContactDisplay                     string  `json:"defenderSecurityCenterITContactDisplay,omitempty"`
	DefenderSecurityCenterNotificationsFromApp                 string  `json:"defenderSecurityCenterNotificationsFromApp,omitempty"`
	DefenderSecurityCenterOrganizationDisplayName              any     `json:"defenderSecurityCenterOrganizationDisplayName,omitempty"`
	DefenderSignatureUpdateIntervalInHours                     any     `json:"defenderSignatureUpdateIntervalInHours,omitempty"`
	DefenderSubmitSamplesConsentType                           any     `json:"defenderSubmitSamplesConsentType,omitempty"`
	DefenderSystemScanSchedule                                 string  `json:"defenderSystemScanSchedule,omitempty"`
	DefenderUntrustedExecutable                                string  `json:"defenderUntrustedExecutable,omitempty"`
	DefenderUntrustedExecutableType                            string  `json:"defenderUntrustedExecutableType,omitempty"`
	DefenderUntrustedUsbProcess                                string  `json:"defenderUntrustedUSBProcess,omitempty"`
	DefenderUntrustedUsbProcessType                            string  `json:"defenderUntrustedUSBProcessType,omitempty"`
	DefinitionLookupBlocked                                    bool    `json:"definitionLookupBlocked,omitempty"`
	DeliveryOptimizationMode                                   string  `json:"deliveryOptimizationMode,omitempty"`
	Description                                                string  `json:"description,omitempty"`
	DesiredOSVersion                                           any     `json:"desiredOsVersion,omitempty"`
	DestinationStore                                           string  `json:"destinationStore,omitempty"`
	DeveloperUnlockSetting                                     string  `json:"developerUnlockSetting,omitempty"`
	DeviceBlockEnableRestrictions                              bool    `json:"deviceBlockEnableRestrictions,omitempty"`
	DeviceBlockEraseContentAndSettings                         bool    `json:"deviceBlockEraseContentAndSettings,omitempty"`
	DeviceBlockNameModification                                bool    `json:"deviceBlockNameModification,omitempty"`
	DeviceGuardEnableSecureBootWithDma                         bool    `json:"deviceGuardEnableSecureBootWithDMA,omitempty"`
	DeviceGuardEnableVirtualizationBasedSecurity               bool    `json:"deviceGuardEnableVirtualizationBasedSecurity,omitempty"`
	DeviceGuardLaunchSystemGuard                               string  `json:"deviceGuardLaunchSystemGuard,omitempty"`
	DeviceGuardLocalSystemAuthorityCredentialGuardSettings     string  `json:"deviceGuardLocalSystemAuthorityCredentialGuardSettings,omitempty"`
	DeviceGuardSecureBootWithDma                               string  `json:"deviceGuardSecureBootWithDMA,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode                any     `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOSEdition                 any     `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOSVersion                 any     `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceManagementBlockFactoryResetOnMobile                  bool    `json:"deviceManagementBlockFactoryResetOnMobile,omitempty"`
	DeviceManagementBlockManualUnenroll                        bool    `json:"deviceManagementBlockManualUnenroll,omitempty"`
	DiagnosticDataBlockSubmission                              bool    `json:"diagnosticDataBlockSubmission,omitempty"`
	DiagnosticDataBlockSubmissionModification                  bool    `json:"diagnosticDataBlockSubmissionModification,omitempty"`
	DiagnosticsDataSubmissionMode                              string  `json:"diagnosticsDataSubmissionMode,omitempty"`
	DisableOnDemandUserOverride                                bool    `json:"disableOnDemandUserOverride,omitempty"`
	DisconnectOnIdle                                           bool    `json:"disconnectOnIdle,omitempty"`
	DisconnectOnIdleTimerInSeconds                             any     `json:"disconnectOnIdleTimerInSeconds,omitempty"`
	DisplayAppListWithGdiDpiScalingTurnedOff                   []any   `json:"displayAppListWithGdiDPIScalingTurnedOff,omitempty"`
	DisplayAppListWithGdiDpiScalingTurnedOn                    []any   `json:"displayAppListWithGdiDPIScalingTurnedOn,omitempty"`
	DisplayName                                                string  `json:"displayName,omitempty"`
	DmaGuardDeviceEnumerationPolicy                            string  `json:"dmaGuardDeviceEnumerationPolicy,omitempty"`
	DocumentsBlockManagedDocumentsInUnmanagedApps              bool    `json:"documentsBlockManagedDocumentsInUnmanagedApps,omitempty"`
	DocumentsBlockUnmanagedDocumentsInManagedApps              bool    `json:"documentsBlockUnmanagedDocumentsInManagedApps,omitempty"`
	DriversExcluded                                            bool    `json:"driversExcluded,omitempty"`
	EdgeAllowStartPagesModification                            bool    `json:"edgeAllowStartPagesModification,omitempty"`
	EdgeBlockAccessToAboutFlags                                bool    `json:"edgeBlockAccessToAboutFlags,omitempty"`
	EdgeBlockAddressBarDropdown                                bool    `json:"edgeBlockAddressBarDropdown,omitempty"`
	EdgeBlockAutofill                                          bool    `json:"edgeBlockAutofill,omitempty"`
	EdgeBlockCompatibilityList                                 bool    `json:"edgeBlockCompatibilityList,omitempty"`
	EdgeBlockDeveloperTools                                    bool    `json:"edgeBlockDeveloperTools,omitempty"`
	EdgeBlockEditFavorites                                     bool    `json:"edgeBlockEditFavorites,omitempty"`
	EdgeBlockExtensions                                        bool    `json:"edgeBlockExtensions,omitempty"`
	EdgeBlockFullScreenMode                                    bool    `json:"edgeBlockFullScreenMode,omitempty"`
	EdgeBlockInPrivateBrowsing                                 bool    `json:"edgeBlockInPrivateBrowsing,omitempty"`
	EdgeBlockJavaScript                                        bool    `json:"edgeBlockJavaScript,omitempty"`
	EdgeBlockLiveTileDataCollection                            bool    `json:"edgeBlockLiveTileDataCollection,omitempty"`
	EdgeBlockPasswordManager                                   bool    `json:"edgeBlockPasswordManager,omitempty"`
	EdgeBlockPopups                                            bool    `json:"edgeBlockPopups,omitempty"`
	EdgeBlockPrelaunch                                         bool    `json:"edgeBlockPrelaunch,omitempty"`
	EdgeBlockPrinting                                          bool    `json:"edgeBlockPrinting,omitempty"`
	EdgeBlockSavingHistory                                     bool    `json:"edgeBlockSavingHistory,omitempty"`
	EdgeBlockSearchEngineCustomization                         bool    `json:"edgeBlockSearchEngineCustomization,omitempty"`
	EdgeBlockSearchSuggestions                                 bool    `json:"edgeBlockSearchSuggestions,omitempty"`
	EdgeBlockSendingDoNotTrackHeader                           bool    `json:"edgeBlockSendingDoNotTrackHeader,omitempty"`
	EdgeBlockSendingIntranetTrafficToInternetExplorer          bool    `json:"edgeBlockSendingIntranetTrafficToInternetExplorer,omitempty"`
	EdgeBlockSideloadingExtensions                             bool    `json:"edgeBlockSideloadingExtensions,omitempty"`
	EdgeBlockTabPreloading                                     bool    `json:"edgeBlockTabPreloading,omitempty"`
	EdgeBlockWebContentOnNewTabPage                            bool    `json:"edgeBlockWebContentOnNewTabPage,omitempty"`
	EdgeBlocked                                                bool    `json:"edgeBlocked,omitempty"`
	EdgeClearBrowsingDataOnExit                                bool    `json:"edgeClearBrowsingDataOnExit,omitempty"`
	EdgeCookiePolicy                                           string  `json:"edgeCookiePolicy,omitempty"`
	EdgeDisableFirstRunPage                                    bool    `json:"edgeDisableFirstRunPage,omitempty"`
	EdgeEnterpriseModeSiteListLocation                         any     `json:"edgeEnterpriseModeSiteListLocation,omitempty"`
	EdgeFavoritesBarVisibility                                 string  `json:"edgeFavoritesBarVisibility,omitempty"`
	EdgeFavoritesListLocation                                  any     `json:"edgeFavoritesListLocation,omitempty"`
	EdgeFirstRunURL                                            any     `json:"edgeFirstRunUrl,omitempty"`
	EdgeHomeButtonConfiguration                                any     `json:"edgeHomeButtonConfiguration,omitempty"`
	EdgeHomeButtonConfigurationEnabled                         bool    `json:"edgeHomeButtonConfigurationEnabled,omitempty"`
	EdgeHomepageUrls                                           []any   `json:"edgeHomepageUrls,omitempty"`
	EdgeKioskModeRestriction                                   string  `json:"edgeKioskModeRestriction,omitempty"`
	EdgeKioskResetAfterIdleTimeInMinutes                       any     `json:"edgeKioskResetAfterIdleTimeInMinutes,omitempty"`
	EdgeNewTabPageURL                                          any     `json:"edgeNewTabPageURL,omitempty"`
	EdgeOpensWith                                              string  `json:"edgeOpensWith,omitempty"`
	EdgePreventCertificateErrorOverride                        bool    `json:"edgePreventCertificateErrorOverride,omitempty"`
	EdgeRequireSmartScreen                                     bool    `json:"edgeRequireSmartScreen,omitempty"`
	EdgeRequiredExtensionPackageFamilyNames                    []any   `json:"edgeRequiredExtensionPackageFamilyNames,omitempty"`
	EdgeSearchEngine                                           any     `json:"edgeSearchEngine,omitempty"`
	EdgeSendIntranetTrafficToInternetExplorer                  bool    `json:"edgeSendIntranetTrafficToInternetExplorer,omitempty"`
	EdgeShowMessageWhenOpeningInternetExplorerSites            string  `json:"edgeShowMessageWhenOpeningInternetExplorerSites,omitempty"`
	EdgeSyncFavoritesWithInternetExplorer                      bool    `json:"edgeSyncFavoritesWithInternetExplorer,omitempty"`
	EdgeTelemetryForMicrosoft365Analytics                      string  `json:"edgeTelemetryForMicrosoft365Analytics,omitempty"`
	EmailInDomainSuffixes                                      []any   `json:"emailInDomainSuffixes,omitempty"`
	EnableAutomaticRedeployment                                bool    `json:"enableAutomaticRedeployment,omitempty"`
	EnablePerApp                                               bool    `json:"enablePerApp,omitempty"`
	EnableSplitTunneling                                       bool    `json:"enableSplitTunneling,omitempty"`
	EnergySaverOnBatteryThresholdPercentage                    any     `json:"energySaverOnBatteryThresholdPercentage,omitempty"`
	EnergySaverPluggedInThresholdPercentage                    any     `json:"energySaverPluggedInThresholdPercentage,omitempty"`
	EnforcedSoftwareUpdateDelayInDays                          any     `json:"enforcedSoftwareUpdateDelayInDays,omitempty"`
	EngagedRestartDeadlineInDays                               any     `json:"engagedRestartDeadlineInDays,omitempty"`
	EngagedRestartSnoozeScheduleInDays                         any     `json:"engagedRestartSnoozeScheduleInDays,omitempty"`
	EngagedRestartTransitionScheduleInDays                     any     `json:"engagedRestartTransitionScheduleInDays,omitempty"`
	EnhancedAntiSpoofingForFacialFeaturesEnabled               bool    `json:"enhancedAntiSpoofingForFacialFeaturesEnabled,omitempty"`
	EnterpriseAppBlockTrust                                    bool    `json:"enterpriseAppBlockTrust,omitempty"`
	EnterpriseAppBlockTrustModification                        bool    `json:"enterpriseAppBlockTrustModification,omitempty"`
	EnterpriseBookBlockBackup                                  bool    `json:"enterpriseBookBlockBackup,omitempty"`
	EnterpriseBookBlockMetadataSync                            bool    `json:"enterpriseBookBlockMetadataSync,omitempty"`
	EnterpriseCloudPrintDiscoveryEndPoint                      any     `json:"enterpriseCloudPrintDiscoveryEndPoint,omitempty"`
	EnterpriseCloudPrintDiscoveryMaxLimit                      any     `json:"enterpriseCloudPrintDiscoveryMaxLimit,omitempty"`
	EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier      any     `json:"enterpriseCloudPrintMopriaDiscoveryResourceIdentifier,omitempty"`
	EnterpriseCloudPrintOAuthAuthority                         any     `json:"enterpriseCloudPrintOAuthAuthority,omitempty"`
	EnterpriseCloudPrintOAuthClientIdentifier                  any     `json:"enterpriseCloudPrintOAuthClientIdentifier,omitempty"`
	EnterpriseCloudPrintResourceIdentifier                     any     `json:"enterpriseCloudPrintResourceIdentifier,omitempty"`
	EsimBlockModification                                      bool    `json:"esimBlockModification,omitempty"`
	ExcludeList                                                []any   `json:"excludeList,omitempty"`
	ExcludedDomains                                            []any   `json:"excludedDomains,omitempty"`
	ExperienceBlockDeviceDiscovery                             bool    `json:"experienceBlockDeviceDiscovery,omitempty"`
	ExperienceBlockErrorDialogWhenNoSim                        bool    `json:"experienceBlockErrorDialogWhenNoSIM,omitempty"`
	ExperienceBlockTaskSwitcher                                bool    `json:"experienceBlockTaskSwitcher,omitempty"`
	ExperienceDoNotSyncBrowserSettings                         string  `json:"experienceDoNotSyncBrowserSettings,omitempty"`
	ExtendedKeyUsages                                          []struct {
		Name             string `json:"name,omitempty"`
		ObjectIdentifier string `json:"objectIdentifier,omitempty"`
	} `json:"extendedKeyUsages,omitempty"`
	FaceTimeBlocked                                                              bool      `json:"faceTimeBlocked,omitempty"`
	FeatureUpdatesDeferralPeriodInDays                                           float64   `json:"featureUpdatesDeferralPeriodInDays,omitempty"`
	FeatureUpdatesPauseExpiryDateTime                                            time.Time `json:"featureUpdatesPauseExpiryDateTime,omitempty"`
	FeatureUpdatesPauseStartDate                                                 any       `json:"featureUpdatesPauseStartDate,omitempty"`
	FeatureUpdatesPaused                                                         bool      `json:"featureUpdatesPaused,omitempty"`
	FeatureUpdatesRollbackStartDateTime                                          time.Time `json:"featureUpdatesRollbackStartDateTime,omitempty"`
	FeatureUpdatesRollbackWindowInDays                                           float64   `json:"featureUpdatesRollbackWindowInDays,omitempty"`
	FeatureUpdatesWillBeRolledBack                                               bool      `json:"featureUpdatesWillBeRolledBack,omitempty"`
	FilesNetworkDriveAccessBlocked                                               bool      `json:"filesNetworkDriveAccessBlocked,omitempty"`
	FilesUsbDriveAccessBlocked                                                   bool      `json:"filesUsbDriveAccessBlocked,omitempty"`
	FindMyDeviceInFindMyAppBlocked                                               bool      `json:"findMyDeviceInFindMyAppBlocked,omitempty"`
	FindMyFiles                                                                  string    `json:"findMyFiles,omitempty"`
	FindMyFriendsBlocked                                                         bool      `json:"findMyFriendsBlocked,omitempty"`
	FindMyFriendsInFindMyAppBlocked                                              bool      `json:"findMyFriendsInFindMyAppBlocked,omitempty"`
	FirewallBlockStatefulFtp                                                     any       `json:"firewallBlockStatefulFTP,omitempty"`
	FirewallCertificateRevocationListCheckMethod                                 string    `json:"firewallCertificateRevocationListCheckMethod,omitempty"`
	FirewallIpSecExemptionsAllowDhcp                                             bool      `json:"firewallIPSecExemptionsAllowDHCP,omitempty"`
	FirewallIpSecExemptionsAllowIcmp                                             bool      `json:"firewallIPSecExemptionsAllowICMP,omitempty"`
	FirewallIpSecExemptionsAllowNeighborDiscovery                                bool      `json:"firewallIPSecExemptionsAllowNeighborDiscovery,omitempty"`
	FirewallIpSecExemptionsAllowRouterDiscovery                                  bool      `json:"firewallIPSecExemptionsAllowRouterDiscovery,omitempty"`
	FirewallIpSecExemptionsNone                                                  bool      `json:"firewallIPSecExemptionsNone,omitempty"`
	FirewallIdleTimeoutForSecurityAssociationInSeconds                           any       `json:"firewallIdleTimeoutForSecurityAssociationInSeconds,omitempty"`
	FirewallMergeKeyingModuleSettings                                            any       `json:"firewallMergeKeyingModuleSettings,omitempty"`
	FirewallPacketQueueingMethod                                                 string    `json:"firewallPacketQueueingMethod,omitempty"`
	FirewallPreSharedKeyEncodingMethod                                           string    `json:"firewallPreSharedKeyEncodingMethod,omitempty"`
	FirewallProfileDomain                                                        any       `json:"firewallProfileDomain,omitempty"`
	FirewallProfilePrivate                                                       any       `json:"firewallProfilePrivate,omitempty"`
	FirewallProfilePublic                                                        any       `json:"firewallProfilePublic,omitempty"`
	FirewallRules                                                                []any     `json:"firewallRules,omitempty"`
	ForegroundDownloadFromHTTPDelayInSeconds                                     float64   `json:"foregroundDownloadFromHttpDelayInSeconds,omitempty"`
	GameCenterBlocked                                                            bool      `json:"gameCenterBlocked,omitempty"`
	GameDvrBlocked                                                               bool      `json:"gameDvrBlocked,omitempty"`
	GamingBlockGameCenterFriends                                                 bool      `json:"gamingBlockGameCenterFriends,omitempty"`
	GamingBlockMultiplayer                                                       bool      `json:"gamingBlockMultiplayer,omitempty"`
	GroupIDSource                                                                any       `json:"groupIdSource,omitempty"`
	HashAlgorithm                                                                string    `json:"hashAlgorithm,omitempty"`
	HomeScreenDockIcons                                                          []any     `json:"homeScreenDockIcons,omitempty"`
	HomeScreenGridHeight                                                         any       `json:"homeScreenGridHeight,omitempty"`
	HomeScreenGridWidth                                                          any       `json:"homeScreenGridWidth,omitempty"`
	HomeScreenPages                                                              []any     `json:"homeScreenPages,omitempty"`
	HostPairingBlocked                                                           bool      `json:"hostPairingBlocked,omitempty"`
	IBooksStoreBlockErotica                                                      bool      `json:"iBooksStoreBlockErotica,omitempty"`
	IBooksStoreBlocked                                                           bool      `json:"iBooksStoreBlocked,omitempty"`
	ICloudBlockActivityContinuation                                              bool      `json:"iCloudBlockActivityContinuation,omitempty"`
	ICloudBlockBackup                                                            bool      `json:"iCloudBlockBackup,omitempty"`
	ICloudBlockDocumentSync                                                      bool      `json:"iCloudBlockDocumentSync,omitempty"`
	ICloudBlockManagedAppsSync                                                   bool      `json:"iCloudBlockManagedAppsSync,omitempty"`
	ICloudBlockPhotoLibrary                                                      bool      `json:"iCloudBlockPhotoLibrary,omitempty"`
	ICloudBlockPhotoStreamSync                                                   bool      `json:"iCloudBlockPhotoStreamSync,omitempty"`
	ICloudBlockSharedPhotoStream                                                 bool      `json:"iCloudBlockSharedPhotoStream,omitempty"`
	ICloudPrivateRelayBlocked                                                    bool      `json:"iCloudPrivateRelayBlocked,omitempty"`
	ICloudRequireEncryptedBackup                                                 bool      `json:"iCloudRequireEncryptedBackup,omitempty"`
	ITunesBlockExplicitContent                                                   bool      `json:"iTunesBlockExplicitContent,omitempty"`
	ITunesBlockMusicService                                                      bool      `json:"iTunesBlockMusicService,omitempty"`
	ITunesBlockRadio                                                             bool      `json:"iTunesBlockRadio,omitempty"`
	ITunesBlocked                                                                bool      `json:"iTunesBlocked,omitempty"`
	ID                                                                           string    `json:"id,omitempty"`
	Identifier                                                                   string    `json:"identifier,omitempty"`
	InkWorkspaceAccess                                                           string    `json:"inkWorkspaceAccess,omitempty"`
	InkWorkspaceAccessState                                                      string    `json:"inkWorkspaceAccessState,omitempty"`
	InkWorkspaceBlockSuggestedApps                                               bool      `json:"inkWorkspaceBlockSuggestedApps,omitempty"`
	InstallationSchedule                                                         any       `json:"installationSchedule,omitempty"`
	InternetSharingBlocked                                                       bool      `json:"internetSharingBlocked,omitempty"`
	IosSingleSignOnExtension                                                     any       `json:"iosSingleSignOnExtension,omitempty"`
	IsEnabled                                                                    bool      `json:"isEnabled,omitempty"`
	KeySize                                                                      string    `json:"keySize,omitempty"`
	KeyStorageProvider                                                           string    `json:"keyStorageProvider,omitempty"`
	KeyUsage                                                                     string    `json:"keyUsage,omitempty"`
	KeyboardBlockAutoCorrect                                                     bool      `json:"keyboardBlockAutoCorrect,omitempty"`
	KeyboardBlockDictation                                                       bool      `json:"keyboardBlockDictation,omitempty"`
	KeyboardBlockPredictive                                                      bool      `json:"keyboardBlockPredictive,omitempty"`
	KeyboardBlockShortcuts                                                       bool      `json:"keyboardBlockShortcuts,omitempty"`
	KeyboardBlockSpellCheck                                                      bool      `json:"keyboardBlockSpellCheck,omitempty"`
	KeychainBlockCloudSync                                                       bool      `json:"keychainBlockCloudSync,omitempty"`
	KioskModeAllowAssistiveSpeak                                                 bool      `json:"kioskModeAllowAssistiveSpeak,omitempty"`
	KioskModeAllowAssistiveTouchSettings                                         bool      `json:"kioskModeAllowAssistiveTouchSettings,omitempty"`
	KioskModeAllowAutoLock                                                       bool      `json:"kioskModeAllowAutoLock,omitempty"`
	KioskModeAllowColorInversionSettings                                         bool      `json:"kioskModeAllowColorInversionSettings,omitempty"`
	KioskModeAllowRingerSwitch                                                   bool      `json:"kioskModeAllowRingerSwitch,omitempty"`
	KioskModeAllowScreenRotation                                                 bool      `json:"kioskModeAllowScreenRotation,omitempty"`
	KioskModeAllowSleepButton                                                    bool      `json:"kioskModeAllowSleepButton,omitempty"`
	KioskModeAllowTouchscreen                                                    bool      `json:"kioskModeAllowTouchscreen,omitempty"`
	KioskModeAllowVoiceControlModification                                       bool      `json:"kioskModeAllowVoiceControlModification,omitempty"`
	KioskModeAllowVoiceOverSettings                                              bool      `json:"kioskModeAllowVoiceOverSettings,omitempty"`
	KioskModeAllowVolumeButtons                                                  bool      `json:"kioskModeAllowVolumeButtons,omitempty"`
	KioskModeAllowZoomSettings                                                   bool      `json:"kioskModeAllowZoomSettings,omitempty"`
	KioskModeAppStoreURL                                                         any       `json:"kioskModeAppStoreUrl,omitempty"`
	KioskModeAppType                                                             string    `json:"kioskModeAppType,omitempty"`
	KioskModeBlockAutoLock                                                       bool      `json:"kioskModeBlockAutoLock,omitempty"`
	KioskModeBlockRingerSwitch                                                   bool      `json:"kioskModeBlockRingerSwitch,omitempty"`
	KioskModeBlockScreenRotation                                                 bool      `json:"kioskModeBlockScreenRotation,omitempty"`
	KioskModeBlockSleepButton                                                    bool      `json:"kioskModeBlockSleepButton,omitempty"`
	KioskModeBlockTouchscreen                                                    bool      `json:"kioskModeBlockTouchscreen,omitempty"`
	KioskModeBlockVolumeButtons                                                  bool      `json:"kioskModeBlockVolumeButtons,omitempty"`
	KioskModeBuiltInAppID                                                        any       `json:"kioskModeBuiltInAppId,omitempty"`
	KioskModeEnableVoiceControl                                                  bool      `json:"kioskModeEnableVoiceControl,omitempty"`
	KioskModeManagedAppID                                                        any       `json:"kioskModeManagedAppId,omitempty"`
	KioskModeRequireAssistiveTouch                                               bool      `json:"kioskModeRequireAssistiveTouch,omitempty"`
	KioskModeRequireColorInversion                                               bool      `json:"kioskModeRequireColorInversion,omitempty"`
	KioskModeRequireMonoAudio                                                    bool      `json:"kioskModeRequireMonoAudio,omitempty"`
	KioskModeRequireVoiceOver                                                    bool      `json:"kioskModeRequireVoiceOver,omitempty"`
	KioskModeRequireZoom                                                         bool      `json:"kioskModeRequireZoom,omitempty"`
	LanManagerAuthenticationLevel                                                string    `json:"lanManagerAuthenticationLevel,omitempty"`
	LanManagerWorkstationDisableInsecureGuestLogons                              bool      `json:"lanManagerWorkstationDisableInsecureGuestLogons,omitempty"`
	LastModifiedDateTime                                                         time.Time `json:"lastModifiedDateTime,omitempty"`
	License                                                                      string    `json:"license,omitempty"`
	LicenseType                                                                  string    `json:"licenseType,omitempty"`
	LocalSecurityOptionsAdministratorAccountName                                 string    `json:"localSecurityOptionsAdministratorAccountName,omitempty"`
	LocalSecurityOptionsAdministratorElevationPromptBehavior                     string    `json:"localSecurityOptionsAdministratorElevationPromptBehavior,omitempty"`
	LocalSecurityOptionsAllowAnonymousEnumerationOfSamAccountsAndShares          bool      `json:"localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares,omitempty"`
	LocalSecurityOptionsAllowPku2UAuthenticationRequests                         bool      `json:"localSecurityOptionsAllowPKU2UAuthenticationRequests,omitempty"`
	LocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager                string    `json:"localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager,omitempty"`
	LocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool      bool      `json:"localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool,omitempty"`
	LocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn              bool      `json:"localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn,omitempty"`
	LocalSecurityOptionsAllowUiAccessApplicationElevation                        bool      `json:"localSecurityOptionsAllowUIAccessApplicationElevation,omitempty"`
	LocalSecurityOptionsAllowUiAccessApplicationsForSecureLocations              bool      `json:"localSecurityOptionsAllowUIAccessApplicationsForSecureLocations,omitempty"`
	LocalSecurityOptionsAllowUndockWithoutHavingToLogon                          bool      `json:"localSecurityOptionsAllowUndockWithoutHavingToLogon,omitempty"`
	LocalSecurityOptionsBlockMicrosoftAccounts                                   bool      `json:"localSecurityOptionsBlockMicrosoftAccounts,omitempty"`
	LocalSecurityOptionsBlockRemoteLogonWithBlankPassword                        bool      `json:"localSecurityOptionsBlockRemoteLogonWithBlankPassword,omitempty"`
	LocalSecurityOptionsBlockRemoteOpticalDriveAccess                            bool      `json:"localSecurityOptionsBlockRemoteOpticalDriveAccess,omitempty"`
	LocalSecurityOptionsBlockUsersInstallingPrinterDrivers                       bool      `json:"localSecurityOptionsBlockUsersInstallingPrinterDrivers,omitempty"`
	LocalSecurityOptionsClearVirtualMemoryPageFile                               bool      `json:"localSecurityOptionsClearVirtualMemoryPageFile,omitempty"`
	LocalSecurityOptionsClientDigitallySignCommunicationsAlways                  bool      `json:"localSecurityOptionsClientDigitallySignCommunicationsAlways,omitempty"`
	LocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySmbServers      bool      `json:"localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers,omitempty"`
	LocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation      bool      `json:"localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation,omitempty"`
	LocalSecurityOptionsDisableAdministratorAccount                              bool      `json:"localSecurityOptionsDisableAdministratorAccount,omitempty"`
	LocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees   bool      `json:"localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees,omitempty"`
	LocalSecurityOptionsDisableGuestAccount                                      bool      `json:"localSecurityOptionsDisableGuestAccount,omitempty"`
	LocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways           bool      `json:"localSecurityOptionsDisableServerDigitallySignCommunicationsAlways,omitempty"`
	LocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees   bool      `json:"localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees,omitempty"`
	LocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSamAccounts              bool      `json:"localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts,omitempty"`
	LocalSecurityOptionsDoNotRequireCtrlAltDel                                   bool      `json:"localSecurityOptionsDoNotRequireCtrlAltDel,omitempty"`
	LocalSecurityOptionsDoNotStoreLanManagerHashValueOnNextPasswordChange        bool      `json:"localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange,omitempty"`
	LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser                string    `json:"localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser,omitempty"`
	LocalSecurityOptionsGuestAccountName                                         string    `json:"localSecurityOptionsGuestAccountName,omitempty"`
	LocalSecurityOptionsHideLastSignedInUser                                     bool      `json:"localSecurityOptionsHideLastSignedInUser,omitempty"`
	LocalSecurityOptionsHideUsernameAtSignIn                                     bool      `json:"localSecurityOptionsHideUsernameAtSignIn,omitempty"`
	LocalSecurityOptionsInformationDisplayedOnLockScreen                         string    `json:"localSecurityOptionsInformationDisplayedOnLockScreen,omitempty"`
	LocalSecurityOptionsInformationShownOnLockScreen                             string    `json:"localSecurityOptionsInformationShownOnLockScreen,omitempty"`
	LocalSecurityOptionsLogOnMessageText                                         any       `json:"localSecurityOptionsLogOnMessageText,omitempty"`
	LocalSecurityOptionsLogOnMessageTitle                                        any       `json:"localSecurityOptionsLogOnMessageTitle,omitempty"`
	LocalSecurityOptionsMachineInactivityLimit                                   float64   `json:"localSecurityOptionsMachineInactivityLimit,omitempty"`
	LocalSecurityOptionsMachineInactivityLimitInMinutes                          float64   `json:"localSecurityOptionsMachineInactivityLimitInMinutes,omitempty"`
	LocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients             string    `json:"localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients,omitempty"`
	LocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers             string    `json:"localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers,omitempty"`
	LocalSecurityOptionsOnlyElevateSignedExecutables                             bool      `json:"localSecurityOptionsOnlyElevateSignedExecutables,omitempty"`
	LocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares             bool      `json:"localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares,omitempty"`
	LocalSecurityOptionsSmartCardRemovalBehavior                                 string    `json:"localSecurityOptionsSmartCardRemovalBehavior,omitempty"`
	LocalSecurityOptionsStandardUserElevationPromptBehavior                      string    `json:"localSecurityOptionsStandardUserElevationPromptBehavior,omitempty"`
	LocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation           bool      `json:"localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation,omitempty"`
	LocalSecurityOptionsUseAdminApprovalMode                                     bool      `json:"localSecurityOptionsUseAdminApprovalMode,omitempty"`
	LocalSecurityOptionsUseAdminApprovalModeForAdministrators                    bool      `json:"localSecurityOptionsUseAdminApprovalModeForAdministrators,omitempty"`
	LocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations bool      `json:"localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations,omitempty"`
	LocationServicesBlocked                                                      bool      `json:"locationServicesBlocked,omitempty"`
	LockScreenActivateAppsWithVoice                                              string    `json:"lockScreenActivateAppsWithVoice,omitempty"`
	LockScreenAllowTimeoutConfiguration                                          bool      `json:"lockScreenAllowTimeoutConfiguration,omitempty"`
	LockScreenBlockActionCenterNotifications                                     bool      `json:"lockScreenBlockActionCenterNotifications,omitempty"`
	LockScreenBlockControlCenter                                                 bool      `json:"lockScreenBlockControlCenter,omitempty"`
	LockScreenBlockCortana                                                       bool      `json:"lockScreenBlockCortana,omitempty"`
	LockScreenBlockNotificationView                                              bool      `json:"lockScreenBlockNotificationView,omitempty"`
	LockScreenBlockPassbook                                                      bool      `json:"lockScreenBlockPassbook,omitempty"`
	LockScreenBlockToastNotifications                                            bool      `json:"lockScreenBlockToastNotifications,omitempty"`
	LockScreenBlockTodayView                                                     bool      `json:"lockScreenBlockTodayView,omitempty"`
	LockScreenFootnote                                                           string    `json:"lockScreenFootnote,omitempty"`
	LockScreenTimeoutInSeconds                                                   any       `json:"lockScreenTimeoutInSeconds,omitempty"`
	LoginGroupOrDomain                                                           any       `json:"loginGroupOrDomain,omitempty"`
	LogonBlockFastUserSwitching                                                  bool      `json:"logonBlockFastUserSwitching,omitempty"`
	ManagedPasteboardRequired                                                    bool      `json:"managedPasteboardRequired,omitempty"`
	MaximumCacheAgeInDays                                                        float64   `json:"maximumCacheAgeInDays,omitempty"`
	MaximumCacheSize                                                             *struct {
		Odata_Type                 string  `json:"@odata.type,omitempty"`
		MaximumCacheSizePercentage float64 `json:"maximumCacheSizePercentage,omitempty"`
	} `json:"maximumCacheSize,omitempty"`
	MediaContentRatingApps                  string  `json:"mediaContentRatingApps,omitempty"`
	MediaContentRatingAustralia             any     `json:"mediaContentRatingAustralia,omitempty"`
	MediaContentRatingCanada                any     `json:"mediaContentRatingCanada,omitempty"`
	MediaContentRatingFrance                any     `json:"mediaContentRatingFrance,omitempty"`
	MediaContentRatingGermany               any     `json:"mediaContentRatingGermany,omitempty"`
	MediaContentRatingIreland               any     `json:"mediaContentRatingIreland,omitempty"`
	MediaContentRatingJapan                 any     `json:"mediaContentRatingJapan,omitempty"`
	MediaContentRatingNewZealand            any     `json:"mediaContentRatingNewZealand,omitempty"`
	MediaContentRatingUnitedKingdom         any     `json:"mediaContentRatingUnitedKingdom,omitempty"`
	MediaContentRatingUnitedStates          any     `json:"mediaContentRatingUnitedStates,omitempty"`
	MessagesBlocked                         bool    `json:"messagesBlocked,omitempty"`
	MessagingBlockMms                       bool    `json:"messagingBlockMMS,omitempty"`
	MessagingBlockRichCommunicationServices bool    `json:"messagingBlockRichCommunicationServices,omitempty"`
	MessagingBlockSync                      bool    `json:"messagingBlockSync,omitempty"`
	MicrosoftAccountBlockSettingsSync       bool    `json:"microsoftAccountBlockSettingsSync,omitempty"`
	MicrosoftAccountBlocked                 bool    `json:"microsoftAccountBlocked,omitempty"`
	MicrosoftAccountSignInAssistantSettings string  `json:"microsoftAccountSignInAssistantSettings,omitempty"`
	MicrosoftTunnelSiteID                   string  `json:"microsoftTunnelSiteId,omitempty"`
	MicrosoftUpdateServiceAllowed           bool    `json:"microsoftUpdateServiceAllowed,omitempty"`
	MinimumBatteryPercentageAllowedToUpload float64 `json:"minimumBatteryPercentageAllowedToUpload,omitempty"`
	MinimumDiskSizeAllowedToPeerInGigabytes float64 `json:"minimumDiskSizeAllowedToPeerInGigabytes,omitempty"`
	MinimumFileSizeToCacheInMegabytes       float64 `json:"minimumFileSizeToCacheInMegabytes,omitempty"`
	MinimumRamAllowedToPeerInGigabytes      float64 `json:"minimumRamAllowedToPeerInGigabytes,omitempty"`
	ModifyCacheLocation                     any     `json:"modifyCacheLocation,omitempty"`
	NetworkProxyApplySettingsDeviceWide     bool    `json:"networkProxyApplySettingsDeviceWide,omitempty"`
	NetworkProxyAutomaticConfigurationURL   any     `json:"networkProxyAutomaticConfigurationUrl,omitempty"`
	NetworkProxyDisableAutoDetect           bool    `json:"networkProxyDisableAutoDetect,omitempty"`
	NetworkProxyServer                      any     `json:"networkProxyServer,omitempty"`
	NetworkUsageRules                       []any   `json:"networkUsageRules,omitempty"`
	NfcBlocked                              bool    `json:"nfcBlocked,omitempty"`
	NotificationSettings                    []any   `json:"notificationSettings,omitempty"`
	NotificationsBlockSettingsModification  bool    `json:"notificationsBlockSettingsModification,omitempty"`
	OmaSettings                             []struct {
		Odata_Type             string `json:"@odata.type,omitempty"`
		Description            string `json:"description,omitempty"`
		DisplayName            string `json:"displayName,omitempty"`
		FileName               string `json:"fileName,omitempty"`
		IsEncrypted            bool   `json:"isEncrypted,omitempty"`
		IsReadOnly             bool   `json:"isReadOnly,omitempty"`
		OmaURI                 string `json:"omaUri,omitempty"`
		SecretReferenceValueID string `json:"secretReferenceValueId,omitempty"`
		Value                  any    `json:"value,omitempty"`
	} `json:"omaSettings,omitempty"`
	OnDemandRules []struct {
		Action                string `json:"action,omitempty"`
		DnsSearchDomains      []any  `json:"dnsSearchDomains,omitempty"`
		DnsServerAddressMatch []any  `json:"dnsServerAddressMatch,omitempty"`
		DomainAction          string `json:"domainAction,omitempty"`
		Domains               []any  `json:"domains,omitempty"`
		InterfaceTypeMatch    string `json:"interfaceTypeMatch,omitempty"`
		ProbeRequiredURL      any    `json:"probeRequiredUrl,omitempty"`
		ProbeURL              any    `json:"probeUrl,omitempty"`
		Ssids                 []any  `json:"ssids,omitempty"`
	} `json:"onDemandRules,omitempty"`
	OnDeviceOnlyDictationForced                     bool      `json:"onDeviceOnlyDictationForced,omitempty"`
	OnDeviceOnlyTranslationForced                   bool      `json:"onDeviceOnlyTranslationForced,omitempty"`
	OneDriveDisableFileSync                         bool      `json:"oneDriveDisableFileSync,omitempty"`
	OptInToDeviceIDSharing                          any       `json:"optInToDeviceIdSharing,omitempty"`
	PasscodeBlockFingerprintModification            bool      `json:"passcodeBlockFingerprintModification,omitempty"`
	PasscodeBlockFingerprintUnlock                  bool      `json:"passcodeBlockFingerprintUnlock,omitempty"`
	PasscodeBlockModification                       bool      `json:"passcodeBlockModification,omitempty"`
	PasscodeBlockSimple                             bool      `json:"passcodeBlockSimple,omitempty"`
	PasscodeExpirationDays                          float64   `json:"passcodeExpirationDays,omitempty"`
	PasscodeMinimumCharacterSetCount                float64   `json:"passcodeMinimumCharacterSetCount,omitempty"`
	PasscodeMinimumLength                           float64   `json:"passcodeMinimumLength,omitempty"`
	PasscodeMinutesOfInactivityBeforeLock           float64   `json:"passcodeMinutesOfInactivityBeforeLock,omitempty"`
	PasscodeMinutesOfInactivityBeforeScreenTimeout  float64   `json:"passcodeMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	PasscodePreviousPasscodeBlockCount              float64   `json:"passcodePreviousPasscodeBlockCount,omitempty"`
	PasscodeRequired                                bool      `json:"passcodeRequired,omitempty"`
	PasscodeRequiredType                            string    `json:"passcodeRequiredType,omitempty"`
	PasscodeSignInFailureCountBeforeWipe            float64   `json:"passcodeSignInFailureCountBeforeWipe,omitempty"`
	PasswordBlockAirDropSharing                     bool      `json:"passwordBlockAirDropSharing,omitempty"`
	PasswordBlockAutoFill                           bool      `json:"passwordBlockAutoFill,omitempty"`
	PasswordBlockProximityRequests                  bool      `json:"passwordBlockProximityRequests,omitempty"`
	PasswordBlockSimple                             bool      `json:"passwordBlockSimple,omitempty"`
	PasswordExpirationDays                          float64   `json:"passwordExpirationDays,omitempty"`
	PasswordMinimumAgeInDays                        any       `json:"passwordMinimumAgeInDays,omitempty"`
	PasswordMinimumCharacterSetCount                float64   `json:"passwordMinimumCharacterSetCount,omitempty"`
	PasswordMinimumLength                           float64   `json:"passwordMinimumLength,omitempty"`
	PasswordMinutesOfInactivityBeforeScreenTimeout  float64   `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	PasswordPreviousPasswordBlockCount              float64   `json:"passwordPreviousPasswordBlockCount,omitempty"`
	PasswordRequireWhenResumeFromIdleState          bool      `json:"passwordRequireWhenResumeFromIdleState,omitempty"`
	PasswordRequired                                bool      `json:"passwordRequired,omitempty"`
	PasswordRequiredType                            string    `json:"passwordRequiredType,omitempty"`
	PasswordSignInFailureCountBeforeFactoryReset    float64   `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
	PersonalizationDesktopImageURL                  string    `json:"personalizationDesktopImageUrl,omitempty"`
	PersonalizationLockScreenImageURL               string    `json:"personalizationLockScreenImageUrl,omitempty"`
	PinExpirationInDays                             float64   `json:"pinExpirationInDays,omitempty"`
	PinLowercaseCharactersUsage                     string    `json:"pinLowercaseCharactersUsage,omitempty"`
	PinMaximumLength                                float64   `json:"pinMaximumLength,omitempty"`
	PinMinimumLength                                float64   `json:"pinMinimumLength,omitempty"`
	PinPreviousBlockCount                           float64   `json:"pinPreviousBlockCount,omitempty"`
	PinRecoveryEnabled                              bool      `json:"pinRecoveryEnabled,omitempty"`
	PinSpecialCharactersUsage                       string    `json:"pinSpecialCharactersUsage,omitempty"`
	PinUppercaseCharactersUsage                     string    `json:"pinUppercaseCharactersUsage,omitempty"`
	PkiBlockOtaUpdates                              bool      `json:"pkiBlockOTAUpdates,omitempty"`
	PodcastsBlocked                                 bool      `json:"podcastsBlocked,omitempty"`
	PostponeRebootUntilAfterDeadline                bool      `json:"postponeRebootUntilAfterDeadline,omitempty"`
	PowerButtonActionOnBattery                      string    `json:"powerButtonActionOnBattery,omitempty"`
	PowerButtonActionPluggedIn                      string    `json:"powerButtonActionPluggedIn,omitempty"`
	PowerHybridSleepOnBattery                       string    `json:"powerHybridSleepOnBattery,omitempty"`
	PowerHybridSleepPluggedIn                       string    `json:"powerHybridSleepPluggedIn,omitempty"`
	PowerLidCloseActionOnBattery                    string    `json:"powerLidCloseActionOnBattery,omitempty"`
	PowerLidCloseActionPluggedIn                    string    `json:"powerLidCloseActionPluggedIn,omitempty"`
	PowerSleepButtonActionOnBattery                 string    `json:"powerSleepButtonActionOnBattery,omitempty"`
	PowerSleepButtonActionPluggedIn                 string    `json:"powerSleepButtonActionPluggedIn,omitempty"`
	PrereleaseFeatures                              string    `json:"prereleaseFeatures,omitempty"`
	PrinterBlockAddition                            bool      `json:"printerBlockAddition,omitempty"`
	PrinterDefaultName                              any       `json:"printerDefaultName,omitempty"`
	PrinterNames                                    []any     `json:"printerNames,omitempty"`
	PrivacyAdvertisingID                            string    `json:"privacyAdvertisingId,omitempty"`
	PrivacyAutoAcceptPairingAndConsentPrompts       bool      `json:"privacyAutoAcceptPairingAndConsentPrompts,omitempty"`
	PrivacyBlockActivityFeed                        bool      `json:"privacyBlockActivityFeed,omitempty"`
	PrivacyBlockInputPersonalization                bool      `json:"privacyBlockInputPersonalization,omitempty"`
	PrivacyBlockPublishUserActivities               bool      `json:"privacyBlockPublishUserActivities,omitempty"`
	PrivacyDisableLaunchExperience                  bool      `json:"privacyDisableLaunchExperience,omitempty"`
	PrivacyForceLimitAdTracking                     bool      `json:"privacyForceLimitAdTracking,omitempty"`
	ProductKey                                      string    `json:"productKey,omitempty"`
	ProviderType                                    any       `json:"providerType,omitempty"`
	ProximityBlockSetupToNewDevice                  bool      `json:"proximityBlockSetupToNewDevice,omitempty"`
	ProxyServer                                     any       `json:"proxyServer,omitempty"`
	QualityUpdatesDeferralPeriodInDays              float64   `json:"qualityUpdatesDeferralPeriodInDays,omitempty"`
	QualityUpdatesPauseExpiryDateTime               time.Time `json:"qualityUpdatesPauseExpiryDateTime,omitempty"`
	QualityUpdatesPauseStartDate                    any       `json:"qualityUpdatesPauseStartDate,omitempty"`
	QualityUpdatesPaused                            bool      `json:"qualityUpdatesPaused,omitempty"`
	QualityUpdatesRollbackStartDateTime             time.Time `json:"qualityUpdatesRollbackStartDateTime,omitempty"`
	QualityUpdatesWillBeRolledBack                  bool      `json:"qualityUpdatesWillBeRolledBack,omitempty"`
	Realm                                           any       `json:"realm,omitempty"`
	RenewalThresholdPercentage                      float64   `json:"renewalThresholdPercentage,omitempty"`
	ResetProtectionModeBlocked                      bool      `json:"resetProtectionModeBlocked,omitempty"`
	RestrictPeerSelectionBy                         string    `json:"restrictPeerSelectionBy,omitempty"`
	Role                                            any       `json:"role,omitempty"`
	RoleScopeTagIds                                 []string  `json:"roleScopeTagIds,omitempty"`
	SafariBlockAutofill                             bool      `json:"safariBlockAutofill,omitempty"`
	SafariBlockJavaScript                           bool      `json:"safariBlockJavaScript,omitempty"`
	SafariBlockPopups                               bool      `json:"safariBlockPopups,omitempty"`
	SafariBlocked                                   bool      `json:"safariBlocked,omitempty"`
	SafariCookieSettings                            string    `json:"safariCookieSettings,omitempty"`
	SafariDomains                                   []any     `json:"safariDomains,omitempty"`
	SafariManagedDomains                            []any     `json:"safariManagedDomains,omitempty"`
	SafariPasswordAutoFillDomains                   []any     `json:"safariPasswordAutoFillDomains,omitempty"`
	SafariRequireFraudWarning                       bool      `json:"safariRequireFraudWarning,omitempty"`
	SafeSearchFilter                                string    `json:"safeSearchFilter,omitempty"`
	ScepServerUrls                                  []string  `json:"scepServerUrls,omitempty"`
	ScheduleImminentRestartWarningInMinutes         any       `json:"scheduleImminentRestartWarningInMinutes,omitempty"`
	ScheduleRestartWarningInHours                   any       `json:"scheduleRestartWarningInHours,omitempty"`
	ScheduledInstallDays                            []any     `json:"scheduledInstallDays,omitempty"`
	ScreenCaptureBlocked                            bool      `json:"screenCaptureBlocked,omitempty"`
	SearchBlockDiacritics                           bool      `json:"searchBlockDiacritics,omitempty"`
	SearchBlockWebResults                           bool      `json:"searchBlockWebResults,omitempty"`
	SearchDisableAutoLanguageDetection              bool      `json:"searchDisableAutoLanguageDetection,omitempty"`
	SearchDisableIndexerBackoff                     bool      `json:"searchDisableIndexerBackoff,omitempty"`
	SearchDisableIndexingEncryptedItems             bool      `json:"searchDisableIndexingEncryptedItems,omitempty"`
	SearchDisableIndexingRemovableDrive             bool      `json:"searchDisableIndexingRemovableDrive,omitempty"`
	SearchDisableLocation                           bool      `json:"searchDisableLocation,omitempty"`
	SearchDisableUseLocation                        bool      `json:"searchDisableUseLocation,omitempty"`
	SearchEnableAutomaticIndexSizeManangement       bool      `json:"searchEnableAutomaticIndexSizeManangement,omitempty"`
	SearchEnableRemoteQueries                       bool      `json:"searchEnableRemoteQueries,omitempty"`
	SecurityBlockAzureAdJoinedDevicesAutoEncryption bool      `json:"securityBlockAzureADJoinedDevicesAutoEncryption,omitempty"`
	SecurityDeviceRequired                          bool      `json:"securityDeviceRequired,omitempty"`
	Server                                          *struct {
		Address         string `json:"address,omitempty"`
		Description     string `json:"description,omitempty"`
		IsDefaultServer bool   `json:"isDefaultServer,omitempty"`
	} `json:"server,omitempty"`
	SettingsBlockAccountsPage                            bool    `json:"settingsBlockAccountsPage,omitempty"`
	SettingsBlockAddProvisioningPackage                  bool    `json:"settingsBlockAddProvisioningPackage,omitempty"`
	SettingsBlockAppsPage                                bool    `json:"settingsBlockAppsPage,omitempty"`
	SettingsBlockChangeLanguage                          bool    `json:"settingsBlockChangeLanguage,omitempty"`
	SettingsBlockChangePowerSleep                        bool    `json:"settingsBlockChangePowerSleep,omitempty"`
	SettingsBlockChangeRegion                            bool    `json:"settingsBlockChangeRegion,omitempty"`
	SettingsBlockChangeSystemTime                        bool    `json:"settingsBlockChangeSystemTime,omitempty"`
	SettingsBlockDevicesPage                             bool    `json:"settingsBlockDevicesPage,omitempty"`
	SettingsBlockEaseOfAccessPage                        bool    `json:"settingsBlockEaseOfAccessPage,omitempty"`
	SettingsBlockEditDeviceName                          bool    `json:"settingsBlockEditDeviceName,omitempty"`
	SettingsBlockGamingPage                              bool    `json:"settingsBlockGamingPage,omitempty"`
	SettingsBlockNetworkInternetPage                     bool    `json:"settingsBlockNetworkInternetPage,omitempty"`
	SettingsBlockPersonalizationPage                     bool    `json:"settingsBlockPersonalizationPage,omitempty"`
	SettingsBlockPrivacyPage                             bool    `json:"settingsBlockPrivacyPage,omitempty"`
	SettingsBlockRemoveProvisioningPackage               bool    `json:"settingsBlockRemoveProvisioningPackage,omitempty"`
	SettingsBlockSettingsApp                             bool    `json:"settingsBlockSettingsApp,omitempty"`
	SettingsBlockSystemPage                              bool    `json:"settingsBlockSystemPage,omitempty"`
	SettingsBlockTimeLanguagePage                        bool    `json:"settingsBlockTimeLanguagePage,omitempty"`
	SettingsBlockUpdateSecurityPage                      bool    `json:"settingsBlockUpdateSecurityPage,omitempty"`
	SharedDeviceBlockTemporarySessions                   bool    `json:"sharedDeviceBlockTemporarySessions,omitempty"`
	SharedUserAppDataAllowed                             bool    `json:"sharedUserAppDataAllowed,omitempty"`
	SingleSignOnExtension                                any     `json:"singleSignOnExtension,omitempty"`
	SingleSignOnSettings                                 any     `json:"singleSignOnSettings,omitempty"`
	SiriBlockUserGeneratedContent                        bool    `json:"siriBlockUserGeneratedContent,omitempty"`
	SiriBlocked                                          bool    `json:"siriBlocked,omitempty"`
	SiriBlockedWhenLocked                                bool    `json:"siriBlockedWhenLocked,omitempty"`
	SiriRequireProfanityFilter                           bool    `json:"siriRequireProfanityFilter,omitempty"`
	SkipChecksBeforeRestart                              bool    `json:"skipChecksBeforeRestart,omitempty"`
	SmartScreenAppInstallControl                         string  `json:"smartScreenAppInstallControl,omitempty"`
	SmartScreenBlockOverrideForFiles                     bool    `json:"smartScreenBlockOverrideForFiles,omitempty"`
	SmartScreenBlockPromptOverride                       bool    `json:"smartScreenBlockPromptOverride,omitempty"`
	SmartScreenBlockPromptOverrideForFiles               bool    `json:"smartScreenBlockPromptOverrideForFiles,omitempty"`
	SmartScreenEnableAppInstallControl                   bool    `json:"smartScreenEnableAppInstallControl,omitempty"`
	SmartScreenEnableInShell                             bool    `json:"smartScreenEnableInShell,omitempty"`
	SoftwareUpdatesEnforcedDelayInDays                   any     `json:"softwareUpdatesEnforcedDelayInDays,omitempty"`
	SoftwareUpdatesForceDelayed                          bool    `json:"softwareUpdatesForceDelayed,omitempty"`
	SpotlightBlockInternetResults                        bool    `json:"spotlightBlockInternetResults,omitempty"`
	StartBlockUnpinningAppsFromTaskbar                   bool    `json:"startBlockUnpinningAppsFromTaskbar,omitempty"`
	StartMenuAppListVisibility                           string  `json:"startMenuAppListVisibility,omitempty"`
	StartMenuHideChangeAccountSettings                   bool    `json:"startMenuHideChangeAccountSettings,omitempty"`
	StartMenuHideFrequentlyUsedApps                      bool    `json:"startMenuHideFrequentlyUsedApps,omitempty"`
	StartMenuHideHibernate                               bool    `json:"startMenuHideHibernate,omitempty"`
	StartMenuHideLock                                    bool    `json:"startMenuHideLock,omitempty"`
	StartMenuHidePowerButton                             bool    `json:"startMenuHidePowerButton,omitempty"`
	StartMenuHideRecentJumpLists                         bool    `json:"startMenuHideRecentJumpLists,omitempty"`
	StartMenuHideRecentlyAddedApps                       bool    `json:"startMenuHideRecentlyAddedApps,omitempty"`
	StartMenuHideRestartOptions                          bool    `json:"startMenuHideRestartOptions,omitempty"`
	StartMenuHideShutDown                                bool    `json:"startMenuHideShutDown,omitempty"`
	StartMenuHideSignOut                                 bool    `json:"startMenuHideSignOut,omitempty"`
	StartMenuHideSleep                                   bool    `json:"startMenuHideSleep,omitempty"`
	StartMenuHideSwitchAccount                           bool    `json:"startMenuHideSwitchAccount,omitempty"`
	StartMenuHideUserTile                                bool    `json:"startMenuHideUserTile,omitempty"`
	StartMenuLayoutEdgeAssetsXML                         any     `json:"startMenuLayoutEdgeAssetsXml,omitempty"`
	StartMenuLayoutXML                                   string  `json:"startMenuLayoutXml,omitempty"`
	StartMenuMode                                        string  `json:"startMenuMode,omitempty"`
	StartMenuPinnedFolderDocuments                       string  `json:"startMenuPinnedFolderDocuments,omitempty"`
	StartMenuPinnedFolderDownloads                       string  `json:"startMenuPinnedFolderDownloads,omitempty"`
	StartMenuPinnedFolderFileExplorer                    string  `json:"startMenuPinnedFolderFileExplorer,omitempty"`
	StartMenuPinnedFolderHomeGroup                       string  `json:"startMenuPinnedFolderHomeGroup,omitempty"`
	StartMenuPinnedFolderMusic                           string  `json:"startMenuPinnedFolderMusic,omitempty"`
	StartMenuPinnedFolderNetwork                         string  `json:"startMenuPinnedFolderNetwork,omitempty"`
	StartMenuPinnedFolderPersonalFolder                  string  `json:"startMenuPinnedFolderPersonalFolder,omitempty"`
	StartMenuPinnedFolderPictures                        string  `json:"startMenuPinnedFolderPictures,omitempty"`
	StartMenuPinnedFolderSettings                        string  `json:"startMenuPinnedFolderSettings,omitempty"`
	StartMenuPinnedFolderVideos                          string  `json:"startMenuPinnedFolderVideos,omitempty"`
	StorageBlockRemovableStorage                         bool    `json:"storageBlockRemovableStorage,omitempty"`
	StorageRequireMobileDeviceEncryption                 bool    `json:"storageRequireMobileDeviceEncryption,omitempty"`
	StorageRestrictAppDataToSystemVolume                 bool    `json:"storageRestrictAppDataToSystemVolume,omitempty"`
	StorageRestrictAppInstallToSystemVolume              bool    `json:"storageRestrictAppInstallToSystemVolume,omitempty"`
	StrictEnforcement                                    any     `json:"strictEnforcement,omitempty"`
	SubjectAlternativeNameFormatString                   any     `json:"subjectAlternativeNameFormatString,omitempty"`
	SubjectAlternativeNameType                           string  `json:"subjectAlternativeNameType,omitempty"`
	SubjectNameFormat                                    string  `json:"subjectNameFormat,omitempty"`
	SubjectNameFormatString                              string  `json:"subjectNameFormatString,omitempty"`
	SupportsScopeTags                                    bool    `json:"supportsScopeTags,omitempty"`
	SystemTelemetryProxyServer                           any     `json:"systemTelemetryProxyServer,omitempty"`
	TargetEdition                                        string  `json:"targetEdition,omitempty"`
	TargetedMobileApps                                   []any   `json:"targetedMobileApps,omitempty"`
	TaskManagerBlockEndTask                              bool    `json:"taskManagerBlockEndTask,omitempty"`
	TenantLockdownRequireNetworkDuringOutOfBoxExperience bool    `json:"tenantLockdownRequireNetworkDuringOutOfBoxExperience,omitempty"`
	TrustedRootCertificate                               string  `json:"trustedRootCertificate,omitempty"`
	UninstallBuiltInApps                                 bool    `json:"uninstallBuiltInApps,omitempty"`
	UnlockWithBiometricsEnabled                          bool    `json:"unlockWithBiometricsEnabled,omitempty"`
	UnpairedExternalBootToRecoveryAllowed                bool    `json:"unpairedExternalBootToRecoveryAllowed,omitempty"`
	UpdateNotificationLevel                              string  `json:"updateNotificationLevel,omitempty"`
	UpdateScheduleType                                   string  `json:"updateScheduleType,omitempty"`
	UpdateWeeks                                          any     `json:"updateWeeks,omitempty"`
	UsbBlocked                                           bool    `json:"usbBlocked,omitempty"`
	UsbRestrictedModeBlocked                             bool    `json:"usbRestrictedModeBlocked,omitempty"`
	UseCertificatesForOnPremisesAuthEnabled              bool    `json:"useCertificatesForOnPremisesAuthEnabled,omitempty"`
	UseSecurityKeyForSignin                              bool    `json:"useSecurityKeyForSignin,omitempty"`
	UserDomain                                           any     `json:"userDomain,omitempty"`
	UserPauseAccess                                      string  `json:"userPauseAccess,omitempty"`
	UserRightsAccessCredentialManagerAsTrustedCaller     any     `json:"userRightsAccessCredentialManagerAsTrustedCaller,omitempty"`
	UserRightsActAsPartOfTheOperatingSystem              any     `json:"userRightsActAsPartOfTheOperatingSystem,omitempty"`
	UserRightsAllowAccessFromNetwork                     any     `json:"userRightsAllowAccessFromNetwork,omitempty"`
	UserRightsBackupData                                 any     `json:"userRightsBackupData,omitempty"`
	UserRightsBlockAccessFromNetwork                     any     `json:"userRightsBlockAccessFromNetwork,omitempty"`
	UserRightsChangeSystemTime                           any     `json:"userRightsChangeSystemTime,omitempty"`
	UserRightsCreateGlobalObjects                        any     `json:"userRightsCreateGlobalObjects,omitempty"`
	UserRightsCreatePageFile                             any     `json:"userRightsCreatePageFile,omitempty"`
	UserRightsCreatePermanentSharedObjects               any     `json:"userRightsCreatePermanentSharedObjects,omitempty"`
	UserRightsCreateSymbolicLinks                        any     `json:"userRightsCreateSymbolicLinks,omitempty"`
	UserRightsCreateToken                                any     `json:"userRightsCreateToken,omitempty"`
	UserRightsDebugPrograms                              any     `json:"userRightsDebugPrograms,omitempty"`
	UserRightsDelegation                                 any     `json:"userRightsDelegation,omitempty"`
	UserRightsDenyLocalLogOn                             any     `json:"userRightsDenyLocalLogOn,omitempty"`
	UserRightsGenerateSecurityAudits                     any     `json:"userRightsGenerateSecurityAudits,omitempty"`
	UserRightsImpersonateClient                          any     `json:"userRightsImpersonateClient,omitempty"`
	UserRightsIncreaseSchedulingPriority                 any     `json:"userRightsIncreaseSchedulingPriority,omitempty"`
	UserRightsLoadUnloadDrivers                          any     `json:"userRightsLoadUnloadDrivers,omitempty"`
	UserRightsLocalLogOn                                 any     `json:"userRightsLocalLogOn,omitempty"`
	UserRightsLockMemory                                 any     `json:"userRightsLockMemory,omitempty"`
	UserRightsManageAuditingAndSecurityLogs              any     `json:"userRightsManageAuditingAndSecurityLogs,omitempty"`
	UserRightsManageVolumes                              any     `json:"userRightsManageVolumes,omitempty"`
	UserRightsModifyFirmwareEnvironment                  any     `json:"userRightsModifyFirmwareEnvironment,omitempty"`
	UserRightsModifyObjectLabels                         any     `json:"userRightsModifyObjectLabels,omitempty"`
	UserRightsProfileSingleProcess                       any     `json:"userRightsProfileSingleProcess,omitempty"`
	UserRightsRemoteDesktopServicesLogOn                 any     `json:"userRightsRemoteDesktopServicesLogOn,omitempty"`
	UserRightsRemoteShutdown                             any     `json:"userRightsRemoteShutdown,omitempty"`
	UserRightsRestoreData                                any     `json:"userRightsRestoreData,omitempty"`
	UserRightsTakeOwnership                              any     `json:"userRightsTakeOwnership,omitempty"`
	UserWindowsUpdateScanAccess                          string  `json:"userWindowsUpdateScanAccess,omitempty"`
	UtcTimeOffsetInMinutes                               any     `json:"utcTimeOffsetInMinutes,omitempty"`
	Version                                              float64 `json:"version,omitempty"`
	VoiceDialingBlocked                                  bool    `json:"voiceDialingBlocked,omitempty"`
	VoiceRecordingBlocked                                bool    `json:"voiceRecordingBlocked,omitempty"`
	VpnBlockCreation                                     bool    `json:"vpnBlockCreation,omitempty"`
	VpnPeerCaching                                       string  `json:"vpnPeerCaching,omitempty"`
	WallpaperBlockModification                           bool    `json:"wallpaperBlockModification,omitempty"`
	WallpaperDisplayLocation                             string  `json:"wallpaperDisplayLocation,omitempty"`
	WallpaperImage                                       any     `json:"wallpaperImage,omitempty"`
	WebRtcBlockLocalhostIpAddress                        bool    `json:"webRtcBlockLocalhostIpAddress,omitempty"`
	WiFiBlockAutomaticConnectHotspots                    bool    `json:"wiFiBlockAutomaticConnectHotspots,omitempty"`
	WiFiBlockManualConfiguration                         bool    `json:"wiFiBlockManualConfiguration,omitempty"`
	WiFiBlocked                                          bool    `json:"wiFiBlocked,omitempty"`
	WiFiConnectOnlyToConfiguredNetworks                  bool    `json:"wiFiConnectOnlyToConfiguredNetworks,omitempty"`
	WiFiConnectToAllowedNetworksOnlyForced               bool    `json:"wiFiConnectToAllowedNetworksOnlyForced,omitempty"`
	WiFiScanInterval                                     any     `json:"wiFiScanInterval,omitempty"`
	WifiPowerOnForced                                    bool    `json:"wifiPowerOnForced,omitempty"`
	Windows10AppsForceUpdateSchedule                     any     `json:"windows10AppsForceUpdateSchedule,omitempty"`
	WindowsDefenderTamperProtection                      string  `json:"windowsDefenderTamperProtection,omitempty"`
	WindowsHelloForBusinessBlocked                       bool    `json:"windowsHelloForBusinessBlocked,omitempty"`
	WindowsNetworkIsolationPolicy                        *struct {
		EnterpriseCloudResources []struct {
			IpAddressOrFqdn string `json:"ipAddressOrFQDN,omitempty"`
			Proxy           any    `json:"proxy,omitempty"`
		} `json:"enterpriseCloudResources,omitempty"`
		EnterpriseIpRanges                     []any `json:"enterpriseIPRanges,omitempty"`
		EnterpriseIpRangesAreAuthoritative     bool  `json:"enterpriseIPRangesAreAuthoritative,omitempty"`
		EnterpriseInternalProxyServers         []any `json:"enterpriseInternalProxyServers,omitempty"`
		EnterpriseNetworkDomainNames           []any `json:"enterpriseNetworkDomainNames,omitempty"`
		EnterpriseProxyServers                 []any `json:"enterpriseProxyServers,omitempty"`
		EnterpriseProxyServersAreAuthoritative bool  `json:"enterpriseProxyServersAreAuthoritative,omitempty"`
		NeutralDomainResources                 []any `json:"neutralDomainResources,omitempty"`
	} `json:"windowsNetworkIsolationPolicy,omitempty"`
	WindowsSMode                                      string `json:"windowsSMode,omitempty"`
	WindowsSpotlightBlockConsumerSpecificFeatures     bool   `json:"windowsSpotlightBlockConsumerSpecificFeatures,omitempty"`
	WindowsSpotlightBlockOnActionCenter               bool   `json:"windowsSpotlightBlockOnActionCenter,omitempty"`
	WindowsSpotlightBlockTailoredExperiences          bool   `json:"windowsSpotlightBlockTailoredExperiences,omitempty"`
	WindowsSpotlightBlockThirdPartyNotifications      bool   `json:"windowsSpotlightBlockThirdPartyNotifications,omitempty"`
	WindowsSpotlightBlockWelcomeExperience            bool   `json:"windowsSpotlightBlockWelcomeExperience,omitempty"`
	WindowsSpotlightBlockWindowsTips                  bool   `json:"windowsSpotlightBlockWindowsTips,omitempty"`
	WindowsSpotlightBlocked                           bool   `json:"windowsSpotlightBlocked,omitempty"`
	WindowsSpotlightConfigureOnLockScreen             string `json:"windowsSpotlightConfigureOnLockScreen,omitempty"`
	WindowsStoreBlockAutoUpdate                       bool   `json:"windowsStoreBlockAutoUpdate,omitempty"`
	WindowsStoreBlocked                               bool   `json:"windowsStoreBlocked,omitempty"`
	WindowsStoreEnablePrivateStoreOnly                bool   `json:"windowsStoreEnablePrivateStoreOnly,omitempty"`
	WirelessDisplayBlockProjectionToThisDevice        bool   `json:"wirelessDisplayBlockProjectionToThisDevice,omitempty"`
	WirelessDisplayBlockUserInputFromReceiver         bool   `json:"wirelessDisplayBlockUserInputFromReceiver,omitempty"`
	WirelessDisplayRequirePinForPairing               bool   `json:"wirelessDisplayRequirePinForPairing,omitempty"`
	XboxServicesAccessoryManagementServiceStartupMode string `json:"xboxServicesAccessoryManagementServiceStartupMode,omitempty"`
	XboxServicesEnableXboxGameSaveTask                bool   `json:"xboxServicesEnableXboxGameSaveTask,omitempty"`
	XboxServicesLiveAuthManagerServiceStartupMode     string `json:"xboxServicesLiveAuthManagerServiceStartupMode,omitempty"`
	XboxServicesLiveGameSaveServiceStartupMode        string `json:"xboxServicesLiveGameSaveServiceStartupMode,omitempty"`
	XboxServicesLiveNetworkingServiceStartupMode      string `json:"xboxServicesLiveNetworkingServiceStartupMode,omitempty"`
}
