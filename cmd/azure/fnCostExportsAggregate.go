package azure

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/jercle/cloudini/lib"
)

// Column positions of the Azure cost export CSV. These match the field order of
// lib.RowData, which is what the reflection based UnmarshalCostExportCSV relies on.
const (
	costColSubscriptionGuid = 3
	costColSubscriptionName = 4
	costColResourceGroup    = 5
	costColUsageDateTime    = 8
	costColProductName      = 9
	costColMeterCategory    = 10
	costColMeterSubcategory = 11
	costColMeterName        = 13
	costColUnitOfMeasure    = 15
	costColUsageQuantity    = 16
	costColResourceRate     = 17
	costColPreTaxCost       = 18
	costColConsumedService  = 20
	costColResourceType     = 21
	costColInstanceId       = 22
	costColTags             = 23
	costColAdditionalInfo   = 25

	// Number of columns in the export. lib.RowData has one extra field (Datafile).
	costColCount = 29
)

// costAggNode is one level of the tenant > subscription > resource group > resource
// tree. Nodes are held by pointer while aggregating so that adding a row is a
// handful of map updates, instead of copying every parent struct back into its
// parent map as AggregateCostData has to do with the value typed lib maps.
type costAggNode struct {
	costPerDay lib.CostPerDay
	total      float64
	children   map[string]*costAggNode            // tenant, subscription and resource group levels
	meters     map[string]*lib.AggregatedCostItem // resource level only
}

func newCostAggNode(leaf bool) *costAggNode {
	n := &costAggNode{costPerDay: make(lib.CostPerDay)}
	if leaf {
		n.meters = make(map[string]*lib.AggregatedCostItem)
	} else {
		n.children = make(map[string]*costAggNode)
	}
	return n
}

func (n *costAggNode) child(key string, leaf bool) *costAggNode {
	c, ok := n.children[key]
	if !ok {
		c = newCostAggNode(leaf)
		n.children[strings.Clone(key)] = c
	}
	return c
}

// merge adds the data of o into n. Nodes that only exist in o are adopted as is.
func (n *costAggNode) merge(o *costAggNode) {
	for day, cost := range o.costPerDay {
		n.costPerDay[day] += cost
	}
	n.total += o.total

	for key, oc := range o.children {
		if nc, ok := n.children[key]; ok {
			nc.merge(oc)
		} else {
			n.children[key] = oc
		}
	}

	for key, om := range o.meters {
		nm, ok := n.meters[key]
		if !ok {
			n.meters[key] = om
			continue
		}
		// Descriptive fields stay as they were on the first row seen for the meter
		for day, cost := range om.CostPerDay {
			nm.CostPerDay[day] += cost
		}
		for day, qty := range om.UsageQuantityPerDay {
			nm.UsageQuantityPerDay[day] += qty
		}
		nm.MonthTotalCost += om.MonthTotalCost
	}
}

// costRowKey holds every raw CSV value that decides which tenant, subscription,
// resource group, resource and meter a row belongs to. Rows for the same meter on
// different days share a key, so all of the lowercasing, splitting and identifier
// building only has to happen once per meter rather than once per row.
type costRowKey struct {
	subscriptionGuid string
	subscriptionName string
	resourceGroup    string
	instanceId       string
	meterCategory    string
	meterSubcategory string
	meterName        string
}

func (k costRowKey) clone() costRowKey {
	return costRowKey{
		subscriptionGuid: strings.Clone(k.subscriptionGuid),
		subscriptionName: strings.Clone(k.subscriptionName),
		resourceGroup:    strings.Clone(k.resourceGroup),
		instanceId:       strings.Clone(k.instanceId),
		meterCategory:    strings.Clone(k.meterCategory),
		meterSubcategory: strings.Clone(k.meterSubcategory),
		meterName:        strings.Clone(k.meterName),
	}
}

type costRowTarget struct {
	nodes [4]*costAggNode // tenant, subscription, resource group, resource
	meter *lib.AggregatedCostItem
}

type costFileAggregator struct {
	datafile      string
	customTenants map[string]string
	tenants       map[string]*costAggNode
	targets       map[costRowKey]*costRowTarget
	days          map[string]string
}

// lowerClone lowercases s into a string that does not share memory with s.
// encoding/csv backs all fields of a record with one string, so keeping a plain
// substring would keep the whole CSV line alive.
func lowerClone(s string) string {
	return strings.Clone(strings.ToLower(s))
}

// day returns the lowercased usage date, interned so that the per day maps of every
// node share the same handful of strings.
func (a *costFileAggregator) day(raw string) string {
	if d, ok := a.days[raw]; ok {
		return d
	}
	d := lowerClone(raw)
	a.days[strings.Clone(raw)] = d
	return d
}

// resolve finds or creates the nodes and meter a row belongs to. The values built
// here are the same ones TransformCostDataNew and AggregateCostData produce.
func (a *costFileAggregator) resolve(record []string, resourceRate float64) *costRowTarget {
	subscriptionGuid := record[costColSubscriptionGuid]
	instanceId := lowerClone(record[costColInstanceId])

	tenantName, ok := a.customTenants[subscriptionGuid]
	if !ok || tenantName == "" {
		tenantName = a.datafile
	}
	tenantName = strings.ToUpper(tenantName)

	subscriptionName := lowerClone(record[costColSubscriptionName])

	resGrp := "nil_res_group"
	if record[costColResourceGroup] != "" {
		resGrp = strings.ToLower(record[costColResourceGroup])
	}

	resourceName := instanceId[strings.LastIndexByte(instanceId, '/')+1:]

	identifierResName := "providers/"
	if i := strings.Index(instanceId, "providers/"); i >= 0 {
		identifierResName = instanceId[i:]
	}

	meterCategory := lowerClone(record[costColMeterCategory])
	meterSubcategory := lowerClone(record[costColMeterSubcategory])
	meterName := lowerClone(record[costColMeterName])

	var id strings.Builder
	id.Grow(len(subscriptionGuid) + len(resGrp) + len(identifierResName) + len(meterCategory) + len(meterSubcategory) + len(meterName) + 5)
	id.WriteString(subscriptionGuid)
	id.WriteByte('/')
	if record[costColResourceGroup] != "" {
		id.WriteString(resGrp)
		id.WriteByte('/')
	}
	id.WriteString(identifierResName)
	id.WriteByte('_')
	id.WriteString(meterCategory)
	id.WriteByte('_')
	id.WriteString(meterSubcategory)
	id.WriteByte('_')
	id.WriteString(meterName)
	resourceMeterIdentifier := id.String()

	tenant, ok := a.tenants[tenantName]
	if !ok {
		tenant = newCostAggNode(false)
		a.tenants[strings.Clone(tenantName)] = tenant
	}
	sub := tenant.child(subscriptionName, false)
	rg := sub.child(resGrp, false)
	res := rg.child(resourceName, true)

	meter, ok := res.meters[resourceMeterIdentifier]
	if !ok {
		// Tags and AdditionalInfo are only kept from the first row of a meter, so
		// they are only parsed here and not for every row.
		var (
			tags           map[string]string
			additionalInfo interface{}
		)
		json.Unmarshal([]byte("{"+record[costColTags]+"}"), &tags)
		if len(tags) == 0 {
			tags = nil
		}
		json.Unmarshal([]byte(record[costColAdditionalInfo]), &additionalInfo)

		meter = &lib.AggregatedCostItem{
			CostPerDay:              make(map[string]float64),
			UsageQuantityPerDay:     make(map[string]float64),
			SubscriptionName:        subscriptionName,
			ProductName:             lowerClone(record[costColProductName]),
			MeterCategory:           meterCategory,
			MeterSubcategory:        meterSubcategory,
			MeterName:               meterName,
			UnitOfMeasure:           strings.Clone(record[costColUnitOfMeasure]),
			ResourceRate:            resourceRate,
			ConsumedService:         lowerClone(record[costColConsumedService]),
			ResourceType:            lowerClone(record[costColResourceType]),
			InstanceId:              instanceId,
			Tags:                    tags,
			AdditionalInfo:          additionalInfo,
			ResourceMeterIdentifier: resourceMeterIdentifier,
		}
		res.meters[resourceMeterIdentifier] = meter
	}

	return &costRowTarget{
		nodes: [4]*costAggNode{tenant, sub, rg, res},
		meter: meter,
	}
}

func (a *costFileAggregator) add(record []string) error {
	if len(record) != costColCount {
		return &lib.FieldMismatch{Expected: costColCount + 1, Found: len(record)}
	}
	// Repeated header row
	if record[costColResourceGroup] == "ResourceGroup" {
		return nil
	}

	usageQuantity, err := strconv.ParseFloat(record[costColUsageQuantity], 64)
	if err != nil {
		return err
	}
	resourceRate, err := strconv.ParseFloat(record[costColResourceRate], 64)
	if err != nil {
		return err
	}
	preTaxCost, err := strconv.ParseFloat(record[costColPreTaxCost], 64)
	if err != nil {
		return err
	}

	key := costRowKey{
		subscriptionGuid: record[costColSubscriptionGuid],
		subscriptionName: record[costColSubscriptionName],
		resourceGroup:    record[costColResourceGroup],
		instanceId:       record[costColInstanceId],
		meterCategory:    record[costColMeterCategory],
		meterSubcategory: record[costColMeterSubcategory],
		meterName:        record[costColMeterName],
	}

	target, ok := a.targets[key]
	if !ok {
		target = a.resolve(record, resourceRate)
		a.targets[key.clone()] = target
	}

	day := a.day(record[costColUsageDateTime])

	for _, n := range target.nodes {
		n.costPerDay[day] += preTaxCost
		n.total += preTaxCost
	}
	target.meter.CostPerDay[day] += preTaxCost
	target.meter.UsageQuantityPerDay[day] += usageQuantity
	target.meter.MonthTotalCost += preTaxCost

	return nil
}

// aggregateCostExportCSVFile streams one cost export CSV and aggregates it as it
// goes, so the rows themselves are never held in memory.
func aggregateCostExportCSVFile(fileName string, customTenants map[string]string) (map[string]*costAggNode, error) {
	// File names look like <prefix>__<TENANT>.csv
	nameSplit := strings.Split(fileName, "__")
	if len(nameSplit) < 2 {
		return nil, fmt.Errorf("cost export file name is missing the '__<tenant>' suffix")
	}
	datafile := strings.Split(nameSplit[1], ".")[0]

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReaderSize(file, 1<<20))
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true
	reader.ReuseRecord = true

	// Header row
	if _, err := reader.Read(); err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}

	agg := costFileAggregator{
		datafile:      datafile,
		customTenants: customTenants,
		tenants:       make(map[string]*costAggNode),
		targets:       make(map[costRowKey]*costRowTarget),
		days:          make(map[string]string, 31),
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if err := agg.add(record); err != nil {
			line, _ := reader.FieldPos(0)
			return nil, fmt.Errorf("line %d: %w", line, err)
		}
	}

	return agg.tenants, nil
}

func costAggTreeToAggregatedCostData(tenants map[string]*costAggNode) lib.AggregatedCostData {
	data := make(lib.AggregatedCostData, len(tenants))

	for tenantName, tenant := range tenants {
		subs := make(map[string]lib.AggregatedCostSubscription, len(tenant.children))

		for subName, sub := range tenant.children {
			resGrps := make(map[string]lib.AggregatedCostResourceGroup, len(sub.children))

			for rgName, rg := range sub.children {
				resources := make(map[string]lib.AggregatedCostResource, len(rg.children))

				for resName, res := range rg.children {
					meters := make(map[string]lib.AggregatedCostItem, len(res.meters))
					for meterId, meter := range res.meters {
						meters[meterId] = *meter
					}
					resources[resName] = lib.AggregatedCostResource{
						CostPerDay:     res.costPerDay,
						MonthTotalCost: res.total,
						MeterData:      meters,
					}
				}

				resGrps[rgName] = lib.AggregatedCostResourceGroup{
					CostPerDay:     rg.costPerDay,
					MonthTotalCost: rg.total,
					Resources:      resources,
				}
			}

			subs[subName] = lib.AggregatedCostSubscription{
				CostPerDay:     sub.costPerDay,
				MonthTotalCost: sub.total,
				ResourceGroups: resGrps,
			}
		}

		data[tenantName] = lib.AggregatedCostTenant{
			CostPerDay:     tenant.costPerDay,
			MonthTotalCost: tenant.total,
			Subscriptions:  subs,
		}
	}

	return data
}

// aggregateCostExportCSVFiles aggregates the given files in parallel, one goroutine
// per file up to the number of CPUs, then merges the per file results in the order
// the files were given.
func aggregateCostExportCSVFiles(csvFiles []string, customTenants map[string]string) (lib.AggregatedCostData, error) {
	var (
		wg      sync.WaitGroup
		results = make([]map[string]*costAggNode, len(csvFiles))
		errs    = make([]error, len(csvFiles))
		limit   = make(chan struct{}, max(runtime.NumCPU(), 1))
	)

	for i, file := range csvFiles {
		wg.Add(1)
		go func() {
			defer wg.Done()
			limit <- struct{}{}
			defer func() { <-limit }()
			results[i], errs[i] = aggregateCostExportCSVFile(file, customTenants)
		}()
	}
	wg.Wait()

	for i, err := range errs {
		if err != nil {
			return nil, fmt.Errorf("%s: %w", csvFiles[i], err)
		}
	}

	merged := make(map[string]*costAggNode)
	for i, tenants := range results {
		for tenantName, tenant := range tenants {
			if existing, ok := merged[tenantName]; ok {
				existing.merge(tenant)
			} else {
				merged[tenantName] = tenant
			}
		}
		results[i] = nil
	}

	return costAggTreeToAggregatedCostData(merged), nil
}

// AggregateCostExportCSVFiles reads every cost export CSV under dataPath and
// aggregates the rows by tenant, subscription, resource group, resource and meter.
//
// It produces the same data as
//
//	TransformCostDataNew(CombineCostExportCSVData(dataPath), ...)
//
// but streams each file and aggregates while reading, rather than loading every row
// of every file into memory and then transforming them one at a time.
func AggregateCostExportCSVFiles(dataPath string) (lib.AggregatedCostData, error) {
	var csvFiles []string
	for _, file := range lib.GetFullFilePaths(dataPath) {
		if strings.HasSuffix(file, ".csv") {
			csvFiles = append(csvFiles, file)
		}
	}
	sort.Strings(csvFiles)

	customTenants := make(map[string]string)
	cfg := lib.GetCldConfig(nil)
	if cfg.Azure != nil {
		for name, subIds := range cfg.Azure.CustomSubIdToTenantNameMap {
			for _, subId := range subIds {
				customTenants[subId] = name
			}
		}
	}

	return aggregateCostExportCSVFiles(csvFiles, customTenants)
}
