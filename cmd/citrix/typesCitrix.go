/*
Copyright © 2024 Evan Colwell ercolwell@gmail.com
*/
package citrix

import "time"

type GetUserInfoResponse struct {
	Customers             []CustomerResponse `json:"Customers"`
	DisplayName           string             `json:"DisplayName"`
	ExpiryTime            time.Time          `json:"ExpiryTime"`
	IsCspCustomer         bool               `json:"IsCspCustomer"`
	RefreshExpirationTime time.Time          `json:"RefreshExpirationTime"`
	UserID                string             `json:"UserId"`
	VerifiedEmail         string             `json:"VerifiedEmail"`
}

type CustomerResponse struct {
	ID    string         `json:"Id"`
	Name  any            `json:"Name"`
	Sites []SiteResponse `json:"Sites"`
}

type SiteResponse struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
}

type GetMachineCatalogsResponse struct {
	Items MachineCatalogs `json:"Items"`
}

type MachineCatalogs []MachineCatalog

type MachineCatalog struct {
	AdminFolder struct {
		ID   string  `json:"Id"`
		Name string  `json:"Name"`
		Uid  float64 `json:"Uid"`
	} `json:"AdminFolder"`
	AllocationType                  string  `json:"AllocationType"`
	AssignedCount                   float64 `json:"AssignedCount"`
	AvailableAssignedCount          float64 `json:"AvailableAssignedCount"`
	AvailableAssignedCountOfSuspend any     `json:"AvailableAssignedCountOfSuspend"`
	AvailableCount                  float64 `json:"AvailableCount"`
	AvailableCountOfSuspend         any     `json:"AvailableCountOfSuspend"`
	AvailableUnassignedCount        float64 `json:"AvailableUnassignedCount"`
	CanRecreateCatalog              bool    `json:"CanRecreateCatalog"`
	CanRollbackVmImage              bool    `json:"CanRollbackVMImage"`
	Description                     *string `json:"Description"`
	Errors                          []any   `json:"Errors"`
	FullName                        string  `json:"FullName"`
	HasBeenPromoted                 bool    `json:"HasBeenPromoted"`
	HasBeenPromotedFrom             string  `json:"HasBeenPromotedFrom"`
	ID                              string  `json:"Id"`
	ImageUpdateStatus               string  `json:"ImageUpdateStatus"`
	IsBroken                        bool    `json:"IsBroken"`
	IsMasterImageAssociated         bool    `json:"IsMasterImageAssociated"`
	IsPowerManaged                  bool    `json:"IsPowerManaged"`
	IsRemotePc                      bool    `json:"IsRemotePC"`
	JobsInProgress                  any     `json:"JobsInProgress"`
	MachineType                     string  `json:"MachineType"`
	Metadata                        []struct {
		Name  string `json:"Name"`
		Value string `json:"Value"`
	} `json:"Metadata"`
	MinimumFunctionalLevel string `json:"MinimumFunctionalLevel"`
	Name                   string `json:"Name"`
	PersistChanges         string `json:"PersistChanges"`
	ProvisioningProgress   any    `json:"ProvisioningProgress"`
	ProvisioningScheme     struct {
		AzureAdTenantID          string   `json:"AzureADTenantId"`
		AzureAdJoinType          string   `json:"AzureAdJoinType"`
		AzureAdSecurityGroupName any      `json:"AzureAdSecurityGroupName"`
		CleanOnBoot              bool     `json:"CleanOnBoot"`
		ControllerAddresses      []string `json:"ControllerAddresses"`
		CoresPerCpuCount         float64  `json:"CoresPerCpuCount"`
		CpuCount                 float64  `json:"CpuCount"`
		CurrentDiskImage         *struct {
			Date            time.Time `json:"Date"`
			FunctionalLevel string    `json:"FunctionalLevel"`
			Image           struct {
				FullRelativePath string `json:"FullRelativePath"`
				ID               any    `json:"Id"`
				Name             string `json:"Name"`
				ObjectTypeName   string `json:"ObjectTypeName"`
				RelativePath     string `json:"RelativePath"`
				ResourceType     string `json:"ResourceType"`
				XdPath           string `json:"XDPath"`
			} `json:"Image"`
			ImageStatus     string `json:"ImageStatus"`
			MasterImageNote any    `json:"MasterImageNote"`
		} `json:"CurrentDiskImage"`
		CurrentImageVersion *struct {
			Date         time.Time `json:"Date"`
			ImageVersion struct {
				Description     any    `json:"Description"`
				ID              string `json:"Id"`
				ImageDefinition struct {
					ID   string `json:"Id"`
					Name string `json:"Name"`
					Uid  any    `json:"Uid"`
				} `json:"ImageDefinition"`
				ImageVersionSpecs any `json:"ImageVersionSpecs"`
				Number            int `json:"Number"`
			} `json:"ImageVersion"`
		} `json:"CurrentImageVersion"`
		CustomProperties []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"CustomProperties"`
		CustomPropertiesInString string `json:"CustomPropertiesInString"`
		DedicatedTenancy         bool   `json:"DedicatedTenancy"`
		DeviceManagementType     string `json:"DeviceManagementType"`
		DiskSizeGb               int    `json:"DiskSizeGB"`
		GpuTypeID                any    `json:"GpuTypeId"`
		HistoricalDiskImages     []struct {
			Date            time.Time `json:"Date"`
			FunctionalLevel string    `json:"FunctionalLevel"`
			Image           struct {
				FullRelativePath any    `json:"FullRelativePath"`
				ID               any    `json:"Id"`
				Name             string `json:"Name"`
				ObjectTypeName   any    `json:"ObjectTypeName"`
				RelativePath     any    `json:"RelativePath"`
				ResourceType     any    `json:"ResourceType"`
				XdPath           string `json:"XDPath"`
			} `json:"Image"`
			ImageStatus     string  `json:"ImageStatus"`
			MasterImageNote *string `json:"MasterImageNote"`
		} `json:"HistoricalDiskImages"`
		HistoricalImageVersions []struct {
			Date         time.Time `json:"Date"`
			ImageVersion struct {
				Description     any `json:"Description"`
				ID              any `json:"Id"`
				ImageDefinition struct {
					ID   any    `json:"Id"`
					Name string `json:"Name"`
					Uid  any    `json:"Uid"`
				} `json:"ImageDefinition"`
				ImageVersionSpecs any     `json:"ImageVersionSpecs"`
				Number            float64 `json:"Number"`
			} `json:"ImageVersion"`
		} `json:"HistoricalImageVersions"`
		ID              string `json:"Id"`
		IdentityContent any    `json:"IdentityContent"`
		IdentityPool    struct {
			ID   string `json:"Id"`
			Name string `json:"Name"`
			Uid  any    `json:"Uid"`
		} `json:"IdentityPool"`
		IdentityType                string `json:"IdentityType"`
		ImageRuntimeEnvironment     any    `json:"ImageRuntimeEnvironment"`
		IsPreviousImageLegacyVda    bool   `json:"IsPreviousImageLegacyVda"`
		MachineAccountCreationRules struct {
			Domain *struct {
				Children                             any     `json:"Children"`
				Controllers                          any     `json:"Controllers"`
				DefaultController                    any     `json:"DefaultController"`
				DistinguishedName                    any     `json:"DistinguishedName"`
				Forest                               string  `json:"Forest"`
				Guid                                 any     `json:"Guid"`
				Name                                 string  `json:"Name"`
				NetBiosName                          any     `json:"NetBiosName"`
				Parent                               any     `json:"Parent"`
				PossibleLookupFailure                bool    `json:"PossibleLookupFailure"`
				PropertiesFetched                    float64 `json:"PropertiesFetched"`
				ServiceConnectionPointConfigurations any     `json:"ServiceConnectionPointConfigurations"`
				Sid                                  any     `json:"Sid"`
				TrustedDomains                       any     `json:"TrustedDomains"`
				UpnSuffixes                          any     `json:"UpnSuffixes"`
			} `json:"Domain"`
			NamingScheme     string  `json:"NamingScheme"`
			NamingSchemeType string  `json:"NamingSchemeType"`
			NextValue        string  `json:"NextValue"`
			Ou               *string `json:"OU"`
		} `json:"MachineAccountCreationRules"`
		MachineCount   float64 `json:"MachineCount"`
		MachineProfile *struct {
			FullRelativePath any    `json:"FullRelativePath"`
			ID               any    `json:"Id"`
			Name             string `json:"Name"`
			ObjectTypeName   string `json:"ObjectTypeName"`
			RelativePath     any    `json:"RelativePath"`
			ResourceType     string `json:"ResourceType"`
			XdPath           string `json:"XDPath"`
		} `json:"MachineProfile"`
		MasterImage struct {
			FullRelativePath any    `json:"FullRelativePath"`
			ID               any    `json:"Id"`
			Name             string `json:"Name"`
			ObjectTypeName   string `json:"ObjectTypeName"`
			RelativePath     any    `json:"RelativePath"`
			ResourceType     string `json:"ResourceType"`
			XdPath           string `json:"XDPath"`
		} `json:"MasterImage"`
		MemoryMb float64 `json:"MemoryMB"`
		Metadata []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"Metadata"`
		Name        string `json:"Name"`
		NetworkMaps []struct {
			DeviceID      string `json:"DeviceId"`
			DeviceName    any    `json:"DeviceName"`
			IsCardEnabled bool   `json:"IsCardEnabled"`
			Network       struct {
				FullRelativePath string `json:"FullRelativePath"`
				ID               any    `json:"Id"`
				Name             any    `json:"Name"`
				ObjectTypeName   string `json:"ObjectTypeName"`
				RelativePath     string `json:"RelativePath"`
				ResourceType     string `json:"ResourceType"`
				XdPath           string `json:"XDPath"`
			} `json:"Network"`
		} `json:"NetworkMaps"`
		NumAvailableMachineAccounts float64 `json:"NumAvailableMachineAccounts"`
		NutanixContainer            any     `json:"NutanixContainer"`
		PvsSite                     any     `json:"PVSSite"`
		PvsvDisk                    any     `json:"PVSVDisk"`
		ProvisioningSchemeType      string  `json:"ProvisioningSchemeType"`
		ResetAdministratorPasswords bool    `json:"ResetAdministratorPasswords"`
		ResourcePool                struct {
			FullRelativePath string `json:"FullRelativePath"`
			Hypervisor       struct {
				ConnectionType    string  `json:"ConnectionType"`
				ID                string  `json:"Id"`
				Name              string  `json:"Name"`
				PluginFactoryName string  `json:"PluginFactoryName"`
				Uid               float64 `json:"Uid"`
			} `json:"Hypervisor"`
			ID     string `json:"Id"`
			Name   string `json:"Name"`
			XdPath string `json:"XDPath"`
		} `json:"ResourcePool"`
		SecurityGroups               []any  `json:"SecurityGroups"`
		ServiceAccountUid            []any  `json:"ServiceAccountUid"`
		ServiceOffering              string `json:"ServiceOffering"`
		TenancyType                  string `json:"TenancyType"`
		UseFullDiskCloneProvisioning bool   `json:"UseFullDiskCloneProvisioning"`
		UseWriteBackCache            bool   `json:"UseWriteBackCache"`
		VmMetadata                   *struct {
			AcceleratedNetwork                bool   `json:"AcceleratedNetwork"`
			BootDiagnostics                   any    `json:"BootDiagnostics"`
			ConfidentialVmDiskEncryptionSetID string `json:"ConfidentialVmDiskEncryptionSetId"`
			DiskSecurityType                  string `json:"DiskSecurityType"`
			EnableSecureBoot                  any    `json:"EnableSecureBoot"`
			EnableVtpm                        any    `json:"EnableVTPM"`
			EncryptionAtHost                  any    `json:"EncryptionAtHost"`
			EncryptionKeyID                   any    `json:"EncryptionKeyId"`
			Labels                            any    `json:"Labels"`
			MachineSize                       string `json:"MachineSize"`
			OSDiskCaching                     string `json:"OsDiskCaching"`
			SecurityType                      string `json:"SecurityType"`
			StorageType                       any    `json:"StorageType"`
			SupportsHibernation               any    `json:"SupportsHibernation"`
			Tags                              string `json:"Tags"`
			ZoneName                          any    `json:"ZoneName"`
		} `json:"VMMetadata"`
		Warning                    any     `json:"Warning"`
		Warnings                   []any   `json:"Warnings"`
		WindowsActivationType      string  `json:"WindowsActivationType"`
		WorkgroupMachines          bool    `json:"WorkgroupMachines"`
		WriteBackCacheDiskIndex    float64 `json:"WriteBackCacheDiskIndex"`
		WriteBackCacheDiskSizeGb   float64 `json:"WriteBackCacheDiskSizeGB"`
		WriteBackCacheDriveLetter  string  `json:"WriteBackCacheDriveLetter"`
		WriteBackCacheMemorySizeMb float64 `json:"WriteBackCacheMemorySizeMB"`
	} `json:"ProvisioningScheme"`
	ProvisioningType         string `json:"ProvisioningType"`
	PvsAddress               any    `json:"PvsAddress"`
	PvsDomain                any    `json:"PvsDomain"`
	RemotePcEnrollmentScopes any    `json:"RemotePCEnrollmentScopes"`
	Scopes                   []struct {
		Description   any    `json:"Description"`
		ID            string `json:"Id"`
		IsAllScope    bool   `json:"IsAllScope"`
		IsBuiltIn     bool   `json:"IsBuiltIn"`
		IsTenantScope bool   `json:"IsTenantScope"`
		Name          string `json:"Name"`
		TenantID      any    `json:"TenantId"`
		TenantName    any    `json:"TenantName"`
		Uid           any    `json:"Uid"`
	} `json:"Scopes"`
	SessionSupport  string  `json:"SessionSupport"`
	SharingKind     string  `json:"SharingKind"`
	Tenants         any     `json:"Tenants"`
	TotalCount      float64 `json:"TotalCount"`
	Uid             float64 `json:"Uid"`
	UnassignedCount float64 `json:"UnassignedCount"`
	UpgradeInfo     *struct {
		UpgradeFailedMachinesCount  float64 `json:"UpgradeFailedMachinesCount"`
		UpgradeOngoingMachinesCount float64 `json:"UpgradeOngoingMachinesCount"`
		UpgradeScheduleStatus       string  `json:"UpgradeScheduleStatus"`
		UpgradeState                string  `json:"UpgradeState"`
		UpgradeType                 string  `json:"UpgradeType"`
	} `json:"UpgradeInfo"`
	UsedCount float64 `json:"UsedCount"`
	Warnings  []any   `json:"Warnings"`
	Zone      struct {
		ID   string `json:"Id"`
		Name string `json:"Name"`
		Uid  any    `json:"Uid"`
	} `json:"Zone"`
}

type GetMachineCatalogDeliveryGroupAssociationsResponse struct {
	Items []MchnCatDelGrpAssociation `json:"Items"`
}

type MchnCatDelGrpAssociation struct {
	Associated bool    `json:"Associated"`
	ID         string  `json:"Id"`
	Name       string  `json:"Name"`
	Priority   any     `json:"Priority"`
	Uid        float64 `json:"Uid"`
}

type MachineCatalogCurrentImage struct {
	ImageDefinitionName  string  `json:"imageDefinition"`
	Version              string  `json:"version"`
	ImageGallery         string  `json:"imageGallery"`
	ResourceGroup        string  `json:"resourceGroup"`
	IsPreparedImage      bool    `json:"isPreparedImage,omitempty"`
	PreparedImageName    *string `json:"preparedImageName,omitempty"`
	PreparedImageVersion *string `json:"preparedImageVersion,omitempty"`
}
