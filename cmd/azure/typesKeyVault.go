package azure

import (
	"time"

	"github.com/r3labs/diff/v3"
)

type ListKeyVaultSecretsResponse struct {
	NextLink string           `json:"nextLink,omitempty,omitzero" bson:"nextLink,omitempty,omitzero"`
	Value    []KeyVaultSecret `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

type ListKeyVaultCertsResponse struct {
	NextLink string                   `json:"nextLink,omitempty,omitzero" bson:"nextLink,omitempty,omitzero"`
	Value    []KeyVaultCertificateMin `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

type KeyVaultCertificateMin struct {
	Attributes struct {
		Created int64 `json:"created,omitempty,omitzero" bson:"created,omitempty,omitzero"`
		Enabled bool  `json:"enabled,omitempty,omitzero" bson:"enabled,omitempty,omitzero"`
		Exp     int64 `json:"exp,omitempty,omitzero" bson:"exp,omitempty,omitzero"`
		Nbf     int64 `json:"nbf,omitempty,omitzero" bson:"nbf,omitempty,omitzero"`
		Updated int64 `json:"updated,omitempty,omitzero" bson:"updated,omitempty,omitzero"`
	} `json:"attributes,omitempty,omitzero" bson:"attributes,omitempty,omitzero"`
	ID      string    `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Subject string    `json:"subject,omitempty,omitzero" bson:"subject,omitempty,omitzero"`
	Tags    *struct{} `json:"tags,omitempty" bson:"tags,omitempty"`
	X5T     string    `json:"x5t,omitempty,omitzero" bson:"x5t,omitempty,omitzero"`
}

type KeyVaultCertificate struct {
	Attributes struct {
		Created         float64 `json:"created,omitempty,omitzero" bson:"created,omitempty,omitzero"`
		Enabled         bool    `json:"enabled,omitempty,omitzero" bson:"enabled,omitempty,omitzero"`
		Exp             float64 `json:"exp,omitempty,omitzero" bson:"exp,omitempty,omitzero"`
		Nbf             float64 `json:"nbf,omitempty,omitzero" bson:"nbf,omitempty,omitzero"`
		RecoverableDays float64 `json:"recoverableDays,omitempty,omitzero" bson:"recoverableDays,omitempty,omitzero"`
		RecoveryLevel   string  `json:"recoveryLevel,omitempty,omitzero" bson:"recoveryLevel,omitempty,omitzero"`
		Updated         float64 `json:"updated,omitempty,omitzero" bson:"updated,omitempty,omitzero"`
	} `json:"attributes,omitempty,omitzero" bson:"attributes,omitempty,omitzero"`
	Cer    string `json:"cer,omitempty,omitzero" bson:"cer,omitempty,omitzero"`
	ID     string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Kid    string `json:"kid,omitempty,omitzero" bson:"kid,omitempty,omitzero"`
	Policy struct {
		Attributes struct {
			Created float64 `json:"created,omitempty,omitzero" bson:"created,omitempty,omitzero"`
			Enabled bool    `json:"enabled,omitempty,omitzero" bson:"enabled,omitempty,omitzero"`
			Updated float64 `json:"updated,omitempty,omitzero" bson:"updated,omitempty,omitzero"`
		} `json:"attributes,omitempty,omitzero" bson:"attributes,omitempty,omitzero"`
		ID     string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
		Issuer struct {
			Name string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
		} `json:"issuer,omitempty,omitzero" bson:"issuer,omitempty,omitzero"`
		KeyProps struct {
			Exportable bool    `json:"exportable,omitempty,omitzero" bson:"exportable,omitempty,omitzero"`
			KeySize    float64 `json:"key_size,omitempty,omitzero" bson:"key_size,omitempty,omitzero"`
			Kty        string  `json:"kty,omitempty,omitzero" bson:"kty,omitempty,omitzero"`
			ReuseKey   bool    `json:"reuse_key,omitempty,omitzero" bson:"reuse_key,omitempty,omitzero"`
		} `json:"key_props,omitempty,omitzero" bson:"key_props,omitempty,omitzero"`
		LifetimeActions []struct {
			Action struct {
				ActionType string `json:"action_type,omitempty,omitzero" bson:"action_type,omitempty,omitzero"`
			} `json:"action,omitempty,omitzero" bson:"action,omitempty,omitzero"`
			Trigger struct {
				LifetimePercentage float64 `json:"lifetime_percentage,omitempty,omitzero" bson:"lifetime_percentage,omitempty,omitzero"`
			} `json:"trigger,omitempty,omitzero" bson:"trigger,omitempty,omitzero"`
		} `json:"lifetime_actions,omitempty,omitzero" bson:"lifetime_actions,omitempty,omitzero"`
		SecretProps struct {
			ContentType string `json:"contentType,omitempty,omitzero" bson:"contentType,omitempty,omitzero"`
		} `json:"secret_props,omitempty,omitzero" bson:"secret_props,omitempty,omitzero"`
		X509Props struct {
			BasicConstraints struct {
				Ca bool `json:"ca,omitempty,omitzero" bson:"ca,omitempty,omitzero"`
			} `json:"basic_constraints,omitempty,omitzero" bson:"basic_constraints,omitempty,omitzero"`
			Ekus           []string `json:"ekus,omitempty,omitzero" bson:"ekus,omitempty,omitzero"`
			KeyUsage       []string `json:"key_usage,omitempty,omitzero" bson:"key_usage,omitempty,omitzero"`
			Subject        string   `json:"subject,omitempty,omitzero" bson:"subject,omitempty,omitzero"`
			ValidityMonths float64  `json:"validity_months,omitempty,omitzero" bson:"validity_months,omitempty,omitzero"`
		} `json:"x509_props,omitempty,omitzero" bson:"x509_props,omitempty,omitzero"`
	} `json:"policy,omitempty,omitzero" bson:"policy,omitempty,omitzero"`
	Sid string `json:"sid,omitempty,omitzero" bson:"sid,omitempty,omitzero"`
	X5T string `json:"x5t,omitempty,omitzero" bson:"x5t,omitempty,omitzero"`
}

type ListKeyVaultsForSubResponse struct {
	NextLink string     `json:"nextLink,omitempty,omitzero" bson:"nextLink,omitempty,omitzero"`
	Value    []KeyVault `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

type KeyVaultSecret struct {
	Attributes  KeyVaultSecretAttributes `json:"attributes,omitempty,omitzero" bson:"attributes,omitempty,omitzero"`
	ContentType string                   `json:"contentType,omitempty,omitzero" bson:"contentType,omitempty,omitzero"`
	ID          string                   `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Managed     bool                     `json:"managed,omitempty" bson:"managed,omitempty"`
	Tags        map[string]string        `json:"tags,omitempty,omitzero" bson:"tags,omitempty,omitzero"`
	Value       string                   `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
	Name        string                   `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
}

type KeyVaultSecretAttributes struct {
	Created         int64  `json:"created,omitempty,omitzero" bson:"created,omitempty,omitzero"`
	Enabled         bool   `json:"enabled,omitempty,omitzero" bson:"enabled,omitempty,omitzero"`
	Exp             *int64 `json:"exp,omitempty" bson:"exp,omitempty"`
	Nbf             int64  `json:"nbf,omitempty" bson:"nbf,omitempty"`
	RecoverableDays int64  `json:"recoverableDays,omitempty,omitzero" bson:"recoverableDays,omitempty,omitzero"`
	RecoveryLevel   string `json:"recoveryLevel,omitempty,omitzero" bson:"recoveryLevel,omitempty,omitzero"`
	Updated         int64  `json:"updated,omitempty,omitzero" bson:"updated,omitempty,omitzero"`
}

type GetKeyVaultsResponse struct {
	NextLink string     `json:"nextLink,omitempty,omitzero" bson:"nextLink,omitempty,omitzero"`
	Value    []KeyVault `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

type KeyVault struct {
	ID               string             `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Location         string             `json:"location,omitempty,omitzero" bson:"location,omitempty,omitzero"`
	Name             string             `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	TenantName       string             `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	TenantId         string             `json:"tenantId,omitempty,omitzero" bson:"tenantId,omitempty,omitzero"`
	SubscriptionName string             `json:"subscriptionName,omitempty,omitzero" bson:"subscriptionName,omitempty,omitzero"`
	SubscriptionId   string             `json:"subscriptionId,omitempty,omitzero" bson:"subscriptionId,omitempty,omitzero"`
	Properties       KeyVaultProperties `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
	SystemData       KeyVaultSystemData `json:"systemData,omitempty,omitzero" bson:"systemData,omitempty,omitzero"`
	Tags             map[string]string  `json:"tags,omitempty,omitzero" bson:"tags,omitempty,omitzero"`
	Type             string             `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
}
type KeyVaultSystemData struct {
	CreatedAt          time.Time `json:"createdAt,omitempty,omitzero" bson:"createdAt,omitempty,omitzero"`
	CreatedBy          string    `json:"createdBy,omitempty,omitzero" bson:"createdBy,omitempty,omitzero"`
	CreatedByType      string    `json:"createdByType,omitempty,omitzero" bson:"createdByType,omitempty,omitzero"`
	LastModifiedAt     time.Time `json:"lastModifiedAt,omitempty,omitzero" bson:"lastModifiedAt,omitempty,omitzero"`
	LastModifiedBy     string    `json:"lastModifiedBy,omitempty,omitzero" bson:"lastModifiedBy,omitempty,omitzero"`
	LastModifiedByType string    `json:"lastModifiedByType,omitempty,omitzero" bson:"lastModifiedByType,omitempty,omitzero"`
}

type KeyVaultProperties struct {
	AccessPolicies               []KeyVaultAccessPolicy               `json:"accessPolicies,omitempty,omitzero" bson:"accessPolicies,omitempty,omitzero"`
	EnablePurgeProtection        bool                                 `json:"enablePurgeProtection,omitempty,omitzero" bson:"enablePurgeProtection,omitempty,omitzero"`
	EnableRbacAuthorization      bool                                 `json:"enableRbacAuthorization,omitempty,omitzero" bson:"enableRbacAuthorization,omitempty,omitzero"`
	EnableSoftDelete             bool                                 `json:"enableSoftDelete,omitempty,omitzero" bson:"enableSoftDelete,omitempty,omitzero"`
	EnabledForDeployment         bool                                 `json:"enabledForDeployment,omitempty,omitzero" bson:"enabledForDeployment,omitempty,omitzero"`
	EnabledForDiskEncryption     bool                                 `json:"enabledForDiskEncryption,omitempty,omitzero" bson:"enabledForDiskEncryption,omitempty,omitzero"`
	EnabledForTemplateDeployment bool                                 `json:"enabledForTemplateDeployment,omitempty,omitzero" bson:"enabledForTemplateDeployment,omitempty,omitzero"`
	NetworkAcls                  KeyVaultNetworkACLs                  `json:"networkAcls,omitempty,omitzero" bson:"networkAcls,omitempty,omitzero"`
	PrivateEndpointConnections   []KeyVaultPrivateEndpointConnections `json:"privateEndpointConnections,omitempty,omitzero" bson:"privateEndpointConnections,omitempty,omitzero"`
	ProvisioningState            string                               `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
	PublicNetworkAccess          string                               `json:"publicNetworkAccess,omitempty,omitzero" bson:"publicNetworkAccess,omitempty,omitzero"`
	Sku                          struct {
		Family string `json:"family,omitempty,omitzero" bson:"family,omitempty,omitzero"`
		Name   string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	} `json:"sku,omitempty,omitzero" bson:"sku,omitempty,omitzero"`
	SoftDeleteRetentionInDays float64 `json:"softDeleteRetentionInDays,omitempty,omitzero" bson:"softDeleteRetentionInDays,omitempty,omitzero"`
	TenantID                  string  `json:"tenantId,omitempty,omitzero" bson:"tenantId,omitempty,omitzero"`
	VaultURI                  string  `json:"vaultUri,omitempty,omitzero" bson:"vaultUri,omitempty,omitzero"`
}

type KeyVaultAccessPolicy struct {
	ObjectID    string                          `json:"objectId,omitempty,omitzero" bson:"objectId,omitempty,omitzero"`
	Permissions KeyVaultAccessPolicyPermissions `json:"permissions,omitempty,omitzero" bson:"permissions,omitempty,omitzero"`
	TenantID    string                          `json:"tenantId,omitempty,omitzero" bson:"tenantId,omitempty,omitzero"`
}

type KeyVaultAccessPolicyPermissions struct {
	Certificates []string `json:"certificates,omitempty,omitzero" bson:"certificates,omitempty,omitzero"`
	Keys         []string `json:"keys,omitempty,omitzero" bson:"keys,omitempty,omitzero"`
	Secrets      []string `json:"secrets,omitempty,omitzero" bson:"secrets,omitempty,omitzero"`
	Storage      []string `json:"storage,omitempty,omitzero" bson:"storage,omitempty,omitzero"`
}

type KeyVaultNetworkACLs struct {
	Bypass        string `json:"bypass,omitempty,omitzero" bson:"bypass,omitempty,omitzero"`
	DefaultAction string `json:"defaultAction,omitempty,omitzero" bson:"defaultAction,omitempty,omitzero"`
	IpRules       []struct {
		Value string `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
	} `json:"ipRules,omitempty,omitzero" bson:"ipRules,omitempty,omitzero"`
	VirtualNetworkRules []KeyVaultNetworkACLVirtualNetworkRule `json:"virtualNetworkRules,omitempty,omitzero" bson:"virtualNetworkRules,omitempty,omitzero"`
}

type KeyVaultNetworkACLVirtualNetworkRule struct {
	ID                               string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	IgnoreMissingVnetServiceEndpoint bool   `json:"ignoreMissingVnetServiceEndpoint,omitempty,omitzero" bson:"ignoreMissingVnetServiceEndpoint,omitempty,omitzero"`
}

type KeyVaultPrivateEndpointConnections struct {
	ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Properties struct {
		PrivateEndpoint struct {
			ID string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
		} `json:"privateEndpoint,omitempty,omitzero" bson:"privateEndpoint,omitempty,omitzero"`
		PrivateLinkServiceConnectionState struct {
			ActionsRequired string `json:"actionsRequired,omitempty,omitzero" bson:"actionsRequired,omitempty,omitzero"`
			Description     string `json:"description,omitempty,omitzero" bson:"description,omitempty,omitzero"`
			Status          string `json:"status,omitempty,omitzero" bson:"status,omitempty,omitzero"`
		} `json:"privateLinkServiceConnectionState,omitempty,omitzero" bson:"privateLinkServiceConnectionState,omitempty,omitzero"`
		ProvisioningState string `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
	} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
}

type T struct {
	ID         string             `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Location   string             `json:"location,omitempty,omitzero" bson:"location,omitempty,omitzero"`
	Name       string             `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Properties KeyVaultProperties `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
	SystemData struct {
		LastModifiedAt     time.Time `json:"lastModifiedAt,omitempty,omitzero" bson:"lastModifiedAt,omitempty,omitzero"`
		LastModifiedBy     string    `json:"lastModifiedBy,omitempty,omitzero" bson:"lastModifiedBy,omitempty,omitzero"`
		LastModifiedByType string    `json:"lastModifiedByType,omitempty,omitzero" bson:"lastModifiedByType,omitempty,omitzero"`
	} `json:"systemData,omitempty,omitzero" bson:"systemData,omitempty,omitzero"`
	Tags struct {
		// "Business Onwer" cannot be unmarshalled into a struct field by encoding/json.
		// "Created By" cannot be unmarshalled into a struct field by encoding/json.
		// "Technical Onwer" cannot be unmarshalled into a struct field by encoding/json.
	} `json:"tags,omitempty,omitzero" bson:"tags,omitempty,omitzero"`
	Type string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
}

type KeyVaultSecretStored struct {
	Name             string    `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Id               string    `json:"id,omitempty,omitzero" bson:"_id,omitempty,omitzero"`
	Expiration       time.Time `json:"expiration,omitempty,omitzero" bson:"expiration,omitempty,omitzero"`
	KeyVaultId       string    `json:"keyVaultId,omitempty,omitzero" bson:"keyVaultId,omitempty,omitzero"`
	KeyVaultName     string    `json:"keyVaultName,omitempty,omitzero" bson:"keyVaultName,omitempty,omitzero"`
	KeyVaultUrl      string    `json:"keyVaultUrl,omitempty,omitzero" bson:"keyVaultUrl,omitempty,omitzero"`
	TenantName       string    `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	SubscriptionName string    `json:"subscriptionName,omitempty,omitzero" bson:"subscriptionName,omitempty,omitzero"`
	Type             string    `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
	ContentType      string    `json:"contentType,omitempty,omitzero" bson:"contentType,omitempty,omitzero"`
}

type KeyVaultUpdateComparison struct {
	Id         string        `json:"id"`
	Original   KeyVault      `json:"original"`
	Updated    KeyVault      `json:"updated"`
	Diff       diff.PatchLog `json:"diff"`
	DiffString string        `json:"diffString"`
}
