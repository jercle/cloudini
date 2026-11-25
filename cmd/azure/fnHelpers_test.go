package azure

import (
	json "encoding/json/v2"
	"fmt"
	"net"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/jercle/cloudini/lib"
)

//
// Auth scopes
//

func TestGetAuthScopeString(t *testing.T) {
	tests := []struct{ scope, want string }{
		{"graph", "https://graph.microsoft.com/.default"},
		{"storage", "https://storage.azure.com/.default"},
		{"monitor", "https://monitor.azure.com//.default"},
		{"defender", "https://api.securitycenter.microsoft.com/.default"},
		{"loganalytics", "https://api.loganalytics.io/.default"},
		{"keyvault", "cfa8b339-82a2-471a-a3c9-0fc0be7a4093/.default"},
		{"m365", "https://outlook.office365.com/.default"},
		// Anything else falls back to Azure Resource Manager
		{"management", "https://management.core.windows.net/.default"},
		{"", "https://management.core.windows.net/.default"},
		{"Graph", "https://management.core.windows.net/.default"},
		{"unknown", "https://management.core.windows.net/.default"},
	}
	for _, tc := range tests {
		t.Run(tc.scope, func(t *testing.T) {
			got := GetAuthScopeString(tc.scope)
			if len(got) != 1 || got[0] != tc.want {
				t.Fatalf("got %v, want [%s]", got, tc.want)
			}
		})
	}
}

//
// IP ranges and CIDRs
//

func TestIpRangeToCidr(t *testing.T) {
	tests := []struct {
		name       string
		start, end string
		want       []string
	}{
		{"aligned /24", "10.0.0.0", "10.0.0.255", []string{"10.0.0.0/24"}},
		{"single address", "10.0.0.5", "10.0.0.5", []string{"10.0.0.5/32"}},
		{"two aligned addresses", "10.0.0.4", "10.0.0.5", []string{"10.0.0.4/31"}},
		{"unaligned range", "10.0.0.1", "10.0.0.6", []string{"10.0.0.1/32", "10.0.0.2/31", "10.0.0.4/31", "10.0.0.6/32"}},
		{"range crossing an octet", "10.0.0.128", "10.0.1.127", []string{"10.0.0.128/25", "10.0.1.0/25"}},
		{"two adjacent /24s that do not merge", "10.0.1.0", "10.0.2.255", []string{"10.0.1.0/24", "10.0.2.0/24"}},
		{"aligned /16", "172.16.0.0", "172.16.255.255", []string{"172.16.0.0/16"}},
		{"ipv6 single address", "2001:db8::1", "2001:db8::1", []string{"2001:db8::1/128"}},
		{"ipv6 aligned range", "2001:db8::", "2001:db8::ff", []string{"2001:db8::/120"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := IpRangeToCidr(tc.start, tc.end)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

// The CIDRs returned for a range should cover every address in it exactly
// once, and nothing outside it.
func TestIpRangeToCidrCoversRangeExactly(t *testing.T) {
	ranges := [][2]string{
		{"10.0.0.3", "10.0.0.200"},
		{"192.168.0.250", "192.168.2.5"},
		{"10.1.255.255", "10.2.0.0"},
	}
	for _, r := range ranges {
		cidrs, err := IpRangeToCidr(r[0], r[1])
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		var covered []string
		for _, cidr := range cidrs {
			ips, err := GetIPsFromCIDR(cidr)
			if err != nil {
				t.Fatalf("returned an invalid CIDR %q: %v", cidr, err)
			}
			for _, ip := range ips {
				covered = append(covered, ip.String())
			}
		}

		var want []string
		end := net.ParseIP(r[1]).To4()
		for ip := net.ParseIP(r[0]).To4(); ; IncrementIp(ip) {
			want = append(want, ip.String())
			if ip.Equal(end) {
				break
			}
		}

		if !reflect.DeepEqual(covered, want) {
			t.Fatalf("%s - %s: CIDRs %v cover %d addresses, want %d", r[0], r[1], cidrs, len(covered), len(want))
		}
	}
}

func TestIpRangeToCidrErrors(t *testing.T) {
	tests := []struct {
		name, start, end string
		wantErr          string
	}{
		{"invalid start", "nope", "10.0.0.1", ""},
		{"invalid end", "10.0.0.1", "nope", ""},
		{"start after end", "10.0.0.10", "10.0.0.1", "start > end"},
		{"mixed families", "10.0.0.1", "2001:db8::1", "types are different"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := IpRangeToCidr(tc.start, tc.end)
			if err == nil {
				t.Fatalf("expected an error, got %v", got)
			}
			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("expected error containing %q, got %q", tc.wantErr, err.Error())
			}
		})
	}
}

func ipStrings(ips []net.IP) []string {
	out := []string{}
	for _, ip := range ips {
		out = append(out, ip.String())
	}
	return out
}

func TestGetIPsFromCIDR(t *testing.T) {
	tests := []struct {
		cidr string
		want []string
	}{
		{"192.168.1.0/30", []string{"192.168.1.0", "192.168.1.1", "192.168.1.2", "192.168.1.3"}},
		{"10.0.0.7/32", []string{"10.0.0.7"}},
		{"10.0.0.254/31", []string{"10.0.0.254", "10.0.0.255"}},
		// A host address is masked back to the start of its network
		{"10.0.0.6/30", []string{"10.0.0.4", "10.0.0.5", "10.0.0.6", "10.0.0.7"}},
		{"2001:db8::/127", []string{"2001:db8::", "2001:db8::1"}},
	}
	for _, tc := range tests {
		t.Run(tc.cidr, func(t *testing.T) {
			got, err := GetIPsFromCIDR(tc.cidr)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(ipStrings(got), tc.want) {
				t.Fatalf("got %v, want %v", ipStrings(got), tc.want)
			}
		})
	}
}

func TestGetIPsFromCIDRIncludesNetworkAndBroadcast(t *testing.T) {
	got, err := GetIPsFromCIDR("10.20.30.0/24")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 256 {
		t.Fatalf("expected all 256 addresses of a /24, got %d", len(got))
	}
	if got[0].String() != "10.20.30.0" || got[255].String() != "10.20.30.255" {
		t.Fatalf("unexpected first and last addresses: %s, %s", got[0], got[255])
	}

	// Each entry is its own slice, not a view of one shared buffer
	seen := map[string]bool{}
	for _, ip := range got {
		seen[ip.String()] = true
	}
	if len(seen) != 256 {
		t.Fatalf("expected 256 distinct addresses, got %d", len(seen))
	}
}

func TestGetIPsFromCIDRInvalid(t *testing.T) {
	for _, cidr := range []string{"", "10.0.0.1", "10.0.0.0/33", "not-a-cidr"} {
		got, err := GetIPsFromCIDR(cidr)
		if err == nil {
			t.Fatalf("%q: expected an error, got %v", cidr, got)
		}
		if !strings.Contains(err.Error(), "invalid CIDR") {
			t.Fatalf("%q: unexpected error %q", cidr, err.Error())
		}
	}
}

func TestIncrementIp(t *testing.T) {
	tests := []struct{ in, want string }{
		{"10.0.0.1", "10.0.0.2"},
		{"10.0.0.255", "10.0.1.0"},
		{"10.0.255.255", "10.1.0.0"},
		{"255.255.255.255", "0.0.0.0"},
		{"2001:db8::ffff", "2001:db8::1:0"},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			ip := net.ParseIP(tc.in)
			if v4 := ip.To4(); v4 != nil {
				ip = v4
			}
			IncrementIp(ip) // increments in place
			if ip.String() != tc.want {
				t.Fatalf("got %s, want %s", ip, tc.want)
			}
		})
	}
}

//
// Resource graph queries for IP addresses
//

func TestGetIPAddressesQueries(t *testing.T) {
	all := GetIPAddressesQueries(nil)
	if len(all) == 0 {
		t.Fatal("expected queries to be returned")
	}
	for name, query := range all {
		if !strings.HasPrefix(name, "GetIPAddressesQuery") {
			t.Fatalf("unexpected query name %q", name)
		}
		if strings.TrimSpace(query) == "" {
			t.Fatalf("query %q is empty", name)
		}
		if !strings.Contains(strings.ToLower(query), "resources") {
			t.Fatalf("query %q does not look like a resource graph query", name)
		}
	}
	for _, name := range []string{
		"GetIPAddressesQueryVirtualMachines",
		"GetIPAddressesQueryVirtualNetworks",
		"GetIPAddressesQueryWebSites",
	} {
		if _, exists := all[name]; !exists {
			t.Fatalf("expected query %q to exist", name)
		}
	}
}

func TestGetIPAddressesQueriesSelection(t *testing.T) {
	all := GetIPAddressesQueries(nil)

	selected := []string{"GetIPAddressesQueryVirtualMachines", "GetIPAddressesQueryVirtualNetworks", "DoesNotExist"}
	got := GetIPAddressesQueries(&selected)

	if len(got) != 2 {
		t.Fatalf("expected the 2 known queries, got %d", len(got))
	}
	for _, name := range selected[:2] {
		if got[name] != all[name] {
			t.Fatalf("selected query %q does not match the full set", name)
		}
	}

	none := []string{}
	if got := GetIPAddressesQueries(&none); len(got) != 0 {
		t.Fatalf("expected no queries for an empty selection, got %d", len(got))
	}
}

func TestIPAddressesAllResourceTypesUnmarshalJSON(t *testing.T) {
	data := `{
		"id": "/subscriptions/sub-1/resourceGroups/rg-net/providers/Microsoft.Network/networkInterfaces/nic-01",
		"name": "nic-01",
		"type": "microsoft.network/networkinterfaces",
		"resourceGroup": "rg-net",
		"subscriptionId": "sub-1",
		"tenantId": "tenant-1",
		"allIpAddresses": ["10.0.0.4", "20.1.2.3"],
		"privateIps": ["10.0.0.4"],
		"publicIps": ["20.1.2.3"],
		"cidrs": ["10.0.0.0/24", "10.0.1.5", "10.0.2.9/24", "not-an-ip"],
		"isAttached": 1,
		"attachedTo": "/subscriptions/sub-1/resourceGroups/rg-app/providers/Microsoft.Compute/virtualMachines/vm-01",
		"tags": "{\"env\":\"prod\",\"owner\":\"platform\"}"
	}`

	var got IPAddressesAllResourceTypes
	if err := got.UnmarshalJSON([]byte(data)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.Name != "nic-01" || got.ResourceGroup != "rg-net" || got.SubscriptionID != "sub-1" || got.TenantID != "tenant-1" {
		t.Fatalf("unexpected identity fields: %+v", got)
	}
	if got.Type != "microsoft.network/networkinterfaces" || !strings.HasSuffix(got.ID, "/nic-01") {
		t.Fatalf("unexpected type or id: %q, %q", got.Type, got.ID)
	}
	if !reflect.DeepEqual(got.AllIpAddresses, []string{"10.0.0.4", "20.1.2.3"}) {
		t.Fatalf("unexpected allIpAddresses %v", got.AllIpAddresses)
	}
	if !reflect.DeepEqual(got.PrivateIps, []string{"10.0.0.4"}) || !reflect.DeepEqual(got.PublicIps, []string{"20.1.2.3"}) {
		t.Fatalf("unexpected private or public IPs: %v, %v", got.PrivateIps, got.PublicIps)
	}
	// The resource graph returns isAttached as a number
	if !got.IsAttached {
		t.Fatal("expected isAttached 1 to become true")
	}
	if !strings.HasSuffix(got.AttachedTo, "/vm-01") {
		t.Fatalf("unexpected attachedTo %q", got.AttachedTo)
	}
	// Tags arrive as a JSON string and are expanded into a map
	if want := map[string]string{"env": "prod", "owner": "platform"}; !reflect.DeepEqual(got.Tags, want) {
		t.Fatalf("got tags %v, want %v", got.Tags, want)
	}
	// CIDRs are normalised to their network, bare addresses become /32 and junk is dropped
	if want := []string{"10.0.0.0/24", "10.0.1.5/32", "10.0.2.0/24"}; !reflect.DeepEqual(got.Cidrs, want) {
		t.Fatalf("got cidrs %v, want %v", got.Cidrs, want)
	}
}

func TestIPAddressesAllResourceTypesUnmarshalJSONNotAttachedAndNoTags(t *testing.T) {
	var got IPAddressesAllResourceTypes
	if err := got.UnmarshalJSON([]byte(`{"name": "pip-unused", "isAttached": 0, "tags": ""}`)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "pip-unused" {
		t.Fatalf("unexpected name %q", got.Name)
	}
	if got.IsAttached {
		t.Fatal("expected isAttached 0 to become false")
	}
	if len(got.Tags) != 0 || len(got.Cidrs) != 0 {
		t.Fatalf("expected no tags or cidrs, got %v and %v", got.Tags, got.Cidrs)
	}
}

func TestIPAddressesAllResourceTypesUnmarshalJSONNullLeavesValueAlone(t *testing.T) {
	for _, data := range []string{`null`, `""`} {
		got := IPAddressesAllResourceTypes{Name: "existing"}
		if err := got.UnmarshalJSON([]byte(data)); err != nil {
			t.Fatalf("%s: unexpected error: %v", data, err)
		}
		if got.Name != "existing" {
			t.Fatalf("%s: expected value to be left alone, got %+v", data, got)
		}
	}
}

func TestIPAddressesAllResourceTypesUnmarshalJSONErrors(t *testing.T) {
	for name, data := range map[string]string{
		"not json":               `{"name": `,
		"tags not a json string": `{"name": "nic", "tags": "this is not json"}`,
	} {
		t.Run(name, func(t *testing.T) {
			var got IPAddressesAllResourceTypes
			if err := got.UnmarshalJSON([]byte(data)); err == nil {
				t.Fatalf("expected an error, got %+v", got)
			}
		})
	}
}

//
// Cost exports
//

func TestSortCostPerDay(t *testing.T) {
	in := lib.CostPerDay{
		"2025-01-03": 3.5,
		"2025-01-01": 1.25,
		"2025-01-02": 0,
	}

	got := SortCostPerDay(in)

	if len(got) != 3 || got["2025-01-01"] != 1.25 || got["2025-01-02"] != 0 || got["2025-01-03"] != 3.5 {
		t.Fatalf("expected the same days and costs, got %v", got)
	}

	// The result is a copy
	got["2025-01-04"] = 9
	if _, exists := in["2025-01-04"]; exists {
		t.Fatal("changing the result should not change the input")
	}

	if got := SortCostPerDay(nil); len(got) != 0 {
		t.Fatalf("expected an empty result for nil input, got %v", got)
	}
}

//
// Key Vault
//

func TestRemoveDuplicateVnetRules(t *testing.T) {
	const subnetA = "/subscriptions/s/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet/subnets/a"
	const subnetB = "/subscriptions/s/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet/subnets/b"

	got := removeDuplicateVnetRules([]KeyVaultNetworkACLVirtualNetworkRule{
		{ID: subnetA},
		{ID: subnetB, IgnoreMissingVnetServiceEndpoint: true},
		{ID: strings.ToUpper(subnetA), IgnoreMissingVnetServiceEndpoint: true},
		{ID: subnetB},
	})

	// Azure resource IDs are case insensitive. The first of each is kept, as it was given
	if len(got) != 2 {
		t.Fatalf("expected 2 unique rules, got %+v", got)
	}
	if got[0].ID != subnetA || got[0].IgnoreMissingVnetServiceEndpoint {
		t.Fatalf("unexpected first rule %+v", got[0])
	}
	if got[1].ID != subnetB || !got[1].IgnoreMissingVnetServiceEndpoint {
		t.Fatalf("unexpected second rule %+v", got[1])
	}

	if got := removeDuplicateVnetRules(nil); got == nil || len(got) != 0 {
		t.Fatalf("expected an empty, non-nil slice for nil input, got %#v", got)
	}
}

func TestRemoveDuplicateAccessPolicies(t *testing.T) {
	got := removeDuplicateAccessPolicies([]KeyVaultAccessPolicy{
		{ObjectID: "aaaaaaaa-0000-0000-0000-000000000001", TenantID: "first"},
		{ObjectID: "bbbbbbbb-0000-0000-0000-000000000002", TenantID: "first"},
		{ObjectID: "AAAAAAAA-0000-0000-0000-000000000001", TenantID: "second"},
	})

	if len(got) != 2 {
		t.Fatalf("expected 2 unique access policies, got %+v", got)
	}
	if got[0].ObjectID != "aaaaaaaa-0000-0000-0000-000000000001" || got[0].TenantID != "first" {
		t.Fatalf("expected the first policy for a repeated object id to be kept, got %+v", got[0])
	}
	if got[1].ObjectID != "bbbbbbbb-0000-0000-0000-000000000002" {
		t.Fatalf("unexpected second policy %+v", got[1])
	}

	if got := removeDuplicateAccessPolicies(nil); got == nil || len(got) != 0 {
		t.Fatalf("expected an empty, non-nil slice for nil input, got %#v", got)
	}
}

//
// NSG flow logs
//

// Flow tuple fields: time, source IP, dest IP, source port, dest port, protocol, direction, decision
func flowLogRecords(t *testing.T) []NsgFlowLogRecord {
	t.Helper()
	data := `[
		{
			"resourceId": "nsg-1",
			"properties": {"Version": 2, "flows": [
				{"rule": "AllowHttps", "flows": [
					{"mac": "000D3AF87856", "flowTuples": [
						"1700000000,10.0.0.4,20.1.1.1,50000,443,T,O,A",
						"1700000001,10.0.0.5,20.1.1.1,50001,443,T,O,A"
					]}
				]},
				{"rule": "DenyAll", "flows": [
					{"mac": "000D3AF87856", "flowTuples": ["1700000002,10.0.0.4,20.2.2.2,50002,22,T,O,D"]}
				]}
			]}
		},
		{
			"resourceId": "nsg-2",
			"properties": {"Version": 2, "flows": [
				{"rule": "AllowDns", "flows": [
					{"mac": "000D3AF87857", "flowTuples": ["1700000003,10.0.9.9,20.3.3.3,50003,53,U,O,A"]}
				]}
			]}
		},
		{
			"resourceId": "nsg-3",
			"properties": {"Version": 2, "flows": []}
		}
	]`
	var records []NsgFlowLogRecord
	if err := json.Unmarshal([]byte(data), &records); err != nil {
		t.Fatal(err)
	}
	return records
}

func TestGetUniqueIpAddresses(t *testing.T) {
	got := GetUniqueIpAddresses(flowLogRecords(t))

	if want := []string{"10.0.0.4", "10.0.0.5", "10.0.9.9"}; !reflect.DeepEqual(got.SourceIps, want) {
		t.Fatalf("source IPs: got %v, want %v", got.SourceIps, want)
	}
	if want := []string{"20.1.1.1", "20.2.2.2", "20.3.3.3"}; !reflect.DeepEqual(got.DestIps, want) {
		t.Fatalf("dest IPs: got %v, want %v", got.DestIps, want)
	}
}

func TestGetUniqueIpAddressesNoRecords(t *testing.T) {
	got := GetUniqueIpAddresses(nil)
	if len(got.SourceIps) != 0 || len(got.DestIps) != 0 {
		t.Fatalf("expected no addresses, got %+v", got)
	}
}

func flowLogResourceIds(logs CombinedFlowLogs) []string {
	ids := []string{}
	for _, record := range logs.NsgFlowLogs {
		ids = append(ids, record.ResourceID)
	}
	return ids
}

func TestCombinedFlowLogsFilterIp(t *testing.T) {
	tests := []struct {
		name      string
		filter    string
		direction string
		want      []string
	}{
		{"source in one record", "10.0.9.9", "source", []string{"nsg-2"}},
		{"source matching several tuples adds the record once", "10.0.0.4", "source", []string{"nsg-1"}},
		{"several source addresses", "10.0.0.5,10.0.9.9", "source", []string{"nsg-1", "nsg-2"}},
		{"direction is case insensitive", "10.0.9.9", "SOURCE", []string{"nsg-2"}},
		{"dest", "20.2.2.2", "dest", []string{"nsg-1"}},
		{"dest address used as source filter", "20.2.2.2", "source", []string{}},
		{"no match", "192.168.0.1", "source", []string{}},
		{"partial addresses do not match", "10.0.0", "source", []string{}},
		{"unknown direction matches nothing", "10.0.0.4", "sideways", []string{}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			logs := CombinedFlowLogs{NsgFlowLogs: flowLogRecords(t), FileCount: 7}

			logs.FilterIp(tc.filter, tc.direction)

			if got := flowLogResourceIds(logs); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			if logs.FileCount != 7 {
				t.Fatalf("expected file count to be kept, got %d", logs.FileCount)
			}
		})
	}
}

func TestIpListFilterSourceIp(t *testing.T) {
	list := IpList{
		SourceIps: []string{"10.0.0.4", "10.0.0.5", "10.0.9.9"},
		DestIps:   []string{"20.1.1.1"},
	}

	list.FilterSourceIp("10.0.9.9,10.0.0.4,192.168.0.1")

	if want := []string{"10.0.0.4", "10.0.9.9"}; !reflect.DeepEqual(list.SourceIps, want) {
		t.Fatalf("got %v, want %v", list.SourceIps, want)
	}
	// Filtering by source drops the destination list
	if len(list.DestIps) != 0 {
		t.Fatalf("expected dest IPs to be cleared, got %v", list.DestIps)
	}

	list.FilterSourceIp("nothing-matches")
	if len(list.SourceIps) != 0 {
		t.Fatalf("expected no source IPs, got %v", list.SourceIps)
	}
}

//
// Log Analytics tables
//

func logAnalyticsTables(t *testing.T) LogAnalyticsTables {
	t.Helper()
	data := `[
		{"name": "AzureActivity", "properties": {"retentionInDays": 90, "retentionInDaysAsDefault": true}},
		{"name": "SecurityEvent", "properties": {"retentionInDays": 365, "retentionInDaysAsDefault": false}},
		{"name": "AzureDiagnostics", "properties": {"retentionInDays": 30, "retentionInDaysAsDefault": true}},
		{"name": "Heartbeat", "properties": {"retentionInDays": 30, "retentionInDaysAsDefault": false}}
	]`
	var tables LogAnalyticsTables
	if err := json.Unmarshal([]byte(data), &tables); err != nil {
		t.Fatal(err)
	}
	return tables
}

func tableNames(tables LogAnalyticsTables) []string {
	names := []string{}
	for _, table := range tables {
		names = append(names, table.Name)
	}
	return names
}

func TestLogAnalyticsTablesFilterByName(t *testing.T) {
	tests := []struct {
		name            string
		filter          string
		caseInsensitive bool
		want            []string
	}{
		{"case sensitive match", "Azure", false, []string{"AzureActivity", "AzureDiagnostics"}},
		{"case sensitive miss", "azure", false, []string{}},
		{"case insensitive match", "azure", true, []string{"AzureActivity", "AzureDiagnostics"}},
		{"case insensitive with upper case filter", "HEART", true, []string{"Heartbeat"}},
		{"substring in the middle", "Event", false, []string{"SecurityEvent"}},
		{"empty filter keeps everything", "", false, []string{"AzureActivity", "SecurityEvent", "AzureDiagnostics", "Heartbeat"}},
		{"no match", "Syslog", true, []string{}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tables := logAnalyticsTables(t)
			tables.filterByName(tc.filter, tc.caseInsensitive)
			if got := tableNames(tables); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestLogAnalyticsTablesFilterByRetentionInDaysAsDefault(t *testing.T) {
	usingDefault := logAnalyticsTables(t)
	usingDefault.filterByRetentionInDaysAsDefault(true)
	if got, want := tableNames(usingDefault), []string{"AzureActivity", "AzureDiagnostics"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("default retention: got %v, want %v", got, want)
	}

	custom := logAnalyticsTables(t)
	custom.filterByRetentionInDaysAsDefault(false)
	if got, want := tableNames(custom), []string{"SecurityEvent", "Heartbeat"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("custom retention: got %v, want %v", got, want)
	}

	// Filters stack
	custom.filterByName("Heart", false)
	if got, want := tableNames(custom), []string{"Heartbeat"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("stacked filters: got %v, want %v", got, want)
	}
}

//
// Resource SKUs
//

func resourceSkus(t *testing.T) []lib.AzureResourceSku {
	t.Helper()
	data := `[
		{"name": "Standard_D4s_v5", "resourceType": "virtualMachines", "capabilities": [
			{"name": "vCPUs", "value": "4"},
			{"name": "vCPUsAvailable", "value": "4"},
			{"name": "vCPUsPerCore", "value": "2"},
			{"name": "MemoryGB", "value": "16"}
		]},
		{"name": "Standard_LRS", "resourceType": "disks", "capabilities": [
			{"name": "MaxSizeGiB", "value": "32767"}
		]},
		{"name": "Standard_HB120rs_v3", "resourceType": "virtualMachines", "capabilities": [
			{"name": "Cores", "value": "120"}
		]},
		{"name": "NoCapabilities", "resourceType": "snapshots"}
	]`
	var skus []lib.AzureResourceSku
	if err := json.Unmarshal([]byte(data), &skus); err != nil {
		t.Fatal(err)
	}
	return skus
}

func TestGetSkusWithVcpus(t *testing.T) {
	got := GetSkusWithVcpus(resourceSkus(t))

	// Only SKUs with at least one vCPU or core capability are kept
	if len(got) != 2 {
		t.Fatalf("expected 2 SKUs with vCPU details, got %d: %+v", len(got), got)
	}

	d4 := got[0]
	if d4.Name != "Standard_D4s_v5" || d4.ResourceType != "virtualMachines" {
		t.Fatalf("unexpected first SKU: %q (%q)", d4.Name, d4.ResourceType)
	}
	if d4.VCPUs != "4" || d4.VCPUsAvailable != "4" || d4.VCPUsPerCore != "2" {
		t.Fatalf("unexpected vCPU details on %s: vCPUs %q, available %q, per core %q", d4.Name, d4.VCPUs, d4.VCPUsAvailable, d4.VCPUsPerCore)
	}

	hb := got[1]
	if hb.Name != "Standard_HB120rs_v3" || hb.Cores != "120" {
		t.Fatalf("unexpected second SKU: %q with cores %q", hb.Name, hb.Cores)
	}
	if hb.VCPUs != "" {
		t.Fatalf("expected no vCPUs value on %s, got %q", hb.Name, hb.VCPUs)
	}
}

func TestGetSkusWithVcpusNoInput(t *testing.T) {
	if got := GetSkusWithVcpus(nil); len(got) != 0 {
		t.Fatalf("expected no SKUs, got %+v", got)
	}
}

func TestMapVmSizeDetails(t *testing.T) {
	skus := []lib.AzureVirtualMachineSku{
		{Name: "Standard_D2s_v5", VCPUs: "2"},
		{Name: "Standard_D4s_v5", VCPUs: "4"},
		{Name: "Standard_D8s_v5", VCPUs: "8"},
	}

	if got := MapVmSizeDetails("Standard_D4s_v5", skus); got.Name != "Standard_D4s_v5" || got.VCPUs != "4" {
		t.Fatalf("unexpected SKU: %q with vCPUs %q", got.Name, got.VCPUs)
	}
	// Names are matched exactly
	if got := MapVmSizeDetails("standard_d4s_v5", skus); got.Name != "" {
		t.Fatalf("expected no match for different casing, got %q", got.Name)
	}
	if got := MapVmSizeDetails("Standard_E96s_v5", skus); got.Name != "" || got.VCPUs != "" {
		t.Fatalf("expected an empty SKU for an unknown size, got %q", got.Name)
	}
	if got := MapVmSizeDetails("Standard_D4s_v5", nil); got.Name != "" {
		t.Fatalf("expected an empty SKU with no SKUs to search, got %q", got.Name)
	}
}

func TestMapVmSizeDetailsLastMatchWins(t *testing.T) {
	skus := []lib.AzureVirtualMachineSku{
		{Name: "Standard_D4s_v5", VCPUs: "4", Size: "first"},
		{Name: "Standard_D4s_v5", VCPUs: "4", Size: "second"},
	}
	if got := MapVmSizeDetails("Standard_D4s_v5", skus); got.Size != "second" {
		t.Fatalf("expected the last matching SKU, got %q", got.Size)
	}
}

//
// Subscription picker list items
//

func TestSubscriptionListItem(t *testing.T) {
	i := item{
		title:      "Production",
		subId:      "sub-123",
		subName:    "Production",
		tenantId:   "tenant-456",
		tenantName: "CONTOSO",
		isDefault:  true,
	}

	if got := i.Title(); got != "Production" {
		t.Fatalf("unexpected title %q", got)
	}
	if got, want := i.Description(), "Sub: sub-123(Production) - Tenant: tenant-456 (CONTOSO)"; got != want {
		t.Fatalf("got description %q, want %q", got, want)
	}
	// Filtering searches the title, subscription id and tenant name
	filter := i.FilterValue()
	for _, part := range []string{"Production", "sub-123", "CONTOSO"} {
		if !strings.Contains(filter, part) {
			t.Fatalf("expected filter value %q to contain %q", filter, part)
		}
	}
	if !i.IsDefault() {
		t.Fatal("expected item to be the default")
	}
	if (item{}).IsDefault() {
		t.Fatal("expected a zero item not to be the default")
	}
}

//
// Token freshness on multi auth tokens
//

// Token sources stay registered for the life of the test binary, so each test
// run needs its own token name for go test -count=N to behave.
func uniqueTokenName(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, time.Now().UnixNano())
}

func TestFreshTokenUnregisteredTokenUnchanged(t *testing.T) {
	mat := lib.AzureMultiAuthToken{
		TenantId:   "tenant-id",
		TenantName: "PROD",
		TokenData:  lib.AzureTokenData{Token: "fresh-token-unregistered", ExpiresOn: time.Now().Add(-time.Hour)},
	}
	if got := FreshToken(mat); got != mat {
		t.Fatalf("expected an unregistered token to be returned as is, got %+v", got)
	}
}

func TestFreshTokenRefreshesExpiredTokenAndKeepsTenantDetails(t *testing.T) {
	name := uniqueTokenName("fresh-token-expired")
	original, refreshes := registerTestToken(name, time.Now().Add(-time.Minute))
	mat := lib.AzureMultiAuthToken{
		TenantId:     "tenant-id",
		ClientId:     "client-id",
		TenantName:   "PROD",
		TenantDomain: "prod.example.com",
		TokenData:    original,
	}

	got := FreshToken(mat)

	if got.TokenData.Token != name+"-1" {
		t.Fatalf("expected the refreshed token, got %q", got.TokenData.Token)
	}
	if !got.TokenData.ExpiresOn.After(time.Now()) {
		t.Fatalf("expected the refreshed token to expire in the future, got %v", got.TokenData.ExpiresOn)
	}
	if got.TenantId != "tenant-id" || got.ClientId != "client-id" || got.TenantName != "PROD" || got.TenantDomain != "prod.example.com" {
		t.Fatalf("expected tenant details to be kept, got %+v", got)
	}
	if mat.TokenData.Token != name {
		t.Fatal("the token passed in should not be modified")
	}

	// Asking again with the stale copy gives the same new token without another refresh
	if again := FreshToken(mat); again.TokenData.Token != name+"-1" {
		t.Fatalf("expected the already refreshed token, got %q", again.TokenData.Token)
	}
	if n := refreshes.Load(); n != 1 {
		t.Fatalf("expected a single refresh, got %d", n)
	}
}
