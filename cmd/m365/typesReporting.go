package m365

import "time"

type GetSubscribedSkusByIdResponse struct {
	Odata_Context string           `json:"@odata.context,omitempty,omitzero" bson:"@odata.context,omitempty,omitzero"`
	Value         []M365LicenseSku `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
}

//
//

type M365LicenseSku struct {
	// AccountID        string `json:"accountId,omitempty,omitzero" bson:"accountId,omitempty,omitzero"`
	// AccountName      string `json:"accountName,omitempty,omitzero" bson:"accountName,omitempty,omitzero"`
	// AppliesTo        string `json:"appliesTo,omitempty,omitzero" bson:"appliesTo,omitempty,omitzero"`
	// CapabilityStatus string `json:"capabilityStatus,omitempty,omitzero" bson:"capabilityStatus,omitempty,omitzero"`
	ConsumedUnits uint32 `json:"consumedUnits" bson:"consumedUnits"`
	// ID               string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	// OverageUnits struct {
	// 	Enabled   float64 `json:"enabled,omitempty,omitzero" bson:"enabled,omitempty,omitzero"`
	// 	LockedOut float64 `json:"lockedOut,omitempty,omitzero" bson:"lockedOut,omitempty,omitzero"`
	// 	Suspended float64 `json:"suspended,omitempty,omitzero" bson:"suspended,omitempty,omitzero"`
	// 	Warning   float64 `json:"warning,omitempty,omitzero" bson:"warning,omitempty,omitzero"`
	// } `json:"overageUnits,omitempty,omitzero" bson:"overageUnits,omitempty,omitzero"`
	// PrepaidUnits struct {
	// 	Enabled   float64 `json:"enabled,omitempty,omitzero" bson:"enabled,omitempty,omitzero"`
	// 	LockedOut float64 `json:"lockedOut,omitempty,omitzero" bson:"lockedOut,omitempty,omitzero"`
	// 	Suspended float64 `json:"suspended,omitempty,omitzero" bson:"suspended,omitempty,omitzero"`
	// 	Warning   float64 `json:"warning,omitempty,omitzero" bson:"warning,omitempty,omitzero"`
	// } `json:"prepaidUnits,omitempty,omitzero" bson:"prepaidUnits,omitempty,omitzero"`
	// SelfServiceSignupUnits struct {
	// 	Enabled   float64 `json:"enabled,omitempty,omitzero" bson:"enabled,omitempty,omitzero"`
	// 	LockedOut float64 `json:"lockedOut,omitempty,omitzero" bson:"lockedOut,omitempty,omitzero"`
	// 	Suspended float64 `json:"suspended,omitempty,omitzero" bson:"suspended,omitempty,omitzero"`
	// 	Warning   float64 `json:"warning,omitempty,omitzero" bson:"warning,omitempty,omitzero"`
	// } `json:"selfServiceSignupUnits,omitempty,omitzero" bson:"selfServiceSignupUnits,omitempty,omitzero"`
	// ServicePlans []struct {
	// 	AppliesTo          string `json:"appliesTo,omitempty,omitzero" bson:"appliesTo,omitempty,omitzero"`
	// 	ProvisioningStatus string `json:"provisioningStatus,omitempty,omitzero" bson:"provisioningStatus,omitempty,omitzero"`
	// 	ServicePlanID      string `json:"servicePlanId,omitempty,omitzero" bson:"servicePlanId,omitempty,omitzero"`
	// 	ServicePlanName    string `json:"servicePlanName,omitempty,omitzero" bson:"servicePlanName,omitempty,omitzero"`
	// 	ServicePlanType    string `json:"servicePlanType,omitempty,omitzero" bson:"servicePlanType,omitempty,omitzero"`
	// } `json:"servicePlans,omitempty,omitzero" bson:"servicePlans,omitempty,omitzero"`
	SkuID         string `json:"skuId,omitempty,omitzero" bson:"skuId,omitempty,omitzero"`
	SkuPartNumber string `json:"skuPartNumber,omitempty,omitzero" bson:"skuPartNumber,omitempty,omitzero"`
	// SubscriptionIds []string `json:"subscriptionIds,omitempty,omitzero" bson:"subscriptionIds,omitempty,omitzero"`
	// TrialUnits      struct {
	// 	Enabled   float64 `json:"enabled,omitempty,omitzero" bson:"enabled,omitempty,omitzero"`
	// 	LockedOut float64 `json:"lockedOut,omitempty,omitzero" bson:"lockedOut,omitempty,omitzero"`
	// 	Suspended float64 `json:"suspended,omitempty,omitzero" bson:"suspended,omitempty,omitzero"`
	// 	Warning   float64 `json:"warning,omitempty,omitzero" bson:"warning,omitempty,omitzero"`
	// } `json:"trialUnits,omitempty,omitzero" bson:"trialUnits,omitempty,omitzero"`
}

//
//

type M365LicenseCounts struct {
	TenantName                   string                          `json:"tenantName" bson:"tenantName"`
	TenantId                     string                          `json:"tenantId" bson:"tenantId"`
	DateDataFetched              time.Time                       `json:"dateDataFetched" bson:"dateDataFetched"`
	UsersCount                   uint32                          `json:"usersCount" bson:"usersCount"`
	UsersCountChecked            uint32                          `json:"usersCountChecked" bson:"usersCountChecked"`
	Users                        []LicenseReportUser             `json:"users" bson:"users"`
	UsersNoLogonLast90DaysCount  uint32                          `json:"usersNoLogonLast90DaysCount" bson:"usersNoLogonLast90DaysCount"`
	UsersNoLogonLast45DaysCount  uint32                          `json:"usersNoLogonLast45DaysCount" bson:"usersNoLogonLast45DaysCount"`
	UsersDisabledCount           uint32                          `json:"usersDisabledCount" bson:"usersDisabledCount"`
	UsersActiveCount             uint32                          `json:"usersActiveCount" bson:"usersActiveCount"`
	LicenseCountsBySkuPartNumber M365LicenseCountBySkuPartNumber `json:"licenseCountsBySkuPartNumber" bson:"licenseCountsBySkuPartNumber"`
}

//
//

type M365LicenseCount struct {
	Users                       []string `json:"users" bson:"users"`
	UsersNoLogonLast90DaysCount uint32   `json:"usersNoLogonLast90DaysCount" bson:"usersNoLogonLast90DaysCount"`
	UsersNoLogonLast45DaysCount uint32   `json:"usersNoLogonLast45DaysCount" bson:"usersNoLogonLast45DaysCount"`
	UsersActiveCount            uint32   `json:"usersActiveCount" bson:"usersActiveCount"`
	UsersDisabledCount          uint32   `json:"usersDisabledCount" bson:"usersDisabledCount"`
	SkuId                       string   `json:"skuId" bson:"skuId"`
	// SkuPartNumber               string   `json:"skuPartNumber" bson:"skuPartNumber"`
	ConsumedUnits        uint32 `json:"consumedUnits" bson:"consumedUnits"`
	ConsumedUnitsChecked uint32 `json:"consumedUnitsChecked" bson:"consumedUnitsChecked"`
}

//
//

type M365LicenseCountBySkuPartNumber map[string]M365LicenseCount

//
//

type LicenseReportUser struct {
	AccountEnabled              bool       `json:"accountEnabled,omitempty,omitzero" bson:"accountEnabled,omitempty,omitzero"`
	IsActiveUser                bool       `json:"isActiveUser,omitempty,omitzero" bson:"isActiveUser,omitempty,omitzero"`
	NoLogonLast45Days           bool       `json:"noLogonLast45Days,omitempty,omitzero" bson:"noLogonLast45Days,omitempty,omitzero"`
	NoLogonLast90Days           bool       `json:"noLogonLast90Days,omitempty,omitzero" bson:"noLogonLast90Days,omitempty,omitzero"`
	AssignedLicenses            []string   `json:"assignedLicenses" bson:"assignedLicenses"`
	CreatedDateTime             time.Time  `json:"createdDateTime,omitempty,omitzero" bson:"createdDateTime,omitempty,omitzero"`
	ID                          string     `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	OnPremisesDistinguishedName string     `json:"onPremisesDistinguishedName,omitempty,omitzero" bson:"onPremisesDistinguishedName,omitempty,omitzero"`
	OnPremisesSamAccountName    string     `json:"onPremisesSamAccountName,omitempty,omitzero" bson:"onPremisesSamAccountName,omitempty,omitzero"`
	LastSignInDateTime          *time.Time `json:"lastSignInDateTime,omitempty,omitzero" bson:"lastSignInDateTime,omitempty,omitzero"`
	UserPrincipalName           string     `json:"userPrincipalName,omitempty,omitzero" bson:"userPrincipalName,omitempty,omitzero"`
}
