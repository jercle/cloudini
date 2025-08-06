package forgerock

import "time"

type LDAPConectorResponse struct {
	PagedResultsCookie      any           `json:"pagedResultsCookie"`
	RemainingPagedResults   float64       `json:"remainingPagedResults"`
	Result                  []interface{} `json:"result"`
	ResultCount             float64       `json:"resultCount"`
	TotalPagedResults       float64       `json:"totalPagedResults"`
	TotalPagedResultsPolicy string        `json:"totalPagedResultsPolicy"`
}

type LDAPConnectorGroup struct {
	ID             string                   `json:"_id"`
	MemberId       []string                 `json:"_memberId"`
	Cn             string                   `json:"cn"`
	Description    string                   `json:"description,omitempty"`
	DisplayName    string                   `json:"displayName,omitempty"`
	Dn             string                   `json:"dn"`
	GroupScope     string                   `json:"groupScope"`
	GroupType      string                   `json:"groupType"`
	Info           string                   `json:"info,omitempty"`
	ManagedBy      string                   `json:"managedBy,omitempty"`
	Member         LDAPConnectorGroupMember `json:"member"`
	MemberOf       []string                 `json:"memberOf"`
	ProxyAddresses []string                 `json:"proxyAddresses"`
	SamAccountName string                   `json:"samAccountName"`
	USnChanged     string                   `json:"uSNChanged"`
	USnCreated     string                   `json:"uSNCreated"`
	WhenChanged    string                   `json:"whenChanged"`
	WhenCreated    string                   `json:"whenCreated"`
}

type LDAPConnectorGroupMember struct {
	SamAccountNames    []string `json:"samAccountNames"`
	DistinguishedNames []string `json:"distinguishedNames"`
}

type LDAPConnectorUser struct {
	ID                         string    `json:"_id"`
	AccountExpiresRead         string    `json:"accountExpiresRead"`
	BadPasswordTime            string    `json:"badPasswordTime,omitempty"`
	BadPwdCount                string    `json:"badPwdCount,omitempty"`
	CanonicalName              string    `json:"canonicalName"`
	Cn                         string    `json:"cn"`
	Company                    string    `json:"company,omitempty"`
	CountryCode                string    `json:"countryCode"`
	Description                string    `json:"description,omitempty"`
	DisplayName                string    `json:"displayName,omitempty"`
	Dn                         string    `json:"dn"`
	DontExpirePassword         bool      `json:"dontExpirePassword"`
	GivenName                  string    `json:"givenName,omitempty"`
	Initials                   string    `json:"initials,omitempty"`
	IsActive                   bool      `json:"isActive"`
	LastLogon                  string    `json:"lastLogon,omitempty"`
	LastLogonTimestamp         time.Time `json:"lastLogonTimestamp,omitempty"`
	LdapGroups                 []string  `json:"ldapGroups"`
	LockOut                    bool      `json:"lockOut"`
	LockoutTime                string    `json:"lockoutTime,omitempty"`
	Mail                       string    `json:"mail,omitempty"`
	MsExchPoliciesExcluded     []string  `json:"msExchPoliciesExcluded"`
	ObjectGuid                 string    `json:"objectGUID"`
	ObjectSid                  string    `json:"objectSid"`
	OtherHomePhone             []any     `json:"otherHomePhone"`
	PasswordExpired            bool      `json:"passwordExpired"`
	PasswordExpiryTimeComputed string    `json:"passwordExpiryTimeComputed"`
	PasswordNotRequired        bool      `json:"passwordNotRequired"`
	ProxyAddresses             []string  `json:"proxyAddresses"`
	PwdLastSet                 string    `json:"pwdLastSet"`
	SAmAccountName             string    `json:"sAMAccountName"`
	SmartcardRequired          bool      `json:"smartcardRequired"`
	Sn                         string    `json:"sn,omitempty"`
	TelephoneNumber            string    `json:"telephoneNumber,omitempty"`
	USnChanged                 string    `json:"uSNChanged"`
	USnCreated                 string    `json:"uSNCreated"`
	UserAccountControl         string    `json:"userAccountControl"`
	UserPrincipalName          string    `json:"userPrincipalName,omitempty"`
	WhenChanged                string    `json:"whenChanged"`
	WhenCreated                string    `json:"whenCreated"`
}
