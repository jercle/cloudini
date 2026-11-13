package azure

import "time"

type GetManagedDevicesResponse struct {
	Context string                 `json:"@odata.context,omitempty,omitzero" bson:"@odata.context,omitempty,omitzero"`
	Count   float64                `json:"@odata.count,omitempty,omitzero" bson:"@odata.count,omitempty,omitzero"`
	Value   []ManagedDeviceMinimal `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

//
//

type ManagedDeviceMinimal struct {
	AzureAdDeviceID   string    `json:"azureADDeviceId,omitempty,omitzero" bson:"azureADDeviceId,omitempty,omitzero"`
	DeviceName        string    `json:"deviceName,omitempty,omitzero" bson:"deviceName,omitempty,omitzero"`
	ID                string    `json:"id,omitempty,omitzero" bson:"_id,omitempty,omitzero"`
	LastSyncDateTime  time.Time `json:"lastSyncDateTime,omitempty,omitzero" bson:"lastSyncDateTime,omitempty,omitzero"`
	SerialNumber      string    `json:"serialNumber,omitempty,omitzero" bson:"serialNumber,omitempty,omitzero"`
	UserPrincipalName string    `json:"userPrincipalName,omitempty,omitzero" bson:"userPrincipalName,omitempty,omitzero"`
}

//
//

type ListManagedDevicesResponse struct {
	OdataContext string          `json:"@odata.context,omitempty,omitzero" bson:"@odata.context,omitempty,omitzero"`
	OdataCount   int             `json:"@odata.count,omitempty,omitzero" bson:"@odata.count,omitempty,omitzero"`
	Value        []ManagedDevice `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

//
//

type ManagedDevice struct {
	TenantName            string    `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	TenantId              string    `json:"tenantId,omitempty,omitzero" bson:"tenantId,omitempty,omitzero"`
	LastDatabaseSync      time.Time `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
	AzureAdDeviceID       string    `json:"azureADDeviceId,omitempty,omitzero" bson:"azureADDeviceId,omitempty,omitzero"`
	DeviceName            string    `json:"deviceName,omitempty,omitzero" bson:"deviceName,omitempty,omitzero"`
	DeviceType            string    `json:"deviceType,omitempty,omitzero" bson:"deviceType,omitempty,omitzero"`
	EnrolledDateTime      time.Time `json:"enrolledDateTime,omitempty,omitzero" bson:"enrolledDateTime,omitempty,omitzero"`
	EnrollmentProfileName any       `json:"enrollmentProfileName,omitempty,omitzero" bson:"enrollmentProfileName,omitempty,omitzero"`
	EthernetMacAddress    any       `json:"ethernetMacAddress,omitempty,omitzero" bson:"ethernetMacAddress,omitempty,omitzero"`
	HardwareInformation   struct {
		BatteryChargeCycles                                            float64 `json:"batteryChargeCycles,omitempty,omitzero" bson:"batteryChargeCycles,omitempty,omitzero"`
		BatteryHealthPercentage                                        float64 `json:"batteryHealthPercentage,omitempty,omitzero" bson:"batteryHealthPercentage,omitempty,omitzero"`
		BatteryLevelPercentage                                         any     `json:"batteryLevelPercentage,omitempty,omitzero" bson:"batteryLevelPercentage,omitempty,omitzero"`
		BatterySerialNumber                                            any     `json:"batterySerialNumber,omitempty,omitzero" bson:"batterySerialNumber,omitempty,omitzero"`
		CellularTechnology                                             any     `json:"cellularTechnology,omitempty,omitzero" bson:"cellularTechnology,omitempty,omitzero"`
		DeviceFullQualifiedDomainName                                  any     `json:"deviceFullQualifiedDomainName,omitempty,omitzero" bson:"deviceFullQualifiedDomainName,omitempty,omitzero"`
		DeviceGuardLocalSystemAuthorityCredentialGuardState            string  `json:"deviceGuardLocalSystemAuthorityCredentialGuardState,omitempty,omitzero" bson:"deviceGuardLocalSystemAuthorityCredentialGuardState,omitempty,omitzero"`
		DeviceGuardVirtualizationBasedSecurityHardwareRequirementState string  `json:"deviceGuardVirtualizationBasedSecurityHardwareRequirementState,omitempty,omitzero" bson:"deviceGuardVirtualizationBasedSecurityHardwareRequirementState,omitempty,omitzero"`
		DeviceGuardVirtualizationBasedSecurityState                    string  `json:"deviceGuardVirtualizationBasedSecurityState,omitempty,omitzero" bson:"deviceGuardVirtualizationBasedSecurityState,omitempty,omitzero"`
		DeviceLicensingLastErrorCode                                   float64 `json:"deviceLicensingLastErrorCode,omitempty,omitzero" bson:"deviceLicensingLastErrorCode,omitempty,omitzero"`
		DeviceLicensingLastErrorDescription                            any     `json:"deviceLicensingLastErrorDescription,omitempty,omitzero" bson:"deviceLicensingLastErrorDescription,omitempty,omitzero"`
		DeviceLicensingStatus                                          string  `json:"deviceLicensingStatus,omitempty,omitzero" bson:"deviceLicensingStatus,omitempty,omitzero"`
		EsimIdentifier                                                 any     `json:"esimIdentifier,omitempty,omitzero" bson:"esimIdentifier,omitempty,omitzero"`
		FreeStorageSpace                                               float64 `json:"freeStorageSpace,omitempty,omitzero" bson:"freeStorageSpace,omitempty,omitzero"`
		Imei                                                           string  `json:"imei,omitempty,omitzero" bson:"imei,omitempty,omitzero"`
		IpAddressV4                                                    any     `json:"ipAddressV4,omitempty,omitzero" bson:"ipAddressV4,omitempty,omitzero"`
		IsEncrypted                                                    bool    `json:"isEncrypted,omitempty,omitzero" bson:"isEncrypted,omitempty,omitzero"`
		IsSharedDevice                                                 bool    `json:"isSharedDevice,omitempty,omitzero" bson:"isSharedDevice,omitempty,omitzero"`
		IsSupervised                                                   bool    `json:"isSupervised,omitempty,omitzero" bson:"isSupervised,omitempty,omitzero"`
		Manufacturer                                                   any     `json:"manufacturer,omitempty,omitzero" bson:"manufacturer,omitempty,omitzero"`
		Meid                                                           any     `json:"meid,omitempty,omitzero" bson:"meid,omitempty,omitzero"`
		Model                                                          any     `json:"model,omitempty,omitzero" bson:"model,omitempty,omitzero"`
		OperatingSystemEdition                                         any     `json:"operatingSystemEdition,omitempty,omitzero" bson:"operatingSystemEdition,omitempty,omitzero"`
		OperatingSystemLanguage                                        any     `json:"operatingSystemLanguage,omitempty,omitzero" bson:"operatingSystemLanguage,omitempty,omitzero"`
		OperatingSystemProductType                                     float64 `json:"operatingSystemProductType,omitempty,omitzero" bson:"operatingSystemProductType,omitempty,omitzero"`
		OSBuildNumber                                                  any     `json:"osBuildNumber,omitempty,omitzero" bson:"osBuildNumber,omitempty,omitzero"`
		PhoneNumber                                                    any     `json:"phoneNumber,omitempty,omitzero" bson:"phoneNumber,omitempty,omitzero"`
		ProductName                                                    any     `json:"productName,omitempty,omitzero" bson:"productName,omitempty,omitzero"`
		ResidentUsersCount                                             any     `json:"residentUsersCount,omitempty,omitzero" bson:"residentUsersCount,omitempty,omitzero"`
		SerialNumber                                                   string  `json:"serialNumber,omitempty,omitzero" bson:"serialNumber,omitempty,omitzero"`
		SharedDeviceCachedUsers                                        []any   `json:"sharedDeviceCachedUsers,omitempty,omitzero" bson:"sharedDeviceCachedUsers,omitempty,omitzero"`
		SubnetAddress                                                  any     `json:"subnetAddress,omitempty,omitzero" bson:"subnetAddress,omitempty,omitzero"`
		SubscriberCarrier                                              any     `json:"subscriberCarrier,omitempty,omitzero" bson:"subscriberCarrier,omitempty,omitzero"`
		SystemManagementBiosVersion                                    any     `json:"systemManagementBIOSVersion,omitempty,omitzero" bson:"systemManagementBIOSVersion,omitempty,omitzero"`
		TotalStorageSpace                                              float64 `json:"totalStorageSpace,omitempty,omitzero" bson:"totalStorageSpace,omitempty,omitzero"`
		TpmManufacturer                                                any     `json:"tpmManufacturer,omitempty,omitzero" bson:"tpmManufacturer,omitempty,omitzero"`
		TpmSpecificationVersion                                        any     `json:"tpmSpecificationVersion,omitempty,omitzero" bson:"tpmSpecificationVersion,omitempty,omitzero"`
		TpmVersion                                                     any     `json:"tpmVersion,omitempty,omitzero" bson:"tpmVersion,omitempty,omitzero"`
		WifiMac                                                        any     `json:"wifiMac,omitempty,omitzero" bson:"wifiMac,omitempty,omitzero"`
		WiredIPv4Addresses                                             []any   `json:"wiredIPv4Addresses,omitempty,omitzero" bson:"wiredIPv4Addresses,omitempty,omitzero"`
	} `json:"hardwareInformation,omitempty,omitzero" bson:"hardwareInformation,omitempty,omitzero"`
	ID                                  string                      `json:"id,omitempty,omitzero" bson:"_id,omitempty,omitzero"`
	Imei                                string                      `json:"imei,omitempty,omitzero" bson:"imei,omitempty,omitzero"`
	LastSyncDateTime                    time.Time                   `json:"lastSyncDateTime,omitempty,omitzero" bson:"lastSyncDateTime,omitempty,omitzero"`
	ManagedDeviceName                   string                      `json:"managedDeviceName,omitempty,omitzero" bson:"managedDeviceName,omitempty,omitzero"`
	ManagementCertificateExpirationDate time.Time                   `json:"managementCertificateExpirationDate,omitempty,omitzero" bson:"managementCertificateExpirationDate,omitempty,omitzero"`
	Model                               string                      `json:"model,omitempty,omitzero" bson:"model,omitempty,omitzero"`
	OSVersion                           string                      `json:"osVersion,omitempty,omitzero" bson:"osVersion,omitempty,omitzero"`
	OperatingSystem                     string                      `json:"operatingSystem,omitempty,omitzero" bson:"operatingSystem,omitempty,omitzero"`
	SerialNumber                        string                      `json:"serialNumber,omitempty,omitzero" bson:"serialNumber,omitempty,omitzero"`
	TotalStorageSpaceInBytes            float64                     `json:"totalStorageSpaceInBytes"`
	UserPrincipalName                   string                      `json:"userPrincipalName,omitempty,omitzero" bson:"userPrincipalName,omitempty,omitzero"`
	UsersLoggedOn                       []ManagedDeviceUserLoggedOn `json:"usersLoggedOn,omitempty,omitzero" bson:"usersLoggedOn,omitempty,omitzero"`
	WiFiMacAddress                      string                      `json:"wiFiMacAddress,omitempty,omitzero" bson:"wiFiMacAddress,omitempty,omitzero"`
	WindowsActiveMalwareCount           float64                     `json:"windowsActiveMalwareCount"`
	WindowsRemediatedMalwareCount       float64                     `json:"windowsRemediatedMalwareCount"`
}

// type ManagedDevice struct {
// 	TenantName                                  string    `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
// 	TenantId                                    string    `json:"tenantId,omitempty,omitzero" bson:"tenantId,omitempty,omitzero"`
// 	LastDatabaseSync                            time.Time `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
// 	Odata_Context                               string    `json:"@odata.context"`
// 	AadRegistered                               bool      `json:"aadRegistered"`
// 	ActivationLockBypassCode                    any       `json:"activationLockBypassCode"`
// 	AndroidSecurityPatchLevel                   any       `json:"androidSecurityPatchLevel"`
// 	AutopilotEnrolled                           bool      `json:"autopilotEnrolled"`
// 	AzureAdDeviceID                             string    `json:"azureADDeviceId"`
// 	AzureAdRegistered                           bool      `json:"azureADRegistered"`
// 	AzureActiveDirectoryDeviceID                string    `json:"azureActiveDirectoryDeviceId"`
// 	BootstrapTokenEscrowed                      bool      `json:"bootstrapTokenEscrowed"`
// 	ChassisType                                 string    `json:"chassisType"`
// 	ChromeOSDeviceInfo                          []any     `json:"chromeOSDeviceInfo"`
// 	ComplianceGracePeriodExpirationDateTime     time.Time `json:"complianceGracePeriodExpirationDateTime"`
// 	ComplianceState                             string    `json:"complianceState"`
// 	ConfigurationManagerClientEnabledFeatures   any       `json:"configurationManagerClientEnabledFeatures"`
// 	ConfigurationManagerClientHealthState       any       `json:"configurationManagerClientHealthState"`
// 	ConfigurationManagerClientInformation       any       `json:"configurationManagerClientInformation"`
// 	DeviceActionResults                         []any     `json:"deviceActionResults"`
// 	DeviceCategoryDisplayName                   string    `json:"deviceCategoryDisplayName"`
// 	DeviceEnrollmentType                        string    `json:"deviceEnrollmentType"`
// 	DeviceFirmwareConfigurationInterfaceManaged bool      `json:"deviceFirmwareConfigurationInterfaceManaged"`
// 	DeviceHealthAttestationState                any       `json:"deviceHealthAttestationState"`
// 	DeviceIdentityAttestationDetail             struct {
// 		DeviceIdentityAttestationStatus string `json:"deviceIdentityAttestationStatus"`
// 	} `json:"deviceIdentityAttestationDetail"`
// 	DeviceName                         string    `json:"deviceName"`
// 	DeviceRegistrationState            string    `json:"deviceRegistrationState"`
// 	DeviceType                         string    `json:"deviceType"`
// 	EasActivated                       bool      `json:"easActivated"`
// 	EasActivationDateTime              time.Time `json:"easActivationDateTime"`
// 	EasDeviceID                        any       `json:"easDeviceId"`
// 	EmailAddress                       string    `json:"emailAddress"`
// 	EnrolledByUserPrincipalName        any       `json:"enrolledByUserPrincipalName"`
// 	EnrolledDateTime                   time.Time `json:"enrolledDateTime"`
// 	EnrollmentProfileName              string    `json:"enrollmentProfileName"`
// 	EthernetMacAddress                 string    `json:"ethernetMacAddress"`
// 	ExchangeAccessState                string    `json:"exchangeAccessState"`
// 	ExchangeAccessStateReason          string    `json:"exchangeAccessStateReason"`
// 	ExchangeLastSuccessfulSyncDateTime time.Time `json:"exchangeLastSuccessfulSyncDateTime"`
// 	FreeStorageSpaceInBytes            float64   `json:"freeStorageSpaceInBytes"`
// 	HardwareInformation                struct {
// 		BatteryChargeCycles                                            float64 `json:"batteryChargeCycles"`
// 		BatteryHealthPercentage                                        float64 `json:"batteryHealthPercentage"`
// 		BatteryLevelPercentage                                         any     `json:"batteryLevelPercentage"`
// 		BatterySerialNumber                                            any     `json:"batterySerialNumber"`
// 		CellularTechnology                                             any     `json:"cellularTechnology"`
// 		DeviceFullQualifiedDomainName                                  any     `json:"deviceFullQualifiedDomainName"`
// 		DeviceGuardLocalSystemAuthorityCredentialGuardState            string  `json:"deviceGuardLocalSystemAuthorityCredentialGuardState"`
// 		DeviceGuardVirtualizationBasedSecurityHardwareRequirementState string  `json:"deviceGuardVirtualizationBasedSecurityHardwareRequirementState"`
// 		DeviceGuardVirtualizationBasedSecurityState                    string  `json:"deviceGuardVirtualizationBasedSecurityState"`
// 		DeviceLicensingLastErrorCode                                   float64 `json:"deviceLicensingLastErrorCode"`
// 		DeviceLicensingLastErrorDescription                            any     `json:"deviceLicensingLastErrorDescription"`
// 		DeviceLicensingStatus                                          string  `json:"deviceLicensingStatus"`
// 		EsimIdentifier                                                 any     `json:"esimIdentifier"`
// 		FreeStorageSpace                                               float64 `json:"freeStorageSpace"`
// 		Imei                                                           string  `json:"imei"`
// 		IpAddressV4                                                    any     `json:"ipAddressV4"`
// 		IsEncrypted                                                    bool    `json:"isEncrypted"`
// 		IsSharedDevice                                                 bool    `json:"isSharedDevice"`
// 		IsSupervised                                                   bool    `json:"isSupervised"`
// 		Manufacturer                                                   any     `json:"manufacturer"`
// 		Meid                                                           any     `json:"meid"`
// 		Model                                                          any     `json:"model"`
// 		OperatingSystemEdition                                         any     `json:"operatingSystemEdition"`
// 		OperatingSystemLanguage                                        any     `json:"operatingSystemLanguage"`
// 		OperatingSystemProductType                                     float64 `json:"operatingSystemProductType"`
// 		OSBuildNumber                                                  any     `json:"osBuildNumber"`
// 		PhoneNumber                                                    any     `json:"phoneNumber"`
// 		ProductName                                                    any     `json:"productName"`
// 		ResidentUsersCount                                             any     `json:"residentUsersCount"`
// 		SerialNumber                                                   any     `json:"serialNumber"`
// 		SharedDeviceCachedUsers                                        []any   `json:"sharedDeviceCachedUsers"`
// 		SubnetAddress                                                  any     `json:"subnetAddress"`
// 		SubscriberCarrier                                              any     `json:"subscriberCarrier"`
// 		SystemManagementBiosVersion                                    any     `json:"systemManagementBIOSVersion"`
// 		TotalStorageSpace                                              float64 `json:"totalStorageSpace"`
// 		TpmManufacturer                                                any     `json:"tpmManufacturer"`
// 		TpmSpecificationVersion                                        any     `json:"tpmSpecificationVersion"`
// 		TpmVersion                                                     any     `json:"tpmVersion"`
// 		WifiMac                                                        any     `json:"wifiMac"`
// 		WiredIPv4Addresses                                             []any   `json:"wiredIPv4Addresses"`
// 	} `json:"hardwareInformation"`
// 	Iccid                                   string                      `json:"iccid"`
// 	ID                                      string                      `json:"id"`
// 	Imei                                    string                      `json:"imei"`
// 	IsEncrypted                             bool                        `json:"isEncrypted"`
// 	IsSupervised                            bool                        `json:"isSupervised"`
// 	JailBroken                              string                      `json:"jailBroken"`
// 	JoinType                                string                      `json:"joinType"`
// 	LastSyncDateTime                        time.Time                   `json:"lastSyncDateTime"`
// 	LostModeState                           string                      `json:"lostModeState"`
// 	ManagedDeviceName                       string                      `json:"managedDeviceName"`
// 	ManagedDeviceOwnerType                  string                      `json:"managedDeviceOwnerType"`
// 	ManagementAgent                         string                      `json:"managementAgent"`
// 	ManagementCertificateExpirationDate     time.Time                   `json:"managementCertificateExpirationDate"`
// 	ManagementFeatures                      string                      `json:"managementFeatures"`
// 	ManagementState                         string                      `json:"managementState"`
// 	Manufacturer                            string                      `json:"manufacturer"`
// 	Meid                                    any                         `json:"meid"`
// 	Model                                   string                      `json:"model"`
// 	Notes                                   any                         `json:"notes"`
// 	OperatingSystem                         string                      `json:"operatingSystem"`
// 	OSVersion                               string                      `json:"osVersion"`
// 	OwnerType                               string                      `json:"ownerType"`
// 	PartnerReportedThreatState              string                      `json:"partnerReportedThreatState"`
// 	PhoneNumber                             any                         `json:"phoneNumber"`
// 	PhysicalMemoryInBytes                   float64                     `json:"physicalMemoryInBytes"`
// 	PreferMdmOverGroupPolicyAppliedDateTime time.Time                   `json:"preferMdmOverGroupPolicyAppliedDateTime"`
// 	ProcessorArchitecture                   string                      `json:"processorArchitecture"`
// 	RemoteAssistanceSessionErrorDetails     string                      `json:"remoteAssistanceSessionErrorDetails"`
// 	RemoteAssistanceSessionURL              string                      `json:"remoteAssistanceSessionUrl"`
// 	RequireUserEnrollmentApproval           any                         `json:"requireUserEnrollmentApproval"`
// 	RetireAfterDateTime                     time.Time                   `json:"retireAfterDateTime"`
// 	RoleScopeTagIds                         []string                    `json:"roleScopeTagIds"`
// 	SecurityPatchLevel                      any                         `json:"securityPatchLevel"`
// 	SerialNumber                            string                      `json:"serialNumber"`
// 	SkuFamily                               string                      `json:"skuFamily"`
// 	SkuNumber                               float64                     `json:"skuNumber"`
// 	SpecificationVersion                    any                         `json:"specificationVersion"`
// 	SubscriberCarrier                       string                      `json:"subscriberCarrier"`
// 	SupplementalDeviceDetails               []any                       `json:"supplementalDeviceDetails"`
// 	TotalStorageSpaceInBytes                float64                     `json:"totalStorageSpaceInBytes"`
// 	Udid                                    string                      `json:"udid"`
// 	UserDisplayName                         string                      `json:"userDisplayName"`
// 	UserID                                  string                      `json:"userId"`
// 	UserPrincipalName                       string                      `json:"userPrincipalName"`
// 	UsersLoggedOn                           []ManagedDeviceUserLoggedOn `json:"usersLoggedOn"`
// 	WiFiMacAddress                          string                      `json:"wiFiMacAddress"`
// 	WindowsActiveMalwareCount               float64                     `json:"windowsActiveMalwareCount"`
// 	WindowsRemediatedMalwareCount           float64                     `json:"windowsRemediatedMalwareCount"`
// }

//
//

type ManagedDeviceUserLoggedOn struct {
	// EntraUser
	LastLogOnDateTime time.Time `json:"lastLogOnDateTime,omitempty,omitzero" bson:"lastLogOnDateTime,omitempty,omitzero"`
	UserID            string    `json:"userId,omitempty,omitzero" bson:"userId,omitempty,omitzero"`
	UserPrincipalName string    `json:"userPrincipalName,omitempty,omitzero" bson:"userPrincipalName,omitempty,omitzero"`
}
