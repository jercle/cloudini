package lib

import "time"

type CldConfigRoot struct {
	// Cloudini     *CloudiniConfig         `json:"cloudini,omitempty"`
	ActiveDirectory       *ActiveDirectoryConfig       `json:"activeDirectory,omitempty"`
	AWS                   *AWSConfig                   `json:"aws,omitempty"`
	Azure                 *AzureConfig                 `json:"azure,omitempty"`
	AzureDevOps           *AzureDevOpsConfig           `json:"azureDevOps,omitempty"`
	CertificateManagement *CertificateManagementConfig `json:"certificateManagement,omitempty"`
	CitrixCloud           *CitrixCloud                 `json:"citrixCloud,omitempty"`
	Domains               *map[string]string           `json:"domains,omitempty" fakesize:"2"`
	Forgerock             *ForgerockConfig             `json:"forgerock,omitempty"`
	MongoDBConfig         *MongoDBConfig               `json:"mongoDbConfig,omitempty"`
	ProxyConfig           *map[string]ProxyConfig      `json:"proxyConfig,omitempty" fakesize:"2"`
	SophosConfig          *SophosConfig                `json:"sophos,omitempty"`
}

type ForgerockConfig struct {
	Domains FRDomains `json:"domains"`
}
type FRDomains map[string]FRDomainConfig

type FRDomainConfig struct {
	ClientID           string `json:"clientId,omitempty"`
	ClientSecretBase64 string `json:"clientSecretBase64,omitempty"`
	UrlBase            string `json:"urlBase,omitempty"`
	AuthScope          string `json:"authScope,omitempty"`
	LDAPConnector      string `json:"ldapConnector,omitempty"`
}

//
//

type ActiveDirectoryConfig struct {
	Domains ADDomains `json:"domains"`
}

type ADDomains map[string]ADDomainConfig

type ADDomainConfig struct {
	DomainController string `json:"domainController,omitempty"`
	Domain           string `json:"domain,omitempty"`
	BindUser         string `json:"bindUser,omitempty"`
	BindPwd          string `json:"bindPwd,omitempty"`
	BaseSearchDn     string `json:"baseSearchDn,omitempty"`
}

type CitrixCloud struct {
	Environments *map[string]CitrixCloudAccountConfig `json:"environments,omitempty" fake:"-"`
}

type CitrixCloudAccountConfig struct {
	CustomerId   string `json:"customerId,omitempty" fake:"{password:true,false,true,false,false,12}"`
	SiteId       string `json:"siteId,omitempty" fake:"{uuid}"`
	ClientId     string `json:"clientId,omitempty" fake:"{uuid}"`
	ClientSecret string `json:"clientSecret,omitempty" fake:"{password:true,true,true,true,false,30}"`
	Region       string `json:"region,omitempty" fake:"{randomstring:[AP,JP,US,EU]}"`
}

type ProxyConfig struct {
	Server    string `json:"server,omitempty"`
	Port      string `json:"port,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
	Overrides string `json:"overrides,omitempty"`
}

type SophosConfig struct {
	Environments map[string]SophosEnvironment `json:"environments,omitempty" fake:"-"`
}

type SophosEnvironment struct {
	Hosts   []string `json:"hosts,omitempty" fake:"-"`
	ApiUser string   `json:"api_user,omitempty" fake:"{username}"`
	ApiKey  string   `json:"api_key,omitempty" fake:"{password:true,true,true,true,false,30}"`
}

type PackerConfig struct {
	Logs struct {
		TenantName    string `json:"tenantName,omitempty"`
		StorageAcct   string `json:"storageAccount,omitempty"`
		BlobContainer struct {
			Hosts     string `json:"hosts,omitempty"`
			Pipelines string `json:"pipelines,omitempty"`
		}
	}
}

type AzureDevOpsConfig struct {
	Packer *PackerConfig `json:"packer,omitempty"`
}

type CertificateManagementConfig struct {
	StorageAccountTenantName string `json:"storageAccountTenantName,omitempty"`
	StorageAccountName       string `json:"storageAccountName,omitempty"`
	ContainerName            string `json:"containerName,omitempty"`
}

type MongoDBConfig struct {
	// Server
	ConnectionString string `json:"connectionString,omitempty"`

	// Databases
	DbAD                      string `json:"dbAD,omitempty"`
	DbAWSMonitoring           string `json:"dbAWSMonitoring,omitempty"`
	DbAzRes                   string `json:"dbAzureResources,omitempty"`
	DbCertificates            string `json:"dbCertificates,omitempty"`
	DbCitrix                  string `json:"dbCitrix,omitempty"`
	DbEntra                   string `json:"dbEntra,omitempty"`
	DbEnvironmentOptimisation string `json:"dbEnvironmentOptimisation,omitempty"`
	DbGeneral                 string `json:"dbGeneral,omitempty"`
	DbIpam                    string `json:"dbIpam,omitempty"`
	DbM365                    string `json:"dbM365,omitempty"`

	// Collections
	CollADUsers string `json:"collADUsers,omitempty"`

	CollAWSMonLogging string `json:"collAWSMonLogging,omitempty"`

	CollAzResImageGalleryImages     string `json:"collAzResImageGalleryImages,omitempty"`
	CollAzResResourceList           string `json:"collAzResResourceList,omitempty"`
	CollAzResGrpsList               string `json:"collAzResGrps,omitempty"`
	CollAzResSKU                    string `json:"collAzResSKU,omitempty"`
	CollAzResTenants                string `json:"collAzResTenants,omitempty"`
	CollAzResVcpuCounts             string `json:"collAzResVcpuCounts,omitempty"`
	CollAzResIPAddresses            string `json:"collAzResIPAddresses,omitempty"`
	CollAzResP2SVpnConnections      string `json:"collAzResP2SVpnConnections,omitempty"`
	CollAzStorageAcctMinTlsVersions string `json:"collAzStorageAcctMinTlsVersions,omitempty"`

	CollCitrixMachineCatalogs string `json:"collCitrixMachineCatalogs,omitempty"`

	CollCertsCaCertInfo     string `json:"collCertsCaCertInfo,omitempty"`
	CollCertsServerCertInfo string `json:"collCertsServerCertInfo,omitempty"`

	CollEntraAppReg                           string `json:"collEntraAppReg,omitempty"`
	CollEntraAppRegCredsExpiring              string `json:"collEntraAppRegCredsExpiring,omitempty"`
	CollEntraRoleAssignmentScheduleInstances  string `json:"collEntraRoleAssignmentScheduleInstances,omitempty"`
	CollEntraRoleEligibilityScheduleInstances string `json:"collEntraRoleEligibilityScheduleInstances,omitempty"`
	CollEntraB2CUsers                         string `json:"collEntraB2CUsers,omitempty"`

	CollEnvOptCosting          string `json:"collEnvOptCosting,omitempty"`
	CollEnvOptCostingMeters    string `json:"collEnvOptCostingMeters,omitempty"`
	CollEnvOptCostingResGrps   string `json:"collEnvOptCostingResGrps,omitempty"`
	CollEnvOptCostingResources string `json:"collEnvOptCostingResources,omitempty"`
	CollEnvOptCostingSubs      string `json:"collEnvOptCostingSubs,omitempty"`
	CollEnvOptCostingTenants   string `json:"collEnvOptCostingTenants,omitempty"`

	CollGenEolTracking   string `json:"collGenEolTracking,omitempty"`
	CollGenSupportAlerts string `json:"collGenSupportAlerts,omitempty"`

	CollIpamIpAddressBlocks string `json:"collIpamIpAddressBlocks,omitempty"`
	CollIpamIpAddresses     string `json:"collIpamIpAddresses,omitempty"`

	CollM365MailboxStatistics string `json:"collM365MailboxStatistics,omitempty"`
}

// type ServerList []string

// type CloudiniConfig struct {
// EncryptConfig bool `json:"encryptConfig,omitempty" fake:"{bool}"`
// }

type AWSConfig struct {
	LogIngestCountQuery string `json:"logIngestCountQuery,omitempty"`
}

//
//

type AzureConfig struct {
	MultiTenantAuth struct {
		Tenants CldConfigTenants `json:"tenants,omitempty" fake:"-"`
	} `json:"multiTenantAuth,omitempty"`
	TenantMap                  map[string]string   `json:"tenantMap,omitempty"`
	CustomSubIdToTenantNameMap map[string][]string `json:"customSubIdToTenantNameMap,omitempty"`
	TenantAliases              map[string]string   `json:"tenantAliases,omitempty"`
	CostDataBlobPrefix         string              `json:"costDataBlobPrefix,omitempty"`
	SkuListSubscription        string              `json:"skuListSubscription,omitempty"`
	SkuListAuthTenant          string              `json:"skuListAuthTenant,omitempty"`
	ResourceLocation           string              `json:"resourceLocation,omitempty"`
	VirtualMachines            map[string]string   `json:"virtualMachines,omitempty"`
	SupportAlerts              SupportAlertsConfig `json:"supportAlerts,omitempty"`
	Ipam                       IpamConfig          `json:"ipam,omitempty"`
}

type IpamConfig struct {
	CidrBlocksToCheck []IpamCidrBlockToCheck `json:"cidrBlocksToCheck"`
}

type IpamCidrBlockToCheck struct {
	CidrBlock string `json:"cidrBlock"`
	BlockTag  string `json:"blockTag"`
}

type CldConfigOptions struct {
	ConfigFile             string
	EncryptUnencryptedFile bool
}

type CldConfigTenants map[string]CldConfigTenantAuth

type CldConfigTenantAuth struct {
	TenantName           string                      `json:"tenantName,omitempty"`
	Default              bool                        `json:"default,omitempty" fake:"-"`
	TenantID             string                      `json:"tenantId,omitempty" fake:"{uuid}"`
	Reader               *CldConfigClientAuthDetails `json:"reader,omitempty"`
	Writer               *CldConfigClientAuthDetails `json:"writer,omitempty"`
	CostExportsLocation  string                      `json:"costExportsLocation,omitempty"`
	CheckExchange        bool                        `json:"checkExchange,omitempty"`
	IsB2C                bool                        `json:"isB2C,omitempty"`
	GetWorkbookAlerts    bool                        `json:"getWorkbookAlerts,omitempty"`
	AWSIngestWorkspaceID string                      `json:"awsIngestWorkspaceID,omitempty"`
	AWSIngestRef         string                      `json:"awsIngestRef,omitempty"`
}

type CldConfigClientAuthDetails struct {
	ClientID     string `json:"clientId,omitempty" fake:"{uuid}"`
	ClientSecret string `json:"clientSecret,omitempty" fake:"{password:true,true,true,true,false,30}"`
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

type SupportAlertsConfig struct {
	DefaultTenant     string            `json:"defaultTenant,omitempty"`
	TenantWorkbookIds map[string]string `json:"tenantWorkbookIds,omitempty"`
	// WorkbookId string `json:"workbookId,omitempty"`
}
