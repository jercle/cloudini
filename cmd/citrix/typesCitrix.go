package citrix

import "time"

type GetUserInfoResponse struct {
	Customers             []CustomerResponse `json:"Customers,omitempty" bson:"Customers,omitempty"`
	DisplayName           string             `json:"DisplayName,omitempty" bson:"DisplayName,omitempty"`
	ExpiryTime            time.Time          `json:"ExpiryTime,omitempty" bson:"ExpiryTime,omitempty"`
	IsCspCustomer         bool               `json:"IsCspCustomer,omitempty" bson:"IsCspCustomer,omitempty"`
	RefreshExpirationTime time.Time          `json:"RefreshExpirationTime,omitempty" bson:"RefreshExpirationTime,omitempty"`
	UserID                string             `json:"UserId,omitempty" bson:"UserId,omitempty"`
	VerifiedEmail         string             `json:"VerifiedEmail,omitempty" bson:"VerifiedEmail,omitempty"`
}

type CustomerResponse struct {
	ID    string         `json:"Id,omitempty" bson:"Id,omitempty"`
	Name  any            `json:"Name,omitempty" bson:"Name,omitempty"`
	Sites []SiteResponse `json:"Sites,omitempty" bson:"Sites,omitempty"`
}

type SiteResponse struct {
	ID   string `json:"Id,omitempty" bson:"Id,omitempty"`
	Name string `json:"Name,omitempty" bson:"Name,omitempty"`
}

type GetMachineCatalogsResponse struct {
	Items MachineCatalogs `json:"Items,omitempty" bson:"Items,omitempty"`
}

type MachineCatalogs []MachineCatalog

type MachineCatalog struct {
	AdminFolder struct {
		ID   string  `json:"id,omitempty" bson:"id,omitempty"`
		Name string  `json:"name,omitempty" bson:"name,omitempty"`
		Uid  float64 `json:"uid,omitempty" bson:"uid,omitempty"`
	} `json:"adminFolder,omitempty" bson:"adminFolder,omitempty"`
	AllocationType                  string  `json:"allocationType,omitempty" bson:"allocationType,omitempty"`
	AssignedCount                   float64 `json:"assignedCount,omitempty" bson:"assignedCount,omitempty"`
	AvailableAssignedCount          float64 `json:"availableAssignedCount,omitempty" bson:"availableAssignedCount,omitempty"`
	AvailableAssignedCountOfSuspend any     `json:"availableAssignedCountOfSuspend,omitempty" bson:"availableAssignedCountOfSuspend,omitempty"`
	AvailableCount                  float64 `json:"availableCount,omitempty" bson:"availableCount,omitempty"`
	AvailableCountOfSuspend         any     `json:"availableCountOfSuspend,omitempty" bson:"availableCountOfSuspend,omitempty"`
	AvailableUnassignedCount        float64 `json:"availableUnassignedCount,omitempty" bson:"availableUnassignedCount,omitempty"`
	CanRecreateCatalog              bool    `json:"canRecreateCatalog,omitempty" bson:"canRecreateCatalog,omitempty"`
	CanRollbackVmImage              bool    `json:"canRollbackVMImage,omitempty" bson:"canRollbackVMImage,omitempty"`
	Description                     *string `json:"description,omitempty" bson:"description,omitempty"`
	Errors                          []any   `json:"errors,omitempty" bson:"errors,omitempty"`
	FullName                        string  `json:"fullName,omitempty" bson:"fullName,omitempty"`
	HasBeenPromoted                 bool    `json:"hasBeenPromoted,omitempty" bson:"hasBeenPromoted,omitempty"`
	HasBeenPromotedFrom             string  `json:"hasBeenPromotedFrom,omitempty" bson:"hasBeenPromotedFrom,omitempty"`
	ID                              string  `json:"id,omitempty" bson:"_id,omitempty"`
	ImageUpdateStatus               string  `json:"imageUpdateStatus,omitempty" bson:"imageUpdateStatus,omitempty"`
	IsBroken                        bool    `json:"isBroken,omitempty" bson:"isBroken,omitempty"`
	IsMasterImageAssociated         bool    `json:"isMasterImageAssociated,omitempty" bson:"isMasterImageAssociated,omitempty"`
	IsPowerManaged                  bool    `json:"isPowerManaged,omitempty" bson:"isPowerManaged,omitempty"`
	IsRemotePc                      bool    `json:"isRemotePC,omitempty" bson:"isRemotePC,omitempty"`
	JobsInProgress                  any     `json:"jobsInProgress,omitempty" bson:"jobsInProgress,omitempty"`
	MachineType                     string  `json:"machineType,omitempty" bson:"machineType,omitempty"`
	Metadata                        []struct {
		Name  string `json:"name,omitempty" bson:"name,omitempty"`
		Value string `json:"value,omitempty" bson:"value,omitempty"`
	} `json:"metadata,omitempty" bson:"metadata,omitempty"`
	MinimumFunctionalLevel string `json:"minimumFunctionalLevel,omitempty" bson:"minimumFunctionalLevel,omitempty"`
	Name                   string `json:"name,omitempty" bson:"name,omitempty"`
	PersistChanges         string `json:"persistChanges,omitempty" bson:"persistChanges,omitempty"`
	ProvisioningProgress   any    `json:"provisioningProgress,omitempty" bson:"provisioningProgress,omitempty"`
	ProvisioningScheme     struct {
		AzureAdTenantID          string   `json:"azureADTenantId,omitempty" bson:"azureADTenantId,omitempty"`
		AzureAdJoinType          string   `json:"azureAdJoinType,omitempty" bson:"azureAdJoinType,omitempty"`
		AzureAdSecurityGroupName any      `json:"azureAdSecurityGroupName,omitempty" bson:"azureAdSecurityGroupName,omitempty"`
		CleanOnBoot              bool     `json:"cleanOnBoot,omitempty" bson:"cleanOnBoot,omitempty"`
		ControllerAddresses      []string `json:"controllerAddresses,omitempty" bson:"controllerAddresses,omitempty"`
		CoresPerCpuCount         float64  `json:"coresPerCpuCount,omitempty" bson:"coresPerCpuCount,omitempty"`
		CpuCount                 float64  `json:"cpuCount,omitempty" bson:"cpuCount,omitempty"`
		CurrentDiskImage         *struct {
			Date            time.Time `json:"date,omitempty" bson:"date,omitempty"`
			FunctionalLevel string    `json:"functionalLevel,omitempty" bson:"functionalLevel,omitempty"`
			Image           struct {
				FullRelativePath string `json:"fullRelativePath,omitempty" bson:"fullRelativePath,omitempty"`
				ID               any    `json:"id,omitempty" bson:"id,omitempty"`
				Name             string `json:"name,omitempty" bson:"name,omitempty"`
				ObjectTypeName   string `json:"objectTypeName,omitempty" bson:"objectTypeName,omitempty"`
				RelativePath     string `json:"relativePath,omitempty" bson:"relativePath,omitempty"`
				ResourceType     string `json:"resourceType,omitempty" bson:"resourceType,omitempty"`
				XdPath           string `json:"xDPath,omitempty" bson:"xDPath,omitempty"`
			} `json:"image,omitempty" bson:"image,omitempty"`
			ImageStatus     string `json:"imageStatus,omitempty" bson:"imageStatus,omitempty"`
			MasterImageNote any    `json:"masterImageNote,omitempty" bson:"masterImageNote,omitempty"`
		} `json:"currentDiskImage,omitempty" bson:"currentDiskImage,omitempty"`
		CurrentImageVersion *struct {
			Date         time.Time `json:"date,omitempty" bson:"date,omitempty"`
			ImageVersion struct {
				Description     any    `json:"description,omitempty" bson:"description,omitempty"`
				ID              string `json:"id,omitempty" bson:"id,omitempty"`
				ImageDefinition struct {
					ID   string `json:"id,omitempty" bson:"id,omitempty"`
					Name string `json:"name,omitempty" bson:"name,omitempty"`
					Uid  any    `json:"uid,omitempty" bson:"uid,omitempty"`
				} `json:"imageDefinition,omitempty" bson:"imageDefinition,omitempty"`
				ImageVersionSpecs any `json:"imageVersionSpecs,omitempty" bson:"imageVersionSpecs,omitempty"`
				Number            int `json:"number,omitempty" bson:"number,omitempty"`
			} `json:"imageVersion,omitempty" bson:"imageVersion,omitempty"`
		} `json:"currentImageVersion,omitempty" bson:"currentImageVersion,omitempty"`
		CustomProperties []struct {
			Name  string `json:"name,omitempty" bson:"name,omitempty"`
			Value string `json:"value,omitempty" bson:"value,omitempty"`
		} `json:"customProperties,omitempty" bson:"customProperties,omitempty"`
		CustomPropertiesInString string `json:"customPropertiesInString,omitempty" bson:"customPropertiesInString,omitempty"`
		DedicatedTenancy         bool   `json:"dedicatedTenancy,omitempty" bson:"dedicatedTenancy,omitempty"`
		DeviceManagementType     string `json:"deviceManagementType,omitempty" bson:"deviceManagementType,omitempty"`
		DiskSizeGb               int    `json:"diskSizeGB,omitempty" bson:"diskSizeGB,omitempty"`
		GpuTypeID                any    `json:"gpuTypeId,omitempty" bson:"gpuTypeId,omitempty"`
		HistoricalDiskImages     []struct {
			Date            time.Time `json:"date,omitempty" bson:"date,omitempty"`
			FunctionalLevel string    `json:"functionalLevel,omitempty" bson:"functionalLevel,omitempty"`
			Image           struct {
				FullRelativePath any    `json:"fullRelativePath,omitempty" bson:"fullRelativePath,omitempty"`
				ID               any    `json:"id,omitempty" bson:"id,omitempty"`
				Name             string `json:"name,omitempty" bson:"name,omitempty"`
				ObjectTypeName   any    `json:"objectTypeName,omitempty" bson:"objectTypeName,omitempty"`
				RelativePath     any    `json:"relativePath,omitempty" bson:"relativePath,omitempty"`
				ResourceType     any    `json:"resourceType,omitempty" bson:"resourceType,omitempty"`
				XdPath           string `json:"xDPath,omitempty" bson:"xDPath,omitempty"`
			} `json:"image,omitempty" bson:"image,omitempty"`
			ImageStatus     string  `json:"imageStatus,omitempty" bson:"imageStatus,omitempty"`
			MasterImageNote *string `json:"masterImageNote,omitempty" bson:"masterImageNote,omitempty"`
		} `json:"historicalDiskImages,omitempty" bson:"historicalDiskImages,omitempty"`
		HistoricalImageVersions []struct {
			Date         time.Time `json:"date,omitempty" bson:"date,omitempty"`
			ImageVersion struct {
				Description     any `json:"description,omitempty" bson:"description,omitempty"`
				ID              any `json:"id,omitempty" bson:"id,omitempty"`
				ImageDefinition struct {
					ID   any    `json:"id,omitempty" bson:"id,omitempty"`
					Name string `json:"name,omitempty" bson:"name,omitempty"`
					Uid  any    `json:"uid,omitempty" bson:"uid,omitempty"`
				} `json:"imageDefinition,omitempty" bson:"imageDefinition,omitempty"`
				ImageVersionSpecs any     `json:"imageVersionSpecs,omitempty" bson:"imageVersionSpecs,omitempty"`
				Number            float64 `json:"number,omitempty" bson:"number,omitempty"`
			} `json:"imageVersion,omitempty" bson:"imageVersion,omitempty"`
		} `json:"historicalImageVersions,omitempty" bson:"historicalImageVersions,omitempty"`
		ID              string `json:"id,omitempty" bson:"id,omitempty"`
		IdentityContent any    `json:"identityContent,omitempty" bson:"identityContent,omitempty"`
		IdentityPool    struct {
			ID   string `json:"id,omitempty" bson:"id,omitempty"`
			Name string `json:"name,omitempty" bson:"name,omitempty"`
			Uid  any    `json:"uid,omitempty" bson:"uid,omitempty"`
		} `json:"identityPool,omitempty" bson:"identityPool,omitempty"`
		IdentityType                string `json:"identityType,omitempty" bson:"identityType,omitempty"`
		ImageRuntimeEnvironment     any    `json:"imageRuntimeEnvironment,omitempty" bson:"imageRuntimeEnvironment,omitempty"`
		IsPreviousImageLegacyVda    bool   `json:"isPreviousImageLegacyVda,omitempty" bson:"isPreviousImageLegacyVda,omitempty"`
		MachineAccountCreationRules struct {
			Domain *struct {
				Children                             any     `json:"children,omitempty" bson:"children,omitempty"`
				Controllers                          any     `json:"controllers,omitempty" bson:"controllers,omitempty"`
				DefaultController                    any     `json:"defaultController,omitempty" bson:"defaultController,omitempty"`
				DistinguishedName                    any     `json:"distinguishedName,omitempty" bson:"distinguishedName,omitempty"`
				Forest                               string  `json:"forest,omitempty" bson:"forest,omitempty"`
				Guid                                 any     `json:"guid,omitempty" bson:"guid,omitempty"`
				Name                                 string  `json:"name,omitempty" bson:"name,omitempty"`
				NetBiosName                          any     `json:"netBiosName,omitempty" bson:"netBiosName,omitempty"`
				Parent                               any     `json:"parent,omitempty" bson:"parent,omitempty"`
				PossibleLookupFailure                bool    `json:"possibleLookupFailure,omitempty" bson:"possibleLookupFailure,omitempty"`
				PropertiesFetched                    float64 `json:"propertiesFetched,omitempty" bson:"propertiesFetched,omitempty"`
				ServiceConnectionPointConfigurations any     `json:"serviceConnectionPointConfigurations,omitempty" bson:"serviceConnectionPointConfigurations,omitempty"`
				Sid                                  any     `json:"sid,omitempty" bson:"sid,omitempty"`
				TrustedDomains                       any     `json:"trustedDomains,omitempty" bson:"trustedDomains,omitempty"`
				UpnSuffixes                          any     `json:"upnSuffixes,omitempty" bson:"upnSuffixes,omitempty"`
			} `json:"domain,omitempty" bson:"domain,omitempty"`
			NamingScheme     string  `json:"namingScheme,omitempty" bson:"namingScheme,omitempty"`
			NamingSchemeType string  `json:"namingSchemeType,omitempty" bson:"namingSchemeType,omitempty"`
			NextValue        string  `json:"nextValue,omitempty" bson:"nextValue,omitempty"`
			Ou               *string `json:"oU,omitempty" bson:"oU,omitempty"`
		} `json:"machineAccountCreationRules,omitempty" bson:"machineAccountCreationRules,omitempty"`
		MachineCount   float64 `json:"machineCount,omitempty" bson:"machineCount,omitempty"`
		MachineProfile *struct {
			FullRelativePath any    `json:"fullRelativePath,omitempty" bson:"fullRelativePath,omitempty"`
			ID               any    `json:"id,omitempty" bson:"id,omitempty"`
			Name             string `json:"name,omitempty" bson:"name,omitempty"`
			ObjectTypeName   string `json:"objectTypeName,omitempty" bson:"objectTypeName,omitempty"`
			RelativePath     any    `json:"relativePath,omitempty" bson:"relativePath,omitempty"`
			ResourceType     string `json:"resourceType,omitempty" bson:"resourceType,omitempty"`
			XdPath           string `json:"xDPath,omitempty" bson:"xDPath,omitempty"`
		} `json:"machineProfile,omitempty" bson:"machineProfile,omitempty"`
		MasterImage struct {
			FullRelativePath any    `json:"fullRelativePath,omitempty" bson:"fullRelativePath,omitempty"`
			ID               any    `json:"id,omitempty" bson:"id,omitempty"`
			Name             string `json:"name,omitempty" bson:"name,omitempty"`
			ObjectTypeName   string `json:"objectTypeName,omitempty" bson:"objectTypeName,omitempty"`
			RelativePath     any    `json:"relativePath,omitempty" bson:"relativePath,omitempty"`
			ResourceType     string `json:"resourceType,omitempty" bson:"resourceType,omitempty"`
			XdPath           string `json:"xDPath,omitempty" bson:"xDPath,omitempty"`
		} `json:"masterImage,omitempty" bson:"masterImage,omitempty"`
		MemoryMb float64 `json:"memoryMB,omitempty" bson:"memoryMB,omitempty"`
		Metadata []struct {
			Name  string `json:"name,omitempty" bson:"name,omitempty"`
			Value string `json:"value,omitempty" bson:"value,omitempty"`
		} `json:"metadata,omitempty" bson:"metadata,omitempty"`
		Name        string `json:"name,omitempty" bson:"name,omitempty"`
		NetworkMaps []struct {
			DeviceID      string `json:"deviceId,omitempty" bson:"deviceId,omitempty"`
			DeviceName    any    `json:"deviceName,omitempty" bson:"deviceName,omitempty"`
			IsCardEnabled bool   `json:"isCardEnabled,omitempty" bson:"isCardEnabled,omitempty"`
			Network       struct {
				FullRelativePath string `json:"fullRelativePath,omitempty" bson:"fullRelativePath,omitempty"`
				ID               any    `json:"id,omitempty" bson:"id,omitempty"`
				Name             any    `json:"name,omitempty" bson:"name,omitempty"`
				ObjectTypeName   string `json:"objectTypeName,omitempty" bson:"objectTypeName,omitempty"`
				RelativePath     string `json:"relativePath,omitempty" bson:"relativePath,omitempty"`
				ResourceType     string `json:"resourceType,omitempty" bson:"resourceType,omitempty"`
				XdPath           string `json:"xDPath,omitempty" bson:"xDPath,omitempty"`
			} `json:"network,omitempty" bson:"network,omitempty"`
		} `json:"networkMaps,omitempty" bson:"networkMaps,omitempty"`
		NumAvailableMachineAccounts float64 `json:"numAvailableMachineAccounts,omitempty" bson:"numAvailableMachineAccounts,omitempty"`
		NutanixContainer            any     `json:"nutanixContainer,omitempty" bson:"nutanixContainer,omitempty"`
		PvsSite                     any     `json:"pVSSite,omitempty" bson:"pVSSite,omitempty"`
		PvsvDisk                    any     `json:"pVSVDisk,omitempty" bson:"pVSVDisk,omitempty"`
		ProvisioningSchemeType      string  `json:"provisioningSchemeType,omitempty" bson:"provisioningSchemeType,omitempty"`
		ResetAdministratorPasswords bool    `json:"resetAdministratorPasswords,omitempty" bson:"resetAdministratorPasswords,omitempty"`
		ResourcePool                struct {
			FullRelativePath string `json:"fullRelativePath,omitempty" bson:"fullRelativePath,omitempty"`
			Hypervisor       struct {
				ConnectionType    string  `json:"connectionType,omitempty" bson:"connectionType,omitempty"`
				ID                string  `json:"id,omitempty" bson:"id,omitempty"`
				Name              string  `json:"name,omitempty" bson:"name,omitempty"`
				PluginFactoryName string  `json:"pluginFactoryName,omitempty" bson:"pluginFactoryName,omitempty"`
				Uid               float64 `json:"uid,omitempty" bson:"uid,omitempty"`
			} `json:"hypervisor,omitempty" bson:"hypervisor,omitempty"`
			ID     string `json:"id,omitempty" bson:"id,omitempty"`
			Name   string `json:"name,omitempty" bson:"name,omitempty"`
			XdPath string `json:"xDPath,omitempty" bson:"xDPath,omitempty"`
		} `json:"resourcePool,omitempty" bson:"resourcePool,omitempty"`
		SecurityGroups               []any  `json:"securityGroups,omitempty" bson:"securityGroups,omitempty"`
		ServiceAccountUid            []any  `json:"serviceAccountUid,omitempty" bson:"serviceAccountUid,omitempty"`
		ServiceOffering              string `json:"serviceOffering,omitempty" bson:"serviceOffering,omitempty"`
		TenancyType                  string `json:"tenancyType,omitempty" bson:"tenancyType,omitempty"`
		UseFullDiskCloneProvisioning bool   `json:"useFullDiskCloneProvisioning,omitempty" bson:"useFullDiskCloneProvisioning,omitempty"`
		UseWriteBackCache            bool   `json:"useWriteBackCache,omitempty" bson:"useWriteBackCache,omitempty"`
		VmMetadata                   *struct {
			AcceleratedNetwork                bool   `json:"acceleratedNetwork,omitempty" bson:"acceleratedNetwork,omitempty"`
			BootDiagnostics                   any    `json:"bootDiagnostics,omitempty" bson:"bootDiagnostics,omitempty"`
			ConfidentialVmDiskEncryptionSetID string `json:"confidentialVmDiskEncryptionSetId,omitempty" bson:"confidentialVmDiskEncryptionSetId,omitempty"`
			DiskSecurityType                  string `json:"diskSecurityType,omitempty" bson:"diskSecurityType,omitempty"`
			EnableSecureBoot                  any    `json:"enableSecureBoot,omitempty" bson:"enableSecureBoot,omitempty"`
			EnableVtpm                        any    `json:"enableVTPM,omitempty" bson:"enableVTPM,omitempty"`
			EncryptionAtHost                  any    `json:"encryptionAtHost,omitempty" bson:"encryptionAtHost,omitempty"`
			EncryptionKeyID                   any    `json:"encryptionKeyId,omitempty" bson:"encryptionKeyId,omitempty"`
			Labels                            any    `json:"labels,omitempty" bson:"labels,omitempty"`
			MachineSize                       string `json:"machineSize,omitempty" bson:"machineSize,omitempty"`
			OSDiskCaching                     string `json:"osDiskCaching,omitempty" bson:"osDiskCaching,omitempty"`
			SecurityType                      string `json:"securityType,omitempty" bson:"securityType,omitempty"`
			StorageType                       any    `json:"storageType,omitempty" bson:"storageType,omitempty"`
			SupportsHibernation               any    `json:"supportsHibernation,omitempty" bson:"supportsHibernation,omitempty"`
			Tags                              string `json:"tags" bson:"tags"`
			ZoneName                          any    `json:"zoneName,omitempty" bson:"zoneName,omitempty"`
		} `json:"vMMetadata,omitempty" bson:"vMMetadata,omitempty"`
		Warning                    any     `json:"warning,omitempty" bson:"warning,omitempty"`
		Warnings                   []any   `json:"warnings,omitempty" bson:"warnings,omitempty"`
		WindowsActivationType      string  `json:"windowsActivationType,omitempty" bson:"windowsActivationType,omitempty"`
		WorkgroupMachines          bool    `json:"workgroupMachines,omitempty" bson:"workgroupMachines,omitempty"`
		WriteBackCacheDiskIndex    float64 `json:"writeBackCacheDiskIndex,omitempty" bson:"writeBackCacheDiskIndex,omitempty"`
		WriteBackCacheDiskSizeGb   float64 `json:"writeBackCacheDiskSizeGB,omitempty" bson:"writeBackCacheDiskSizeGB,omitempty"`
		WriteBackCacheDriveLetter  string  `json:"writeBackCacheDriveLetter,omitempty" bson:"writeBackCacheDriveLetter,omitempty"`
		WriteBackCacheMemorySizeMb float64 `json:"writeBackCacheMemorySizeMB,omitempty" bson:"writeBackCacheMemorySizeMB,omitempty"`
	} `json:"provisioningScheme,omitempty" bson:"provisioningScheme,omitempty"`
	ProvisioningType         string `json:"provisioningType,omitempty" bson:"provisioningType,omitempty"`
	PvsAddress               any    `json:"pvsAddress,omitempty" bson:"pvsAddress,omitempty"`
	PvsDomain                any    `json:"pvsDomain,omitempty" bson:"pvsDomain,omitempty"`
	RemotePcEnrollmentScopes any    `json:"remotePCEnrollmentScopes,omitempty" bson:"remotePCEnrollmentScopes,omitempty"`
	Scopes                   []struct {
		Description   any    `json:"description,omitempty" bson:"description,omitempty"`
		ID            string `json:"id,omitempty" bson:"id,omitempty"`
		IsAllScope    bool   `json:"isAllScope,omitempty" bson:"isAllScope,omitempty"`
		IsBuiltIn     bool   `json:"isBuiltIn,omitempty" bson:"isBuiltIn,omitempty"`
		IsTenantScope bool   `json:"isTenantScope,omitempty" bson:"isTenantScope,omitempty"`
		Name          string `json:"name,omitempty" bson:"name,omitempty"`
		TenantID      any    `json:"tenantId,omitempty" bson:"tenantId,omitempty"`
		TenantName    any    `json:"tenantName,omitempty" bson:"tenantName,omitempty"`
		Uid           any    `json:"uid,omitempty" bson:"uid,omitempty"`
	} `json:"scopes,omitempty" bson:"scopes,omitempty"`
	SessionSupport  string  `json:"sessionSupport,omitempty" bson:"sessionSupport,omitempty"`
	SharingKind     string  `json:"sharingKind,omitempty" bson:"sharingKind,omitempty"`
	Tenants         any     `json:"tenants,omitempty" bson:"tenants,omitempty"`
	TotalCount      float64 `json:"totalCount,omitempty" bson:"totalCount,omitempty"`
	Uid             float64 `json:"uid,omitempty" bson:"uid,omitempty"`
	UnassignedCount float64 `json:"unassignedCount,omitempty" bson:"unassignedCount,omitempty"`
	UpgradeInfo     *struct {
		UpgradeFailedMachinesCount  float64 `json:"upgradeFailedMachinesCount,omitempty" bson:"upgradeFailedMachinesCount,omitempty"`
		UpgradeOngoingMachinesCount float64 `json:"upgradeOngoingMachinesCount,omitempty" bson:"upgradeOngoingMachinesCount,omitempty"`
		UpgradeScheduleStatus       string  `json:"upgradeScheduleStatus,omitempty" bson:"upgradeScheduleStatus,omitempty"`
		UpgradeState                string  `json:"upgradeState,omitempty" bson:"upgradeState,omitempty"`
		UpgradeType                 string  `json:"upgradeType,omitempty" bson:"upgradeType,omitempty"`
	} `json:"upgradeInfo,omitempty" bson:"upgradeInfo,omitempty"`
	UsedCount float64 `json:"usedCount,omitempty" bson:"usedCount,omitempty"`
	Warnings  []any   `json:"warnings,omitempty" bson:"warnings,omitempty"`
	Zone      struct {
		ID   string `json:"id,omitempty" bson:"id,omitempty"`
		Name string `json:"name,omitempty" bson:"name,omitempty"`
		Uid  any    `json:"uid,omitempty" bson:"uid,omitempty"`
	} `json:"zone,omitempty" bson:"zone,omitempty"`
	TenantName     string    `json:"tenantName,omitempty" bson:"tenantName,omitempty"`
	LastCitrixSync time.Time `json:"lastCitrixSync,omitempty" bson:"lastCitrixSync,omitempty"`
	LastDBSync     time.Time `json:"lastDatabaseSync,omitempty" bson:"lastDatabaseSync,omitempty"`
}

type GetMachineCatalogDeliveryGroupAssociationsResponse struct {
	Items []MchnCatDelGrpAssociation `json:"Items,omitempty" bson:"Items,omitempty"`
}

type MchnCatDelGrpAssociation struct {
	Associated bool    `json:"Associated,omitempty" bson:"Associated,omitempty"`
	ID         string  `json:"Id,omitempty" bson:"Id,omitempty"`
	Name       string  `json:"Name,omitempty" bson:"Name,omitempty"`
	Priority   any     `json:"Priority,omitempty" bson:"Priority,omitempty"`
	Uid        float64 `json:"Uid,omitempty" bson:"Uid,omitempty"`
}

type MachineCatalogCurrentImage struct {
	ImageDefinitionName  string    `json:"imageDefinition,omitempty" bson:"imageDefinition,omitempty"`
	Version              string    `json:"version,omitempty" bson:"version,omitempty"`
	ImageGallery         string    `json:"imageGallery,omitempty" bson:"imageGallery,omitempty"`
	ResourceGroup        string    `json:"resourceGroup,omitempty" bson:"resourceGroup,omitempty"`
	IsPreparedImage      bool      `json:"isPreparedImage,omitempty" bson:"isPreparedImage,omitempty"`
	PreparedImageName    string    `json:"preparedImageName,omitempty" bson:"preparedImageName,omitempty"`
	PreparedImageVersion string    `json:"preparedImageVersion,omitempty" bson:"preparedImageVersion,omitempty"`
	LastCitrixSync       time.Time `json:"lastCitrixSync,omitempty" bson:"lastCitrixSync,omitempty"`
	LastDBSync           string    `json:"lastDatabaseSync,omitempty" bson:"lastDatabaseSync,omitempty"`
}
