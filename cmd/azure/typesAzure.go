/*
Copyright © 2024 Evan Colwell ercolwell@gmail.com
*/
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

type EntraListApplicationsResponse struct {
	OdataContext string             `json:"@odata.context,omitempty" bson:"@odata.context,omitempty"`
	NextLink     string             `json:"@odata.nextLink,omitempty" bson:"@odata.nextLink,omitempty"`
	Value        []EntraApplication `json:"value,omitempty" bson:"value,omitempty"`
}

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
	AppID           string `json:"appId,omitempty" bson:"_id,omitempty"`
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
	Description               any       `json:"description,omitempty" bson:"description,omitempty"`
	DisabledByMicrosoftStatus any       `json:"disabledByMicrosoftStatus,omitempty" bson:"disabledByMicrosoftStatus,omitempty"`
	DisplayName               string    `json:"displayName,omitempty" bson:"displayName,omitempty"`
	GroupMembershipClaims     *string   `json:"groupMembershipClaims,omitempty" bson:"groupMembershipClaims,omitempty"`
	ID                        string    `json:"id,omitempty" bson:"id,omitempty"`
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
		CustomKeyIdentifier *string   `json:"customKeyIdentifier,omitempty" bson:"customKeyIdentifier,omitempty"`
		DisplayName         *string   `json:"displayName,omitempty" bson:"displayName,omitempty"`
		EndDateTime         time.Time `json:"endDateTime,omitempty" bson:"endDateTime,omitempty"`
		Hint                *string   `json:"hint,omitempty" bson:"hint,omitempty"`
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
	Tags                 *[]string `json:"tags,omitempty" bson:"tags,omitempty"`
	TenantName           string    `json:"tenantName,omitempty" bson:"tenantName,omitempty"`
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
	Windows any `json:"windows,omitempty" bson:"windows,omitempty"`
}
