package lib

import "time"

type MultiAuthTokenRequestOptions struct {
	// unicorn
	TenantID                     string `json:"tenantID"`
	TenantName                   string `json:"tenantName"`
	GetWriteToken                bool   `json:"getWriteToken"`
	ConfigFilePath               string `json:"configFilePath"`
	ClientID                     string `json:"clientId"`
	ClientSecret                 string `json:"clientSecret"`
	Scope                        string `json:"scope"`
	AzureContainerRepositoryName string `json:"azureContainerRepositoryName"`
}

type MultiAuthToken struct {
	TenantId   string `json:"tenantId"`
	TenantName string `json:"tenantName"`
	TokenData  TokenData
}

type Request struct {
	Url     string
	Outfile string
}

type TokenData struct {
	Token     string
	ExpiresOn string
}

type AcrAccessToken struct {
	AccessToken string
}

type TokenRequestResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	ExpiresOn    string `json:"expires_on"`
	ExtExpiresIn string `json:"ext_expires_in"`
	NotBefore    string `json:"not_before"`
	Resource     string `json:"resource"`
	TokenType    string `json:"token_type"`
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
	AuthorizationSource  string   `json:"authorizationSource"`
	DisplayName          string   `json:"displayName"`
	ID                   string   `json:"id"`
	ManagedByTenants     []string `json:"managedByTenants"`
	State                string   `json:"state"`
	SubscriptionID       string   `json:"subscriptionId"`
	SubscriptionPolicies struct {
		LocationPlacementID string `json:"locationPlacementId"`
		QuotaID             string `json:"quotaId"`
		SpendingLimit       string `json:"spendingLimit"`
	} `json:"subscriptionPolicies"`
	TenantID   string `json:"tenantId"`
	TenantName string `json:"tenantName"`
}

type SubsReqResBody struct {
	Count struct {
		Type  string  `json:"type"`
		Value float64 `json:"value"`
	} `json:"count"`
	Value []FetchedSubscription `json:"value"`
}

type AllTenantTokens []MultiAuthToken

type ListGalleryImageVersionsResponse struct {
	Value    []GalleryImageVersionResponse `json:"value"`
	NextLink string                        `json:"nextLink",omitempty`
}
type GalleryImageVersionResponse struct {
	ID         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		ProvisioningState string `json:"provisioningState"`
		PublishingProfile struct {
			ExcludeFromLatest  bool      `json:"excludeFromLatest"`
			PublishedDate      time.Time `json:"publishedDate"`
			ReplicaCount       float64   `json:"replicaCount"`
			ReplicationMode    string    `json:"replicationMode"`
			StorageAccountType string    `json:"storageAccountType"`
			TargetRegions      []struct {
				Name                 string  `json:"name"`
				RegionalReplicaCount float64 `json:"regionalReplicaCount"`
				StorageAccountType   string  `json:"storageAccountType"`
			} `json:"targetRegions"`
		} `json:"publishingProfile"`
		SafetyProfile struct {
			AllowDeletionOfReplicatedLocations bool `json:"allowDeletionOfReplicatedLocations"`
			ReportedForPolicyViolation         bool `json:"reportedForPolicyViolation"`
		} `json:"safetyProfile"`
		StorageProfile struct {
			OSDiskImage struct {
				HostCaching string   `json:"hostCaching"`
				SizeInGb    float64  `json:"sizeInGB"`
				Source      struct{} `json:"source"`
			} `json:"osDiskImage"`
			Source struct {
				VirtualMachineID string `json:"virtualMachineId"`
			} `json:"source"`
		} `json:"storageProfile"`
	} `json:"properties"`
	Tags struct {
		CostGroup string `json:"cost_group"`
		Env       string `json:"env"`
		ManagedBy string `json:"managed_by"`
	} `json:"tags"`
	Type string `json:"type"`
}

type GalleryImageVersion struct {
	ID string `json:"id"`

	Name       string `json:"name"`
	Properties struct {
		ProvisioningState string `json:"provisioningState"`
		PublishingProfile struct {
			ExcludeFromLatest bool `json:"excludeFromLatest"`
		} `json:"publishingProfile"`
	} `json:"properties"`
	SuffixAdded bool `json:"suffixAdded,omitempty"`
}

type GalleryImageVersionList struct {
	Versions []GalleryImageVersion
	Sorted   bool
}

type GalleryImageResponse struct {
	ID         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		Architecture string `json:"architecture"`
		Description  string `json:"description"`
		Disallowed   struct {
			DiskTypes []any `json:"diskTypes"`
		} `json:"disallowed"`
		HyperVGeneration string `json:"hyperVGeneration"`
		Identifier       struct {
			Offer     string `json:"offer"`
			Publisher string `json:"publisher"`
			Sku       string `json:"sku"`
		} `json:"identifier"`
		OSState           string `json:"osState"`
		OSType            string `json:"osType"`
		ProvisioningState string `json:"provisioningState"`
		Recommended       struct {
			Memory struct{} `json:"memory"`
			VCpUs  struct{} `json:"vCPUs"`
		} `json:"recommended"`
	} `json:"properties"`
	Tags struct {
		CostGroup         string `json:"cost_group"`
		Env               string `json:"env"`
		ManagedBy         string `json:"managed_by"`
		VersionsManagedBy string `json:"versions_managed_by"`
	} `json:"tags"`
	Type string `json:"type"`
}
