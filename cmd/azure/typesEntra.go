package azure

import (
	"time"

	"encoding/json/jsontext"
)

type EntraRoleDefinition struct {
	Description     string `json:"description,omitempty" bson:"description,omitempty"`
	DisplayName     string `json:"displayName,omitempty" bson:"displayName,omitempty"`
	ID              string `json:"id,omitempty" bson:"id,omitempty"`
	IsBuiltIn       bool   `json:"isBuiltIn,omitempty" bson:"isBuiltIn,omitempty"`
	IsEnabled       bool   `json:"isEnabled,omitempty" bson:"isEnabled,omitempty"`
	RolePermissions []struct {
		AllowedResourceActions []string `json:"allowedResourceActions,omitempty" bson:"allowedResourceActions,omitempty"`
		Condition              *string  `json:"condition,omitempty" bson:"condition,omitempty"`
	} `json:"rolePermissions,omitempty" bson:"rolePermissions,omitempty"`
}

//
//

type ListEntraRoleDefinitionsResponse struct {
	Odata_Context string                `json:"@odata.context,omitempty" bson:"@odata.context,omitempty"`
	Value         []EntraRoleDefinition `json:"value,omitempty" bson:"value,omitempty"`
}

//
//

type ListRoleEligibilityScheduleInstancesResponse struct {
	Value []RoleEligibilityScheduleInstance `json:"value,omitempty" bson:"value,omitempty"`
}

//
//

type RoleEligibilityScheduleInstance struct {
	ID         string `json:"id,omitempty" bson:"id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		CreatedOn          time.Time `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
		EndDateTime        time.Time `json:"endDateTime,omitempty" bson:"endDateTime,omitempty"`
		ExpandedProperties struct {
			Principal struct {
				DisplayName string `json:"displayName,omitempty" bson:"displayName,omitempty"`
				ID          string `json:"id,omitempty" bson:"id,omitempty"`
				Type        string `json:"type,omitempty" bson:"type,omitempty"`
			} `json:"principal,omitempty" bson:"principal,omitempty"`
			RoleDefinition struct {
				DisplayName string `json:"displayName,omitempty" bson:"displayName,omitempty"`
				ID          string `json:"id,omitempty" bson:"id,omitempty"`
				Type        string `json:"type,omitempty" bson:"type,omitempty"`
			} `json:"roleDefinition,omitempty" bson:"roleDefinition,omitempty"`
			Scope struct {
				DisplayName string `json:"displayName,omitempty" bson:"displayName,omitempty"`
				ID          string `json:"id,omitempty" bson:"id,omitempty"`
				Type        string `json:"type,omitempty" bson:"type,omitempty"`
			} `json:"scope,omitempty" bson:"scope,omitempty"`
		} `json:"expandedProperties,omitempty" bson:"expandedProperties,omitempty"`
		MemberType                string    `json:"memberType,omitempty" bson:"memberType,omitempty"`
		PrincipalID               string    `json:"principalId,omitempty" bson:"principalId,omitempty"`
		PrincipalType             string    `json:"principalType,omitempty" bson:"principalType,omitempty"`
		RoleDefinitionID          string    `json:"roleDefinitionId,omitempty" bson:"roleDefinitionId,omitempty"`
		RoleEligibilityScheduleID string    `json:"roleEligibilityScheduleId,omitempty" bson:"roleEligibilityScheduleId,omitempty"`
		Scope                     string    `json:"scope,omitempty" bson:"scope,omitempty"`
		StartDateTime             time.Time `json:"startDateTime,omitempty" bson:"startDateTime,omitempty"`
		Status                    string    `json:"status,omitempty" bson:"status,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type          string    `json:"type,omitempty" bson:"type,omitempty"`
	ScopeType     string    `json:"scopeType,omitempty" bson:"scopeType,omitempty"`
	LastAzureSync time.Time `json:"lastAzureSync,omitempty" bson:"lastAzureSync,omitempty" fake:"-"`
	LastDBSync    time.Time `json:"lastDatabaseSync,omitempty" bson:"lastDatabaseSync,omitempty" fake:"-"`
	TenantName    string    `json:"tenantName,omitempty" bson:"tenantName,omitempty" fake:"-"`
}

//
//

type ListRoleAssignmentScheduleInstancesResponse struct {
	Value []RoleAssignmentScheduleInstance `json:"value,omitempty" bson:"value,omitempty"`
}

//
//

type RoleAssignmentScheduleInstance struct {
	ID         string `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string `json:"name,omitempty" bson:"name,omitempty"`
	Properties struct {
		AssignmentType     string    `json:"assignmentType,omitempty" bson:"assignmentType,omitempty"`
		CreatedOn          time.Time `json:"createdOn,omitempty" bson:"createdOn,omitempty"`
		EndDateTime        time.Time `json:"endDateTime,omitempty" bson:"endDateTime,omitempty"`
		ExpandedProperties struct {
			Principal struct {
				DisplayName string `json:"displayName,omitempty" bson:"displayName,omitempty"`
				ID          string `json:"id,omitempty" bson:"id,omitempty"`
				Type        string `json:"type,omitempty" bson:"type,omitempty"`
			} `json:"principal,omitempty" bson:"principal,omitempty"`
			RoleDefinition struct {
				DisplayName string `json:"displayName,omitempty" bson:"displayName,omitempty"`
				ID          string `json:"id,omitempty" bson:"id,omitempty"`
				Type        string `json:"type,omitempty" bson:"type,omitempty"`
			} `json:"roleDefinition,omitempty" bson:"roleDefinition,omitempty"`
			Scope struct {
				DisplayName string `json:"displayName,omitempty" bson:"displayName,omitempty"`
				ID          string `json:"id,omitempty" bson:"id,omitempty"`
				Type        string `json:"type,omitempty" bson:"type,omitempty"`
			} `json:"scope,omitempty" bson:"scope,omitempty"`
		} `json:"expandedProperties,omitempty" bson:"expandedProperties,omitempty"`
		LinkedRoleEligibilityScheduleID         string    `json:"linkedRoleEligibilityScheduleId,omitempty" bson:"linkedRoleEligibilityScheduleId,omitempty"`
		LinkedRoleEligibilityScheduleInstanceID string    `json:"linkedRoleEligibilityScheduleInstanceId,omitempty" bson:"linkedRoleEligibilityScheduleInstanceId,omitempty"`
		MemberType                              string    `json:"memberType,omitempty" bson:"memberType,omitempty"`
		OriginRoleAssignmentID                  string    `json:"originRoleAssignmentId,omitempty" bson:"originRoleAssignmentId,omitempty"`
		PrincipalID                             string    `json:"principalId,omitempty" bson:"principalId,omitempty"`
		PrincipalType                           string    `json:"principalType,omitempty" bson:"principalType,omitempty"`
		RoleAssignmentScheduleID                string    `json:"roleAssignmentScheduleId,omitempty" bson:"roleAssignmentScheduleId,omitempty"`
		RoleDefinitionID                        string    `json:"roleDefinitionId,omitempty" bson:"roleDefinitionId,omitempty"`
		Scope                                   string    `json:"scope,omitempty" bson:"scope,omitempty"`
		StartDateTime                           time.Time `json:"startDateTime,omitempty" bson:"startDateTime,omitempty"`
		Status                                  string    `json:"status,omitempty" bson:"status,omitempty"`
	} `json:"properties,omitempty" bson:"properties,omitempty"`
	Type          string    `json:"type,omitempty" bson:"type,omitempty"`
	ScopeType     string    `json:"scopeType,omitempty" bson:"scopeType,omitempty"`
	LastAzureSync time.Time `json:"lastAzureSync,omitempty" bson:"lastAzureSync,omitempty" fake:"-"`
	LastDBSync    time.Time `json:"lastDatabaseSync,omitempty" bson:"lastDatabaseSync,omitempty" fake:"-"`
	TenantName    string    `json:"tenantName,omitempty" bson:"tenantName,omitempty" fake:"-"`
}

//
//

type ListRoleAssignmentSchedulesResponse struct {
	Odata_Context  string                   `json:"@odata.context,omitempty" bson:"@odata.context,omitempty"`
	Odata_NextLink string                   `json:"@odata.nextLink,omitempty" bson:"@odata.nextLink,omitempty"`
	Value          []RoleAssignmentSchedule `json:"value,omitempty" bson:"value,omitempty"`
}

//
//

type RoleAssignmentScheduleProcessed struct {
	ActivatedUsing any `json:"activatedUsing,omitempty" bson:"activatedUsing,omitempty"`
	Principal      struct {
		Odata_Type                string   `json:"@odata.type,omitempty" bson:"@odata.type,omitempty"`
		AccountEnabled            bool     `json:"accountEnabled,omitempty" bson:"accountEnabled,omitempty"`
		AppDescription            any      `json:"appDescription,omitempty" bson:"appDescription,omitempty"`
		AppDisplayName            string   `json:"appDisplayName,omitempty" bson:"appDisplayName,omitempty"`
		AppID                     string   `json:"appId,omitempty" bson:"appId,omitempty"`
		Description               *string  `json:"description,omitempty" bson:"description,omitempty"`
		DisplayName               string   `json:"displayName,omitempty" bson:"displayName,omitempty"`
		ID                        string   `json:"id,omitempty" bson:"id,omitempty"`
		ManagedIdentityResourceID any      `json:"managedIdentityResourceId,omitempty" bson:"managedIdentityResourceId,omitempty"`
		ServicePrincipalNames     []string `json:"servicePrincipalNames,omitempty" bson:"servicePrincipalNames,omitempty"`
		ServicePrincipalType      string   `json:"servicePrincipalType,omitempty" bson:"servicePrincipalType,omitempty"`
		Tags                      []string `json:"tags" bson:"tags"`
		UserPrincipalName         string   `json:"userPrincipalName,omitempty" bson:"userPrincipalName,omitempty"`
		UserType                  string   `json:"userType,omitempty" bson:"userType,omitempty"`
	} `json:"principal,omitempty" bson:"principal,omitempty"`
	PrincipalID    string `json:"principalId,omitempty" bson:"principalId,omitempty"`
	RoleDefinition struct {
		Description     string `json:"description,omitempty" bson:"description,omitempty"`
		DisplayName     string `json:"displayName,omitempty" bson:"displayName,omitempty"`
		ID              string `json:"id,omitempty" bson:"id,omitempty"`
		IsBuiltIn       bool   `json:"isBuiltIn,omitempty" bson:"isBuiltIn,omitempty"`
		IsEnabled       bool   `json:"isEnabled,omitempty" bson:"isEnabled,omitempty"`
		ResourceScopes  []any  `json:"resourceScopes,omitempty" bson:"resourceScopes,omitempty"`
		RolePermissions []any  `json:"rolePermissions,omitempty" bson:"rolePermissions,omitempty"`
	} `json:"roleDefinition,omitempty" bson:"roleDefinition,omitempty"`
	RoleDefinitionID string `json:"roleDefinitionId,omitempty" bson:"roleDefinitionId,omitempty"`
}

//
//

type RoleAssignmentSchedule struct {
	ActivatedUsing   any           `json:"activatedUsing,omitempty" bson:"activatedUsing,omitempty"`
	AppScopeID       any           `json:"appScopeId,omitempty" bson:"appScopeId,omitempty"`
	AssignmentType   string        `json:"assignmentType,omitempty" bson:"assignmentType,omitempty"`
	CreatedDateTime  *time.Time    `json:"createdDateTime,omitempty" bson:"createdDateTime,omitempty"`
	CreatedUsing     *string       `json:"createdUsing,omitempty" bson:"createdUsing,omitempty"`
	DirectoryScopeID string        `json:"directoryScopeId,omitempty" bson:"directoryScopeId,omitempty"`
	ID               string        `json:"id,omitempty" bson:"id,omitempty"`
	MemberType       string        `json:"memberType,omitempty" bson:"memberType,omitempty"`
	ModifiedDateTime any           `json:"modifiedDateTime,omitempty" bson:"modifiedDateTime,omitempty"`
	Principal        RolePrinciple `json:"principal,omitempty" bson:"principal,omitempty"`
	PrincipalID      string        `json:"principalId,omitempty" bson:"principalId,omitempty"`
	RoleDefinition   struct {
		Description     string `json:"description,omitempty" bson:"description,omitempty"`
		DisplayName     string `json:"displayName,omitempty" bson:"displayName,omitempty"`
		ID              string `json:"id,omitempty" bson:"id,omitempty"`
		IsBuiltIn       bool   `json:"isBuiltIn,omitempty" bson:"isBuiltIn,omitempty"`
		IsEnabled       bool   `json:"isEnabled,omitempty" bson:"isEnabled,omitempty"`
		ResourceScopes  []any  `json:"resourceScopes,omitempty" bson:"resourceScopes,omitempty"`
		RolePermissions []any  `json:"rolePermissions,omitempty" bson:"rolePermissions,omitempty"`
		TemplateID      string `json:"templateId,omitempty" bson:"templateId,omitempty"`
		Version         any    `json:"version,omitempty" bson:"version,omitempty"`
	} `json:"roleDefinition,omitempty" bson:"roleDefinition,omitempty"`
	RoleDefinitionID string `json:"roleDefinitionId,omitempty" bson:"roleDefinitionId,omitempty"`
	ScheduleInfo     struct {
		Expiration struct {
			Duration    any        `json:"duration,omitempty" bson:"duration,omitempty"`
			EndDateTime *time.Time `json:"endDateTime,omitempty" bson:"endDateTime,omitempty"`
			Type        string     `json:"type,omitempty" bson:"type,omitempty"`
		} `json:"expiration,omitempty" bson:"expiration,omitempty"`
		Recurrence    any       `json:"recurrence,omitempty" bson:"recurrence,omitempty"`
		StartDateTime time.Time `json:"startDateTime,omitempty" bson:"startDateTime,omitempty"`
	} `json:"scheduleInfo,omitempty" bson:"scheduleInfo,omitempty"`
	Status string `json:"status,omitempty" bson:"status,omitempty"`
}

//
//

type RolePrinciple struct {
	Odata_Type     string `json:"@odata.type,omitempty" bson:"@odata.type,omitempty"`
	AccountEnabled bool   `json:"accountEnabled,omitempty" bson:"accountEnabled,omitempty"`
	// AddIns                 []any  `json:"addIns,omitempty" bson:"addIns,omitempty"`
	// AgeGroup               any    `json:"ageGroup,omitempty" bson:"ageGroup,omitempty"`
	// AlternativeNames       []any  `json:"alternativeNames,omitempty" bson:"alternativeNames,omitempty"`
	// AlternativeSecurityIds []any  `json:"alternativeSecurityIds,omitempty" bson:"alternativeSecurityIds,omitempty"`
	// API                    *struct {
	// 	ResourceSpecificApplicationPermissions []any `json:"resourceSpecificApplicationPermissions,omitempty" bson:"resourceSpecificApplicationPermissions,omitempty"`
	// } `json:"api,omitempty" bson:"api,omitempty"`
	// AppData                   any    `json:"appData,omitempty" bson:"appData,omitempty"`
	AppDescription any    `json:"appDescription,omitempty" bson:"appDescription,omitempty"`
	AppDisplayName string `json:"appDisplayName,omitempty" bson:"appDisplayName,omitempty"`
	AppID          string `json:"appId,omitempty" bson:"appId,omitempty"`
	// AppMetadata               any    `json:"appMetadata,omitempty" bson:"appMetadata,omitempty"`
	// AppOwnerOrganizationID    string `json:"appOwnerOrganizationId,omitempty" bson:"appOwnerOrganizationId,omitempty"`
	// AppRoleAssignmentRequired bool   `json:"appRoleAssignmentRequired,omitempty" bson:"appRoleAssignmentRequired,omitempty"`
	// AppRoles                  []any  `json:"appRoles,omitempty" bson:"appRoles,omitempty"`
	// ApplicationTemplateID     any    `json:"applicationTemplateId,omitempty" bson:"applicationTemplateId,omitempty"`
	// AssignedLicenses          []struct {
	// 	DisabledPlans []string `json:"disabledPlans,omitempty" bson:"disabledPlans,omitempty"`
	// 	SkuID         string   `json:"skuId,omitempty" bson:"skuId,omitempty"`
	// } `json:"assignedLicenses,omitempty" bson:"assignedLicenses,omitempty"`
	// AssignedPlans []struct {
	// 	AssignedDateTime time.Time `json:"assignedDateTime,omitempty" bson:"assignedDateTime,omitempty"`
	// 	CapabilityStatus string    `json:"capabilityStatus,omitempty" bson:"capabilityStatus,omitempty"`
	// 	Service          string    `json:"service,omitempty" bson:"service,omitempty"`
	// 	ServicePlanID    string    `json:"servicePlanId,omitempty" bson:"servicePlanId,omitempty"`
	// } `json:"assignedPlans,omitempty" bson:"assignedPlans,omitempty"`
	// AuthorizationInfo *struct {
	// 	CertificateUserIds []any `json:"certificateUserIds,omitempty" bson:"certificateUserIds,omitempty"`
	// } `json:"authorizationInfo,omitempty" bson:"authorizationInfo,omitempty"`
	// BusinessPhones                 []any `json:"businessPhones,omitempty" bson:"businessPhones,omitempty"`
	// Certification                  any   `json:"certification,omitempty" bson:"certification,omitempty"`
	// City                           any   `json:"city,omitempty" bson:"city,omitempty"`
	// Classification                 any   `json:"classification,omitempty" bson:"classification,omitempty"`
	// CloudRealtimeCommunicationInfo *struct {
	// 	CloudMsRtcOwnerUrn          any      `json:"cloudMSRtcOwnerUrn,omitempty" bson:"cloudMSRtcOwnerUrn,omitempty"`
	// 	CloudMsRtcPolicyAssignments []string `json:"cloudMSRtcPolicyAssignments,omitempty" bson:"cloudMSRtcPolicyAssignments,omitempty"`
	// 	CloudMsRtcPool              *string  `json:"cloudMSRtcPool,omitempty" bson:"cloudMSRtcPool,omitempty"`
	// 	CloudMsRtcServiceAttributes *struct {
	// 		ApplicationOption   any     `json:"applicationOption,omitempty" bson:"applicationOption,omitempty"`
	// 		DeploymentLocator   string  `json:"deploymentLocator,omitempty" bson:"deploymentLocator,omitempty"`
	// 		HideFromAddressList any     `json:"hideFromAddressList,omitempty" bson:"hideFromAddressList,omitempty"`
	// 		OptionFlag          float64 `json:"optionFlag,omitempty" bson:"optionFlag,omitempty"`
	// 	} `json:"cloudMSRtcServiceAttributes,omitempty" bson:"cloudMSRtcServiceAttributes,omitempty"`
	// 	CloudSipLine         any     `json:"cloudSipLine,omitempty" bson:"cloudSipLine,omitempty"`
	// 	CloudSipProxyAddress *string `json:"cloudSipProxyAddress,omitempty" bson:"cloudSipProxyAddress,omitempty"`
	// 	IsSipEnabled         *bool   `json:"isSipEnabled,omitempty" bson:"isSipEnabled,omitempty"`
	// } `json:"cloudRealtimeCommunicationInfo,omitempty" bson:"cloudRealtimeCommunicationInfo,omitempty"`
	// CompanyName               any       `json:"companyName,omitempty" bson:"companyName,omitempty"`
	// ConsentProvidedForMinor   any       `json:"consentProvidedForMinor,omitempty" bson:"consentProvidedForMinor,omitempty"`
	// Country                   any       `json:"country,omitempty" bson:"country,omitempty"`
	// CreatedByAppID            string    `json:"createdByAppId,omitempty" bson:"createdByAppId,omitempty"`
	// CreatedDateTime           time.Time `json:"createdDateTime,omitempty" bson:"createdDateTime,omitempty"`
	// CreationType              any       `json:"creationType,omitempty" bson:"creationType,omitempty"`
	// DeletedDateTime           any       `json:"deletedDateTime,omitempty" bson:"deletedDateTime,omitempty"`
	// Department                *string   `json:"department,omitempty" bson:"department,omitempty"`
	Description *string `json:"description,omitempty" bson:"description,omitempty"`
	// DeviceKeys                []any     `json:"deviceKeys,omitempty" bson:"deviceKeys,omitempty"`
	// DeviceManagementAppType   any       `json:"deviceManagementAppType,omitempty" bson:"deviceManagementAppType,omitempty"`
	// DisabledByMicrosoftStatus any       `json:"disabledByMicrosoftStatus,omitempty" bson:"disabledByMicrosoftStatus,omitempty"`
	DisplayName string `json:"displayName,omitempty" bson:"displayName,omitempty"`
	// EmployeeHireDate          any       `json:"employeeHireDate,omitempty" bson:"employeeHireDate,omitempty"`
	// EmployeeID                any       `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	// EmployeeLeaveDateTime     any       `json:"employeeLeaveDateTime,omitempty" bson:"employeeLeaveDateTime,omitempty"`
	// EmployeeOrgData           any       `json:"employeeOrgData,omitempty" bson:"employeeOrgData,omitempty"`
	// EmployeeType              any       `json:"employeeType,omitempty" bson:"employeeType,omitempty"`
	// ErrorURL                  any       `json:"errorUrl,omitempty" bson:"errorUrl,omitempty"`
	// ExpirationDateTime        any       `json:"expirationDateTime,omitempty" bson:"expirationDateTime,omitempty"`
	// ExternalUserConvertedOn   any       `json:"externalUserConvertedOn,omitempty" bson:"externalUserConvertedOn,omitempty"`
	// ExternalUserInformation   *struct {
	// 	AcceptedAsMail   any       `json:"acceptedAsMail,omitempty" bson:"acceptedAsMail,omitempty"`
	// 	AcceptedDateTime any       `json:"acceptedDateTime,omitempty" bson:"acceptedDateTime,omitempty"`
	// 	InviteReplyUrls  []any     `json:"inviteReplyUrls,omitempty" bson:"inviteReplyUrls,omitempty"`
	// 	InviteResources  []any     `json:"inviteResources,omitempty" bson:"inviteResources,omitempty"`
	// 	InvitedAsMail    any       `json:"invitedAsMail,omitempty" bson:"invitedAsMail,omitempty"`
	// 	InvitedDateTime  time.Time `json:"invitedDateTime,omitempty" bson:"invitedDateTime,omitempty"`
	// 	SignInNames      []string  `json:"signInNames,omitempty" bson:"signInNames,omitempty"`
	// } `json:"externalUserInformation,omitempty" bson:"externalUserInformation,omitempty"`
	// ExternalUserState               any     `json:"externalUserState,omitempty" bson:"externalUserState,omitempty"`
	// ExternalUserStateChangeDateTime any     `json:"externalUserStateChangeDateTime,omitempty" bson:"externalUserStateChangeDateTime,omitempty"`
	// FaxNumber                       any     `json:"faxNumber,omitempty" bson:"faxNumber,omitempty"`
	// GivenName                       *string `json:"givenName,omitempty" bson:"givenName,omitempty"`
	// GroupTypes                      []any   `json:"groupTypes,omitempty" bson:"groupTypes,omitempty"`
	// Homepage                        *string `json:"homepage,omitempty" bson:"homepage,omitempty"`
	ID string `json:"id,omitempty" bson:"id,omitempty"`
	// Identities                      []struct {
	// 	Issuer           string `json:"issuer,omitempty" bson:"issuer,omitempty"`
	// 	IssuerAssignedID string `json:"issuerAssignedId,omitempty" bson:"issuerAssignedId,omitempty"`
	// 	SignInType       string `json:"signInType,omitempty" bson:"signInType,omitempty"`
	// } `json:"identities,omitempty" bson:"identities,omitempty"`
	// ImAddresses []string `json:"imAddresses,omitempty" bson:"imAddresses,omitempty"`
	// Info        *struct {
	// 	LogoURL             any `json:"logoUrl,omitempty" bson:"logoUrl,omitempty"`
	// 	MarketingURL        any `json:"marketingUrl,omitempty" bson:"marketingUrl,omitempty"`
	// 	PrivacyStatementURL any `json:"privacyStatementUrl,omitempty" bson:"privacyStatementUrl,omitempty"`
	// 	SupportURL          any `json:"supportUrl,omitempty" bson:"supportUrl,omitempty"`
	// 	TermsOfServiceURL   any `json:"termsOfServiceUrl,omitempty" bson:"termsOfServiceUrl,omitempty"`
	// } `json:"info,omitempty" bson:"info,omitempty"`
	// InfoCatalogs                  []any   `json:"infoCatalogs,omitempty" bson:"infoCatalogs,omitempty"`
	// IsAssignableToRole            bool    `json:"isAssignableToRole,omitempty" bson:"isAssignableToRole,omitempty"`
	// IsAuthorizationServiceEnabled bool    `json:"isAuthorizationServiceEnabled,omitempty" bson:"isAuthorizationServiceEnabled,omitempty"`
	// IsLicenseReconciliationNeeded bool    `json:"isLicenseReconciliationNeeded,omitempty" bson:"isLicenseReconciliationNeeded,omitempty"`
	// IsManagementRestricted        any     `json:"isManagementRestricted,omitempty" bson:"isManagementRestricted,omitempty"`
	IsResourceAccount any `json:"isResourceAccount,omitempty" bson:"isResourceAccount,omitempty"`
	// JobTitle                      any     `json:"jobTitle,omitempty" bson:"jobTitle,omitempty"`
	// KeyCredentials                []any   `json:"keyCredentials,omitempty" bson:"keyCredentials,omitempty"`
	// LegalAgeGroupClassification   any     `json:"legalAgeGroupClassification,omitempty" bson:"legalAgeGroupClassification,omitempty"`
	// LoginURL                      any     `json:"loginUrl,omitempty" bson:"loginUrl,omitempty"`
	// LogoutURL                     any     `json:"logoutUrl,omitempty" bson:"logoutUrl,omitempty"`
	// Mail                          *string `json:"mail,omitempty" bson:"mail,omitempty"`
	// MailEnabled                   bool    `json:"mailEnabled,omitempty" bson:"mailEnabled,omitempty"`
	// MailNickname                  string  `json:"mailNickname,omitempty" bson:"mailNickname,omitempty"`
	ManagedIdentityResourceID any `json:"managedIdentityResourceId,omitempty" bson:"managedIdentityResourceId,omitempty"`
	// MembershipRule                any     `json:"membershipRule,omitempty" bson:"membershipRule,omitempty"`
	// MembershipRuleProcessingState any     `json:"membershipRuleProcessingState,omitempty" bson:"membershipRuleProcessingState,omitempty"`
	// MicrosoftPolicyGroup          any     `json:"microsoftPolicyGroup,omitempty" bson:"microsoftPolicyGroup,omitempty"`
	// MobilePhone                   any     `json:"mobilePhone,omitempty" bson:"mobilePhone,omitempty"`
	// NetID                         string  `json:"netId,omitempty" bson:"netId,omitempty"`
	// Notes                         any     `json:"notes,omitempty" bson:"notes,omitempty"`
	// NotificationEmailAddresses    []any   `json:"notificationEmailAddresses,omitempty" bson:"notificationEmailAddresses,omitempty"`
	// OfficeLocation                any     `json:"officeLocation,omitempty" bson:"officeLocation,omitempty"`
	// OnPremisesDistinguishedName   *string `json:"onPremisesDistinguishedName,omitempty" bson:"onPremisesDistinguishedName,omitempty"`
	// OnPremisesDomainName          *string `json:"onPremisesDomainName,omitempty" bson:"onPremisesDomainName,omitempty"`
	// OnPremisesExtensionAttributes *struct {
	// 	ExtensionAttribute1  any `json:"extensionAttribute1,omitempty" bson:"extensionAttribute1,omitempty"`
	// 	ExtensionAttribute10 any `json:"extensionAttribute10,omitempty" bson:"extensionAttribute10,omitempty"`
	// 	ExtensionAttribute11 any `json:"extensionAttribute11,omitempty" bson:"extensionAttribute11,omitempty"`
	// 	ExtensionAttribute12 any `json:"extensionAttribute12,omitempty" bson:"extensionAttribute12,omitempty"`
	// 	ExtensionAttribute13 any `json:"extensionAttribute13,omitempty" bson:"extensionAttribute13,omitempty"`
	// 	ExtensionAttribute14 any `json:"extensionAttribute14,omitempty" bson:"extensionAttribute14,omitempty"`
	// 	ExtensionAttribute15 any `json:"extensionAttribute15,omitempty" bson:"extensionAttribute15,omitempty"`
	// 	ExtensionAttribute2  any `json:"extensionAttribute2,omitempty" bson:"extensionAttribute2,omitempty"`
	// 	ExtensionAttribute3  any `json:"extensionAttribute3,omitempty" bson:"extensionAttribute3,omitempty"`
	// 	ExtensionAttribute4  any `json:"extensionAttribute4,omitempty" bson:"extensionAttribute4,omitempty"`
	// 	ExtensionAttribute5  any `json:"extensionAttribute5,omitempty" bson:"extensionAttribute5,omitempty"`
	// 	ExtensionAttribute6  any `json:"extensionAttribute6,omitempty" bson:"extensionAttribute6,omitempty"`
	// 	ExtensionAttribute7  any `json:"extensionAttribute7,omitempty" bson:"extensionAttribute7,omitempty"`
	// 	ExtensionAttribute8  any `json:"extensionAttribute8,omitempty" bson:"extensionAttribute8,omitempty"`
	// 	ExtensionAttribute9  any `json:"extensionAttribute9,omitempty" bson:"extensionAttribute9,omitempty"`
	// } `json:"onPremisesExtensionAttributes,omitempty" bson:"onPremisesExtensionAttributes,omitempty"`
	// OnPremisesImmutableID        *string    `json:"onPremisesImmutableId,omitempty" bson:"onPremisesImmutableId,omitempty"`
	// OnPremisesLastSyncDateTime   *time.Time `json:"onPremisesLastSyncDateTime,omitempty" bson:"onPremisesLastSyncDateTime,omitempty"`
	// OnPremisesNetBiosName        any        `json:"onPremisesNetBiosName,omitempty" bson:"onPremisesNetBiosName,omitempty"`
	// OnPremisesObjectIdentifier   *string    `json:"onPremisesObjectIdentifier,omitempty" bson:"onPremisesObjectIdentifier,omitempty"`
	// OnPremisesProvisioningErrors []any      `json:"onPremisesProvisioningErrors,omitempty" bson:"onPremisesProvisioningErrors,omitempty"`
	// OnPremisesSamAccountName     *string    `json:"onPremisesSamAccountName,omitempty" bson:"onPremisesSamAccountName,omitempty"`
	// OnPremisesSecurityIdentifier *string    `json:"onPremisesSecurityIdentifier,omitempty" bson:"onPremisesSecurityIdentifier,omitempty"`
	// OnPremisesSipInfo            *struct {
	// 	IsSipEnabled          bool `json:"isSipEnabled,omitempty" bson:"isSipEnabled,omitempty"`
	// 	SipDeploymentLocation any  `json:"sipDeploymentLocation,omitempty" bson:"sipDeploymentLocation,omitempty"`
	// 	SipPrimaryAddress     any  `json:"sipPrimaryAddress,omitempty" bson:"sipPrimaryAddress,omitempty"`
	// } `json:"onPremisesSipInfo,omitempty" bson:"onPremisesSipInfo,omitempty"`
	// OnPremisesSyncEnabled       *bool   `json:"onPremisesSyncEnabled,omitempty" bson:"onPremisesSyncEnabled,omitempty"`
	// OnPremisesUserPrincipalName *string `json:"onPremisesUserPrincipalName,omitempty" bson:"onPremisesUserPrincipalName,omitempty"`
	// OrganizationID              string  `json:"organizationId,omitempty" bson:"organizationId,omitempty"`
	// OtherMails                  []any   `json:"otherMails,omitempty" bson:"otherMails,omitempty"`
	// PasswordCredentials         []any   `json:"passwordCredentials,omitempty" bson:"passwordCredentials,omitempty"`
	// PasswordPolicies            *string `json:"passwordPolicies,omitempty" bson:"passwordPolicies,omitempty"`
	// PasswordProfile             *struct {
	// 	ForceChangePasswordNextSignIn        bool `json:"forceChangePasswordNextSignIn,omitempty" bson:"forceChangePasswordNextSignIn,omitempty"`
	// 	ForceChangePasswordNextSignInWithMfa bool `json:"forceChangePasswordNextSignInWithMfa,omitempty" bson:"forceChangePasswordNextSignInWithMfa,omitempty"`
	// 	Password                             any  `json:"password,omitempty" bson:"password,omitempty"`
	// } `json:"passwordProfile,omitempty" bson:"passwordProfile,omitempty"`
	// PortalSetting                       any `json:"portalSetting,omitempty" bson:"portalSetting,omitempty"`
	// PostalCode                          any `json:"postalCode,omitempty" bson:"postalCode,omitempty"`
	// PreferredDataLocation               any `json:"preferredDataLocation,omitempty" bson:"preferredDataLocation,omitempty"`
	// PreferredLanguage                   any `json:"preferredLanguage,omitempty" bson:"preferredLanguage,omitempty"`
	// PreferredSingleSignOnMode           any `json:"preferredSingleSignOnMode,omitempty" bson:"preferredSingleSignOnMode,omitempty"`
	// PreferredTokenSigningKeyEndDateTime any `json:"preferredTokenSigningKeyEndDateTime,omitempty" bson:"preferredTokenSigningKeyEndDateTime,omitempty"`
	// PreferredTokenSigningKeyThumbprint  any `json:"preferredTokenSigningKeyThumbprint,omitempty" bson:"preferredTokenSigningKeyThumbprint,omitempty"`
	// ProvisionedPlans                    []struct {
	// 	CapabilityStatus   string `json:"capabilityStatus,omitempty" bson:"capabilityStatus,omitempty"`
	// 	ProvisioningStatus string `json:"provisioningStatus,omitempty" bson:"provisioningStatus,omitempty"`
	// 	Service            string `json:"service,omitempty" bson:"service,omitempty"`
	// } `json:"provisionedPlans,omitempty" bson:"provisionedPlans,omitempty"`
	// ProxyAddresses            []string `json:"proxyAddresses,omitempty" bson:"proxyAddresses,omitempty"`
	// PublishedPermissionScopes []struct {
	// 	AdminConsentDescription string `json:"adminConsentDescription,omitempty" bson:"adminConsentDescription,omitempty"`
	// 	AdminConsentDisplayName string `json:"adminConsentDisplayName,omitempty" bson:"adminConsentDisplayName,omitempty"`
	// 	ID                      string `json:"id,omitempty" bson:"id,omitempty"`
	// 	IsEnabled               bool   `json:"isEnabled,omitempty" bson:"isEnabled,omitempty"`
	// 	IsPrivate               bool   `json:"isPrivate,omitempty" bson:"isPrivate,omitempty"`
	// 	Type                    string `json:"type,omitempty" bson:"type,omitempty"`
	// 	UserConsentDescription  string `json:"userConsentDescription,omitempty" bson:"userConsentDescription,omitempty"`
	// 	UserConsentDisplayName  string `json:"userConsentDisplayName,omitempty" bson:"userConsentDisplayName,omitempty"`
	// 	Value                   string `json:"value,omitempty" bson:"value,omitempty"`
	// } `json:"publishedPermissionScopes,omitempty" bson:"publishedPermissionScopes,omitempty"`
	// PublisherName                          string    `json:"publisherName,omitempty" bson:"publisherName,omitempty"`
	// RefreshTokensValidFromDateTime         time.Time `json:"refreshTokensValidFromDateTime,omitempty" bson:"refreshTokensValidFromDateTime,omitempty"`
	// ReleaseTrack                           any       `json:"releaseTrack,omitempty" bson:"releaseTrack,omitempty"`
	// RenewedDateTime                        time.Time `json:"renewedDateTime,omitempty" bson:"renewedDateTime,omitempty"`
	// ReplyUrls                              []string  `json:"replyUrls,omitempty" bson:"replyUrls,omitempty"`
	// ResourceBehaviorOptions                []any     `json:"resourceBehaviorOptions,omitempty" bson:"resourceBehaviorOptions,omitempty"`
	// ResourceProvisioningOptions            []any     `json:"resourceProvisioningOptions,omitempty" bson:"resourceProvisioningOptions,omitempty"`
	// ResourceSpecificApplicationPermissions []any     `json:"resourceSpecificApplicationPermissions,omitempty" bson:"resourceSpecificApplicationPermissions,omitempty"`
	// SamlMetadataURL                        any       `json:"samlMetadataUrl,omitempty" bson:"samlMetadataUrl,omitempty"`
	// SamlSloBindingType                     string    `json:"samlSLOBindingType,omitempty" bson:"samlSLOBindingType,omitempty"`
	// SamlSingleSignOnSettings               any       `json:"samlSingleSignOnSettings,omitempty" bson:"samlSingleSignOnSettings,omitempty"`
	// SecurityEnabled                        bool      `json:"securityEnabled,omitempty" bson:"securityEnabled,omitempty"`
	// SecurityIdentifier    string   `json:"securityIdentifier,omitempty" bson:"securityIdentifier,omitempty"`
	ServicePrincipalNames []string `json:"servicePrincipalNames,omitempty" bson:"servicePrincipalNames,omitempty"`
	ServicePrincipalType  string   `json:"servicePrincipalType,omitempty" bson:"servicePrincipalType,omitempty"`
	// ServiceProvisioningErrors              []any     `json:"serviceProvisioningErrors,omitempty" bson:"serviceProvisioningErrors,omitempty"`
	// ShowInAddressList                      any       `json:"showInAddressList,omitempty" bson:"showInAddressList,omitempty"`
	// SignInAudience                         string    `json:"signInAudience,omitempty" bson:"signInAudience,omitempty"`
	// SignInSessionsValidFromDateTime        time.Time `json:"signInSessionsValidFromDateTime,omitempty" bson:"signInSessionsValidFromDateTime,omitempty"`
	// State                                  any       `json:"state,omitempty" bson:"state,omitempty"`
	// StreetAddress                          any       `json:"streetAddress,omitempty" bson:"streetAddress,omitempty"`
	// Surname                                *string   `json:"surname,omitempty" bson:"surname,omitempty"`
	Tags []string `json:"tags" bson:"tags"`
	// Theme                                  any       `json:"theme,omitempty" bson:"theme,omitempty"`
	// TokenEncryptionKeyID                   any       `json:"tokenEncryptionKeyId,omitempty" bson:"tokenEncryptionKeyId,omitempty"`
	// TokensRevocationDateTime               any       `json:"tokensRevocationDateTime,omitempty" bson:"tokensRevocationDateTime,omitempty"`
	// UniqueName                             any       `json:"uniqueName,omitempty" bson:"uniqueName,omitempty"`
	// UsageLocation                          *string   `json:"usageLocation,omitempty" bson:"usageLocation,omitempty"`
	UserPrincipalName string `json:"userPrincipalName,omitempty" bson:"userPrincipalName,omitempty"`
	UserType          string `json:"userType,omitempty" bson:"userType,omitempty"`
	// VerifiedPublisher                      *struct {
	// 	AddedDateTime       any `json:"addedDateTime,omitempty" bson:"addedDateTime,omitempty"`
	// 	DisplayName         any `json:"displayName,omitempty" bson:"displayName,omitempty"`
	// 	VerifiedPublisherID any `json:"verifiedPublisherId,omitempty" bson:"verifiedPublisherId,omitempty"`
	// } `json:"verifiedPublisher,omitempty" bson:"verifiedPublisher,omitempty"`
	// Visibility             string `json:"visibility,omitempty" bson:"visibility,omitempty"`
	// WellKnownObject        any    `json:"wellKnownObject,omitempty" bson:"wellKnownObject,omitempty"`
	// WritebackConfiguration *struct {
	// 	IsEnabled           any `json:"isEnabled,omitempty" bson:"isEnabled,omitempty"`
	// 	OnPremisesGroupType any `json:"onPremisesGroupType,omitempty" bson:"onPremisesGroupType,omitempty"`
	// } `json:"writebackConfiguration,omitempty" bson:"writebackConfiguration,omitempty"`
}

//
//

type B2CUser struct {
	AccountEnabled         bool `json:"accountEnabled"`
	AgeGroup               any  `json:"ageGroup"`
	AlternativeSecurityIds []struct {
		IdentityProvider any     `json:"identityProvider"`
		Key              string  `json:"key"`
		Type             float64 `json:"type"`
	} `json:"alternativeSecurityIds"`
	AssignedLicenses  []any `json:"assignedLicenses"`
	AssignedPlans     []any `json:"assignedPlans"`
	AuthorizationInfo struct {
		CertificateUserIds []any `json:"certificateUserIds"`
	} `json:"authorizationInfo"`
	BusinessPhones                 []any `json:"businessPhones"`
	City                           any   `json:"city"`
	CloudRealtimeCommunicationInfo struct {
		CloudMsRtcOwnerUrn          any   `json:"cloudMSRtcOwnerUrn"`
		CloudMsRtcPolicyAssignments []any `json:"cloudMSRtcPolicyAssignments"`
		CloudMsRtcPool              any   `json:"cloudMSRtcPool"`
		CloudMsRtcServiceAttributes any   `json:"cloudMSRtcServiceAttributes"`
		CloudSipLine                any   `json:"cloudSipLine"`
		CloudSipProxyAddress        any   `json:"cloudSipProxyAddress"`
		IsSipEnabled                any   `json:"isSipEnabled"`
	} `json:"cloudRealtimeCommunicationInfo"`
	CompanyName             any       `json:"companyName"`
	ConsentProvidedForMinor any       `json:"consentProvidedForMinor"`
	Country                 any       `json:"country"`
	CreatedDateTime         time.Time `json:"createdDateTime"`
	CreationType            *string   `json:"creationType"`
	DeletedDateTime         any       `json:"deletedDateTime"`
	Department              any       `json:"department"`
	DeviceKeys              []any     `json:"deviceKeys"`
	DisplayName             string    `json:"displayName"`
	EmployeeHireDate        any       `json:"employeeHireDate"`
	EmployeeID              any       `json:"employeeId"`
	EmployeeLeaveDateTime   any       `json:"employeeLeaveDateTime"`
	EmployeeOrgData         any       `json:"employeeOrgData"`
	EmployeeType            any       `json:"employeeType"`
	UnknownFields                 jsontext.Value `json:",unknown"`
	ExternalUserConvertedOn any            `json:"externalUserConvertedOn"`
	ExternalUserInformation struct {
		AcceptedAsMail   *string    `json:"acceptedAsMail"`
		AcceptedDateTime *time.Time `json:"acceptedDateTime"`
		InviteReplyUrls  []any      `json:"inviteReplyUrls"`
		InviteResources  []any      `json:"inviteResources"`
		InvitedAsMail    *string    `json:"invitedAsMail"`
		InvitedDateTime  time.Time  `json:"invitedDateTime"`
		SignInNames      []string   `json:"signInNames"`
	} `json:"externalUserInformation"`
	ExternalUserState               *string    `json:"externalUserState"`
	ExternalUserStateChangeDateTime *time.Time `json:"externalUserStateChangeDateTime"`
	FaxNumber                       any        `json:"faxNumber"`
	GivenName                       *string    `json:"givenName"`
	ID                              string     `json:"id"`
	Identities                      []struct {
		Issuer           string  `json:"issuer"`
		IssuerAssignedID *string `json:"issuerAssignedId"`
		SignInType       string  `json:"signInType"`
	} `json:"identities"`
	ImAddresses                   []any   `json:"imAddresses"`
	InfoCatalogs                  []any   `json:"infoCatalogs"`
	IsLicenseReconciliationNeeded bool    `json:"isLicenseReconciliationNeeded"`
	IsManagementRestricted        any     `json:"isManagementRestricted"`
	IsResourceAccount             any     `json:"isResourceAccount"`
	JobTitle                      any     `json:"jobTitle"`
	LegalAgeGroupClassification   any     `json:"legalAgeGroupClassification"`
	Mail                          *string `json:"mail"`
	MailNickname                  string  `json:"mailNickname"`
	MobilePhone                   any     `json:"mobilePhone"`
	NetID                         string  `json:"netId"`
	OfficeLocation                any     `json:"officeLocation"`
	OnPremisesDistinguishedName   any     `json:"onPremisesDistinguishedName"`
	OnPremisesDomainName          any     `json:"onPremisesDomainName"`
	OnPremisesExtensionAttributes struct {
		ExtensionAttribute1  any `json:"extensionAttribute1"`
		ExtensionAttribute10 any `json:"extensionAttribute10"`
		ExtensionAttribute11 any `json:"extensionAttribute11"`
		ExtensionAttribute12 any `json:"extensionAttribute12"`
		ExtensionAttribute13 any `json:"extensionAttribute13"`
		ExtensionAttribute14 any `json:"extensionAttribute14"`
		ExtensionAttribute15 any `json:"extensionAttribute15"`
		ExtensionAttribute2  any `json:"extensionAttribute2"`
		ExtensionAttribute3  any `json:"extensionAttribute3"`
		ExtensionAttribute4  any `json:"extensionAttribute4"`
		ExtensionAttribute5  any `json:"extensionAttribute5"`
		ExtensionAttribute6  any `json:"extensionAttribute6"`
		ExtensionAttribute7  any `json:"extensionAttribute7"`
		ExtensionAttribute8  any `json:"extensionAttribute8"`
		ExtensionAttribute9  any `json:"extensionAttribute9"`
	} `json:"onPremisesExtensionAttributes"`
	OnPremisesImmutableID        any   `json:"onPremisesImmutableId"`
	OnPremisesLastSyncDateTime   any   `json:"onPremisesLastSyncDateTime"`
	OnPremisesObjectIdentifier   any   `json:"onPremisesObjectIdentifier"`
	OnPremisesProvisioningErrors []any `json:"onPremisesProvisioningErrors"`
	OnPremisesSamAccountName     any   `json:"onPremisesSamAccountName"`
	OnPremisesSecurityIdentifier any   `json:"onPremisesSecurityIdentifier"`
	OnPremisesSipInfo            struct {
		IsSipEnabled          bool `json:"isSipEnabled"`
		SipDeploymentLocation any  `json:"sipDeploymentLocation"`
		SipPrimaryAddress     any  `json:"sipPrimaryAddress"`
	} `json:"onPremisesSipInfo"`
	OnPremisesSyncEnabled       any      `json:"onPremisesSyncEnabled"`
	OnPremisesUserPrincipalName any      `json:"onPremisesUserPrincipalName"`
	OtherMails                  []string `json:"otherMails"`
	PasswordPolicies            *string  `json:"passwordPolicies"`
	PasswordProfile             *struct {
		ForceChangePasswordNextSignIn        bool `json:"forceChangePasswordNextSignIn"`
		ForceChangePasswordNextSignInWithMfa bool `json:"forceChangePasswordNextSignInWithMfa"`
		Password                             any  `json:"password"`
	} `json:"passwordProfile"`
	PortalSetting                   any       `json:"portalSetting"`
	PostalCode                      any       `json:"postalCode"`
	PreferredDataLocation           any       `json:"preferredDataLocation"`
	PreferredLanguage               any       `json:"preferredLanguage"`
	ProvisionedPlans                []any     `json:"provisionedPlans"`
	ProxyAddresses                  []string  `json:"proxyAddresses"`
	RefreshTokensValidFromDateTime  time.Time `json:"refreshTokensValidFromDateTime"`
	ReleaseTrack                    any       `json:"releaseTrack"`
	SecurityIdentifier              string    `json:"securityIdentifier"`
	ServiceProvisioningErrors       []any     `json:"serviceProvisioningErrors"`
	ShowInAddressList               *bool     `json:"showInAddressList"`
	SignInSessionsValidFromDateTime time.Time `json:"signInSessionsValidFromDateTime"`
	State                           any       `json:"state"`
	StreetAddress                   any       `json:"streetAddress"`
	Surname                         *string   `json:"surname"`
	UsageLocation                   any       `json:"usageLocation"`
	UserPrincipalName               string    `json:"userPrincipalName"`
	UserType                        string    `json:"userType"`
}

//
//

type GetAllB2CTenantUsersResponse struct {
	Context  string    `json:"@odata.context"`
	NextLink string    `json:"@odata.nextLink"`
	Value    []B2CUser `json:"value"`
}

//
//

type B2CUserMinimal struct {
	AccountEnabled  bool      `json:"accountEnabled,omitempty,omitzero" bson:"accountEnabled,omitempty,omitzero"`
	AccountLocked   bool      `json:"accountLocked,omitempty,omitzero" bson:"accountLocked,omitempty,omitzero"`
	CreatedDateTime time.Time `json:"createdDateTime,omitempty,omitzero" bson:"createdDateTime,omitempty,omitzero"`
	CreationType    *string   `json:"creationType,omitempty,omitzero" bson:"creationType,omitempty,omitzero"`
	DisplayName     string    `json:"displayName,omitempty,omitzero" bson:"displayName,omitempty,omitzero"`
	ExternalUserInformation struct {
		AcceptedAsMail   any       `json:"acceptedAsMail,omitempty,omitzero" bson:"acceptedAsMail,omitempty,omitzero"`
		AcceptedDateTime any       `json:"acceptedDateTime,omitempty,omitzero" bson:"acceptedDateTime,omitempty,omitzero"`
		InviteReplyUrls  []any     `json:"inviteReplyUrls,omitempty,omitzero" bson:"inviteReplyUrls,omitempty,omitzero"`
		InviteResources  []any     `json:"inviteResources,omitempty,omitzero" bson:"inviteResources,omitempty,omitzero"`
		InvitedAsMail    any       `json:"invitedAsMail,omitempty,omitzero" bson:"invitedAsMail,omitempty,omitzero"`
		InvitedDateTime  time.Time `json:"invitedDateTime,omitempty,omitzero" bson:"invitedDateTime,omitempty,omitzero"`
		SignInNames      []string  `json:"signInNames,omitempty,omitzero" bson:"signInNames,omitempty,omitzero"`
	} `json:"externalUserInformation,omitempty,omitzero" bson:"externalUserInformation,omitempty,omitzero"`
	ID         string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Identities []struct {
		Issuer           string  `json:"issuer,omitempty,omitzero" bson:"issuer,omitempty,omitzero"`
		IssuerAssignedID *string `json:"issuerAssignedId,omitempty,omitzero" bson:"issuerAssignedId,omitempty,omitzero"`
		SignInType       string  `json:"signInType,omitempty,omitzero" bson:"signInType,omitempty,omitzero"`
	} `json:"identities,omitempty,omitzero" bson:"identities,omitempty,omitzero"`
	MailNickname     string `json:"mailNickname,omitempty,omitzero" bson:"mailNickname,omitempty,omitzero"`
	PasswordPolicies any    `json:"passwordPolicies,omitempty,omitzero" bson:"passwordPolicies,omitempty,omitzero"`
	PasswordProfile  *struct {
		ForceChangePasswordNextSignIn        bool `json:"forceChangePasswordNextSignIn,omitempty,omitzero" bson:"forceChangePasswordNextSignIn,omitempty,omitzero"`
		ForceChangePasswordNextSignInWithMfa bool `json:"forceChangePasswordNextSignInWithMfa,omitempty,omitzero" bson:"forceChangePasswordNextSignInWithMfa,omitempty,omitzero"`
		Password                             any  `json:"password,omitempty,omitzero" bson:"password,omitempty,omitzero"`
	} `json:"passwordProfile,omitempty,omitzero" bson:"passwordProfile,omitempty,omitzero"`
	RefreshTokensValidFromDateTime  time.Time `json:"refreshTokensValidFromDateTime,omitempty,omitzero" bson:"refreshTokensValidFromDateTime,omitempty,omitzero"`
	SecurityIdentifier              string    `json:"securityIdentifier,omitempty,omitzero" bson:"securityIdentifier,omitempty,omitzero"`
	SignInSessionsValidFromDateTime time.Time `json:"signInSessionsValidFromDateTime,omitempty,omitzero" bson:"signInSessionsValidFromDateTime,omitempty,omitzero"`
	UserPrincipalName               string    `json:"userPrincipalName,omitempty,omitzero" bson:"userPrincipalName,omitempty,omitzero"`
	UserType                        string    `json:"userType,omitempty,omitzero" bson:"userType,omitempty,omitzero"`
	B2CTenant                       string    `json:"b2cTenant,omitempty,omitzero" bson:"b2cTenant,omitempty,omitzero"`
	LastDBSync                      time.Time `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
	ExtensionLastLogonTime          time.Time `json:"extensionLastLogonTime,omitempty,omitzero" bson:"extensionLastLogonTime,omitempty,omitzero"`
	ExtensionPasswordResetOn        time.Time `json:"extensionPasswordResetOn,omitempty,omitzero" bson:"extensionPasswordResetOn,omitempty,omitzero"`
	// TenantName                      string    `json:"tenantName" bson:"tenantName"`
}
