package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jercle/cloudini/cmd/ad"
	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/cmd/citrix"
	"github.com/jercle/cloudini/cmd/m365"
	"github.com/jercle/cloudini/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpsertMonthlyTenantSubResGrpCosts(
	costData lib.AggregatedCostData,
	costExportMonth string,
	costingTenantsColl *mongo.Collection,
	costingSubsColl *mongo.Collection,
	costingResGrpsColl *mongo.Collection,
	costingResourcesColl *mongo.Collection,
	costingMetersColl *mongo.Collection,
	tenantsColl *mongo.Collection,
	allResourcesColl *mongo.Collection,
) (results UpsertMonthlyTenantSubResGrpCostsResults) {
	var (
		tenants                   []MongoDbAzureTenant
		updateTenants             []mongo.WriteModel
		updateTenantsCostData     []mongo.WriteModel
		mongoTenantsCostData      []lib.MongoDbCostTenant
		tenantsProcessedUpdates   []mongo.WriteModel
		updateSubs                []mongo.WriteModel
		updateSubsCostData        []mongo.WriteModel
		mongoSubsCostData         []lib.MongoDbCostSubscription
		subsProcessedUpdates      []mongo.WriteModel
		updateResGrps             []mongo.WriteModel
		updateResGrpsCostData     []mongo.WriteModel
		mongoResGrpsCostData      []lib.MongoDbCostResourceGroup
		resGrpsProcessedUpdates   []mongo.WriteModel
		updateResources           []mongo.WriteModel
		updateResourcesCostData   []mongo.WriteModel
		mongoResourcesCostData    []lib.MongoDbCostResourceGroup
		resourcesProcessedUpdates []mongo.WriteModel
		updateMeters              []mongo.WriteModel
		updateMetersCostData      []mongo.WriteModel
		mongoMetersCostData       []lib.MongoDbCostMeter
		metersProcessedUpdates    []mongo.WriteModel
	)

	_ = tenantsProcessedUpdates
	_ = subsProcessedUpdates
	_ = resGrpsProcessedUpdates
	_ = resourcesProcessedUpdates
	_ = metersProcessedUpdates

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
	// aliasStr, _ := json.MarshalIndent(tenantAliases, "", "  ")
	// fmt.Println(string(aliasStr))
	// os.Exit(0)

	// s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	// s.Start()

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
		// jsonStr, _ := json.MarshalIndent(tenantDetails, "", "  ")
		// fmt.Println(string(jsonStr))
		// fmt.Println("tenantName")
		// fmt.Println(tenantName)
		// fmt.Println("matchName")
		// fmt.Println(matchName)
		// os.Exit(0)
		// continue

		var currTenant lib.MongoDbCostTenant
		var tcd lib.MongoDbCostData
		tcd.CostPerDay = tenantData.CostPerDay
		tcd.MonthTotalCost = tenantData.MonthTotalCost
		currTenant.CostGroups = tenantData.CostGroups
		currTenant.TenantName = tenantName
		currTenant.TenantId = tenantDetails.TenantId
		currTenant.LastDBSync = time.Now()

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
			currSub.LastDBSync = time.Now()

			// fmt.Println(currSub.SubscriptionId)
			// os.Exit(0)

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
				currResGrp.LastDBSync = time.Now()

				for resName, resData := range rgData.Resources {
					// GetResource(resName, rgName, subName, tenantName, allResourcesColl)
					var currRes lib.MongoDbCostResource
					var rcd lib.MongoDbCostData
					rcd.CostPerDay = resData.CostPerDay
					rcd.MonthTotalCost = resData.MonthTotalCost

					currRes.CostGroups = resData.CostGroups
					currRes.TenantId = currSub.TenantId
					currRes.TenantName = currSub.TenantName
					currRes.SubscriptionId = currSub.SubscriptionId
					currRes.SubscriptionName = currSub.SubscriptionName
					currRes.ResourceGroupName = rgName
					currRes.ResGrpMongoId = currResGrp.MongoId
					currRes.Name = resName
					currRes.MongoId = strings.ToLower(tenantName + "_" + currRes.SubscriptionId + "_" + currRes.ResourceGroupName + "_" + resName)
					currRes.LastDBSync = time.Now()

					for _, meterData := range resData.MeterData {
						var currMeter lib.MongoDbCostMeter
						var mcd lib.MongoDbCostData
						mcd.CostPerDay = meterData.CostPerDay
						// mcd.UsageQuantityPerDay = meterData.quan
						mcd.MonthTotalCost = meterData.MonthTotalCost
						mcd.ResourceRate = meterData.ResourceRate
						mcd.UnitOfMeasure = meterData.UnitOfMeasure

						// jsonStr, _ := json.MarshalIndent(mcd, "", "  ")
						// fmt.Println(string(jsonStr))
						// os.Exit(0)

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
						currMeter.LastDBSync = time.Now()

						// TODO: Add the below, but unique items
						// currTenant.RelatedCostMeters = append(currTenant.RelatedCostMeters, meterData.ResourceMeterIdentifier)
						// currSub.RelatedCostMeters = append(currSub.RelatedCostMeters, meterData.ResourceMeterIdentifier)
						// currResGrp.RelatedCostMeters = append(currResGrp.RelatedCostMeters, meterData.ResourceMeterIdentifier)
						// currRes.RelatedCostMeters = append(currRes.RelatedCostMeters, meterData.ResourceMeterIdentifier)

						filterMeter := bson.D{{"_id", currMeter.ResourceMeterIdentifier}}
						updateMeter := bson.D{{"$set", currMeter}}
						updateMeters = append(updateMeters, mongo.NewUpdateOneModel().SetFilter(filterMeter).SetUpdate(updateMeter).SetUpsert(true))
						updateMeterCostData := bson.D{{"$set", bson.D{{"costData." + costExportMonth, mcd}}}}
						updateMetersCostData = append(updateMetersCostData, mongo.NewUpdateOneModel().SetFilter(filterMeter).SetUpdate(updateMeterCostData).SetUpsert(true))
					}

					filterRes := bson.D{{"_id", currRes.MongoId}}
					updateRes := bson.D{{"$set", currRes}}
					updateResources = append(updateResources, mongo.NewUpdateOneModel().SetFilter(filterRes).SetUpdate(updateRes).SetUpsert(true))
					updateResCostData := bson.D{{"$set", bson.D{{"costData." + costExportMonth, rcd}}}}
					updateResourcesCostData = append(updateResourcesCostData, mongo.NewUpdateOneModel().SetFilter(filterRes).SetUpdate(updateResCostData).SetUpsert(true))
				}

				filterResGrp := bson.D{{"_id", currResGrp.MongoId}}
				updateResGrp := bson.D{{"$set", currResGrp}}
				updateResGrps = append(updateResGrps, mongo.NewUpdateOneModel().SetFilter(filterResGrp).SetUpdate(updateResGrp).SetUpsert(true))
				updateResGrpCostData := bson.D{{"$set", bson.D{{"costData." + costExportMonth, rgcd}}}}
				updateResGrpsCostData = append(updateResGrpsCostData, mongo.NewUpdateOneModel().SetFilter(filterResGrp).SetUpdate(updateResGrpCostData).SetUpsert(true))
			}

			filterSub := bson.D{{"_id", currSub.SubscriptionId}}
			updateSub := bson.D{{"$set", currSub}}
			updateSubs = append(updateSubs, mongo.NewUpdateOneModel().SetFilter(filterSub).SetUpdate(updateSub).SetUpsert(true))
			updateSubCostData := bson.D{{"$set", bson.D{{"costData." + costExportMonth, scd}}}}
			updateSubsCostData = append(updateSubsCostData, mongo.NewUpdateOneModel().SetFilter(filterSub).SetUpdate(updateSubCostData).SetUpsert(true))
		}

		filterTenant := bson.D{{"_id", currTenant.TenantName}}
		updateTenant := bson.D{{"$set", currTenant}}
		updateTenants = append(updateTenants, mongo.NewUpdateOneModel().SetFilter(filterTenant).SetUpdate(updateTenant).SetUpsert(true))
		updateTenantCostData := bson.D{{"$set", bson.D{{"costData." + costExportMonth, tcd}}}}
		updateTenantsCostData = append(updateTenantsCostData, mongo.NewUpdateOneModel().SetFilter(filterTenant).SetUpdate(updateTenantCostData).SetUpsert(true))
	}
	// os.Exit(0)

	fmt.Println("Upserting Tenant data...")
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	s.Start()
	results.UpdateTenants, err = costingTenantsColl.BulkWrite(ctx, updateTenants)
	lib.CheckFatalError(err)
	results.UpdateTenantsCostData, err = costingTenantsColl.BulkWrite(ctx, updateTenantsCostData)
	lib.CheckFatalError(err)
	s.Stop()

	fmt.Println("Upserting Subscription data...")
	s.Start()
	results.UpdateSubs, err = costingSubsColl.BulkWrite(ctx, updateSubs)
	lib.CheckFatalError(err)
	results.UpdateSubsCostData, err = costingSubsColl.BulkWrite(ctx, updateSubsCostData)
	lib.CheckFatalError(err)
	s.Stop()

	fmt.Println("Upserting Resource Group data...")
	s.Start()
	results.UpdateResGrps, err = costingResGrpsColl.BulkWrite(ctx, updateResGrps)
	lib.CheckFatalError(err)
	results.UpdateResGrpsCostData, err = costingResGrpsColl.BulkWrite(ctx, updateResGrpsCostData)
	lib.CheckFatalError(err)
	s.Stop()

	fmt.Println("Upserting Resource data...")
	s.Start()
	results.UpdateResources, err = costingResourcesColl.BulkWrite(ctx, updateResources)
	lib.CheckFatalError(err)
	results.UpdateResourcesCostData, err = costingResourcesColl.BulkWrite(ctx, updateResourcesCostData)
	lib.CheckFatalError(err)
	s.Stop()

	fmt.Println("Upserting cost meter data...")
	s.Start()
	results.UpdateMeters, err = costingMetersColl.BulkWrite(ctx, updateMeters)
	lib.CheckFatalError(err)
	results.UpdateMetersCostData, err = costingMetersColl.BulkWrite(ctx, updateMetersCostData)
	lib.CheckFatalError(err)
	s.Stop()

	fmt.Println("Pulling all cost data from database for sync...")
	s.Start()
	rsp, err = costingTenantsColl.Find(ctx, bson.D{{}})
	lib.CheckFatalError(err)
	err = rsp.All(ctx, &mongoTenantsCostData)
	lib.CheckFatalError(err)
	rsp, err = costingSubsColl.Find(ctx, bson.D{{}})
	lib.CheckFatalError(err)
	err = rsp.All(ctx, &mongoSubsCostData)
	lib.CheckFatalError(err)
	rsp, err = costingResGrpsColl.Find(ctx, bson.D{{}})
	lib.CheckFatalError(err)
	err = rsp.All(ctx, &mongoResGrpsCostData)
	lib.CheckFatalError(err)
	rsp, err = costingResourcesColl.Find(ctx, bson.D{{}})
	lib.CheckFatalError(err)
	err = rsp.All(ctx, &mongoResourcesCostData)
	lib.CheckFatalError(err)
	rsp, err = costingMetersColl.Find(ctx, bson.D{{}})
	lib.CheckFatalError(err)
	err = rsp.All(ctx, &mongoMetersCostData)
	lib.CheckFatalError(err)
	s.Stop()

	fmt.Println("Updating monthly cost values...")
	s.Start()
	for _, tenantData := range mongoTenantsCostData {
		currTenant := tenantData
		lifetimeCost := float64(0)
		for _, costData := range tenantData.CostData {
			lifetimeCost += costData.MonthTotalCost
		}
		// currTenant.LifetimeTotalCost = lifetimeCost
		// tenantsProcessed = append(tenantsProcessed)
		filterTenant := bson.D{{"_id", currTenant.TenantName}}

		updateTenant := bson.D{{"$set", bson.D{{"lifetimeTotalCost", lifetimeCost}}}}
		tenantsProcessedUpdates = append(tenantsProcessedUpdates, mongo.NewUpdateOneModel().SetFilter(filterTenant).SetUpdate(updateTenant).SetUpsert(true))
	}
	for _, subData := range mongoSubsCostData {
		currSub := subData
		lifetimeCost := float64(0)
		for _, costData := range subData.CostData {

			// fmt.Println(costMeterId)
			// os.Exit(0)
			lifetimeCost += costData.MonthTotalCost
		}
		// currSub.LifetimeTotalCost = lifetimeCost
		// subsProcessed = append(subsProcessed)
		filterSub := bson.D{{"_id", currSub.SubscriptionId}}

		updateSub := bson.D{{"$set", bson.D{{"lifetimeTotalCost", lifetimeCost}}}}
		subsProcessedUpdates = append(subsProcessedUpdates, mongo.NewUpdateOneModel().SetFilter(filterSub).SetUpdate(updateSub).SetUpsert(true))
	}
	for _, resGrpData := range mongoResGrpsCostData {
		currResGrp := resGrpData
		lifetimeCost := float64(0)
		for _, costData := range resGrpData.CostData {
			lifetimeCost += costData.MonthTotalCost
		}
		// currResGrp.LifetimeTotalCost = lifetimeCost
		// resGrpsProcessed = append(resGrpsProcessed)

		filterResGrp := bson.D{{"_id", currResGrp.MongoId}}

		updateResGrp := bson.D{{"$set", bson.D{{"lifetimeTotalCost", lifetimeCost}}}}
		resGrpsProcessedUpdates = append(resGrpsProcessedUpdates, mongo.NewUpdateOneModel().SetFilter(filterResGrp).SetUpdate(updateResGrp).SetUpsert(true))
	}
	for _, resData := range mongoResourcesCostData {
		currRes := resData
		lifetimeCost := float64(0)
		for _, costData := range resData.CostData {
			lifetimeCost += costData.MonthTotalCost
		}
		currRes.LifetimeTotalCost = lifetimeCost
		// resGrpsProcessed = append(resGrpsProcessed)

		filterRes := bson.D{{"_id", currRes.MongoId}}

		updateRes := bson.D{{"$set", bson.D{{"lifetimeTotalCost", lifetimeCost}}}}
		resourcesProcessedUpdates = append(resourcesProcessedUpdates, mongo.NewUpdateOneModel().SetFilter(filterRes).SetUpdate(updateRes).SetUpsert(true))
	}
	for _, meterData := range mongoMetersCostData {
		currMeter := meterData
		lifetimeCost := float64(0)
		for _, costData := range meterData.CostData {
			lifetimeCost += costData.MonthTotalCost
		}
		currMeter.LifetimeTotalCost = lifetimeCost
		// resGrpsProcessed = append(resGrpsProcessed)

		filterMeter := bson.D{{"_id", meterData.ResourceMeterIdentifier}}

		updateMeter := bson.D{{"$set", bson.D{{"lifetimeTotalCost", lifetimeCost}}}}
		metersProcessedUpdates = append(metersProcessedUpdates, mongo.NewUpdateOneModel().SetFilter(filterMeter).SetUpdate(updateMeter).SetUpsert(true))
	}
	s.Stop()

	fmt.Println("Pushing processed data back to database...")
	s.Start()
	results.UpdateTenantsProcessedUpdates, err = costingTenantsColl.BulkWrite(ctx, tenantsProcessedUpdates)
	lib.CheckFatalError(err)

	results.UpdateSubsProcessedUpdates, err = costingSubsColl.BulkWrite(ctx, subsProcessedUpdates)
	lib.CheckFatalError(err)

	results.UpdateResGrpsProcessedUpdates, err = costingResGrpsColl.BulkWrite(ctx, resGrpsProcessedUpdates)
	lib.CheckFatalError(err)

	results.UpdateResourcesProcessedUpdates, err = costingResGrpsColl.BulkWrite(ctx, resGrpsProcessedUpdates)
	lib.CheckFatalError(err)
	s.Stop()
	// // fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
	// fmt.Printf("Number of documents matched: %d\n", results.MatchedCount)
	// // fmt.Printf("Number of documents matched: %d\n", )
	// fmt.Printf("Number of documents inserted: %d\n", results.UpsertedCount)
	// fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
	// fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
	// fmt.Println("Upserted IDs:")

	// jsonStr, _ := json.MarshalIndent(results, "", "  ")
	// fmt.Println(string(jsonStr))

	// os.WriteFile("cost-exports/UpsertMonthlyTenantSubResGrpCosts-"+costExportMonth+".json", jsonStr, 0644)
	// s.Stop()
	return results
}

//
//

// func UpsertMultipleMonthlyCostData(
// 	costExportMonths []string,
// 	costingTenantsColl *mongo.Collection,
// 	costingSubsColl *mongo.Collection,
// 	costingResGrpsColl *mongo.Collection,
// 	costingResourcesColl *mongo.Collection,
// 	costingMetersColl *mongo.Collection,
// 	resourcesTenantsColl *mongo.Collection,
// ) {
// 	for _, costExportMonth := range costExportMonths {
// 		fmt.Println(costExportMonth)
// 		monthCostFile, err := os.ReadFile("cost-exports/" + costExportMonth + "/MonthlyCostReport-" + costExportMonth + ".json")
// 		lib.CheckFatalError(err)
// 		var monthCostData lib.AggregatedCostData
// 		// fmt.Println(costExportMonth)
// 		err = json.Unmarshal(monthCostFile, &monthCostData)
// 		lib.CheckFatalError(err)

// 		UpsertMonthlyTenantSubResGrpCosts(monthCostData,
// 			costExportMonth,
// 			costingTenantsColl,
// 			costingSubsColl,
// 			costingResGrpsColl,
// 			costingResourcesColl,
// 			costingMetersColl,
// 			resourcesTenantsColl,
// 			allResou
// 		)
// 	}
// }

//
//

func UpsertTenantAndSubs(tenantsColl *mongo.Collection, tokenReq *lib.AllTenantTokens) (results *mongo.BulkWriteResult) {

	allSubs := azure.ListAllAuthenticatedSubscriptions(tokenReq)

	config := lib.GetCldConfig(nil)

	ctx := context.TODO()

	var updates []mongo.WriteModel
	for _, tData := range allSubs {
		var currTenant MongoDbAzureTenant
		currTenant.Subscriptions = make(map[string]MongoDbAzureSubscription)
		currTenant.TenantName = tData.TenantName
		currTenant.TenantId = tData.TenantId
		currTenant.CostExportsLocation = config.Azure.MultiTenantAuth.Tenants[tData.TenantName].CostExportsLocation
		// if config.Azure.TenantAliases[tData.TenantName] != "" {
		// 	currTenant.Aliases = append(currTenant.Aliases, config.Azure.TenantAliases[tData.TenantName])
		// }

		for alias, tName := range config.Azure.TenantAliases {
			if tName == tData.TenantName {
				currTenant.Aliases = append(currTenant.Aliases, alias)
			}
		}

		token, err := tokenReq.SelectTenant(tData.TenantName)
		lib.CheckFatalError(err)

		for _, sId := range tData.Subscriptions {
			var currSub MongoDbAzureSubscription
			res, err := azure.HttpGet("https://management.azure.com/subscriptions/"+sId+"?api-version=2022-12-01", *token)
			lib.CheckFatalError(err)

			err = json.Unmarshal(res, &currSub)
			lib.CheckFatalError(err)
			currTenant.Subscriptions[strings.ToLower(currSub.DisplayName)] = currSub
			currTenant.TenantId = currSub.TenantId
		}

		allTenantsFilter := bson.D{{}}

		rsp, err := tenantsColl.Find(ctx, allTenantsFilter)
		lib.CheckFatalError(err)
		var retrievedTenantData []MongoDbAzureTenant
		err = rsp.All(ctx, &retrievedTenantData)
		lib.CheckFatalError(err)

		filter := bson.D{{"_id", currTenant.TenantId}}
		update := bson.D{{"$set", currTenant}}
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))

		// jsonStr, _ := json.MarshalIndent(r, "", "  ")
		// fmt.Println(string(jsonStr))
		// UpsertResource()
	}
	results, err := tenantsColl.BulkWrite(ctx, updates)
	lib.CheckFatalError(err)
	return
	// jsonStr, _ := json.MarshalIndent(results, "", "  ")
	// fmt.Println(string(jsonStr))
}

//
//

func UpsertImageGalleryImages(images []lib.GalleryImage, collection *mongo.Collection) (results []mongo.BulkWriteResult) {
	if len(images) == 0 {
		fmt.Println("No images in slice")
		return nil
	}
	ctx := context.TODO()

	var updates []mongo.WriteModel

	for _, img := range images {
		curr := img
		curr.LastDBSync = time.Now()
		filter := bson.D{{"_id", img.ID}}
		update := bson.D{{"$set", curr}}
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	var opts options.BulkWriteOptions
	opts.SetOrdered(false)

	chunkSize := 100
	var chunks [][]mongo.WriteModel
	for i := 0; i < len(updates); i += chunkSize {
		end := i + chunkSize
		if end > len(updates) {
			end = len(updates)
		}
		chunks = append(chunks, updates[i:end])
	}

	for _, chunk := range chunks {
		res, err := collection.BulkWrite(ctx, chunk, &opts)
		results = append(results, *res)
		lib.CheckFatalError(err)
	}

	// fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
	// fmt.Printf("Number of documents matched: %d\n", results.MatchedCount)
	// // fmt.Printf("Number of documents matched: %d\n", )
	// fmt.Printf("Number of documents inserted: %d\n", results.UpsertedCount)
	// fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
	// fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
	// fmt.Println("Upserted IDs:")
	// jsonStr, _ := json.MarshalIndent(results.UpsertedIDs, "", "  ")
	// fmt.Println(string(jsonStr))
	return results
}

//
//

func UpsertResourceSKUs(skus []lib.AzureResourceSku, collection *mongo.Collection) (results []mongo.BulkWriteResult) {
	if len(skus) == 0 {
		fmt.Println("No apps in slice")
		return nil
	}
	ctx := context.TODO()

	var updates []mongo.WriteModel

	for _, sku := range skus {
		curr := sku
		curr.LastDBSync = time.Now()
		filter := bson.D{{"_id", sku.Name}}
		update := bson.D{{"$set", curr}}

		// .SetUpsert(true)
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	var opts options.BulkWriteOptions
	opts.SetOrdered(false)

	chunkSize := 100
	var chunks [][]mongo.WriteModel
	for i := 0; i < len(updates); i += chunkSize {
		end := i + chunkSize
		if end > len(updates) {
			end = len(updates)
		}
		chunks = append(chunks, updates[i:end])
	}

	for _, chunk := range chunks {
		res, err := collection.BulkWrite(ctx, chunk, &opts)
		results = append(results, *res)
		lib.CheckFatalError(err)
	}

	// fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
	// fmt.Printf("Number of documents matched: %d\n", results.MatchedCount)
	// fmt.Printf("Number of documents matched: %d\n", )
	// fmt.Printf("Number of documents inserted: %d\n", results.UpsertedCount)
	// fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
	// fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
	// fmt.Println("Upserted IDs:")
	// jsonStr, _ := json.MarshalIndent(results.UpsertedIDs, "", "  ")
	// fmt.Println(string(jsonStr))
	return results
}

//
//

func UpsertCitrixMachineCatalogs(machineCatalogs citrix.MachineCatalogs, coll *mongo.Collection) (results []mongo.BulkWriteResult) {
	if len(machineCatalogs) == 0 {
		fmt.Println("No apps in slice")
		return nil
	}
	ctx := context.TODO()

	var updates []mongo.WriteModel

	for _, mc := range machineCatalogs {
		curr := mc
		curr.LastDBSync = time.Now()
		filter := bson.D{{"_id", mc.ID}}
		update := bson.D{{"$set", curr}}

		// .SetUpsert(true)
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	var opts options.BulkWriteOptions
	opts.SetOrdered(false)

	chunkSize := 100
	var chunks [][]mongo.WriteModel
	for i := 0; i < len(updates); i += chunkSize {
		end := i + chunkSize
		if end > len(updates) {
			end = len(updates)
		}
		chunks = append(chunks, updates[i:end])
	}

	for _, chunk := range chunks {
		res, err := coll.BulkWrite(ctx, chunk, &opts)
		results = append(results, *res)
		lib.CheckFatalError(err)
	}

	// fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
	// fmt.Printf("Number of documents matched: %d\n", results.MatchedCount)
	// fmt.Printf("Number of documents matched: %d\n", )
	// fmt.Printf("Number of documents inserted: %d\n", results.UpsertedCount)
	// fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
	// fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
	// fmt.Println("Upserted IDs:")
	// jsonStr, _ := json.MarshalIndent(results.UpsertedIDs, "", "  ")
	// fmt.Println(string(jsonStr))
	return results
}

//
//

func UpsertMultipleEntraApps[T azure.EntraApplication | azure.EntraExpiringCredential](apps []T, collection *mongo.Collection) (results []mongo.BulkWriteResult) {
	if len(apps) == 0 {
		fmt.Println("No apps in slice")
		return nil
	}
	ctx := context.TODO()

	var updates []mongo.WriteModel

	for _, app := range apps {
		var filter bson.D
		var update bson.D
		switch any(app).(type) {
		case azure.EntraApplication:
			currStr, _ := json.Marshal(app)
			var curr azure.EntraApplication
			err := json.Unmarshal(currStr, &curr)
			curr.LastDBSync = time.Now()

			lib.CheckFatalError(err)
			filter = bson.D{{"_id", curr.AppID}}
			update = bson.D{{"$set", curr}}
		case azure.EntraExpiringCredential:
			currStr, _ := json.Marshal(app)
			var curr azure.EntraExpiringCredential
			err := json.Unmarshal(currStr, &curr)
			curr.LastDBSync = time.Now()

			lib.CheckFatalError(err)
			filter = bson.D{{"_id", curr.MongoDbId}}
			update = bson.D{{"$set", curr}}
		default:
			panic("This should not happen")
		}
		// curr := app
		// curr.LastDBSync = time.Now()
		// filter := bson.D{{"_id", app.AppID}}
		// update := bson.D{{"$set", curr}}

		// .SetUpsert(true)
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	var opts options.BulkWriteOptions
	opts.SetOrdered(false)

	chunkSize := 100
	var chunks [][]mongo.WriteModel
	for i := 0; i < len(updates); i += chunkSize {
		end := i + chunkSize
		if end > len(updates) {
			end = len(updates)
		}
		chunks = append(chunks, updates[i:end])
	}

	for _, chunk := range chunks {
		res, err := collection.BulkWrite(ctx, chunk, &opts)
		results = append(results, *res)
		lib.CheckFatalError(err)
	}

	// fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
	// fmt.Printf("Number of documents matched: %d\n", results.MatchedCount)
	// fmt.Printf("Number of documents matched: %d\n", )
	// fmt.Printf("Number of documents inserted: %d\n", results.UpsertedCount)
	// fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
	// fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
	// fmt.Println("Upserted IDs:")
	// jsonStr, _ := json.MarshalIndent(results.UpsertedIDs, "", "  ")
	// fmt.Println(string(jsonStr))
	return results
}

//
//

func UpsertVcpuCounts(vcpuCountData lib.VCpuCountByTenant, collection *mongo.Collection) (results []mongo.BulkWriteResult) {
	if len(vcpuCountData) == 0 {
		fmt.Println("No data in slice")
		return nil
	}
	ctx := context.TODO()

	var updates []mongo.WriteModel

	for tName, tData := range vcpuCountData {
		curr := tData
		curr.LastDBSync = time.Now()
		filter := bson.D{{"_id", tName}}
		update := bson.D{{"$set", curr}}

		// .SetUpsert(true)
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	var opts options.BulkWriteOptions
	opts.SetOrdered(false)

	chunkSize := 100
	var chunks [][]mongo.WriteModel
	for i := 0; i < len(updates); i += chunkSize {
		end := i + chunkSize
		if end > len(updates) {
			end = len(updates)
		}
		chunks = append(chunks, updates[i:end])
	}

	for _, chunk := range chunks {
		res, err := collection.BulkWrite(ctx, chunk, &opts)
		results = append(results, *res)
		lib.CheckFatalError(err)
	}

	// fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
	// fmt.Printf("Number of documents matched: %d\n", results.MatchedCount)
	// fmt.Printf("Number of documents matched: %d\n", )
	// fmt.Printf("Number of documents inserted: %d\n", results.UpsertedCount)
	// fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
	// fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
	// fmt.Println("Upserted IDs:")
	// jsonStr, _ := json.MarshalIndent(results.UpsertedIDs, "", "  ")
	// fmt.Println(string(jsonStr))
	return results
}

//
//

// func UpsertMultipleResources(resources []lib.AzureResourceDetails, resourcesListColl *mongo.Collection) {
func UpsertMultipleResources(resources []lib.AzureResourceDetails, resourcesListColl *mongo.Collection) (results []mongo.BulkWriteResult) {
	// for _, res := range resources {
	// 	if res.Type == "microsoft.network/virtualnetworks/subnets" {
	// 		lib.JsonMarshalAndPrint(res)
	// 	}
	// }
	// os.Exit(0)
	ctx := context.TODO()

	var updates []mongo.WriteModel

	for _, res := range resources {
		resource := res
		// if !strings.EqualFold(resource.Type, "microsoft.compute/virtualmachines") {
		// 	resource.Properties.Sku = nil
		// }
		resource.LastDBSync = time.Now()
		resource.ID = strings.ToLower(res.ID)
		filter := bson.D{{"_id", resource.ID}}
		update := bson.D{{"$set", resource}}

		// .SetUpsert(true)
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
		// 	_, err := resourcesListColl.UpdateOne(ctx, filter, update, nil)
		// 	lib.CheckFatalError(err)
		// 	if err != nil {
		// 		// fmt.Println(err)
		// 		_, _, cachePath := lib.InitConfig(nil)
		// 		_ = updates
		// 		allResStr, _ := json.MarshalIndent(resources, "", "  ")
		// 		os.WriteFile(cachePath+"/mongo.updateOne-error.resources.json", allResStr, 0644)
		// 		jsonStr, _ := json.MarshalIndent(res, "", "  ")
		// 		os.WriteFile(cachePath+"/mongo.updateOne-error.err.json", jsonStr, 0644)
		// 		// fmt.Println(string(jsonStr))
		// 		fmt.Println(res.ID)
		// 		lib.CheckFatalError(err)
		// 		// os.Exit(1)
		// 	}
	}

	// results, err := resourcesListColl.BulkWrite(ctx, updates)
	// lib.CheckFatalError(err)
	// return results

	if len(updates) > 0 {
		var opts options.BulkWriteOptions
		opts.SetOrdered(true)

		chunkSize := 5000
		var chunks [][]mongo.WriteModel
		for i := 0; i < len(updates); i += chunkSize {
			end := i + chunkSize
			if end > len(updates) {
				end = len(updates)
			}
			chunks = append(chunks, updates[i:end])
		}

		for _, chunk := range chunks {
			res, err := resourcesListColl.BulkWrite(ctx, chunk, &opts)
			results = append(results, *res)
			lib.CheckFatalError(err)
		}
		return results
	} else {
		results := []mongo.BulkWriteResult{}
		return results
	}

	// fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
	// fmt.Printf("Number of documents matched: %d\n", results.MatchedCount)
	// fmt.Printf("Number of documents matched: %d\n", )
	// fmt.Printf("Number of documents inserted: %d\n", results.UpsertedCount)
	// fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
	// fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
	// fmt.Println("Upserted IDs:")
	// jsonStr, _ := json.MarshalIndent(results.UpsertedIDs, "", "  ")
	// fmt.Println(string(jsonStr))

}

//
//

func UpsertMultipleResGrps(resGrps []azure.ResourceGroup, resourcesListColl *mongo.Collection) *mongo.BulkWriteResult {
	ctx := context.TODO()

	var updates []mongo.WriteModel

	for _, res := range resGrps {
		rg := res
		rg.LastDBSync = time.Now()
		rg.ID = strings.ToLower(res.ID)
		filter := bson.D{{"_id", rg.ID}}
		update := bson.D{{"$set", rg}}

		// .SetUpsert(true)
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	results, err := resourcesListColl.BulkWrite(ctx, updates)
	lib.CheckFatalError(err)

	// fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
	// fmt.Printf("Number of documents matched: %d\n", results.MatchedCount)
	// fmt.Printf("Number of documents matched: %d\n", )
	// fmt.Printf("Number of documents inserted: %d\n", results.UpsertedCount)
	// fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
	// fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
	// fmt.Println("Upserted IDs:")
	// jsonStr, _ := json.MarshalIndent(results.UpsertedIDs, "", "  ")
	// fmt.Println(string(jsonStr))
	return results
}

//
//

func UpsertMultipleRoleAssignmentScheduleInstances(ras []azure.RoleAssignmentScheduleInstance, coll *mongo.Collection) *mongo.BulkWriteResult {
	ctx := context.TODO()

	var updates []mongo.WriteModel

	for _, ra := range ras {
		curr := ra
		curr.LastDBSync = time.Now()
		curr.ID = strings.ToLower(ra.ID)
		filter := bson.D{{"_id", curr.ID}}
		update := bson.D{{"$set", curr}}

		// .SetUpsert(true)
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	results, err := coll.BulkWrite(ctx, updates)
	lib.CheckFatalError(err)

	// fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
	// fmt.Printf("Number of documents matched: %d\n", results.MatchedCount)
	// fmt.Printf("Number of documents matched: %d\n", )
	// fmt.Printf("Number of documents inserted: %d\n", results.UpsertedCount)
	// fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
	// fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
	// fmt.Println("Upserted IDs:")
	// jsonStr, _ := json.MarshalIndent(results.UpsertedIDs, "", "  ")
	// fmt.Println(string(jsonStr))
	return results
}

//
//

func UpsertMultipleRoleEligibilityScheduleInstances(res []azure.RoleEligibilityScheduleInstance, coll *mongo.Collection) *mongo.BulkWriteResult {
	ctx := context.TODO()

	var updates []mongo.WriteModel

	for _, re := range res {
		curr := re
		curr.LastDBSync = time.Now()
		curr.ID = strings.ToLower(re.ID)
		filter := bson.D{{"_id", curr.ID}}
		update := bson.D{{"$set", curr}}

		// .SetUpsert(true)
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	results, err := coll.BulkWrite(ctx, updates)
	lib.CheckFatalError(err)

	// fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
	// fmt.Printf("Number of documents matched: %d\n", results.MatchedCount)
	// fmt.Printf("Number of documents matched: %d\n", )
	// fmt.Printf("Number of documents inserted: %d\n", results.UpsertedCount)
	// fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
	// fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
	// fmt.Println("Upserted IDs:")
	// jsonStr, _ := json.MarshalIndent(results.UpsertedIDs, "", "  ")
	// fmt.Println(string(jsonStr))
	return results
}

//
//

func UpsertCertInfo(caCertInfo []lib.CertAuthorityCertInfo, serverCertInfo []lib.ServerCertInfo, caCertsColl *mongo.Collection, serverCertsColl *mongo.Collection) (caResults *mongo.BulkWriteResult, svrResults *mongo.BulkWriteResult) {
	caCtx := context.TODO()

	var caUpdates []mongo.WriteModel

	for _, cert := range caCertInfo {
		curr := cert
		currTime := time.Now()
		curr.LastDBSync = &currTime
		filter := bson.D{{"_id", curr.ID}}
		update := bson.D{{"$set", curr}}

		caUpdates = append(caUpdates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	caResults, err := caCertsColl.BulkWrite(caCtx, caUpdates)
	lib.CheckFatalError(err)

	svrCtx := context.TODO()

	var svrUpdates []mongo.WriteModel

	for _, cert := range serverCertInfo {
		curr := cert
		currTime := time.Now()
		curr.LastDBSync = &currTime
		filter := bson.D{{"_id", curr.ID}}
		update := bson.D{{"$set", curr}}

		svrUpdates = append(svrUpdates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	svrResults, err = serverCertsColl.BulkWrite(svrCtx, svrUpdates)
	lib.CheckFatalError(err)

	// fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
	// fmt.Printf("Number of documents matched: %d\n", results.MatchedCount)
	// fmt.Printf("Number of documents matched: %d\n", )
	// fmt.Printf("Number of documents inserted: %d\n", results.UpsertedCount)
	// fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
	// fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
	// fmt.Println("Upserted IDs:")
	// jsonStr, _ := json.MarshalIndent(results.UpsertedIDs, "", "  ")
	// fmt.Println(string(jsonStr))
	return
}

//
//

func UpsertAzureIPAddresses(resources []azure.IPAddressesAllResourceTypes, ipamIpAddressesColl *mongo.Collection) *mongo.BulkWriteResult {
	ctx := context.TODO()

	DeleteAllDocumentsInCollection(ipamIpAddressesColl)

	var updates []mongo.WriteModel

	for _, res := range resources {
		resource := res
		resource.LastDBSync = time.Now()
		resource.ID = strings.ToLower(res.ID)

		updates = append(updates, mongo.NewInsertOneModel().SetDocument(resource))
	}

	results, err := ipamIpAddressesColl.BulkWrite(ctx, updates)
	lib.CheckFatalError(err)

	return results
}

//
//

func UpdateIpamAddressBlocks(ipAddressBlocks []azure.IpAddressBlocksByBlockTag, ipamIpAddressBlocksColl *mongo.Collection) *mongo.BulkWriteResult {
	ctx := context.TODO()

	DeleteAllDocumentsInCollection(ipamIpAddressBlocksColl)

	var inserts []mongo.WriteModel

	for _, ipBlock := range ipAddressBlocks {
		inserts = append(inserts, mongo.NewInsertOneModel().SetDocument(ipBlock))
	}

	results, err := ipamIpAddressBlocksColl.BulkWrite(ctx, inserts)
	lib.CheckFatalError(err)

	return results
}

//
//

func UpsertServerCertificates(serverCertInfo []lib.ServerCertInfo, coll *mongo.Collection) (results []mongo.BulkWriteResult) {
	if len(serverCertInfo) == 0 {
		fmt.Println("No apps in slice")
		return nil
	}
	ctx := context.TODO()

	var updates []mongo.WriteModel

	// fmt.Println(len(serverCertInfo))

	for _, cert := range serverCertInfo {
		curr := cert

		timeNow := time.Now()
		curr.LastDBSync = &timeNow
		filter := bson.D{{"_id", curr.SerialNumber}}
		update := bson.D{{"$set", curr}}

		// .SetUpsert(true)
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	var opts options.BulkWriteOptions
	opts.SetOrdered(false)

	chunkSize := 100
	var chunks [][]mongo.WriteModel
	for i := 0; i < len(updates); i += chunkSize {
		end := i + chunkSize
		if end > len(updates) {
			end = len(updates)
		}
		chunks = append(chunks, updates[i:end])
	}

	for _, chunk := range chunks {
		res, err := coll.BulkWrite(ctx, chunk, &opts)
		results = append(results, *res)
		lib.CheckFatalError(err)
	}

	// fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
	// fmt.Printf("Number of documents matched: %d\n", results.MatchedCount)
	// fmt.Printf("Number of documents matched: %d\n", )
	// fmt.Printf("Number of documents inserted: %d\n", results.UpsertedCount)
	// fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
	// fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
	// fmt.Println("Upserted IDs:")
	// jsonStr, _ := json.MarshalIndent(results.UpsertedIDs, "", "  ")
	// fmt.Println(string(jsonStr))
	return results
}

func UpsertCACertificates(caCertInfo []lib.CertAuthorityCertInfo, coll *mongo.Collection) (results []mongo.BulkWriteResult) {
	if len(caCertInfo) == 0 {
		fmt.Println("No apps in slice")
		return nil
	}
	ctx := context.TODO()

	var updates []mongo.WriteModel

	// fmt.Println(len(caCertInfo))

	for _, cert := range caCertInfo {
		curr := cert

		timeNow := time.Now()
		curr.LastDBSync = &timeNow
		filter := bson.D{{"_id", curr.SerialNumber}}
		update := bson.D{{"$set", curr}}

		// .SetUpsert(true)
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	var opts options.BulkWriteOptions
	opts.SetOrdered(false)

	chunkSize := 100
	var chunks [][]mongo.WriteModel
	for i := 0; i < len(updates); i += chunkSize {
		end := i + chunkSize
		if end > len(updates) {
			end = len(updates)
		}
		chunks = append(chunks, updates[i:end])
	}

	for _, chunk := range chunks {
		res, err := coll.BulkWrite(ctx, chunk, &opts)
		results = append(results, *res)
		lib.CheckFatalError(err)
	}

	// fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
	// fmt.Printf("Number of documents matched: %d\n", results.MatchedCount)
	// fmt.Printf("Number of documents matched: %d\n", )
	// fmt.Printf("Number of documents inserted: %d\n", results.UpsertedCount)
	// fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
	// fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
	// fmt.Println("Upserted IDs:")
	// jsonStr, _ := json.MarshalIndent(results.UpsertedIDs, "", "  ")
	// fmt.Println(string(jsonStr))
	return results
}

// func UpsertMultipleResources(resources []lib.AzureResourceDetails, resourcesListColl *mongo.Collection) {
func UpsertStorageAccountMinTlsVersions(resources []azure.StorageAccountTlsVersion, mongoColl *mongo.Collection) *mongo.BulkWriteResult {
	ctx := context.TODO()

	var updates []mongo.WriteModel

	for _, res := range resources {
		resource := res

		resource.LastDBSync = time.Now()
		resource.ID = strings.ToLower(res.ID)
		filter := bson.D{{"_id", resource.ID}}
		update := bson.D{{"$set", resource}}

		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))

	}

	if len(updates) > 0 {
		results, err := mongoColl.BulkWrite(ctx, updates)
		lib.CheckFatalError(err)
		return results
	} else {
		results := mongo.BulkWriteResult{}
		return &results
	}
}

//
//

func UpsertB2CUsers(users []azure.B2CUserMinimal, coll *mongo.Collection) (results *mongo.BulkWriteResult) {
	if len(users) == 0 {
		fmt.Println("No users in slice")
		return nil
	}
	ctx := context.TODO()

	var updates []mongo.WriteModel

	for _, user := range users {
		curr := user

		timeNow := time.Now()
		curr.LastDBSync = timeNow

		filter := bson.D{{"_id", curr.B2CTenant + "-" + curr.ID}}
		update := bson.D{{"$set", curr}}

		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	// var opts options.BulkWriteOptions
	// opts.SetOrdered(false)

	results, err := coll.BulkWrite(ctx, updates)
	lib.CheckFatalError(err)

	return
}

//
//

func UpsertADUsers(users []ad.ADUser, coll *mongo.Collection) (results *mongo.BulkWriteResult) {
	if len(users) == 0 {
		fmt.Println("No users in slice")
		return nil
	}
	ctx := context.TODO()

	var updates []mongo.WriteModel

	for _, cert := range users {
		curr := cert

		timeNow := time.Now()
		curr.LastDBSync = timeNow
		filter := bson.D{{"_id", curr.UserPrincipalName}}
		update := bson.D{{"$set", curr}}

		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	var opts options.BulkWriteOptions
	opts.SetOrdered(false)

	results, err := coll.BulkWrite(ctx, updates, &opts)
	lib.CheckFatalError(err)

	return
}

//
//

func UpsertMailboxStatistics(mailboxStats []m365.MailboxUsageDetail, coll *mongo.Collection) (results *mongo.BulkWriteResult) {
	if len(mailboxStats) == 0 {
		fmt.Println("No data in slice")
		return nil
	}
	ctx := context.TODO()

	var updates []mongo.WriteModel

	// fmt.Println(len(serverCertInfo))

	for _, cert := range mailboxStats {
		curr := cert

		filter := bson.D{{"_id", curr.UserPrincipalName}}
		update := bson.D{{"$set", curr}}

		// .SetUpsert(true)
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	res, err := coll.BulkWrite(ctx, updates, nil)
	lib.CheckFatalError(err)

	results = res

	return
}

//
//

func UpsertSupportAlerts(alerts []azure.AzureAlertProcessed, coll *mongo.Collection) (results *mongo.BulkWriteResult) {
	if len(alerts) == 0 {
		fmt.Println("No data in slice")
		return nil
	}
	ctx := context.TODO()

	var updates []mongo.WriteModel

	// fmt.Println(len(serverCertInfo))

	for _, alert := range alerts {
		curr := alert

		filter := bson.D{{"_id", curr.ID}}
		update := bson.D{{"$set", curr}}

		// .SetUpsert(true)
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	res, err := coll.BulkWrite(ctx, updates, nil)
	lib.CheckFatalError(err)

	results = res

	return
}

//
//

func UpsertAWSMontoringData(data interface{}, coll *mongo.Collection) {
	jsonStr, _ := json.Marshal(data)
	var props struct {
		ID          string `json:"id"`
		Environment string `json:"environment"`
		Monitor     string `json:"monitor"`
	}
	err := json.Unmarshal(jsonStr, &props)

	filter := bson.D{{"_id", props.ID}}
	update := bson.D{{"$set", data}}
	mdbOpts := options.Update().SetUpsert(true)
	_, err = coll.UpdateOne(context.TODO(), filter, update, mdbOpts)
	lib.CheckFatalError(err)
}
