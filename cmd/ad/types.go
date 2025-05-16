package ad

import "time"

type SearchResponse struct {
	Controls []any `json:"Controls"`
	Entries  []struct {
		Attributes []struct {
			ByteValues []string `json:"ByteValues"`
			Name       string   `json:"Name"`
			Values     []string `json:"Values"`
		} `json:"Attributes"`
		Dn string `json:"DN"`
	} `json:"Entries"`
	Referrals []any `json:"Referrals"`
}

type ADUser struct {
	ObjectClass                   []string  `json:"objectClass"`
	CN                            string    `json:"cn"`
	Description                   string    `json:"description"`
	Domain                        string    `json:"domain"`
	GivenName                     string    `json:"givenName"`
	Initials                      string    `json:"initials"`
	DistinguishedName             string    `json:"distinguishedName"`
	WhenCreated                   string    `json:"whenCreated"`
	WhenChanged                   string    `json:"whenChanged"`
	DisplayName                   string    `json:"displayName"`
	MemberOf                      []string  `json:"memberOf"`
	Name                          string    `json:""`
	ObjectGUID                    string    `json:"objectGUID"`
	LastLogon                     string    `json:"lastLogon"`
	PwdLastSet                    string    `json:"pwdLastSet"`
	ObjectSid                     string    `json:"objectSid"`
	AccountExpires                string    `json:"accountExpires"`
	SAMAccountName                string    `json:"sAMAccountName"`
	UserPrincipalName             string    `json:"userPrincipalName"`
	LockoutTime                   string    `json:"lockoutTime"`
	LastLogonTimestamp            string    `json:"lastLogonTimestamp"`
	MsDSExternalDirectoryObjectId string    `json:"msDS-ExternalDirectoryObjectId"`
	Mail                          string    `json:"mail"`
	MailNickname                  string    `json:"mailNickname"`
	LastDBSync                    time.Time `json:"lastDBSync"`
}
