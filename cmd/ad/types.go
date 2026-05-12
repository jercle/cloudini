package ad

import "time"

type SearchResponse struct {
	Controls []any `json:"Controls,omitempty,omitzero" bson:"Controls,omitempty,omitzero"`
	Entries  []struct {
		Attributes []struct {
			ByteValues []string `json:"ByteValues,omitempty,omitzero" bson:"ByteValues,omitempty,omitzero"`
			Name       string   `json:"Name,omitempty,omitzero" bson:"Name,omitempty,omitzero"`
			Values     []string `json:"Values,omitempty,omitzero" bson:"Values,omitempty,omitzero"`
		} `json:"Attributes,omitempty,omitzero" bson:"Attributes,omitempty,omitzero"`
		Dn string `json:"DN,omitempty,omitzero" bson:"DN,omitempty,omitzero"`
	} `json:"Entries,omitempty,omitzero" bson:"Entries,omitempty,omitzero"`
	Referrals []any `json:"Referrals,omitempty,omitzero" bson:"Referrals,omitempty,omitzero"`
}

type ADUser struct {
	ObjectClass                   []string  `json:"objectClass,omitempty,omitzero" bson:"objectClass,omitempty,omitzero"`
	CN                            string    `json:"cn,omitempty,omitzero" bson:"cn,omitempty,omitzero"`
	Description                   string    `json:"description,omitempty,omitzero" bson:"description,omitempty,omitzero"`
	Domain                        string    `json:"domain,omitempty,omitzero" bson:"domain,omitempty,omitzero"`
	GivenName                     string    `json:"givenName,omitempty,omitzero" bson:"givenName,omitempty,omitzero"`
	Initials                      string    `json:"initials,omitempty,omitzero" bson:"initials,omitempty,omitzero"`
	DistinguishedName             string    `json:"distinguishedName,omitempty,omitzero" bson:"distinguishedName,omitempty,omitzero"`
	WhenCreated                   string    `json:"whenCreated,omitempty,omitzero" bson:"whenCreated,omitempty,omitzero"`
	WhenChanged                   string    `json:"whenChanged,omitempty,omitzero" bson:"whenChanged,omitempty,omitzero"`
	DisplayName                   string    `json:"displayName,omitempty,omitzero" bson:"displayName,omitempty,omitzero"`
	MemberOf                      []string  `json:"memberOf,omitempty,omitzero" bson:"memberOf,omitempty,omitzero"`
	Name                          string    `json:",omitempty,omitzero" bson:",omitempty,omitzero"`
	ObjectGUID                    string    `json:"objectGUID,omitempty,omitzero" bson:"objectGUID,omitempty,omitzero"`
	LastLogon                     string    `json:"lastLogon,omitempty,omitzero" bson:"lastLogon,omitempty,omitzero"`
	PwdLastSet                    string    `json:"pwdLastSet,omitempty,omitzero" bson:"pwdLastSet,omitempty,omitzero"`
	ObjectSid                     string    `json:"objectSid,omitempty,omitzero" bson:"objectSid,omitempty,omitzero"`
	AccountExpires                string    `json:"accountExpires,omitempty,omitzero" bson:"accountExpires,omitempty,omitzero"`
	SAMAccountName                string    `json:"sAMAccountName,omitempty,omitzero" bson:"sAMAccountName,omitempty,omitzero"`
	UserPrincipalName             string    `json:"userPrincipalName,omitempty,omitzero" bson:"userPrincipalName,omitempty,omitzero"`
	LockoutTime                   string    `json:"lockoutTime,omitempty,omitzero" bson:"lockoutTime,omitempty,omitzero"`
	LastLogonTimestamp            string    `json:"lastLogonTimestamp,omitempty,omitzero" bson:"lastLogonTimestamp,omitempty,omitzero"`
	MsDSExternalDirectoryObjectId string    `json:"msDS-ExternalDirectoryObjectId,omitempty,omitzero" bson:"msDS-ExternalDirectoryObjectId,omitempty,omitzero"`
	Mail                          string    `json:"mail,omitempty,omitzero" bson:"mail,omitempty,omitzero"`
	MailNickname                  string    `json:"mailNickname,omitempty,omitzero" bson:"mailNickname,omitempty,omitzero"`
	LastDBSync                    time.Time `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
}
