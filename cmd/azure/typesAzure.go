package azure

type AzureRequestOptions struct {
	SubscriptionId    string
	ResourceId        string
	ResourceGroupName string
	ResourceName      string
	TenantId          string
	TenantName        string

	ConfigFilePath string
}

type ListAllResourcesResponse struct {
	Value    []ListRspResource `json:"value"`
	NextLink string            `json:"nextLink"`
}

type ListRspResource struct {
	ID       string `json:"id"`
	Identity *struct {
		PrincipalID string `json:"principalId"`
		TenantID    string `json:"tenantId"`
		Type        string `json:"type"`
	} `json:"identity,omitempty"`
	Location  string `json:"location"`
	ManagedBy string `json:"managedBy,omitempty"`
	Name      string `json:"name"`
	Sku       *struct {
		Name string `json:"name"`
		Tier string `json:"tier"`
	} `json:"sku,omitempty"`
	Tags  map[string]string `json:"tags,omitempty"`
	Type  string            `json:"type"`
	Zones []string          `json:"zones,omitempty"`
}
