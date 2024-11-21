/*
Copyright © 2024 Evan Colwell ercolwell@gmail.com
*/
package lib

import "time"

type CldConfigRoot struct {
	Cloudini     CloudiniConfig         `json:"cloudini"`
	Azure        AzureConfig            `json:"azure"`
	SophosConfig SophosConfig           `json:"sophos"`
	CitrixCloud  CitrixCloud            `json:"citrixCloud"`
	ProxyConfig  map[string]ProxyConfig `json:"proxyConfig" fakesize:"2"`
	Domains      map[string]string      `json:"domains" fakesize:"2"`
}

type CitrixCloud struct {
	Environments map[string]CitrixCloudAccountConfig `json:"environments" fake:"-"`
}

type CitrixCloudAccountConfig struct {
	CustomerId   string `json:"customerId" fake:"{password:true,false,true,false,false,12}"`
	SiteId       string `json:"siteId" fake:"{uuid}"`
	ClientId     string `json:"clientId" fake:"{uuid}"`
	ClientSecret string `json:"clientSecret" fake:"{password:true,true,true,true,false,30}"`
}

type ProxyConfig struct {
	Server    string `json:"server"`
	Port      string `json:"port"`
	Enabled   bool   `json:"enabled"`
	Overrides string `json:"overrides"`
}

type SophosConfig struct {
	Environments map[string]SophosEnvironment `json:"environments" fake:"-"`
}

type SophosEnvironment struct {
	Hosts   []string `json:"hosts" fake:"-"`
	ApiUser string   `json:"api_user" fake:"{username}"`
	ApiKey  string   `json:"api_key" fake:"{password:true,true,true,true,false,30}"`
}

// type ServerList []string

type CloudiniConfig struct {
	// EncryptConfig bool `json:"encryptConfig" fake:"{bool}"`
}

type AzureConfig struct {
	MultiTenantAuth struct {
		Tenants CldConfigTenants `json:"tenants" fake:"-"`
	} `json:"multiTenantAuth"`
	TenantMap                  map[string]string   `json:"tenantMap,omitempty"`
	CustomSubIdToTenantNameMap map[string][]string `json:"customSubIdToTenantNameMap,omitempty"`
}

type CldConfigOptions struct {
	ConfigFile             string
	EncryptUnencryptedFile bool
}

type CldConfigTenants map[string]CldConfigTenantAuth

type CldConfigTenantAuth struct {
	TenantName          string                     `json:"tenantName"`
	Default             bool                       `json:"default" fake:"-"`
	TenantID            string                     `json:"tenantId" fake:"{uuid}"`
	Reader              CldConfigClientAuthDetails `json:"reader"`
	Writer              CldConfigClientAuthDetails `json:"writer"`
	CostExportsLocation string                     `json:"costExportsLocation"`
}

type CldConfigClientAuthDetails struct {
	ClientID     string `json:"clientId" fake:"{uuid}"`
	ClientSecret string `json:"clientSecret" fake:"{password:true,true,true,true,false,30}"`
}

type EncryptedTokenData struct {
	TokenData string
	TokenType string
	Expiry    time.Time
}

type TokenCache map[string]string

// type TokenCacheTypes interface {
// 	AzureMultiAuthToken | AzureTokenData | CitrixTokenData
// }
