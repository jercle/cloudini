package lib

import "time"

type MultiAuthTokenRequestOptions struct {
	// unicorn
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

type ListSIGImageVersionsResponse struct {
	Value []struct {
		ID         string `json:"id"`
		Location   string `json:"location"`
		Name       string `json:"name"`
		Properties struct {
			ProvisioningState string `json:"provisioningState"`
			PublishingProfile struct {
				ExcludeFromLatest  bool      `json:"excludeFromLatest"`
				PublishedDate      time.Time `json:"publishedDate"`
				ReplicaCount       float64   `json:"replicaCount"`
				ReplicationMode    string    `json:"replicationMode,omitempty"`
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
				OSDiskImage *struct {
					HostCaching string   `json:"hostCaching"`
					SizeInGb    float64  `json:"sizeInGB"`
					Source      struct{} `json:"source"`
				} `json:"osDiskImage,omitempty"`
				Source struct {
					VirtualMachineID string `json:"virtualMachineId"`
				} `json:"source"`
			} `json:"storageProfile"`
		} `json:"properties"`
		Tags struct {
			CreationDate string `json:"CreationDate,omitempty"`
			CostGroup    string `json:"cost_group,omitempty"`
			Dept         string `json:"dept,omitempty"`
			Env          string `json:"env,omitempty"`
			ManagedBy    string `json:"managed_by,omitempty"`
			Task         string `json:"task,omitempty"`
		} `json:"tags"`
		Type string `json:"type"`
	} `json:"value"`
}

type SIGImageVersion struct {
	ID         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		ProvisioningState string `json:"provisioningState"`
		PublishingProfile struct {
			ExcludeFromLatest bool      `json:"excludeFromLatest"`
			PublishedDate     time.Time `json:"publishedDate"`
		} `json:"publishingProfile"`
	} `json:"properties"`
	Tags   map[string]string `json:"tags"`
	Latest bool              `json:"latest"`
}

type SIGImageVersionList []SIGImageVersion
