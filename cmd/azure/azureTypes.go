package azure

type AzureRequestOptions struct {
	SubscriptionId    string
	ResourceId        string
	ResourceGroupName string
	ResourceName      string
	TenantId          string
	TenantName        string

	GetWriteToken  bool
	ConfigFilePath string
}