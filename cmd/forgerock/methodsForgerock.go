package forgerock

import (
	"encoding/json"
	"strings"
)

func (g *LDAPConnectorGroupMembers) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	var (
		memberStr                []string
		memberSamAccountNames    []string
		memberDistinguishedNames []string
	)

	namesToDN := make(map[string]string)

	err := json.Unmarshal(data, &memberStr)
	if err != nil {
		return err
	}

	for _, member := range memberStr {
		splitStr := strings.Split(member, ",")
		samAccountName := strings.Replace(splitStr[0], "CN=", "", 1)
		memberSamAccountNames = append(memberSamAccountNames, samAccountName)
		memberDistinguishedNames = append(memberDistinguishedNames, member)
		namesToDN[samAccountName] = member
	}

	// *g = LDAPConnectorGroupMember{
	// 	SamAccountNames:    memberSamAccountNames,
	// 	DistinguishedNames: memberDistinguishedNames,
	// 	NamesToDN:          namesToDN,
	// }

	*g = namesToDN

	return nil
}

// func (g *LDAPConnectorGroupMember) MarshalJSON() ([]byte, error) {

// }
