package lib

type CldConfigRoot struct {
	CldConfig CldConfig   `json:"cldConfig"`
	Azure     AzureConfig `json:"azure"`
}

type CldConfig struct {
	EncodeConfig bool `json:"encodeConfig"`
}

type AzureConfig struct {
	MultiTenantAuth struct {
		Tenants CldConfigTenants `json:"tenants"`
	} `json:"multiTenantAuth"`
}

type CldConfigOptions struct {
	ConfigFilePath string
}

type CldConfigTenants map[string]CldConfigTenantAuth

type CldConfigTenantAuth struct {
	TenantName string                     `json:"tenantName"`
	TenantID   string                     `json:"tenantId"`
	Reader     CldConfigClientAuthDetails `json:"reader"`
	Writer     CldConfigClientAuthDetails `json:"writer"`
}

type CldConfigClientAuthDetails struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}
