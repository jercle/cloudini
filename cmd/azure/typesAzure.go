package azure

import "time"

type AzureRequestOptions struct {
	SubscriptionId    string
	ResourceId        string
	ResourceGroupName string
	ResourceName      string
	TenantId          string
	TenantName        string

	ConfigFilePath string
}

//
//

type ListAllResourcesResponse struct {
	Value    []ListRspResource `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
	NextLink string            `json:"nextLink,omitempty,omitzero" bson:"nextLink,omitempty,omitzero"`
}

//
//

//
//

type ListRspResource struct {
	ID       string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Identity *struct {
		PrincipalID string `json:"principalId,omitempty,omitzero" bson:"principalId,omitempty,omitzero"`
		TenantID    string `json:"tenantId,omitempty,omitzero" bson:"tenantId,omitempty,omitzero"`
		Type        string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
	} `json:"identity,omitempty" bson:"identity,omitempty"`
	Location  string `json:"location,omitempty,omitzero" bson:"location,omitempty,omitzero"`
	ManagedBy string `json:"managedBy,omitempty" bson:"managedBy,omitempty"`
	Name      string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Sku       *struct {
		Name string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
		Tier string `json:"tier,omitempty,omitzero" bson:"tier,omitempty,omitzero"`
	} `json:"sku,omitempty" bson:"sku,omitempty"`
	Tags  map[string]string `json:"tags," bson:"tags"`
	Type  string            `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
	Zones []string          `json:"zones,omitempty" bson:"zones,omitempty"`
}

//
//

type EntraListApplicationsResponse struct {
	OdataContext string             `json:"@odata.context,omitempty" bson:"@odata.context,omitempty"`
	NextLink     string             `json:"@odata.nextLink,omitempty" bson:"@odata.nextLink,omitempty"`
	Value        []EntraApplication `json:"value,omitempty" bson:"value,omitempty"`
}

//
//

type EntraApplication struct {
	AddIns *[]struct {
		ID         string `json:"id,omitempty" bson:"id,omitempty"`
		Properties []struct {
			Key   string `json:"key,omitempty" bson:"key,omitempty"`
			Value string `json:"value,omitempty" bson:"value,omitempty"`
		} `json:"properties,omitempty" bson:"properties,omitempty"`
		Type string `json:"type,omitempty" bson:"type,omitempty"`
	} `json:"addIns,omitempty" bson:"addIns,omitempty"`
	API *struct {
		AcceptMappedClaims      *bool     `json:"acceptMappedClaims,omitempty" bson:"acceptMappedClaims,omitempty"`
		KnownClientApplications *[]string `json:"knownClientApplications,omitempty" bson:"knownClientApplications,omitempty"`
		Oauth2PermissionScopes  *[]struct {
			AdminConsentDescription string `json:"adminConsentDescription,omitempty" bson:"adminConsentDescription,omitempty"`
			AdminConsentDisplayName string `json:"adminConsentDisplayName,omitempty" bson:"adminConsentDisplayName,omitempty"`
			ID                      string `json:"id,omitempty" bson:"id,omitempty"`
			IsEnabled               bool   `json:"isEnabled,omitempty" bson:"isEnabled,omitempty"`
			IsPrivate               bool   `json:"isPrivate,omitempty" bson:"isPrivate,omitempty"`
			Type                    string `json:"type,omitempty" bson:"type,omitempty"`
			UserConsentDescription  string `json:"userConsentDescription,omitempty" bson:"userConsentDescription,omitempty"`
			UserConsentDisplayName  string `json:"userConsentDisplayName,omitempty" bson:"userConsentDisplayName,omitempty"`
			Value                   string `json:"value,omitempty" bson:"value,omitempty"`
		} `json:"oauth2PermissionScopes,omitempty" bson:"oauth2PermissionScopes,omitempty"`
		PreAuthorizedApplications *[]struct {
			AppID         string    `json:"appId,omitempty" bson:"appId,omitempty"`
			PermissionIds *[]string `json:"permissionIds,omitempty" bson:"permissionIds,omitempty"`
		} `json:"preAuthorizedApplications,omitempty" bson:"preAuthorizedApplications,omitempty"`
		RequestedAccessTokenVersion *float64 `json:"requestedAccessTokenVersion,omitempty" bson:"requestedAccessTokenVersion,omitempty"`
		TokenEncryptionSetting      *struct {
			Audience              any `json:"audience,omitempty" bson:"audience,omitempty"`
			AutomatedTokenVersion *struct {
				Available []any `json:"available,omitempty" bson:"available,omitempty"`
				Current   any   `json:"current,omitempty" bson:"current,omitempty"`
			} `json:"automatedTokenVersion,omitempty" bson:"automatedTokenVersion,omitempty"`
			Scheme any `json:"scheme,omitempty" bson:"scheme,omitempty"`
		} `json:"tokenEncryptionSetting,omitempty" bson:"tokenEncryptionSetting,omitempty"`
	} `json:"api,omitempty" bson:"api,omitempty"`
	AppCapabilities []any  `json:"appCapabilities,omitempty" bson:"appCapabilities,omitempty"`
	AppID           string `json:"appId,omitempty" bson:"appId,omitempty"`
	AppRoles        *[]struct {
		AllowedMemberTypes         []string `json:"allowedMemberTypes,omitempty" bson:"allowedMemberTypes,omitempty"`
		Description                string   `json:"description,omitempty" bson:"description,omitempty"`
		DisplayName                string   `json:"displayName,omitempty" bson:"displayName,omitempty"`
		ID                         string   `json:"id,omitempty" bson:"id,omitempty"`
		IsEnabled                  bool     `json:"isEnabled,omitempty" bson:"isEnabled,omitempty"`
		IsPreAuthorizationRequired bool     `json:"isPreAuthorizationRequired,omitempty" bson:"isPreAuthorizationRequired,omitempty"`
		IsPrivate                  bool     `json:"isPrivate,omitempty" bson:"isPrivate,omitempty"`
		Origin                     string   `json:"origin,omitempty" bson:"origin,omitempty"`
		Value                      *string  `json:"value,omitempty" bson:"value,omitempty"`
	} `json:"appRoles,omitempty" bson:"appRoles,omitempty"`
	ApplicationTemplateID     *string   `json:"applicationTemplateId,omitempty" bson:"applicationTemplateId,omitempty"`
	BillingInformation        any       `json:"billingInformation,omitempty" bson:"billingInformation,omitempty"`
	Certification             any       `json:"certification,omitempty" bson:"certification,omitempty"`
	CreatedDateTime           time.Time `json:"createdDateTime,omitempty" bson:"createdDateTime,omitempty"`
	DefaultRedirectURI        *string   `json:"defaultRedirectUri,omitempty" bson:"defaultRedirectUri,omitempty"`
	DeletedDateTime           any       `json:"deletedDateTime,omitempty" bson:"deletedDateTime,omitempty"`
	Description               string    `json:"description,omitempty" bson:"description,omitempty"`
	DisabledByMicrosoftStatus any       `json:"disabledByMicrosoftStatus,omitempty" bson:"disabledByMicrosoftStatus,omitempty"`
	DisplayName               string    `json:"displayName,omitempty" bson:"displayName,omitempty"`
	GroupMembershipClaims     *string   `json:"groupMembershipClaims,omitempty" bson:"groupMembershipClaims,omitempty"`
	ID                        string    `json:"id,omitempty" bson:"objectId,omitempty"`
	IdentifierUris            []string  `json:"identifierUris,omitempty" bson:"identifierUris,omitempty"`
	Info                      *struct {
		LogoURL             any `json:"logoUrl,omitempty" bson:"logoUrl,omitempty"`
		MarketingURL        any `json:"marketingUrl,omitempty" bson:"marketingUrl,omitempty"`
		PrivacyStatementURL any `json:"privacyStatementUrl,omitempty" bson:"privacyStatementUrl,omitempty"`
		SupportURL          any `json:"supportUrl,omitempty" bson:"supportUrl,omitempty"`
		TermsOfServiceURL   any `json:"termsOfServiceUrl,omitempty" bson:"termsOfServiceUrl,omitempty"`
	} `json:"info,omitempty" bson:"info,omitempty"`
	IsAuthorizationServiceEnabled bool  `json:"isAuthorizationServiceEnabled,omitempty" bson:"isAuthorizationServiceEnabled,omitempty"`
	IsDeviceOnlyAuthSupported     *bool `json:"isDeviceOnlyAuthSupported,omitempty" bson:"isDeviceOnlyAuthSupported,omitempty"`
	IsDisabled                    any   `json:"isDisabled,omitempty" bson:"isDisabled,omitempty"`
	IsFallbackPublicClient        *bool `json:"isFallbackPublicClient,omitempty" bson:"isFallbackPublicClient,omitempty"`
	IsManagementRestricted        any   `json:"isManagementRestricted,omitempty" bson:"isManagementRestricted,omitempty"`
	KeyCredentials                *[]struct {
		CustomKeyIdentifier string    `json:"customKeyIdentifier,omitempty" bson:"customKeyIdentifier,omitempty"`
		DisplayName         string    `json:"displayName,omitempty" bson:"displayName,omitempty"`
		EndDateTime         time.Time `json:"endDateTime,omitempty" bson:"endDateTime,omitempty"`
		HasExtendedValue    any       `json:"hasExtendedValue,omitempty" bson:"hasExtendedValue,omitempty"`
		Key                 any       `json:"key,omitempty" bson:"key,omitempty"`
		KeyID               string    `json:"keyId,omitempty" bson:"keyId,omitempty"`
		StartDateTime       time.Time `json:"startDateTime,omitempty" bson:"startDateTime,omitempty"`
		Type                string    `json:"type,omitempty" bson:"type,omitempty"`
		Usage               string    `json:"usage,omitempty" bson:"usage,omitempty"`
	} `json:"keyCredentials,omitempty" bson:"keyCredentials,omitempty"`
	MigrationStatus                 any  `json:"migrationStatus,omitempty" bson:"migrationStatus,omitempty"`
	NativeAuthenticationApisEnabled any  `json:"nativeAuthenticationApisEnabled,omitempty" bson:"nativeAuthenticationApisEnabled,omitempty"`
	Notes                           any  `json:"notes,omitempty" bson:"notes,omitempty"`
	Oauth2RequirePostResponse       bool `json:"oauth2RequirePostResponse,omitempty" bson:"oauth2RequirePostResponse,omitempty"`
	OptionalClaims                  *struct {
		AccessToken []any `json:"accessToken,omitempty" bson:"accessToken,omitempty"`
		IDToken     []any `json:"idToken,omitempty" bson:"idToken,omitempty"`
		Saml2Token  []struct {
			AdditionalProperties []any  `json:"additionalProperties,omitempty" bson:"additionalProperties,omitempty"`
			Essential            bool   `json:"essential,omitempty" bson:"essential,omitempty"`
			Name                 string `json:"name,omitempty" bson:"name,omitempty"`
			Source               any    `json:"source,omitempty" bson:"source,omitempty"`
		} `json:"saml2Token,omitempty" bson:"saml2Token,omitempty"`
	} `json:"optionalClaims,omitempty" bson:"optionalClaims,omitempty"`
	OrgRestrictions         []any `json:"orgRestrictions,omitempty" bson:"orgRestrictions,omitempty"`
	ParentalControlSettings *struct {
		CountriesBlockedForMinors []any  `json:"countriesBlockedForMinors,omitempty" bson:"countriesBlockedForMinors,omitempty"`
		LegalAgeGroupRule         string `json:"legalAgeGroupRule,omitempty" bson:"legalAgeGroupRule,omitempty"`
	} `json:"parentalControlSettings,omitempty" bson:"parentalControlSettings,omitempty"`
	PasswordCredentials *[]struct {
		CustomKeyIdentifier string    `json:"customKeyIdentifier,omitempty" bson:"customKeyIdentifier,omitempty"`
		DisplayName         string    `json:"displayName,omitempty" bson:"displayName,omitempty"`
		EndDateTime         time.Time `json:"endDateTime,omitempty" bson:"endDateTime,omitempty"`
		Hint                string    `json:"hint,omitempty" bson:"hint,omitempty"`
		KeyID               string    `json:"keyId,omitempty" bson:"keyId,omitempty"`
		SecretText          any       `json:"secretText,omitempty" bson:"secretText,omitempty"`
		StartDateTime       time.Time `json:"startDateTime,omitempty" bson:"startDateTime,omitempty"`
	} `json:"passwordCredentials,omitempty" bson:"passwordCredentials,omitempty"`
	PublicClient *struct {
		RedirectUris []any `json:"redirectUris,omitempty" bson:"redirectUris,omitempty"`
	} `json:"publicClient,omitempty" bson:"publicClient,omitempty"`
	PublisherDomain              string `json:"publisherDomain,omitempty" bson:"publisherDomain,omitempty"`
	RequestSignatureVerification *struct {
		AllowedWeakAlgorithms   any  `json:"allowedWeakAlgorithms,omitempty" bson:"allowedWeakAlgorithms,omitempty"`
		IsSignedRequestRequired bool `json:"isSignedRequestRequired,omitempty" bson:"isSignedRequestRequired,omitempty"`
	} `json:"requestSignatureVerification,omitempty" bson:"requestSignatureVerification,omitempty"`
	RequiredResourceAccess *[]struct {
		ResourceAccess *[]struct {
			ID   string `json:"id,omitempty" bson:"id,omitempty"`
			Type string `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"resourceAccess,omitempty" bson:"resourceAccess,omitempty"`
		ResourceAppID string `json:"resourceAppId,omitempty" bson:"resourceAppId,omitempty"`
	} `json:"requiredResourceAccess,omitempty" bson:"requiredResourceAccess,omitempty"`
	SamlMetadataURL                   any `json:"samlMetadataUrl,omitempty" bson:"samlMetadataUrl,omitempty"`
	ServiceManagementReference        any `json:"serviceManagementReference,omitempty" bson:"serviceManagementReference,omitempty"`
	ServicePrincipalLockConfiguration *struct {
		AllProperties              bool `json:"allProperties,omitempty" bson:"allProperties,omitempty"`
		CredentialsWithUsageSign   bool `json:"credentialsWithUsageSign,omitempty" bson:"credentialsWithUsageSign,omitempty"`
		CredentialsWithUsageVerify bool `json:"credentialsWithUsageVerify,omitempty" bson:"credentialsWithUsageVerify,omitempty"`
		IdentifierUris             bool `json:"identifierUris,omitempty" bson:"identifierUris,omitempty"`
		IsEnabled                  bool `json:"isEnabled,omitempty" bson:"isEnabled,omitempty"`
		TokenEncryptionKeyID       bool `json:"tokenEncryptionKeyId,omitempty" bson:"tokenEncryptionKeyId,omitempty"`
	} `json:"servicePrincipalLockConfiguration,omitempty" bson:"servicePrincipalLockConfiguration,omitempty"`
	SignInAudience string `json:"signInAudience,omitempty" bson:"signInAudience,omitempty"`
	Spa            *struct {
		RedirectUris *[]string `json:"redirectUris,omitempty" bson:"redirectUris,omitempty"`
	} `json:"spa,omitempty" bson:"spa,omitempty"`
	Tags                 *[]string `json:"tags" bson:"tags"`
	TenantName           string    `json:"tenantName,omitempty" bson:"tenantName,omitempty"`
	TenantId             string    `json:"tenantId,omitempty" bson:"tenantId,omitempty"`
	TokenEncryptionKeyID *string   `json:"tokenEncryptionKeyId,omitempty" bson:"tokenEncryptionKeyId,omitempty"`
	UniqueName           any       `json:"uniqueName,omitempty" bson:"uniqueName,omitempty"`
	VerifiedPublisher    *struct {
		AddedDateTime       any `json:"addedDateTime,omitempty" bson:"addedDateTime,omitempty"`
		DisplayName         any `json:"displayName,omitempty" bson:"displayName,omitempty"`
		VerifiedPublisherID any `json:"verifiedPublisherId,omitempty" bson:"verifiedPublisherId,omitempty"`
	} `json:"verifiedPublisher,omitempty" bson:"verifiedPublisher,omitempty"`
	Web *struct {
		HomePageURL           *string `json:"homePageUrl,omitempty" bson:"homePageUrl,omitempty"`
		ImplicitGrantSettings struct {
			EnableAccessTokenIssuance bool `json:"enableAccessTokenIssuance,omitempty" bson:"enableAccessTokenIssuance,omitempty"`
			EnableIDTokenIssuance     bool `json:"enableIdTokenIssuance,omitempty" bson:"enableIdTokenIssuance,omitempty"`
		} `json:"implicitGrantSettings,omitempty" bson:"implicitGrantSettings,omitempty"`
		LogoutURL           any `json:"logoutUrl,omitempty" bson:"logoutUrl,omitempty"`
		RedirectURISettings *[]struct {
			Index any    `json:"index,omitempty" bson:"index,omitempty"`
			URI   string `json:"uri,omitempty" bson:"uri,omitempty"`
		} `json:"redirectUriSettings,omitempty" bson:"redirectUriSettings,omitempty"`
		RedirectUris []string `json:"redirectUris,omitempty" bson:"redirectUris,omitempty"`
	} `json:"web,omitempty" bson:"web,omitempty"`
	Windows       any       `json:"windows,omitempty" bson:"windows,omitempty"`
	MongoDbId     string    `json:"_id,omitempty" bson:"_id,omitempty"`
	LastAzureSync time.Time `json:"lastAzureSync,omitempty,omitzero" bson:"lastAzureSync,omitempty,omitzero" fake:"-"`
	LastDBSync    time.Time `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero" fake:"-"`
}

//
//

type EntraExpiringCredential struct {
	AppRegAppID             string    `json:"appId,omitempty" bson:"appId,omitempty"`
	AppRegDescription       string    `json:"description,omitempty" bson:"description,omitempty"`
	AppRegDisplayName       string    `json:"appRegDisplayName,omitempty" bson:"appRegDisplayName,omitempty"`
	AppRegCreatedDateTime   time.Time `json:"appRegCreatedDateTime,omitempty" bson:"appRegCreatedDateTime,omitempty"`
	AppRegObjectID          string    `json:"appRegObjectId,omitempty" bson:"appRegObjectId,omitempty"`
	CredCustomKeyIdentifier string    `json:"credCustomKeyIdentifier,omitempty" bson:"credCustomKeyIdentifier,omitempty"`
	CredDisplayName         string    `json:"credDisplayName,omitempty" bson:"credDisplayName,omitempty"`
	CredEndDateTime         time.Time `json:"credEndDateTime,omitempty" bson:"credEndDateTime,omitempty"`
	CredID                  string    `json:"credId,omitempty" bson:"credId,omitempty"`
	CredKeyType             string    `json:"credKeyType,omitempty" bson:"credKeyType,omitempty"`
	CredKeyUsage            string    `json:"credKeyUsage,omitempty" bson:"credKeyUsage,omitempty"`
	CredStartDateTime       time.Time `json:"credStartDateTime,omitempty" bson:"credStartDateTime,omitempty"`
	CredType                string    `json:"credType,omitempty" bson:"credType,omitempty"`
	MongoDbId               string    `json:"_id,omitempty" bson:"_id,omitempty"`
	TenantName              string    `json:"tenantName,omitempty" bson:"tenantName,omitempty"`
	TenantId                string    `json:"tenantId,omitempty" bson:"tenantId,omitempty"`
	LastAzureSync           time.Time `json:"lastAzureSync,omitempty,omitzero" bson:"lastAzureSync,omitempty,omitzero" fake:"-"`
	LastDBSync              time.Time `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero" fake:"-"`
}

//
//

type AzureB2CUser struct {
	AccountEnabled                                           bool      `json:"accountEnabled,omitempty,omitzero" bson:"accountEnabled,omitempty,omitzero"`
	CreatedDateTime                                          time.Time `json:"createdDateTime,omitempty,omitzero" bson:"createdDateTime,omitempty,omitzero"`
	CreationType                                             string    `json:"creationType,omitempty,omitzero" bson:"creationType,omitempty,omitzero"`
	DisplayName                                              string    `json:"displayName,omitempty,omitzero" bson:"displayName,omitempty,omitzero"`
	Extension4e4fa41c1d3246639764b37ff949534dLastLogonTime   time.Time `json:"extension_4e4fa41c1d3246639764b37ff949534d_lastLogonTime,omitempty,omitzero" bson:"extension_4e4fa41c1d3246639764b37ff949534d_lastLogonTime,omitempty,omitzero"`
	Extension4e4fa41c1d3246639764b37ff949534dPasswordResetOn time.Time `json:"extension_4e4fa41c1d3246639764b37ff949534d_passwordResetOn,omitempty,omitzero" bson:"extension_4e4fa41c1d3246639764b37ff949534d_passwordResetOn,omitempty,omitzero"`
	ID                                                       string    `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	UserPrincipalName                                        string    `json:"userPrincipalName,omitempty,omitzero" bson:"userPrincipalName,omitempty,omitzero"`
}

//
//

type TenantList []TenantDetails

//
//

type TenantDetails struct {
	TenantId      string            `json:"id" bson:"_i,omitempty,omitzero"`
	TenantName    string            `json:"tenantName" bson:"tenantNam,omitempty,omitzero"`
	Subscriptions map[string]string `json:"subscriptions" bson:"subscription,omitempty,omitzero"`
}

//
//

type GetAllImageGalleriesForSubscriptionResponse struct {
	Value []ImageGallery `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

//
//

type ImageGallery struct {
	ID             string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Location       string `json:"location,omitempty,omitzero" bson:"location,omitempty,omitzero"`
	Name           string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	SubscriptionId string `json:"subscriptionId,omitempty,omitzero" bson:"subscriptionId,omitempty,omitzero"`
	ResourceGroup  string `json:"resourceGroup,omitempty,omitzero" bson:"resourceGroup,omitempty,omitzero"`
	TenantName     string `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	Properties     struct {
		Description string `json:"description,omitempty" bson:"description,omitempty"`
		Identifier  struct {
			UniqueName string `json:"uniqueName,omitempty,omitzero" bson:"uniqueName,omitempty,omitzero"`
		} `json:"identifier,omitempty,omitzero" bson:"identifier,omitempty,omitzero"`
		ProvisioningState string `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
		SoftDeletePolicy  *struct {
			IsSoftDeleteEnabled bool `json:"isSoftDeleteEnabled,omitempty,omitzero" bson:"isSoftDeleteEnabled,omitempty,omitzero"`
		} `json:"softDeletePolicy,omitempty" bson:"softDeletePolicy,omitempty"`
	} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
	Tags map[string]string `json:"tags" bson:"tags"`
	Type string            `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
}

//
//

type ListManagementGroupsResponse struct {
	Value []ManagementGroup `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

//
//

type ManagementGroup struct {
	ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Properties struct {
		DisplayName string `json:"displayName,omitempty,omitzero" bson:"displayName,omitempty,omitzero"`
		TenantID    string `json:"tenantId,omitempty,omitzero" bson:"tenantId,omitempty,omitzero"`
	} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
	Type string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
}

//
//

type ResourceRoleDefinition struct {
	ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Properties struct {
		AssignableScopes []string  `json:"assignableScopes,omitempty,omitzero" bson:"assignableScopes,omitempty,omitzero"`
		CreatedBy        any       `json:"createdBy,omitempty,omitzero" bson:"createdBy,omitempty,omitzero"`
		CreatedOn        time.Time `json:"createdOn,omitempty,omitzero" bson:"createdOn,omitempty,omitzero"`
		Description      string    `json:"description,omitempty,omitzero" bson:"description,omitempty,omitzero"`
		Permissions      []struct {
			Actions        []string `json:"actions,omitempty,omitzero" bson:"actions,omitempty,omitzero"`
			DataActions    []any    `json:"dataActions,omitempty,omitzero" bson:"dataActions,omitempty,omitzero"`
			NotActions     []any    `json:"notActions,omitempty,omitzero" bson:"notActions,omitempty,omitzero"`
			NotDataActions []any    `json:"notDataActions,omitempty,omitzero" bson:"notDataActions,omitempty,omitzero"`
		} `json:"permissions,omitempty,omitzero" bson:"permissions,omitempty,omitzero"`
		RoleName  string    `json:"roleName,omitempty,omitzero" bson:"roleName,omitempty,omitzero"`
		Type      string    `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
		UpdatedBy any       `json:"updatedBy,omitempty,omitzero" bson:"updatedBy,omitempty,omitzero"`
		UpdatedOn time.Time `json:"updatedOn,omitempty,omitzero" bson:"updatedOn,omitempty,omitzero"`
	} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
	Type string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
}

//
//

type ListResourceRoleDefinitionsResponse struct {
	Value []ResourceRoleDefinition `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

//
//

type AzureAsyncRequestResponse struct {
	AzureAsyncnotification                        []string `json:"Azure-Asyncnotification,omitempty,omitzero" bson:"Azure-Asyncnotification,omitempty,omitzero"`
	AzureAsyncoperation                           []string `json:"Azure-Asyncoperation,omitempty,omitzero" bson:"Azure-Asyncoperation,omitempty,omitzero"`
	CacheControl                                  []string `json:"Cache-Control,omitempty,omitzero" bson:"Cache-Control,omitempty,omitzero"`
	ContentLength                                 []string `json:"Content-Length,omitempty,omitzero" bson:"Content-Length,omitempty,omitzero"`
	Date                                          []string `json:"Date,omitempty,omitzero" bson:"Date,omitempty,omitzero"`
	Expires                                       []string `json:"Expires,omitempty,omitzero" bson:"Expires,omitempty,omitzero"`
	Location                                      []string `json:"Location,omitempty,omitzero" bson:"Location,omitempty,omitzero"`
	Pragma                                        []string `json:"Pragma,omitempty,omitzero" bson:"Pragma,omitempty,omitzero"`
	StrictTransportSecurity                       []string `json:"Strict-Transport-Security,omitempty,omitzero" bson:"Strict-Transport-Security,omitempty,omitzero"`
	XCache                                        []string `json:"X-Cache,omitempty,omitzero" bson:"X-Cache,omitempty,omitzero"`
	XContentTypeOptions                           []string `json:"X-Content-Type-Options,omitempty,omitzero" bson:"X-Content-Type-Options,omitempty,omitzero"`
	XMsCorrelationRequestID                       []string `json:"X-Ms-Correlation-Request-Id,omitempty,omitzero" bson:"X-Ms-Correlation-Request-Id,omitempty,omitzero"`
	XMsRatelimitRemainingResource                 []string `json:"X-Ms-Ratelimit-Remaining-Resource,omitempty,omitzero" bson:"X-Ms-Ratelimit-Remaining-Resource,omitempty,omitzero"`
	XMsRatelimitRemainingSubscriptionGlobalWrites []string `json:"X-Ms-Ratelimit-Remaining-Subscription-Global-Writes,omitempty,omitzero" bson:"X-Ms-Ratelimit-Remaining-Subscription-Global-Writes,omitempty,omitzero"`
	XMsRatelimitRemainingSubscriptionWrites       []string `json:"X-Ms-Ratelimit-Remaining-Subscription-Writes,omitempty,omitzero" bson:"X-Ms-Ratelimit-Remaining-Subscription-Writes,omitempty,omitzero"`
	XMsRequestID                                  []string `json:"X-Ms-Request-Id,omitempty,omitzero" bson:"X-Ms-Request-Id,omitempty,omitzero"`
	XMsRoutingRequestID                           []string `json:"X-Ms-Routing-Request-Id,omitempty,omitzero" bson:"X-Ms-Routing-Request-Id,omitempty,omitzero"`
	XMsedgeRef                                    []string `json:"X-Msedge-Ref,omitempty,omitzero" bson:"X-Msedge-Ref,omitempty,omitzero"`
}

//
//

type AzureAsyncOpUpdateResponse struct {
	EndTime   time.Time `json:"endTime,omitempty,omitzero" bson:"endTime,omitempty,omitzero"`
	Name      string    `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	StartTime time.Time `json:"startTime,omitempty,omitzero" bson:"startTime,omitempty,omitzero"`
	Status    string    `json:"status,omitempty,omitzero" bson:"status,omitempty,omitzero"`
}

//
//

type VirtualMachine struct {
	ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Location   string `json:"location,omitempty,omitzero" bson:"location,omitempty,omitzero"`
	Name       string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Properties struct {
		AdditionalCapabilities struct {
			HibernationEnabled bool `json:"hibernationEnabled,omitempty,omitzero" bson:"hibernationEnabled,omitempty,omitzero"`
		} `json:"additionalCapabilities,omitempty,omitzero" bson:"additionalCapabilities,omitempty,omitzero"`
		DiagnosticsProfile struct {
			BootDiagnostics struct {
				Enabled bool `json:"enabled,omitempty,omitzero" bson:"enabled,omitempty,omitzero"`
			} `json:"bootDiagnostics,omitempty,omitzero" bson:"bootDiagnostics,omitempty,omitzero"`
		} `json:"diagnosticsProfile,omitempty,omitzero" bson:"diagnosticsProfile,omitempty,omitzero"`
		HardwareProfile struct {
			VmSize string `json:"vmSize,omitempty,omitzero" bson:"vmSize,omitempty,omitzero"`
		} `json:"hardwareProfile,omitempty,omitzero" bson:"hardwareProfile,omitempty,omitzero"`
		NetworkProfile struct {
			NetworkInterfaces []struct {
				ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
				Properties struct {
					DeleteOption string `json:"deleteOption,omitempty,omitzero" bson:"deleteOption,omitempty,omitzero"`
				} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
			} `json:"networkInterfaces,omitempty,omitzero" bson:"networkInterfaces,omitempty,omitzero"`
		} `json:"networkProfile,omitempty,omitzero" bson:"networkProfile,omitempty,omitzero"`
		OSProfile struct {
			AdminUsername            string `json:"adminUsername,omitempty,omitzero" bson:"adminUsername,omitempty,omitzero"`
			AllowExtensionOperations bool   `json:"allowExtensionOperations,omitempty,omitzero" bson:"allowExtensionOperations,omitempty,omitzero"`
			ComputerName             string `json:"computerName,omitempty,omitzero" bson:"computerName,omitempty,omitzero"`
			LinuxConfiguration       struct {
				DisablePasswordAuthentication bool `json:"disablePasswordAuthentication,omitempty,omitzero" bson:"disablePasswordAuthentication,omitempty,omitzero"`
				EnableVmAgentPlatformUpdates  bool `json:"enableVMAgentPlatformUpdates,omitempty,omitzero" bson:"enableVMAgentPlatformUpdates,omitempty,omitzero"`
				PatchSettings                 struct {
					AssessmentMode              string `json:"assessmentMode,omitempty,omitzero" bson:"assessmentMode,omitempty,omitzero"`
					AutomaticByPlatformSettings struct {
						BypassPlatformSafetyChecksOnUserSchedule bool   `json:"bypassPlatformSafetyChecksOnUserSchedule,omitempty,omitzero" bson:"bypassPlatformSafetyChecksOnUserSchedule,omitempty,omitzero"`
						RebootSetting                            string `json:"rebootSetting,omitempty,omitzero" bson:"rebootSetting,omitempty,omitzero"`
					} `json:"automaticByPlatformSettings,omitempty,omitzero" bson:"automaticByPlatformSettings,omitempty,omitzero"`
					PatchMode string `json:"patchMode,omitempty,omitzero" bson:"patchMode,omitempty,omitzero"`
				} `json:"patchSettings,omitempty,omitzero" bson:"patchSettings,omitempty,omitzero"`
				ProvisionVmAgent bool `json:"provisionVMAgent,omitempty,omitzero" bson:"provisionVMAgent,omitempty,omitzero"`
				SSH              struct {
					PublicKeys []struct {
						KeyData string `json:"keyData,omitempty,omitzero" bson:"keyData,omitempty,omitzero"`
						Path    string `json:"path,omitempty,omitzero" bson:"path,omitempty,omitzero"`
					} `json:"publicKeys,omitempty,omitzero" bson:"publicKeys,omitempty,omitzero"`
				} `json:"ssh,omitempty,omitzero" bson:"ssh,omitempty,omitzero"`
			} `json:"linuxConfiguration,omitempty,omitzero" bson:"linuxConfiguration,omitempty,omitzero"`
			RequireGuestProvisionSignal bool  `json:"requireGuestProvisionSignal,omitempty,omitzero" bson:"requireGuestProvisionSignal,omitempty,omitzero"`
			Secrets                     []any `json:"secrets,omitempty,omitzero" bson:"secrets,omitempty,omitzero"`
		} `json:"osProfile,omitempty,omitzero" bson:"osProfile,omitempty,omitzero"`
		ProvisioningState string `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
		SecurityProfile   struct {
			SecurityType string `json:"securityType,omitempty,omitzero" bson:"securityType,omitempty,omitzero"`
			UefiSettings struct {
				SecureBootEnabled bool `json:"secureBootEnabled,omitempty,omitzero" bson:"secureBootEnabled,omitempty,omitzero"`
				VTpmEnabled       bool `json:"vTpmEnabled,omitempty,omitzero" bson:"vTpmEnabled,omitempty,omitzero"`
			} `json:"uefiSettings,omitempty,omitzero" bson:"uefiSettings,omitempty,omitzero"`
		} `json:"securityProfile,omitempty,omitzero" bson:"securityProfile,omitempty,omitzero"`
		StorageProfile struct {
			DataDisks          []any  `json:"dataDisks,omitempty,omitzero" bson:"dataDisks,omitempty,omitzero"`
			DiskControllerType string `json:"diskControllerType,omitempty,omitzero" bson:"diskControllerType,omitempty,omitzero"`
			ImageReference     struct {
				ExactVersion string `json:"exactVersion,omitempty,omitzero" bson:"exactVersion,omitempty,omitzero"`
				Offer        string `json:"offer,omitempty,omitzero" bson:"offer,omitempty,omitzero"`
				Publisher    string `json:"publisher,omitempty,omitzero" bson:"publisher,omitempty,omitzero"`
				Sku          string `json:"sku,omitempty,omitzero" bson:"sku,omitempty,omitzero"`
				Version      string `json:"version,omitempty,omitzero" bson:"version,omitempty,omitzero"`
			} `json:"imageReference,omitempty,omitzero" bson:"imageReference,omitempty,omitzero"`
			OSDisk struct {
				Caching      string  `json:"caching,omitempty,omitzero" bson:"caching,omitempty,omitzero"`
				CreateOption string  `json:"createOption,omitempty,omitzero" bson:"createOption,omitempty,omitzero"`
				DeleteOption string  `json:"deleteOption,omitempty,omitzero" bson:"deleteOption,omitempty,omitzero"`
				DiskSizeGb   float64 `json:"diskSizeGB,omitempty,omitzero" bson:"diskSizeGB,omitempty,omitzero"`
				ManagedDisk  struct {
					StorageAccountType string `json:"storageAccountType,omitempty,omitzero" bson:"storageAccountType,omitempty,omitzero"`
				} `json:"managedDisk,omitempty,omitzero" bson:"managedDisk,omitempty,omitzero"`
				OSType string `json:"osType,omitempty,omitzero" bson:"osType,omitempty,omitzero"`
			} `json:"osDisk,omitempty,omitzero" bson:"osDisk,omitempty,omitzero"`
		} `json:"storageProfile,omitempty,omitzero" bson:"storageProfile,omitempty,omitzero"`
		TimeCreated time.Time `json:"timeCreated,omitempty,omitzero" bson:"timeCreated,omitempty,omitzero"`
		VmID        string    `json:"vmId,omitempty,omitzero" bson:"vmId,omitempty,omitzero"`
	} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
	Type  string   `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
	Zones []string `json:"zones,omitempty,omitzero" bson:"zones,omitempty,omitzero"`
}

//
//

type VirtualMachineInstanceView struct {
	BootDiagnostics struct{} `json:"bootDiagnostics,omitempty,omitzero" bson:"bootDiagnostics,omitempty,omitzero"`
	ComputerName    string   `json:"computerName,omitempty,omitzero" bson:"computerName,omitempty,omitzero"`
	Disks           []struct {
		Name     string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
		Statuses []struct {
			Code          string    `json:"code,omitempty,omitzero" bson:"code,omitempty,omitzero"`
			DisplayStatus string    `json:"displayStatus,omitempty,omitzero" bson:"displayStatus,omitempty,omitzero"`
			Level         string    `json:"level,omitempty,omitzero" bson:"level,omitempty,omitzero"`
			Time          time.Time `json:"time,omitempty,omitzero" bson:"time,omitempty,omitzero"`
		} `json:"statuses,omitempty,omitzero" bson:"statuses,omitempty,omitzero"`
	} `json:"disks,omitempty,omitzero" bson:"disks,omitempty,omitzero"`
	Extensions []struct {
		Name     string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
		Statuses []struct {
			Code          string `json:"code,omitempty,omitzero" bson:"code,omitempty,omitzero"`
			DisplayStatus string `json:"displayStatus,omitempty,omitzero" bson:"displayStatus,omitempty,omitzero"`
			Level         string `json:"level,omitempty,omitzero" bson:"level,omitempty,omitzero"`
			Message       string `json:"message,omitempty,omitzero" bson:"message,omitempty,omitzero"`
		} `json:"statuses,omitempty,omitzero" bson:"statuses,omitempty,omitzero"`
		Type               string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
		TypeHandlerVersion string `json:"typeHandlerVersion,omitempty,omitzero" bson:"typeHandlerVersion,omitempty,omitzero"`
	} `json:"extensions,omitempty,omitzero" bson:"extensions,omitempty,omitzero"`
	HyperVGeneration string `json:"hyperVGeneration,omitempty,omitzero" bson:"hyperVGeneration,omitempty,omitzero"`
	OSName           string `json:"osName,omitempty,omitzero" bson:"osName,omitempty,omitzero"`
	OSVersion        string `json:"osVersion,omitempty,omitzero" bson:"osVersion,omitempty,omitzero"`
	PatchStatus      struct {
		AvailablePatchSummary struct {
			AssessmentActivityID          string  `json:"assessmentActivityId,omitempty,omitzero" bson:"assessmentActivityId,omitempty,omitzero"`
			CriticalAndSecurityPatchCount float64 `json:"criticalAndSecurityPatchCount,omitempty,omitzero" bson:"criticalAndSecurityPatchCount,omitempty,omitzero"`
			Error                         struct {
				Code    string `json:"code,omitempty,omitzero" bson:"code,omitempty,omitzero"`
				Details []struct {
					Code    string `json:"code,omitempty,omitzero" bson:"code,omitempty,omitzero"`
					Message string `json:"message,omitempty,omitzero" bson:"message,omitempty,omitzero"`
				} `json:"details,omitempty,omitzero" bson:"details,omitempty,omitzero"`
				Message string `json:"message,omitempty,omitzero" bson:"message,omitempty,omitzero"`
			} `json:"error,omitempty,omitzero" bson:"error,omitempty,omitzero"`
			LastModifiedTime time.Time `json:"lastModifiedTime,omitempty,omitzero" bson:"lastModifiedTime,omitempty,omitzero"`
			OtherPatchCount  float64   `json:"otherPatchCount,omitempty,omitzero" bson:"otherPatchCount,omitempty,omitzero"`
			RebootPending    bool      `json:"rebootPending,omitempty,omitzero" bson:"rebootPending,omitempty,omitzero"`
			StartTime        time.Time `json:"startTime,omitempty,omitzero" bson:"startTime,omitempty,omitzero"`
			Status           string    `json:"status,omitempty,omitzero" bson:"status,omitempty,omitzero"`
		} `json:"availablePatchSummary,omitempty,omitzero" bson:"availablePatchSummary,omitempty,omitzero"`
		ConfigurationStatuses []struct {
			Code          string    `json:"code,omitempty,omitzero" bson:"code,omitempty,omitzero"`
			DisplayStatus string    `json:"displayStatus,omitempty,omitzero" bson:"displayStatus,omitempty,omitzero"`
			Level         string    `json:"level,omitempty,omitzero" bson:"level,omitempty,omitzero"`
			Time          time.Time `json:"time,omitempty,omitzero" bson:"time,omitempty,omitzero"`
		} `json:"configurationStatuses,omitempty,omitzero" bson:"configurationStatuses,omitempty,omitzero"`
	} `json:"patchStatus,omitempty,omitzero" bson:"patchStatus,omitempty,omitzero"`
	Statuses []struct {
		Code          string    `json:"code,omitempty,omitzero" bson:"code,omitempty,omitzero"`
		DisplayStatus string    `json:"displayStatus,omitempty,omitzero" bson:"displayStatus,omitempty,omitzero"`
		Level         string    `json:"level,omitempty,omitzero" bson:"level,omitempty,omitzero"`
		Time          time.Time `json:"time,omitempty" bson:"time,omitempty"`
	} `json:"statuses,omitempty,omitzero" bson:"statuses,omitempty,omitzero"`
	VmAgent struct {
		ExtensionHandlers []struct {
			Status struct {
				Code          string `json:"code,omitempty,omitzero" bson:"code,omitempty,omitzero"`
				DisplayStatus string `json:"displayStatus,omitempty,omitzero" bson:"displayStatus,omitempty,omitzero"`
				Level         string `json:"level,omitempty,omitzero" bson:"level,omitempty,omitzero"`
				Message       string `json:"message,omitempty,omitzero" bson:"message,omitempty,omitzero"`
			} `json:"status,omitempty,omitzero" bson:"status,omitempty,omitzero"`
			Type               string `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
			TypeHandlerVersion string `json:"typeHandlerVersion,omitempty,omitzero" bson:"typeHandlerVersion,omitempty,omitzero"`
		} `json:"extensionHandlers,omitempty,omitzero" bson:"extensionHandlers,omitempty,omitzero"`
		Statuses []struct {
			Code          string    `json:"code,omitempty,omitzero" bson:"code,omitempty,omitzero"`
			DisplayStatus string    `json:"displayStatus,omitempty,omitzero" bson:"displayStatus,omitempty,omitzero"`
			Level         string    `json:"level,omitempty,omitzero" bson:"level,omitempty,omitzero"`
			Message       string    `json:"message,omitempty,omitzero" bson:"message,omitempty,omitzero"`
			Time          time.Time `json:"time,omitempty,omitzero" bson:"time,omitempty,omitzero"`
		} `json:"statuses,omitempty,omitzero" bson:"statuses,omitempty,omitzero"`
		VmAgentVersion string `json:"vmAgentVersion,omitempty,omitzero" bson:"vmAgentVersion,omitempty,omitzero"`
	} `json:"vmAgent,omitempty,omitzero" bson:"vmAgent,omitempty,omitzero"`
}

//
//

type CheckStorageAccountTlsVersionsResponse struct {
	Count           float64                    `json:"count,omitempty,omitzero" bson:"count,omitempty,omitzero"`
	Data            []StorageAccountTlsVersion `json:"data,omitempty,omitzero" bson:"data,omitempty,omitzero"`
	Facets          []any                      `json:"facets,omitempty,omitzero" bson:"facets,omitempty,omitzero"`
	ResultTruncated string                     `json:"resultTruncated,omitempty,omitzero" bson:"resultTruncated,omitempty,omitzero"`
	TotalRecords    float64                    `json:"totalRecords,omitempty,omitzero" bson:"totalRecords,omitempty,omitzero"`
	SkipToken       string                     `json:"$skipToken,omitempty,omitzero" bson:"$skipToken,omitempty,omitzero"`
}

//
//

type StorageAccountTlsVersion struct {
	ID                string    `json:"id,omitempty,omitzero" bson:"_id,omitempty,omitzero"`
	Name              string    `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	MinimumTlsVersion string    `json:"minimumTlsVersion,omitempty,omitzero" bson:"minimumTlsVersion,omitempty,omitzero"`
	ResourceGroup     string    `json:"resourceGroup,omitempty,omitzero" bson:"resourceGroup,omitempty,omitzero"`
	SubscriptionID    string    `json:"subscriptionId,omitempty,omitzero" bson:"subscriptionId,omitempty,omitzero"`
	SubscriptionName  string    `json:"subscriptionName,omitempty,omitzero" bson:"subscriptionName,omitempty,omitzero"`
	TenantID          string    `json:"tenantId,omitempty,omitzero" bson:"tenantId,omitempty,omitzero"`
	TenantName        string    `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	LastDBSync        time.Time `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
	LastAzureSync     time.Time `json:"lastAzureSync,omitempty,omitzero" bson:"lastAzureSync,omitempty,omitzero"`
}

//
//
