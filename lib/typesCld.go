package lib

import "fmt"

type CldConfigRoot struct {
	CldConfig    CldConfig              `json:"cldConfig"`
	Azure        AzureConfig            `json:"azure"`
	ProxyConfig  map[string]ProxyConfig `json:"proxyConfig"`
	SophosConfig SophosConfig           `json:"sophos"`
	Domains      map[string]string
}

type ProxyConfig struct {
	Server    string `json:"server"`
	Port      string `json:"port"`
	Enabled   bool   `json:"enabled"`
	Overrides string `json:"overrides"`
}

type SophosConfig struct {
	Servers map[string]ServerList `json:"servers"`
	ApiKey  string                `json:"api_key"`
	ApiUser string                `json:"api_user"`
}

type ServerList map[string]string

type CldConfig struct {
	EncodeConfig bool `json:"encodeConfig"`
}

type AzureConfig struct {
	MultiTenantAuth struct {
		Tenants CldConfigTenants `json:"tenants"`
	} `json:"multiTenantAuth"`
}

func (config AzureConfig) GetDefaultTenant() CldConfigTenantAuth {
	var tenant *CldConfigTenantAuth
	for _, tConf := range config.MultiTenantAuth.Tenants {
		if tConf.Default {
			tenant = &tConf
		}
	}
	if tenant == nil {
		CheckFatalError(fmt.Errorf("No default Azure tenant configured"))
	}
	return *tenant
}

type CldConfigOptions struct {
	ConfigFile string
}

type CldConfigTenants map[string]CldConfigTenantAuth

type CldConfigTenantAuth struct {
	TenantName          string                     `json:"tenantName"`
	Default             bool                       `json:"default"`
	TenantID            string                     `json:"tenantId"`
	Reader              CldConfigClientAuthDetails `json:"reader"`
	Writer              CldConfigClientAuthDetails `json:"writer"`
	CostExportsLocation string                     `json:"costExportsLocation"`
}

type CldConfigClientAuthDetails struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}
