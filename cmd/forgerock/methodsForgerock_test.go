package forgerock

import (
	json "encoding/json/v2"
	"reflect"
	"testing"
)

func TestLDAPConnectorGroupMembersUnmarshalJSON(t *testing.T) {
	data := `[
		"CN=jsmith,OU=Users,DC=corp,DC=example,DC=com",
		"CN=adoe,OU=Admins,DC=corp,DC=example,DC=com"
	]`

	var members LDAPConnectorGroupMembers
	if err := json.Unmarshal([]byte(data), &members); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := LDAPConnectorGroupMembers{
		"jsmith": "CN=jsmith,OU=Users,DC=corp,DC=example,DC=com",
		"adoe":   "CN=adoe,OU=Admins,DC=corp,DC=example,DC=com",
	}
	if !reflect.DeepEqual(members, want) {
		t.Fatalf("got %v, want %v", members, want)
	}
}

func TestLDAPConnectorGroupMembersUnmarshalJSONEmptyValues(t *testing.T) {
	for _, data := range []string{`null`, `""`} {
		t.Run(data, func(t *testing.T) {
			var members LDAPConnectorGroupMembers
			if err := json.Unmarshal([]byte(data), &members); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(members) != 0 {
				t.Fatalf("expected no members, got %v", members)
			}
		})
	}

	var members LDAPConnectorGroupMembers
	if err := json.Unmarshal([]byte(`[]`), &members); err != nil {
		t.Fatalf("unexpected error for an empty array: %v", err)
	}
	if members == nil || len(members) != 0 {
		t.Fatalf("expected an empty, non-nil map for an empty array, got %#v", members)
	}
}

func TestLDAPConnectorGroupMembersUnmarshalJSONOnlyStripsLeadingCN(t *testing.T) {
	// Only the first CN= is removed, and only the first DN component is used as the name
	data := `["CN=svc-CN=odd,CN=Managed Service Accounts,DC=corp,DC=example,DC=com", "no-cn-prefix,DC=corp"]`

	var members LDAPConnectorGroupMembers
	if err := json.Unmarshal([]byte(data), &members); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := LDAPConnectorGroupMembers{
		"svc-CN=odd":   "CN=svc-CN=odd,CN=Managed Service Accounts,DC=corp,DC=example,DC=com",
		"no-cn-prefix": "no-cn-prefix,DC=corp",
	}
	if !reflect.DeepEqual(members, want) {
		t.Fatalf("got %v, want %v", members, want)
	}
}

func TestLDAPConnectorGroupMembersUnmarshalJSONInvalid(t *testing.T) {
	for _, data := range []string{`{"member":"CN=a"}`, `"CN=a,DC=b"`, `[1,2,3]`, `42`} {
		t.Run(data, func(t *testing.T) {
			var members LDAPConnectorGroupMembers
			if err := json.Unmarshal([]byte(data), &members); err == nil {
				t.Fatalf("expected an error, got %v", members)
			}
		})
	}
}

func TestLDAPConnectorGroupMembersUnmarshalJSONAsStructField(t *testing.T) {
	type group struct {
		Name    string                    `json:"name"`
		Members LDAPConnectorGroupMembers `json:"members"`
		Empty   LDAPConnectorGroupMembers `json:"empty"`
	}
	data := `{"name":"Domain Admins","members":["CN=adoe,OU=Admins,DC=corp"],"empty":null}`

	var got group
	if err := json.Unmarshal([]byte(data), &got); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "Domain Admins" {
		t.Fatalf("unexpected name %q", got.Name)
	}
	if got.Members["adoe"] != "CN=adoe,OU=Admins,DC=corp" || len(got.Members) != 1 {
		t.Fatalf("unexpected members %v", got.Members)
	}
	if len(got.Empty) != 0 {
		t.Fatalf("expected no members for null, got %v", got.Empty)
	}
}
