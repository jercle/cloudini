package azure

import (
	"math"
	"os"
	"path/filepath"
	"testing"
)

const costExportTestHeader = "DepartmentName,AccountName,AccountOwnerId,SubscriptionGuid,SubscriptionName,ResourceGroup,ResourceLocation,AvailabilityZone,UsageDateTime,ProductName,MeterCategory,MeterSubcategory,MeterId,MeterName,MeterRegion,UnitOfMeasure,UsageQuantity,ResourceRate,PreTaxCost,CostCenter,ConsumedService,ResourceType,InstanceId,Tags,OfferId,AdditionalInfo,ServiceInfo1,ServiceInfo2,Currency\n"

func writeCostExportTestFile(t *testing.T, dir string, tenant string, rows string) string {
	t.Helper()
	path := filepath.Join(dir, "cost-export__"+tenant+".csv")
	if err := os.WriteFile(path, []byte(costExportTestHeader+rows), 0644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestAggregateCostExportCSVFiles(t *testing.T) {
	dir := t.TempDir()

	vm := "/subscriptions/sub-1/resourceGroups/RG-One/providers/Microsoft.Compute/virtualMachines/VM01"
	fileA := writeCostExportTestFile(t, dir, "TenantA", ""+
		`d,a,o,sub-1,Sub One,RG-One,aue,,2026-10-01,Prod,Virtual Machines,D Series,m,D2 v3,AU,1 Hour,24,0.5,12,cc,Microsoft.Compute,Microsoft.Compute/virtualMachines,`+vm+`,"""env"": ""prod""",o,"{""VCPUs"":2}",,,AUD`+"\n"+
		// Same meter on another day, with different casing
		`d,a,o,sub-1,SUB ONE,rg-one,aue,,2026-10-02,Prod,Virtual Machines,D Series,m,D2 v3,AU,1 Hour,12,0.9,6,cc,Microsoft.Compute,Microsoft.Compute/virtualMachines,`+vm+`,,o,,,,AUD`+"\n"+
		// No resource group
		`d,a,o,sub-1,Sub One,,aue,,2026-10-01,Prod,Support,,m,Standard,AU,1/Month,1,100,100,cc,Microsoft.Support,,,,o,,,,AUD`+"\n")
	// A subscription that is mapped to TenantA, exported from another tenant
	fileB := writeCostExportTestFile(t, dir, "TenantB", ""+
		`d,a,o,sub-2,Sub Two,RG-Two,aue,,2026-10-01,Prod,Storage,,m,LRS,AU,1 GB,5,0.1,0.5,cc,Microsoft.Storage,Microsoft.Storage/storageAccounts,/subscriptions/sub-2/resourceGroups/RG-Two/providers/Microsoft.Storage/storageAccounts/sa1,,o,,,,AUD`+"\n")

	data, err := aggregateCostExportCSVFiles([]string{fileA, fileB}, map[string]string{"sub-2": "TenantA"})
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 1 {
		t.Fatalf("expected 1 tenant, got %d", len(data))
	}
	tenant := data["TENANTA"]
	if math.Abs(tenant.MonthTotalCost-118.5) > 1e-9 || tenant.CostPerDay["2026-10-01"] != 112.5 || tenant.CostPerDay["2026-10-02"] != 6 {
		t.Fatalf("unexpected tenant totals: %+v", tenant)
	}
	if len(tenant.Subscriptions) != 2 {
		t.Fatalf("expected 2 subscriptions, got %d", len(tenant.Subscriptions))
	}

	sub := tenant.Subscriptions["sub one"]
	if sub.MonthTotalCost != 118 || len(sub.ResourceGroups) != 2 {
		t.Fatalf("unexpected subscription: %+v", sub)
	}
	if sub.ResourceGroups["nil_res_group"].MonthTotalCost != 100 {
		t.Fatalf("unexpected nil_res_group: %+v", sub.ResourceGroups["nil_res_group"])
	}

	res := sub.ResourceGroups["rg-one"].Resources["vm01"]
	if res.MonthTotalCost != 18 || len(res.MeterData) != 1 {
		t.Fatalf("unexpected resource: %+v", res)
	}

	meterId := "sub-1/rg-one/providers/microsoft.compute/virtualmachines/vm01_virtual machines_d series_d2 v3"
	meter, ok := res.MeterData[meterId]
	if !ok {
		t.Fatalf("meter %q not found in %+v", meterId, res.MeterData)
	}
	if meter.MonthTotalCost != 18 || meter.CostPerDay["2026-10-02"] != 6 || meter.UsageQuantityPerDay["2026-10-01"] != 24 {
		t.Fatalf("unexpected meter totals: %+v", meter)
	}
	// Descriptive fields come from the first row of the meter
	if meter.ResourceRate != 0.5 || meter.Tags["env"] != "prod" || meter.SubscriptionName != "sub one" || meter.ResourceType != "microsoft.compute/virtualmachines" {
		t.Fatalf("unexpected meter details: %+v", meter)
	}
	if info, ok := meter.AdditionalInfo.(map[string]interface{}); !ok || info["VCPUs"] != float64(2) {
		t.Fatalf("unexpected AdditionalInfo: %#v", meter.AdditionalInfo)
	}
}

func TestAggregateCostExportCSVFilesBadRow(t *testing.T) {
	dir := t.TempDir()
	file := writeCostExportTestFile(t, dir, "TenantA",
		`d,a,o,sub-1,Sub One,RG-One,aue,,2026-10-01,Prod,Cat,Sub,m,Name,AU,1 Hour,not-a-number,0.5,12,cc,svc,type,/x,,o,,,,AUD`+"\n")

	if _, err := aggregateCostExportCSVFiles([]string{file}, nil); err == nil {
		t.Fatal("expected an error for a row with an invalid number")
	}
}
