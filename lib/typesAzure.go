package lib

import (
	"time"
)

type AzureMultiAuthTokenRequestOptions struct {
	// unicorn
	TenantID                     string `json:"tenantID,omitempty" bson:"tenantID,omitempty"`
	TenantName                   string `json:"tenantName,omitempty" bson:"tenantName,omitempty"`
	GetWriteToken                bool   `json:"getWriteToken,omitempty" bson:"getWriteToken,omitempty"`
	ConfigFilePath               string `json:"configFilePath,omitempty" bson:"configFilePath,omitempty"`
	ClientID                     string `json:"clientId,omitempty" bson:"clientId,omitempty"`
	ClientSecret                 string `json:"clientSecret,omitempty" bson:"clientSecret,omitempty"`
	Scope                        string `json:"scope,omitempty" bson:"scope,omitempty"`
	AzureContainerRepositoryName string `json:"azureContainerRepositoryName,omitempty" bson:"azureContainerRepositoryName,omitempty"`
	NoCache                      bool   `json:"noCache,omitempty" bson:"noCache,omitempty"` // Does not use cached token
}

type AzureMultiAuthToken struct {
	TenantId   string `json:"tenantId,omitempty" bson:"tenantId,omitempty"`
	ClientId   string `json:"clientId,omitempty" bson:"clientId,omitempty"`
	TenantName string `json:"tenantName,omitempty" bson:"tenantName,omitempty"`
	TokenData  AzureTokenData
}

type Request struct {
	Url     string
	Outfile string
}

type AzureTokenData struct {
	Token     string    `json:"token,omitempty" bson:"token,omitempty"`
	ExpiresOn time.Time `json:"expiresOn,omitempty" bson:"expiresOn,omitempty"`
}

type AcrAccessToken struct {
	AccessToken string `json:"accessToken,omitempty" bson:"accessToken,omitempty"`
}

type TokenRequestResponse struct {
	AccessToken  string `json:"access_token,omitempty" bson:"access_token,omitempty"`
	ExpiresIn    string `json:"expires_in,omitempty" bson:"expires_in,omitempty"`
	ExpiresOn    string `json:"expires_on,omitempty" bson:"expires_on,omitempty"`
	ExtExpiresIn string `json:"ext_expires_in,omitempty" bson:"ext_expires_in,omitempty"`
	NotBefore    string `json:"not_before,omitempty" bson:"not_before,omitempty"`
	Resource     string `json:"resource,omitempty" bson:"resource,omitempty"`
	TokenType    string `json:"token_type,omitempty" bson:"token_type,omitempty"`
}

type AzureAuthDetails struct {
	AZURE_TENANT_ID       string
	AZURE_SUBSCRIPTION_ID string
	AZURE_CLIENT_ID       string
	AZURE_CLIENT_SECRET   string
	AZURE_RESOURCE_GROUP  string
	AZURE_RESOURCE_NAME   string
}

type AzureAuthRequirements struct {
	AZURE_TENANT_ID       bool
	AZURE_SUBSCRIPTION_ID bool
	AZURE_CLIENT_ID       bool
	AZURE_CLIENT_SECRET   bool
	AZURE_RESOURCE_GROUP  bool
	AZURE_RESOURCE_NAME   bool
}

type FetchedSubscription struct {
	AuthorizationSource  string   `json:"authorizationSource,omitempty" bson:"authorizationSource,omitempty"`
	DisplayName          string   `json:"displayName,omitempty" bson:"displayName,omitempty"`
	ID                   string   `json:"id,omitempty" bson:"id,omitempty"`
	ManagedByTenants     []string `json:"managedByTenants,omitempty" bson:"managedByTenants,omitempty"`
	State                string   `json:"state,omitempty" bson:"state,omitempty"`
	SubscriptionID       string   `json:"subscriptionId,omitempty" bson:"subscriptionId,omitempty"`
	SubscriptionPolicies struct {
		LocationPlacementID string `json:"locationPlacementId,omitempty" bson:"locationPlacementId,omitempty"`
		QuotaID             string `json:"quotaId,omitempty" bson:"quotaId,omitempty"`
		SpendingLimit       string `json:"spendingLimit,omitempty" bson:"spendingLimit,omitempty"`
	} `json:"subscriptionPolicies,omitempty" bson:"subscriptionPolicies,omitempty"`
	TenantID   string `json:"tenantId,omitempty" bson:"tenantId,omitempty"`
	TenantName string `json:"tenantName,omitempty" bson:"tenantName,omitempty"`
}

type SubsReqResBody struct {
	Count struct {
		Type  string  `json:"type,omitempty" bson:"type,omitempty"`
		Value float64 `json:"value,omitempty" bson:"value,omitempty"`
	} `json:"count,omitempty" bson:"count,omitempty"`
	Value []FetchedSubscription `json:"value,omitempty" bson:"value,omitempty"`
}

type AllTenantTokens []AzureMultiAuthToken

type ListGalleryImageVersionsResponse struct {
	Value    []GalleryImageVersionDetailed `json:"value,omitempty" bson:"value,omitempty"`
	NextLink string                        `json:"nextLink",omitempt" bson:"nextLink",omitempt"`
}
type GalleryImageVersionDetailed struct {
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Location   string `json:"location,omitempty" bson:"location,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		ProvisioningState string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
		PublishingProfile struct {
			ExcludeFromLatest  bool      `json:"excludeFromLatest,omitempty" bson:"excludeFromLatest,omitempty"`
			PublishedDate      time.Time `json:"publishedDate,omitempty" bson:"publishedDate,omitempty"`
			ReplicaCount       float64   `json:"replicaCount,omitempty" bson:"replicaCount,omitempty"`
			ReplicationMode    string    `json:"replicationMode,omitempty" bson:"replicationMode,omitempty"`
			StorageAccountType string    `json:"storageAccountType,omitempty" bson:"storageAccountType,omitempty"`
			TargetRegions      []struct {
				Name                 string  `json:"name,omitempty" bson:"name,omitempty"`
				RegionalReplicaCount float64 `json:"regionalReplicaCount,omitempty" bson:"regionalReplicaCount,omitempty"`
				StorageAccountType   string  `json:"storageAccountType,omitempty" bson:"storageAccountType,omitempty"`
			} `json:"targetRegions,omitempty" bson:"targetRegions,omitempty"`
		} `json:"publishingProfile,omitempty" bson:"publishingProfile,omitempty"`
		SafetyProfile struct {
			AllowDeletionOfReplicatedLocations bool `json:"allowDeletionOfReplicatedLocations,omitempty" bson:"allowDeletionOfReplicatedLocations,omitempty"`
			ReportedForPolicyViolation         bool `json:"reportedForPolicyViolation,omitempty" bson:"reportedForPolicyViolation,omitempty"`
		} `json:"safetyProfile,omitempty" bson:"safetyProfile,omitempty"`
		StorageProfile struct {
			OSDiskImage struct {
				HostCaching string   `json:"hostCaching,omitempty" bson:"hostCaching,omitempty"`
				SizeInGb    float64  `json:"sizeInGB,omitempty" bson:"sizeInGB,omitempty"`
				Source      struct{} `json:"source,omitempty" bson:"source,omitempty"`
			} `json:"osDiskImage,omitempty" bson:"osDiskImage,omitempty"`
			Source struct {
				VirtualMachineID string `json:"virtualMachineId,omitempty" bson:"virtualMachineId,omitempty"`
			} `json:"source,omitempty" bson:"source,omitempty"`
		} `json:"storageProfile,omitempty" bson:"storageProfile,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Tags                      map[string]string `json:"tags" bson:"tags"`
	Type                      string            `json:"type,omitempty" bson:"type,omitempty"`
	UsedByCitrix              bool              `json:"usedByCitrix,omitempty" bson:"usedByCitrix,omitempty"`
	MachineCatalogsUsingImage []string          `json:"machineCatalogsUsingImage,omitempty" bson:"machineCatalogsUsingImage,omitempty"`
	// UsedByInCitrix []string `json:"usedByInCitrix,omitempty" bson:"usedByInCitrix,omitempty"`
	LastCitrixSync time.Time          `json:"lastCitrixSync,omitempty" bson:"lastCitrixSync,omitempty"`
	LastDBSync     time.Time          `json:"lastDatabaseSync,omitempty" bson:"lastDatabaseSync,omitempty"`
	AzDoBuildData  PackerLogBuildData `json:"azDoBuildData,omitempty" bson:"azDoBuildData,omitempty"`
}

//
//

type GalleryImageVersionFlat struct {
	ID                  string    `json:"id,omitempty" bson:"id,omitempty"`
	Location            string    `json:"location,omitempty" bson:"location,omitempty"`
	ImageDefinition     string    `json:"imageDefinition,omitempty" bson:"imageDefinition,omitempty"`
	ImageDefinitionName string    `json:"imageDefinitionName,omitempty" bson:"imageDefinitionName,omitempty"`
	Name                string    `json:"name,omitempty" bson:"name,omitempty"`
	PublishedDate       time.Time `json:"publishedDate,omitempty" bson:"publishedDate,omitempty"`
	Tags                string    `json:"tags" bson:"tags"`
}

type GalleryImageVersion struct {
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		ProvisioningState string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
		PublishingProfile struct {
			ExcludeFromLatest bool `json:"excludeFromLatest,omitempty" bson:"excludeFromLatest,omitempty"`
		} `json:"publishingProfile,omitempty" bson:"publishingProfile,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	SuffixAdded  bool `json:"suffixAdded,omitempty" bson:"suffixAdded,omitempty"`
	UsedNyCitrix bool `json:"usedByCitrix,omitempty" bson:"usedByCitrix,omitempty"`
}

type GalleryImageVersionList struct {
	Versions []GalleryImageVersion
	Sorted   bool
}

type GetAllGalleryImagesResponse struct {
	Value []GalleryImage `json:"value,omitempty" bson:"value,omitempty"`
}

type GalleryImage struct {
	ID             string `json:"id,omitempty" bson:"_id,omitempty"`
	Location       string `json:"location,omitempty" bson:"location,omitempty"`
	Name           string `json:"name,omitempty" bson:"name,omitempty"`
	SubscriptionId string `json:"subscriptionId,omitempty" bson:"subscriptionId,omitempty"`
	ResourceGroup  string `json:"resourceGroup,omitempty" bson:"resourceGroup,omitempty"`
	TenantName     string `json:"tenantName,omitempty" bson:"tenantName,omitempty"`
	GalleryName    string `json:"galleryName,omitempty" bson:"galleryName,omitempty"`
	Properties     *struct {
		Architecture string `json:"architecture,omitempty" bson:"architecture,omitempty"`
		Description  string `json:"description,omitempty" bson:"description,omitempty"`
		Disallowed   *struct {
			DiskTypes *[]string `json:"diskTypes,omitempty" bson:"diskTypes,omitempty"`
		} `json:"disallowed,omitempty" bson:"disallowed,omitempty"`
		HyperVGeneration string `json:"hyperVGeneration,omitempty" bson:"hyperVGeneration,omitempty"`
		Identifier       *struct {
			Offer     string `json:"offer,omitempty" bson:"offer,omitempty"`
			Publisher string `json:"publisher,omitempty" bson:"publisher,omitempty"`
			Sku       string `json:"sku,omitempty" bson:"sku,omitempty"`
		} `json:"identifier,omitempty" bson:"identifier,omitempty"`
		OSState           string `json:"osState,omitempty" bson:"osState,omitempty"`
		OSType            string `json:"osType,omitempty" bson:"osType,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
		Recommended       *struct {
			Memory string `json:"memory,omitempty" bson:"memory,omitempty"`
			VCpUs  string `json:"vCPUs,omitempty" bson:"vCPUs,omitempty"`
		} `json:"recommended,omitempty" bson:"recommended,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Tags                      *map[string]string                     `json:"tags" bson:"tags"`
	Type                      string                                 `json:"type,omitempty" bson:"type,omitempty"`
	ImageVersions             map[string]GalleryImageVersionDetailed `json:"imageVersions,omitempty" bson:"imageVersions,omitempty"`
	UsedByCitrix              bool                                   `json:"usedByCitrix,omitempty" bson:"usedByCitrix,omitempty"`
	MachineCatalogsUsingImage []string                               `json:"machineCatalogsUsingImage,omitempty" bson:"machineCatalogsUsingImage,omitempty"`
	LastAzureSync             time.Time                              `json:"lastAzureSync,omitempty,omitzero" bson:"lastAzureSync,omitempty,omitzero"`
	LastDBSync                time.Time                              `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
}

//
//

type GalleryImageFlat struct {
	ID             string `json:"id,omitempty" bson:"_id,omitempty"`
	Location       string `json:"location,omitempty" bson:"location,omitempty"`
	Name           string `json:"name,omitempty" bson:"name,omitempty"`
	SubscriptionId string `json:"subscriptionId,omitempty" bson:"subscriptionId,omitempty"`
	ResourceGroup  string `json:"resourceGroup,omitempty" bson:"resourceGroup,omitempty"`
	TenantName     string `json:"tenantName,omitempty" bson:"tenantName,omitempty"`
	GalleryName    string `json:"galleryName,omitempty" bson:"galleryName,omitempty"`
	Description    string `json:"description,omitempty" bson:"description,omitempty"`
	Offer          string `json:"offer,omitempty" bson:"offer,omitempty"`
	Publisher      string `json:"publisher,omitempty" bson:"publisher,omitempty"`
	Sku            string `json:"sku,omitempty" bson:"sku,omitempty"`
	OSType         string `json:"osType,omitempty" bson:"osType,omitempty"`
	Tags           string `json:"tags" bson:"tags"`
}

//
//

//
//

type GetAllResourcesForAllConfiguredTenantsOptions struct {
	SubscriptionId                  string
	TenantName                      string
	AzureAuth                       CldConfigTenantAuth
	Location                        string
	OutputFilePath                  string
	SuppressSteps                   bool
	GetAllStorageAccountsInTlsCheck bool
	SelectedIPAddressQueries        *[]string
}

//
//

type GetAllM365LicenseCountsForAllConfiguredTenantsOptions struct {
	OutputFilePath string
	SuppressSteps  bool
}

type VCpuCountByTenant map[string]struct {
	VmResources               []string  `json:"vmResources,omitempty" bson:"vmResources,omitempty"`
	VmResourcesDeallocated    []string  `json:"vmResourcesDeallocated,omitempty" bson:"vmResourcesDeallocated,omitempty"`
	VmCoreCount               int       `json:"vmCoreCount,omitempty" bson:"vmCoreCount,omitempty"`
	VmCoreCountDeallocated    int       `json:"VmCoreCountDeallocated,omitempty" bson:"vmCoreCountDeallocated,omitempty"`
	VmResourcesSql            []string  `json:"vmResourcesSql,omitempty" bson:"vmResourcesSql,omitempty"`
	VmCoreCountSql            int       `json:"vmCoreCountSql,omitempty" bson:"vmCoreCountSql,omitempty"`
	VmResourcesSqlDeallocated []string  `json:"vmResourcesSqlDeallocated,omitempty" bson:"vmResourcesSqlDeallocated,omitempty"`
	VmCoreCountSqlDeallocated int       `json:"vmCoreCountSqlDeallocated,omitempty" bson:"vmCoreCountSqlDeallocated,omitempty"`
	LastDBSync                time.Time `json:"lastDatabaseSync,omitempty" bson:"lastDatabaseSync,omitempty"`
}

type VCpuCountByTenantVmResource struct {
	Id               string                   `json:"id,omitempty" bson:"id,omitempty"`
	Name             string                   `json:"name,omitempty" bson:"name,omitempty"`
	Size             string                   `json:"size,omitempty" bson:"size,omitempty"`
	VCPUs            int                      `json:"vCPUs,omitempty" bson:"vCPUs,omitempty"`
	ResourceGroup    string                   `json:"resourceGroup,omitempty" bson:"resourceGroup,omitempty"`
	SubscriptionName string                   `json:"subscriptionName,omitempty" bson:"subscriptionName,omitempty"`
	Properties       *AzureResourceProperties `json:"properties,omitempty" bson:"properties,omitempty"`
	PowerState       string                   `json:"powerState,omitempty" bson:"powerState,omitempty"`
}
