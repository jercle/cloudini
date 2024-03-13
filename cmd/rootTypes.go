package cmd

type CldConfig struct {
	Azure struct {
		TenantAuth struct {
			Tenants []CldConfigTenantAuth `json:"tenants"`
		} `json:"multiTenantAuth"`
	} `json:"azure"`
}

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
