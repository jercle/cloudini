package lib

import (
	"time"
)

type AzureResourceUserAssignedIdentity struct {
	ClientID    string `json:"clientId,omitempty" bson:"clientId,omitempty"`
	PrincipalID string `json:"principalId,omitempty" bson:"principalId,omitempty"`
}

type AzureResourceIdentity struct {
	PrincipalID            string                                       `json:"principalId,omitempty" bson:"principalId,omitempty"`
	TenantID               string                                       `json:"tenantId,omitempty" bson:"tenantId,omitempty"`
	Type                   string                                       `json:"type,omitempty" bson:"type,omitempty"`
	UserAssignedIdentities map[string]AzureResourceUserAssignedIdentity `json:"userAssignedIdentities,omitempty" bson:"userAssignedIdentities,omitempty"`
}

type AzureResourcePlan struct {
	Name          string `json:"name,omitempty" bson:"name,omitempty"`
	Product       string `json:"product,omitempty" bson:"product,omitempty"`
	PromotionCode string `json:"promotionCode,omitempty" bson:"promotionCode,omitempty"`
	Publisher     string `json:"publisher,omitempty" bson:"publisher,omitempty"`
}

type AzureResourcePrivateLinkScopedResource struct {
	ResourceID string `json:"ResourceId,omitempty" bson:"ResourceId,omitempty"`
	ScopeID    string `json:"ScopeId,omitempty" bson:"ScopeId,omitempty"`
}

type AzureResourceRuntimeConfiguration struct {
	Powershell struct {
		BuiltinModules struct {
			Az string `json:"Az,omitempty" bson:"Az,omitempty"`
		} `json:"builtinModules,omitempty" bson:"builtinModules,omitempty"`
	} `json:"powershell,omitempty" bson:"powershell,omitempty"`
	Powershell7 struct {
		BuiltinModules struct {
			Az string `json:"Az,omitempty" bson:"Az,omitempty"`
		} `json:"builtinModules,omitempty" bson:"builtinModules,omitempty"`
	} `json:"powershell7,omitempty" bson:"powershell7,omitempty"`
	Powershell72 struct {
		BuiltinModules struct {
			Az string `json:"Az,omitempty" bson:"Az,omitempty"`
		} `json:"builtinModules,omitempty" bson:"builtinModules,omitempty"`
	} `json:"powershell72,omitempty" bson:"powershell72,omitempty"`
}

type AzureResourceAadAuthenticationParameters struct {
	AadAudience string `json:"aadAudience,omitempty" bson:"aadAudience,omitempty"`
	AadIssuer   string `json:"aadIssuer,omitempty" bson:"aadIssuer,omitempty"`
	AadTenant   string `json:"aadTenant,omitempty" bson:"aadTenant,omitempty"`
}

type AzureResourceAadProfile struct {
	AdminGroupObjectIDs any    `json:"adminGroupObjectIDs,omitempty" bson:"adminGroupObjectIDs,omitempty"`
	AdminUsers          any    `json:"adminUsers,omitempty" bson:"adminUsers,omitempty"`
	EnableAzureRbac     bool   `json:"enableAzureRBAC,omitempty" bson:"enableAzureRBAC,omitempty"`
	Managed             bool   `json:"managed,omitempty" bson:"managed,omitempty"`
	TenantID            string `json:"tenantID,omitempty" bson:"tenantID,omitempty"`
}

type AzureResourceAccessModeSettings struct {
	Exclusions          []any  `json:"exclusions,omitempty" bson:"exclusions,omitempty"`
	IngestionAccessMode string `json:"ingestionAccessMode,omitempty" bson:"ingestionAccessMode,omitempty"`
	QueryAccessMode     string `json:"queryAccessMode,omitempty" bson:"queryAccessMode,omitempty"`
}

type AzureResourceActiveDirectory struct {
	ActiveDirectoryID          string      `json:"activeDirectoryId,omitempty" bson:"activeDirectoryId,omitempty"`
	AesEncryption              bool        `json:"aesEncryption,omitempty" bson:"aesEncryption,omitempty"`
	AllowLocalNfsUsersWithLdap bool        `json:"allowLocalNfsUsersWithLdap,omitempty" bson:"allowLocalNfsUsersWithLdap,omitempty"`
	Dns                        string      `json:"dns,omitempty" bson:"dns,omitempty"`
	Domain                     string      `json:"domain,omitempty" bson:"domain,omitempty"`
	EncryptDcConnections       bool        `json:"encryptDCConnections,omitempty" bson:"encryptDCConnections,omitempty"`
	LdapOverTls                bool        `json:"ldapOverTLS,omitempty" bson:"ldapOverTLS,omitempty"`
	LdapSearchScope            interface{} `json:"ldapSearchScope,omitempty" bson:"ldapSearchScope,omitempty"`
	LdapSigning                bool        `json:"ldapSigning,omitempty" bson:"ldapSigning,omitempty"`
	OrganizationalUnit         string      `json:"organizationalUnit,omitempty" bson:"organizationalUnit,omitempty"`
	Password                   string      `json:"password,omitempty" bson:"password,omitempty"`
	SmbServerName              string      `json:"smbServerName,omitempty" bson:"smbServerName,omitempty"`
	Status                     string      `json:"status,omitempty" bson:"status,omitempty"`
	Username                   string      `json:"username,omitempty" bson:"username,omitempty"`
}

type AzureResourceAddonProfiles struct {
	AciConnectorLinux struct {
		Config  struct{} `json:"config,omitempty" bson:"config,omitempty"`
		Enabled bool     `json:"enabled,omitempty" bson:"enabled,omitempty"`
	} `json:"aciConnectorLinux,omitempty" bson:"aciConnectorLinux,omitempty"`
	Azurepolicy struct {
		Config   any  `json:"config,omitempty" bson:"config,omitempty"`
		Enabled  bool `json:"enabled,omitempty" bson:"enabled,omitempty"`
		Identity struct {
			ClientID   string `json:"clientId,omitempty" bson:"clientId,omitempty"`
			ObjectID   string `json:"objectId,omitempty" bson:"objectId,omitempty"`
			ResourceID string `json:"resourceId,omitempty" bson:"resourceId,omitempty"`
		} `json:"identity,omitempty" bson:"identity,omitempty"`
	} `json:"azurepolicy,omitempty" bson:"azurepolicy,omitempty"`
	HTTPApplicationRouting struct {
		Config  any  `json:"config,omitempty" bson:"config,omitempty"`
		Enabled bool `json:"enabled,omitempty" bson:"enabled,omitempty"`
	} `json:"httpApplicationRouting,omitempty" bson:"httpApplicationRouting,omitempty"`
	IngressApplicationGateway *struct {
		Config  any  `json:"config,omitempty" bson:"config,omitempty"`
		Enabled bool `json:"enabled,omitempty" bson:"enabled,omitempty"`
	} `json:"ingressApplicationGateway,omitempty" bson:"ingressApplicationGateway,omitempty"`
	KubeDashboard struct {
		Config  any  `json:"config,omitempty" bson:"config,omitempty"`
		Enabled bool `json:"enabled,omitempty" bson:"enabled,omitempty"`
	} `json:"kubeDashboard,omitempty" bson:"kubeDashboard,omitempty"`
	Omsagent struct {
		Config struct {
			LogAnalyticsWorkspaceResourceID string `json:"logAnalyticsWorkspaceResourceID,omitempty" bson:"logAnalyticsWorkspaceResourceID,omitempty"`
		} `json:"config,omitempty" bson:"config,omitempty"`
		Enabled  bool `json:"enabled,omitempty" bson:"enabled,omitempty"`
		Identity struct {
			ClientID   string `json:"clientId,omitempty" bson:"clientId,omitempty"`
			ObjectID   string `json:"objectId,omitempty" bson:"objectId,omitempty"`
			ResourceID string `json:"resourceId,omitempty" bson:"resourceId,omitempty"`
		} `json:"identity,omitempty" bson:"identity,omitempty"`
	} `json:"omsagent,omitempty" bson:"omsagent,omitempty"`
}

type AzureResourceAdministrators struct {
	AdministratorType         string `json:"administratorType,omitempty" bson:"administratorType,omitempty"`
	AzureAdOnlyAuthentication bool   `json:"azureADOnlyAuthentication,omitempty" bson:"azureADOnlyAuthentication,omitempty"`
	Login                     string `json:"login,omitempty" bson:"login,omitempty"`
	PrincipalType             string `json:"principalType,omitempty" bson:"principalType,omitempty"`
	Sid                       string `json:"sid,omitempty" bson:"sid,omitempty"`
	TenantID                  string `json:"tenantId,omitempty" bson:"tenantId,omitempty"`
}

type AzureResourceAutoScalerProfile struct {
	BalanceSimilarNodeGroups      string `json:"balance-similar-node-groups,omitempty" bson:"balance-similar-node-groups,omitempty"`
	Expander                      string `json:"expander,omitempty" bson:"expander,omitempty"`
	MaxEmptyBulkDelete            string `json:"max-empty-bulk-delete,omitempty" bson:"max-empty-bulk-delete,omitempty"`
	MaxGracefulTerminationSec     string `json:"max-graceful-termination-sec,omitempty" bson:"max-graceful-termination-sec,omitempty"`
	MaxNodeProvisionTime          string `json:"max-node-provision-time,omitempty" bson:"max-node-provision-time,omitempty"`
	MaxTotalUnreadyPercentage     string `json:"max-total-unready-percentage,omitempty" bson:"max-total-unready-percentage,omitempty"`
	NewPodScaleUpDelay            string `json:"new-pod-scale-up-delay,omitempty" bson:"new-pod-scale-up-delay,omitempty"`
	OkTotalUnreadyCount           string `json:"ok-total-unready-count,omitempty" bson:"ok-total-unready-count,omitempty"`
	ScaleDownDelayAfterAdd        string `json:"scale-down-delay-after-add,omitempty" bson:"scale-down-delay-after-add,omitempty"`
	ScaleDownDelayAfterDelete     string `json:"scale-down-delay-after-delete,omitempty" bson:"scale-down-delay-after-delete,omitempty"`
	ScaleDownDelayAfterFailure    string `json:"scale-down-delay-after-failure,omitempty" bson:"scale-down-delay-after-failure,omitempty"`
	ScaleDownUnneededTime         string `json:"scale-down-unneeded-time,omitempty" bson:"scale-down-unneeded-time,omitempty"`
	ScaleDownUnreadyTime          string `json:"scale-down-unready-time,omitempty" bson:"scale-down-unready-time,omitempty"`
	ScaleDownUtilizationThreshold string `json:"scale-down-utilization-threshold,omitempty" bson:"scale-down-utilization-threshold,omitempty"`
	ScanInterval                  string `json:"scan-interval,omitempty" bson:"scan-interval,omitempty"`
	SkipNodesWithLocalStorage     string `json:"skip-nodes-with-local-storage,omitempty" bson:"skip-nodes-with-local-storage,omitempty"`
	SkipNodesWithSystemPods       string `json:"skip-nodes-with-system-pods,omitempty" bson:"skip-nodes-with-system-pods,omitempty"`
}

type AzureResourceAgentPoolProfile struct {
	AvailabilityZones          []string `json:"availabilityZones,omitempty" bson:"availabilityZones,omitempty"`
	Count                      float64  `json:"count,omitempty" bson:"count,omitempty"`
	CurrentOrchestratorVersion string   `json:"currentOrchestratorVersion,omitempty" bson:"currentOrchestratorVersion,omitempty"`
	EnableAutoScaling          bool     `json:"enableAutoScaling,omitempty" bson:"enableAutoScaling,omitempty"`
	EnableEncryptionAtHost     bool     `json:"enableEncryptionAtHost,omitempty" bson:"enableEncryptionAtHost,omitempty"`
	EnableFips                 bool     `json:"enableFIPS,omitempty" bson:"enableFIPS,omitempty"`
	EnableNodePublicIp         bool     `json:"enableNodePublicIP,omitempty" bson:"enableNodePublicIP,omitempty"`
	EnableUltraSsd             bool     `json:"enableUltraSSD,omitempty" bson:"enableUltraSSD,omitempty"`
	KubeletDiskType            string   `json:"kubeletDiskType,omitempty" bson:"kubeletDiskType,omitempty"`
	MaxCount                   float64  `json:"maxCount,omitempty" bson:"maxCount,omitempty"`
	MaxPods                    float64  `json:"maxPods,omitempty" bson:"maxPods,omitempty"`
	MinCount                   float64  `json:"minCount,omitempty" bson:"minCount,omitempty"`
	Mode                       string   `json:"mode,omitempty" bson:"mode,omitempty"`
	Name                       string   `json:"name,omitempty" bson:"name,omitempty"`
	NodeImageVersion           string   `json:"nodeImageVersion,omitempty" bson:"nodeImageVersion,omitempty"`
	OrchestratorVersion        string   `json:"orchestratorVersion,omitempty" bson:"orchestratorVersion,omitempty"`
	OSDiskSizeGb               float64  `json:"osDiskSizeGB,omitempty" bson:"osDiskSizeGB,omitempty"`
	OSDiskType                 string   `json:"osDiskType,omitempty" bson:"osDiskType,omitempty"`
	OSSku                      string   `json:"osSKU,omitempty" bson:"osSKU,omitempty"`
	OSType                     string   `json:"osType,omitempty" bson:"osType,omitempty"`
	PowerState                 struct {
		Code string `json:"code,omitempty" bson:"code,omitempty"`
	} `json:"powerState,omitempty" bson:"powerState,omitempty"`
	ProvisioningState string    `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
	ScaleDownMode     string    `json:"scaleDownMode,omitempty" bson:"scaleDownMode,omitempty"`
	Type              string    `json:"type,omitempty" bson:"type,omitempty"`
	UpgradeSettings   *struct{} `json:"upgradeSettings,omitempty" bson:"upgradeSettings,omitempty"`
	VmSize            string    `json:"vmSize,omitempty" bson:"vmSize,omitempty"`
	VnetSubnetID      string    `json:"vnetSubnetID,omitempty" bson:"vnetSubnetID,omitempty"`
	WorkloadRuntime   string    `json:"workloadRuntime,omitempty" bson:"workloadRuntime,omitempty"`
}

type AzureResourceAPI struct {
	BrandColor  string `json:"brandColor,omitempty" bson:"brandColor,omitempty"`
	Category    string `json:"category,omitempty" bson:"category,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	DisplayName string `json:"displayName,omitempty" bson:"displayName,omitempty"`
	IconURI     string `json:"iconUri,omitempty" bson:"iconUri,omitempty"`
	ID          string `json:"id,omitempty" bson:"id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Type        string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceAPIServerAccessProfile struct {
	EnablePrivateCluster           bool   `json:"enablePrivateCluster,omitempty" bson:"enablePrivateCluster,omitempty"`
	EnablePrivateClusterPublicFqdn bool   `json:"enablePrivateClusterPublicFQDN,omitempty" bson:"enablePrivateClusterPublicFQDN,omitempty"`
	PrivateDnsZone                 string `json:"privateDNSZone,omitempty" bson:"privateDNSZone,omitempty"`
}

type AzureResourceAppLogsConfiguration struct {
	Destination               *string `json:"destination,omitempty" bson:"destination,omitempty"`
	LogAnalyticsConfiguration *struct {
		CustomerID         string `json:"customerId,omitempty" bson:"customerId,omitempty"`
		DynamicJSONColumns bool   `json:"dynamicJsonColumns,omitempty" bson:"dynamicJsonColumns,omitempty"`
		SharedKey          any    `json:"sharedKey,omitempty" bson:"sharedKey,omitempty"`
	} `json:"logAnalyticsConfiguration,omitempty" bson:"logAnalyticsConfiguration,omitempty"`
}

type AzureResourceAuthorization struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		AuthorizationKey       string `json:"authorizationKey,omitempty" bson:"authorizationKey,omitempty"`
		AuthorizationUseStatus string `json:"authorizationUseStatus,omitempty" bson:"authorizationUseStatus,omitempty"`
		ConnectionResourceURI  string `json:"connectionResourceUri,omitempty" bson:"connectionResourceUri,omitempty"`
		ProvisioningState      string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceAutoScaleConfiguration struct {
	Bounds struct {
		Min float64 `json:"min,omitempty" bson:"min,omitempty"`
	} `json:"bounds,omitempty" bson:"bounds,omitempty"`
}

type AzureResourceBackendAddressPool struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		BackendIpConfigurations []struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"backendIPConfigurations,omitempty" bson:"backendIPConfigurations,omitempty"`
		LoadBalancerBackendAddresses []struct {
			Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
			ID         string `json:"id,omitempty" bson:"id,omitempty"`
			Name       string `json:"name,omitempty" bson:"name,omitempty"`
			Properties struct {
				IpAddress                       string `json:"ipAddress,omitempty" bson:"ipAddress,omitempty"`
				NetworkInterfaceIpConfiguration *struct {
					ID string `json:"id,omitempty" bson:"id,omitempty"`
				} `json:"networkInterfaceIPConfiguration,omitempty" bson:"networkInterfaceIPConfiguration,omitempty"`
				ProvisioningState string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
				Subnet            *struct {
					ID string `json:"id,omitempty" bson:"id,omitempty"`
				} `json:"subnet,omitempty" bson:"subnet,omitempty"`
				VirtualNetwork *struct {
					ID string `json:"id,omitempty" bson:"id,omitempty"`
				} `json:"virtualNetwork,omitempty" bson:"virtualNetwork,omitempty"`
			} `json:"properties,omitempty" bson:"properties,omitempty"`
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"loadBalancerBackendAddresses,omitempty" bson:"loadBalancerBackendAddresses,omitempty"`
		LoadBalancingRules []struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"loadBalancingRules,omitempty" bson:"loadBalancingRules,omitempty"`
		OutboundRules []struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"outboundRules,omitempty" bson:"outboundRules,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceBackup struct {
	BackupRetentionDays float64 `json:"backupRetentionDays,omitempty" bson:"backupRetentionDays,omitempty"`
	EarliestRestoreDate string  `json:"earliestRestoreDate,omitempty" bson:"earliestRestoreDate,omitempty"`
	GeoRedundantBackup  string  `json:"geoRedundantBackup,omitempty" bson:"geoRedundantBackup,omitempty"`
}

type AzureResourceBackupPolicy struct {
	PeriodicModeProperties struct {
		BackupIntervalInMinutes        float64 `json:"backupIntervalInMinutes,omitempty" bson:"backupIntervalInMinutes,omitempty"`
		BackupRetentionIntervalInHours float64 `json:"backupRetentionIntervalInHours,omitempty" bson:"backupRetentionIntervalInHours,omitempty"`
		BackupStorageRedundancy        string  `json:"backupStorageRedundancy,omitempty" bson:"backupStorageRedundancy,omitempty"`
	} `json:"periodicModeProperties,omitempty" bson:"periodicModeProperties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceCallRateLimit struct {
	Rules []AzureResourceCallRateLimitRule `json:"rules,omitempty" bson:"rules,omitempty"`
}

type AzureResourceCallRateLimitRule struct {
	Count                    float64 `json:"count,omitempty" bson:"count,omitempty"`
	DynamicThrottlingEnabled bool    `json:"dynamicThrottlingEnabled,omitempty" bson:"dynamicThrottlingEnabled,omitempty"`
	Key                      string  `json:"key,omitempty" bson:"key,omitempty"`
	MatchPatterns            []struct {
		Method string `json:"method,omitempty" bson:"method,omitempty"`
		Path   string `json:"path,omitempty" bson:"path,omitempty"`
	} `json:"matchPatterns,omitempty" bson:"matchPatterns,omitempty"`
	RenewalPeriod float64 `json:"renewalPeriod,omitempty" bson:"renewalPeriod,omitempty"`
}

type AzureResourceConfiguration struct {
	Dapr               any `json:"dapr,omitempty" bson:"dapr,omitempty"`
	EventTriggerConfig *struct {
		Parallelism            float64 `json:"parallelism,omitempty" bson:"parallelism,omitempty"`
		ReplicaCompletionCount float64 `json:"replicaCompletionCount,omitempty" bson:"replicaCompletionCount,omitempty"`
		Scale                  struct {
			MaxExecutions   float64 `json:"maxExecutions,omitempty" bson:"maxExecutions,omitempty"`
			MinExecutions   float64 `json:"minExecutions,omitempty" bson:"minExecutions,omitempty"`
			PollingInterval float64 `json:"pollingInterval,omitempty" bson:"pollingInterval,omitempty"`
			Rules           []struct {
				Auth []struct {
					SecretRef        string `json:"secretRef,omitempty" bson:"secretRef,omitempty"`
					TriggerParameter string `json:"triggerParameter,omitempty" bson:"triggerParameter,omitempty"`
				} `json:"auth,omitempty" bson:"auth,omitempty"`
				Metadata struct {
					PoolName                   string `json:"poolName,omitempty" bson:"poolName,omitempty"`
					TargetPipelinesQueueLength string `json:"targetPipelinesQueueLength,omitempty" bson:"targetPipelinesQueueLength,omitempty"`
				} `json:"metadata,omitempty" bson:"metadata,omitempty"`
				Name string `json:"name,omitempty" bson:"name,omitempty"`
				Type string `json:"type,omitempty" bson:"type,omitempty"`
			} `json:"rules,omitempty" bson:"rules,omitempty"`
		} `json:"scale,omitempty" bson:"scale,omitempty"`
	} `json:"eventTriggerConfig,omitempty" bson:"eventTriggerConfig,omitempty"`
	IdentitySettings    []any `json:"identitySettings,omitempty" bson:"identitySettings,omitempty"`
	ManualTriggerConfig *struct {
		Parallelism            float64 `json:"parallelism,omitempty" bson:"parallelism,omitempty"`
		ReplicaCompletionCount float64 `json:"replicaCompletionCount,omitempty" bson:"replicaCompletionCount,omitempty"`
	} `json:"manualTriggerConfig,omitempty" bson:"manualTriggerConfig,omitempty"`
	Registries []struct {
		Identity          string `json:"identity,omitempty" bson:"identity,omitempty"`
		PasswordSecretRef string `json:"passwordSecretRef,omitempty" bson:"passwordSecretRef,omitempty"`
		Server            string `json:"server,omitempty" bson:"server,omitempty"`
		Username          string `json:"username,omitempty" bson:"username,omitempty"`
	} `json:"registries,omitempty" bson:"registries,omitempty"`
	ReplicaRetryLimit     float64 `json:"replicaRetryLimit,omitempty" bson:"replicaRetryLimit,omitempty"`
	ReplicaTimeout        float64 `json:"replicaTimeout,omitempty" bson:"replicaTimeout,omitempty"`
	ScheduleTriggerConfig any     `json:"scheduleTriggerConfig,omitempty" bson:"scheduleTriggerConfig,omitempty"`
	Secrets               []struct {
		Name string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"secrets,omitempty" bson:"secrets,omitempty"`
	TriggerType string `json:"triggerType,omitempty" bson:"triggerType,omitempty"`
}

type AzureResourceCreationData struct {
	CreateOption          string `json:"createOption,omitempty" bson:"createOption,omitempty"`
	GalleryImageReference *struct {
		ID string `json:"id,omitempty" bson:"id,omitempty"`
	} `json:"galleryImageReference,omitempty" bson:"galleryImageReference,omitempty"`
	ImageReference *struct {
		ID  string  `json:"id,omitempty" bson:"id,omitempty"`
		Lun float64 `json:"lun,omitempty" bson:"lun,omitempty"`
	} `json:"imageReference,omitempty" bson:"imageReference,omitempty"`
	SecurityDataURI  string  `json:"securityDataUri,omitempty" bson:"securityDataUri,omitempty"`
	SourceResourceID string  `json:"sourceResourceId,omitempty" bson:"sourceResourceId,omitempty"`
	SourceUniqueID   string  `json:"sourceUniqueId,omitempty" bson:"sourceUniqueId,omitempty"`
	SourceURI        string  `json:"sourceUri,omitempty" bson:"sourceUri,omitempty"`
	StorageAccountID string  `json:"storageAccountId,omitempty" bson:"storageAccountId,omitempty"`
	UploadSizeBytes  float64 `json:"uploadSizeBytes,omitempty" bson:"uploadSizeBytes,omitempty"`
}

type AzureResourceCriteria struct {
	AllOf []struct {
		CriterionType string `json:"criterionType,omitempty" bson:"criterionType,omitempty"`
		Dimensions    []struct {
			Name     string   `json:"name,omitempty" bson:"name,omitempty"`
			Operator string   `json:"operator,omitempty" bson:"operator,omitempty"`
			Values   []string `json:"values,omitempty" bson:"values,omitempty"`
		} `json:"dimensions,omitempty" bson:"dimensions,omitempty"`
		FailingPeriods *struct {
			MinFailingPeriodsToAlert  float64 `json:"minFailingPeriodsToAlert,omitempty" bson:"minFailingPeriodsToAlert,omitempty"`
			NumberOfEvaluationPeriods float64 `json:"numberOfEvaluationPeriods,omitempty" bson:"numberOfEvaluationPeriods,omitempty"`
		} `json:"failingPeriods,omitempty" bson:"failingPeriods,omitempty"`
		MetricMeasureColumn string  `json:"metricMeasureColumn,omitempty" bson:"metricMeasureColumn,omitempty"`
		MetricName          string  `json:"metricName,omitempty" bson:"metricName,omitempty"`
		MetricNamespace     string  `json:"metricNamespace,omitempty" bson:"metricNamespace,omitempty"`
		Name                string  `json:"name,omitempty" bson:"name,omitempty"`
		Operator            string  `json:"operator,omitempty" bson:"operator,omitempty"`
		Query               string  `json:"query,omitempty" bson:"query,omitempty"`
		ResourceIDColumn    string  `json:"resourceIdColumn,omitempty" bson:"resourceIdColumn,omitempty"`
		Threshold           float64 `json:"threshold,omitempty" bson:"threshold,omitempty"`
		TimeAggregation     string  `json:"timeAggregation,omitempty" bson:"timeAggregation,omitempty"`
	} `json:"allOf,omitempty" bson:"allOf,omitempty"`
	Odata_Type string `json:"odata.type,omitempty" bson:"odata.type,omitempty"`
}

type AzureResourceCustomDomainConfiguration struct {
	CertificateKeyVaultProperties any    `json:"certificateKeyVaultProperties,omitempty" bson:"certificateKeyVaultProperties,omitempty"`
	CertificatePassword           any    `json:"certificatePassword,omitempty" bson:"certificatePassword,omitempty"`
	CertificateValue              any    `json:"certificateValue,omitempty" bson:"certificateValue,omitempty"`
	CustomDomainVerificationID    string `json:"customDomainVerificationId,omitempty" bson:"customDomainVerificationId,omitempty"`
	DnsSuffix                     any    `json:"dnsSuffix,omitempty" bson:"dnsSuffix,omitempty"`
	ExpirationDate                any    `json:"expirationDate,omitempty" bson:"expirationDate,omitempty"`
	SubjectName                   any    `json:"subjectName,omitempty" bson:"subjectName,omitempty"`
	Thumbprint                    any    `json:"thumbprint,omitempty" bson:"thumbprint,omitempty"`
}

type AzureResourceDataProtection struct {
	Backup struct {
		BackupEnabled  bool   `json:"backupEnabled,omitempty" bson:"backupEnabled,omitempty"`
		BackupPolicyID string `json:"backupPolicyId,omitempty" bson:"backupPolicyId,omitempty"`
		PolicyEnforced bool   `json:"policyEnforced,omitempty" bson:"policyEnforced,omitempty"`
		VaultID        string `json:"vaultId,omitempty" bson:"vaultId,omitempty"`
	} `json:"backup,omitempty" bson:"backup,omitempty"`
	Snapshot struct {
		SnapshotPolicyID string `json:"snapshotPolicyId,omitempty" bson:"snapshotPolicyId,omitempty"`
	} `json:"snapshot,omitempty" bson:"snapshot,omitempty"`
}

type AzureResourceDataSources struct {
	Extensions []struct {
		ExtensionName     string `json:"extensionName,omitempty" bson:"extensionName,omitempty"`
		ExtensionSettings struct {
			Filters []any `json:"Filters,omitempty" bson:"Filters,omitempty"`
		} `json:"extensionSettings,omitempty" bson:"extensionSettings,omitempty"`
		Name    string   `json:"name,omitempty" bson:"name,omitempty"`
		Streams []string `json:"streams,omitempty" bson:"streams,omitempty"`
	} `json:"extensions,omitempty" bson:"extensions,omitempty"`
	LogFiles []struct {
		FilePatterns []string `json:"filePatterns,omitempty" bson:"filePatterns,omitempty"`
		Format       string   `json:"format,omitempty" bson:"format,omitempty"`
		Name         string   `json:"name,omitempty" bson:"name,omitempty"`
		Settings     struct {
			Text struct {
				RecordStartTimestampFormat string `json:"recordStartTimestampFormat,omitempty" bson:"recordStartTimestampFormat,omitempty"`
			} `json:"text,omitempty" bson:"text,omitempty"`
		} `json:"settings,omitempty" bson:"settings,omitempty"`
		Streams []string `json:"streams,omitempty" bson:"streams,omitempty"`
	} `json:"logFiles,omitempty" bson:"logFiles,omitempty"`
	PerformanceCounters []struct {
		CounterSpecifiers          []string `json:"counterSpecifiers,omitempty" bson:"counterSpecifiers,omitempty"`
		Name                       string   `json:"name,omitempty" bson:"name,omitempty"`
		SamplingFrequencyInSeconds float64  `json:"samplingFrequencyInSeconds,omitempty" bson:"samplingFrequencyInSeconds,omitempty"`
		Streams                    []string `json:"streams,omitempty" bson:"streams,omitempty"`
	} `json:"performanceCounters,omitempty" bson:"performanceCounters,omitempty"`
	Syslog []struct {
		FacilityNames []string `json:"facilityNames,omitempty" bson:"facilityNames,omitempty"`
		LogLevels     []string `json:"logLevels,omitempty" bson:"logLevels,omitempty"`
		Name          string   `json:"name,omitempty" bson:"name,omitempty"`
		Streams       []string `json:"streams,omitempty" bson:"streams,omitempty"`
	} `json:"syslog,omitempty" bson:"syslog,omitempty"`
	WindowsEventLogs []struct {
		Name         string   `json:"name,omitempty" bson:"name,omitempty"`
		Streams      []string `json:"streams,omitempty" bson:"streams,omitempty"`
		XPathQueries []string `json:"xPathQueries,omitempty" bson:"xPathQueries,omitempty"`
	} `json:"windowsEventLogs,omitempty" bson:"windowsEventLogs,omitempty"`
	WindowsFirewallLogs []struct {
		Name          string   `json:"name,omitempty" bson:"name,omitempty"`
		ProfileFilter []string `json:"profileFilter,omitempty" bson:"profileFilter,omitempty"`
		Streams       []string `json:"streams,omitempty" bson:"streams,omitempty"`
	} `json:"windowsFirewallLogs,omitempty" bson:"windowsFirewallLogs,omitempty"`
}

type AzureResourceDefaultSecurityRules struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		Access                     string  `json:"access,omitempty" bson:"access,omitempty"`
		Description                string  `json:"description,omitempty" bson:"description,omitempty"`
		DestinationAddressPrefix   string  `json:"destinationAddressPrefix,omitempty" bson:"destinationAddressPrefix,omitempty"`
		DestinationAddressPrefixes []any   `json:"destinationAddressPrefixes,omitempty" bson:"destinationAddressPrefixes,omitempty"`
		DestinationPortRange       string  `json:"destinationPortRange,omitempty" bson:"destinationPortRange,omitempty"`
		DestinationPortRanges      []any   `json:"destinationPortRanges,omitempty" bson:"destinationPortRanges,omitempty"`
		Direction                  string  `json:"direction,omitempty" bson:"direction,omitempty"`
		Priority                   float64 `json:"priority,omitempty" bson:"priority,omitempty"`
		Protocol                   string  `json:"protocol,omitempty" bson:"protocol,omitempty"`
		ProvisioningState          string  `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
		SourceAddressPrefix        string  `json:"sourceAddressPrefix,omitempty" bson:"sourceAddressPrefix,omitempty"`
		SourceAddressPrefixes      []any   `json:"sourceAddressPrefixes,omitempty" bson:"sourceAddressPrefixes,omitempty"`
		SourcePortRange            string  `json:"sourcePortRange,omitempty" bson:"sourcePortRange,omitempty"`
		SourcePortRanges           []any   `json:"sourcePortRanges,omitempty" bson:"sourcePortRanges,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceDefinition struct {
	Schema  string `json:"$schema,omitempty" bson:"_schema,omitempty"`
	Actions struct {
		ComposeEmailResponse *struct {
			Inputs   string `json:"inputs,omitempty" bson:"inputs,omitempty"`
			RunAfter struct {
				CreateHtmlTableWithAlerts []string `json:"Create_HTML_table_with_Alerts,omitempty" bson:"Create_HTML_table_with_Alerts,omitempty"`
			} `json:"runAfter,omitempty" bson:"runAfter,omitempty"`
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"Compose_Email_response,omitempty" bson:"Compose_Email_response,omitempty"`
		ComposeIncidentLink *struct {
			Inputs   string `json:"inputs,omitempty" bson:"inputs,omitempty"`
			RunAfter struct {
				CreateHtmlTableWithEntities []string `json:"Create_HTML_table_with_Entities,omitempty" bson:"Create_HTML_table_with_Entities,omitempty"`
			} `json:"runAfter,omitempty" bson:"runAfter,omitempty"`
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"Compose_Incident_link,omitempty" bson:"Compose_Incident_link,omitempty"`
		CreateHtmlTableWithAlerts *struct {
			Inputs struct {
				Format string `json:"format,omitempty" bson:"format,omitempty"`
				From   string `json:"from,omitempty" bson:"from,omitempty"`
			} `json:"inputs,omitempty" bson:"inputs,omitempty"`
			RunAfter struct {
				SelectAlerts []string `json:"Select_Alerts,omitempty" bson:"Select_Alerts,omitempty"`
			} `json:"runAfter,omitempty" bson:"runAfter,omitempty"`
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"Create_HTML_table_with_Alerts,omitempty" bson:"Create_HTML_table_with_Alerts,omitempty"`
		CreateHtmlTableWithEntities *struct {
			Inputs struct {
				Format string `json:"format,omitempty" bson:"format,omitempty"`
				From   string `json:"from,omitempty" bson:"from,omitempty"`
			} `json:"inputs,omitempty" bson:"inputs,omitempty"`
			RunAfter struct {
				SelectEntities []string `json:"Select_Entities,omitempty" bson:"Select_Entities,omitempty"`
			} `json:"runAfter,omitempty" bson:"runAfter,omitempty"`
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"Create_HTML_table_with_Entities,omitempty" bson:"Create_HTML_table_with_Entities,omitempty"`
		SelectAlerts *struct {
			Inputs struct {
				From   string `json:"from,omitempty" bson:"from,omitempty"`
				Select struct {
					Alerts string `json:"Alerts,omitempty" bson:"Alerts,omitempty"`
				} `json:"select,omitempty" bson:"select,omitempty"`
			} `json:"inputs,omitempty" bson:"inputs,omitempty"`
			RunAfter struct {
				CreateHtmlTableWithEntities []string `json:"Create_HTML_table_with_Entities,omitempty" bson:"Create_HTML_table_with_Entities,omitempty"`
			} `json:"runAfter,omitempty" bson:"runAfter,omitempty"`
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"Select_Alerts,omitempty" bson:"Select_Alerts,omitempty"`
		SelectEntities *struct {
			Inputs struct {
				From   string `json:"from,omitempty" bson:"from,omitempty"`
				Select struct {
					Entity string `json:"Entity,omitempty" bson:"Entity,omitempty"`
					// "Entity Type" cannot be unmarshalled into a struct field by encoding/json.
					// "Entity type" cannot be unmarshalled into a struct field by encoding/json.
				} `json:"select,omitempty" bson:"select,omitempty"`
			} `json:"inputs,omitempty" bson:"inputs,omitempty"`
			RunAfter *struct{} `json:"runAfter,omitempty" bson:"runAfter,omitempty"`
			Type     string    `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"Select_Entities,omitempty" bson:"Select_Entities,omitempty"`
		SendAnEmailWithIncidentDetails *struct {
			Inputs struct {
				Body struct {
					Body       string `json:"Body,omitempty" bson:"Body,omitempty"`
					Importance string `json:"Importance,omitempty" bson:"Importance,omitempty"`
					Subject    string `json:"Subject,omitempty" bson:"Subject,omitempty"`
					To         string `json:"To,omitempty" bson:"To,omitempty"`
				} `json:"body,omitempty" bson:"body,omitempty"`
				Host struct {
					Connection struct {
						Name string `json:"name,omitempty" bson:"name,omitempty"`
					} `json:"connection,omitempty" bson:"connection,omitempty"`
				} `json:"host,omitempty" bson:"host,omitempty"`
				Method string `json:"method,omitempty" bson:"method,omitempty"`
				Path   string `json:"path,omitempty" bson:"path,omitempty"`
			} `json:"inputs,omitempty" bson:"inputs,omitempty"`
			RunAfter struct {
				ComposeEmailResponse []string `json:"Compose_Email_response,omitempty" bson:"Compose_Email_response,omitempty"`
				ComposeIncidentLink  []string `json:"Compose_Incident_link,omitempty" bson:"Compose_Incident_link,omitempty"`
			} `json:"runAfter,omitempty" bson:"runAfter,omitempty"`
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"Send_an_email_with_Incident_details,omitempty" bson:"Send_an_email_with_Incident_details,omitempty"`
	} `json:"actions,omitempty" bson:"actions,omitempty"`
	ContentVersion string    `json:"contentVersion,omitempty" bson:"contentVersion,omitempty"`
	Outputs        *struct{} `json:"outputs,omitempty" bson:"outputs,omitempty"`
	Parameters     struct {
		Connections *struct {
			DefaultValue *struct{} `json:"defaultValue,omitempty" bson:"defaultValue,omitempty"`
			Type         string    `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"$connections,omitempty" bson:"_connections,omitempty"`
		// "Company logo link" cannot be unmarshalled into a struct field by encoding/json.
		// "Report name" cannot be unmarshalled into a struct field by encoding/json.
	} `json:"parameters,omitempty" bson:"parameters,omitempty"`
	Triggers struct {
		MicrosoftSentinelAlert *struct {
			Inputs struct {
				Body struct {
					CallbackURL string `json:"callback_url,omitempty" bson:"callback_url,omitempty"`
				} `json:"body,omitempty" bson:"body,omitempty"`
				Host struct {
					Connection struct {
						Name string `json:"name,omitempty" bson:"name,omitempty"`
					} `json:"connection,omitempty" bson:"connection,omitempty"`
				} `json:"host,omitempty" bson:"host,omitempty"`
				Path string `json:"path,omitempty" bson:"path,omitempty"`
			} `json:"inputs,omitempty" bson:"inputs,omitempty"`
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"Microsoft_Sentinel_alert,omitempty" bson:"Microsoft_Sentinel_alert,omitempty"`
		MicrosoftSentinelIncident *struct {
			Inputs struct {
				Body struct {
					CallbackURL string `json:"callback_url,omitempty" bson:"callback_url,omitempty"`
				} `json:"body,omitempty" bson:"body,omitempty"`
				Host struct {
					Connection struct {
						Name string `json:"name,omitempty" bson:"name,omitempty"`
					} `json:"connection,omitempty" bson:"connection,omitempty"`
				} `json:"host,omitempty" bson:"host,omitempty"`
				Path string `json:"path,omitempty" bson:"path,omitempty"`
			} `json:"inputs,omitempty" bson:"inputs,omitempty"`
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"Microsoft_Sentinel_incident,omitempty" bson:"Microsoft_Sentinel_incident,omitempty"`
	} `json:"triggers,omitempty" bson:"triggers,omitempty"`
}

type AzureResourceDestinations struct {
	LogAnalytics []struct {
		Name                string `json:"name,omitempty" bson:"name,omitempty"`
		WorkspaceID         string `json:"workspaceId,omitempty" bson:"workspaceId,omitempty"`
		WorkspaceResourceID string `json:"workspaceResourceId,omitempty" bson:"workspaceResourceId,omitempty"`
	} `json:"logAnalytics,omitempty" bson:"logAnalytics,omitempty"`
}

type AzureResourceDiagnosticsProfile *struct {
	BootDiagnostics struct {
		Enabled    bool   `json:"enabled,omitempty" bson:"enabled,omitempty"`
		StorageURI string `json:"storageUri,omitempty" bson:"storageUri,omitempty"`
	} `json:"bootDiagnostics,omitempty" bson:"bootDiagnostics,omitempty"`
}

type AzureResourceDistribute struct {
	ArtifactTags struct {
		Baseosimg string `json:"baseosimg,omitempty" bson:"baseosimg,omitempty"`
		Source    string `json:"source,omitempty" bson:"source,omitempty"`
	} `json:"artifactTags,omitempty" bson:"artifactTags,omitempty"`
	ExcludeFromLatest  bool     `json:"excludeFromLatest,omitempty" bson:"excludeFromLatest,omitempty"`
	GalleryImageID     string   `json:"galleryImageId,omitempty" bson:"galleryImageId,omitempty"`
	ReplicationRegions []string `json:"replicationRegions,omitempty" bson:"replicationRegions,omitempty"`
	RunOutputName      string   `json:"runOutputName,omitempty" bson:"runOutputName,omitempty"`
	Type               string   `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceDnsSettings struct {
	AppliedDnsServers        []string `json:"appliedDnsServers,omitempty" bson:"appliedDnsServers,omitempty"`
	DnsServers               []string `json:"dnsServers,omitempty" bson:"dnsServers,omitempty"`
	DomainNameLabel          string   `json:"domainNameLabel,omitempty" bson:"domainNameLabel,omitempty"`
	EnableProxy              bool     `json:"enableProxy,omitempty" bson:"enableProxy,omitempty"`
	Fqdn                     string   `json:"fqdn,omitempty" bson:"fqdn,omitempty"`
	InternalDomainNameSuffix string   `json:"internalDomainNameSuffix,omitempty" bson:"internalDomainNameSuffix,omitempty"`
	Servers                  []string `json:"servers,omitempty" bson:"servers,omitempty"`
}

type AzureResourceEmailReceivers struct {
	EmailAddress         string `json:"emailAddress,omitempty" bson:"emailAddress,omitempty"`
	Name                 string `json:"name,omitempty" bson:"name,omitempty"`
	Status               string `json:"status,omitempty" bson:"status,omitempty"`
	UseCommonAlertSchema bool   `json:"useCommonAlertSchema,omitempty" bson:"useCommonAlertSchema,omitempty"`
}

type AzureResourceEncryption struct {
	Identity *struct {
		UserAssignedIdentity any `json:"userAssignedIdentity,omitempty" bson:"userAssignedIdentity,omitempty"`
	} `json:"identity,omitempty" bson:"identity,omitempty"`
	KeySource                       string `json:"keySource,omitempty" bson:"keySource,omitempty"`
	RequireInfrastructureEncryption bool   `json:"requireInfrastructureEncryption,omitempty" bson:"requireInfrastructureEncryption,omitempty"`
	Services                        *struct {
		Blob struct {
			Enabled         bool   `json:"enabled,omitempty" bson:"enabled,omitempty"`
			KeyType         string `json:"keyType,omitempty" bson:"keyType,omitempty"`
			LastEnabledTime string `json:"lastEnabledTime,omitempty" bson:"lastEnabledTime,omitempty"`
		} `json:"blob,omitempty" bson:"blob,omitempty"`
		File struct {
			Enabled         bool   `json:"enabled,omitempty" bson:"enabled,omitempty"`
			KeyType         string `json:"keyType,omitempty" bson:"keyType,omitempty"`
			LastEnabledTime string `json:"lastEnabledTime,omitempty" bson:"lastEnabledTime,omitempty"`
		} `json:"file,omitempty" bson:"file,omitempty"`
	} `json:"services,omitempty" bson:"services,omitempty"`
	Status string `json:"status,omitempty" bson:"status,omitempty"`
	Type   string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceEncryptionSettingsCollection struct {
	Enabled            bool `json:"enabled,omitempty" bson:"enabled,omitempty"`
	EncryptionSettings []struct {
		DiskEncryptionKey struct {
			SecretURL   string `json:"secretUrl,omitempty" bson:"secretUrl,omitempty"`
			SourceVault struct {
				ID string `json:"id,omitempty" bson:"id,omitempty"`
			} `json:"sourceVault,omitempty" bson:"sourceVault,omitempty"`
		} `json:"diskEncryptionKey,omitempty" bson:"diskEncryptionKey,omitempty"`
		KeyEncryptionKey *struct {
			KeyURL      string `json:"keyUrl,omitempty" bson:"keyUrl,omitempty"`
			SourceVault struct {
				ID string `json:"id,omitempty" bson:"id,omitempty"`
			} `json:"sourceVault,omitempty" bson:"sourceVault,omitempty"`
		} `json:"keyEncryptionKey,omitempty" bson:"keyEncryptionKey,omitempty"`
	} `json:"encryptionSettings,omitempty" bson:"encryptionSettings,omitempty"`
	EncryptionSettingsVersion string `json:"encryptionSettingsVersion,omitempty" bson:"encryptionSettingsVersion,omitempty"`
}

type AzureResourceEndpointsConfiguration struct {
	Connector struct {
		OutgoingIpAddresses []struct {
			Address string `json:"address,omitempty" bson:"address,omitempty"`
		} `json:"outgoingIpAddresses,omitempty" bson:"outgoingIpAddresses,omitempty"`
	} `json:"connector,omitempty" bson:"connector,omitempty"`
	Workflow struct {
		AccessEndpointIpAddresses []struct {
			Address string `json:"address,omitempty" bson:"address,omitempty"`
		} `json:"accessEndpointIpAddresses,omitempty" bson:"accessEndpointIpAddresses,omitempty"`
		OutgoingIpAddresses []struct {
			Address string `json:"address,omitempty" bson:"address,omitempty"`
		} `json:"outgoingIpAddresses,omitempty" bson:"outgoingIpAddresses,omitempty"`
	} `json:"workflow,omitempty" bson:"workflow,omitempty"`
}

type AzureResourceExpressRouteConnections struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		EnableInternetSecurity     bool `json:"enableInternetSecurity,omitempty" bson:"enableInternetSecurity,omitempty"`
		ExpressRouteCircuitPeering struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"expressRouteCircuitPeering,omitempty" bson:"expressRouteCircuitPeering,omitempty"`
		ExpressRouteGatewayBypass bool   `json:"expressRouteGatewayBypass,omitempty" bson:"expressRouteGatewayBypass,omitempty"`
		ProvisioningState         string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
		ResourceGuid              string `json:"resourceGuid,omitempty" bson:"resourceGuid,omitempty"`
		RoutingConfiguration      struct {
			AssociatedRouteTable struct {
				ID string `json:"id,omitempty" bson:"id,omitempty"`
			} `json:"associatedRouteTable,omitempty" bson:"associatedRouteTable,omitempty"`
			PropagatedRouteTables struct {
				Ids []struct {
					ID string `json:"id,omitempty" bson:"id,omitempty"`
				} `json:"ids,omitempty" bson:"ids,omitempty"`
				Labels []string `json:"labels,omitempty" bson:"labels,omitempty"`
			} `json:"propagatedRouteTables,omitempty" bson:"propagatedRouteTables,omitempty"`
		} `json:"routingConfiguration,omitempty" bson:"routingConfiguration,omitempty"`
		RoutingWeight float64 `json:"routingWeight,omitempty" bson:"routingWeight,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceExtended struct {
	InstanceView *struct {
		ComputerName     string `json:"computerName,omitempty" bson:"computerName,omitempty"`
		HyperVGeneration string `json:"hyperVGeneration,omitempty" bson:"hyperVGeneration,omitempty"`
		OSName           string `json:"osName,omitempty" bson:"osName,omitempty"`
		OSVersion        string `json:"osVersion,omitempty" bson:"osVersion,omitempty"`
		PowerState       struct {
			Code          string `json:"code,omitempty" bson:"code,omitempty"`
			DisplayStatus string `json:"displayStatus,omitempty" bson:"displayStatus,omitempty"`
			Level         string `json:"level,omitempty" bson:"level,omitempty"`
		} `json:"powerState,omitempty" bson:"powerState,omitempty"`
	} `json:"instanceView,omitempty" bson:"instanceView,omitempty"`
}

type AzureResourceFlowAnalyticsConfiguration struct {
	NetworkWatcherFlowAnalyticsConfiguration *struct {
		Enabled                  bool    `json:"enabled,omitempty" bson:"enabled,omitempty"`
		TrafficAnalyticsInterval float64 `json:"trafficAnalyticsInterval,omitempty" bson:"trafficAnalyticsInterval,omitempty"`
		WorkspaceID              string  `json:"workspaceId,omitempty" bson:"workspaceId,omitempty"`
		WorkspaceRegion          string  `json:"workspaceRegion,omitempty" bson:"workspaceRegion,omitempty"`
		WorkspaceResourceID      string  `json:"workspaceResourceId,omitempty" bson:"workspaceResourceId,omitempty"`
	} `json:"networkWatcherFlowAnalyticsConfiguration,omitempty" bson:"networkWatcherFlowAnalyticsConfiguration,omitempty"`
}

type AzureResourceFrontendIpConfiguration struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		InboundNatRules []struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"inboundNatRules,omitempty" bson:"inboundNatRules,omitempty"`
		LoadBalancingRules []struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"loadBalancingRules,omitempty" bson:"loadBalancingRules,omitempty"`
		OutboundRules []struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"outboundRules,omitempty" bson:"outboundRules,omitempty"`
		PrivateIpAddress          string `json:"privateIPAddress,omitempty" bson:"privateIPAddress,omitempty"`
		PrivateIpAddressVersion   string `json:"privateIPAddressVersion,omitempty" bson:"privateIPAddressVersion,omitempty"`
		PrivateIpAllocationMethod string `json:"privateIPAllocationMethod,omitempty" bson:"privateIPAllocationMethod,omitempty"`
		ProvisioningState         string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
		PublicIpAddress           *struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"publicIPAddress,omitempty" bson:"publicIPAddress,omitempty"`
		Subnet *struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"subnet,omitempty" bson:"subnet,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type  string   `json:"type,omitempty" bson:"type,omitempty"`
	Zones []string `json:"zones,omitempty" bson:"zones,omitempty"`
}

type AzureResourceGeoDataReplication struct {
	Locations []struct {
		LocationName string `json:"locationName,omitempty" bson:"locationName,omitempty"`
		ReplicaState string `json:"replicaState,omitempty" bson:"replicaState,omitempty"`
		RoleType     string `json:"roleType,omitempty" bson:"roleType,omitempty"`
	} `json:"locations,omitempty" bson:"locations,omitempty"`
	MaxReplicationLagDurationInSeconds float64 `json:"maxReplicationLagDurationInSeconds,omitempty" bson:"maxReplicationLagDurationInSeconds,omitempty"`
}

type AzureResourceHostNameSslStates struct {
	CertificateResourceID any    `json:"certificateResourceId,omitempty" bson:"certificateResourceId,omitempty"`
	HostType              string `json:"hostType,omitempty" bson:"hostType,omitempty"`
	IpBasedSslResult      any    `json:"ipBasedSslResult,omitempty" bson:"ipBasedSslResult,omitempty"`
	IpBasedSslState       string `json:"ipBasedSslState,omitempty" bson:"ipBasedSslState,omitempty"`
	Name                  string `json:"name,omitempty" bson:"name,omitempty"`
	SslState              string `json:"sslState,omitempty" bson:"sslState,omitempty"`
	Thumbprint            any    `json:"thumbprint,omitempty" bson:"thumbprint,omitempty"`
	ToUpdate              any    `json:"toUpdate,omitempty" bson:"toUpdate,omitempty"`
	ToUpdateIpBasedSsl    any    `json:"toUpdateIpBasedSsl,omitempty" bson:"toUpdateIpBasedSsl,omitempty"`
	VirtualIp             any    `json:"virtualIP,omitempty" bson:"virtualIP,omitempty"`
	VirtualIPv6           any    `json:"virtualIPv6,omitempty" bson:"virtualIPv6,omitempty"`
}

type AzureResourceHubIpAddresses struct {
	PrivateIpAddress string `json:"privateIPAddress,omitempty" bson:"privateIPAddress,omitempty"`
	PublicIPs        struct {
		Addresses []struct {
			Address string `json:"address,omitempty" bson:"address,omitempty"`
		} `json:"addresses,omitempty" bson:"addresses,omitempty"`
		Count float64 `json:"count,omitempty" bson:"count,omitempty"`
	} `json:"publicIPs,omitempty" bson:"publicIPs,omitempty"`
}

type AzureResourceIdentifier struct {
	Offer      string `json:"offer,omitempty" bson:"offer,omitempty"`
	Publisher  string `json:"publisher,omitempty" bson:"publisher,omitempty"`
	Sku        string `json:"sku,omitempty" bson:"sku,omitempty"`
	UniqueName string `json:"uniqueName,omitempty" bson:"uniqueName,omitempty"`
}

type AzureResourceIdentityProfile struct {
	Kubeletidentity struct {
		ClientID   string `json:"clientId,omitempty" bson:"clientId,omitempty"`
		ObjectID   string `json:"objectId,omitempty" bson:"objectId,omitempty"`
		ResourceID string `json:"resourceId,omitempty" bson:"resourceId,omitempty"`
	} `json:"kubeletidentity,omitempty" bson:"kubeletidentity,omitempty"`
}

type AzureResourceInboundNatRule struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		AllowBackendPortConflict         bool    `json:"allowBackendPortConflict,omitempty" bson:"allowBackendPortConflict,omitempty"`
		BackendPort                      float64 `json:"backendPort,omitempty" bson:"backendPort,omitempty"`
		EnableDestinationServiceEndpoint bool    `json:"enableDestinationServiceEndpoint,omitempty" bson:"enableDestinationServiceEndpoint,omitempty"`
		EnableFloatingIp                 bool    `json:"enableFloatingIP,omitempty" bson:"enableFloatingIP,omitempty"`
		EnableTcpReset                   bool    `json:"enableTcpReset,omitempty" bson:"enableTcpReset,omitempty"`
		FrontendIpConfiguration          struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"frontendIPConfiguration,omitempty" bson:"frontendIPConfiguration,omitempty"`
		FrontendPort         float64 `json:"frontendPort,omitempty" bson:"frontendPort,omitempty"`
		IdleTimeoutInMinutes float64 `json:"idleTimeoutInMinutes,omitempty" bson:"idleTimeoutInMinutes,omitempty"`
		Protocol             string  `json:"protocol,omitempty" bson:"protocol,omitempty"`
		ProvisioningState    string  `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceInstallPatches struct {
	LinuxParameters *struct {
		ClassificationsToInclude  []string `json:"classificationsToInclude,omitempty" bson:"classificationsToInclude,omitempty"`
		PackageNameMasksToExclude []string `json:"packageNameMasksToExclude,omitempty" bson:"packageNameMasksToExclude,omitempty"`
	} `json:"linuxParameters,omitempty" bson:"linuxParameters,omitempty"`
	RebootSetting     string `json:"rebootSetting,omitempty" bson:"rebootSetting,omitempty"`
	WindowsParameters struct {
		ClassificationsToInclude []string `json:"classificationsToInclude,omitempty" bson:"classificationsToInclude,omitempty"`
		KbNumbersToExclude       []any    `json:"kbNumbersToExclude,omitempty" bson:"kbNumbersToExclude,omitempty"`
	} `json:"windowsParameters,omitempty" bson:"windowsParameters,omitempty"`
}

type AzureResourceIntrusionDetection struct {
	Configuration struct {
		BypassTrafficSettings []struct {
			Description          string   `json:"description,omitempty" bson:"description,omitempty"`
			DestinationAddresses []string `json:"destinationAddresses,omitempty" bson:"destinationAddresses,omitempty"`
			DestinationIpGroups  []any    `json:"destinationIpGroups,omitempty" bson:"destinationIpGroups,omitempty"`
			DestinationPorts     []string `json:"destinationPorts,omitempty" bson:"destinationPorts,omitempty"`
			Name                 string   `json:"name,omitempty" bson:"name,omitempty"`
			Protocol             string   `json:"protocol,omitempty" bson:"protocol,omitempty"`
			SourceAddresses      []string `json:"sourceAddresses,omitempty" bson:"sourceAddresses,omitempty"`
			SourceIpGroups       []any    `json:"sourceIpGroups,omitempty" bson:"sourceIpGroups,omitempty"`
		} `json:"bypassTrafficSettings,omitempty" bson:"bypassTrafficSettings,omitempty"`
		SignatureOverrides []struct {
			ID   string `json:"id,omitempty" bson:"id,omitempty"`
			Mode string `json:"mode,omitempty" bson:"mode,omitempty"`
		} `json:"signatureOverrides,omitempty" bson:"signatureOverrides,omitempty"`
	} `json:"configuration,omitempty" bson:"configuration,omitempty"`
	Mode string `json:"mode,omitempty" bson:"mode,omitempty"`
}

type AzureResourceIpConfiguration struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		LoadBalancerBackendAddressPools []struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"loadBalancerBackendAddressPools,omitempty" bson:"loadBalancerBackendAddressPools,omitempty"`
		Primary                         bool   `json:"primary,omitempty" bson:"primary,omitempty"`
		PrivateIpAddress                string `json:"privateIPAddress,omitempty" bson:"privateIPAddress,omitempty"`
		PrivateIpAddressVersion         string `json:"privateIPAddressVersion,omitempty" bson:"privateIPAddressVersion,omitempty"`
		PrivateIpAllocationMethod       string `json:"privateIPAllocationMethod,omitempty" bson:"privateIPAllocationMethod,omitempty"`
		PrivateLinkConnectionProperties *struct {
			Fqdns              []string `json:"fqdns,omitempty" bson:"fqdns,omitempty"`
			GroupID            string   `json:"groupId,omitempty" bson:"groupId,omitempty"`
			RequiredMemberName string   `json:"requiredMemberName,omitempty" bson:"requiredMemberName,omitempty"`
		} `json:"privateLinkConnectionProperties,omitempty" bson:"privateLinkConnectionProperties,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
		PublicIpAddress   *struct {
			ID         string `json:"id,omitempty" bson:"id,omitempty"`
			Properties *struct {
				DeleteOption string `json:"deleteOption,omitempty" bson:"deleteOption,omitempty"`
			} `json:"properties,omitempty" bson:"properties,omitempty"`
		} `json:"publicIPAddress,omitempty" bson:"publicIPAddress,omitempty"`
		Subnet *struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"subnet,omitempty" bson:"subnet,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceKeysMetadata struct {
	PrimaryMasterKey struct {
		GenerationTime string `json:"generationTime,omitempty" bson:"generationTime,omitempty"`
	} `json:"primaryMasterKey,omitempty" bson:"primaryMasterKey,omitempty"`
	PrimaryReadonlyMasterKey struct {
		GenerationTime string `json:"generationTime,omitempty" bson:"generationTime,omitempty"`
	} `json:"primaryReadonlyMasterKey,omitempty" bson:"primaryReadonlyMasterKey,omitempty"`
	SecondaryMasterKey struct {
		GenerationTime string `json:"generationTime,omitempty" bson:"generationTime,omitempty"`
	} `json:"secondaryMasterKey,omitempty" bson:"secondaryMasterKey,omitempty"`
	SecondaryReadonlyMasterKey struct {
		GenerationTime string `json:"generationTime,omitempty" bson:"generationTime,omitempty"`
	} `json:"secondaryReadonlyMasterKey,omitempty" bson:"secondaryReadonlyMasterKey,omitempty"`
}

type AzureResourceLense struct {
	Order float64 `json:"order,omitempty" bson:"order,omitempty"`
	Parts []struct {
		Metadata struct {
			DeepLink          string `json:"deepLink,omitempty" bson:"deepLink,omitempty"`
			DefaultMenuItemID string `json:"defaultMenuItemId,omitempty" bson:"defaultMenuItemId,omitempty"`
			Filters           *struct {
				EntityName *struct {
					Model struct {
						Operator string   `json:"operator,omitempty" bson:"operator,omitempty"`
						Values   []string `json:"values,omitempty" bson:"values,omitempty"`
					} `json:"model,omitempty" bson:"model,omitempty"`
				} `json:"EntityName,omitempty" bson:"EntityName,omitempty"`
				MsPortalFxTimeRange struct {
					Model struct {
						Format      string `json:"format,omitempty" bson:"format,omitempty"`
						Granularity string `json:"granularity,omitempty" bson:"granularity,omitempty"`
						Relative    string `json:"relative,omitempty" bson:"relative,omitempty"`
					} `json:"model,omitempty" bson:"model,omitempty"`
				} `json:"MsPortalFx_TimeRange,omitempty" bson:"MsPortalFx_TimeRange,omitempty"`
			} `json:"filters,omitempty" bson:"filters,omitempty"`
			Inputs []struct {
				IsOptional bool   `json:"isOptional,omitempty" bson:"isOptional,omitempty"`
				Name       string `json:"name,omitempty" bson:"name,omitempty"`
				Value      any    `json:"value,omitempty" bson:"value,omitempty"`
			} `json:"inputs,omitempty" bson:"inputs,omitempty"`
			Settings *struct {
				Content struct {
					Content        *string `json:"content,omitempty" bson:"content,omitempty"`
					MarkdownSource float64 `json:"markdownSource,omitempty" bson:"markdownSource,omitempty"`
					MarkdownURI    string  `json:"markdownUri,omitempty" bson:"markdownUri,omitempty"`
					Options        *struct {
						Chart struct {
							Grouping *struct {
								Dimension string  `json:"dimension,omitempty" bson:"dimension,omitempty"`
								Sort      float64 `json:"sort,omitempty" bson:"sort,omitempty"`
								Top       float64 `json:"top,omitempty" bson:"top,omitempty"`
							} `json:"grouping,omitempty" bson:"grouping,omitempty"`
							Metrics []struct {
								AggregationType     float64 `json:"aggregationType,omitempty" bson:"aggregationType,omitempty"`
								MetricVisualization struct {
									DisplayName         string `json:"displayName,omitempty" bson:"displayName,omitempty"`
									ResourceDisplayName string `json:"resourceDisplayName,omitempty" bson:"resourceDisplayName,omitempty"`
								} `json:"metricVisualization,omitempty" bson:"metricVisualization,omitempty"`
								Name             string `json:"name,omitempty" bson:"name,omitempty"`
								Namespace        string `json:"namespace,omitempty" bson:"namespace,omitempty"`
								ResourceMetadata struct {
									ID string `json:"id,omitempty" bson:"id,omitempty"`
								} `json:"resourceMetadata,omitempty" bson:"resourceMetadata,omitempty"`
							} `json:"metrics,omitempty" bson:"metrics,omitempty"`
							Title         string  `json:"title,omitempty" bson:"title,omitempty"`
							TitleKind     float64 `json:"titleKind,omitempty" bson:"titleKind,omitempty"`
							Visualization struct {
								AxisVisualization struct {
									X struct {
										AxisType  float64 `json:"axisType,omitempty" bson:"axisType,omitempty"`
										IsVisible bool    `json:"isVisible,omitempty" bson:"isVisible,omitempty"`
									} `json:"x,omitempty" bson:"x,omitempty"`
									Y struct {
										AxisType  float64 `json:"axisType,omitempty" bson:"axisType,omitempty"`
										IsVisible bool    `json:"isVisible,omitempty" bson:"isVisible,omitempty"`
									} `json:"y,omitempty" bson:"y,omitempty"`
								} `json:"axisVisualization,omitempty" bson:"axisVisualization,omitempty"`
								ChartType           float64 `json:"chartType,omitempty" bson:"chartType,omitempty"`
								DisablePinning      bool    `json:"disablePinning,omitempty" bson:"disablePinning,omitempty"`
								LegendVisualization struct {
									HideSubtitle bool    `json:"hideSubtitle,omitempty" bson:"hideSubtitle,omitempty"`
									IsVisible    bool    `json:"isVisible,omitempty" bson:"isVisible,omitempty"`
									Position     float64 `json:"position,omitempty" bson:"position,omitempty"`
								} `json:"legendVisualization,omitempty" bson:"legendVisualization,omitempty"`
							} `json:"visualization,omitempty" bson:"visualization,omitempty"`
						} `json:"chart,omitempty" bson:"chart,omitempty"`
					} `json:"options,omitempty" bson:"options,omitempty"`
					Settings *struct {
						Content  string `json:"content,omitempty" bson:"content,omitempty"`
						Subtitle string `json:"subtitle,omitempty" bson:"subtitle,omitempty"`
						Title    string `json:"title,omitempty" bson:"title,omitempty"`
					} `json:"settings,omitempty" bson:"settings,omitempty"`
					Subtitle *string `json:"subtitle,omitempty" bson:"subtitle,omitempty"`
					Title    string  `json:"title,omitempty" bson:"title,omitempty"`
				} `json:"content,omitempty" bson:"content,omitempty"`
			} `json:"settings,omitempty" bson:"settings,omitempty"`
			Type      string `json:"type,omitempty" bson:"type,omitempty"`
			ViewState *struct {
				Content struct {
					ConfigurationID string `json:"configurationId,omitempty" bson:"configurationId,omitempty"`
				} `json:"content,omitempty" bson:"content,omitempty"`
			} `json:"viewState,omitempty" bson:"viewState,omitempty"`
		} `json:"metadata,omitempty" bson:"metadata,omitempty"`
		Position struct {
			ColSpan float64 `json:"colSpan,omitempty" bson:"colSpan,omitempty"`
			RowSpan float64 `json:"rowSpan,omitempty" bson:"rowSpan,omitempty"`
			X       float64 `json:"x,omitempty" bson:"x,omitempty"`
			Y       float64 `json:"y,omitempty" bson:"y,omitempty"`
		} `json:"position,omitempty" bson:"position,omitempty"`
	} `json:"parts,omitempty" bson:"parts,omitempty"`
}

type AzureResourceLink struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		AdminState    string `json:"adminState,omitempty" bson:"adminState,omitempty"`
		ConnectorType string `json:"connectorType,omitempty" bson:"connectorType,omitempty"`
		InterfaceName string `json:"interfaceName,omitempty" bson:"interfaceName,omitempty"`
		MacSecConfig  struct {
			CakSecretIdentifier string `json:"cakSecretIdentifier,omitempty" bson:"cakSecretIdentifier,omitempty"`
			Cipher              string `json:"cipher,omitempty" bson:"cipher,omitempty"`
			CknSecretIdentifier string `json:"cknSecretIdentifier,omitempty" bson:"cknSecretIdentifier,omitempty"`
			SciState            string `json:"sciState,omitempty" bson:"sciState,omitempty"`
		} `json:"macSecConfig,omitempty" bson:"macSecConfig,omitempty"`
		PatchPanelID      string `json:"patchPanelId,omitempty" bson:"patchPanelId,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
		RackID            string `json:"rackId,omitempty" bson:"rackId,omitempty"`
		RouterName        string `json:"routerName,omitempty" bson:"routerName,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceLinuxProfile struct {
	AdminUsername string `json:"adminUsername,omitempty" bson:"adminUsername,omitempty"`
	SSH           struct {
		PublicKeys []struct {
			KeyData string `json:"keyData,omitempty" bson:"keyData,omitempty"`
		} `json:"publicKeys,omitempty" bson:"publicKeys,omitempty"`
	} `json:"ssh,omitempty" bson:"ssh,omitempty"`
}

type AzureResourceLoadBalancingRule struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		AllowBackendPortConflict bool `json:"allowBackendPortConflict,omitempty" bson:"allowBackendPortConflict,omitempty"`
		BackendAddressPool       struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"backendAddressPool,omitempty" bson:"backendAddressPool,omitempty"`
		BackendAddressPools []struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"backendAddressPools,omitempty" bson:"backendAddressPools,omitempty"`
		BackendPort                      float64 `json:"backendPort,omitempty" bson:"backendPort,omitempty"`
		DisableOutboundSnat              bool    `json:"disableOutboundSnat,omitempty" bson:"disableOutboundSnat,omitempty"`
		EnableDestinationServiceEndpoint bool    `json:"enableDestinationServiceEndpoint,omitempty" bson:"enableDestinationServiceEndpoint,omitempty"`
		EnableFloatingIp                 bool    `json:"enableFloatingIP,omitempty" bson:"enableFloatingIP,omitempty"`
		EnableTcpReset                   bool    `json:"enableTcpReset,omitempty" bson:"enableTcpReset,omitempty"`
		FrontendIpConfiguration          struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"frontendIPConfiguration,omitempty" bson:"frontendIPConfiguration,omitempty"`
		FrontendPort         float64 `json:"frontendPort,omitempty" bson:"frontendPort,omitempty"`
		IdleTimeoutInMinutes float64 `json:"idleTimeoutInMinutes,omitempty" bson:"idleTimeoutInMinutes,omitempty"`
		LoadDistribution     string  `json:"loadDistribution,omitempty" bson:"loadDistribution,omitempty"`
		Probe                *struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"probe,omitempty" bson:"probe,omitempty"`
		Protocol          string `json:"protocol,omitempty" bson:"protocol,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceLocation struct {
	DocumentEndpoint  string  `json:"documentEndpoint,omitempty" bson:"documentEndpoint,omitempty"`
	FailoverPriority  float64 `json:"failoverPriority,omitempty" bson:"failoverPriority,omitempty"`
	ID                string  `json:"id,omitempty" bson:"id,omitempty"`
	IsZoneRedundant   bool    `json:"isZoneRedundant,omitempty" bson:"isZoneRedundant,omitempty"`
	LocationName      string  `json:"locationName,omitempty" bson:"locationName,omitempty"`
	ProvisioningState string  `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
}

type AzureResourceMaintenanceWindow struct {
	CustomWindow  string  `json:"customWindow,omitempty" bson:"customWindow,omitempty"`
	DayOfWeek     float64 `json:"dayOfWeek,omitempty" bson:"dayOfWeek,omitempty"`
	Duration      string  `json:"duration,omitempty" bson:"duration,omitempty"`
	RecurEvery    string  `json:"recurEvery,omitempty" bson:"recurEvery,omitempty"`
	StartDateTime string  `json:"startDateTime,omitempty" bson:"startDateTime,omitempty"`
	StartHour     float64 `json:"startHour,omitempty" bson:"startHour,omitempty"`
	StartMinute   float64 `json:"startMinute,omitempty" bson:"startMinute,omitempty"`
	TimeZone      string  `json:"timeZone,omitempty" bson:"timeZone,omitempty"`
}

type AzureResourceMetadata struct {
	CreatedBy              string `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	CreatedDateTimeUtc     string `json:"createdDateTimeUtc,omitempty" bson:"createdDateTimeUtc,omitempty"`
	LastUpdatedBy          string `json:"lastUpdatedBy,omitempty" bson:"lastUpdatedBy,omitempty"`
	LastUpdatedDateTimeUtc string `json:"lastUpdatedDateTimeUtc,omitempty" bson:"lastUpdatedDateTimeUtc,omitempty"`
	Model                  *struct {
		FilterLocale *struct {
			Value string `json:"value,omitempty" bson:"value,omitempty"`
		} `json:"filterLocale,omitempty" bson:"filterLocale,omitempty"`
		Filters *struct {
			Value struct {
				MsPortalFxTimeRange struct {
					DisplayCache struct {
						Name  string `json:"name,omitempty" bson:"name,omitempty"`
						Value string `json:"value,omitempty" bson:"value,omitempty"`
					} `json:"displayCache,omitempty" bson:"displayCache,omitempty"`
					FilteredPartIds []string `json:"filteredPartIds,omitempty" bson:"filteredPartIds,omitempty"`
					Model           struct {
						Format      string `json:"format,omitempty" bson:"format,omitempty"`
						Granularity string `json:"granularity,omitempty" bson:"granularity,omitempty"`
						Relative    string `json:"relative,omitempty" bson:"relative,omitempty"`
					} `json:"model,omitempty" bson:"model,omitempty"`
				} `json:"MsPortalFx_TimeRange,omitempty" bson:"MsPortalFx_TimeRange,omitempty"`
			} `json:"value,omitempty" bson:"value,omitempty"`
		} `json:"filters,omitempty" bson:"filters,omitempty"`
		TimeRange *struct {
			Type  string `json:"type,omitempty" bson:"type,omitempty"`
			Value struct {
				Relative struct {
					Duration float64 `json:"duration,omitempty" bson:"duration,omitempty"`
					TimeUnit float64 `json:"timeUnit,omitempty" bson:"timeUnit,omitempty"`
				} `json:"relative,omitempty" bson:"relative,omitempty"`
			} `json:"value,omitempty" bson:"value,omitempty"`
		} `json:"timeRange,omitempty" bson:"timeRange,omitempty"`
	} `json:"model,omitempty" bson:"model,omitempty"`
}

type AzureResourceNetwork struct {
	DelegatedSubnetResourceID   string `json:"delegatedSubnetResourceId,omitempty" bson:"delegatedSubnetResourceId,omitempty"`
	PrivateDnsZoneArmResourceID string `json:"privateDnsZoneArmResourceId,omitempty" bson:"privateDnsZoneArmResourceId,omitempty"`
	PublicNetworkAccess         string `json:"publicNetworkAccess,omitempty" bson:"publicNetworkAccess,omitempty"`
}

type AzureResourceNetworkAcls struct {
	Bypass        string `json:"bypass,omitempty" bson:"bypass,omitempty"`
	DefaultAction string `json:"defaultAction,omitempty" bson:"defaultAction,omitempty"`
	IpRules       []struct {
		Action string `json:"action,omitempty" bson:"action,omitempty"`
		Value  string `json:"value,omitempty" bson:"value,omitempty"`
	} `json:"ipRules,omitempty" bson:"ipRules,omitempty"`
	Ipv6Rules           []any  `json:"ipv6Rules,omitempty" bson:"ipv6Rules,omitempty"`
	PublicNetworkAccess string `json:"publicNetworkAccess,omitempty" bson:"publicNetworkAccess,omitempty"`
	ResourceAccessRules []struct {
		ResourceID string `json:"resourceId,omitempty" bson:"resourceId,omitempty"`
		TenantID   string `json:"tenantId,omitempty" bson:"tenantId,omitempty"`
	} `json:"resourceAccessRules,omitempty" bson:"resourceAccessRules,omitempty"`
	VirtualNetworkRules []struct {
		Action                           string `json:"action,omitempty" bson:"action,omitempty"`
		ID                               string `json:"id,omitempty" bson:"id,omitempty"`
		IgnoreMissingVnetServiceEndpoint bool   `json:"ignoreMissingVnetServiceEndpoint,omitempty" bson:"ignoreMissingVnetServiceEndpoint,omitempty"`
		State                            string `json:"state,omitempty" bson:"state,omitempty"`
	} `json:"virtualNetworkRules,omitempty" bson:"virtualNetworkRules,omitempty"`
}

type AzureResourceNetworkProfile struct {
	AccountAccess *struct {
		DefaultAction string `json:"defaultAction,omitempty" bson:"defaultAction,omitempty"`
	} `json:"accountAccess,omitempty" bson:"accountAccess,omitempty"`
	DnsServiceIp        string   `json:"dnsServiceIP,omitempty" bson:"dnsServiceIP,omitempty"`
	IpFamilies          []string `json:"ipFamilies,omitempty" bson:"ipFamilies,omitempty"`
	LoadBalancerProfile *struct {
		BackendPoolType      string `json:"backendPoolType,omitempty" bson:"backendPoolType,omitempty"`
		EffectiveOutboundIPs []struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"effectiveOutboundIPs,omitempty" bson:"effectiveOutboundIPs,omitempty"`
		ManagedOutboundIPs struct {
			Count float64 `json:"count,omitempty" bson:"count,omitempty"`
		} `json:"managedOutboundIPs,omitempty" bson:"managedOutboundIPs,omitempty"`
	} `json:"loadBalancerProfile,omitempty" bson:"loadBalancerProfile,omitempty"`
	LoadBalancerSku   string `json:"loadBalancerSku,omitempty" bson:"loadBalancerSku,omitempty"`
	NetworkDataplane  string `json:"networkDataplane,omitempty" bson:"networkDataplane,omitempty"`
	NetworkInterfaces []struct {
		ID         string `json:"id,omitempty" bson:"id,omitempty"`
		Properties *struct {
			DeleteOption string `json:"deleteOption,omitempty" bson:"deleteOption,omitempty"`
			Primary      bool   `json:"primary,omitempty" bson:"primary,omitempty"`
		} `json:"properties,omitempty" bson:"properties,omitempty"`
	} `json:"networkInterfaces,omitempty" bson:"networkInterfaces,omitempty"`
	NetworkPlugin string   `json:"networkPlugin,omitempty" bson:"networkPlugin,omitempty"`
	NetworkPolicy string   `json:"networkPolicy,omitempty" bson:"networkPolicy,omitempty"`
	OutboundType  string   `json:"outboundType,omitempty" bson:"outboundType,omitempty"`
	ServiceCidr   string   `json:"serviceCidr,omitempty" bson:"serviceCidr,omitempty"`
	ServiceCidrs  []string `json:"serviceCidrs,omitempty" bson:"serviceCidrs,omitempty"`
}

type AzureResourceNetworkRuleSet struct {
	DefaultAction string `json:"defaultAction,omitempty" bson:"defaultAction,omitempty"`
	IpRules       []struct {
		Action string `json:"action,omitempty" bson:"action,omitempty"`
		Value  string `json:"value,omitempty" bson:"value,omitempty"`
	} `json:"ipRules,omitempty" bson:"ipRules,omitempty"`
}

type AzureResourceNotificationSettings struct {
	EmailRecipient     string  `json:"emailRecipient,omitempty" bson:"emailRecipient,omitempty"`
	NotificationLocale string  `json:"notificationLocale,omitempty" bson:"notificationLocale,omitempty"`
	Status             string  `json:"status,omitempty" bson:"status,omitempty"`
	TimeInMinutes      float64 `json:"timeInMinutes,omitempty" bson:"timeInMinutes,omitempty"`
	WebhookURL         string  `json:"webhookUrl,omitempty" bson:"webhookUrl,omitempty"`
}

type AzureResourceOSProfile struct {
	AdminUsername            string `json:"adminUsername,omitempty" bson:"adminUsername,omitempty"`
	AllowExtensionOperations bool   `json:"allowExtensionOperations,omitempty" bson:"allowExtensionOperations,omitempty"`
	ComputerName             string `json:"computerName,omitempty" bson:"computerName,omitempty"`
	LinuxConfiguration       *struct {
		DisablePasswordAuthentication bool `json:"disablePasswordAuthentication,omitempty" bson:"disablePasswordAuthentication,omitempty"`
		EnableVmAgentPlatformUpdates  bool `json:"enableVMAgentPlatformUpdates,omitempty" bson:"enableVMAgentPlatformUpdates,omitempty"`
		PatchSettings                 struct {
			AssessmentMode              string `json:"assessmentMode,omitempty" bson:"assessmentMode,omitempty"`
			AutomaticByPlatformSettings *struct {
				BypassPlatformSafetyChecksOnUserSchedule bool   `json:"bypassPlatformSafetyChecksOnUserSchedule,omitempty" bson:"bypassPlatformSafetyChecksOnUserSchedule,omitempty"`
				RebootSetting                            string `json:"rebootSetting,omitempty" bson:"rebootSetting,omitempty"`
			} `json:"automaticByPlatformSettings,omitempty" bson:"automaticByPlatformSettings,omitempty"`
			PatchMode string `json:"patchMode,omitempty" bson:"patchMode,omitempty"`
		} `json:"patchSettings,omitempty" bson:"patchSettings,omitempty"`
		ProvisionVmAgent bool `json:"provisionVMAgent,omitempty" bson:"provisionVMAgent,omitempty"`
		SSH              *struct {
			PublicKeys []struct {
				KeyData string `json:"keyData,omitempty" bson:"keyData,omitempty"`
				Path    string `json:"path,omitempty" bson:"path,omitempty"`
			} `json:"publicKeys,omitempty" bson:"publicKeys,omitempty"`
		} `json:"ssh,omitempty" bson:"ssh,omitempty"`
	} `json:"linuxConfiguration,omitempty" bson:"linuxConfiguration,omitempty"`
	RequireGuestProvisionSignal bool `json:"requireGuestProvisionSignal,omitempty" bson:"requireGuestProvisionSignal,omitempty"`
	Secrets                     []struct {
		SourceVault struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"sourceVault,omitempty" bson:"sourceVault,omitempty"`
		VaultCertificates []struct {
			CertificateStore string `json:"certificateStore,omitempty" bson:"certificateStore,omitempty"`
			CertificateURL   string `json:"certificateUrl,omitempty" bson:"certificateUrl,omitempty"`
		} `json:"vaultCertificates,omitempty" bson:"vaultCertificates,omitempty"`
	} `json:"secrets,omitempty" bson:"secrets,omitempty"`
	WindowsConfiguration *struct {
		EnableAutomaticUpdates       bool `json:"enableAutomaticUpdates,omitempty" bson:"enableAutomaticUpdates,omitempty"`
		EnableVmAgentPlatformUpdates bool `json:"enableVMAgentPlatformUpdates,omitempty" bson:"enableVMAgentPlatformUpdates,omitempty"`
		PatchSettings                struct {
			AssessmentMode              string `json:"assessmentMode,omitempty" bson:"assessmentMode,omitempty"`
			AutomaticByPlatformSettings *struct {
				BypassPlatformSafetyChecksOnUserSchedule bool   `json:"bypassPlatformSafetyChecksOnUserSchedule,omitempty" bson:"bypassPlatformSafetyChecksOnUserSchedule,omitempty"`
				RebootSetting                            string `json:"rebootSetting,omitempty" bson:"rebootSetting,omitempty"`
			} `json:"automaticByPlatformSettings,omitempty" bson:"automaticByPlatformSettings,omitempty"`
			EnableHotpatching bool   `json:"enableHotpatching,omitempty" bson:"enableHotpatching,omitempty"`
			PatchMode         string `json:"patchMode,omitempty" bson:"patchMode,omitempty"`
		} `json:"patchSettings,omitempty" bson:"patchSettings,omitempty"`
		ProvisionVmAgent bool   `json:"provisionVMAgent,omitempty" bson:"provisionVMAgent,omitempty"`
		TimeZone         string `json:"timeZone,omitempty" bson:"timeZone,omitempty"`
	} `json:"windowsConfiguration,omitempty" bson:"windowsConfiguration,omitempty"`
}

type AzureResourceOutboundRules struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		AllocatedOutboundPorts float64 `json:"allocatedOutboundPorts,omitempty" bson:"allocatedOutboundPorts,omitempty"`
		AllocationPolicy       struct {
			OnDemandAllocation bool   `json:"onDemandAllocation,omitempty" bson:"onDemandAllocation,omitempty"`
			PortReuse          string `json:"portReuse,omitempty" bson:"portReuse,omitempty"`
		} `json:"allocationPolicy,omitempty" bson:"allocationPolicy,omitempty"`
		BackendAddressPool struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"backendAddressPool,omitempty" bson:"backendAddressPool,omitempty"`
		EnableTcpReset           bool `json:"enableTcpReset,omitempty" bson:"enableTcpReset,omitempty"`
		FrontendIpConfigurations []struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"frontendIPConfigurations,omitempty" bson:"frontendIPConfigurations,omitempty"`
		IdleTimeoutInMinutes float64 `json:"idleTimeoutInMinutes,omitempty" bson:"idleTimeoutInMinutes,omitempty"`
		Protocol             string  `json:"protocol,omitempty" bson:"protocol,omitempty"`
		ProvisioningState    string  `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceP2SConnectionConfiguration struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		EnableInternetSecurity bool   `json:"enableInternetSecurity,omitempty" bson:"enableInternetSecurity,omitempty"`
		ProvisioningState      string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
		RoutingConfiguration   struct {
			AssociatedRouteTable struct {
				ID string `json:"id,omitempty" bson:"id,omitempty"`
			} `json:"associatedRouteTable,omitempty" bson:"associatedRouteTable,omitempty"`
			PropagatedRouteTables struct {
				Ids []struct {
					ID string `json:"id,omitempty" bson:"id,omitempty"`
				} `json:"ids,omitempty" bson:"ids,omitempty"`
				Labels []string `json:"labels,omitempty" bson:"labels,omitempty"`
			} `json:"propagatedRouteTables,omitempty" bson:"propagatedRouteTables,omitempty"`
		} `json:"routingConfiguration,omitempty" bson:"routingConfiguration,omitempty"`
		VpnClientAddressPool struct {
			AddressPrefixes []string `json:"addressPrefixes,omitempty" bson:"addressPrefixes,omitempty"`
		} `json:"vpnClientAddressPool,omitempty" bson:"vpnClientAddressPool,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceParameterConnection struct {
	ConnectionID         string `json:"connectionId,omitempty" bson:"connectionId,omitempty"`
	ConnectionName       string `json:"connectionName,omitempty" bson:"connectionName,omitempty"`
	ConnectionProperties struct {
		Authentication struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"authentication,omitempty" bson:"authentication,omitempty"`
	} `json:"connectionProperties,omitempty" bson:"connectionProperties,omitempty"`
	ID string `json:"id,omitempty" bson:"id,omitempty"`
}

type AzureResourceParameter struct {
	DefaultValue any     `json:"defaultValue,omitempty" bson:"defaultValue,omitempty"`
	IsMandatory  bool    `json:"isMandatory,omitempty" bson:"isMandatory,omitempty"`
	Position     float64 `json:"position,omitempty" bson:"position,omitempty"`
	Type         string  `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceParameters struct {
	Action                   *AzureResourceParameter `json:"Action,omitempty" bson:"Action,omitempty"`
	Alert                    *AzureResourceParameter `json:"Alert,omitempty" bson:"Alert,omitempty"`
	AlertsVariable           *AzureResourceParameter `json:"alertsVariable,omitempty" bson:"alertsVariable,omitempty"`
	AutomationAccountName    *AzureResourceParameter `json:"AutomationAccountName,omitempty" bson:"AutomationAccountName,omitempty"`
	AzureConnectionAssetName *AzureResourceParameter `json:"AzureConnectionAssetName,omitempty" bson:"AzureConnectionAssetName,omitempty"`
	ComplianceTableName      *AzureResourceParameter `json:"ComplianceTableName,omitempty" bson:"ComplianceTableName,omitempty"`
	Connections              *struct {
		Value map[string]AzureResourceParameterConnection `json:"value,omitempty" bson:"value,omitempty"`
	} `json:"$connections,omitempty" bson:"_connections,omitempty"`
	Days                       *AzureResourceParameter `json:"days,omitempty" bson:"days,omitempty"`
	Details                    *AzureResourceParameter `json:"Details,omitempty" bson:"Details,omitempty"`
	DomainCredentialName       *AzureResourceParameter `json:"DomainCredentialName,omitempty" bson:"DomainCredentialName,omitempty"`
	DomainJoinCred             *AzureResourceParameter `json:"domainJoinCred,omitempty" bson:"domainJoinCred,omitempty"`
	DomainJoinUser             *AzureResourceParameter `json:"domainJoinUser,omitempty" bson:"domainJoinUser,omitempty"`
	DomainName                 *AzureResourceParameter `json:"DomainName,omitempty" bson:"DomainName,omitempty"`
	EmergencySendTo            *AzureResourceParameter `json:"EmergencySendTo,omitempty" bson:"EmergencySendTo,omitempty"`
	GetIgnoreAlertsOnSpoke     *AzureResourceParameter `json:"getIgnoreAlertsOnSpoke,omitempty" bson:"getIgnoreAlertsOnSpoke,omitempty"`
	GetSpokeAlertConfiguration *AzureResourceParameter `json:"getSpokeAlertConfiguration,omitempty" bson:"getSpokeAlertConfiguration,omitempty"`
	Hostname                   *AzureResourceParameter `json:"hostname,omitempty" bson:"hostname,omitempty"`
	HostName                   *AzureResourceParameter `json:"hostName,omitempty" bson:"hostName,omitempty"`
	LdapAccountCred            *AzureResourceParameter `json:"ldapAccountCred,omitempty" bson:"ldapAccountCred,omitempty"`
	LdapAccountUser            *AzureResourceParameter `json:"ldapAccountUser,omitempty" bson:"ldapAccountUser,omitempty"`
	NetBiosName                *AzureResourceParameter `json:"NetBiosName,omitempty" bson:"NetBiosName,omitempty"`
	OrgName                    *AzureResourceParameter `json:"OrgName,omitempty" bson:"OrgName,omitempty"`
	Password                   *AzureResourceParameter `json:"password,omitempty" bson:"password,omitempty"`
	PasswordOther              *AzureResourceParameter `json:"Password,omitempty" bson:"Password,omitempty"`
	Portagw                    *AzureResourceParameter `json:"portagw,omitempty" bson:"portagw,omitempty"`
	Portaip                    *AzureResourceParameter `json:"portaip,omitempty" bson:"portaip,omitempty"`
	Regions                    *AzureResourceParameter `json:"Regions,omitempty" bson:"Regions,omitempty"`
	ResourceGroup              *AzureResourceParameter `json:"ResourceGroup,omitempty" bson:"ResourceGroup,omitempty"`
	ResourceGroupName          *AzureResourceParameter `json:"ResourceGroupName,omitempty" bson:"ResourceGroupName,omitempty"`
	RestartCount               *AzureResourceParameter `json:"RestartCount,omitempty" bson:"RestartCount,omitempty"`
	RetryCount                 *AzureResourceParameter `json:"RetryCount,omitempty" bson:"RetryCount,omitempty"`
	RetryIntervalSec           *AzureResourceParameter `json:"RetryIntervalSec,omitempty" bson:"RetryIntervalSec,omitempty"`
	SafeModeCredentialName     *AzureResourceParameter `json:"SafeModeCredentialName,omitempty" bson:"SafeModeCredentialName,omitempty"`
	SendTo                     *AzureResourceParameter `json:"sendTo,omitempty" bson:"sendTo,omitempty"`
	SendToOther                *AzureResourceParameter `json:"SendTo,omitempty" bson:"SendTo,omitempty"`
	Sshport                    *AzureResourceParameter `json:"sshport,omitempty" bson:"sshport,omitempty"`
	Subject                    *AzureResourceParameter `json:"subject,omitempty" bson:"subject,omitempty"`
	SubjectOther               *AzureResourceParameter `json:"Subject,omitempty" bson:"Subject,omitempty"`
	SubscriptionID             *AzureResourceParameter `json:"SubscriptionId,omitempty" bson:"SubscriptionId,omitempty"`
	TriggerRunbook             *AzureResourceParameter `json:"TriggerRunbook,omitempty" bson:"TriggerRunbook,omitempty"`
	VerbosePreference          *AzureResourceParameter `json:"VerbosePreference,omitempty" bson:"VerbosePreference,omitempty"`
	WaitTimeout                *AzureResourceParameter `json:"WaitTimeout,omitempty" bson:"WaitTimeout,omitempty"`
	WebhookData                *AzureResourceParameter `json:"WebhookData,omitempty" bson:"WebhookData,omitempty"`
	WeeklyReport               *AzureResourceParameter `json:"WeeklyReport,omitempty" bson:"WeeklyReport,omitempty"`
}

type AzureResourcePolicies struct {
	AzureAdAuthenticationAsArmPolicy struct {
		Status string `json:"status,omitempty" bson:"status,omitempty"`
	} `json:"azureADAuthenticationAsArmPolicy,omitempty" bson:"azureADAuthenticationAsArmPolicy,omitempty"`
	ExportPolicy struct {
		Status string `json:"status,omitempty" bson:"status,omitempty"`
	} `json:"exportPolicy,omitempty" bson:"exportPolicy,omitempty"`
	QuarantinePolicy struct {
		Status string `json:"status,omitempty" bson:"status,omitempty"`
	} `json:"quarantinePolicy,omitempty" bson:"quarantinePolicy,omitempty"`
	RetentionPolicy struct {
		Days            float64 `json:"days,omitempty" bson:"days,omitempty"`
		LastUpdatedTime string  `json:"lastUpdatedTime,omitempty" bson:"lastUpdatedTime,omitempty"`
		Status          string  `json:"status,omitempty" bson:"status,omitempty"`
	} `json:"retentionPolicy,omitempty" bson:"retentionPolicy,omitempty"`
	SoftDeletePolicy struct {
		LastUpdatedTime string  `json:"lastUpdatedTime,omitempty" bson:"lastUpdatedTime,omitempty"`
		RetentionDays   float64 `json:"retentionDays,omitempty" bson:"retentionDays,omitempty"`
		Status          string  `json:"status,omitempty" bson:"status,omitempty"`
	} `json:"softDeletePolicy,omitempty" bson:"softDeletePolicy,omitempty"`
	TrustPolicy struct {
		Status string `json:"status,omitempty" bson:"status,omitempty"`
		Type   string `json:"type,omitempty" bson:"type,omitempty"`
	} `json:"trustPolicy,omitempty" bson:"trustPolicy,omitempty"`
}

type AzureResourcePeerings struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		AzureAsn               float64 `json:"azureASN,omitempty" bson:"azureASN,omitempty"`
		Connections            []any   `json:"connections,omitempty" bson:"connections,omitempty"`
		GatewayManagerEtag     string  `json:"gatewayManagerEtag,omitempty" bson:"gatewayManagerEtag,omitempty"`
		LastModifiedBy         string  `json:"lastModifiedBy,omitempty" bson:"lastModifiedBy,omitempty"`
		MicrosoftPeeringConfig *struct {
			AdvertisedCommunities         []any   `json:"advertisedCommunities,omitempty" bson:"advertisedCommunities,omitempty"`
			AdvertisedPublicPrefixes      []any   `json:"advertisedPublicPrefixes,omitempty" bson:"advertisedPublicPrefixes,omitempty"`
			AdvertisedPublicPrefixesState string  `json:"advertisedPublicPrefixesState,omitempty" bson:"advertisedPublicPrefixesState,omitempty"`
			CustomerAsn                   float64 `json:"customerASN,omitempty" bson:"customerASN,omitempty"`
			LegacyMode                    float64 `json:"legacyMode,omitempty" bson:"legacyMode,omitempty"`
			RoutingRegistryName           string  `json:"routingRegistryName,omitempty" bson:"routingRegistryName,omitempty"`
		} `json:"microsoftPeeringConfig,omitempty" bson:"microsoftPeeringConfig,omitempty"`
		PeerAsn                    float64 `json:"peerASN,omitempty" bson:"peerASN,omitempty"`
		PeeredConnections          []any   `json:"peeredConnections,omitempty" bson:"peeredConnections,omitempty"`
		PeeringType                string  `json:"peeringType,omitempty" bson:"peeringType,omitempty"`
		PrimaryAzurePort           string  `json:"primaryAzurePort,omitempty" bson:"primaryAzurePort,omitempty"`
		PrimaryPeerAddressPrefix   string  `json:"primaryPeerAddressPrefix,omitempty" bson:"primaryPeerAddressPrefix,omitempty"`
		ProvisioningState          string  `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
		SecondaryAzurePort         string  `json:"secondaryAzurePort,omitempty" bson:"secondaryAzurePort,omitempty"`
		SecondaryPeerAddressPrefix string  `json:"secondaryPeerAddressPrefix,omitempty" bson:"secondaryPeerAddressPrefix,omitempty"`
		State                      string  `json:"state,omitempty" bson:"state,omitempty"`
		VlanID                     float64 `json:"vlanId,omitempty" bson:"vlanId,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourcePrimaryEndpoints struct {
	Blob              string `json:"blob,omitempty" bson:"blob,omitempty"`
	Dfs               string `json:"dfs,omitempty" bson:"dfs,omitempty"`
	File              string `json:"file,omitempty" bson:"file,omitempty"`
	InternetEndpoints *struct {
		Blob string `json:"blob,omitempty" bson:"blob,omitempty"`
		Dfs  string `json:"dfs,omitempty" bson:"dfs,omitempty"`
		File string `json:"file,omitempty" bson:"file,omitempty"`
		Web  string `json:"web,omitempty" bson:"web,omitempty"`
	} `json:"internetEndpoints,omitempty" bson:"internetEndpoints,omitempty"`
	MicrosoftEndpoints *struct {
		Blob  string `json:"blob,omitempty" bson:"blob,omitempty"`
		Dfs   string `json:"dfs,omitempty" bson:"dfs,omitempty"`
		File  string `json:"file,omitempty" bson:"file,omitempty"`
		Queue string `json:"queue,omitempty" bson:"queue,omitempty"`
		Table string `json:"table,omitempty" bson:"table,omitempty"`
		Web   string `json:"web,omitempty" bson:"web,omitempty"`
	} `json:"microsoftEndpoints,omitempty" bson:"microsoftEndpoints,omitempty"`
	Queue string `json:"queue,omitempty" bson:"queue,omitempty"`
	Table string `json:"table,omitempty" bson:"table,omitempty"`
	Web   string `json:"web,omitempty" bson:"web,omitempty"`
}

type AzureResourcePrivateEndpointConnections struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Location   string `json:"location,omitempty" bson:"location,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		GroupIds        []string `json:"groupIds,omitempty" bson:"groupIds,omitempty"`
		IpAddresses     []string `json:"ipAddresses,omitempty" bson:"ipAddresses,omitempty"`
		PrivateEndpoint struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"privateEndpoint,omitempty" bson:"privateEndpoint,omitempty"`
		PrivateLinkServiceConnectionState struct {
			ActionRequired  string `json:"actionRequired,omitempty" bson:"actionRequired,omitempty"`
			ActionsRequired string `json:"actionsRequired,omitempty" bson:"actionsRequired,omitempty"`
			Description     string `json:"description,omitempty" bson:"description,omitempty"`
			Status          string `json:"status,omitempty" bson:"status,omitempty"`
		} `json:"privateLinkServiceConnectionState,omitempty" bson:"privateLinkServiceConnectionState,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourcePrivateLinkResources struct {
	GroupID         string   `json:"groupId,omitempty" bson:"groupId,omitempty"`
	ID              string   `json:"id,omitempty" bson:"id,omitempty"`
	Name            string   `json:"name,omitempty" bson:"name,omitempty"`
	RequiredMembers []string `json:"requiredMembers,omitempty" bson:"requiredMembers,omitempty"`
	Type            string   `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourcePrivateLinkServiceConnection struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		GroupIds                          []string `json:"groupIds,omitempty" bson:"groupIds,omitempty"`
		PrivateLinkServiceConnectionState struct {
			ActionsRequired string `json:"actionsRequired,omitempty" bson:"actionsRequired,omitempty"`
			Description     string `json:"description,omitempty" bson:"description,omitempty"`
			Status          string `json:"status,omitempty" bson:"status,omitempty"`
		} `json:"privateLinkServiceConnectionState,omitempty" bson:"privateLinkServiceConnectionState,omitempty"`
		PrivateLinkServiceID               string `json:"privateLinkServiceId,omitempty" bson:"privateLinkServiceId,omitempty"`
		ProvisioningState                  string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
		RequestMessage                     string `json:"requestMessage,omitempty" bson:"requestMessage,omitempty"`
		ResolvedPrivateLinkServiceLocation string `json:"resolvedPrivateLinkServiceLocation,omitempty" bson:"resolvedPrivateLinkServiceLocation,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceProbe struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		IntervalInSeconds  float64 `json:"intervalInSeconds,omitempty" bson:"intervalInSeconds,omitempty"`
		LoadBalancingRules []struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"loadBalancingRules,omitempty" bson:"loadBalancingRules,omitempty"`
		NumberOfProbes    float64 `json:"numberOfProbes,omitempty" bson:"numberOfProbes,omitempty"`
		Port              float64 `json:"port,omitempty" bson:"port,omitempty"`
		ProbeThreshold    float64 `json:"probeThreshold,omitempty" bson:"probeThreshold,omitempty"`
		Protocol          string  `json:"protocol,omitempty" bson:"protocol,omitempty"`
		ProvisioningState string  `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
		RequestPath       string  `json:"requestPath,omitempty" bson:"requestPath,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourcePublishingProfile struct {
	ExcludeFromLatest bool    `json:"excludeFromLatest,omitempty" bson:"excludeFromLatest,omitempty"`
	PublishedDate     string  `json:"publishedDate,omitempty" bson:"publishedDate,omitempty"`
	ReplicaCount      float64 `json:"replicaCount,omitempty" bson:"replicaCount,omitempty"`
	Source            struct {
		ManagedImage struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"managedImage,omitempty" bson:"managedImage,omitempty"`
	} `json:"source,omitempty" bson:"source,omitempty"`
	StorageAccountType string `json:"storageAccountType,omitempty" bson:"storageAccountType,omitempty"`
	TargetRegions      []struct {
		Name                 string  `json:"name,omitempty" bson:"name,omitempty"`
		RegionalReplicaCount float64 `json:"regionalReplicaCount,omitempty" bson:"regionalReplicaCount,omitempty"`
		StorageAccountType   string  `json:"storageAccountType,omitempty" bson:"storageAccountType,omitempty"`
	} `json:"targetRegions,omitempty" bson:"targetRegions,omitempty"`
}

type AzureResourcePurchasePlan struct {
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	Product   string `json:"product,omitempty" bson:"product,omitempty"`
	Publisher string `json:"publisher,omitempty" bson:"publisher,omitempty"`
}

type AzureResourceReadLocation struct {
	DocumentEndpoint  string  `json:"documentEndpoint,omitempty" bson:"documentEndpoint,omitempty"`
	FailoverPriority  float64 `json:"failoverPriority,omitempty" bson:"failoverPriority,omitempty"`
	ID                string  `json:"id,omitempty" bson:"id,omitempty"`
	IsZoneRedundant   bool    `json:"isZoneRedundant,omitempty" bson:"isZoneRedundant,omitempty"`
	LocationName      string  `json:"locationName,omitempty" bson:"locationName,omitempty"`
	ProvisioningState string  `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
}

type AzureResourceRecommended struct {
	Memory struct {
		Max float64 `json:"max,omitempty" bson:"max,omitempty"`
		Min float64 `json:"min,omitempty" bson:"min,omitempty"`
	} `json:"memory,omitempty" bson:"memory,omitempty"`
	VCpUs struct {
		Max float64 `json:"max,omitempty" bson:"max,omitempty"`
		Min float64 `json:"min,omitempty" bson:"min,omitempty"`
	} `json:"vCPUs,omitempty" bson:"vCPUs,omitempty"`
}

type AzureResourceRoute struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		AddressPrefix     string `json:"addressPrefix,omitempty" bson:"addressPrefix,omitempty"`
		HasBgpOverride    bool   `json:"hasBgpOverride,omitempty" bson:"hasBgpOverride,omitempty"`
		NextHopIpAddress  string `json:"nextHopIpAddress,omitempty" bson:"nextHopIpAddress,omitempty"`
		NextHopType       string `json:"nextHopType,omitempty" bson:"nextHopType,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceRoutingPreference struct {
	PublishInternetEndpoints  bool   `json:"publishInternetEndpoints,omitempty" bson:"publishInternetEndpoints,omitempty"`
	PublishMicrosoftEndpoints bool   `json:"publishMicrosoftEndpoints,omitempty" bson:"publishMicrosoftEndpoints,omitempty"`
	RoutingChoice             string `json:"routingChoice,omitempty" bson:"routingChoice,omitempty"`
}

type AzureResourceSecurityProfile struct {
	Defender *struct {
		LogAnalyticsWorkspaceResourceID string `json:"logAnalyticsWorkspaceResourceId,omitempty" bson:"logAnalyticsWorkspaceResourceId,omitempty"`
		SecurityMonitoring              struct {
			Enabled bool `json:"enabled,omitempty" bson:"enabled,omitempty"`
		} `json:"securityMonitoring,omitempty" bson:"securityMonitoring,omitempty"`
	} `json:"defender,omitempty" bson:"defender,omitempty"`
	SecurityType string `json:"securityType,omitempty" bson:"securityType,omitempty"`
	UefiSettings *struct {
		SecureBootEnabled bool `json:"secureBootEnabled,omitempty" bson:"secureBootEnabled,omitempty"`
		VTpmEnabled       bool `json:"vTpmEnabled,omitempty" bson:"vTpmEnabled,omitempty"`
	} `json:"uefiSettings,omitempty" bson:"uefiSettings,omitempty"`
}

type AzureResourceSecurityRule struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		Access                     string   `json:"access,omitempty" bson:"access,omitempty"`
		Description                string   `json:"description,omitempty" bson:"description,omitempty"`
		DestinationAddressPrefix   string   `json:"destinationAddressPrefix,omitempty" bson:"destinationAddressPrefix,omitempty"`
		DestinationAddressPrefixes []string `json:"destinationAddressPrefixes,omitempty" bson:"destinationAddressPrefixes,omitempty"`
		DestinationPortRange       string   `json:"destinationPortRange,omitempty" bson:"destinationPortRange,omitempty"`
		DestinationPortRanges      []string `json:"destinationPortRanges,omitempty" bson:"destinationPortRanges,omitempty"`
		Direction                  string   `json:"direction,omitempty" bson:"direction,omitempty"`
		Priority                   float64  `json:"priority,omitempty" bson:"priority,omitempty"`
		Protocol                   string   `json:"protocol,omitempty" bson:"protocol,omitempty"`
		ProvisioningState          string   `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
		SourceAddressPrefix        string   `json:"sourceAddressPrefix,omitempty" bson:"sourceAddressPrefix,omitempty"`
		SourceAddressPrefixes      []string `json:"sourceAddressPrefixes,omitempty" bson:"sourceAddressPrefixes,omitempty"`
		SourcePortRange            string   `json:"sourcePortRange,omitempty" bson:"sourcePortRange,omitempty"`
		SourcePortRanges           []string `json:"sourcePortRanges,omitempty" bson:"sourcePortRanges,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceSecuritySettings struct {
	ImmutabilitySettings   any    `json:"immutabilitySettings,omitempty" bson:"immutabilitySettings,omitempty"`
	MultiUserAuthorization string `json:"multiUserAuthorization,omitempty" bson:"multiUserAuthorization,omitempty"`
	SoftDeleteSettings     struct {
		EnhancedSecurityState           string  `json:"enhancedSecurityState,omitempty" bson:"enhancedSecurityState,omitempty"`
		RetentionDurationInDays         float64 `json:"retentionDurationInDays,omitempty" bson:"retentionDurationInDays,omitempty"`
		SoftDeleteRetentionPeriodInDays float64 `json:"softDeleteRetentionPeriodInDays,omitempty" bson:"softDeleteRetentionPeriodInDays,omitempty"`
		SoftDeleteState                 string  `json:"softDeleteState,omitempty" bson:"softDeleteState,omitempty"`
		State                           string  `json:"state,omitempty" bson:"state,omitempty"`
	} `json:"softDeleteSettings,omitempty" bson:"softDeleteSettings,omitempty"`
}

type AzureResourceSettingsAttestationConfig struct {
	AscSettings struct {
		AscReportingEndpoint  string `json:"ascReportingEndpoint,omitempty" bson:"ascReportingEndpoint,omitempty"`
		AscReportingFrequency string `json:"ascReportingFrequency,omitempty" bson:"ascReportingFrequency,omitempty"`
	} `json:"AscSettings,omitempty" bson:"AscSettings,omitempty"`
	MaaSettings struct {
		MaaEndpoint   string `json:"maaEndpoint,omitempty" bson:"maaEndpoint,omitempty"`
		MaaTenantName string `json:"maaTenantName,omitempty" bson:"maaTenantName,omitempty"`
	} `json:"MaaSettings,omitempty" bson:"MaaSettings,omitempty"`
	DisableAlerts  string `json:"disableAlerts,omitempty" bson:"disableAlerts,omitempty"`
	UseCustomToken string `json:"useCustomToken,omitempty" bson:"useCustomToken,omitempty"`
}

type AzureResourceSettingsAutoPatching struct {
	AdditionalVmPatch             string `json:"AdditionalVmPatch,omitempty" bson:"AdditionalVmPatch,omitempty"`
	DayOfWeek                     string `json:"DayOfWeek,omitempty" bson:"DayOfWeek,omitempty"`
	Enable                        bool   `json:"Enable,omitempty" bson:"Enable,omitempty"`
	MaintenanceWindowDuration     string `json:"MaintenanceWindowDuration,omitempty" bson:"MaintenanceWindowDuration,omitempty"`
	MaintenanceWindowStartingHour string `json:"MaintenanceWindowStartingHour,omitempty" bson:"MaintenanceWindowStartingHour,omitempty"`
	PatchCategory                 string `json:"PatchCategory,omitempty" bson:"PatchCategory,omitempty"`
}

type AzureResourceSettingsServerConfigurationsManagement struct {
	AdditionalFeaturesServerConfigurations struct {
		BackupPermissionsForAzureBackupSvc bool `json:"BackupPermissionsForAzureBackupSvc,omitempty" bson:"BackupPermissionsForAzureBackupSvc,omitempty"`
		IsRServicesEnabled                 bool `json:"IsRServicesEnabled,omitempty" bson:"IsRServicesEnabled,omitempty"`
	} `json:"AdditionalFeaturesServerConfigurations,omitempty" bson:"AdditionalFeaturesServerConfigurations,omitempty"`
	SQLConnectivityUpdateSettings *struct {
		ConnectivityType string `json:"ConnectivityType,omitempty" bson:"ConnectivityType,omitempty"`
		Port             string `json:"Port,omitempty" bson:"Port,omitempty"`
	} `json:"SQLConnectivityUpdateSettings,omitempty" bson:"SQLConnectivityUpdateSettings,omitempty"`
	SQLInstanceSettings *struct {
		Collation                          string  `json:"Collation,omitempty" bson:"Collation,omitempty"`
		IsIfiEnabled                       bool    `json:"IsIFIEnabled,omitempty" bson:"IsIFIEnabled,omitempty"`
		IsLpimEnabled                      bool    `json:"IsLPIMEnabled,omitempty" bson:"IsLPIMEnabled,omitempty"`
		IsOptimizeForAdHocWorkloadsEnabled bool    `json:"IsOptimizeForAdHocWorkloadsEnabled,omitempty" bson:"IsOptimizeForAdHocWorkloadsEnabled,omitempty"`
		MaxDop                             float64 `json:"MaxDop,omitempty" bson:"MaxDop,omitempty"`
		MaxServerMemoryMb                  float64 `json:"MaxServerMemoryMB,omitempty" bson:"MaxServerMemoryMB,omitempty"`
		MinServerMemoryMb                  float64 `json:"MinServerMemoryMB,omitempty" bson:"MinServerMemoryMB,omitempty"`
	} `json:"SQLInstanceSettings,omitempty" bson:"SQLInstanceSettings,omitempty"`
	SQLStorageUpdateSettingsV2 *struct {
		DiskConfigurationType string `json:"DiskConfigurationType,omitempty" bson:"DiskConfigurationType,omitempty"`
		SQLDataSettings       struct {
			DefaultFilePath string    `json:"DefaultFilePath,omitempty" bson:"DefaultFilePath,omitempty"`
			LuNs            []float64 `json:"LUNs,omitempty" bson:"LUNs,omitempty"`
		} `json:"SQLDataSettings,omitempty" bson:"SQLDataSettings,omitempty"`
		SQLLogSettings struct {
			DefaultFilePath string    `json:"DefaultFilePath,omitempty" bson:"DefaultFilePath,omitempty"`
			LuNs            []float64 `json:"LUNs,omitempty" bson:"LUNs,omitempty"`
		} `json:"SQLLogSettings,omitempty" bson:"SQLLogSettings,omitempty"`
		SQLSystemDBOnDataDisk bool `json:"SQLSystemDbOnDataDisk,omitempty" bson:"SQLSystemDbOnDataDisk,omitempty"`
		SQLTempDBSettings     struct {
			DataFileCount   string `json:"DataFileCount,omitempty" bson:"DataFileCount,omitempty"`
			DataFileSize    string `json:"DataFileSize,omitempty" bson:"DataFileSize,omitempty"`
			DataGrowth      string `json:"DataGrowth,omitempty" bson:"DataGrowth,omitempty"`
			DefaultFilePath string `json:"DefaultFilePath,omitempty" bson:"DefaultFilePath,omitempty"`
			LogFileSize     string `json:"LogFileSize,omitempty" bson:"LogFileSize,omitempty"`
			LogGrowth       string `json:"LogGrowth,omitempty" bson:"LogGrowth,omitempty"`
		} `json:"SQLTempDbSettings,omitempty" bson:"SQLTempDbSettings,omitempty"`
	} `json:"SQLStorageUpdateSettingsV2,omitempty" bson:"SQLStorageUpdateSettingsV2,omitempty"`
	SQLWorkloadTypeUpdateSettings *struct {
		SQLWorkloadType float64 `json:"SQLWorkloadType,omitempty" bson:"SQLWorkloadType,omitempty"`
	} `json:"SQLWorkloadTypeUpdateSettings,omitempty" bson:"SQLWorkloadTypeUpdateSettings,omitempty"`
}

type AzureResourceSettingsWadCfg struct {
	DiagnosticMonitorConfiguration struct {
		DiagnosticInfrastructureLogs struct {
			ScheduledTransferLogLevelFilter string `json:"scheduledTransferLogLevelFilter,omitempty" bson:"scheduledTransferLogLevelFilter,omitempty"`
			ScheduledTransferPeriod         string `json:"scheduledTransferPeriod,omitempty" bson:"scheduledTransferPeriod,omitempty"`
		} `json:"DiagnosticInfrastructureLogs,omitempty" bson:"DiagnosticInfrastructureLogs,omitempty"`
		Directories *struct {
			ScheduledTransferPeriod string `json:"scheduledTransferPeriod,omitempty" bson:"scheduledTransferPeriod,omitempty"`
		} `json:"Directories,omitempty" bson:"Directories,omitempty"`
		Metrics struct {
			MetricAggregation []struct {
				ScheduledTransferPeriod string `json:"scheduledTransferPeriod,omitempty" bson:"scheduledTransferPeriod,omitempty"`
			} `json:"MetricAggregation,omitempty" bson:"MetricAggregation,omitempty"`
			ResourceID string `json:"resourceId,omitempty" bson:"resourceId,omitempty"`
		} `json:"Metrics,omitempty" bson:"Metrics,omitempty"`
		PerformanceCounters struct {
			PerformanceCounterConfiguration []struct {
				CounterSpecifier string `json:"counterSpecifier,omitempty" bson:"counterSpecifier,omitempty"`
				SampleRate       string `json:"sampleRate,omitempty" bson:"sampleRate,omitempty"`
				Unit             string `json:"unit,omitempty" bson:"unit,omitempty"`
			} `json:"PerformanceCounterConfiguration,omitempty" bson:"PerformanceCounterConfiguration,omitempty"`
			ScheduledTransferPeriod string `json:"scheduledTransferPeriod,omitempty" bson:"scheduledTransferPeriod,omitempty"`
		} `json:"PerformanceCounters,omitempty" bson:"PerformanceCounters,omitempty"`
		WindowsEventLog struct {
			DataSource []struct {
				Name string `json:"name,omitempty" bson:"name,omitempty"`
			} `json:"DataSource,omitempty" bson:"DataSource,omitempty"`
			ScheduledTransferPeriod string `json:"scheduledTransferPeriod,omitempty" bson:"scheduledTransferPeriod,omitempty"`
		} `json:"WindowsEventLog,omitempty" bson:"WindowsEventLog,omitempty"`
		OverallQuotaInMb float64 `json:"overallQuotaInMB,omitempty" bson:"overallQuotaInMB,omitempty"`
	} `json:"DiagnosticMonitorConfiguration,omitempty" bson:"DiagnosticMonitorConfiguration,omitempty"`
}

type AzureResourceSettingsConfigurationArguments struct {
	ActionAfterReboot              string  `json:"ActionAfterReboot,omitempty" bson:"ActionAfterReboot,omitempty"`
	AllowModuleOverwrite           bool    `json:"AllowModuleOverwrite,omitempty" bson:"AllowModuleOverwrite,omitempty"`
	ConfigurationMode              string  `json:"ConfigurationMode,omitempty" bson:"ConfigurationMode,omitempty"`
	ConfigurationModeFrequencyMins float64 `json:"ConfigurationModeFrequencyMins,omitempty" bson:"ConfigurationModeFrequencyMins,omitempty"`
	NodeConfigurationName          string  `json:"NodeConfigurationName,omitempty" bson:"NodeConfigurationName,omitempty"`
	RebootNodeIfNeeded             bool    `json:"RebootNodeIfNeeded,omitempty" bson:"RebootNodeIfNeeded,omitempty"`
	RefreshFrequencyMins           float64 `json:"RefreshFrequencyMins,omitempty" bson:"RefreshFrequencyMins,omitempty"`
	RegistrationURL                string  `json:"RegistrationUrl,omitempty" bson:"RegistrationUrl,omitempty"`
}

type AzureResourceSettings struct {
	AadClientCertThumbprint string                                  `json:"AADClientCertThumbprint,omitempty" bson:"AADClientCertThumbprint,omitempty"`
	AadClientID             string                                  `json:"AADClientID,omitempty" bson:"AADClientID,omitempty"`
	AntimalwareEnabled      any                                     `json:"AntimalwareEnabled,omitempty" bson:"AntimalwareEnabled,omitempty"`
	AttestationConfig       *AzureResourceSettingsAttestationConfig `json:"AttestationConfig,omitempty" bson:"AttestationConfig,omitempty"`
	AutoPatchingSettings    *AzureResourceSettingsAutoPatching      `json:"AutoPatchingSettings,omitempty" bson:"AutoPatchingSettings,omitempty"`
	DeploymentTokenSettings *struct {
		DeploymentToken any `json:"DeploymentToken,omitempty" bson:"DeploymentToken,omitempty"`
	} `json:"DeploymentTokenSettings,omitempty" bson:"DeploymentTokenSettings,omitempty"`
	EncryptionOperation string `json:"EncryptionOperation,omitempty" bson:"EncryptionOperation,omitempty"`
	Exclusions          *struct {
		Extensions string `json:"Extensions,omitempty" bson:"Extensions,omitempty"`
		Paths      string `json:"Paths,omitempty" bson:"Paths,omitempty"`
		Processes  string `json:"Processes,omitempty" bson:"Processes,omitempty"`
	} `json:"Exclusions,omitempty" bson:"Exclusions,omitempty"`
	KekVaultResourceID         string `json:"KekVaultResourceId,omitempty" bson:"KekVaultResourceId,omitempty"`
	KeyEncryptionAlgorithm     string `json:"KeyEncryptionAlgorithm,omitempty" bson:"KeyEncryptionAlgorithm,omitempty"`
	KeyEncryptionKeyURL        string `json:"KeyEncryptionKeyURL,omitempty" bson:"KeyEncryptionKeyURL,omitempty"`
	KeyVaultCredentialSettings *struct {
		Enable bool `json:"Enable,omitempty" bson:"Enable,omitempty"`
	} `json:"KeyVaultCredentialSettings,omitempty" bson:"KeyVaultCredentialSettings,omitempty"`
	KeyVaultResourceID string  `json:"KeyVaultResourceId,omitempty" bson:"KeyVaultResourceId,omitempty"`
	KeyVaultURL        string  `json:"KeyVaultURL,omitempty" bson:"KeyVaultURL,omitempty"`
	Name               string  `json:"Name,omitempty" bson:"Name,omitempty"`
	OuPath             string  `json:"OUPath,omitempty" bson:"OUPath,omitempty"`
	Options            float64 `json:"Options,omitempty" bson:"Options,omitempty"`
	Properties         []struct {
		Name     string `json:"Name,omitempty" bson:"Name,omitempty"`
		TypeName string `json:"TypeName,omitempty" bson:"TypeName,omitempty"`
		Value    any    `json:"Value,omitempty" bson:"Value,omitempty"`
	} `json:"Properties,omitempty" bson:"Properties,omitempty"`
	RealtimeProtectionEnabled string `json:"RealtimeProtectionEnabled,omitempty" bson:"RealtimeProtectionEnabled,omitempty"`
	RegistrationSettings      *struct {
		ProvisionExtensionWithNoSQLServer bool   `json:"ProvisionExtensionWithNoSQLServer,omitempty" bson:"ProvisionExtensionWithNoSQLServer,omitempty"`
		RegistrationSource                string `json:"RegistrationSource,omitempty" bson:"RegistrationSource,omitempty"`
	} `json:"RegistrationSettings,omitempty" bson:"RegistrationSettings,omitempty"`
	ResizeOSDisk          bool   `json:"ResizeOSDisk,omitempty" bson:"ResizeOSDisk,omitempty"`
	Restart               string `json:"Restart,omitempty" bson:"Restart,omitempty"`
	ScheduledScanSettings *struct {
		Day       string `json:"day,omitempty" bson:"day,omitempty"`
		IsEnabled string `json:"isEnabled,omitempty" bson:"isEnabled,omitempty"`
		ScanType  string `json:"scanType,omitempty" bson:"scanType,omitempty"`
		Time      string `json:"time,omitempty" bson:"time,omitempty"`
	} `json:"ScheduledScanSettings,omitempty" bson:"ScheduledScanSettings,omitempty"`
	SequenceVersion                        string                                               `json:"SequenceVersion,omitempty" bson:"SequenceVersion,omitempty"`
	ServerConfigurationsManagementSettings *AzureResourceSettingsServerConfigurationsManagement `json:"ServerConfigurationsManagementSettings,omitempty" bson:"ServerConfigurationsManagementSettings,omitempty"`
	SQLManagement                          *struct {
		IsEnabled bool `json:"IsEnabled,omitempty" bson:"IsEnabled,omitempty"`
	} `json:"SqlManagement,omitempty" bson:"SqlManagement,omitempty"`
	AutoUpdate                    bool                                         `json:"autoUpdate,omitempty" bson:"autoUpdate,omitempty"`
	AzureResourceID               *string                                      `json:"azureResourceId,omitempty" bson:"azureResourceId,omitempty"`
	CommandStartTimeUtcTicks      string                                       `json:"commandStartTimeUTCTicks,omitempty" bson:"commandStartTimeUTCTicks,omitempty"`
	CommandToExecute              string                                       `json:"commandToExecute,omitempty" bson:"commandToExecute,omitempty"`
	ConfigurationArguments        *AzureResourceSettingsConfigurationArguments `json:"configurationArguments,omitempty" bson:"configurationArguments,omitempty"`
	DefenderForServersWorkspaceID string                                       `json:"defenderForServersWorkspaceId,omitempty" bson:"defenderForServersWorkspaceId,omitempty"`
	EnableAma                     string                                       `json:"enableAMA,omitempty" bson:"enableAMA,omitempty"`
	FileUris                      []string                                     `json:"fileUris,omitempty" bson:"fileUris,omitempty"`
	ForceReOnboarding             bool                                         `json:"forceReOnboarding,omitempty" bson:"forceReOnboarding,omitempty"`
	Locale                        string                                       `json:"locale,omitempty" bson:"locale,omitempty"`
	ObjectStr                     string                                       `json:"objectStr,omitempty" bson:"objectStr,omitempty"`
	Port                          any                                          `json:"port,omitempty" bson:"port,omitempty"`
	Protocol                      string                                       `json:"protocol,omitempty" bson:"protocol,omitempty"`
	RequestPath                   string                                       `json:"requestPath,omitempty" bson:"requestPath,omitempty"`
	Salt                          string                                       `json:"salt,omitempty" bson:"salt,omitempty"`
	SkipDos2Unix                  bool                                         `json:"skipDos2Unix,omitempty" bson:"skipDos2Unix,omitempty"`
	StopOnMultipleConnections     any                                          `json:"stopOnMultipleConnections,omitempty" bson:"stopOnMultipleConnections,omitempty"`
	StorageAccount                string                                       `json:"storageAccount,omitempty" bson:"storageAccount,omitempty"`
	StorageAccountOther           string                                       `json:"StorageAccount,omitempty" bson:"StorageAccount,omitempty"`
	TaskID                        string                                       `json:"taskId,omitempty" bson:"taskId,omitempty"`
	Timestamp                     float64                                      `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	TimeStamp                     string                                       `json:"timeStamp,omitempty" bson:"timeStamp,omitempty"`
	TriggerForceUpgrade           bool                                         `json:"triggerForceUpgrade,omitempty" bson:"triggerForceUpgrade,omitempty"`
	User                          string                                       `json:"User,omitempty" bson:"User,omitempty"`
	UserName                      string                                       `json:"userName,omitempty" bson:"userName,omitempty"`
	UserNameOther                 string                                       `json:"UserName,omitempty" bson:"UserName,omitempty"`
	VmType                        string                                       `json:"vmType,omitempty" bson:"vmType,omitempty"`
	VNextEnabled                  bool                                         `json:"vNextEnabled,omitempty" bson:"vNextEnabled,omitempty"`
	VolumeType                    string                                       `json:"VolumeType,omitempty" bson:"VolumeType,omitempty"`
	WadCfg                        *AzureResourceSettingsWadCfg                 `json:"WadCfg,omitempty" bson:"WadCfg,omitempty"`
	WorkspaceID                   string                                       `json:"workspaceId,omitempty" bson:"workspaceId,omitempty"`
	XMLCfg                        string                                       `json:"xmlCfg,omitempty" bson:"xmlCfg,omitempty"`
}

type AzureResourceSiteConfig struct {
	AcrUseManagedIdentityCreds             bool    `json:"acrUseManagedIdentityCreds,omitempty" bson:"acrUseManagedIdentityCreds,omitempty"`
	AcrUserManagedIdentityID               any     `json:"acrUserManagedIdentityID,omitempty" bson:"acrUserManagedIdentityID,omitempty"`
	AlwaysOn                               bool    `json:"alwaysOn,omitempty" bson:"alwaysOn,omitempty"`
	AntivirusScanEnabled                   any     `json:"antivirusScanEnabled,omitempty" bson:"antivirusScanEnabled,omitempty"`
	APIDefinition                          any     `json:"apiDefinition,omitempty" bson:"apiDefinition,omitempty"`
	APIManagementConfig                    any     `json:"apiManagementConfig,omitempty" bson:"apiManagementConfig,omitempty"`
	AppCommandLine                         any     `json:"appCommandLine,omitempty" bson:"appCommandLine,omitempty"`
	AppSettings                            any     `json:"appSettings,omitempty" bson:"appSettings,omitempty"`
	AutoHealEnabled                        any     `json:"autoHealEnabled,omitempty" bson:"autoHealEnabled,omitempty"`
	AutoHealRules                          any     `json:"autoHealRules,omitempty" bson:"autoHealRules,omitempty"`
	AutoSwapSlotName                       any     `json:"autoSwapSlotName,omitempty" bson:"autoSwapSlotName,omitempty"`
	AzureMonitorLogCategories              any     `json:"azureMonitorLogCategories,omitempty" bson:"azureMonitorLogCategories,omitempty"`
	AzureStorageAccounts                   any     `json:"azureStorageAccounts,omitempty" bson:"azureStorageAccounts,omitempty"`
	ClusteringEnabled                      bool    `json:"clusteringEnabled,omitempty" bson:"clusteringEnabled,omitempty"`
	ConnectionStrings                      any     `json:"connectionStrings,omitempty" bson:"connectionStrings,omitempty"`
	Cors                                   any     `json:"cors,omitempty" bson:"cors,omitempty"`
	CustomAppPoolIdentityAdminState        any     `json:"customAppPoolIdentityAdminState,omitempty" bson:"customAppPoolIdentityAdminState,omitempty"`
	CustomAppPoolIdentityTenantState       any     `json:"customAppPoolIdentityTenantState,omitempty" bson:"customAppPoolIdentityTenantState,omitempty"`
	DefaultDocuments                       any     `json:"defaultDocuments,omitempty" bson:"defaultDocuments,omitempty"`
	DetailedErrorLoggingEnabled            any     `json:"detailedErrorLoggingEnabled,omitempty" bson:"detailedErrorLoggingEnabled,omitempty"`
	DocumentRoot                           any     `json:"documentRoot,omitempty" bson:"documentRoot,omitempty"`
	ElasticWebAppScaleLimit                any     `json:"elasticWebAppScaleLimit,omitempty" bson:"elasticWebAppScaleLimit,omitempty"`
	Experiments                            any     `json:"experiments,omitempty" bson:"experiments,omitempty"`
	FileChangeAuditEnabled                 any     `json:"fileChangeAuditEnabled,omitempty" bson:"fileChangeAuditEnabled,omitempty"`
	FtpsState                              any     `json:"ftpsState,omitempty" bson:"ftpsState,omitempty"`
	FunctionAppScaleLimit                  float64 `json:"functionAppScaleLimit,omitempty" bson:"functionAppScaleLimit,omitempty"`
	FunctionsRuntimeScaleMonitoringEnabled any     `json:"functionsRuntimeScaleMonitoringEnabled,omitempty" bson:"functionsRuntimeScaleMonitoringEnabled,omitempty"`
	HandlerMappings                        any     `json:"handlerMappings,omitempty" bson:"handlerMappings,omitempty"`
	HealthCheckPath                        any     `json:"healthCheckPath,omitempty" bson:"healthCheckPath,omitempty"`
	HTTP20Enabled                          bool    `json:"http20Enabled,omitempty" bson:"http20Enabled,omitempty"`
	HTTP20ProxyFlag                        any     `json:"http20ProxyFlag,omitempty" bson:"http20ProxyFlag,omitempty"`
	HTTPLoggingEnabled                     any     `json:"httpLoggingEnabled,omitempty" bson:"httpLoggingEnabled,omitempty"`
	IpSecurityRestrictions                 any     `json:"ipSecurityRestrictions,omitempty" bson:"ipSecurityRestrictions,omitempty"`
	IpSecurityRestrictionsDefaultAction    any     `json:"ipSecurityRestrictionsDefaultAction,omitempty" bson:"ipSecurityRestrictionsDefaultAction,omitempty"`
	JavaContainer                          any     `json:"javaContainer,omitempty" bson:"javaContainer,omitempty"`
	JavaContainerVersion                   any     `json:"javaContainerVersion,omitempty" bson:"javaContainerVersion,omitempty"`
	JavaVersion                            any     `json:"javaVersion,omitempty" bson:"javaVersion,omitempty"`
	KeyVaultReferenceIdentity              any     `json:"keyVaultReferenceIdentity,omitempty" bson:"keyVaultReferenceIdentity,omitempty"`
	Limits                                 any     `json:"limits,omitempty" bson:"limits,omitempty"`
	LinuxFxVersion                         string  `json:"linuxFxVersion,omitempty" bson:"linuxFxVersion,omitempty"`
	LoadBalancing                          any     `json:"loadBalancing,omitempty" bson:"loadBalancing,omitempty"`
	LocalMySQLEnabled                      any     `json:"localMySqlEnabled,omitempty" bson:"localMySqlEnabled,omitempty"`
	LogsDirectorySizeLimit                 any     `json:"logsDirectorySizeLimit,omitempty" bson:"logsDirectorySizeLimit,omitempty"`
	MachineKey                             any     `json:"machineKey,omitempty" bson:"machineKey,omitempty"`
	ManagedPipelineMode                    any     `json:"managedPipelineMode,omitempty" bson:"managedPipelineMode,omitempty"`
	ManagedServiceIdentityID               any     `json:"managedServiceIdentityId,omitempty" bson:"managedServiceIdentityId,omitempty"`
	Metadata                               any     `json:"metadata,omitempty" bson:"metadata,omitempty"`
	MinTlsCipherSuite                      any     `json:"minTlsCipherSuite,omitempty" bson:"minTlsCipherSuite,omitempty"`
	MinTlsVersion                          any     `json:"minTlsVersion,omitempty" bson:"minTlsVersion,omitempty"`
	MinimumElasticInstanceCount            float64 `json:"minimumElasticInstanceCount,omitempty" bson:"minimumElasticInstanceCount,omitempty"`
	NetFrameworkVersion                    any     `json:"netFrameworkVersion,omitempty" bson:"netFrameworkVersion,omitempty"`
	NodeVersion                            any     `json:"nodeVersion,omitempty" bson:"nodeVersion,omitempty"`
	NumberOfWorkers                        float64 `json:"numberOfWorkers,omitempty" bson:"numberOfWorkers,omitempty"`
	PhpVersion                             any     `json:"phpVersion,omitempty" bson:"phpVersion,omitempty"`
	PowerShellVersion                      any     `json:"powerShellVersion,omitempty" bson:"powerShellVersion,omitempty"`
	PreWarmedInstanceCount                 any     `json:"preWarmedInstanceCount,omitempty" bson:"preWarmedInstanceCount,omitempty"`
	PublicNetworkAccess                    any     `json:"publicNetworkAccess,omitempty" bson:"publicNetworkAccess,omitempty"`
	PublishingPassword                     any     `json:"publishingPassword,omitempty" bson:"publishingPassword,omitempty"`
	PublishingUsername                     any     `json:"publishingUsername,omitempty" bson:"publishingUsername,omitempty"`
	Push                                   any     `json:"push,omitempty" bson:"push,omitempty"`
	PythonVersion                          any     `json:"pythonVersion,omitempty" bson:"pythonVersion,omitempty"`
	RemoteDebuggingEnabled                 any     `json:"remoteDebuggingEnabled,omitempty" bson:"remoteDebuggingEnabled,omitempty"`
	RemoteDebuggingVersion                 any     `json:"remoteDebuggingVersion,omitempty" bson:"remoteDebuggingVersion,omitempty"`
	RequestTracingEnabled                  any     `json:"requestTracingEnabled,omitempty" bson:"requestTracingEnabled,omitempty"`
	RoutingRules                           any     `json:"routingRules,omitempty" bson:"routingRules,omitempty"`
	RuntimeAdUser                          any     `json:"runtimeADUser,omitempty" bson:"runtimeADUser,omitempty"`
	RuntimeAdUserPassword                  any     `json:"runtimeADUserPassword,omitempty" bson:"runtimeADUserPassword,omitempty"`
	ScmIpSecurityRestrictions              any     `json:"scmIpSecurityRestrictions,omitempty" bson:"scmIpSecurityRestrictions,omitempty"`
	ScmIpSecurityRestrictionsDefaultAction any     `json:"scmIpSecurityRestrictionsDefaultAction,omitempty" bson:"scmIpSecurityRestrictionsDefaultAction,omitempty"`
	ScmIpSecurityRestrictionsUseMain       any     `json:"scmIpSecurityRestrictionsUseMain,omitempty" bson:"scmIpSecurityRestrictionsUseMain,omitempty"`
	ScmMinTlsCipherSuite                   any     `json:"scmMinTlsCipherSuite,omitempty" bson:"scmMinTlsCipherSuite,omitempty"`
	ScmMinTlsVersion                       any     `json:"scmMinTlsVersion,omitempty" bson:"scmMinTlsVersion,omitempty"`
	ScmSupportedTlsCipherSuites            any     `json:"scmSupportedTlsCipherSuites,omitempty" bson:"scmSupportedTlsCipherSuites,omitempty"`
	ScmType                                any     `json:"scmType,omitempty" bson:"scmType,omitempty"`
	SitePort                               any     `json:"sitePort,omitempty" bson:"sitePort,omitempty"`
	SitePrivateLinkHostEnabled             any     `json:"sitePrivateLinkHostEnabled,omitempty" bson:"sitePrivateLinkHostEnabled,omitempty"`
	StorageType                            any     `json:"storageType,omitempty" bson:"storageType,omitempty"`
	SupportedTlsCipherSuites               any     `json:"supportedTlsCipherSuites,omitempty" bson:"supportedTlsCipherSuites,omitempty"`
	TracingOptions                         any     `json:"tracingOptions,omitempty" bson:"tracingOptions,omitempty"`
	Use32BitWorkerProcess                  any     `json:"use32BitWorkerProcess,omitempty" bson:"use32BitWorkerProcess,omitempty"`
	VirtualApplications                    any     `json:"virtualApplications,omitempty" bson:"virtualApplications,omitempty"`
	VnetName                               any     `json:"vnetName,omitempty" bson:"vnetName,omitempty"`
	VnetPrivatePortsCount                  any     `json:"vnetPrivatePortsCount,omitempty" bson:"vnetPrivatePortsCount,omitempty"`
	VnetRouteAllEnabled                    any     `json:"vnetRouteAllEnabled,omitempty" bson:"vnetRouteAllEnabled,omitempty"`
	WebSocketsEnabled                      any     `json:"webSocketsEnabled,omitempty" bson:"webSocketsEnabled,omitempty"`
	WebsiteTimeZone                        any     `json:"websiteTimeZone,omitempty" bson:"websiteTimeZone,omitempty"`
	WinAuthAdminState                      any     `json:"winAuthAdminState,omitempty" bson:"winAuthAdminState,omitempty"`
	WinAuthTenantState                     any     `json:"winAuthTenantState,omitempty" bson:"winAuthTenantState,omitempty"`
	WindowsConfiguredStacks                any     `json:"windowsConfiguredStacks,omitempty" bson:"windowsConfiguredStacks,omitempty"`
	WindowsFxVersion                       any     `json:"windowsFxVersion,omitempty" bson:"windowsFxVersion,omitempty"`
	XManagedServiceIdentityID              any     `json:"xManagedServiceIdentityId,omitempty" bson:"xManagedServiceIdentityId,omitempty"`
}

type AzureResourceSiteProperties struct {
	AppSettings any `json:"appSettings,omitempty" bson:"appSettings,omitempty"`
	Metadata    any `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Properties  []struct {
		Name  string  `json:"name,omitempty" bson:"name,omitempty"`
		Value *string `json:"value,omitempty" bson:"value,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
}

type AzureResourceSources struct {
	EventSource string `json:"eventSource,omitempty" bson:"eventSource,omitempty"`
	RuleSets    []struct {
		Rules []struct {
			ExpectedValue string `json:"expectedValue,omitempty" bson:"expectedValue,omitempty"`
			Operator      string `json:"operator,omitempty" bson:"operator,omitempty"`
			PropertyJPath string `json:"propertyJPath,omitempty" bson:"propertyJPath,omitempty"`
			PropertyType  string `json:"propertyType,omitempty" bson:"propertyType,omitempty"`
		} `json:"rules,omitempty" bson:"rules,omitempty"`
	} `json:"ruleSets,omitempty" bson:"ruleSets,omitempty"`
}

type AzureResourceStatus struct {
	Error *struct {
		Code    string `json:"code,omitempty" bson:"code,omitempty"`
		Message string `json:"message,omitempty" bson:"message,omitempty"`
	} `json:"error,omitempty" bson:"error,omitempty"`
	Status string `json:"status,omitempty" bson:"status,omitempty"`
	Target string `json:"target,omitempty" bson:"target,omitempty"`
}

type AzureResourceStorage struct {
	AutoGrow      string  `json:"autoGrow,omitempty" bson:"autoGrow,omitempty"`
	Iops          float64 `json:"iops,omitempty" bson:"iops,omitempty"`
	StorageSizeGb float64 `json:"storageSizeGB,omitempty" bson:"storageSizeGB,omitempty"`
	Tier          string  `json:"tier,omitempty" bson:"tier,omitempty"`
	Type          string  `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceStorageProfileDataDisk struct {
	Caching      string  `json:"caching,omitempty" bson:"caching,omitempty"`
	CreateOption string  `json:"createOption,omitempty" bson:"createOption,omitempty"`
	DeleteOption string  `json:"deleteOption,omitempty" bson:"deleteOption,omitempty"`
	DiskSizeGb   float64 `json:"diskSizeGB,omitempty" bson:"diskSizeGB,omitempty"`
	Lun          float64 `json:"lun,omitempty" bson:"lun,omitempty"`
	ManagedDisk  struct {
		ID                 string `json:"id,omitempty" bson:"id,omitempty"`
		StorageAccountType string `json:"storageAccountType,omitempty" bson:"storageAccountType,omitempty"`
	} `json:"managedDisk,omitempty" bson:"managedDisk,omitempty"`
	Name                    string `json:"name,omitempty" bson:"name,omitempty"`
	ToBeDetached            bool   `json:"toBeDetached,omitempty" bson:"toBeDetached,omitempty"`
	WriteAcceleratorEnabled bool   `json:"writeAcceleratorEnabled,omitempty" bson:"writeAcceleratorEnabled,omitempty"`
}

type AzureResourceStorageProfileImageReference struct {
	ExactVersion string `json:"exactVersion,omitempty" bson:"exactVersion,omitempty"`
	ID           string `json:"id,omitempty" bson:"id,omitempty"`
	Offer        string `json:"offer,omitempty" bson:"offer,omitempty"`
	Publisher    string `json:"publisher,omitempty" bson:"publisher,omitempty"`
	Sku          string `json:"sku,omitempty" bson:"sku,omitempty"`
	Version      string `json:"version,omitempty" bson:"version,omitempty"`
}

type AzureResourceStorageProfileOSDisk struct {
	BlobURI          string `json:"blobUri,omitempty" bson:"blobUri,omitempty"`
	Caching          string `json:"caching,omitempty" bson:"caching,omitempty"`
	CreateOption     string `json:"createOption,omitempty" bson:"createOption,omitempty"`
	DeleteOption     string `json:"deleteOption,omitempty" bson:"deleteOption,omitempty"`
	DiffDiskSettings *struct {
		Option    string `json:"option,omitempty" bson:"option,omitempty"`
		Placement string `json:"placement,omitempty" bson:"placement,omitempty"`
	} `json:"diffDiskSettings,omitempty" bson:"diffDiskSettings,omitempty"`
	DiskSizeGb  float64 `json:"diskSizeGB,omitempty" bson:"diskSizeGB,omitempty"`
	ManagedDisk *struct {
		ID                 string `json:"id,omitempty" bson:"id,omitempty"`
		StorageAccountType string `json:"storageAccountType,omitempty" bson:"storageAccountType,omitempty"`
	} `json:"managedDisk,omitempty" bson:"managedDisk,omitempty"`
	Name               string `json:"name,omitempty" bson:"name,omitempty"`
	OSState            string `json:"osState,omitempty" bson:"osState,omitempty"`
	OSType             string `json:"osType,omitempty" bson:"osType,omitempty"`
	StorageAccountType string `json:"storageAccountType,omitempty" bson:"storageAccountType,omitempty"`
	Vhd                *struct {
		URI string `json:"uri,omitempty" bson:"uri,omitempty"`
	} `json:"vhd,omitempty" bson:"vhd,omitempty"`
	WriteAcceleratorEnabled bool `json:"writeAcceleratorEnabled,omitempty" bson:"writeAcceleratorEnabled,omitempty"`
}

type AzureResourceStorageProfile struct {
	DataDisks     []AzureResourceStorageProfileDataDisk `json:"dataDisks,omitempty" bson:"dataDisks,omitempty"`
	DiskCsiDriver *struct {
		Enabled bool `json:"enabled,omitempty" bson:"enabled,omitempty"`
	} `json:"diskCSIDriver,omitempty" bson:"diskCSIDriver,omitempty"`
	DiskControllerType string `json:"diskControllerType,omitempty" bson:"diskControllerType,omitempty"`
	FileCsiDriver      *struct {
		Enabled bool `json:"enabled,omitempty" bson:"enabled,omitempty"`
	} `json:"fileCSIDriver,omitempty" bson:"fileCSIDriver,omitempty"`
	ImageReference *AzureResourceStorageProfileImageReference `json:"imageReference,omitempty" bson:"imageReference,omitempty"`
	OSDisk         *AzureResourceStorageProfileOSDisk         `json:"osDisk,omitempty" bson:"osDisk,omitempty"`
	OSDiskImage    *struct {
		HostCaching string  `json:"hostCaching,omitempty" bson:"hostCaching,omitempty"`
		SizeInGb    float64 `json:"sizeInGB,omitempty" bson:"sizeInGB,omitempty"`
	} `json:"osDiskImage,omitempty" bson:"osDiskImage,omitempty"`
	SnapshotController *struct {
		Enabled bool `json:"enabled,omitempty" bson:"enabled,omitempty"`
	} `json:"snapshotController,omitempty" bson:"snapshotController,omitempty"`
	ZoneResilient bool `json:"zoneResilient,omitempty" bson:"zoneResilient,omitempty"`
}

type AzureResourceStreamDeclarations struct {
	CustomTextLoki_CL *struct {
		Columns []struct {
			Name string `json:"name,omitempty" bson:"name,omitempty"`
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"columns,omitempty" bson:"columns,omitempty"`
	} `json:"Custom-Text-Loki_CL,omitempty" bson:"Custom-Text-Loki_CL,omitempty"`
}

type AzureResourceSubnet struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties *struct {
		AddressPrefix         string `json:"addressPrefix,omitempty" bson:"addressPrefix,omitempty"`
		DefaultOutboundAccess bool   `json:"defaultOutboundAccess,omitempty" bson:"defaultOutboundAccess,omitempty"`
		Delegations           []struct {
			Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
			ID         string `json:"id,omitempty" bson:"id,omitempty"`
			Name       string `json:"name,omitempty" bson:"name,omitempty"`
			Properties struct {
				Actions           []string `json:"actions,omitempty" bson:"actions,omitempty"`
				ProvisioningState string   `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
				ServiceName       string   `json:"serviceName,omitempty" bson:"serviceName,omitempty"`
			} `json:"properties,omitempty" bson:"properties,omitempty"`
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"delegations,omitempty" bson:"delegations,omitempty"`
		IpConfigurations []struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"ipConfigurations,omitempty" bson:"ipConfigurations,omitempty"`
		NetworkSecurityGroup *struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"networkSecurityGroup,omitempty" bson:"networkSecurityGroup,omitempty"`
		PrivateEndpointNetworkPolicies string `json:"privateEndpointNetworkPolicies,omitempty" bson:"privateEndpointNetworkPolicies,omitempty"`
		PrivateEndpoints               []struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"privateEndpoints,omitempty" bson:"privateEndpoints,omitempty"`
		PrivateLinkServiceNetworkPolicies string `json:"privateLinkServiceNetworkPolicies,omitempty" bson:"privateLinkServiceNetworkPolicies,omitempty"`
		ProvisioningState                 string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
		Purpose                           string `json:"purpose,omitempty" bson:"purpose,omitempty"`
		RouteTable                        *struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"routeTable,omitempty" bson:"routeTable,omitempty"`
		ServiceAssociationLinks []struct {
			Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
			ID         string `json:"id,omitempty" bson:"id,omitempty"`
			Name       string `json:"name,omitempty" bson:"name,omitempty"`
			Properties struct {
				AllowDelete              bool   `json:"allowDelete,omitempty" bson:"allowDelete,omitempty"`
				EnabledForArmDeployments bool   `json:"enabledForArmDeployments,omitempty" bson:"enabledForArmDeployments,omitempty"`
				Link                     string `json:"link,omitempty" bson:"link,omitempty"`
				LinkedResourceType       string `json:"linkedResourceType,omitempty" bson:"linkedResourceType,omitempty"`
				Locations                []any  `json:"locations,omitempty" bson:"locations,omitempty"`
				ProvisioningState        string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
				SubnetID                 string `json:"subnetId,omitempty" bson:"subnetId,omitempty"`
			} `json:"properties,omitempty" bson:"properties,omitempty"`
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"serviceAssociationLinks,omitempty" bson:"serviceAssociationLinks,omitempty"`
		ServiceEndpoints []struct {
			Locations         []string `json:"locations,omitempty" bson:"locations,omitempty"`
			ProvisioningState string   `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
			Service           string   `json:"service,omitempty" bson:"service,omitempty"`
		} `json:"serviceEndpoints,omitempty" bson:"serviceEndpoints,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceSystemData struct {
	CreatedAt          string `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	CreatedBy          string `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	CreatedByType      string `json:"createdByType,omitempty" bson:"createdByType,omitempty"`
	LastModifiedAt     string `json:"lastModifiedAt,omitempty" bson:"lastModifiedAt,omitempty"`
	LastModifiedBy     string `json:"lastModifiedBy,omitempty" bson:"lastModifiedBy,omitempty"`
	LastModifiedByType string `json:"lastModifiedByType,omitempty" bson:"lastModifiedByType,omitempty"`
}

type AzureResourceTemplate struct {
	Schema     string `json:"$schema,omitempty" bson:"_schema,omitempty"`
	Containers []struct {
		Env []struct {
			Name      string `json:"name,omitempty" bson:"name,omitempty"`
			SecretRef string `json:"secretRef,omitempty" bson:"secretRef,omitempty"`
			Value     string `json:"value,omitempty" bson:"value,omitempty"`
		} `json:"env,omitempty" bson:"env,omitempty"`
		Image     string `json:"image,omitempty" bson:"image,omitempty"`
		ImageType string `json:"imageType,omitempty" bson:"imageType,omitempty"`
		Name      string `json:"name,omitempty" bson:"name,omitempty"`
		Resources struct {
			Cpu              float64 `json:"cpu,omitempty" bson:"cpu,omitempty"`
			EphemeralStorage string  `json:"ephemeralStorage,omitempty" bson:"ephemeralStorage,omitempty"`
			Memory           string  `json:"memory,omitempty" bson:"memory,omitempty"`
		} `json:"resources,omitempty" bson:"resources,omitempty"`
	} `json:"containers,omitempty" bson:"containers,omitempty"`
	ContentVersion string `json:"contentVersion,omitempty" bson:"contentVersion,omitempty"`
	InitContainers any    `json:"initContainers,omitempty" bson:"initContainers,omitempty"`
	Outputs        *struct {
		AdminUsername struct {
			Type  string `json:"type,omitempty" bson:"type,omitempty"`
			Value string `json:"value,omitempty" bson:"value,omitempty"`
		} `json:"adminUsername,omitempty" bson:"adminUsername,omitempty"`
	} `json:"outputs,omitempty" bson:"outputs,omitempty"`
	Parameters *struct {
		AdminPassword *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"adminPassword,omitempty" bson:"adminPassword,omitempty"`
		AdminUsername *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"adminUsername,omitempty" bson:"adminUsername,omitempty"`
		AutoShutdownNotificationEmail *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"autoShutdownNotificationEmail,omitempty" bson:"autoShutdownNotificationEmail,omitempty"`
		AutoShutdownNotificationLocale *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"autoShutdownNotificationLocale,omitempty" bson:"autoShutdownNotificationLocale,omitempty"`
		AutoShutdownNotificationStatus *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"autoShutdownNotificationStatus,omitempty" bson:"autoShutdownNotificationStatus,omitempty"`
		AutoShutdownStatus *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"autoShutdownStatus,omitempty" bson:"autoShutdownStatus,omitempty"`
		AutoShutdownTime *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"autoShutdownTime,omitempty" bson:"autoShutdownTime,omitempty"`
		AutoShutdownTimeZone *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"autoShutdownTimeZone,omitempty" bson:"autoShutdownTimeZone,omitempty"`
		DataDiskResources *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"dataDiskResources,omitempty" bson:"dataDiskResources,omitempty"`
		DataDisks *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"dataDisks,omitempty" bson:"dataDisks,omitempty"`
		EnableAcceleratedNetworking *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"enableAcceleratedNetworking,omitempty" bson:"enableAcceleratedNetworking,omitempty"`
		Location *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"location,omitempty" bson:"location,omitempty"`
		NetworkInterfaceName *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"networkInterfaceName,omitempty" bson:"networkInterfaceName,omitempty"`
		NicDeleteOption *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"nicDeleteOption,omitempty" bson:"nicDeleteOption,omitempty"`
		OSDiskDeleteOption *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"osDiskDeleteOption,omitempty" bson:"osDiskDeleteOption,omitempty"`
		OSDiskType *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"osDiskType,omitempty" bson:"osDiskType,omitempty"`
		SecureBoot *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"secureBoot,omitempty" bson:"secureBoot,omitempty"`
		SecurityType *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"securityType,omitempty" bson:"securityType,omitempty"`
		SubnetName *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"subnetName,omitempty" bson:"subnetName,omitempty"`
		VTpm *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"vTPM,omitempty" bson:"vTPM,omitempty"`
		VirtualMachineComputerName *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"virtualMachineComputerName,omitempty" bson:"virtualMachineComputerName,omitempty"`
		VirtualMachineName *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"virtualMachineName,omitempty" bson:"virtualMachineName,omitempty"`
		VirtualMachineRg *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"virtualMachineRG,omitempty" bson:"virtualMachineRG,omitempty"`
		VirtualMachineSize *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"virtualMachineSize,omitempty" bson:"virtualMachineSize,omitempty"`
		VirtualNetworkID *struct {
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"virtualNetworkId,omitempty" bson:"virtualNetworkId,omitempty"`
	} `json:"parameters,omitempty" bson:"parameters,omitempty"`
	Resources []struct {
		APIVersion string `json:"apiVersion,omitempty" bson:"apiVersion,omitempty"`
		Copy       *struct {
			Count string `json:"count,omitempty" bson:"count,omitempty"`
			Name  string `json:"name,omitempty" bson:"name,omitempty"`
		} `json:"copy,omitempty" bson:"copy,omitempty"`
		DependsOn []string `json:"dependsOn,omitempty" bson:"dependsOn,omitempty"`
		Kind      string   `json:"kind,omitempty" bson:"kind,omitempty"`
		Location  string   `json:"location,omitempty" bson:"location,omitempty"`
		Metadata  *struct {
			Description string `json:"description,omitempty" bson:"description,omitempty"`
		} `json:"metadata,omitempty" bson:"metadata,omitempty"`
		Name       string `json:"name,omitempty" bson:"name,omitempty"`
		Properties any    `json:"properties,omitempty" bson:"properties,omitempty"`
		Sku        *struct {
			Name string `json:"name,omitempty" bson:"name,omitempty"`
		} `json:"sku,omitempty" bson:"sku,omitempty"`
		Tags  map[string]string `json:"tags" bson:"tags"`
		Type  string            `json:"type,omitempty" bson:"type,omitempty"`
		Zones []string          `json:"zones,omitempty" bson:"zones,omitempty"`
	} `json:"resources,omitempty" bson:"resources,omitempty"`
	Variables *struct {
		SubnetRef string `json:"subnetRef,omitempty" bson:"subnetRef,omitempty"`
		VnetID    string `json:"vnetId,omitempty" bson:"vnetId,omitempty"`
		VnetName  string `json:"vnetName,omitempty" bson:"vnetName,omitempty"`
	} `json:"variables,omitempty" bson:"variables,omitempty"`
	Volumes any `json:"volumes,omitempty" bson:"volumes,omitempty"`
}

type AzureResourceTestConfigurations struct {
	Name             string `json:"name,omitempty" bson:"name,omitempty"`
	Protocol         string `json:"protocol,omitempty" bson:"protocol,omitempty"`
	TcpConfiguration struct {
		DisableTraceRoute bool    `json:"disableTraceRoute,omitempty" bson:"disableTraceRoute,omitempty"`
		Port              float64 `json:"port,omitempty" bson:"port,omitempty"`
	} `json:"tcpConfiguration,omitempty" bson:"tcpConfiguration,omitempty"`
	TestFrequencySec float64 `json:"testFrequencySec,omitempty" bson:"testFrequencySec,omitempty"`
}

type AzureResourceTestGroups struct {
	Destinations       []string `json:"destinations,omitempty" bson:"destinations,omitempty"`
	Disable            bool     `json:"disable,omitempty" bson:"disable,omitempty"`
	Name               string   `json:"name,omitempty" bson:"name,omitempty"`
	Sources            []string `json:"sources,omitempty" bson:"sources,omitempty"`
	TestConfigurations []string `json:"testConfigurations,omitempty" bson:"testConfigurations,omitempty"`
}

type AzureResourceTestRequests struct {
	Body struct {
		Request struct {
			Method string `json:"method,omitempty" bson:"method,omitempty"`
			Path   string `json:"path,omitempty" bson:"path,omitempty"`
		} `json:"request,omitempty" bson:"request,omitempty"`
	} `json:"body,omitempty" bson:"body,omitempty"`
	Method     string `json:"method,omitempty" bson:"method,omitempty"`
	RequestURI string `json:"requestUri,omitempty" bson:"requestUri,omitempty"`
}

type AzureResourceTransportSecurity struct {
	CertificateAuthority struct {
		KeyVaultSecretID string `json:"keyVaultSecretId,omitempty" bson:"keyVaultSecretId,omitempty"`
		Name             string `json:"name,omitempty" bson:"name,omitempty"`
	} `json:"certificateAuthority,omitempty" bson:"certificateAuthority,omitempty"`
}

type AzureResourceVerificationRecords struct {
	Dkim *struct {
		Name  string  `json:"name,omitempty" bson:"name,omitempty"`
		Ttl   float64 `json:"ttl,omitempty" bson:"ttl,omitempty"`
		Type  string  `json:"type,omitempty" bson:"type,omitempty"`
		Value string  `json:"value,omitempty" bson:"value,omitempty"`
	} `json:"DKIM,omitempty" bson:"DKIM,omitempty"`
	Dkim2 *struct {
		Name  string  `json:"name,omitempty" bson:"name,omitempty"`
		Ttl   float64 `json:"ttl,omitempty" bson:"ttl,omitempty"`
		Type  string  `json:"type,omitempty" bson:"type,omitempty"`
		Value string  `json:"value,omitempty" bson:"value,omitempty"`
	} `json:"DKIM2,omitempty" bson:"DKIM2,omitempty"`
	Domain *struct {
		Name  string  `json:"name,omitempty" bson:"name,omitempty"`
		Ttl   float64 `json:"ttl,omitempty" bson:"ttl,omitempty"`
		Type  string  `json:"type,omitempty" bson:"type,omitempty"`
		Value string  `json:"value,omitempty" bson:"value,omitempty"`
	} `json:"Domain,omitempty" bson:"Domain,omitempty"`
	Spf *struct {
		Name  string  `json:"name,omitempty" bson:"name,omitempty"`
		Ttl   float64 `json:"ttl,omitempty" bson:"ttl,omitempty"`
		Type  string  `json:"type,omitempty" bson:"type,omitempty"`
		Value string  `json:"value,omitempty" bson:"value,omitempty"`
	} `json:"SPF,omitempty" bson:"SPF,omitempty"`
}

type AzureResourceVerificationStates struct {
	Dkim struct {
		ErrorCode string `json:"errorCode,omitempty" bson:"errorCode,omitempty"`
		Status    string `json:"status,omitempty" bson:"status,omitempty"`
	} `json:"DKIM,omitempty" bson:"DKIM,omitempty"`
	Dkim2 struct {
		ErrorCode string `json:"errorCode,omitempty" bson:"errorCode,omitempty"`
		Status    string `json:"status,omitempty" bson:"status,omitempty"`
	} `json:"DKIM2,omitempty" bson:"DKIM2,omitempty"`
	Dmarc struct {
		Status string `json:"status,omitempty" bson:"status,omitempty"`
	} `json:"DMARC,omitempty" bson:"DMARC,omitempty"`
	Domain struct {
		ErrorCode string `json:"errorCode,omitempty" bson:"errorCode,omitempty"`
		Status    string `json:"status,omitempty" bson:"status,omitempty"`
	} `json:"Domain,omitempty" bson:"Domain,omitempty"`
	Spf struct {
		ErrorCode string `json:"errorCode,omitempty" bson:"errorCode,omitempty"`
		Status    string `json:"status,omitempty" bson:"status,omitempty"`
	} `json:"SPF,omitempty" bson:"SPF,omitempty"`
}

type AzureResourceVirtualMachineProfile struct {
	ExtensionProfile *struct {
		Extensions *[]struct {
			Name       string `json:"name,omitempty" bson:"name,omitempty"`
			Properties *struct {
				AutoUpgradeMinorVersion bool   `json:"autoUpgradeMinorVersion,omitempty" bson:"autoUpgradeMinorVersion,omitempty"`
				Publisher               string `json:"publisher,omitempty" bson:"publisher,omitempty"`
				Settings                *struct {
					DisableUu       string `json:"disable-uu,omitempty" bson:"disable-uu,omitempty"`
					EnableUu        string `json:"enable-uu,omitempty" bson:"enable-uu,omitempty"`
					NodeExporterTls string `json:"node-exporter-tls,omitempty" bson:"node-exporter-tls,omitempty"`
				} `json:"settings,omitempty" bson:"settings,omitempty"`
				SuppressFailures   bool   `json:"suppressFailures,omitempty" bson:"suppressFailures,omitempty"`
				Type               string `json:"type,omitempty" bson:"type,omitempty"`
				TypeHandlerVersion string `json:"typeHandlerVersion,omitempty" bson:"typeHandlerVersion,omitempty"`
			} `json:"properties,omitempty" bson:"properties,omitempty"`
		} `json:"extensions,omitempty" bson:"extensions,omitempty"`
		ExtensionsTimeBudget string `json:"extensionsTimeBudget,omitempty" bson:"extensionsTimeBudget,omitempty"`
	} `json:"extensionProfile,omitempty" bson:"extensionProfile,omitempty"`
	NetworkProfile *struct {
		NetworkInterfaceConfigurations *[]struct {
			Name       string `json:"name,omitempty" bson:"name,omitempty"`
			Properties *struct {
				DisableTcpStateTracking bool `json:"disableTcpStateTracking,omitempty" bson:"disableTcpStateTracking,omitempty"`
				DnsSettings             struct {
					DnsServers []any `json:"dnsServers,omitempty" bson:"dnsServers,omitempty"`
				} `json:"dnsSettings,omitempty" bson:"dnsSettings,omitempty"`
				EnableAcceleratedNetworking bool `json:"enableAcceleratedNetworking,omitempty" bson:"enableAcceleratedNetworking,omitempty"`
				EnableIpForwarding          bool `json:"enableIPForwarding,omitempty" bson:"enableIPForwarding,omitempty"`
				IpConfigurations            *[]struct {
					Name       string `json:"name,omitempty" bson:"name,omitempty"`
					Properties *struct {
						LoadBalancerBackendAddressPools *[]struct {
							ID string `json:"id,omitempty" bson:"id,omitempty"`
						} `json:"loadBalancerBackendAddressPools,omitempty" bson:"loadBalancerBackendAddressPools,omitempty"`
						Primary                 bool   `json:"primary,omitempty" bson:"primary,omitempty"`
						PrivateIpAddressVersion string `json:"privateIPAddressVersion,omitempty" bson:"privateIPAddressVersion,omitempty"`
						Subnet                  *struct {
							ID string `json:"id,omitempty" bson:"id,omitempty"`
						} `json:"subnet,omitempty" bson:"subnet,omitempty"`
					} `json:"properties,omitempty" bson:"properties,omitempty"`
				} `json:"ipConfigurations,omitempty" bson:"ipConfigurations,omitempty"`
				NetworkSecurityGroup *struct {
					ID string `json:"id,omitempty" bson:"id,omitempty"`
				} `json:"networkSecurityGroup,omitempty" bson:"networkSecurityGroup,omitempty"`
				Primary bool `json:"primary,omitempty" bson:"primary,omitempty"`
			} `json:"properties,omitempty" bson:"properties,omitempty"`
		} `json:"networkInterfaceConfigurations,omitempty" bson:"networkInterfaceConfigurations,omitempty"`
	} `json:"networkProfile,omitempty" bson:"networkProfile,omitempty"`
	OSProfile *struct {
		AdminUsername            string `json:"adminUsername,omitempty" bson:"adminUsername,omitempty"`
		AllowExtensionOperations bool   `json:"allowExtensionOperations,omitempty" bson:"allowExtensionOperations,omitempty"`
		ComputerNamePrefix       string `json:"computerNamePrefix,omitempty" bson:"computerNamePrefix,omitempty"`
		LinuxConfiguration       *struct {
			DisablePasswordAuthentication bool `json:"disablePasswordAuthentication,omitempty" bson:"disablePasswordAuthentication,omitempty"`
			EnableVmAgentPlatformUpdates  bool `json:"enableVMAgentPlatformUpdates,omitempty" bson:"enableVMAgentPlatformUpdates,omitempty"`
			ProvisionVmAgent              bool `json:"provisionVMAgent,omitempty" bson:"provisionVMAgent,omitempty"`
			SSH                           *struct {
				PublicKeys *[]struct {
					KeyData string `json:"keyData,omitempty" bson:"keyData,omitempty"`
					Path    string `json:"path,omitempty" bson:"path,omitempty"`
				} `json:"publicKeys,omitempty" bson:"publicKeys,omitempty"`
			} `json:"ssh,omitempty" bson:"ssh,omitempty"`
		} `json:"linuxConfiguration,omitempty" bson:"linuxConfiguration,omitempty"`
		RequireGuestProvisionSignal bool  `json:"requireGuestProvisionSignal,omitempty" bson:"requireGuestProvisionSignal,omitempty"`
		Secrets                     []any `json:"secrets,omitempty" bson:"secrets,omitempty"`
	} `json:"osProfile,omitempty" bson:"osProfile,omitempty"`
	StorageProfile *struct {
		DiskControllerType string `json:"diskControllerType,omitempty" bson:"diskControllerType,omitempty"`
		ImageReference     *struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"imageReference,omitempty" bson:"imageReference,omitempty"`
		OSDisk *struct {
			Caching      string  `json:"caching,omitempty" bson:"caching,omitempty"`
			CreateOption string  `json:"createOption,omitempty" bson:"createOption,omitempty"`
			DiskSizeGb   float64 `json:"diskSizeGB,omitempty" bson:"diskSizeGB,omitempty"`
			ManagedDisk  *struct {
				StorageAccountType string `json:"storageAccountType,omitempty" bson:"storageAccountType,omitempty"`
			} `json:"managedDisk,omitempty" bson:"managedDisk,omitempty"`
			OSType string `json:"osType,omitempty" bson:"osType,omitempty"`
		} `json:"osDisk,omitempty" bson:"osDisk,omitempty"`
	} `json:"storageProfile,omitempty" bson:"storageProfile,omitempty"`
	TimeCreated string `json:"timeCreated,omitempty" bson:"timeCreated,omitempty"`
}

type AzureResourceVirtualNetworkPeering struct {
	Etag       string `json:"etag,omitempty" bson:"etag,omitempty"`
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		AllowForwardedTraffic     bool   `json:"allowForwardedTraffic,omitempty" bson:"allowForwardedTraffic,omitempty"`
		AllowGatewayTransit       bool   `json:"allowGatewayTransit,omitempty" bson:"allowGatewayTransit,omitempty"`
		AllowVirtualNetworkAccess bool   `json:"allowVirtualNetworkAccess,omitempty" bson:"allowVirtualNetworkAccess,omitempty"`
		DoNotVerifyRemoteGateways bool   `json:"doNotVerifyRemoteGateways,omitempty" bson:"doNotVerifyRemoteGateways,omitempty"`
		PeerCompleteVnets         bool   `json:"peerCompleteVnets,omitempty" bson:"peerCompleteVnets,omitempty"`
		PeeringState              string `json:"peeringState,omitempty" bson:"peeringState,omitempty"`
		PeeringSyncLevel          string `json:"peeringSyncLevel,omitempty" bson:"peeringSyncLevel,omitempty"`
		ProvisioningState         string `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
		RemoteAddressSpace        struct {
			AddressPrefixes []string `json:"addressPrefixes,omitempty" bson:"addressPrefixes,omitempty"`
		} `json:"remoteAddressSpace,omitempty" bson:"remoteAddressSpace,omitempty"`
		RemoteGateways []struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"remoteGateways,omitempty" bson:"remoteGateways,omitempty"`
		RemoteVirtualNetwork struct {
			ID string `json:"id,omitempty" bson:"id,omitempty"`
		} `json:"remoteVirtualNetwork,omitempty" bson:"remoteVirtualNetwork,omitempty"`
		RemoteVirtualNetworkAddressSpace struct {
			AddressPrefixes []string `json:"addressPrefixes,omitempty" bson:"addressPrefixes,omitempty"`
		} `json:"remoteVirtualNetworkAddressSpace,omitempty" bson:"remoteVirtualNetworkAddressSpace,omitempty"`
		ResourceGuid      string            `json:"resourceGuid,omitempty" bson:"resourceGuid,omitempty"`
		RouteServiceVips  map[string]string `json:"routeServiceVips,omitempty" bson:"routeServiceVips,omitempty"`
		UseRemoteGateways bool              `json:"useRemoteGateways,omitempty" bson:"useRemoteGateways,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type AzureResourceVnetConfiguration struct {
	DockerBridgeCidr       any    `json:"dockerBridgeCidr,omitempty" bson:"dockerBridgeCidr,omitempty"`
	InfrastructureSubnetID string `json:"infrastructureSubnetId,omitempty" bson:"infrastructureSubnetId,omitempty"`
	Internal               bool   `json:"internal,omitempty" bson:"internal,omitempty"`
	PlatformReservedCidr   any    `json:"platformReservedCidr,omitempty" bson:"platformReservedCidr,omitempty"`
	PlatformReservedDnsIp  any    `json:"platformReservedDnsIP,omitempty" bson:"platformReservedDnsIP,omitempty"`
}

type AzureResourceVpnClientIpsecPolicy struct {
	DhGroup             string  `json:"dhGroup,omitempty" bson:"dhGroup,omitempty"`
	IkeEncryption       string  `json:"ikeEncryption,omitempty" bson:"ikeEncryption,omitempty"`
	IkeIntegrity        string  `json:"ikeIntegrity,omitempty" bson:"ikeIntegrity,omitempty"`
	IpsecEncryption     string  `json:"ipsecEncryption,omitempty" bson:"ipsecEncryption,omitempty"`
	IpsecIntegrity      string  `json:"ipsecIntegrity,omitempty" bson:"ipsecIntegrity,omitempty"`
	PfsGroup            string  `json:"pfsGroup,omitempty" bson:"pfsGroup,omitempty"`
	SaDataSizeKilobytes float64 `json:"saDataSizeKilobytes,omitempty" bson:"saDataSizeKilobytes,omitempty"`
	SaLifeTimeSeconds   float64 `json:"saLifeTimeSeconds,omitempty" bson:"saLifeTimeSeconds,omitempty"`
}

type AzureResourceWebhookReceivers struct {
	IdentifierURI        any    `json:"identifierUri,omitempty" bson:"identifierUri,omitempty"`
	Name                 string `json:"name,omitempty" bson:"name,omitempty"`
	ObjectID             any    `json:"objectId,omitempty" bson:"objectId,omitempty"`
	ServiceURI           string `json:"serviceUri,omitempty" bson:"serviceUri,omitempty"`
	TenantID             any    `json:"tenantId,omitempty" bson:"tenantId,omitempty"`
	UseAadAuth           bool   `json:"useAadAuth,omitempty" bson:"useAadAuth,omitempty"`
	UseCommonAlertSchema bool   `json:"useCommonAlertSchema,omitempty" bson:"useCommonAlertSchema,omitempty"`
}

type AzureResourceWriteLocations struct {
	DocumentEndpoint  string  `json:"documentEndpoint,omitempty" bson:"documentEndpoint,omitempty"`
	FailoverPriority  float64 `json:"failoverPriority,omitempty" bson:"failoverPriority,omitempty"`
	ID                string  `json:"id,omitempty" bson:"id,omitempty"`
	IsZoneRedundant   bool    `json:"isZoneRedundant,omitempty" bson:"isZoneRedundant,omitempty"`
	LocationName      string  `json:"locationName,omitempty" bson:"locationName,omitempty"`
	ProvisioningState string  `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
}

type AzureResourceAzureMonitorProfile struct {
	Metrics *struct {
		Enabled          bool `json:"enabled,omitempty" bson:"enabled,omitempty"`
		KubeStateMetrics struct {
			MetricAnnotationsAllowList string `json:"metricAnnotationsAllowList,omitempty" bson:"metricAnnotationsAllowList,omitempty"`
			MetricLabelsAllowlist      string `json:"metricLabelsAllowlist,omitempty" bson:"metricLabelsAllowlist,omitempty"`
		} `json:"kubeStateMetrics,omitempty" bson:"kubeStateMetrics,omitempty"`
	} `json:"metrics,omitempty" bson:"metrics,omitempty"`
}

type AzureResourceAccessPolicy struct {
	ObjectID    string `json:"objectId,omitempty" bson:"objectId,omitempty"`
	Permissions struct {
		Certificates []string `json:"certificates,omitempty" bson:"certificates,omitempty"`
		Keys         []string `json:"keys,omitempty" bson:"keys,omitempty"`
		Secrets      []string `json:"secrets,omitempty" bson:"secrets,omitempty"`
		Storage      []string `json:"storage,omitempty" bson:"storage,omitempty"`
	} `json:"permissions,omitempty" bson:"permissions,omitempty"`
	TenantID string `json:"tenantId,omitempty" bson:"tenantId,omitempty"`
}

type AzureResourceCondition struct {
	AllOf []struct {
		AnyOf []struct {
			Equals string `json:"equals,omitempty" bson:"equals,omitempty"`
			Field  string `json:"field,omitempty" bson:"field,omitempty"`
		} `json:"anyOf,omitempty" bson:"anyOf,omitempty"`
		ContainsAny []string `json:"containsAny,omitempty" bson:"containsAny,omitempty"`
		Equals      string   `json:"equals,omitempty" bson:"equals,omitempty"`
		Field       string   `json:"field,omitempty" bson:"field,omitempty"`
	} `json:"allOf,omitempty" bson:"allOf,omitempty"`
}

type AzureResourceCustomize struct {
	Filters             []string `json:"filters,omitempty" bson:"filters,omitempty"`
	Name                string   `json:"name,omitempty" bson:"name,omitempty"`
	RestartCheckCommand string   `json:"restartCheckCommand,omitempty" bson:"restartCheckCommand,omitempty"`
	RestartCommand      string   `json:"restartCommand,omitempty" bson:"restartCommand,omitempty"`
	RestartTimeout      string   `json:"restartTimeout,omitempty" bson:"restartTimeout,omitempty"`
	RunAsSystem         bool     `json:"runAsSystem,omitempty" bson:"runAsSystem,omitempty"`
	RunElevated         bool     `json:"runElevated,omitempty" bson:"runElevated,omitempty"`
	ScriptURI           string   `json:"scriptUri,omitempty" bson:"scriptUri,omitempty"`
	SearchCriteria      string   `json:"searchCriteria,omitempty" bson:"searchCriteria,omitempty"`
	Sha256Checksum      string   `json:"sha256Checksum,omitempty" bson:"sha256Checksum,omitempty"`
	Type                string   `json:"type,omitempty" bson:"type,omitempty"`
	UpdateLimit         float64  `json:"updateLimit,omitempty" bson:"updateLimit,omitempty"`
}

type AzureResourceDataFlows struct {
	Destinations []string `json:"destinations,omitempty" bson:"destinations,omitempty"`
	OutputStream string   `json:"outputStream,omitempty" bson:"outputStream,omitempty"`
	Streams      []string `json:"streams,omitempty" bson:"streams,omitempty"`
	TransformKql string   `json:"transformKql,omitempty" bson:"transformKql,omitempty"`
}

type AzureResourceFactoryStatistics struct {
	FactorySizeInGbUnits           float64 `json:"factorySizeInGbUnits,omitempty" bson:"factorySizeInGbUnits,omitempty"`
	MaxAllowedFactorySizeInGbUnits float64 `json:"maxAllowedFactorySizeInGbUnits,omitempty" bson:"maxAllowedFactorySizeInGbUnits,omitempty"`
	MaxAllowedResourceCount        float64 `json:"maxAllowedResourceCount,omitempty" bson:"maxAllowedResourceCount,omitempty"`
	TotalResourceCount             float64 `json:"totalResourceCount,omitempty" bson:"totalResourceCount,omitempty"`
}

type AzureResourceFailoverPolicy struct {
	FailoverPriority float64 `json:"failoverPriority,omitempty" bson:"failoverPriority,omitempty"`
	ID               string  `json:"id,omitempty" bson:"id,omitempty"`
	LocationName     string  `json:"locationName,omitempty" bson:"locationName,omitempty"`
}

type AzureResourceFeatureSettings struct {
	CrossSubscriptionRestoreSettings struct {
		State string `json:"state,omitempty" bson:"state,omitempty"`
	} `json:"crossSubscriptionRestoreSettings,omitempty" bson:"crossSubscriptionRestoreSettings,omitempty"`
}

type AzureResourceLastRunStatus struct {
	EndTime     string `json:"endTime,omitempty" bson:"endTime,omitempty"`
	Message     string `json:"message,omitempty" bson:"message,omitempty"`
	RunState    string `json:"runState,omitempty" bson:"runState,omitempty"`
	RunSubState string `json:"runSubState,omitempty" bson:"runSubState,omitempty"`
	StartTime   string `json:"startTime,omitempty" bson:"startTime,omitempty"`
}

type AzureResourceMountTargets struct {
	FileSystemID  string `json:"fileSystemId,omitempty" bson:"fileSystemId,omitempty"`
	IpAddress     string `json:"ipAddress,omitempty" bson:"ipAddress,omitempty"`
	MountTargetID string `json:"mountTargetId,omitempty" bson:"mountTargetId,omitempty"`
	SmbServerFqdn string `json:"smbServerFqdn,omitempty" bson:"smbServerFqdn,omitempty"`
}

type AzureResourceSecondaryEndpoints struct {
	Blob  string `json:"blob,omitempty" bson:"blob,omitempty"`
	Dfs   string `json:"dfs,omitempty" bson:"dfs,omitempty"`
	Queue string `json:"queue,omitempty" bson:"queue,omitempty"`
	Table string `json:"table,omitempty" bson:"table,omitempty"`
	Web   string `json:"web,omitempty" bson:"web,omitempty"`
}

type AzureResourceServiceProviderProperties struct {
	BandwidthInMbps     float64 `json:"bandwidthInMbps,omitempty" bson:"bandwidthInMbps,omitempty"`
	PeeringLocation     string  `json:"peeringLocation,omitempty" bson:"peeringLocation,omitempty"`
	ServiceProviderName string  `json:"serviceProviderName,omitempty" bson:"serviceProviderName,omitempty"`
}

type AzureResourceVolumeBackups struct {
	BackupsCount     float64 `json:"backupsCount,omitempty" bson:"backupsCount,omitempty"`
	PolicyEnabled    bool    `json:"policyEnabled,omitempty" bson:"policyEnabled,omitempty"`
	VolumeName       string  `json:"volumeName,omitempty" bson:"volumeName,omitempty"`
	VolumeResourceID string  `json:"volumeResourceId,omitempty" bson:"volumeResourceId,omitempty"`
}

type AzureResourceWeeklySchedule struct {
	Day             string  `json:"day,omitempty" bson:"day,omitempty"`
	Hour            float64 `json:"hour,omitempty" bson:"hour,omitempty"`
	Minute          float64 `json:"minute,omitempty" bson:"minute,omitempty"`
	SnapshotsToKeep float64 `json:"snapshotsToKeep,omitempty" bson:"snapshotsToKeep,omitempty"`
}

type AzureResourceWorkbookTemplate struct {
	ID struct {
		Public string `json:"public,omitempty" bson:"public,omitempty"`
	} `json:"id,omitempty" bson:"id,omitempty"`
	Source string `json:"source,omitempty" bson:"source,omitempty"`
}

type AzureResourceWorkspaceCapping struct {
	DailyQuotaGb        float64 `json:"dailyQuotaGb,omitempty" bson:"dailyQuotaGb,omitempty"`
	DataIngestionStatus string  `json:"dataIngestionStatus,omitempty" bson:"dataIngestionStatus,omitempty"`
	QuotaNextResetTime  string  `json:"quotaNextResetTime,omitempty" bson:"quotaNextResetTime,omitempty"`
}

type AzureResourceProperties struct {
	// Other any `json:",unknown" bson:",unknown"`
	Other map[string]any `json:",unknown" bson:",unknown"`
	// Other any `json:"other" bson:"other"`
	// AccountURL                      string         `json:"AccountURL,omitempty" bson:"AccountURL,omitempty"`
	AppID         string `json:"AppId,omitempty" bson:"AppId,omitempty"`
	ApplicationID string `json:"ApplicationId,omitempty" bson:"ApplicationId,omitempty"`
	// ApplicationType                 string         `json:"Application_Type,omitempty" bson:"Application_Type,omitempty"`
	// ConnectionString                string         `json:"ConnectionString,omitempty" bson:"ConnectionString,omitempty"`
	// CreationDateOther               string         `json:"CreationDate,omitempty" bson:"CreationDate,omitempty"`
	// DisableIpMasking                bool           `json:"DisableIpMasking,omitempty" bson:"DisableIpMasking,omitempty"`
	// DisableLocalAuthOther           bool           `json:"DisableLocalAuth,omitempty" bson:"DisableLocalAuth,omitempty"`
	// EnabledAPITypes                 string         `json:"EnabledApiTypes,omitempty" bson:"EnabledApiTypes,omitempty"`
	// FlowType                        string         `json:"Flow_Type,omitempty" bson:"Flow_Type,omitempty"`
	// ForceCustomerStorageForProfiler bool           `json:"ForceCustomerStorageForProfiler,omitempty" bson:"ForceCustomerStorageForProfiler,omitempty"`
	// IngestionMode                   string         `json:"IngestionMode,omitempty" bson:"IngestionMode,omitempty"`
	// InstrumentationKey              string         `json:"InstrumentationKey,omitempty" bson:"InstrumentationKey,omitempty"`
	// LastOwnershipUpdateTime         string         `json:"LastOwnershipUpdateTime,omitempty" bson:"LastOwnershipUpdateTime,omitempty"`
	// LinkedStorages                  *struct {
	// 	ServiceProfilerLinkedStorage string `json:"ServiceProfilerLinkedStorage,omitempty" bson:"ServiceProfilerLinkedStorage,omitempty"`
	// } `json:"LinkedStorages,omitempty" bson:"LinkedStorages,omitempty"`
	// NameOther                       string                                    `json:"Name,omitempty" bson:"Name,omitempty"`
	// PrivateLinkScopedResourcesOther []*AzureResourcePrivateLinkScopedResource `json:"PrivateLinkScopedResources,omitempty" bson:"PrivateLinkScopedResources,omitempty"`
	// RegistrationURL                 string                                    `json:"RegistrationUrl,omitempty" bson:"RegistrationUrl,omitempty"`
	// RequestSource                   string                                    `json:"Request_Source,omitempty" bson:"Request_Source,omitempty"`
	// Retention                       string                                    `json:"Retention,omitempty" bson:"Retention,omitempty"`
	// RetentionInDaysOther            float64                                   `json:"RetentionInDays,omitempty" bson:"RetentionInDays,omitempty"`
	// RuntimeConfiguration            *AzureResourceRuntimeConfiguration        `json:"RuntimeConfiguration,omitempty" bson:"RuntimeConfiguration,omitempty"`
	// SamplingPercentage              *float64                                  `json:"SamplingPercentage,omitempty" bson:"SamplingPercentage,omitempty"`
	// TenantID                        string                                    `json:"TenantId,omitempty" bson:"TenantId,omitempty"`
	// Ver                             string                                    `json:"Ver,omitempty" bson:"Ver,omitempty"`
	// WorkspaceResourceID             string                                    `json:"WorkspaceResourceId,omitempty" bson:"WorkspaceResourceId,omitempty"`
	// AadAuthenticationParameters     *AzureResourceAadAuthenticationParameters `json:"aadAuthenticationParameters,omitempty" bson:"aadAuthenticationParameters,omitempty"`
	// AadProfile                      *AzureResourceAadProfile                  `json:"aadProfile,omitempty" bson:"aadProfile,omitempty"`
	// AccessEndpoint                  string                                    `json:"accessEndpoint,omitempty" bson:"accessEndpoint,omitempty"`
	// AccessModeSettings              *AzureResourceAccessModeSettings          `json:"accessModeSettings,omitempty" bson:"accessModeSettings,omitempty"`
	// AccessPolicies                  []*AzureResourceAccessPolicy              `json:"accessPolicies,omitempty" bson:"accessPolicies,omitempty"`
	// AccessTier                      string                                    `json:"accessTier,omitempty" bson:"accessTier,omitempty"`
	// AccountEndpoint                 string                                    `json:"accountEndpoint,omitempty" bson:"accountEndpoint,omitempty"`
	// Actions                         any                                       `json:"actions,omitempty" bson:"actions,omitempty"`
	// ActiveActive                    bool                                      `json:"activeActive,omitempty" bson:"activeActive,omitempty"`
	// ActiveDirectories               []*AzureResourceActiveDirectory           `json:"activeDirectories,omitempty" bson:"activeDirectories,omitempty"`
	// ActiveJobAndJobScheduleQuota    float64                                   `json:"activeJobAndJobScheduleQuota,omitempty" bson:"activeJobAndJobScheduleQuota,omitempty"`
	// AdditionalCapabilities          map[string]bool                           `json:"additionalCapabilities,omitempty" bson:"additionalCapabilities,omitempty"`
	// AdditionalProperties            map[string]string                         `json:"additionalProperties,omitempty" bson:"additionalProperties,omitempty"`
	// AddonProfiles                   *AzureResourceAddonProfiles               `json:"addonProfiles,omitempty" bson:"addonProfiles,omitempty"`
	// AddressPrefix                   string                                    `json:"addressPrefix,omitempty" bson:"addressPrefix,omitempty"`
	// AddressSpace                    *struct {
	// 	AddressPrefixes []string `json:"addressPrefixes,omitempty" bson:"addressPrefixes,omitempty"`
	// } `json:"addressSpace,omitempty" bson:"addressSpace,omitempty"`
	// AdminEnabled                   bool                             `json:"adminEnabled,omitempty" bson:"adminEnabled,omitempty"`
	// AdminRuntimeSiteName           any                              `json:"adminRuntimeSiteName,omitempty" bson:"adminRuntimeSiteName,omitempty"`
	// AdminSiteName                  any                              `json:"adminSiteName,omitempty" bson:"adminSiteName,omitempty"`
	// AdminUserEnabled               bool                             `json:"adminUserEnabled,omitempty" bson:"adminUserEnabled,omitempty"`
	// AdministratorLogin             string                           `json:"administratorLogin,omitempty" bson:"administratorLogin,omitempty"`
	// Administrators                 *AzureResourceAdministrators     `json:"administrators,omitempty" bson:"administrators,omitempty"`
	// AfdEnabled                     bool                             `json:"afdEnabled,omitempty" bson:"afdEnabled,omitempty"`
	// AgentPoolProfiles              []*AzureResourceAgentPoolProfile `json:"agentPoolProfiles,omitempty" bson:"agentPoolProfiles,omitempty"`
	// AllocationDate                 string                           `json:"allocationDate,omitempty" bson:"allocationDate,omitempty"`
	// AllowBlobPublicAccess          bool                             `json:"allowBlobPublicAccess,omitempty" bson:"allowBlobPublicAccess,omitempty"`
	// AllowBranchToBranchTraffic     bool                             `json:"allowBranchToBranchTraffic,omitempty" bson:"allowBranchToBranchTraffic,omitempty"`
	// AllowClassicOperations         bool                             `json:"allowClassicOperations,omitempty" bson:"allowClassicOperations,omitempty"`
	// AllowCrossTenantReplication    bool                             `json:"allowCrossTenantReplication,omitempty" bson:"allowCrossTenantReplication,omitempty"`
	// AllowGlobalReach               bool                             `json:"allowGlobalReach,omitempty" bson:"allowGlobalReach,omitempty"`
	// AllowNonVirtualWanTraffic      bool                             `json:"allowNonVirtualWanTraffic,omitempty" bson:"allowNonVirtualWanTraffic,omitempty"`
	// AllowPort25Out                 bool                             `json:"allowPort25Out,omitempty" bson:"allowPort25Out,omitempty"`
	// AllowRemoteVnetTraffic         bool                             `json:"allowRemoteVnetTraffic,omitempty" bson:"allowRemoteVnetTraffic,omitempty"`
	// AllowSharedKeyAccess           bool                             `json:"allowSharedKeyAccess,omitempty" bson:"allowSharedKeyAccess,omitempty"`
	// AllowVirtualWanTraffic         bool                             `json:"allowVirtualWanTraffic,omitempty" bson:"allowVirtualWanTraffic,omitempty"`
	// AllowVnetToVnetTraffic         bool                             `json:"allowVnetToVnetTraffic,omitempty" bson:"allowVnetToVnetTraffic,omitempty"`
	// AllowedAuthenticationModes     []string                         `json:"allowedAuthenticationModes,omitempty" bson:"allowedAuthenticationModes,omitempty"`
	// AllowedCopyScope               string                           `json:"allowedCopyScope,omitempty" bson:"allowedCopyScope,omitempty"`
	// AlternativeParameterValues     *map[string]string               `json:"alternativeParameterValues,omitempty" bson:"alternativeParameterValues,omitempty"`
	// AnalyticalStorageConfiguration *struct {
	// 	SchemaType string `json:"schemaType,omitempty" bson:"schemaType,omitempty"`
	// } `json:"analyticalStorageConfiguration,omitempty" bson:"analyticalStorageConfiguration,omitempty"`
	// AnonymousPullEnabled       bool                                 `json:"anonymousPullEnabled,omitempty" bson:"anonymousPullEnabled,omitempty"`
	// API                        *AzureResourceAPI                    `json:"api,omitempty" bson:"api,omitempty"`
	// APIServerAccessProfile     *AzureResourceAPIServerAccessProfile `json:"apiServerAccessProfile,omitempty" bson:"apiServerAccessProfile,omitempty"`
	// AppInsightsConfiguration   any                                  `json:"appInsightsConfiguration,omitempty" bson:"appInsightsConfiguration,omitempty"`
	// AppLogsConfiguration       *AzureResourceAppLogsConfiguration   `json:"appLogsConfiguration,omitempty" bson:"appLogsConfiguration,omitempty"`
	// ApplicationRuleCollections []any                                `json:"applicationRuleCollections,omitempty" bson:"applicationRuleCollections,omitempty"`
	// Architecture               string                               `json:"architecture,omitempty" bson:"architecture,omitempty"`
	// ArmRoleReceivers           []struct {
	// 	Name                 string `json:"name,omitempty" bson:"name,omitempty"`
	// 	RoleID               string `json:"roleId,omitempty" bson:"roleId,omitempty"`
	// 	UseCommonAlertSchema bool   `json:"useCommonAlertSchema,omitempty" bson:"useCommonAlertSchema,omitempty"`
	// } `json:"armRoleReceivers,omitempty" bson:"armRoleReceivers,omitempty"`
	// AuthConfig *struct {
	// 	ActiveDirectoryAuth string `json:"activeDirectoryAuth,omitempty" bson:"activeDirectoryAuth,omitempty"`
	// 	PasswordAuth        string `json:"passwordAuth,omitempty" bson:"passwordAuth,omitempty"`
	// } `json:"authConfig,omitempty" bson:"authConfig,omitempty"`
	// AuthenticatedUser *struct {
	// 	Name string `json:"name,omitempty" bson:"name,omitempty"`
	// } `json:"authenticatedUser,omitempty" bson:"authenticatedUser,omitempty"`
	// AuthenticationType                   string                               `json:"authenticationType,omitempty" bson:"authenticationType,omitempty"`
	// Authorizations                       []*AzureResourceAuthorization        `json:"authorizations,omitempty" bson:"authorizations,omitempty"`
	// AutoCreateTopicWithFirstSubscription bool                                 `json:"autoCreateTopicWithFirstSubscription,omitempty" bson:"autoCreateTopicWithFirstSubscription,omitempty"`
	// AutoDeleteTopicWithLastSubscription  bool                                 `json:"autoDeleteTopicWithLastSubscription,omitempty" bson:"autoDeleteTopicWithLastSubscription,omitempty"`
	// AutoGeneratedDomainNameLabelScope    any                                  `json:"autoGeneratedDomainNameLabelScope,omitempty" bson:"autoGeneratedDomainNameLabelScope,omitempty"`
	// AutoMitigate                         bool                                 `json:"autoMitigate,omitempty" bson:"autoMitigate,omitempty"`
	// AutoPauseDelay                       float64                              `json:"autoPauseDelay,omitempty" bson:"autoPauseDelay,omitempty"`
	// AutoScaleConfiguration               *AzureResourceAutoScaleConfiguration `json:"autoScaleConfiguration,omitempty" bson:"autoScaleConfiguration,omitempty"`
	// AutoScalerProfile                    *AzureResourceAutoScalerProfile      `json:"autoScalerProfile,omitempty" bson:"autoScalerProfile,omitempty"`
	// AutoUpgradeMinorVersion              bool                                 `json:"autoUpgradeMinorVersion,omitempty" bson:"autoUpgradeMinorVersion,omitempty"`
	// AutoUpgradeProfile                   *struct {
	// 	UpgradeChannel string `json:"upgradeChannel,omitempty" bson:"upgradeChannel,omitempty"`
	// } `json:"autoUpgradeProfile,omitempty" bson:"autoUpgradeProfile,omitempty"`
	// AutomationHybridServiceURL string `json:"automationHybridServiceUrl,omitempty" bson:"automationHybridServiceUrl,omitempty"`
	// AutomationRunbookReceivers []any  `json:"automationRunbookReceivers,omitempty" bson:"automationRunbookReceivers,omitempty"`
	// AuxiliaryMode              string `json:"auxiliaryMode,omitempty" bson:"auxiliaryMode,omitempty"`
	// AuxiliarySku               string `json:"auxiliarySku,omitempty" bson:"auxiliarySku,omitempty"`
	// AvailabilitySet            *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"availabilitySet,omitempty" bson:"availabilitySet,omitempty"`
	// AvailabilityState                     string `json:"availabilityState,omitempty" bson:"availabilityState,omitempty"`
	// AvailabilityZone                      string `json:"availabilityZone,omitempty" bson:"availabilityZone,omitempty"`
	// AvsDataStore                          string `json:"avsDataStore,omitempty" bson:"avsDataStore,omitempty"`
	// AzureAppPushReceivers                 []any  `json:"azureAppPushReceivers,omitempty" bson:"azureAppPushReceivers,omitempty"`
	// AzureFilesIdentityBasedAuthentication *struct {
	// 	DirectoryServiceOptions string `json:"directoryServiceOptions,omitempty" bson:"directoryServiceOptions,omitempty"`
	// } `json:"azureFilesIdentityBasedAuthentication,omitempty" bson:"azureFilesIdentityBasedAuthentication,omitempty"`
	// AzureFirewall *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"azureFirewall,omitempty" bson:"azureFirewall,omitempty"`
	// AzureFunctionReceivers []any                              `json:"azureFunctionReceivers,omitempty" bson:"azureFunctionReceivers,omitempty"`
	// AzureMonitorProfile    *AzureResourceAzureMonitorProfile  `json:"azureMonitorProfile,omitempty" bson:"azureMonitorProfile,omitempty"`
	// AzurePortalFqdn        string                             `json:"azurePortalFQDN,omitempty" bson:"azurePortalFQDN,omitempty"`
	// BackendAddressPools    []*AzureResourceBackendAddressPool `json:"backendAddressPools,omitempty" bson:"backendAddressPools,omitempty"`
	// Backup                 *AzureResourceBackup               `json:"backup,omitempty" bson:"backup,omitempty"`
	// BackupPolicy           *AzureResourceBackupPolicy         `json:"backupPolicy,omitempty" bson:"backupPolicy,omitempty"`
	// BackupPolicyID         string                             `json:"backupPolicyId,omitempty" bson:"backupPolicyId,omitempty"`
	// BackupStorageVersion   string                             `json:"backupStorageVersion,omitempty" bson:"backupStorageVersion,omitempty"`
	// BandwidthInGbps        float64                            `json:"bandwidthInGbps,omitempty" bson:"bandwidthInGbps,omitempty"`
	// BareMetalServer        *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"bareMetalServer,omitempty" bson:"bareMetalServer,omitempty"`
	// BaremetalTenantID string `json:"baremetalTenantId,omitempty" bson:"baremetalTenantId,omitempty"`
	// BcdrSecurityLevel string `json:"bcdrSecurityLevel,omitempty" bson:"bcdrSecurityLevel,omitempty"`
	// BillingConfig     *struct {
	// 	BillingType           string `json:"billingType,omitempty" bson:"billingType,omitempty"`
	// 	EffectiveStartDateUtc string `json:"effectiveStartDateUtc,omitempty" bson:"effectiveStartDateUtc,omitempty"`
	// } `json:"billingConfig,omitempty" bson:"billingConfig,omitempty"`
	// BillingModel   string `json:"billingModel,omitempty" bson:"billingModel,omitempty"`
	// BillingProfile *struct {
	// 	MaxPrice float64 `json:"maxPrice,omitempty" bson:"maxPrice,omitempty"`
	// } `json:"billingProfile,omitempty" bson:"billingProfile,omitempty"`
	// BlockPathTraversal    bool                        `json:"blockPathTraversal,omitempty" bson:"blockPathTraversal,omitempty"`
	// BuildTimeoutInMinutes float64                     `json:"buildTimeoutInMinutes,omitempty" bson:"buildTimeoutInMinutes,omitempty"`
	// BuildVersion          any                         `json:"buildVersion,omitempty" bson:"buildVersion,omitempty"`
	// CallRateLimit         *AzureResourceCallRateLimit `json:"callRateLimit,omitempty" bson:"callRateLimit,omitempty"`
	// Capabilities          []struct {
	// 	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	// 	Value string `json:"value,omitempty" bson:"value,omitempty"`
	// } `json:"capabilities,omitempty" bson:"capabilities,omitempty"`
	// CatalogCollation         string `json:"catalogCollation,omitempty" bson:"catalogCollation,omitempty"`
	// Category                 string `json:"category,omitempty" bson:"category,omitempty"`
	// Cers                     any    `json:"cers,omitempty" bson:"cers,omitempty"`
	// ChangedTime              string `json:"changedTime,omitempty" bson:"changedTime,omitempty"`
	// ChildPolicies            []any  `json:"childPolicies,omitempty" bson:"childPolicies,omitempty"`
	// CircuitProvisioningState string `json:"circuitProvisioningState,omitempty" bson:"circuitProvisioningState,omitempty"`
	// Circuits                 []struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"circuits,omitempty" bson:"circuits,omitempty"`
	// ClientAffinityEnabled      bool   `json:"clientAffinityEnabled,omitempty" bson:"clientAffinityEnabled,omitempty"`
	// ClientAffinityProxyEnabled bool   `json:"clientAffinityProxyEnabled,omitempty" bson:"clientAffinityProxyEnabled,omitempty"`
	// ClientCertEnabled          bool   `json:"clientCertEnabled,omitempty" bson:"clientCertEnabled,omitempty"`
	// ClientCertExclusionPaths   any    `json:"clientCertExclusionPaths,omitempty" bson:"clientCertExclusionPaths,omitempty"`
	// ClientCertMode             string `json:"clientCertMode,omitempty" bson:"clientCertMode,omitempty"`
	// ClientID                   string `json:"clientId,omitempty" bson:"clientId,omitempty"`
	// CloningInfo                any    `json:"cloningInfo,omitempty" bson:"cloningInfo,omitempty"`
	// CloudConnectors            *struct {
	// 	AwsExternalID string `json:"awsExternalId,omitempty" bson:"awsExternalId,omitempty"`
	// } `json:"cloudConnectors,omitempty" bson:"cloudConnectors,omitempty"`
	// CloudID       string `json:"cloudId,omitempty" bson:"cloudId,omitempty"`
	// CloudServices []struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"cloudServices,omitempty" bson:"cloudServices,omitempty"`
	// Collation           string                      `json:"collation,omitempty" bson:"collation,omitempty"`
	// ComputeMode         *string                     `json:"computeMode,omitempty" bson:"computeMode,omitempty"`
	// Condition           *AzureResourceCondition     `json:"condition,omitempty" bson:"condition,omitempty"`
	// Configuration       *AzureResourceConfiguration `json:"configuration,omitempty" bson:"configuration,omitempty"`
	// ConfigurationAccess *struct {
	// 	Endpoint string `json:"endpoint,omitempty" bson:"endpoint,omitempty"`
	// } `json:"configurationAccess,omitempty" bson:"configurationAccess,omitempty"`
	// ConfigurationOverrides    *struct{} `json:"configurationOverrides,omitempty" bson:"configurationOverrides,omitempty"`
	// ConfigurationPolicyGroups []any     `json:"configurationPolicyGroups,omitempty" bson:"configurationPolicyGroups,omitempty"`
	// ConfigurationType         string    `json:"configurationType,omitempty" bson:"configurationType,omitempty"`
	// Configurations            []struct {
	// 	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	// 	Value string `json:"value,omitempty" bson:"value,omitempty"`
	// } `json:"configurations,omitempty" bson:"configurations,omitempty"`
	// ConnectionMode        string `json:"connectionMode,omitempty" bson:"connectionMode,omitempty"`
	// ConnectionMonitorType string `json:"connectionMonitorType,omitempty" bson:"connectionMonitorType,omitempty"`
	// ConnectionState       string `json:"connectionState,omitempty" bson:"connectionState,omitempty"`
	// ConnectionType        string `json:"connectionType,omitempty" bson:"connectionType,omitempty"`
	// ConsistencyPolicy     *struct {
	// 	DefaultConsistencyLevel string  `json:"defaultConsistencyLevel,omitempty" bson:"defaultConsistencyLevel,omitempty"`
	// 	MaxIntervalInSeconds    float64 `json:"maxIntervalInSeconds,omitempty" bson:"maxIntervalInSeconds,omitempty"`
	// 	MaxStalenessPrefix      float64 `json:"maxStalenessPrefix,omitempty" bson:"maxStalenessPrefix,omitempty"`
	// } `json:"consistencyPolicy,omitempty" bson:"consistencyPolicy,omitempty"`
	// ContainedResources        []string `json:"containedResources,omitempty" bson:"containedResources,omitempty"`
	// ContainerAllocationSubnet any      `json:"containerAllocationSubnet,omitempty" bson:"containerAllocationSubnet,omitempty"`
	// ContainerSize             float64  `json:"containerSize,omitempty" bson:"containerSize,omitempty"`
	// ContentAvailabilityState  string   `json:"contentAvailabilityState,omitempty" bson:"contentAvailabilityState,omitempty"`
	// CoolAccess                bool     `json:"coolAccess,omitempty" bson:"coolAccess,omitempty"`
	// Cors                      []any    `json:"cors,omitempty" bson:"cors,omitempty"`
	// CreateTenantProperties    *struct {
	// 	CountryCode string `json:"countryCode,omitempty" bson:"countryCode,omitempty"`
	// 	DisplayName string `json:"displayName,omitempty" bson:"displayName,omitempty"`
	// } `json:"createTenantProperties,omitempty" bson:"createTenantProperties,omitempty"`
	// CreateTime                     string                     `json:"createTime,omitempty" bson:"createTime,omitempty"`
	// CreatedAt                      string                     `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	// CreatedBy                      string                     `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	// CreatedByObjectID              string                     `json:"createdByObjectId,omitempty" bson:"createdByObjectId,omitempty"`
	// CreatedDate                    string                     `json:"createdDate,omitempty" bson:"createdDate,omitempty"`
	// CreatedTime                    string                     `json:"createdTime,omitempty" bson:"createdTime,omitempty"`
	// CreatedWithAPIVersion          string                     `json:"createdWithApiVersion,omitempty" bson:"createdWithApiVersion,omitempty"`
	// CreationData                   *AzureResourceCreationData `json:"creationData,omitempty" bson:"creationData,omitempty"`
	// CreationDate                   string                     `json:"creationDate,omitempty" bson:"creationDate,omitempty"`
	// CreationTime                   string                     `json:"creationTime,omitempty" bson:"creationTime,omitempty"`
	// CreationToken                  string                     `json:"creationToken,omitempty" bson:"creationToken,omitempty"`
	// Criteria                       *AzureResourceCriteria     `json:"criteria,omitempty" bson:"criteria,omitempty"`
	// Csrs                           []any                      `json:"csrs,omitempty" bson:"csrs,omitempty"`
	// CurrentBackupStorageRedundancy string                     `json:"currentBackupStorageRedundancy,omitempty" bson:"currentBackupStorageRedundancy,omitempty"`
	// CurrentKubernetesVersion       string                     `json:"currentKubernetesVersion,omitempty" bson:"currentKubernetesVersion,omitempty"`
	// CurrentNumberOfWorkers         float64                    `json:"currentNumberOfWorkers,omitempty" bson:"currentNumberOfWorkers,omitempty"`
	// CurrentServiceObjectiveName    string                     `json:"currentServiceObjectiveName,omitempty" bson:"currentServiceObjectiveName,omitempty"`
	// CurrentSku                     *struct {
	// 	Capacity float64 `json:"capacity,omitempty" bson:"capacity,omitempty"`
	// 	Family   string  `json:"family,omitempty" bson:"family,omitempty"`
	// 	Name     string  `json:"name,omitempty" bson:"name,omitempty"`
	// 	Tier     string  `json:"tier,omitempty" bson:"tier,omitempty"`
	// } `json:"currentSku,omitempty" bson:"currentSku,omitempty"`
	// CurrentWorkerSize   string  `json:"currentWorkerSize,omitempty" bson:"currentWorkerSize,omitempty"`
	// CurrentWorkerSizeID float64 `json:"currentWorkerSizeId,omitempty" bson:"currentWorkerSizeId,omitempty"`
	// CustomDnsConfigs    []struct {
	// 	Fqdn        string   `json:"fqdn,omitempty" bson:"fqdn,omitempty"`
	// 	IpAddresses []string `json:"ipAddresses,omitempty" bson:"ipAddresses,omitempty"`
	// } `json:"customDnsConfigs,omitempty" bson:"customDnsConfigs,omitempty"`
	// CustomDnsServers           []string                                `json:"customDnsServers,omitempty" bson:"customDnsServers,omitempty"`
	// CustomDomainConfiguration  *AzureResourceCustomDomainConfiguration `json:"customDomainConfiguration,omitempty" bson:"customDomainConfiguration,omitempty"`
	// CustomDomainVerificationID string                                  `json:"customDomainVerificationId,omitempty" bson:"customDomainVerificationId,omitempty"`
	// CustomNetworkInterfaceName string                                  `json:"customNetworkInterfaceName,omitempty" bson:"customNetworkInterfaceName,omitempty"`
	// CustomParameterValues      *struct{}                               `json:"customParameterValues,omitempty" bson:"customParameterValues,omitempty"`
	// CustomSubDomainName        string                                  `json:"customSubDomainName,omitempty" bson:"customSubDomainName,omitempty"`
	// CustomerID                 string                                  `json:"customerId,omitempty" bson:"customerId,omitempty"`
	// Customize                  []*AzureResourceCustomize               `json:"customize,omitempty" bson:"customize,omitempty"`
	// DailyBackupsToKeep         float64                                 `json:"dailyBackupsToKeep,omitempty" bson:"dailyBackupsToKeep,omitempty"`
	// DailyMemoryTimeQuota       float64                                 `json:"dailyMemoryTimeQuota,omitempty" bson:"dailyMemoryTimeQuota,omitempty"`
	// DailyRecurrence            *struct {
	// 	Time string `json:"time,omitempty" bson:"time,omitempty"`
	// } `json:"dailyRecurrence,omitempty" bson:"dailyRecurrence,omitempty"`
	// DailySchedule *struct {
	// 	Hour            float64 `json:"hour,omitempty" bson:"hour,omitempty"`
	// 	Minute          float64 `json:"minute,omitempty" bson:"minute,omitempty"`
	// 	SnapshotsToKeep float64 `json:"snapshotsToKeep,omitempty" bson:"snapshotsToKeep,omitempty"`
	// } `json:"dailySchedule,omitempty" bson:"dailySchedule,omitempty"`
	// DaprAiConnectionString   any `json:"daprAIConnectionString,omitempty" bson:"daprAIConnectionString,omitempty"`
	// DaprAiInstrumentationKey any `json:"daprAIInstrumentationKey,omitempty" bson:"daprAIInstrumentationKey,omitempty"`
	// DaprConfig               any `json:"daprConfig,omitempty" bson:"daprConfig,omitempty"`
	// DaprConfiguration        *struct {
	// 	Version string `json:"version,omitempty" bson:"version,omitempty"`
	// } `json:"daprConfiguration,omitempty" bson:"daprConfiguration,omitempty"`
	// DataAccessAuthMode       string `json:"dataAccessAuthMode,omitempty" bson:"dataAccessAuthMode,omitempty"`
	// DataCollectionEndpointID string `json:"dataCollectionEndpointId,omitempty" bson:"dataCollectionEndpointId,omitempty"`
	// DataEncryption           *struct {
	// 	Type string `json:"type,omitempty" bson:"type,omitempty"`
	// } `json:"dataEncryption,omitempty" bson:"dataEncryption,omitempty"`
	// DataEndpointEnabled      bool                         `json:"dataEndpointEnabled,omitempty" bson:"dataEndpointEnabled,omitempty"`
	// DataEndpointHostNames    []string                     `json:"dataEndpointHostNames,omitempty" bson:"dataEndpointHostNames,omitempty"`
	// DataFlows                []AzureResourceDataFlows     `json:"dataFlows,omitempty" bson:"dataFlows,omitempty"`
	// DataLocation             string                       `json:"dataLocation,omitempty" bson:"dataLocation,omitempty"`
	// DataProtection           *AzureResourceDataProtection `json:"dataProtection,omitempty" bson:"dataProtection,omitempty"`
	// DataResidencyBoundary    string                       `json:"dataResidencyBoundary,omitempty" bson:"dataResidencyBoundary,omitempty"`
	// DataSources              *AzureResourceDataSources    `json:"dataSources,omitempty" bson:"dataSources,omitempty"`
	// DatabaseAccountOfferType string                       `json:"databaseAccountOfferType,omitempty" bson:"databaseAccountOfferType,omitempty"`
	// DatabaseID               string                       `json:"databaseId,omitempty" bson:"databaseId,omitempty"`
	// DateCreated              string                       `json:"dateCreated,omitempty" bson:"dateCreated,omitempty"`
	// DdosSettings             *struct {
	// 	ProtectionMode string `json:"protectionMode,omitempty" bson:"protectionMode,omitempty"`
	// } `json:"ddosSettings,omitempty" bson:"ddosSettings,omitempty"`
	// DedicatedCoreQuota            float64 `json:"dedicatedCoreQuota,omitempty" bson:"dedicatedCoreQuota,omitempty"`
	// DedicatedCoreQuotaPerVmFamily []*struct {
	// 	CoreQuota float64 `json:"coreQuota,omitempty" bson:"coreQuota,omitempty"`
	// 	Name      string  `json:"name,omitempty" bson:"name,omitempty"`
	// } `json:"dedicatedCoreQuotaPerVMFamily,omitempty" bson:"dedicatedCoreQuotaPerVMFamily,omitempty"`
	// DedicatedCoreQuotaPerVmFamilyEnforced bool                                 `json:"dedicatedCoreQuotaPerVMFamilyEnforced,omitempty" bson:"dedicatedCoreQuotaPerVMFamilyEnforced,omitempty"`
	// DefaultDomain                         string                               `json:"defaultDomain,omitempty" bson:"defaultDomain,omitempty"`
	// DefaultGroupQuotaInKiBs               float64                              `json:"defaultGroupQuotaInKiBs,omitempty" bson:"defaultGroupQuotaInKiBs,omitempty"`
	// DefaultHostName                       string                               `json:"defaultHostName,omitempty" bson:"defaultHostName,omitempty"`
	// DefaultHostNameScope                  string                               `json:"defaultHostNameScope,omitempty" bson:"defaultHostNameScope,omitempty"`
	// DefaultIdentity                       string                               `json:"defaultIdentity,omitempty" bson:"defaultIdentity,omitempty"`
	// DefaultSecondaryLocation              string                               `json:"defaultSecondaryLocation,omitempty" bson:"defaultSecondaryLocation,omitempty"`
	// DefaultSecurityRules                  []*AzureResourceDefaultSecurityRules `json:"defaultSecurityRules,omitempty" bson:"defaultSecurityRules,omitempty"`
	// DefaultToOAuthAuthentication          bool                                 `json:"defaultToOAuthAuthentication,omitempty" bson:"defaultToOAuthAuthentication,omitempty"`
	// DefaultUserQuotaInKiBs                float64                              `json:"defaultUserQuotaInKiBs,omitempty" bson:"defaultUserQuotaInKiBs,omitempty"`
	// Definition                            *AzureResourceDefinition             `json:"definition,omitempty" bson:"definition,omitempty"`
	// DeploymentID                          string                               `json:"deploymentId,omitempty" bson:"deploymentId,omitempty"`
	// Description                           *string                              `json:"description,omitempty" bson:"description,omitempty"`
	// Destinations                          *AzureResourceDestinations           `json:"destinations,omitempty" bson:"destinations,omitempty"`
	// DhcpOptions                           *struct {
	// 	DnsServers []string `json:"dnsServers,omitempty" bson:"dnsServers,omitempty"`
	// } `json:"dhcpOptions,omitempty" bson:"dhcpOptions,omitempty"`
	// DiagnosticsProfile                 *AzureResourceDiagnosticsProfile `json:"diagnosticsProfile,omitempty" bson:"diagnosticsProfile,omitempty"`
	// DisableBgpRoutePropagation         bool                             `json:"disableBgpRoutePropagation,omitempty" bson:"disableBgpRoutePropagation,omitempty"`
	// DisableCopyPaste                   bool                             `json:"disableCopyPaste,omitempty" bson:"disableCopyPaste,omitempty"`
	// DisableIpSecReplayProtection       bool                             `json:"disableIPSecReplayProtection,omitempty" bson:"disableIPSecReplayProtection,omitempty"`
	// DisableKeyBasedMetadataWriteAccess bool                             `json:"disableKeyBasedMetadataWriteAccess,omitempty" bson:"disableKeyBasedMetadataWriteAccess,omitempty"`
	// DisableLocalAuth                   bool                             `json:"disableLocalAuth,omitempty" bson:"disableLocalAuth,omitempty"`
	// DisableTcpStateTracking            bool                             `json:"disableTcpStateTracking,omitempty" bson:"disableTcpStateTracking,omitempty"`
	// DisableVpnEncryption               bool                             `json:"disableVpnEncryption,omitempty" bson:"disableVpnEncryption,omitempty"`
	// Disallowed                         *struct {
	// 	DiskTypes []any `json:"diskTypes,omitempty" bson:"diskTypes,omitempty"`
	// } `json:"disallowed,omitempty" bson:"disallowed,omitempty"`
	// DiskIopsReadWrite float64                    `json:"diskIOPSReadWrite,omitempty" bson:"diskIOPSReadWrite,omitempty"`
	// DiskMBpsReadWrite float64                    `json:"diskMBpsReadWrite,omitempty" bson:"diskMBpsReadWrite,omitempty"`
	// DiskSizeBytes     float64                    `json:"diskSizeBytes,omitempty" bson:"diskSizeBytes,omitempty"`
	// DiskSizeGb        float64                    `json:"diskSizeGB,omitempty" bson:"diskSizeGB,omitempty"`
	// DiskState         string                     `json:"diskState,omitempty" bson:"diskState,omitempty"`
	// DisplayName       string                     `json:"displayName,omitempty" bson:"displayName,omitempty"`
	// Distribute        []*AzureResourceDistribute `json:"distribute,omitempty" bson:"distribute,omitempty"`
	// DnsConfiguration  *struct {
	// 	DnsLegacySortOrder bool `json:"dnsLegacySortOrder,omitempty" bson:"dnsLegacySortOrder,omitempty"`
	// } `json:"dnsConfiguration,omitempty" bson:"dnsConfiguration,omitempty"`
	// DnsEndpointType                        string                                     `json:"dnsEndpointType,omitempty" bson:"dnsEndpointType,omitempty"`
	// DnsName                                string                                     `json:"dnsName,omitempty" bson:"dnsName,omitempty"`
	// DnsPrefix                              string                                     `json:"dnsPrefix,omitempty" bson:"dnsPrefix,omitempty"`
	// DnsSettings                            *AzureResourceDnsSettings                  `json:"dnsSettings,omitempty" bson:"dnsSettings,omitempty"`
	// DoNotRunExtensionsOnOverprovisionedVMs bool                                       `json:"doNotRunExtensionsOnOverprovisionedVMs,omitempty" bson:"doNotRunExtensionsOnOverprovisionedVMs,omitempty"`
	// DocumentEndpoint                       string                                     `json:"documentEndpoint,omitempty" bson:"documentEndpoint,omitempty"`
	// DomainManagement                       string                                     `json:"domainManagement,omitempty" bson:"domainManagement,omitempty"`
	// DomainName                             string                                     `json:"domainName,omitempty" bson:"domainName,omitempty"`
	// DomainVerificationIdentifiers          any                                        `json:"domainVerificationIdentifiers,omitempty" bson:"domainVerificationIdentifiers,omitempty"`
	// DpdTimeoutSeconds                      float64                                    `json:"dpdTimeoutSeconds,omitempty" bson:"dpdTimeoutSeconds,omitempty"`
	// EarliestRestoreDate                    string                                     `json:"earliestRestoreDate,omitempty" bson:"earliestRestoreDate,omitempty"`
	// EgressBytesTransferred                 float64                                    `json:"egressBytesTransferred,omitempty" bson:"egressBytesTransferred,omitempty"`
	// ElasticScaleEnabled                    bool                                       `json:"elasticScaleEnabled,omitempty" bson:"elasticScaleEnabled,omitempty"`
	// EligibleLogCategories                  string                                     `json:"eligibleLogCategories,omitempty" bson:"eligibleLogCategories,omitempty"`
	// EmailReceivers                         []*AzureResourceEmailReceivers             `json:"emailReceivers,omitempty" bson:"emailReceivers,omitempty"`
	// EnableAcceleratedNetworking            bool                                       `json:"enableAcceleratedNetworking,omitempty" bson:"enableAcceleratedNetworking,omitempty"`
	// EnableAnalyticalStorage                bool                                       `json:"enableAnalyticalStorage,omitempty" bson:"enableAnalyticalStorage,omitempty"`
	// EnableAutomaticFailover                bool                                       `json:"enableAutomaticFailover,omitempty" bson:"enableAutomaticFailover,omitempty"`
	// EnableAutomaticUpgrade                 bool                                       `json:"enableAutomaticUpgrade,omitempty" bson:"enableAutomaticUpgrade,omitempty"`
	// EnableBgp                              bool                                       `json:"enableBgp,omitempty" bson:"enableBgp,omitempty"`
	// EnableBgpRouteTranslationForNat        bool                                       `json:"enableBgpRouteTranslationForNat,omitempty" bson:"enableBgpRouteTranslationForNat,omitempty"`
	// EnableBurstCapacity                    bool                                       `json:"enableBurstCapacity,omitempty" bson:"enableBurstCapacity,omitempty"`
	// EnableClientTelemetry                  bool                                       `json:"enableClientTelemetry,omitempty" bson:"enableClientTelemetry,omitempty"`
	// EnableDdosProtection                   bool                                       `json:"enableDdosProtection,omitempty" bson:"enableDdosProtection,omitempty"`
	// EnableDirectPortRateLimit              bool                                       `json:"enableDirectPortRateLimit,omitempty" bson:"enableDirectPortRateLimit,omitempty"`
	// EnableFileCopy                         bool                                       `json:"enableFileCopy,omitempty" bson:"enableFileCopy,omitempty"`
	// EnableFreeTier                         bool                                       `json:"enableFreeTier,omitempty" bson:"enableFreeTier,omitempty"`
	// EnableIpForwarding                     bool                                       `json:"enableIPForwarding,omitempty" bson:"enableIPForwarding,omitempty"`
	// EnableIpConnect                        bool                                       `json:"enableIpConnect,omitempty" bson:"enableIpConnect,omitempty"`
	// EnableKerberos                         bool                                       `json:"enableKerberos,omitempty" bson:"enableKerberos,omitempty"`
	// EnableMultipleWriteLocations           bool                                       `json:"enableMultipleWriteLocations,omitempty" bson:"enableMultipleWriteLocations,omitempty"`
	// EnablePartitionKeyMonitor              bool                                       `json:"enablePartitionKeyMonitor,omitempty" bson:"enablePartitionKeyMonitor,omitempty"`
	// EnablePartitionMerge                   bool                                       `json:"enablePartitionMerge,omitempty" bson:"enablePartitionMerge,omitempty"`
	// EnablePrivateIpAddress                 bool                                       `json:"enablePrivateIpAddress,omitempty" bson:"enablePrivateIpAddress,omitempty"`
	// EnablePrivateLinkFastPath              bool                                       `json:"enablePrivateLinkFastPath,omitempty" bson:"enablePrivateLinkFastPath,omitempty"`
	// EnablePurgeProtection                  bool                                       `json:"enablePurgeProtection,omitempty" bson:"enablePurgeProtection,omitempty"`
	// EnableRbac                             bool                                       `json:"enableRBAC,omitempty" bson:"enableRBAC,omitempty"`
	// EnableRbacAuthorization                bool                                       `json:"enableRbacAuthorization,omitempty" bson:"enableRbacAuthorization,omitempty"`
	// EnableShareableLink                    bool                                       `json:"enableShareableLink,omitempty" bson:"enableShareableLink,omitempty"`
	// EnableSoftDelete                       bool                                       `json:"enableSoftDelete,omitempty" bson:"enableSoftDelete,omitempty"`
	// EnableSubvolumes                       string                                     `json:"enableSubvolumes,omitempty" bson:"enableSubvolumes,omitempty"`
	// EnableTunneling                        bool                                       `json:"enableTunneling,omitempty" bson:"enableTunneling,omitempty"`
	// Enabled                                bool                                       `json:"enabled,omitempty" bson:"enabled,omitempty"`
	// EnabledForDeployment                   bool                                       `json:"enabledForDeployment,omitempty" bson:"enabledForDeployment,omitempty"`
	// EnabledForDiskEncryption               bool                                       `json:"enabledForDiskEncryption,omitempty" bson:"enabledForDiskEncryption,omitempty"`
	// EnabledForTemplateDeployment           bool                                       `json:"enabledForTemplateDeployment,omitempty" bson:"enabledForTemplateDeployment,omitempty"`
	// EnabledHostNames                       []string                                   `json:"enabledHostNames,omitempty" bson:"enabledHostNames,omitempty"`
	// Encapsulation                          string                                     `json:"encapsulation,omitempty" bson:"encapsulation,omitempty"`
	// ResourceEncryption                     *AzureResourceEncryption                   `json:"encryption,omitempty" bson:"encryption,omitempty"`
	// EncryptionKeySource                    string                                     `json:"encryptionKeySource,omitempty" bson:"encryptionKeySource,omitempty"`
	// EncryptionSettingsCollection           *AzureResourceEncryptionSettingsCollection `json:"encryptionSettingsCollection,omitempty" bson:"encryptionSettingsCollection,omitempty"`
	// EncryptionType                         string                                     `json:"encryptionType,omitempty" bson:"encryptionType,omitempty"`
	// EndToEndEncryptionEnabled              bool                                       `json:"endToEndEncryptionEnabled,omitempty" bson:"endToEndEncryptionEnabled,omitempty"`
	// Endpoint                               string                                     `json:"endpoint,omitempty" bson:"endpoint,omitempty"`
	// Endpoints                              any                                        `json:"endpoints,omitempty" bson:"endpoints,omitempty"`
	// EndpointsConfiguration                 *AzureResourceEndpointsConfiguration       `json:"endpointsConfiguration,omitempty" bson:"endpointsConfiguration,omitempty"`
	// EnvironmentID                          string                                     `json:"environmentId,omitempty" bson:"environmentId,omitempty"`
	// EtherType                              string                                     `json:"etherType,omitempty" bson:"etherType,omitempty"`
	// EvaluationFrequency                    string                                     `json:"evaluationFrequency,omitempty" bson:"evaluationFrequency,omitempty"`
	// EventHubReceivers                      []any                                      `json:"eventHubReceivers,omitempty" bson:"eventHubReceivers,omitempty"`
	// EventStreamEndpoint                    string                                     `json:"eventStreamEndpoint,omitempty" bson:"eventStreamEndpoint,omitempty"`
	// EvictionPolicy                         string                                     `json:"evictionPolicy,omitempty" bson:"evictionPolicy,omitempty"`
	// ExactStagingResourceGroup              string                                     `json:"exactStagingResourceGroup,omitempty" bson:"exactStagingResourceGroup,omitempty"`
	// ExistingServerFarmIds                  any                                        `json:"existingServerFarmIds,omitempty" bson:"existingServerFarmIds,omitempty"`
	// ExportPolicy                           *struct {
	// 	Rules []any `json:"rules,omitempty" bson:"rules,omitempty"`
	// } `json:"exportPolicy,omitempty" bson:"exportPolicy,omitempty"`
	// ExpressRouteConnections []*AzureResourceExpressRouteConnections `json:"expressRouteConnections,omitempty" bson:"expressRouteConnections,omitempty"`
	// ExpressRouteGateway     *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"expressRouteGateway,omitempty" bson:"expressRouteGateway,omitempty"`
	// ExpressRouteGatewayBypass bool `json:"expressRouteGatewayBypass,omitempty" bson:"expressRouteGatewayBypass,omitempty"`
	// ExpressRoutePort          *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"expressRoutePort,omitempty" bson:"expressRoutePort,omitempty"`
	Extended *AzureResourceExtended `json:"extended,omitempty" bson:"extended,omitempty"`
	// ExtensionProperties *struct {
	// 	InGuestPatchMode string `json:"InGuestPatchMode,omitempty" bson:"InGuestPatchMode,omitempty"`
	// } `json:"extensionProperties,omitempty" bson:"extensionProperties,omitempty"`
	// ExtensionsTimeBudget     string                          `json:"extensionsTimeBudget,omitempty" bson:"extensionsTimeBudget,omitempty"`
	// ExternalGovernanceStatus string                          `json:"externalGovernanceStatus,omitempty" bson:"externalGovernanceStatus,omitempty"`
	// FactoryStatistics        *AzureResourceFactoryStatistics `json:"factoryStatistics,omitempty" bson:"factoryStatistics,omitempty"`
	// FailoverPolicies         []*AzureResourceFailoverPolicy  `json:"failoverPolicies,omitempty" bson:"failoverPolicies,omitempty"`
	// FeatureSettings          *AzureResourceFeatureSettings   `json:"featureSettings,omitempty" bson:"featureSettings,omitempty"`
	// Features                 any                             `json:"features,omitempty" bson:"features,omitempty"`
	// FileSystemID             string                          `json:"fileSystemId,omitempty" bson:"fileSystemId,omitempty"`
	// FirewallPolicies         []struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"firewallPolicies,omitempty" bson:"firewallPolicies,omitempty"`
	// FirewallPolicy *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"firewallPolicy,omitempty" bson:"firewallPolicy,omitempty"`
	// Firewalls []struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"firewalls,omitempty" bson:"firewalls,omitempty"`
	// FlowAnalyticsConfiguration *AzureResourceFlowAnalyticsConfiguration `json:"flowAnalyticsConfiguration,omitempty" bson:"flowAnalyticsConfiguration,omitempty"`
	// FlowLogs                   []struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"flowLogs,omitempty" bson:"flowLogs,omitempty"`
	// ForceCmkForQuery bool   `json:"forceCmkForQuery,omitempty" bson:"forceCmkForQuery,omitempty"`
	// ForceUpdateTag   string `json:"forceUpdateTag,omitempty" bson:"forceUpdateTag,omitempty"`
	// Format           *struct {
	// 	Type    string  `json:"type,omitempty" bson:"type,omitempty"`
	// 	Version float64 `json:"version,omitempty" bson:"version,omitempty"`
	// } `json:"format,omitempty" bson:"format,omitempty"`
	// Fqdn                                  string                                  `json:"fqdn,omitempty" bson:"fqdn,omitempty"`
	// FreeOfferExpirationTime               any                                     `json:"freeOfferExpirationTime,omitempty" bson:"freeOfferExpirationTime,omitempty"`
	// FriendlyName                          string                                  `json:"friendlyName,omitempty" bson:"friendlyName,omitempty"`
	// FromSenderDomain                      string                                  `json:"fromSenderDomain,omitempty" bson:"fromSenderDomain,omitempty"`
	// FrontendIpConfigurations              []*AzureResourceFrontendIpConfiguration `json:"frontendIPConfigurations,omitempty" bson:"frontendIPConfigurations,omitempty"`
	// FtpUsername                           string                                  `json:"ftpUsername,omitempty" bson:"ftpUsername,omitempty"`
	// FtpsHostName                          string                                  `json:"ftpsHostName,omitempty" bson:"ftpsHostName,omitempty"`
	// FullyQualifiedDomainName              string                                  `json:"fullyQualifiedDomainName,omitempty" bson:"fullyQualifiedDomainName,omitempty"`
	// FunctionAppConfig                     any                                     `json:"functionAppConfig,omitempty" bson:"functionAppConfig,omitempty"`
	// FunctionExecutionUnitsCache           any                                     `json:"functionExecutionUnitsCache,omitempty" bson:"functionExecutionUnitsCache,omitempty"`
	// FunctionsRuntimeAdminIsolationEnabled bool                                    `json:"functionsRuntimeAdminIsolationEnabled,omitempty" bson:"functionsRuntimeAdminIsolationEnabled,omitempty"`
	// GatewayCustomBgpIpAddresses           []any                                   `json:"gatewayCustomBgpIpAddresses,omitempty" bson:"gatewayCustomBgpIpAddresses,omitempty"`
	// GatewayManagerEtag                    string                                  `json:"gatewayManagerEtag,omitempty" bson:"gatewayManagerEtag,omitempty"`
	// GatewayType                           string                                  `json:"gatewayType,omitempty" bson:"gatewayType,omitempty"`
	// GeoDataReplication                    *AzureResourceGeoDataReplication        `json:"geoDataReplication,omitempty" bson:"geoDataReplication,omitempty"`
	// GeoDistributions                      any                                     `json:"geoDistributions,omitempty" bson:"geoDistributions,omitempty"`
	// GeoRegion                             string                                  `json:"geoRegion,omitempty" bson:"geoRegion,omitempty"`
	// GlobalParameters                      *struct {
	// 	Owner struct {
	// 		Type  string `json:"type,omitempty" bson:"type,omitempty"`
	// 		Value string `json:"value,omitempty" bson:"value,omitempty"`
	// 	} `json:"Owner,omitempty" bson:"Owner,omitempty"`
	// } `json:"globalParameters,omitempty" bson:"globalParameters,omitempty"`
	// GlobalReachEnabled bool   `json:"globalReachEnabled,omitempty" bson:"globalReachEnabled,omitempty"`
	// GroupShortName     string `json:"groupShortName,omitempty" bson:"groupShortName,omitempty"`
	HardwareProfile *struct {
		VmSize    string                  `json:"vmSize,omitempty" bson:"vmSize,omitempty"`
		VmSizeSku *AzureVirtualMachineSku `json:"vmSizeSku,omitempty" bson:"vmSizeSku,omitempty"`
	} `json:"hardwareProfile,omitempty" bson:"hardwareProfile,omitempty"`
	// HighAvailability *struct {
	// 	Mode  string `json:"mode,omitempty" bson:"mode,omitempty"`
	// 	State string `json:"state,omitempty" bson:"state,omitempty"`
	// } `json:"highAvailability,omitempty" bson:"highAvailability,omitempty"`
	// HnsOnMigrationInProgress    bool                              `json:"hnsOnMigrationInProgress,omitempty" bson:"hnsOnMigrationInProgress,omitempty"`
	// HomeStamp                   string                            `json:"homeStamp,omitempty" bson:"homeStamp,omitempty"`
	// HostName                    string                            `json:"hostName,omitempty" bson:"hostName,omitempty"`
	// HostNameSslStates           []*AzureResourceHostNameSslStates `json:"hostNameSslStates,omitempty" bson:"hostNameSslStates,omitempty"`
	// HostNames                   []string                          `json:"hostNames,omitempty" bson:"hostNames,omitempty"`
	// HostNamesDisabled           bool                              `json:"hostNamesDisabled,omitempty" bson:"hostNamesDisabled,omitempty"`
	// HostedWorkloads             []string                          `json:"hostedWorkloads,omitempty" bson:"hostedWorkloads,omitempty"`
	// HostingEnvironment          any                               `json:"hostingEnvironment,omitempty" bson:"hostingEnvironment,omitempty"`
	// HostingEnvironmentID        any                               `json:"hostingEnvironmentId,omitempty" bson:"hostingEnvironmentId,omitempty"`
	// HostingEnvironmentProfile   any                               `json:"hostingEnvironmentProfile,omitempty" bson:"hostingEnvironmentProfile,omitempty"`
	// HourlySchedule              *struct{}                         `json:"hourlySchedule,omitempty" bson:"hourlySchedule,omitempty"`
	// HTTPSOnly                   bool                              `json:"httpsOnly,omitempty" bson:"httpsOnly,omitempty"`
	// HubIpAddresses              *AzureResourceHubIpAddresses      `json:"hubIPAddresses,omitempty" bson:"hubIPAddresses,omitempty"`
	// HubRoutingPreference        string                            `json:"hubRoutingPreference,omitempty" bson:"hubRoutingPreference,omitempty"`
	// HyperV                      bool                              `json:"hyperV,omitempty" bson:"hyperV,omitempty"`
	// HyperVGeneration            string                            `json:"hyperVGeneration,omitempty" bson:"hyperVGeneration,omitempty"`
	// Identifier                  *AzureResourceIdentifier          `json:"identifier,omitempty" bson:"identifier,omitempty"`
	// IdentityProfile             *AzureResourceIdentityProfile     `json:"identityProfile,omitempty" bson:"identityProfile,omitempty"`
	// IdleTimeoutInMinutes        float64                           `json:"idleTimeoutInMinutes,omitempty" bson:"idleTimeoutInMinutes,omitempty"`
	// ImmutableID                 string                            `json:"immutableId,omitempty" bson:"immutableId,omitempty"`
	// ImmutableResourceID         string                            `json:"immutableResourceId,omitempty" bson:"immutableResourceId,omitempty"`
	// InFlightFeatures            []string                          `json:"inFlightFeatures,omitempty" bson:"inFlightFeatures,omitempty"`
	// InProgressOperationID       any                               `json:"inProgressOperationId,omitempty" bson:"inProgressOperationId,omitempty"`
	// InboundIpAddress            string                            `json:"inboundIpAddress,omitempty" bson:"inboundIpAddress,omitempty"`
	// InboundNatPools             []any                             `json:"inboundNatPools,omitempty" bson:"inboundNatPools,omitempty"`
	// InboundNatRules             []*AzureResourceInboundNatRule    `json:"inboundNatRules,omitempty" bson:"inboundNatRules,omitempty"`
	// Incremental                 bool                              `json:"incremental,omitempty" bson:"incremental,omitempty"`
	// IncrementalSnapshotFamilyID string                            `json:"incrementalSnapshotFamilyId,omitempty" bson:"incrementalSnapshotFamilyId,omitempty"`
	// InfrastructureResourceGroup *string                           `json:"infrastructureResourceGroup,omitempty" bson:"infrastructureResourceGroup,omitempty"`
	// IngressBytesTransferred     float64                           `json:"ingressBytesTransferred,omitempty" bson:"ingressBytesTransferred,omitempty"`
	// InputSchema                 string                            `json:"inputSchema,omitempty" bson:"inputSchema,omitempty"`
	// InstallPatches              *AzureResourceInstallPatches      `json:"installPatches,omitempty" bson:"installPatches,omitempty"`
	// InstanceID                  string                            `json:"instanceId,omitempty" bson:"instanceId,omitempty"`
	// InternalID                  string                            `json:"internalId,omitempty" bson:"internalId,omitempty"`
	// IntrusionDetection          *AzureResourceIntrusionDetection  `json:"intrusionDetection,omitempty" bson:"intrusionDetection,omitempty"`
	// IpAddress                   string                            `json:"ipAddress,omitempty" bson:"ipAddress,omitempty"`
	// IpAddresses                 []string                          `json:"ipAddresses,omitempty" bson:"ipAddresses,omitempty"`
	// IpConfiguration             *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"ipConfiguration,omitempty" bson:"ipConfiguration,omitempty"`
	IpConfigurations                  []*AzureResourceIpConfiguration `json:"ipConfigurations,omitempty" bson:"ipConfigurations,omitempty"`
	IpCidrBlock                       string                          `json:"ipCidrBlock,omitempty,omitzero" bson:"ipCidrBlock,omitempty,omitzero"`
	IpNumberAddresses                 int                             `json:"ipNumberAddresses,omitempty,omitzero" bson:"ipNumberAddresses,omitempty,omitzero"`
	IpNumberAddressableHosts          int                             `json:"ipNumberAddressableHosts,omitempty,omitzero" bson:"ipNumberAddressableHosts,omitempty,omitzero"`
	IpAddressesUsed                   int                             `json:"ipAddressesUsed,omitempty,omitzero" bson:"ipAddressesUsed,omitempty,omitzero"`
	IpNumberAddressableHostsRemaining int                             `json:"ipNumberAddressableHostsRemaining,omitempty,omitzero" bson:"ipNumberAddressableHostsRemaining,omitempty,omitzero"`
	IpPercentAddressableHostsUsed     float64                         `json:"ipPercentAddressableHostsUsed,omitempty,omitzero" bson:"ipPercentAddressableHostsUsed,omitempty,omitzero"`
	IpRange                           []string                        `json:"IpRange,omitempty,omitzero" bson:"IpRange,omitempty,omitzero"`
	// IpMode           string                          `json:"ipMode,omitempty" bson:"ipMode,omitempty"`
	// IpRules          []struct {
	// 	IpAddressOrRange string `json:"ipAddressOrRange,omitempty" bson:"ipAddressOrRange,omitempty"`
	// } `json:"ipRules,omitempty" bson:"ipRules,omitempty"`
	// IpTags                          []any   `json:"ipTags,omitempty" bson:"ipTags,omitempty"`
	// IpsecPolicies                   []any   `json:"ipsecPolicies,omitempty" bson:"ipsecPolicies,omitempty"`
	// IsAutoInflateEnabled            bool    `json:"isAutoInflateEnabled,omitempty" bson:"isAutoInflateEnabled,omitempty"`
	// IsDefaultQuotaEnabled           bool    `json:"isDefaultQuotaEnabled,omitempty" bson:"isDefaultQuotaEnabled,omitempty"`
	// IsEnabled                       bool    `json:"isEnabled,omitempty" bson:"isEnabled,omitempty"`
	// IsHnsEnabled                    bool    `json:"isHnsEnabled,omitempty" bson:"isHnsEnabled,omitempty"`
	// IsIPv6EnabledPrivateEndpoint    bool    `json:"isIPv6EnabledPrivateEndpoint,omitempty" bson:"isIPv6EnabledPrivateEndpoint,omitempty"`
	// IsInfraEncryptionEnabled        bool    `json:"isInfraEncryptionEnabled,omitempty" bson:"isInfraEncryptionEnabled,omitempty"`
	// IsLedgerOn                      bool    `json:"isLedgerOn,omitempty" bson:"isLedgerOn,omitempty"`
	// IsLocalUserEnabled              bool    `json:"isLocalUserEnabled,omitempty" bson:"isLocalUserEnabled,omitempty"`
	// IsMigrateToCses                 bool    `json:"isMigrateToCSES,omitempty" bson:"isMigrateToCSES,omitempty"`
	// IsNfsV3Enabled                  bool    `json:"isNfsV3Enabled,omitempty" bson:"isNfsV3Enabled,omitempty"`
	// IsRestoring                     bool    `json:"isRestoring,omitempty" bson:"isRestoring,omitempty"`
	// IsRoutingPreferenceInternet     bool    `json:"isRoutingPreferenceInternet,omitempty" bson:"isRoutingPreferenceInternet,omitempty"`
	// IsSftpEnabled                   bool    `json:"isSftpEnabled,omitempty" bson:"isSftpEnabled,omitempty"`
	// IsSpot                          bool    `json:"isSpot,omitempty" bson:"isSpot,omitempty"`
	// IsVaultProtectedByResourceGuard bool    `json:"isVaultProtectedByResourceGuard,omitempty" bson:"isVaultProtectedByResourceGuard,omitempty"`
	// IsVirtualNetworkFilterEnabled   bool    `json:"isVirtualNetworkFilterEnabled,omitempty" bson:"isVirtualNetworkFilterEnabled,omitempty"`
	// IsXenon                         bool    `json:"isXenon,omitempty" bson:"isXenon,omitempty"`
	// ItsmReceivers                   []any   `json:"itsmReceivers,omitempty" bson:"itsmReceivers,omitempty"`
	// JobCount                        float64 `json:"jobCount,omitempty" bson:"jobCount,omitempty"`
	// KafkaEnabled                    bool    `json:"kafkaEnabled,omitempty" bson:"kafkaEnabled,omitempty"`
	// KedaConfiguration               *struct {
	// 	Version string `json:"version,omitempty" bson:"version,omitempty"`
	// } `json:"kedaConfiguration,omitempty" bson:"kedaConfiguration,omitempty"`
	// KerberosEnabled bool `json:"kerberosEnabled,omitempty" bson:"kerberosEnabled,omitempty"`
	// KeyCreationTime *struct {
	// 	Key1 *string `json:"key1,omitempty" bson:"key1,omitempty"`
	// 	Key2 *string `json:"key2,omitempty" bson:"key2,omitempty"`
	// } `json:"keyCreationTime,omitempty" bson:"keyCreationTime,omitempty"`
	// KeyVaultReferenceIdentity string                            `json:"keyVaultReferenceIdentity,omitempty" bson:"keyVaultReferenceIdentity,omitempty"`
	// KeysMetadata              *AzureResourceKeysMetadata        `json:"keysMetadata,omitempty" bson:"keysMetadata,omitempty"`
	// Kind                      string                            `json:"kind,omitempty" bson:"kind,omitempty"`
	// KubeEnvironmentProfile    any                               `json:"kubeEnvironmentProfile,omitempty" bson:"kubeEnvironmentProfile,omitempty"`
	// KubernetesVersion         string                            `json:"kubernetesVersion,omitempty" bson:"kubernetesVersion,omitempty"`
	// LargeFileSharesState      string                            `json:"largeFileSharesState,omitempty" bson:"largeFileSharesState,omitempty"`
	// LastModifiedBy            any                               `json:"lastModifiedBy,omitempty" bson:"lastModifiedBy,omitempty"`
	// LastModifiedTime          string                            `json:"lastModifiedTime,omitempty" bson:"lastModifiedTime,omitempty"`
	// LastModifiedTimeUtc       string                            `json:"lastModifiedTimeUtc,omitempty" bson:"lastModifiedTimeUtc,omitempty"`
	// LastRunStatus             *AzureResourceLastRunStatus       `json:"lastRunStatus,omitempty" bson:"lastRunStatus,omitempty"`
	// LdapEnabled               bool                              `json:"ldapEnabled,omitempty" bson:"ldapEnabled,omitempty"`
	// LeastPrivilegeMode        string                            `json:"leastPrivilegeMode,omitempty" bson:"leastPrivilegeMode,omitempty"`
	// Lenses                    []*AzureResourceLense             `json:"lenses,omitempty" bson:"lenses,omitempty"`
	// LicenseType               string                            `json:"licenseType,omitempty" bson:"licenseType,omitempty"`
	// LinkedDomains             []string                          `json:"linkedDomains,omitempty" bson:"linkedDomains,omitempty"`
	// LinkedResourceType        string                            `json:"linkedResourceType,omitempty" bson:"linkedResourceType,omitempty"`
	// Links                     []*AzureResourceLink              `json:"links,omitempty" bson:"links,omitempty"`
	// LinuxProfile              *AzureResourceLinuxProfile        `json:"linuxProfile,omitempty" bson:"linuxProfile,omitempty"`
	// LoadBalancingRules        []*AzureResourceLoadBalancingRule `json:"loadBalancingRules,omitempty" bson:"loadBalancingRules,omitempty"`
	// Locations                 []*AzureResourceLocation          `json:"locations,omitempty" bson:"locations,omitempty"`
	// LogActivityTrace          float64                           `json:"logActivityTrace,omitempty" bson:"logActivityTrace,omitempty"`
	// LogProgress               bool                              `json:"logProgress,omitempty" bson:"logProgress,omitempty"`
	// LogVerbose                bool                              `json:"logVerbose,omitempty" bson:"logVerbose,omitempty"`
	// LogicAppReceivers         []any                             `json:"logicAppReceivers,omitempty" bson:"logicAppReceivers,omitempty"`
	// LoginServer               string                            `json:"loginServer,omitempty" bson:"loginServer,omitempty"`
	// LogsIngestion             *struct {
	// 	Endpoint string `json:"endpoint,omitempty" bson:"endpoint,omitempty"`
	// } `json:"logsIngestion,omitempty" bson:"logsIngestion,omitempty"`
	// LowPriorityCoreQuota       float64                         `json:"lowPriorityCoreQuota,omitempty" bson:"lowPriorityCoreQuota,omitempty"`
	// MacAddress                 string                          `json:"macAddress,omitempty" bson:"macAddress,omitempty"`
	// MailFromSenderDomain       string                          `json:"mailFromSenderDomain,omitempty" bson:"mailFromSenderDomain,omitempty"`
	// MaintenanceConfigurationID string                          `json:"maintenanceConfigurationId,omitempty" bson:"maintenanceConfigurationId,omitempty"`
	// MaintenanceScope           string                          `json:"maintenanceScope,omitempty" bson:"maintenanceScope,omitempty"`
	// MaintenanceWindow          *AzureResourceMaintenanceWindow `json:"maintenanceWindow,omitempty" bson:"maintenanceWindow,omitempty"`
	// ManagedEnvironmentID       any                             `json:"managedEnvironmentId,omitempty" bson:"managedEnvironmentId,omitempty"`
	// ManagedResourceGroupName   string                          `json:"managedResourceGroupName,omitempty" bson:"managedResourceGroupName,omitempty"`
	// ManagedResources           *struct {
	// 	ResourceGroup  string `json:"resourceGroup,omitempty" bson:"resourceGroup,omitempty"`
	// 	StorageAccount string `json:"storageAccount,omitempty" bson:"storageAccount,omitempty"`
	// } `json:"managedResources,omitempty" bson:"managedResources,omitempty"`
	// ManualPrivateLinkServiceConnections            []any                  `json:"manualPrivateLinkServiceConnections,omitempty" bson:"manualPrivateLinkServiceConnections,omitempty"`
	// MaxAgentPools                                  float64                `json:"maxAgentPools,omitempty" bson:"maxAgentPools,omitempty"`
	// MaxLogSizeBytes                                float64                `json:"maxLogSizeBytes,omitempty" bson:"maxLogSizeBytes,omitempty"`
	// MaxNumberOfRecordSets                          float64                `json:"maxNumberOfRecordSets,omitempty" bson:"maxNumberOfRecordSets,omitempty"`
	// MaxNumberOfRecordsPerRecordSet                 any                    `json:"maxNumberOfRecordsPerRecordSet,omitempty" bson:"maxNumberOfRecordsPerRecordSet,omitempty"`
	// MaxNumberOfVirtualNetworkLinks                 float64                `json:"maxNumberOfVirtualNetworkLinks,omitempty" bson:"maxNumberOfVirtualNetworkLinks,omitempty"`
	// MaxNumberOfVirtualNetworkLinksWithRegistration float64                `json:"maxNumberOfVirtualNetworkLinksWithRegistration,omitempty" bson:"maxNumberOfVirtualNetworkLinksWithRegistration,omitempty"`
	// MaxNumberOfWorkers                             any                    `json:"maxNumberOfWorkers,omitempty" bson:"maxNumberOfWorkers,omitempty"`
	// MaxShares                                      float64                `json:"maxShares,omitempty" bson:"maxShares,omitempty"`
	// MaxSizeBytes                                   float64                `json:"maxSizeBytes,omitempty" bson:"maxSizeBytes,omitempty"`
	// MaximumElasticWorkerCount                      float64                `json:"maximumElasticWorkerCount,omitempty" bson:"maximumElasticWorkerCount,omitempty"`
	// MaximumNumberOfFiles                           float64                `json:"maximumNumberOfFiles,omitempty" bson:"maximumNumberOfFiles,omitempty"`
	// MaximumNumberOfWorkers                         float64                `json:"maximumNumberOfWorkers,omitempty" bson:"maximumNumberOfWorkers,omitempty"`
	// MaximumThroughputUnits                         float64                `json:"maximumThroughputUnits,omitempty" bson:"maximumThroughputUnits,omitempty"`
	// MdmID                                          string                 `json:"mdmId,omitempty" bson:"mdmId,omitempty"`
	// Metadata                                       *AzureResourceMetadata `json:"metadata,omitempty" bson:"metadata,omitempty"`
	// MetadataSearch                                 string                 `json:"metadataSearch,omitempty" bson:"metadataSearch,omitempty"`
	// MetricID                                       string                 `json:"metricId,omitempty" bson:"metricId,omitempty"`
	// MetricResourceID                               string                 `json:"metricResourceId,omitempty" bson:"metricResourceId,omitempty"`
	// MetricsIngestion                               *struct {
	// 	Endpoint string `json:"endpoint,omitempty" bson:"endpoint,omitempty"`
	// } `json:"metricsIngestion,omitempty" bson:"metricsIngestion,omitempty"`
	// MigrateToVmss        any     `json:"migrateToVMSS,omitempty" bson:"migrateToVMSS,omitempty"`
	// MigrationPhase       string  `json:"migrationPhase,omitempty" bson:"migrationPhase,omitempty"`
	// MigrationState       any     `json:"migrationState,omitempty" bson:"migrationState,omitempty"`
	// MinCapacity          float64 `json:"minCapacity,omitempty" bson:"minCapacity,omitempty"`
	// MinimalTlsVersion    string  `json:"minimalTlsVersion,omitempty" bson:"minimalTlsVersion,omitempty"`
	// MinimumTlsVersion    string  `json:"minimumTlsVersion,omitempty" bson:"minimumTlsVersion,omitempty"`
	// MinorVersion         string  `json:"minorVersion,omitempty" bson:"minorVersion,omitempty"`
	// ModifiedDate         string  `json:"modifiedDate,omitempty" bson:"modifiedDate,omitempty"`
	// MonitoringStatus     string  `json:"monitoringStatus,omitempty" bson:"monitoringStatus,omitempty"`
	// MonthlyBackupsToKeep float64 `json:"monthlyBackupsToKeep,omitempty" bson:"monthlyBackupsToKeep,omitempty"`
	// MonthlySchedule      *struct {
	// 	DaysOfMonth string `json:"daysOfMonth,omitempty" bson:"daysOfMonth,omitempty"`
	// } `json:"monthlySchedule,omitempty" bson:"monthlySchedule,omitempty"`
	// MountTargets                []*AzureResourceMountTargets `json:"mountTargets,omitempty" bson:"mountTargets,omitempty"`
	// Mtu                         string                       `json:"mtu,omitempty" bson:"mtu,omitempty"`
	// MuteActionsDuration         string                       `json:"muteActionsDuration,omitempty" bson:"muteActionsDuration,omitempty"`
	// Name                        string                       `json:"name,omitempty" bson:"name,omitempty"`
	// NameServers                 []string                     `json:"nameServers,omitempty" bson:"nameServers,omitempty"`
	// NatRuleCollections          []any                        `json:"natRuleCollections,omitempty" bson:"natRuleCollections,omitempty"`
	// NatRules                    []any                        `json:"natRules,omitempty" bson:"natRules,omitempty"`
	// Network                     *AzureResourceNetwork        `json:"network,omitempty" bson:"network,omitempty"`
	// NetworkAccessPolicy         string                       `json:"networkAccessPolicy,omitempty" bson:"networkAccessPolicy,omitempty"`
	// NetworkACLBypass            string                       `json:"networkAclBypass,omitempty" bson:"networkAclBypass,omitempty"`
	// NetworkACLBypassResourceIds []any                        `json:"networkAclBypassResourceIds,omitempty" bson:"networkAclBypassResourceIds,omitempty"`
	// NetworkAcls                 *AzureResourceNetworkAcls    `json:"networkAcls,omitempty" bson:"networkAcls,omitempty"`
	// NetworkFeatures             string                       `json:"networkFeatures,omitempty" bson:"networkFeatures,omitempty"`
	// NetworkInterfaces           []struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"networkInterfaces,omitempty" bson:"networkInterfaces,omitempty"`
	// NetworkProfile           *AzureResourceNetworkProfile `json:"networkProfile,omitempty" bson:"networkProfile,omitempty"`
	// NetworkRuleBypassOptions string                       `json:"networkRuleBypassOptions,omitempty" bson:"networkRuleBypassOptions,omitempty"`
	// NetworkRuleCollections   []any                        `json:"networkRuleCollections,omitempty" bson:"networkRuleCollections,omitempty"`
	// NetworkRuleSet           *AzureResourceNetworkRuleSet `json:"networkRuleSet,omitempty" bson:"networkRuleSet,omitempty"`
	// NetworkSecurityGroup     *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"networkSecurityGroup,omitempty" bson:"networkSecurityGroup,omitempty"`
	// NetworkSiblingSetID                         string                             `json:"networkSiblingSetId,omitempty" bson:"networkSiblingSetId,omitempty"`
	// NetworkVirtualAppliances                    []any                              `json:"networkVirtualAppliances,omitempty" bson:"networkVirtualAppliances,omitempty"`
	// NicType                                     string                             `json:"nicType,omitempty" bson:"nicType,omitempty"`
	// NodeConfigurationCount                      float64                            `json:"nodeConfigurationCount,omitempty" bson:"nodeConfigurationCount,omitempty"`
	// NodeManagementEndpoint                      string                             `json:"nodeManagementEndpoint,omitempty" bson:"nodeManagementEndpoint,omitempty"`
	// NodeResourceGroup                           string                             `json:"nodeResourceGroup,omitempty" bson:"nodeResourceGroup,omitempty"`
	// NotificationSettings                        *AzureResourceNotificationSettings `json:"notificationSettings,omitempty" bson:"notificationSettings,omitempty"`
	// NumberOfRecordSets                          float64                            `json:"numberOfRecordSets,omitempty" bson:"numberOfRecordSets,omitempty"`
	// NumberOfSites                               float64                            `json:"numberOfSites,omitempty" bson:"numberOfSites,omitempty"`
	// NumberOfVirtualNetworkLinks                 float64                            `json:"numberOfVirtualNetworkLinks,omitempty" bson:"numberOfVirtualNetworkLinks,omitempty"`
	// NumberOfVirtualNetworkLinksWithRegistration float64                            `json:"numberOfVirtualNetworkLinksWithRegistration,omitempty" bson:"numberOfVirtualNetworkLinksWithRegistration,omitempty"`
	// NumberOfWorkers                             float64                            `json:"numberOfWorkers,omitempty" bson:"numberOfWorkers,omitempty"`
	// Office365LocalBreakoutCategory              string                             `json:"office365LocalBreakoutCategory,omitempty" bson:"office365LocalBreakoutCategory,omitempty"`
	// OidcIssuerProfile                           *struct {
	// 	Enabled   bool   `json:"enabled,omitempty" bson:"enabled,omitempty"`
	// 	IssuerURL string `json:"issuerURL,omitempty" bson:"issuerURL,omitempty"`
	// } `json:"oidcIssuerProfile,omitempty" bson:"oidcIssuerProfile,omitempty"`
	// OpenTelemetryConfiguration  any                                        `json:"openTelemetryConfiguration,omitempty" bson:"openTelemetryConfiguration,omitempty"`
	// OrchestrationMode           string                                     `json:"orchestrationMode,omitempty" bson:"orchestrationMode,omitempty"`
	// OSProfile                   *AzureResourceOSProfile                    `json:"osProfile,omitempty" bson:"osProfile,omitempty"`
	// OSState                     string                                     `json:"osState,omitempty" bson:"osState,omitempty"`
	// OSType                      string                                     `json:"osType,omitempty" bson:"osType,omitempty"`
	// OutboundIpAddresses         string                                     `json:"outboundIpAddresses,omitempty" bson:"outboundIpAddresses,omitempty"`
	// OutboundRules               []*AzureResourceOutboundRules              `json:"outboundRules,omitempty" bson:"outboundRules,omitempty"`
	// OutboundVnetRouting         any                                        `json:"outboundVnetRouting,omitempty" bson:"outboundVnetRouting,omitempty"`
	// OutputTypes                 []any                                      `json:"outputTypes,omitempty" bson:"outputTypes,omitempty"`
	// Outputs                     []any                                      `json:"outputs,omitempty" bson:"outputs,omitempty"`
	// OverallStatus               string                                     `json:"overallStatus,omitempty" bson:"overallStatus,omitempty"`
	// Overprovision               bool                                       `json:"overprovision,omitempty" bson:"overprovision,omitempty"`
	// OverrideQueryTimeRange      string                                     `json:"overrideQueryTimeRange,omitempty" bson:"overrideQueryTimeRange,omitempty"`
	// Owner                       any                                        `json:"owner,omitempty" bson:"owner,omitempty"`
	// P2SConnectionConfigurations []*AzureResourceP2SConnectionConfiguration `json:"p2SConnectionConfigurations,omitempty" bson:"p2SConnectionConfigurations,omitempty"`
	// P2SVpnGateway               *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"p2SVpnGateway,omitempty" bson:"p2SVpnGateway,omitempty"`
	// P2SVpnGateways []struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"p2SVpnGateways,omitempty" bson:"p2SVpnGateways,omitempty"`
	// PacketCaptureDiagnosticState string `json:"packetCaptureDiagnosticState,omitempty" bson:"packetCaptureDiagnosticState,omitempty"`
	// ParameterValueType           string `json:"parameterValueType,omitempty" bson:"parameterValueType,omitempty"`
	// ParameterValues              *struct {
	// 	Token_TenantID  string `json:"token:TenantId,omitempty" bson:"token:TenantId,omitempty"`
	// 	Token_GrantType string `json:"token:grantType,omitempty" bson:"token:grantType,omitempty"`
	// } `json:"parameterValues,omitempty" bson:"parameterValues,omitempty"`
	// Parameters *AzureResourceParameters `json:"parameters,omitempty" bson:"parameters,omitempty"`
	ParentVnet string `json:"ParentVnet,omitempty" bson:"ParentVnet,omitempty"`
	// PausedDate string                   `json:"pausedDate,omitempty" bson:"pausedDate,omitempty"`
	// Peer       *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"peer,omitempty" bson:"peer,omitempty"`
	// PeerAuthentication *struct {
	// 	Mtls struct {
	// 		Enabled bool `json:"enabled,omitempty" bson:"enabled,omitempty"`
	// 	} `json:"mtls,omitempty" bson:"mtls,omitempty"`
	// } `json:"peerAuthentication,omitempty" bson:"peerAuthentication,omitempty"`
	// PeerTrafficConfiguration *struct {
	// 	Encryption struct {
	// 		Enabled bool `json:"enabled,omitempty" bson:"enabled,omitempty"`
	// 	} `json:"encryption,omitempty" bson:"encryption,omitempty"`
	// } `json:"peerTrafficConfiguration,omitempty" bson:"peerTrafficConfiguration,omitempty"`
	// PeeringLocation             string                   `json:"peeringLocation,omitempty" bson:"peeringLocation,omitempty"`
	// Peerings                    []*AzureResourcePeerings `json:"peerings,omitempty" bson:"peerings,omitempty"`
	// PerSiteScaling              bool                     `json:"perSiteScaling,omitempty" bson:"perSiteScaling,omitempty"`
	// PlanName                    string                   `json:"planName,omitempty" bson:"planName,omitempty"`
	// PlatformFaultDomainCount    float64                  `json:"platformFaultDomainCount,omitempty" bson:"platformFaultDomainCount,omitempty"`
	// PlatformUpdateDomainCount   float64                  `json:"platformUpdateDomainCount,omitempty" bson:"platformUpdateDomainCount,omitempty"`
	// Policies                    *AzureResourcePolicies   `json:"policies,omitempty" bson:"policies,omitempty"`
	// PoolAllocationMode          string                   `json:"poolAllocationMode,omitempty" bson:"poolAllocationMode,omitempty"`
	// PoolID                      string                   `json:"poolId,omitempty" bson:"poolId,omitempty"`
	// PoolQuota                   float64                  `json:"poolQuota,omitempty" bson:"poolQuota,omitempty"`
	// PossibleInboundIpAddresses  string                   `json:"possibleInboundIpAddresses,omitempty" bson:"possibleInboundIpAddresses,omitempty"`
	// PossibleOutboundIpAddresses string                   `json:"possibleOutboundIpAddresses,omitempty" bson:"possibleOutboundIpAddresses,omitempty"`
	// PowerState                  *struct {
	// 	Code string `json:"code,omitempty" bson:"code,omitempty"`
	// } `json:"powerState,omitempty" bson:"powerState,omitempty"`
	// Primary          bool                           `json:"primary,omitempty" bson:"primary,omitempty"`
	// PrimaryEndpoints *AzureResourcePrimaryEndpoints `json:"primaryEndpoints,omitempty" bson:"primaryEndpoints,omitempty"`
	// PrimaryLocation  string                         `json:"primaryLocation,omitempty" bson:"primaryLocation,omitempty"`
	// PrincipalID      string                         `json:"principalId,omitempty" bson:"principalId,omitempty"`
	// Priority         string                         `json:"priority,omitempty" bson:"priority,omitempty"`
	// PrivateEndpoint  *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"privateEndpoint,omitempty" bson:"privateEndpoint,omitempty"`
	// PrivateEndpointConnections          []*AzureResourcePrivateEndpointConnections   `json:"privateEndpointConnections,omitempty" bson:"privateEndpointConnections,omitempty"`
	// PrivateEndpointStateForBackup       string                                       `json:"privateEndpointStateForBackup,omitempty" bson:"privateEndpointStateForBackup,omitempty"`
	// PrivateEndpointStateForSiteRecovery string                                       `json:"privateEndpointStateForSiteRecovery,omitempty" bson:"privateEndpointStateForSiteRecovery,omitempty"`
	// PrivateEndpointVNetPolicies         string                                       `json:"privateEndpointVNetPolicies,omitempty" bson:"privateEndpointVNetPolicies,omitempty"`
	// PrivateFqdn                         string                                       `json:"privateFQDN,omitempty" bson:"privateFQDN,omitempty"`
	// PrivateLinkIdentifiers              *string                                      `json:"privateLinkIdentifiers,omitempty" bson:"privateLinkIdentifiers,omitempty"`
	// PrivateLinkResources                []*AzureResourcePrivateLinkResources         `json:"privateLinkResources,omitempty" bson:"privateLinkResources,omitempty"`
	// PrivateLinkScopedResources          []*AzureResourcePrivateLinkScopedResource    `json:"privateLinkScopedResources,omitempty" bson:"privateLinkScopedResources,omitempty"`
	// PrivateLinkServiceConnections       []*AzureResourcePrivateLinkServiceConnection `json:"privateLinkServiceConnections,omitempty" bson:"privateLinkServiceConnections,omitempty"`
	// Probes                              []*AzureResourceProbe                        `json:"probes,omitempty" bson:"probes,omitempty"`
	// ProtocolTypes                       []string                                     `json:"protocolTypes,omitempty" bson:"protocolTypes,omitempty"`
	// ProvisionedBandwidthInGbps          float64                                      `json:"provisionedBandwidthInGbps,omitempty" bson:"provisionedBandwidthInGbps,omitempty"`
	// ProvisioningState                   string                                       `json:"provisioningState,omitempty" bson:"provisioningState,omitempty"`
	// PublicIpAddressVersion              string                                       `json:"publicIPAddressVersion,omitempty" bson:"publicIPAddressVersion,omitempty"`
	// PublicIpAllocationMethod            string                                       `json:"publicIPAllocationMethod,omitempty" bson:"publicIPAllocationMethod,omitempty"`
	// PublicKey                           string                                       `json:"publicKey,omitempty" bson:"publicKey,omitempty"`
	// PublicNetworkAccess                 any                                          `json:"publicNetworkAccess,omitempty" bson:"publicNetworkAccess,omitempty"`
	// PublicNetworkAccessForIngestion     string                                       `json:"publicNetworkAccessForIngestion,omitempty" bson:"publicNetworkAccessForIngestion,omitempty"`
	// PublicNetworkAccessForQuery         string                                       `json:"publicNetworkAccessForQuery,omitempty" bson:"publicNetworkAccessForQuery,omitempty"`
	// Publisher                           string                                       `json:"publisher,omitempty" bson:"publisher,omitempty"`
	// PublishingProfile                   *AzureResourcePublishingProfile              `json:"publishingProfile,omitempty" bson:"publishingProfile,omitempty"`
	// PurchasePlan                        *AzureResourcePurchasePlan                   `json:"purchasePlan,omitempty" bson:"purchasePlan,omitempty"`
	// QosType                             string                                       `json:"qosType,omitempty" bson:"qosType,omitempty"`
	// QueryPackID                         string                                       `json:"queryPackId,omitempty" bson:"queryPackId,omitempty"`
	// RadiusClientRootCertificates        []any                                        `json:"radiusClientRootCertificates,omitempty" bson:"radiusClientRootCertificates,omitempty"`
	// RadiusProxyIPs                      []any                                        `json:"radiusProxyIPs,omitempty" bson:"radiusProxyIPs,omitempty"`
	// RadiusServerAddress                 string                                       `json:"radiusServerAddress,omitempty" bson:"radiusServerAddress,omitempty"`
	// RadiusServerRootCertificates        []any                                        `json:"radiusServerRootCertificates,omitempty" bson:"radiusServerRootCertificates,omitempty"`
	// RadiusServerSecret                  string                                       `json:"radiusServerSecret,omitempty" bson:"radiusServerSecret,omitempty"`
	// RadiusServers                       []any                                        `json:"radiusServers,omitempty" bson:"radiusServers,omitempty"`
	// RawTags                             any                                          `json:"rawTags,omitempty" bson:"rawTags,omitempty"`
	// ReadLocations                       []*AzureResourceReadLocation                 `json:"readLocations,omitempty" bson:"readLocations,omitempty"`
	// ReadScale                           string                                       `json:"readScale,omitempty" bson:"readScale,omitempty"`
	// Recommended                         *AzureResourceRecommended                    `json:"recommended,omitempty" bson:"recommended,omitempty"`
	// RedundancyMode                      string                                       `json:"redundancyMode,omitempty" bson:"redundancyMode,omitempty"`
	// RedundancySettings                  *struct {
	// 	CrossRegionRestore            string `json:"crossRegionRestore,omitempty" bson:"crossRegionRestore,omitempty"`
	// 	StandardTierStorageRedundancy string `json:"standardTierStorageRedundancy,omitempty" bson:"standardTierStorageRedundancy,omitempty"`
	// } `json:"redundancySettings,omitempty" bson:"redundancySettings,omitempty"`
	// RegistrationEnabled bool `json:"registrationEnabled,omitempty" bson:"registrationEnabled,omitempty"`
	// Replica             *struct {
	// 	Capacity float64 `json:"capacity,omitempty" bson:"capacity,omitempty"`
	// 	Role     string  `json:"role,omitempty" bson:"role,omitempty"`
	// } `json:"replica,omitempty" bson:"replica,omitempty"`
	// ReplicaCapacity                  float64 `json:"replicaCapacity,omitempty" bson:"replicaCapacity,omitempty"`
	// ReplicatedRegions                []any   `json:"replicatedRegions,omitempty" bson:"replicatedRegions,omitempty"`
	// ReplicationRole                  string  `json:"replicationRole,omitempty" bson:"replicationRole,omitempty"`
	// RepositorySiteName               string  `json:"repositorySiteName,omitempty" bson:"repositorySiteName,omitempty"`
	// RequestedBackupStorageRedundancy string  `json:"requestedBackupStorageRedundancy,omitempty" bson:"requestedBackupStorageRedundancy,omitempty"`
	// RequestedServiceObjectiveName    string  `json:"requestedServiceObjectiveName,omitempty" bson:"requestedServiceObjectiveName,omitempty"`
	// Reserved                         bool    `json:"reserved,omitempty" bson:"reserved,omitempty"`
	// ResolutionPolicy                 string  `json:"resolutionPolicy,omitempty" bson:"resolutionPolicy,omitempty"`
	// ResourceConfig                   any     `json:"resourceConfig,omitempty" bson:"resourceConfig,omitempty"`
	// ResourceGroup                    string  `json:"resourceGroup,omitempty" bson:"resourceGroup,omitempty"`
	// ResourceGuid                     string  `json:"resourceGuid,omitempty" bson:"resourceGuid,omitempty"`
	// ResourceUid                      string  `json:"resourceUID,omitempty" bson:"resourceUID,omitempty"`
	// RestorePointCollectionID         string  `json:"restorePointCollectionId,omitempty" bson:"restorePointCollectionId,omitempty"`
	// RestoreSettings                  *struct {
	// 	CrossSubscriptionRestoreSettings struct {
	// 		CrossSubscriptionRestoreState string `json:"crossSubscriptionRestoreState,omitempty" bson:"crossSubscriptionRestoreState,omitempty"`
	// 	} `json:"crossSubscriptionRestoreSettings,omitempty" bson:"crossSubscriptionRestoreSettings,omitempty"`
	// } `json:"restoreSettings,omitempty" bson:"restoreSettings,omitempty"`
	// RestrictOutboundNetworkAccess string  `json:"restrictOutboundNetworkAccess,omitempty" bson:"restrictOutboundNetworkAccess,omitempty"`
	// RetentionInDays               float64 `json:"retentionInDays,omitempty" bson:"retentionInDays,omitempty"`
	// RetentionPolicy               *struct {
	// 	Days    float64 `json:"days,omitempty" bson:"days,omitempty"`
	// 	Enabled bool    `json:"enabled,omitempty" bson:"enabled,omitempty"`
	// } `json:"retentionPolicy,omitempty" bson:"retentionPolicy,omitempty"`
	// Revision   string `json:"revision,omitempty" bson:"revision,omitempty"`
	// RouteTable *struct {
	// 	Routes []any `json:"routes,omitempty" bson:"routes,omitempty"`
	// } `json:"routeTable,omitempty" bson:"routeTable,omitempty"`
	// Routes               []*AzureResourceRoute           `json:"routes,omitempty" bson:"routes,omitempty"`
	// RoutingPreference    *AzureResourceRoutingPreference `json:"routingPreference,omitempty" bson:"routingPreference,omitempty"`
	// RoutingState         string                          `json:"routingState,omitempty" bson:"routingState,omitempty"`
	// RoutingWeight        float64                         `json:"routingWeight,omitempty" bson:"routingWeight,omitempty"`
	// RuleCollectionGroups []struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"ruleCollectionGroups,omitempty" bson:"ruleCollectionGroups,omitempty"`
	// RunbookType              string                           `json:"runbookType,omitempty" bson:"runbookType,omitempty"`
	// RunningOperationIds      []any                            `json:"runningOperationIds,omitempty" bson:"runningOperationIds,omitempty"`
	// RuntimeAvailabilityState string                           `json:"runtimeAvailabilityState,omitempty" bson:"runtimeAvailabilityState,omitempty"`
	// ScaleUnits               float64                          `json:"scaleUnits,omitempty" bson:"scaleUnits,omitempty"`
	// ScmSiteAlsoStopped       bool                             `json:"scmSiteAlsoStopped,omitempty" bson:"scmSiteAlsoStopped,omitempty"`
	// ScopeID                  string                           `json:"scopeId,omitempty" bson:"scopeId,omitempty"`
	// Scopes                   []any                            `json:"scopes,omitempty" bson:"scopes,omitempty"`
	// SecondaryEndpoints       *AzureResourceSecondaryEndpoints `json:"secondaryEndpoints,omitempty" bson:"secondaryEndpoints,omitempty"`
	// SecondaryLocation        string                           `json:"secondaryLocation,omitempty" bson:"secondaryLocation,omitempty"`
	// SecureScore              string                           `json:"secureScore,omitempty" bson:"secureScore,omitempty"`
	// SecurityProfile          *AzureResourceSecurityProfile    `json:"securityProfile,omitempty" bson:"securityProfile,omitempty"`
	// SecurityRules            []*AzureResourceSecurityRule     `json:"securityRules,omitempty" bson:"securityRules,omitempty"`
	// SecuritySettings         *AzureResourceSecuritySettings   `json:"securitySettings,omitempty" bson:"securitySettings,omitempty"`
	// SecurityStyle            string                           `json:"securityStyle,omitempty" bson:"securityStyle,omitempty"`
	// SelfLink                 string                           `json:"selfLink,omitempty" bson:"selfLink,omitempty"`
	// SerializedData           any                              `json:"serializedData,omitempty" bson:"serializedData,omitempty"`
	// ServerFarm               any                              `json:"serverFarm,omitempty" bson:"serverFarm,omitempty"`
	// ServerFarmID             any                              `json:"serverFarmId,omitempty" bson:"serverFarmId,omitempty"`
	// ServiceBusEndpoint       string                           `json:"serviceBusEndpoint,omitempty" bson:"serviceBusEndpoint,omitempty"`
	// ServiceKey               string                           `json:"serviceKey,omitempty" bson:"serviceKey,omitempty"`
	// ServiceLevel             string                           `json:"serviceLevel,omitempty" bson:"serviceLevel,omitempty"`
	// ServiceManagementTags    any                              `json:"serviceManagementTags,omitempty" bson:"serviceManagementTags,omitempty"`
	// ServicePrincipalProfile  *struct {
	// 	ClientID string `json:"clientId,omitempty" bson:"clientId,omitempty"`
	// } `json:"servicePrincipalProfile,omitempty" bson:"servicePrincipalProfile,omitempty"`
	// ServiceProviderProperties        *AzureResourceServiceProviderProperties `json:"serviceProviderProperties,omitempty" bson:"serviceProviderProperties,omitempty"`
	// ServiceProviderProvisioningState string                                  `json:"serviceProviderProvisioningState,omitempty" bson:"serviceProviderProvisioningState,omitempty"`
	// Settings                         *AzureResourceSettings                  `json:"settings,omitempty" bson:"settings,omitempty"`
	// Severity                         float64                                 `json:"severity,omitempty" bson:"severity,omitempty"`
	// SinglePlacementGroup             bool                                    `json:"singlePlacementGroup,omitempty" bson:"singlePlacementGroup,omitempty"`
	// SiteConfig                       *AzureResourceSiteConfig                `json:"siteConfig,omitempty" bson:"siteConfig,omitempty"`
	// SiteDisabledReason               float64                                 `json:"siteDisabledReason,omitempty" bson:"siteDisabledReason,omitempty"`
	// SiteMode                         any                                     `json:"siteMode,omitempty" bson:"siteMode,omitempty"`
	// SiteProperties                   *AzureResourceSiteProperties            `json:"siteProperties,omitempty" bson:"siteProperties,omitempty"`
	// SiteScopedCertificatesEnabled    bool                                    `json:"siteScopedCertificatesEnabled,omitempty" bson:"siteScopedCertificatesEnabled,omitempty"`
	// Size                             any                                     `json:"size,omitempty" bson:"size,omitempty"`
	Sku any `json:"sku,omitempty" bson:"sku,omitempty"`
	// SlotName                         any                                     `json:"slotName,omitempty" bson:"slotName,omitempty"`
	// SlotSwapStatus                   any                                     `json:"slotSwapStatus,omitempty" bson:"slotSwapStatus,omitempty"`
	// SmbAccessBasedEnumeration        string                                  `json:"smbAccessBasedEnumeration,omitempty" bson:"smbAccessBasedEnumeration,omitempty"`
	// SmbContinuouslyAvailable         bool                                    `json:"smbContinuouslyAvailable,omitempty" bson:"smbContinuouslyAvailable,omitempty"`
	// SmbEncryption                    bool                                    `json:"smbEncryption,omitempty" bson:"smbEncryption,omitempty"`
	// SmbNonBrowsable                  string                                  `json:"smbNonBrowsable,omitempty" bson:"smbNonBrowsable,omitempty"`
	// SmsReceivers                     []any                                   `json:"smsReceivers,omitempty" bson:"smsReceivers,omitempty"`
	// SnapshotDirectoryVisible         bool                                    `json:"snapshotDirectoryVisible,omitempty" bson:"snapshotDirectoryVisible,omitempty"`
	// Snat                             *struct {
	// 	PrivateRanges []string `json:"privateRanges,omitempty" bson:"privateRanges,omitempty"`
	// } `json:"snat,omitempty" bson:"snat,omitempty"`
	// SoftDeletePolicy *struct {
	// 	IsSoftDeleteEnabled bool `json:"isSoftDeleteEnabled,omitempty" bson:"isSoftDeleteEnabled,omitempty"`
	// } `json:"softDeletePolicy,omitempty" bson:"softDeletePolicy,omitempty"`
	// SoftDeleteRetentionInDays float64     `json:"softDeleteRetentionInDays,omitempty" bson:"softDeleteRetentionInDays,omitempty"`
	Source interface{} `json:"source,omitempty" bson:"source,omitempty"`
	// SourceID                  string      `json:"sourceId,omitempty" bson:"sourceId,omitempty"`
	// SourceVirtualMachine      *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"sourceVirtualMachine,omitempty" bson:"sourceVirtualMachine,omitempty"`
	// Sources                     []*AzureResourceSources      `json:"sources,omitempty" bson:"sources,omitempty"`
	// SpotExpirationTime          any                          `json:"spotExpirationTime,omitempty" bson:"spotExpirationTime,omitempty"`
	// SQLEndpoint                 string                       `json:"sqlEndpoint,omitempty" bson:"sqlEndpoint,omitempty"`
	// SQLImageOffer               string                       `json:"sqlImageOffer,omitempty" bson:"sqlImageOffer,omitempty"`
	// SQLImageSku                 string                       `json:"sqlImageSku,omitempty" bson:"sqlImageSku,omitempty"`
	// SQLManagement               string                       `json:"sqlManagement,omitempty" bson:"sqlManagement,omitempty"`
	// SQLServerLicenseType        string                       `json:"sqlServerLicenseType,omitempty" bson:"sqlServerLicenseType,omitempty"`
	// SSHEnabled                  any                          `json:"sshEnabled,omitempty" bson:"sshEnabled,omitempty"`
	// SslCertificates             any                          `json:"sslCertificates,omitempty" bson:"sslCertificates,omitempty"`
	// Stag                        float64                      `json:"stag,omitempty" bson:"stag,omitempty"`
	// StagingResourceGroup        string                       `json:"stagingResourceGroup,omitempty" bson:"stagingResourceGroup,omitempty"`
	// StartTime                   string                       `json:"startTime,omitempty" bson:"startTime,omitempty"`
	// State                       string                       `json:"state,omitempty" bson:"state,omitempty"`
	// StaticIp                    string                       `json:"staticIp,omitempty" bson:"staticIp,omitempty"`
	// Status                      string                       `json:"status,omitempty" bson:"status,omitempty"`
	// StatusOfPrimary             string                       `json:"statusOfPrimary,omitempty" bson:"statusOfPrimary,omitempty"`
	// StatusOfSecondary           string                       `json:"statusOfSecondary,omitempty" bson:"statusOfSecondary,omitempty"`
	// Statuses                    []*AzureResourceStatus       `json:"statuses,omitempty" bson:"statuses,omitempty"`
	// Storage                     *AzureResourceStorage        `json:"storage,omitempty" bson:"storage,omitempty"`
	// StorageAccountRequired      bool                         `json:"storageAccountRequired,omitempty" bson:"storageAccountRequired,omitempty"`
	// StorageID                   string                       `json:"storageId,omitempty" bson:"storageId,omitempty"`
	StorageProfile *AzureResourceStorageProfile `json:"storageProfile,omitempty" bson:"storageProfile,omitempty"`
	// StorageRecoveryDefaultState string                       `json:"storageRecoveryDefaultState,omitempty" bson:"storageRecoveryDefaultState,omitempty"`
	// StorageSettings             []struct {
	// 	DatastoreType string `json:"datastoreType,omitempty" bson:"datastoreType,omitempty"`
	// 	Type          string `json:"type,omitempty" bson:"type,omitempty"`
	// } `json:"storageSettings,omitempty" bson:"storageSettings,omitempty"`
	// StorageToNetworkProximity string                           `json:"storageToNetworkProximity,omitempty" bson:"storageToNetworkProximity,omitempty"`
	// StorageURI                any                              `json:"storageUri,omitempty" bson:"storageUri,omitempty"`
	// StreamDeclarations        *AzureResourceStreamDeclarations `json:"streamDeclarations,omitempty" bson:"streamDeclarations,omitempty"`
	// Subnet                    *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"subnet,omitempty" bson:"subnet,omitempty"`
	// SubnetID              string                 `json:"subnetId,omitempty" bson:"subnetId,omitempty"`
	// Subnets []AzureResourceSubnet `json:"subnets,omitempty" bson:"subnets,omitempty"`
	SubnetIds []string `json:"subnetIds,omitempty" bson:"subnetIds,omitempty"`
	// Subnets []AzureResourceDetails `json:"subnets,omitempty" bson:"subnets,omitempty"`
	// Subscription          string                 `json:"subscription,omitempty" bson:"subscription,omitempty"`
	// SupportPlan           string                 `json:"supportPlan,omitempty" bson:"supportPlan,omitempty"`
	// SupportedCapabilities *struct {
	// 	AcceleratedNetwork  bool   `json:"acceleratedNetwork,omitempty" bson:"acceleratedNetwork,omitempty"`
	// 	Architecture        string `json:"architecture,omitempty" bson:"architecture,omitempty"`
	// 	DiskControllerTypes string `json:"diskControllerTypes,omitempty" bson:"diskControllerTypes,omitempty"`
	// } `json:"supportedCapabilities,omitempty" bson:"supportedCapabilities,omitempty"`
	// SupportsHibernation      bool                               `json:"supportsHibernation,omitempty" bson:"supportsHibernation,omitempty"`
	// SupportsHTTPSTrafficOnly bool                               `json:"supportsHttpsTrafficOnly,omitempty" bson:"supportsHttpsTrafficOnly,omitempty"`
	// SuppressFailures         bool                               `json:"suppressFailures,omitempty" bson:"suppressFailures,omitempty"`
	// SuspendedTill            any                                `json:"suspendedTill,omitempty" bson:"suspendedTill,omitempty"`
	// SystemData               *AzureResourceSystemData           `json:"systemData,omitempty" bson:"systemData,omitempty"`
	// Tags                     any                                `json:"tags,omitempty" bson:"tags,omitempty"`
	// TapConfigurations        []any                              `json:"tapConfigurations,omitempty" bson:"tapConfigurations,omitempty"`
	// TargetBuildVersion       any                                `json:"targetBuildVersion,omitempty" bson:"targetBuildVersion,omitempty"`
	// TargetResourceGuid       string                             `json:"targetResourceGuid,omitempty" bson:"targetResourceGuid,omitempty"`
	// TargetResourceID         string                             `json:"targetResourceId,omitempty" bson:"targetResourceId,omitempty"`
	// TargetResourceRegion     string                             `json:"targetResourceRegion,omitempty" bson:"targetResourceRegion,omitempty"`
	// TargetResourceType       string                             `json:"targetResourceType,omitempty" bson:"targetResourceType,omitempty"`
	// TargetResourceTypes      []string                           `json:"targetResourceTypes,omitempty" bson:"targetResourceTypes,omitempty"`
	// TargetSwapSlot           any                                `json:"targetSwapSlot,omitempty" bson:"targetSwapSlot,omitempty"`
	// TargetWorkerCount        float64                            `json:"targetWorkerCount,omitempty" bson:"targetWorkerCount,omitempty"`
	// TargetWorkerSizeID       float64                            `json:"targetWorkerSizeId,omitempty" bson:"targetWorkerSizeId,omitempty"`
	// TaskType                 string                             `json:"taskType,omitempty" bson:"taskType,omitempty"`
	// Template                 *AzureResourceTemplate             `json:"template,omitempty" bson:"template,omitempty"`
	// TenantIDOther            string                             `json:"tenantId,omitempty" bson:"tenantId,omitempty"`
	// TestConfigurations       []*AzureResourceTestConfigurations `json:"testConfigurations,omitempty" bson:"testConfigurations,omitempty"`
	// TestGroups               []*AzureResourceTestGroups         `json:"testGroups,omitempty" bson:"testGroups,omitempty"`
	// TestLinks                []*struct {
	// 	Method     string `json:"method,omitempty" bson:"method,omitempty"`
	// 	RequestURI string `json:"requestUri,omitempty" bson:"requestUri,omitempty"`
	// } `json:"testLinks,omitempty" bson:"testLinks,omitempty"`
	// TestRequests            []*AzureResourceTestRequests    `json:"testRequests,omitempty" bson:"testRequests,omitempty"`
	// ThreatIntelMode         string                          `json:"threatIntelMode,omitempty" bson:"threatIntelMode,omitempty"`
	// ThroughputMibps         float64                         `json:"throughputMibps,omitempty" bson:"throughputMibps,omitempty"`
	// Tier                    string                          `json:"tier,omitempty" bson:"tier,omitempty"`
	// TimeCreated             string                          `json:"timeCreated,omitempty" bson:"timeCreated,omitempty"`
	// TimeModified            string                          `json:"timeModified,omitempty" bson:"timeModified,omitempty"`
	// TimeZoneID              string                          `json:"timeZoneId,omitempty" bson:"timeZoneId,omitempty"`
	// TopicType               string                          `json:"topicType,omitempty" bson:"topicType,omitempty"`
	// TotalThroughputMibps    float64                         `json:"totalThroughputMibps,omitempty" bson:"totalThroughputMibps,omitempty"`
	// TrafficManagerHostNames any                             `json:"trafficManagerHostNames,omitempty" bson:"trafficManagerHostNames,omitempty"`
	// TrafficSelectorPolicies []any                           `json:"trafficSelectorPolicies,omitempty" bson:"trafficSelectorPolicies,omitempty"`
	// TransportSecurity       *AzureResourceTransportSecurity `json:"transportSecurity,omitempty" bson:"transportSecurity,omitempty"`
	// Type                    string                          `json:"type,omitempty" bson:"type,omitempty"`
	// TypeHandlerVersion      string                          `json:"typeHandlerVersion,omitempty" bson:"typeHandlerVersion,omitempty"`
	// UniqueID                string                          `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	// UniqueIdentifier        string                          `json:"uniqueIdentifier,omitempty" bson:"uniqueIdentifier,omitempty"`
	// UpdatedAt               string                          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	// UpgradePolicy           *struct {
	// 	Mode string `json:"mode,omitempty" bson:"mode,omitempty"`
	// } `json:"upgradePolicy,omitempty" bson:"upgradePolicy,omitempty"`
	// UpgradeSettings *struct {
	// 	OverrideSettings struct {
	// 		ForceUpgrade bool `json:"forceUpgrade,omitempty" bson:"forceUpgrade,omitempty"`
	// 	} `json:"overrideSettings,omitempty" bson:"overrideSettings,omitempty"`
	// } `json:"upgradeSettings,omitempty" bson:"upgradeSettings,omitempty"`
	// UsageState                     string                            `json:"usageState,omitempty" bson:"usageState,omitempty"`
	// UsageThreshold                 float64                           `json:"usageThreshold,omitempty" bson:"usageThreshold,omitempty"`
	// UseContainerLocalhostBindings  any                               `json:"useContainerLocalhostBindings,omitempty" bson:"useContainerLocalhostBindings,omitempty"`
	// UseLocalAzureIpAddress         bool                              `json:"useLocalAzureIpAddress,omitempty" bson:"useLocalAzureIpAddress,omitempty"`
	// UsePolicyBasedTrafficSelectors bool                              `json:"usePolicyBasedTrafficSelectors,omitempty" bson:"usePolicyBasedTrafficSelectors,omitempty"`
	// UseRadiusProxyIPs              bool                              `json:"useRadiusProxyIPs,omitempty" bson:"useRadiusProxyIPs,omitempty"`
	// UserEngagementTracking         string                            `json:"userEngagementTracking,omitempty" bson:"userEngagementTracking,omitempty"`
	// UserID                         string                            `json:"userId,omitempty" bson:"userId,omitempty"`
	// UtilizedThroughputMibps        float64                           `json:"utilizedThroughputMibps,omitempty" bson:"utilizedThroughputMibps,omitempty"`
	// VaultURI                       string                            `json:"vaultUri,omitempty" bson:"vaultUri,omitempty"`
	// VerificationRecords            *AzureResourceVerificationRecords `json:"verificationRecords,omitempty" bson:"verificationRecords,omitempty"`
	// VerificationStates             *AzureResourceVerificationStates  `json:"verificationStates,omitempty" bson:"verificationStates,omitempty"`
	// Version                        string                            `json:"version,omitempty" bson:"version,omitempty"`
	// VirtualHub                     *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"virtualHub,omitempty" bson:"virtualHub,omitempty"`
	// VirtualHubRouteTableV2S []any `json:"virtualHubRouteTableV2s,omitempty" bson:"virtualHubRouteTableV2s,omitempty"`
	// VirtualHubs             []struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"virtualHubs,omitempty" bson:"virtualHubs,omitempty"`
	VirtualMachine *struct {
		ID string `json:"id,omitempty" bson:"id,omitempty"`
	} `json:"virtualMachine,omitempty" bson:"virtualMachine,omitempty"`
	VirtualMachineProfile    *AzureResourceVirtualMachineProfile `json:"virtualMachineProfile,omitempty" bson:"virtualMachineProfile,omitempty"`
	VirtualMachineResourceID string                              `json:"virtualMachineResourceId,omitempty" bson:"virtualMachineResourceId,omitempty"`
	// VirtualMachines          []struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"virtualMachines,omitempty" bson:"virtualMachines,omitempty"`
	// VirtualNetwork *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"virtualNetwork,omitempty" bson:"virtualNetwork,omitempty"`
	// VirtualNetworkGateway1 *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"virtualNetworkGateway1,omitempty" bson:"virtualNetworkGateway1,omitempty"`
	// VirtualNetworkGatewayMigrationStatus *struct {
	// 	ErrorMessage string `json:"errorMessage,omitempty" bson:"errorMessage,omitempty"`
	// 	Phase        string `json:"phase,omitempty" bson:"phase,omitempty"`
	// 	State        string `json:"state,omitempty" bson:"state,omitempty"`
	// } `json:"virtualNetworkGatewayMigrationStatus,omitempty" bson:"virtualNetworkGatewayMigrationStatus,omitempty"`
	// VirtualNetworkGatewayPolicyGroups []any                                 `json:"virtualNetworkGatewayPolicyGroups,omitempty" bson:"virtualNetworkGatewayPolicyGroups,omitempty"`
	// VirtualNetworkLinkState           string                                `json:"virtualNetworkLinkState,omitempty" bson:"virtualNetworkLinkState,omitempty"`
	// VirtualNetworkPeerings            []*AzureResourceVirtualNetworkPeering `json:"virtualNetworkPeerings,omitempty" bson:"virtualNetworkPeerings,omitempty"`
	// VirtualNetworkRules               []struct {
	// 	ID                               string `json:"id,omitempty" bson:"id,omitempty"`
	// 	IgnoreMissingVNetServiceEndpoint bool   `json:"ignoreMissingVNetServiceEndpoint,omitempty" bson:"ignoreMissingVNetServiceEndpoint,omitempty"`
	// } `json:"virtualNetworkRules,omitempty" bson:"virtualNetworkRules,omitempty"`
	// VirtualNetworkSubnetID              *string `json:"virtualNetworkSubnetId,omitempty" bson:"virtualNetworkSubnetId,omitempty"`
	// VirtualRouterAsn                    float64 `json:"virtualRouterAsn,omitempty" bson:"virtualRouterAsn,omitempty"`
	// VirtualRouterAutoScaleConfiguration *struct {
	// 	MinCapacity float64 `json:"minCapacity,omitempty" bson:"minCapacity,omitempty"`
	// } `json:"virtualRouterAutoScaleConfiguration,omitempty" bson:"virtualRouterAutoScaleConfiguration,omitempty"`
	// VirtualRouterIps []string `json:"virtualRouterIps,omitempty" bson:"virtualRouterIps,omitempty"`
	// VirtualWan       *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"virtualWan,omitempty" bson:"virtualWan,omitempty"`
	// Visibility any    `json:"visibility,omitempty" bson:"visibility,omitempty"`
	// VmID       string `json:"vmId,omitempty" bson:"vmId,omitempty"`
	// VmProfile  *struct {
	// 	OSDiskSizeGb float64 `json:"osDiskSizeGB,omitempty" bson:"osDiskSizeGB,omitempty"`
	// 	VmSize       string  `json:"vmSize,omitempty" bson:"vmSize,omitempty"`
	// } `json:"vmProfile,omitempty" bson:"vmProfile,omitempty"`
	// VnetBackupRestoreEnabled     bool                                `json:"vnetBackupRestoreEnabled,omitempty" bson:"vnetBackupRestoreEnabled,omitempty"`
	// VnetConfiguration            *AzureResourceVnetConfiguration     `json:"vnetConfiguration,omitempty" bson:"vnetConfiguration,omitempty"`
	// VnetConnectionsMax           float64                             `json:"vnetConnectionsMax,omitempty" bson:"vnetConnectionsMax,omitempty"`
	// VnetConnectionsUsed          float64                             `json:"vnetConnectionsUsed,omitempty" bson:"vnetConnectionsUsed,omitempty"`
	// VnetContentShareEnabled      bool                                `json:"vnetContentShareEnabled,omitempty" bson:"vnetContentShareEnabled,omitempty"`
	// VnetEncryptionSupported      bool                                `json:"vnetEncryptionSupported,omitempty" bson:"vnetEncryptionSupported,omitempty"`
	// VnetImagePullEnabled         bool                                `json:"vnetImagePullEnabled,omitempty" bson:"vnetImagePullEnabled,omitempty"`
	// VnetRouteAllEnabled          bool                                `json:"vnetRouteAllEnabled,omitempty" bson:"vnetRouteAllEnabled,omitempty"`
	// VoiceReceivers               []any                               `json:"voiceReceivers,omitempty" bson:"voiceReceivers,omitempty"`
	// VolumeBackups                []AzureResourceVolumeBackups        `json:"volumeBackups,omitempty" bson:"volumeBackups,omitempty"`
	// VolumeSpecName               string                              `json:"volumeSpecName,omitempty" bson:"volumeSpecName,omitempty"`
	// VolumeType                   string                              `json:"volumeType,omitempty" bson:"volumeType,omitempty"`
	// VolumesAssigned              float64                             `json:"volumesAssigned,omitempty" bson:"volumesAssigned,omitempty"`
	// VpnAuthenticationTypes       []string                            `json:"vpnAuthenticationTypes,omitempty" bson:"vpnAuthenticationTypes,omitempty"`
	// VpnClientIpsecPolicies       []AzureResourceVpnClientIpsecPolicy `json:"vpnClientIpsecPolicies,omitempty" bson:"vpnClientIpsecPolicies,omitempty"`
	// VpnClientRevokedCertificates []any                               `json:"vpnClientRevokedCertificates,omitempty" bson:"vpnClientRevokedCertificates,omitempty"`
	// VpnClientRootCertificates    []struct {
	// 	Name           string `json:"name,omitempty" bson:"name,omitempty"`
	// 	PublicCertData string `json:"publicCertData,omitempty" bson:"publicCertData,omitempty"`
	// } `json:"vpnClientRootCertificates,omitempty" bson:"vpnClientRootCertificates,omitempty"`
	// VpnGatewayDetachStatus string   `json:"vpnGatewayDetachStatus,omitempty" bson:"vpnGatewayDetachStatus,omitempty"`
	// VpnGatewayGeneration   string   `json:"vpnGatewayGeneration,omitempty" bson:"vpnGatewayGeneration,omitempty"`
	// VpnGatewayScaleUnit    float64  `json:"vpnGatewayScaleUnit,omitempty" bson:"vpnGatewayScaleUnit,omitempty"`
	// VpnProtocols           []string `json:"vpnProtocols,omitempty" bson:"vpnProtocols,omitempty"`
	// VpnServerConfiguration *struct {
	// 	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// } `json:"vpnServerConfiguration,omitempty" bson:"vpnServerConfiguration,omitempty"`
	// VpnServerConfigurationLocation string                          `json:"vpnServerConfigurationLocation,omitempty" bson:"vpnServerConfigurationLocation,omitempty"`
	// VpnType                        string                          `json:"vpnType,omitempty" bson:"vpnType,omitempty"`
	// WebSiteID                      any                             `json:"webSiteId,omitempty" bson:"webSiteId,omitempty"`
	// WebSpace                       string                          `json:"webSpace,omitempty" bson:"webSpace,omitempty"`
	// WebhookReceivers               []AzureResourceWebhookReceivers `json:"webhookReceivers,omitempty" bson:"webhookReceivers,omitempty"`
	// WeeklyBackupsToKeep            float64                         `json:"weeklyBackupsToKeep,omitempty" bson:"weeklyBackupsToKeep,omitempty"`
	// WeeklySchedule                 *AzureResourceWeeklySchedule    `json:"weeklySchedule,omitempty" bson:"weeklySchedule,omitempty"`
	// WindowSize                     string                          `json:"windowSize,omitempty" bson:"windowSize,omitempty"`
	// WindowsProfile                 *struct {
	// 	AdminUsername  string `json:"adminUsername,omitempty" bson:"adminUsername,omitempty"`
	// 	EnableCsiProxy bool   `json:"enableCSIProxy,omitempty" bson:"enableCSIProxy,omitempty"`
	// } `json:"windowsProfile,omitempty" bson:"windowsProfile,omitempty"`
	// WorkbookTemplates         []AzureResourceWorkbookTemplate `json:"workbookTemplates,omitempty" bson:"workbookTemplates,omitempty"`
	// WorkerSize                string                          `json:"workerSize,omitempty" bson:"workerSize,omitempty"`
	// WorkerSizeID              float64                         `json:"workerSizeId,omitempty" bson:"workerSizeId,omitempty"`
	// WorkerTierName            any                             `json:"workerTierName,omitempty" bson:"workerTierName,omitempty"`
	// WorkloadAutoScalerProfile *struct{}                       `json:"workloadAutoScalerProfile,omitempty" bson:"workloadAutoScalerProfile,omitempty"`
	// WorkloadProfileName       *string                         `json:"workloadProfileName,omitempty" bson:"workloadProfileName,omitempty"`
	// WorkloadProfiles          []struct {
	// 	Name                string `json:"name,omitempty" bson:"name,omitempty"`
	// 	WorkloadProfileType string `json:"workloadProfileType,omitempty" bson:"workloadProfileType,omitempty"`
	// } `json:"workloadProfiles,omitempty" bson:"workloadProfiles,omitempty"`
	// WorkspaceCapping         *AzureResourceWorkspaceCapping `json:"workspaceCapping,omitempty" bson:"workspaceCapping,omitempty"`
	// WorkspaceResourceIDOther string                         `json:"workspaceResourceId,omitempty" bson:"workspaceResourceId,omitempty"`
	// WriteLocations           []*AzureResourceWriteLocations `json:"writeLocations,omitempty" bson:"writeLocations,omitempty"`
	// ZoneRedundancy           string                         `json:"zoneRedundancy,omitempty" bson:"zoneRedundancy,omitempty"`
	// ZoneRedundant            bool                           `json:"zoneRedundant,omitempty" bson:"zoneRedundant,omitempty"`
	// ZoneType                 string                         `json:"zoneType,omitempty" bson:"zoneType,omitempty"`
}

type AzureVirtualMachineSku struct {
	LocationInfo []struct {
		Location string   `json:"location,omitempty"`
		Zones    []string `json:"zones,omitempty"`
	} `json:"locationInfo,omitempty"`
	Locations                                    []string  `json:"locations,omitempty"`
	Name                                         string    `json:"name,omitempty"`
	ResourceType                                 string    `json:"resourceType,omitempty"`
	Size                                         string    `json:"size,omitempty"`
	Tier                                         string    `json:"tier,omitempty"`
	MaxResourceVolumeMB                          string    `json:"maxResourceVolumeMB,omitempty"`
	OSVhdSizeMB                                  string    `json:"oSVhdSizeMB,omitempty"`
	VCPUs                                        string    `json:"vCPUs,omitempty"`
	MemoryPreservingMaintenanceSupported         string    `json:"memoryPreservingMaintenanceSupported,omitempty"`
	HyperVGenerations                            string    `json:"hyperVGenerations,omitempty"`
	MemoryGB                                     string    `json:"memoryGB,omitempty"`
	MaxDataDiskCount                             string    `json:"maxDataDiskCount,omitempty"`
	CpuArchitectureType                          string    `json:"cpuArchitectureType,omitempty"`
	LowPriorityCapable                           string    `json:"lowPriorityCapable,omitempty"`
	PremiumIO                                    string    `json:"premiumIO,omitempty"`
	VMDeploymentTypes                            string    `json:"vMDeploymentTypes,omitempty"`
	VCPUsAvailable                               string    `json:"vCPUsAvailable,omitempty"`
	ACUs                                         string    `json:"acus,omitempty"`
	VCPUsPerCore                                 string    `json:"vCPUsPerCore,omitempty"`
	CombinedTempDiskAndCachedIOPS                string    `json:"combinedTempDiskAndCachedIOPS,omitempty"`
	CombinedTempDiskAndCachedReadBytesPerSecond  string    `json:"combinedTempDiskAndCachedReadBytesPerSecond,omitempty"`
	CombinedTempDiskAndCachedWriteBytesPerSecond string    `json:"combinedTempDiskAndCachedWriteBytesPerSecond,omitempty"`
	UncachedDiskIOPS                             string    `json:"uncachedDiskIOPS,omitempty"`
	UncachedDiskBytesPerSecond                   string    `json:"uncachedDiskBytesPerSecond,omitempty"`
	EphemeralOSDiskSupported                     string    `json:"ephemeralOSDiskSupported,omitempty"`
	EncryptionAtHostSupported                    string    `json:"encryptionAtHostSupported,omitempty"`
	CapacityReservationSupported                 string    `json:"capacityReservationSupported,omitempty"`
	AcceleratedNetworkingEnabled                 string    `json:"acceleratedNetworkingEnabled,omitempty"`
	RdmaEnabled                                  string    `json:"rdmaEnabled,omitempty"`
	MaxNetworkInterfaces                         string    `json:"maxNetworkInterfaces,omitempty"`
	Cores                                        string    `json:"cores,omitempty"`
	SupportsAutoplacement                        string    `json:"supportsAutoplacement,omitempty"`
	LastAzureSync                                time.Time `json:"lastAzureSync,omitempty,omitzero" bson:"lastAzureSync,omitempty,omitzero"`
	LastDBSync                                   time.Time `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
}

// type AzureResourceSku struct {
// 	Capacity float64 `json:"capacity,omitempty" bson:"capacity,omitempty"`
// 	Family   string  `json:"family,omitempty" bson:"family,omitempty"`
// 	Name     string  `json:"name,omitempty" bson:"name,omitempty"`
// 	Size     string  `json:"size,omitempty" bson:"size,omitempty"`
// 	Tier     string  `json:"tier,omitempty" bson:"tier,omitempty"`
// }

type AzureResourceDetails struct {
	CostData                  map[string][]AggregatedCostItem `json:"costData,omitempty" bson:"costData,omitempty" fake:"-"`
	ExistsInAzure             bool                            `json:"existsInAzure" bson:"existsInAzure"`
	ExtendedLocation          any                             `json:"extendedLocation,omitempty" bson:"extendedLocation,omitempty"`
	ID                        string                          `json:"id,omitempty" bson:"_id,omitempty" fake:"{uuid}"`
	Identity                  *AzureResourceIdentity          `json:"identity,omitempty" bson:"identity,omitempty" fake:"-"`
	IsSqlRelated              bool                            `json:"isSqlRelated,omitempty" bson:"isSqlRelated,omitempty" fake:"{bool}"`
	Kind                      string                          `json:"kind,omitempty" bson:"kind,omitempty"`
	LastAzureSync             time.Time                       `json:"lastAzureSync,omitempty,omitzero" bson:"lastAzureSync,omitempty,omitzero" fake:"-"`
	LastDBSync                time.Time                       `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero" fake:"-"`
	Location                  string                          `json:"location,omitempty" bson:"location,omitempty" fake:"-"`
	ManagedBy                 string                          `json:"managedBy,omitempty" bson:"managedBy,omitempty" fake:"-"`
	Name                      string                          `json:"name,omitempty" bson:"name,omitempty" fake:"{username}"`
	Plan                      *AzureResourcePlan              `json:"plan,omitempty" bson:"plan,omitempty" fake:"-"`
	Properties                *AzureResourceProperties        `json:"properties,omitempty" bson:"properties,omitempty"`
	RelatedCostMeters         []string                        `json:"relatedCostMeters,omitempty" bson:"relatedCostMeters,omitempty" fake:"-"`
	RelatedCostMetersExpanded []MongoDbCostMeter              `json:"relatedCostMetersExpanded,omitempty" bson:"relatedCostMetersExpanded,omitempty" fake:"-"`
	RelatedResources          []string                        `json:"relatedResources,omitempty" bson:"relatedResources,omitempty" fake:"-"`
	RelatedResourcesExpanded  []AzureResourceDetails          `json:"relatedResourcesExpanded,omitempty" bson:"relatedResourcesExpanded,omitempty" fake:"-"`
	ResourceGroup             string                          `json:"resourceGroup,omitempty" bson:"resourceGroup,omitempty" fake:"{username}"`
	ResourceId                string                          `json:"resourceId,omitempty" bson:"resourceId,omitempty" fake:"{uuid}"`
	Sku                       *AzureResourceSku               `json:"sku,omitempty" bson:"sku,omitempty" fake:"-"`
	SubscriptionID            string                          `json:"subscriptionId,omitempty" bson:"subscriptionId,omitempty" fake:"{uuid}"`
	SubscriptionName          string                          `json:"subscriptionName,omitempty" bson:"subscriptionName,omitempty" fake:"{username}"`
	Tags                      map[string]string               `json:"tags" bson:"tags" fake:"-"`
	TenantID                  string                          `json:"tenantId,omitempty" bson:"tenantId,omitempty" fake:"{uuid}"`
	TenantName                string                          `json:"tenantName,omitempty" bson:"tenantName,omitempty" fake:"{username}"`
	Type                      string                          `json:"type,omitempty" bson:"type,omitempty" fake:"{username}"`
	WindowsType               string                          `json:"windowsType,omitempty" bson:"windowsType,omitempty" fake:"{randomstring:[desktop,server]}"`
	Zones                     []string                        `json:"zones,omitempty" bson:"zones,omitempty" fake:"-"`
}

// type AzureResourceDetailsWithCosting struct {
// 	ResourceDetails
// 	MeterData                  []AggregatedCostItem
// 	RelatedResources           []string
// 	RelatedCostMeters          []string
// 	RelatedCostMetersTotalCost float64
// 	// MeterData Aggre
// }

type AzureResourceSkuResp struct {
	Value []AzureResourceSku `json:"value" bson:"value"`
}

type AzureResourceSku struct {
	Capabilities []*struct {
		Name  string `json:"name,omitempty" bson:"name,omitempty"`
		Value string `json:"value,omitempty" bson:"value,omitempty"`
	} `json:"capabilities,omitempty" bson:"capabilities,omitempty"`
	Family       string `json:"family,omitempty" bson:"family,omitempty"`
	LocationInfo []*struct {
		Location    string `json:"location,omitempty" bson:"location,omitempty"`
		ZoneDetails []*struct {
			Capabilities []*struct {
				Name  string `json:"name,omitempty" bson:"name,omitempty"`
				Value string `json:"value,omitempty" bson:"value,omitempty"`
			} `json:"capabilities,omitempty" bson:"capabilities,omitempty"`
		} `json:"zoneDetails,omitempty" bson:"zoneDetails,omitempty"`
		Zones []*string `json:"zones,omitempty" bson:"zones,omitempty"`
	} `json:"locationInfo,omitempty" bson:"locationInfo,omitempty"`
	Locations    []*string `json:"locations,omitempty" bson:"locations,omitempty"`
	Name         string    `json:"name,omitempty" bson:"name,omitempty"`
	ResourceType string    `json:"resourceType,omitempty" bson:"resourceType,omitempty"`
	Restrictions []*struct {
		ReasonCode      string `json:"reasonCode,omitempty" bson:"reasonCode,omitempty"`
		RestrictionInfo *struct {
			Locations []string `json:"locations,omitempty" bson:"locations,omitempty"`
			Zones     []string `json:"zones,omitempty" bson:"zones,omitempty"`
		} `json:"restrictionInfo,omitempty" bson:"restrictionInfo,omitempty"`
		Type   string   `json:"type,omitempty" bson:"type,omitempty"`
		Values []string `json:"values,omitempty" bson:"values,omitempty"`
	} `json:"restrictions,omitempty" bson:"restrictions,omitempty"`
	Size           string    `json:"size,omitempty" bson:"size,omitempty"`
	Tier           string    `json:"tier,omitempty" bson:"tier,omitempty"`
	VMvCPUs        int       `json:"vMvCPUs,omitempty" bson:"vMvCPUs,omitempty"`
	VMCores        int       `json:"vMCores,omitempty" bson:"vMCores,omitempty"`
	VMvCPUsPerCore int       `json:"vMvCPUsPerCore,omitempty" bson:"vMvCPUsPerCore,omitempty"`
	LastAzureSync  time.Time `json:"lastAzureSync,omitempty,omitzero" bson:"lastAzureSync,omitempty,omitzero"`
	LastDBSync     time.Time `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
}

type AzureVirtualMachineSize struct {
	// 	Capabilities []*struct {
	// 		Name  string `json:"name,omitempty" bson:"name,omitempty"`
	// 		Value string `json:"value,omitempty" bson:"value,omitempty"`
	// 	} `json:"capabilities,omitempty" bson:"capabilities,omitempty"`
	// 	Family       string `json:"family,omitempty" bson:"family,omitempty"`
	// 	LocationInfo []*struct {
	// 		Location    string `json:"location,omitempty" bson:"location,omitempty"`
	// 		ZoneDetails []*struct {
	// 			Capabilities []*struct {
	// 				Name  string `json:"name,omitempty" bson:"name,omitempty"`
	// 				Value string `json:"value,omitempty" bson:"value,omitempty"`
	// 			} `json:"capabilities,omitempty" bson:"capabilities,omitempty"`
	// 		} `json:"zoneDetails,omitempty" bson:"zoneDetails,omitempty"`
	// 		Zones []*string `json:"zones,omitempty" bson:"zones,omitempty"`
	// 	} `json:"locationInfo,omitempty" bson:"locationInfo,omitempty"`
	// 	Locations    []*string `json:"locations,omitempty" bson:"locations,omitempty"`
	Name              string   `json:"name,omitempty" bson:"name,omitempty"`
	ResourceType      string   `json:"resourceType,omitempty" bson:"resourceType,omitempty"`
	MemoryGB          string   `json:"memoryGb,omitempty" bson:"memoryGb,omitempty"`
	HyperVGenerations []string `json:"hyperVGenerations,omitempty" bson:"hyperVGenerations,omitempty"`
	VCPUs             string   `json:"vCPUs,omitempty" bson:"vCPUs,omitempty"`
	// 	Restrictions []*struct {
	// 		ReasonCode      string `json:"reasonCode,omitempty" bson:"reasonCode,omitempty"`
	// 		RestrictionInfo *struct {
	// 			Locations []string `json:"locations,omitempty" bson:"locations,omitempty"`
	// 			Zones     []string `json:"zones,omitempty" bson:"zones,omitempty"`
	// 		} `json:"restrictionInfo,omitempty" bson:"restrictionInfo,omitempty"`
	// 		Type   string   `json:"type,omitempty" bson:"type,omitempty"`
	// 		Values []string `json:"values,omitempty" bson:"values,omitempty"`
	// 	} `json:"restrictions,omitempty" bson:"restrictions,omitempty"`
	Size string `json:"size,omitempty" bson:"size,omitempty"`
	Tier string `json:"tier,omitempty" bson:"tier,omitempty"`
}

type AzureRestorePointCollectionSource struct {
	ID       string `json:"id,omitempty" bson:"id,omitempty"`
	Location string `json:"location,omitempty" bson:"location,omitempty"`
}
