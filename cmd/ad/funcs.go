package ad

import (
	"encoding/json"

	"github.com/go-ldap/ldap/v3"
	"github.com/jercle/cloudini/lib"
)

func SearchResponseUserTransform(sr *ldap.SearchResult) (users []ADUser) {
	jsonStr, _ := json.Marshal(sr)

	var rspData SearchResponse
	err := json.Unmarshal(jsonStr, &rspData)
	lib.CheckFatalError(err)

	for _, u := range rspData.Entries {
		var user ADUser
		for _, attr := range u.Attributes {
			switch attr.Name {
			case "givenName":
				user.GivenName = attr.Values[0]
			case "initials":
				user.Initials = attr.Values[0]
			case "distinguishedName":
				user.DistinguishedName = attr.Values[0]
			case "whenCreated":
				user.WhenCreated = attr.Values[0]
			case "whenChanged":
				user.WhenChanged = attr.Values[0]
			case "displayName":
				user.DisplayName = attr.Values[0]
			case "memberOf":
				user.MemberOf = attr.Values
			case "name":
				user.Name = attr.Values[0]
			case "objectGUID":
				user.ObjectGUID = attr.Values[0]
			case "lastLogon":
				user.LastLogon = attr.Values[0]
			case "pwdLastSet":
				user.PwdLastSet = attr.Values[0]
			case "objectSid":
				user.ObjectSid = attr.Values[0]
			case "accountExpires":
				user.AccountExpires = attr.Values[0]
			case "sAMAccountName":
				user.SAMAccountName = attr.Values[0]
			case "userPrincipalName":
				user.UserPrincipalName = attr.Values[0]
			case "lockoutTime":
				user.LockoutTime = attr.Values[0]
			case "lastLogonTimestamp":
				user.LastLogonTimestamp = attr.Values[0]
			case "msDS-xternalDirectoryObjectId":
				user.MsDSExternalDirectoryObjectId = attr.Values[0]
			case "mail":
				user.Mail = attr.Values[0]
			case "mailNickname":
				user.MailNickname = attr.Values[0]
			case "description":
				user.Description = attr.Values[0]
			case "cn":
				user.CN = attr.Values[0]
			case "objectClass":
				user.ObjectClass = attr.Values
			}
		}
		users = append(users, user)
	}

	return
}
