package lib

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestAllTenantTokensSelectTenant(t *testing.T) {
	tokens := AllTenantTokens{
		{TenantName: "PROD", TenantId: "t-prod", TokenData: AzureTokenData{Token: "prod-token"}},
		{TenantName: "DEV", TenantId: "t-dev", TokenData: AzureTokenData{Token: "dev-token"}},
	}

	got, err := tokens.SelectTenant("DEV")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.TenantId != "t-dev" || got.TokenData.Token != "dev-token" {
		t.Fatalf("selected the wrong tenant: %+v", got)
	}

	// Lookup is an exact, case sensitive match
	if _, err := tokens.SelectTenant("dev"); err == nil {
		t.Fatal("expected an error for a tenant name with different casing")
	}
	if got, err := tokens.SelectTenant("MISSING"); err == nil || got != nil {
		t.Fatalf("expected nil token and an error for a missing tenant, got %+v, %v", got, err)
	}
	if _, err := (AllTenantTokens{}).SelectTenant("PROD"); err == nil {
		t.Fatal("expected an error when there are no tokens")
	}
}

func TestAllTenantTokensSelectTenantReturnsFirstMatch(t *testing.T) {
	tokens := AllTenantTokens{
		{TenantName: "DUP", TenantId: "first"},
		{TenantName: "DUP", TenantId: "second"},
	}
	got, err := tokens.SelectTenant("DUP")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.TenantId != "first" {
		t.Fatalf("expected first matching tenant, got %q", got.TenantId)
	}
}

func TestSubsReqResBodyUpdateTenantName(t *testing.T) {
	subs := SubsReqResBody{
		Value: []FetchedSubscription{
			{DisplayName: "sub-a", SubscriptionID: "1"},
			{DisplayName: "sub-b", SubscriptionID: "2", TenantName: "OLD"},
		},
	}
	subs.Count.Type = "Total"
	subs.Count.Value = 2

	subs.UpdateTenantName("NEW")

	if len(subs.Value) != 2 {
		t.Fatalf("expected 2 subscriptions, got %d", len(subs.Value))
	}
	for _, sub := range subs.Value {
		if sub.TenantName != "NEW" {
			t.Fatalf("expected tenant name NEW on %q, got %q", sub.DisplayName, sub.TenantName)
		}
	}
	if subs.Value[0].SubscriptionID != "1" || subs.Value[1].SubscriptionID != "2" {
		t.Fatalf("expected order and other fields to be kept, got %+v", subs.Value)
	}
	if subs.Count.Type != "Total" || subs.Count.Value != 2 {
		t.Fatalf("expected count to be kept, got %+v", subs.Count)
	}
}

func galleryVersions(names ...string) GalleryImageVersionList {
	var list GalleryImageVersionList
	for _, name := range names {
		list.Versions = append(list.Versions, GalleryImageVersion{Name: name, ID: "id-" + name})
	}
	return list
}

func TestGalleryImageVersionListSort(t *testing.T) {
	list := galleryVersions("1.0.10", "1.0.2", "0.9.0", "2.0.0", "1.1.0")

	sorted := list.Sort()

	var got []string
	for _, v := range sorted.Versions {
		got = append(got, v.Name)
	}
	// Semantic, not alphabetical, ordering: 1.0.2 comes before 1.0.10
	want := []string{"0.9.0", "1.0.2", "1.0.10", "1.1.0", "2.0.0"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	if !sorted.Sorted {
		t.Fatal("expected Sorted to be set on the result")
	}
	if list.Sorted {
		t.Fatal("Sort should not modify the receiver")
	}
	if list.Versions[0].Name != "1.0.10" {
		t.Fatalf("Sort should not reorder the receiver, got %v", list.Versions)
	}
	for _, v := range sorted.Versions {
		if v.ID != "id-"+v.Name {
			t.Fatalf("version %q lost its other fields: %+v", v.Name, v)
		}
		if !v.SuffixAdded {
			t.Fatalf("expected SuffixAdded on %q, which has no v prefix", v.Name)
		}
	}
}

func TestGalleryImageVersionListSortMixedPrefixes(t *testing.T) {
	list := galleryVersions("v1.2.0", "1.10.0", "v1.3.0")

	sorted := list.Sort()

	var got []string
	flags := map[string]bool{}
	for _, v := range sorted.Versions {
		got = append(got, v.Name)
		flags[v.Name] = v.SuffixAdded
	}
	// Names keep the form they came in with
	want := []string{"v1.2.0", "v1.3.0", "1.10.0"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	if flags["v1.2.0"] || flags["v1.3.0"] || !flags["1.10.0"] {
		t.Fatalf("unexpected SuffixAdded flags: %v", flags)
	}
}

func TestGalleryImageVersionListLatest(t *testing.T) {
	list := galleryVersions("1.0.10", "1.0.2", "1.0.9")

	latest, name := list.Latest()
	if name != "1.0.10" || latest.Name != "1.0.10" {
		t.Fatalf("expected 1.0.10 as latest, got %q (%+v)", name, latest)
	}
	if latest.ID != "id-1.0.10" {
		t.Fatalf("expected full version to be returned, got %+v", latest)
	}

	single := galleryVersions("3.4.5")
	if _, name := single.Latest(); name != "3.4.5" {
		t.Fatalf("expected only version to be latest, got %q", name)
	}
}

func TestGalleryImageVersionListLatestTrustsSortedFlag(t *testing.T) {
	// When the list says it is already sorted, the last entry is returned as is
	list := galleryVersions("2.0.0", "1.0.0")
	list.Sorted = true
	if _, name := list.Latest(); name != "1.0.0" {
		t.Fatalf("expected last entry of a pre-sorted list, got %q", name)
	}
}

func TestGalleryImageVersionListListVersions(t *testing.T) {
	list := galleryVersions("1.0.10", "1.0.2", "1.0.1")
	got := list.ListVersions()
	want := []string{"1.0.1", "1.0.2", "1.0.10"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	empty := GalleryImageVersionList{}
	if got := empty.ListVersions(); len(got) != 0 {
		t.Fatalf("expected no versions for an empty list, got %v", got)
	}
}

func TestGalleryImageVersionListCheckVersionExists(t *testing.T) {
	list := galleryVersions("1.0.0", "1.0.1", "v2.0.0")

	if !list.CheckVersionExists("1.0.1") {
		t.Fatal("expected 1.0.1 to exist")
	}
	if !list.CheckVersionExists("v2.0.0") {
		t.Fatal("expected v2.0.0 to exist")
	}
	if list.CheckVersionExists("1.0.2") {
		t.Fatal("expected 1.0.2 not to exist")
	}
	// Comparison is on the name exactly as stored
	if list.CheckVersionExists("2.0.0") {
		t.Fatal("expected 2.0.0 not to match the stored v2.0.0")
	}

	empty := GalleryImageVersionList{}
	if empty.CheckVersionExists("1.0.0") {
		t.Fatal("expected nothing to exist in an empty list")
	}
}

func TestGalleryImageVersionIncrementPatchVersion(t *testing.T) {
	tests := []struct{ in, want string }{
		{"1.2.3", "1.2.4"},
		{"1.2.9", "1.2.10"},
		{"0.0.0", "0.0.1"},
		{"2025.1.99", "2025.1.100"},
		{"v1.2.3", "v1.2.4"},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			v := GalleryImageVersion{Name: tc.in}
			if got := v.IncrementPatchVersion(); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
			if v.Name != tc.in {
				t.Fatalf("receiver should not be modified, got %q", v.Name)
			}
		})
	}
}

func azureConfigWithTenants(tenants CldConfigTenants) AzureConfig {
	var config AzureConfig
	config.MultiTenantAuth.Tenants = tenants
	return config
}

func TestAzureConfigGetDefaultTenant(t *testing.T) {
	config := azureConfigWithTenants(CldConfigTenants{
		"dev":  {TenantName: "DEV", TenantID: "t-dev"},
		"prod": {TenantName: "PROD", TenantID: "t-prod", Default: true},
		"test": {TenantName: "TEST", TenantID: "t-test"},
	})

	got, err := config.GetDefaultTenant()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.TenantName != "PROD" || got.TenantID != "t-prod" {
		t.Fatalf("expected PROD as default, got %+v", got)
	}
}

func TestAzureConfigGetDefaultTenantNoneConfigured(t *testing.T) {
	config := azureConfigWithTenants(CldConfigTenants{
		"dev": {TenantName: "DEV"},
	})
	got, err := config.GetDefaultTenant()
	if err == nil || got != nil {
		t.Fatalf("expected nil tenant and an error, got %+v, %v", got, err)
	}
	if !strings.Contains(err.Error(), "No default Azure tenant") {
		t.Fatalf("unexpected error message %q", err.Error())
	}

	if _, err := (AzureConfig{}).GetDefaultTenant(); err == nil {
		t.Fatal("expected an error when no tenants are configured")
	}
}

func TestTransformedCostItemsByTenantAppendTenantData(t *testing.T) {
	tenants := TransformedCostItemsByTenant{}

	tenants.AppendTenantData(TransformedCostItem{Tenant: "PROD", ResourceGroup: "rg-a", PreTaxCost: 10.5})
	tenants.AppendTenantData(TransformedCostItem{Tenant: "PROD", ResourceGroup: "rg-b", PreTaxCost: 4.5})
	tenants.AppendTenantData(TransformedCostItem{Tenant: "DEV", ResourceGroup: "rg-c", PreTaxCost: 1})

	if len(tenants) != 2 {
		t.Fatalf("expected 2 tenants, got %d", len(tenants))
	}
	prod := tenants["PROD"]
	if prod.PreTaxCost != 15 {
		t.Fatalf("expected PROD cost of 15, got %v", prod.PreTaxCost)
	}
	if len(prod.ResGroups) != 2 || prod.ResGroups[0].ResourceGroup != "rg-a" || prod.ResGroups[1].ResourceGroup != "rg-b" {
		t.Fatalf("unexpected PROD resource groups: %+v", prod.ResGroups)
	}
	dev := tenants["DEV"]
	if dev.PreTaxCost != 1 || len(dev.ResGroups) != 1 {
		t.Fatalf("unexpected DEV data: %+v", dev)
	}
}

func TestCostExportErrorMessages(t *testing.T) {
	mismatch := &FieldMismatch{Expected: 12, Found: 9}
	if got := mismatch.Error(); got != "CSV line fields mismatch. Expected 12 found 9" {
		t.Fatalf("unexpected FieldMismatch message %q", got)
	}
	unsupported := &UnsupportedType{Type: "chan int"}
	if got := unsupported.Error(); got != "Unsupported type: chan int" {
		t.Fatalf("unexpected UnsupportedType message %q", got)
	}

	// Both need to be usable as errors
	var _ error = mismatch
	var _ error = unsupported
}

func blobList(names ...string) BlobList {
	base := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	var list BlobList
	for i, name := range names {
		item := BlobItem{Name: name}
		// Creation dates deliberately do not follow slice order
		item.Properties.CreationTime = base.Add(time.Duration((i*7)%len(names)) * 24 * time.Hour)
		list = append(list, item)
	}
	return list
}

func blobNames(list BlobList) []string {
	names := []string{}
	for _, blob := range list {
		names = append(names, blob.Name)
	}
	return names
}

func TestBlobListFilter(t *testing.T) {
	list := blobList("monthly/202404/a.csv", "daily/202404/b.csv", "monthly/202405/c.csv", "monthly/202404/d.csv")

	list.Filter(BlobListFilterOptions{FilterPrefix: "monthly/202404"})

	want := []string{"monthly/202404/a.csv", "monthly/202404/d.csv"}
	if got := blobNames(list); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestBlobListFilterEdgeCases(t *testing.T) {
	list := blobList("a", "b")
	list.Filter(BlobListFilterOptions{})
	if len(list) != 2 {
		t.Fatalf("expected an empty prefix to keep everything, got %v", blobNames(list))
	}

	list.Filter(BlobListFilterOptions{FilterPrefix: "zzz"})
	if len(list) != 0 {
		t.Fatalf("expected no matches, got %v", blobNames(list))
	}
}

func TestBlobListSortByCreateDate(t *testing.T) {
	day := func(d int) time.Time { return time.Date(2025, time.January, d, 0, 0, 0, 0, time.UTC) }
	newList := func() BlobList {
		list := BlobList{{Name: "middle"}, {Name: "newest"}, {Name: "oldest"}}
		list[0].Properties.CreationTime = day(15)
		list[1].Properties.CreationTime = day(30)
		list[2].Properties.CreationTime = day(1)
		return list
	}

	asc := newList()
	asc.SortByCreateDate("ascending")
	if got, want := blobNames(asc), []string{"oldest", "middle", "newest"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("ascending: got %v, want %v", got, want)
	}

	desc := newList()
	desc.SortByCreateDate("descending")
	if got, want := blobNames(desc), []string{"newest", "middle", "oldest"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("descending: got %v, want %v", got, want)
	}

	// An unknown sort order leaves the list as it was
	unchanged := newList()
	unchanged.SortByCreateDate("sideways")
	if got, want := blobNames(unchanged), []string{"middle", "newest", "oldest"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("invalid order: got %v, want %v", got, want)
	}
}
