package forgerock

import "time"

type LDAPConectorResponse struct {
	PagedResultsCookie      any           `json:"pagedResultsCookie,omitempty,omitzero" bson:"pagedResultsCookie,omitempty,omitzero"`
	RemainingPagedResults   float64       `json:"remainingPagedResults,omitempty,omitzero" bson:"remainingPagedResults,omitempty,omitzero"`
	Result                  []interface{} `json:"result,omitempty,omitzero" bson:"result,omitempty,omitzero"`
	ResultCount             float64       `json:"resultCount,omitempty,omitzero" bson:"resultCount,omitempty,omitzero"`
	TotalPagedResults       float64       `json:"totalPagedResults,omitempty,omitzero" bson:"totalPagedResults,omitempty,omitzero"`
	TotalPagedResultsPolicy string        `json:"totalPagedResultsPolicy,omitempty,omitzero" bson:"totalPagedResultsPolicy,omitempty,omitzero"`
}
type LDAPConectorGrpResponse struct {
	PagedResultsCookie      any                  `json:"pagedResultsCookie,omitempty,omitzero" bson:"pagedResultsCookie,omitempty,omitzero"`
	RemainingPagedResults   float64              `json:"remainingPagedResults,omitempty,omitzero" bson:"remainingPagedResults,omitempty,omitzero"`
	Result                  []LDAPConnectorGroup `json:"result,omitempty,omitzero" bson:"result,omitempty,omitzero"`
	ResultCount             float64              `json:"resultCount,omitempty,omitzero" bson:"resultCount,omitempty,omitzero"`
	TotalPagedResults       float64              `json:"totalPagedResults,omitempty,omitzero" bson:"totalPagedResults,omitempty,omitzero"`
	TotalPagedResultsPolicy string               `json:"totalPagedResultsPolicy,omitempty,omitzero" bson:"totalPagedResultsPolicy,omitempty,omitzero"`
}

type LDAPConnectorGroup struct {
	ID             string                   `json:"_id,omitempty,omitzero" bson:"_id,omitempty,omitzero"`
	MemberId       []string                 `json:"_memberId,omitempty,omitzero" bson:"memberId,omitempty,omitzero"`
	Cn             string                   `json:"cn,omitempty,omitzero" bson:"cn,omitempty,omitzero"`
	Description    string                   `json:"description,omitempty" bson:"description,omitempty"`
	DisplayName    string                   `json:"displayName,omitempty" bson:"displayName,omitempty"`
	Dn             string                   `json:"dn,omitempty,omitzero" bson:"dn,omitempty,omitzero"`
	GroupScope     string                   `json:"groupScope,omitempty,omitzero" bson:"groupScope,omitempty,omitzero"`
	GroupType      string                   `json:"groupType,omitempty,omitzero" bson:"groupType,omitempty,omitzero"`
	Info           string                   `json:"info,omitempty" bson:"info,omitempty"`
	ManagedBy      string                   `json:"managedBy,omitempty" bson:"managedBy,omitempty"`
	Member         LDAPConnectorGroupMember `json:"member,omitempty,omitzero" bson:"member,omitempty,omitzero"`
	MemberOf       []string                 `json:"memberOf,omitempty,omitzero" bson:"memberOf,omitempty,omitzero"`
	ProxyAddresses []string                 `json:"proxyAddresses,omitempty,omitzero" bson:"proxyAddresses,omitempty,omitzero"`
	SamAccountName string                   `json:"samAccountName,omitempty,omitzero" bson:"samAccountName,omitempty,omitzero"`
	USnChanged     string                   `json:"uSNChanged,omitempty,omitzero" bson:"uSNChanged,omitempty,omitzero"`
	USnCreated     string                   `json:"uSNCreated,omitempty,omitzero" bson:"uSNCreated,omitempty,omitzero"`
	WhenChanged    string                   `json:"whenChanged,omitempty,omitzero" bson:"whenChanged,omitempty,omitzero"`
	WhenCreated    string                   `json:"whenCreated,omitempty,omitzero" bson:"whenCreated,omitempty,omitzero"`
	LastDbSync     time.Time                `json:"lastDbSync,omitempty,omitzero" bson:"lastDbSync,omitempty,omitzero"`
}

type LDAPConnectorGroupMember struct {
	SamAccountNames    []string `json:"samAccountNames,omitempty,omitzero" bson:"samAccountNames,omitempty,omitzero"`
	DistinguishedNames []string `json:"distinguishedNames,omitempty,omitzero" bson:"distinguishedNames,omitempty,omitzero"`
}

type LDAPConnectorUser struct {
	ID                         string    `json:"_id,omitempty,omitzero" bson:"_id,omitempty,omitzero"`
	AccountExpiresRead         string    `json:"accountExpiresRead,omitempty,omitzero" bson:"accountExpiresRead,omitempty,omitzero"`
	BadPasswordTime            string    `json:"badPasswordTime,omitempty" bson:"badPasswordTime,omitempty"`
	BadPwdCount                string    `json:"badPwdCount,omitempty" bson:"badPwdCount,omitempty"`
	CanonicalName              string    `json:"canonicalName,omitempty,omitzero" bson:"canonicalName,omitempty,omitzero"`
	Cn                         string    `json:"cn,omitempty,omitzero" bson:"cn,omitempty,omitzero"`
	Company                    string    `json:"company,omitempty" bson:"company,omitempty"`
	CountryCode                string    `json:"countryCode,omitempty,omitzero" bson:"countryCode,omitempty,omitzero"`
	Description                string    `json:"description,omitempty" bson:"description,omitempty"`
	DisplayName                string    `json:"displayName,omitempty" bson:"displayName,omitempty"`
	Dn                         string    `json:"dn,omitempty,omitzero" bson:"dn,omitempty,omitzero"`
	DontExpirePassword         bool      `json:"dontExpirePassword,omitempty,omitzero" bson:"dontExpirePassword,omitempty,omitzero"`
	GivenName                  string    `json:"givenName,omitempty" bson:"givenName,omitempty"`
	Initials                   string    `json:"initials,omitempty" bson:"initials,omitempty"`
	IsActive                   bool      `json:"isActive,omitempty,omitzero" bson:"isActive,omitempty,omitzero"`
	LastLogon                  string    `json:"lastLogon,omitempty" bson:"lastLogon,omitempty"`
	LastLogonTimestamp         time.Time `json:"lastLogonTimestamp,omitempty" bson:"lastLogonTimestamp,omitempty"`
	LdapGroups                 []string  `json:"ldapGroups,omitempty,omitzero" bson:"ldapGroups,omitempty,omitzero"`
	LockOut                    bool      `json:"lockOut,omitempty,omitzero" bson:"lockOut,omitempty,omitzero"`
	LockoutTime                string    `json:"lockoutTime,omitempty" bson:"lockoutTime,omitempty"`
	Mail                       string    `json:"mail,omitempty" bson:"mail,omitempty"`
	MsExchPoliciesExcluded     []string  `json:"msExchPoliciesExcluded,omitempty,omitzero" bson:"msExchPoliciesExcluded,omitempty,omitzero"`
	ObjectGuid                 string    `json:"objectGUID,omitempty,omitzero" bson:"objectGUID,omitempty,omitzero"`
	ObjectSid                  string    `json:"objectSid,omitempty,omitzero" bson:"objectSid,omitempty,omitzero"`
	OtherHomePhone             []any     `json:"otherHomePhone,omitempty,omitzero" bson:"otherHomePhone,omitempty,omitzero"`
	PasswordExpired            bool      `json:"passwordExpired,omitempty,omitzero" bson:"passwordExpired,omitempty,omitzero"`
	PasswordExpiryTimeComputed string    `json:"passwordExpiryTimeComputed,omitempty,omitzero" bson:"passwordExpiryTimeComputed,omitempty,omitzero"`
	PasswordNotRequired        bool      `json:"passwordNotRequired,omitempty,omitzero" bson:"passwordNotRequired,omitempty,omitzero"`
	ProxyAddresses             []string  `json:"proxyAddresses,omitempty,omitzero" bson:"proxyAddresses,omitempty,omitzero"`
	PwdLastSet                 string    `json:"pwdLastSet,omitempty,omitzero" bson:"pwdLastSet,omitempty,omitzero"`
	SAmAccountName             string    `json:"sAMAccountName,omitempty,omitzero" bson:"sAMAccountName,omitempty,omitzero"`
	SmartcardRequired          bool      `json:"smartcardRequired,omitempty,omitzero" bson:"smartcardRequired,omitempty,omitzero"`
	Sn                         string    `json:"sn,omitempty" bson:"sn,omitempty"`
	TelephoneNumber            string    `json:"telephoneNumber,omitempty" bson:"telephoneNumber,omitempty"`
	USnChanged                 string    `json:"uSNChanged,omitempty,omitzero" bson:"uSNChanged,omitempty,omitzero"`
	USnCreated                 string    `json:"uSNCreated,omitempty,omitzero" bson:"uSNCreated,omitempty,omitzero"`
	UserAccountControl         string    `json:"userAccountControl,omitempty,omitzero" bson:"userAccountControl,omitempty,omitzero"`
	UserPrincipalName          string    `json:"userPrincipalName,omitempty" bson:"userPrincipalName,omitempty"`
	WhenChanged                string    `json:"whenChanged,omitempty,omitzero" bson:"whenChanged,omitempty,omitzero"`
	WhenCreated                string    `json:"whenCreated,omitempty,omitzero" bson:"whenCreated,omitempty,omitzero"`
}
