package mongodb

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/jercle/cloudini/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Number of write models sent per BulkWrite call
	costBulkWriteChunkSize = 1000
	// Number of BulkWrite calls in flight at once for a single collection
	costBulkWriteWorkers = 4
)

// newCostUpsertModel builds a single upsert that sets both the fields of doc and
// the cost data for the month. Having both in one $set means one write per
// document, rather than one write for the details and another for the cost data.
func newCostUpsertModel(id string, doc interface{}, costDataKey string, costData lib.MongoDbCostData) mongo.WriteModel {
	raw, err := bson.Marshal(doc)
	lib.CheckFatalError(err)

	var set bson.D
	err = bson.Unmarshal(raw, &set)
	lib.CheckFatalError(err)

	set = append(set, bson.E{Key: costDataKey, Value: costData})

	return mongo.NewUpdateOneModel().
		SetFilter(bson.D{{Key: "_id", Value: id}}).
		SetUpdate(bson.D{{Key: "$set", Value: set}}).
		SetUpsert(true)
}

// bulkWriteCostModels writes models in unordered chunks, a few chunks at a time
func bulkWriteCostModels(ctx context.Context, coll *mongo.Collection, models []mongo.WriteModel) {
	if len(models) == 0 {
		return
	}

	var (
		wg    sync.WaitGroup
		limit = make(chan struct{}, costBulkWriteWorkers)
		opts  = options.BulkWrite().SetOrdered(false)
	)

	for i := 0; i < len(models); i += costBulkWriteChunkSize {
		chunk := models[i:min(i+costBulkWriteChunkSize, len(models))]

		wg.Add(1)
		limit <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-limit }()
			_, err := coll.BulkWrite(ctx, chunk, opts)
			lib.CheckFatalError(err)
		}()
	}
	wg.Wait()
}

// updateLifetimeTotalCosts sets lifetimeTotalCost on every document in coll to the
// sum of monthTotalCost across all of the months in its costData.
//
// The sum is done by the database with an aggregation pipeline update, so no cost
// data has to be pulled down and pushed back. Pipeline updates need MongoDB 4.2 or
// later. If the server rejects the update, the totals are calculated client side.
func updateLifetimeTotalCosts(ctx context.Context, coll *mongo.Collection) {
	pipeline := mongo.Pipeline{
		{{Key: "$set", Value: bson.D{
			{Key: "lifetimeTotalCost", Value: bson.D{
				{Key: "$toDouble", Value: bson.D{
					{Key: "$sum", Value: bson.D{
						{Key: "$map", Value: bson.D{
							{Key: "input", Value: bson.D{
								{Key: "$objectToArray", Value: bson.D{
									{Key: "$ifNull", Value: bson.A{"$costData", bson.D{}}},
								}},
							}},
							{Key: "as", Value: "month"},
							{Key: "in", Value: bson.D{
								{Key: "$ifNull", Value: bson.A{"$$month.v.monthTotalCost", 0}},
							}},
						}},
					}},
				}},
			}},
		}}},
	}

	_, err := coll.UpdateMany(ctx, bson.D{}, pipeline)
	if err == nil {
		return
	}

	fmt.Println("Could not calculate lifetime costs in the database for '" + coll.Name() + "', calculating locally instead: " + err.Error())
	updateLifetimeTotalCostsClientSide(ctx, coll)
}

// updateLifetimeTotalCostsClientSide is the fallback for updateLifetimeTotalCosts.
// It streams the collection, only decoding the month totals, and writes the
// lifetime totals back in batches.
func updateLifetimeTotalCostsClientSide(ctx context.Context, coll *mongo.Collection) {
	type costTotals struct {
		Id       interface{} `bson:"_id"`
		CostData map[string]struct {
			MonthTotalCost float64 `bson:"monthTotalCost"`
		} `bson:"costData"`
	}

	cursor, err := coll.Find(ctx, bson.D{}, options.Find().SetProjection(bson.D{{Key: "costData", Value: 1}}))
	lib.CheckFatalError(err)
	defer cursor.Close(ctx)

	var updates []mongo.WriteModel

	for cursor.Next(ctx) {
		var doc costTotals
		err := cursor.Decode(&doc)
		lib.CheckFatalError(err)

		lifetimeCost := float64(0)
		for _, monthData := range doc.CostData {
			lifetimeCost += monthData.MonthTotalCost
		}

		updates = append(updates, mongo.NewUpdateOneModel().
			SetFilter(bson.D{{Key: "_id", Value: doc.Id}}).
			SetUpdate(bson.D{{Key: "$set", Value: bson.D{{Key: "lifetimeTotalCost", Value: lifetimeCost}}}}))

		if len(updates) >= costBulkWriteChunkSize*costBulkWriteWorkers {
			bulkWriteCostModels(ctx, coll, updates)
			updates = nil
		}
	}
	lib.CheckFatalError(cursor.Err())

	bulkWriteCostModels(ctx, coll, updates)
}

// citrixVmNameFromResourceName strips the last two '-' separated parts from the
// name of a disk or network interface, leaving the name of the VM it belongs to
func citrixVmNameFromResourceName(name string) string {
	nameSpl := strings.Split(name, "-")
	return strings.Join(nameSpl[:max(len(nameSpl)-2, 0)], "-")
}

func UpsertMonthlyTenantSubResGrpCosts(
	costData lib.AggregatedCostData,
	ctxDelGrpsByResGrpAndName map[string]string,
	costExportMonth string,
	costingTenantsColl *mongo.Collection,
	costingSubsColl *mongo.Collection,
	costingResGrpsColl *mongo.Collection,
	costingResourcesColl *mongo.Collection,
	costingMetersColl *mongo.Collection,
	tenantsColl *mongo.Collection,
	allResourcesColl *mongo.Collection,
) {
	if len(costData) == 0 {
		fmt.Println("No cost data to upsert for " + costExportMonth)
		return
	}

	var (
		tenants         []MongoDbAzureTenant
		updateTenants   []mongo.WriteModel
		updateSubs      []mongo.WriteModel
		updateResGrps   []mongo.WriteModel
		updateResources []mongo.WriteModel
		updateMeters    []mongo.WriteModel
	)

	ctx := context.TODO()
	rsp, err := tenantsColl.Find(ctx, bson.D{{}})
	lib.CheckFatalError(err)
	err = rsp.All(ctx, &tenants)
	lib.CheckFatalError(err)

	tenantAliases := make(map[string]string)

	for _, tData := range tenants {
		for _, alias := range tData.Aliases {
			tenantAliases[alias] = tData.TenantName
		}
	}

	costDataKey := "costData." + costExportMonth
	now := time.Now()

	for tenantName, tenantData := range costData {

		matchName := ""
		alias, ok := tenantAliases[tenantName]
		if ok {
			matchName = alias
		} else {
			matchName = tenantName
		}
		var tenantDetails MongoDbAzureTenant
		for _, t := range tenants {
			if strings.EqualFold(t.TenantName, matchName) {
				tenantDetails = t
			}
		}

		var currTenant lib.MongoDbCostTenant
		var tcd lib.MongoDbCostData
		tcd.CostPerDay = tenantData.CostPerDay
		tcd.MonthTotalCost = tenantData.MonthTotalCost
		currTenant.CostGroups = tenantData.CostGroups
		currTenant.TenantName = tenantName
		currTenant.TenantId = tenantDetails.TenantId
		currTenant.LastDBSync = now

		for subName, subData := range tenantData.Subscriptions {
			var currSub lib.MongoDbCostSubscription
			var scd lib.MongoDbCostData
			scd.CostPerDay = subData.CostPerDay
			scd.MonthTotalCost = subData.MonthTotalCost

			currSub.CostGroups = subData.CostGroups
			currSub.TenantId = currTenant.TenantId
			currSub.TenantName = tenantName
			currSub.SubscriptionId = tenantDetails.Subscriptions[strings.ToLower(subName)].SubscriptionID
			currSub.SubscriptionName = subName
			currSub.LastDBSync = now

			for rgName, rgData := range subData.ResourceGroups {
				var currResGrp lib.MongoDbCostResourceGroup
				var rgcd lib.MongoDbCostData
				rgcd.CostPerDay = rgData.CostPerDay
				rgcd.MonthTotalCost = rgData.MonthTotalCost

				currResGrp.CostGroups = rgData.CostGroups
				currResGrp.TenantId = currSub.TenantId
				currResGrp.TenantName = currSub.TenantName
				currResGrp.SubscriptionId = currSub.SubscriptionId
				currResGrp.SubscriptionName = currSub.SubscriptionName
				currResGrp.Name = rgName
				currResGrp.MongoId = strings.ToLower(tenantName + "_" + currResGrp.SubscriptionId + "_" + currResGrp.Name)
				currResGrp.LastDBSync = now

				rgNameLower := strings.ToLower(rgName)

				for resName, resData := range rgData.Resources {
					var currRes lib.MongoDbCostResource
					var rcd lib.MongoDbCostData
					rcd.CostPerDay = resData.CostPerDay
					rcd.MonthTotalCost = resData.MonthTotalCost

					resType := ""

					currRes.CostGroups = resData.CostGroups
					currRes.TenantId = currSub.TenantId
					currRes.TenantName = currSub.TenantName
					currRes.SubscriptionId = currSub.SubscriptionId
					currRes.SubscriptionName = currSub.SubscriptionName
					currRes.ResourceGroupName = rgName
					currRes.ResGrpMongoId = currResGrp.MongoId
					currRes.Name = resName
					currRes.CitrixDeliveryGroup = ctxDelGrpsByResGrpAndName[rgNameLower+"_"+strings.ToLower(resName)]
					currRes.MongoId = strings.ToLower(tenantName + "_" + currRes.SubscriptionId + "_" + currRes.ResourceGroupName + "_" + resName)
					currRes.LastDBSync = now

					for _, meterData := range resData.MeterData {
						var currMeter lib.MongoDbCostMeter
						var mcd lib.MongoDbCostData
						mcd.CostPerDay = meterData.CostPerDay
						mcd.MonthTotalCost = meterData.MonthTotalCost
						mcd.ResourceRate = meterData.ResourceRate
						mcd.UnitOfMeasure = meterData.UnitOfMeasure

						currMeter.TenantId = currSub.TenantId
						currMeter.TenantName = currSub.TenantName
						currMeter.SubscriptionId = currSub.SubscriptionId
						currMeter.SubscriptionName = currSub.SubscriptionName
						currMeter.ResourceGroupName = rgName
						currMeter.ResGrpMongoId = currResGrp.MongoId
						currMeter.ResourceMongoId = currRes.MongoId
						currMeter.ResourceMeterIdentifier = meterData.ResourceMeterIdentifier
						currMeter.MeterCategory = meterData.MeterCategory
						currMeter.ProductName = meterData.ProductName
						currMeter.ConsumedService = meterData.ConsumedService
						currMeter.MeterName = meterData.MeterName
						currMeter.ResourceType = meterData.ResourceType
						currMeter.LastDBSync = now
						currMeter.ResourceName = resName

						resType = currMeter.ResourceType

						// TODO: Add the below, but unique items
						// currTenant.RelatedCostMeters = append(currTenant.RelatedCostMeters, meterData.ResourceMeterIdentifier)
						// currSub.RelatedCostMeters = append(currSub.RelatedCostMeters, meterData.ResourceMeterIdentifier)
						// currResGrp.RelatedCostMeters = append(currResGrp.RelatedCostMeters, meterData.ResourceMeterIdentifier)
						// currRes.RelatedCostMeters = append(currRes.RelatedCostMeters, meterData.ResourceMeterIdentifier)

						updateMeters = append(updateMeters, newCostUpsertModel(currMeter.ResourceMeterIdentifier, currMeter, costDataKey, mcd))
					}

					currRes.ResourceType = resType

					if strings.Contains(currRes.Name, "dtw") {
						switch currRes.ResourceType {
						case "microsoft.compute/disks", "microsoft.network/networkinterfaces":
							vmName := citrixVmNameFromResourceName(currRes.Name)
							dg := ctxDelGrpsByResGrpAndName[rgNameLower+"_"+strings.ToLower(vmName)]
							if dg == "" {
								dg = "UNKNOWN_DG-" + currRes.TenantName + "\\" + vmName
							}
							currRes.CitrixDeliveryGroup = dg
						case "microsoft.compute/virtualmachines":
							dg := currRes.CitrixDeliveryGroup
							if dg == "" {
								dg = "UNKNOWN_DG-" + currRes.TenantName + "\\" + currRes.Name
							}
							currRes.CitrixDeliveryGroup = dg
						case "microsoft.compute/snapshots":
							currRes.CitrixDeliveryGroup = "UNKNOWN_DG-" + currRes.TenantName + "\\" + currRes.Name
						}
					}

					updateResources = append(updateResources, newCostUpsertModel(currRes.MongoId, currRes, costDataKey, rcd))
				}

				updateResGrps = append(updateResGrps, newCostUpsertModel(currResGrp.MongoId, currResGrp, costDataKey, rgcd))
			}

			updateSubs = append(updateSubs, newCostUpsertModel(currSub.SubscriptionId, currSub, costDataKey, scd))
		}

		updateTenants = append(updateTenants, newCostUpsertModel(currTenant.TenantName, currTenant, costDataKey, tcd))
	}

	// Each collection is independent of the others, so they are all written at the
	// same time. The lifetime totals of a collection are recalculated as soon as its
	// own upserts are done.
	collections := []struct {
		name    string
		coll    *mongo.Collection
		updates []mongo.WriteModel
	}{
		{"tenant", costingTenantsColl, updateTenants},
		{"subscription", costingSubsColl, updateSubs},
		{"resource group", costingResGrpsColl, updateResGrps},
		{"resource", costingResourcesColl, updateResources},
		{"cost meter", costingMetersColl, updateMeters},
	}
	updateTenants, updateSubs, updateResGrps, updateResources, updateMeters = nil, nil, nil, nil, nil

	var wg sync.WaitGroup
	for i := range collections {
		c := &collections[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			startTime := time.Now()
			count := len(c.updates)

			bulkWriteCostModels(ctx, c.coll, c.updates)
			c.updates = nil

			updateLifetimeTotalCosts(ctx, c.coll)

			fmt.Printf("Upserted %d %s cost documents and updated lifetime costs in %s\n", count, c.name, time.Since(startTime).Round(time.Millisecond))
		}()
	}
	wg.Wait()
}
