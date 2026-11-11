package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jercle/cloudini/cmd/ad"
	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/cmd/citrix"
	"github.com/jercle/cloudini/cmd/forgerock"
	"github.com/jercle/cloudini/cmd/m365"
	"github.com/jercle/cloudini/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

		// Subscription details are fetched a few at a time rather than one by one
		var (
			wg    sync.WaitGroup
			mutex sync.Mutex
			limit = make(chan struct{}, 8)
		)
		for _, sId := range tData.Subscriptions {
			wg.Add(1)
			limit <- struct{}{}
			go func() {
				defer wg.Done()
				defer func() { <-limit }()

				var currSub MongoDbAzureSubscription
				res, err := azure.HttpGet("https://management.azure.com/subscriptions/"+sId+"?api-version=2022-12-01", *token)
				lib.CheckFatalError(err)

				err = json.Unmarshal(res, &currSub)
				lib.CheckFatalError(err)

				mutex.Lock()
				currTenant.Subscriptions[strings.ToLower(currSub.DisplayName)] = currSub
				currTenant.TenantId = currSub.TenantId
				mutex.Unlock()
			}()
		}
		wg.Wait()

		filter := bson.D{{"_id", currTenant.TenantId}}
		update := bson.D{{"$set", currTenant}}
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))

		// jsonStr, _ := json.MarshalIndent(r, "", "  ")
		// fmt.Println(string(jsonStr))
		// UpsertResource()
	}
	allSubs = nil
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

func UpsertCitrixMachineCatalogs(tenantName string, machineCatalogs citrix.MachineCatalogs, coll *mongo.Collection) (results []mongo.BulkWriteResult) {
	if len(machineCatalogs) == 0 {
		fmt.Println("No apps in slice")
		return nil
	}
	ctx := context.TODO()

	var updates []mongo.WriteModel

	for _, mc := range machineCatalogs {
		curr := mc
		curr.TenantName = tenantName
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

func UpsertCitrixPolicySettingDefs(settingDefs []citrix.PolicySettingDefinition, coll *mongo.Collection) (results *mongo.BulkWriteResult) {
	if len(settingDefs) == 0 {
		fmt.Println("No data in slice")
		return nil
	}
	ctx := context.TODO()

	var updates []mongo.WriteModel

	for _, sd := range settingDefs {
		curr := sd
		curr.LastDBSync = time.Now()
		filter := bson.D{{"_id", sd.SettingName}}
		update := bson.D{{"$set", curr}}

		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	results, err := coll.BulkWrite(ctx, updates)
	lib.CheckFatalError(err)

	return results
}

//
//

func UpsertMultipleEntraApps[T azure.EntraApplication | azure.EntraExpiringCredential](apps []T, collection *mongo.Collection) *mongo.BulkWriteResult {
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

	// var opts options.BulkWriteOptions
	// opts.SetOrdered(false)

	// chunkSize := 100
	// var chunks [][]mongo.WriteModel
	// for i := 0; i < len(updates); i += chunkSize {
	// 	end := i + chunkSize
	// 	if end > len(updates) {
	// 		end = len(updates)
	// 	}
	// 	chunks = append(chunks, updates[i:end])
	// }

	results, err := collection.BulkWrite(ctx, updates, nil)
	lib.CheckFatalError(err)

	// for _, chunk := range chunks {
	// 	res, err := collection.BulkWrite(ctx, chunk, &opts)
	// 	results = append(results, *res)
	// 	lib.CheckFatalError(err)
	// }

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
func UpsertMultipleResources(resources *[]lib.AzureResourceDetails, resourcesListColl *mongo.Collection) (results []mongo.BulkWriteResult) {
	ctx := context.TODO()

	var updates []mongo.WriteModel

	for _, res := range *resources {
		resource := res
		resource.LastDBSync = time.Now()
		resource.ID = strings.ToLower(res.ID)
		filter := bson.D{{"_id", resource.ID}}
		update := bson.D{{"$set", resource}}
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	resources = nil

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

}

//
//

func UpsertMultipleResGrps(resGrps *[]azure.ResourceGroup, resourcesListColl *mongo.Collection) *mongo.BulkWriteResult {
	ctx := context.TODO()

	var updates []mongo.WriteModel

	for _, res := range *resGrps {
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

func UpsertServerCertificatesNew(serverCertInfo []lib.FormattedServerCertInfo, coll *mongo.Collection) (results *mongo.BulkWriteResult) {
	if len(serverCertInfo) == 0 {
		fmt.Println("No apps in slice")
		return nil
	}
	ctx := context.TODO()
	ctx2 := context.TODO()

	var updates []mongo.WriteModel

	// fmt.Println(len(serverCertInfo))

	for _, cert := range serverCertInfo {
		curr := cert

		timeNow := time.Now()
		curr.LastDBSync = &timeNow
		curr.ExistsInEnvironment = true
		filter := bson.D{{Key: "_id", Value: curr.ID}}
		update := bson.D{{Key: "$set", Value: curr}}

		// .SetUpsert(true)
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
		// updates = append(updates, mongo.NewInsertOneModel().SetDocument(curr))
	}

	var currIds []string
	for _, c := range serverCertInfo {
		currIds = append(currIds, c.ID)
	}

	notCurrentFilter := bson.M{
		"_id": bson.M{
			"$nin": currIds,
		},
		"existsInEnvironment": true,
	}
	notCurrentUpdate := bson.D{{Key: "$set", Value: bson.M{"existsInEnvironment": false}}}
	coll.UpdateMany(ctx2, notCurrentFilter, notCurrentUpdate)

	results, err := coll.BulkWrite(ctx, updates, nil)
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

// func UpsertServerCertificates(serverCertInfo []lib.ServerCertInfo, coll *mongo.Collection) (results []mongo.BulkWriteResult) {
// 	if len(serverCertInfo) == 0 {
// 		fmt.Println("No apps in slice")
// 		return nil
// 	}
// 	ctx := context.TODO()

// 	var updates []mongo.WriteModel

// 	// fmt.Println(len(serverCertInfo))

// 	for _, cert := range serverCertInfo {
// 		curr := cert

// 		timeNow := time.Now()
// 		curr.LastDBSync = &timeNow
// 		filter := bson.D{{"_id", curr.SerialNumber}}
// 		update := bson.D{{"$set", curr}}

// 		// .SetUpsert(true)
// 		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
// 	}

// 	var opts options.BulkWriteOptions
// 	opts.SetOrdered(false)

// 	chunkSize := 100
// 	var chunks [][]mongo.WriteModel
// 	for i := 0; i < len(updates); i += chunkSize {
// 		end := i + chunkSize
// 		if end > len(updates) {
// 			end = len(updates)
// 		}
// 		chunks = append(chunks, updates[i:end])
// 	}

// 	for _, chunk := range chunks {
// 		res, err := coll.BulkWrite(ctx, chunk, &opts)
// 		results = append(results, *res)
// 		lib.CheckFatalError(err)
// 	}

// 	// fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
// 	// fmt.Printf("Number of documents matched: %d\n", results.MatchedCount)
// 	// fmt.Printf("Number of documents matched: %d\n", )
// 	// fmt.Printf("Number of documents inserted: %d\n", results.UpsertedCount)
// 	// fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
// 	// fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
// 	// fmt.Println("Upserted IDs:")
// 	// jsonStr, _ := json.MarshalIndent(results.UpsertedIDs, "", "  ")
// 	// fmt.Println(string(jsonStr))
// 	return results
// }

//
//

func UpsertCACertificates(caCertInfo []lib.CertAuthorityCertInfo, coll *mongo.Collection) (results *mongo.BulkWriteResult) {
	if len(caCertInfo) == 0 {
		fmt.Println("No apps in slice")
		return nil
	}
	ctx := context.TODO()
	ctx2 := context.TODO()

	var updates []mongo.WriteModel

	// fmt.Println(len(caCertInfo))

	for _, cert := range caCertInfo {
		curr := cert

		timeNow := time.Now()
		curr.LastDBSync = &timeNow
		curr.ExistsInEnvironment = true
		filter := bson.D{{"_id", curr.SerialNumber}}
		update := bson.D{{"$set", curr}}

		// .SetUpsert(true)
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	var currIds []string
	for _, c := range caCertInfo {
		currIds = append(currIds, c.ID)
	}

	notCurrentFilter := bson.M{
		"_id": bson.M{
			"$nin": currIds,
		},
		"existsInEnvironment": true,
	}
	notCurrentUpdate := bson.D{{Key: "$set", Value: bson.M{"existsInEnvironment": false}}}
	coll.UpdateMany(ctx2, notCurrentFilter, notCurrentUpdate)

	results, err := coll.BulkWrite(ctx, updates, nil)
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

func UpsertMailboxes(mailboxes []m365.EXOMailBox, coll *mongo.Collection) (results *mongo.BulkWriteResult) {
	if len(mailboxes) == 0 {
		fmt.Println("No data in slice")
		return nil
	}
	ctx := context.TODO()

	var updates []mongo.WriteModel

	// fmt.Println(len(serverCertInfo))

	for _, cert := range mailboxes {
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

//
//

func InsertM365LicenseCounts(licenseCounts map[string]m365.M365LicenseCounts, coll *mongo.Collection) (results *mongo.BulkWriteResult) {
	ctx := context.TODO()

	var updates []mongo.WriteModel

	// fmt.Println(len(serverCertInfo))

	for _, tData := range licenseCounts {
		curr := tData

		// filter := bson.D{{"_id", curr.ID}}
		// update := bson.D{{"$set", curr}}

		// .SetUpsert(true)
		updates = append(updates, mongo.NewInsertOneModel().SetDocument(curr))
	}

	res, err := coll.BulkWrite(ctx, updates, nil)
	lib.CheckFatalError(err)

	results = res

	return
}

//
//

func InsertIntuneManagedDevices(managedDevices []azure.ManagedDevice, coll *mongo.Collection) (results *mongo.BulkWriteResult) {
	if len(managedDevices) == 0 {
		fmt.Println("No data in slice")
		return nil
	}
	ctx := context.TODO()

	var updates []mongo.WriteModel

	// fmt.Println(len(serverCertInfo))

	for _, md := range managedDevices {
		// filter := bson.D{{"_id", md.ID}}
		// update := bson.D{{"$set", md}}

		// .SetUpsert(true)
		// updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
		updates = append(updates, mongo.NewInsertOneModel().SetDocument(md))
	}

	res, err := coll.BulkWrite(ctx, updates, nil)
	lib.CheckFatalError(err)

	return res
}

//
//

func InsertFrADGroups(groups []forgerock.LDAPConnectorGroup, coll *mongo.Collection) (results *mongo.BulkWriteResult) {
	if len(groups) == 0 {
		fmt.Println("No data in slice")
		return nil
	}
	ctx := context.TODO()

	var updates []mongo.WriteModel

	// fmt.Println(len(serverCertInfo))

	for _, grp := range groups {
		// filter := bson.D{{"_id", grp.ID}}
		// update := bson.D{{"$set", grp}}
		curr := grp
		curr.LastDbSync = time.Now()

		// .SetUpsert(true)
		// updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
		updates = append(updates, mongo.NewInsertOneModel().SetDocument(curr))
	}

	res, err := coll.BulkWrite(ctx, updates, nil)
	lib.CheckFatalError(err)

	return res
}

//
//

func InsertCitrixMonitorData(machineData map[string]citrix.MonitorMachine, metricsData []citrix.MachineMetric, resUtil []citrix.MachineResourceUtilisation, loadIndexes []citrix.MachineLoadIndex, collMachines *mongo.Collection, collMetrics *mongo.Collection, collResUtil *mongo.Collection, collLoadIndexes *mongo.Collection) (resMachines *mongo.BulkWriteResult, resMetrics *mongo.BulkWriteResult, resResUtil *mongo.BulkWriteResult, resLoadIndexes *mongo.BulkWriteResult) {
	ctx := context.TODO()

	var machineUpdates []mongo.WriteModel
	for _, m := range machineData {

		if m.ID == "" {
			continue
		}
		// machineUpdates = append(machineUpdates, mongo.NewUpdateOneModel().SetDocument(m))
		filter := bson.D{{"_id", m.ID}}
		update := bson.D{{"$set", m}}
		machineUpdates = append(machineUpdates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}
	resMachines, err := collMachines.BulkWrite(ctx, machineUpdates, nil)
	// if err != nil {
	// 	lib.JsonMarshalAndPrint()
	// }
	lib.CheckFatalError(err)

	bulkWriteOpts := options.BulkWriteOptions{
		Ordered: new(false),
	}

	if len(metricsData) > 0 {
		var metricUpdates []mongo.WriteModel
		for _, m := range metricsData {
			// metricUpdates = append(metricUpdates, mongo.NewUpdateOneModel().SetDocument(m))
			// filter := bson.D{{"MachineId", m.MachineID}, {"CollectedDate", m.CollectedDate}, {"tenantName", m.TenantName}}
			// update := bson.D{{"$set", m}}
			// metricUpdates = append(metricUpdates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
			metricUpdates = append(metricUpdates, mongo.NewInsertOneModel().SetDocument(m))
		}
		resMetrics, err = collMetrics.BulkWrite(ctx, metricUpdates, &bulkWriteOpts)
		if err != nil {
			if !strings.Contains(err.Error(), "E11000 duplicate key") && !strings.Contains(err.Error(), "Duplicate key violation") {
				lib.CheckFatalError(err)
			}
		}
	}

	if len(resUtil) > 0 {
		var resUtilUpdates []mongo.WriteModel
		for _, m := range resUtil {
			// resUtilUpdates = append(resUtilUpdates, mongo.NewUpdateOneModel().SetDocument(m))
			// filter := bson.D{{"MachineId", m.MachineID}, {"CollectedDate", m.CollectedDate}, {"tenantName", m.TenantName}}
			// update := bson.D{{"$set", m}}
			// resUtilUpdates = append(resUtilUpdates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
			resUtilUpdates = append(resUtilUpdates, mongo.NewInsertOneModel().SetDocument(m))
		}
		resResUtil, err = collResUtil.BulkWrite(ctx, resUtilUpdates, &bulkWriteOpts)
		if err != nil {
			if !strings.Contains(err.Error(), "E11000 duplicate key") && !strings.Contains(err.Error(), "Duplicate key violation") {
				lib.CheckFatalError(err)
			}
		}
	}

	if len(loadIndexes) > 0 {
		var loadIndexesUpdates []mongo.WriteModel
		for _, m := range loadIndexes {
			curr := m
			curr.MongoID = m.TenantName + "-" + strconv.Itoa(int(m.ID))
			// loadIndexesUpdates = append(loadIndexesUpdates, mongo.NewUpdateOneModel().SetDocument(m))
			// filter := bson.D{{Key: "_id", Value: m.ID}}
			// filter := bson.D{{Key: "MachineId", Value: m.MachineID}, {Key: "CreatedDate", Value: m.CreatedDate}, {Key: "tenantName", Value: m.TenantName}}
			// update := bson.D{{Key: "$set", Value: m}}
			// loadIndexesUpdates = append(loadIndexesUpdates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
			loadIndexesUpdates = append(loadIndexesUpdates, mongo.NewInsertOneModel().SetDocument(curr))

		}
		resLoadIndexes, err = collLoadIndexes.BulkWrite(ctx, loadIndexesUpdates, &bulkWriteOpts)
		if err != nil {
			if !strings.Contains(err.Error(), "E11000 duplicate key") && !strings.Contains(err.Error(), "Duplicate key violation") {
				lib.CheckFatalError(err)
			}
		}
	}

	// return resMachines, resMetrics, resResUtil
	return
}

//
//

func UpsertKeyVaultSecrets(allSecrets []azure.KeyVaultSecretStored, coll *mongo.Collection) (results *mongo.BulkWriteResult) {
	ctx := context.TODO()

	var updates []mongo.WriteModel
	for _, s := range allSecrets {
		filter := bson.D{{"_id", s.Id}}
		update := bson.D{{"$set", s}}
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}
	results, err := coll.BulkWrite(ctx, updates, nil)
	lib.CheckFatalError(err)

	return
}

//
//

func UpsertResourceChanges(resourceChanges []azure.ResourceChange, coll *mongo.Collection) (results *mongo.BulkWriteResult) {
	ctx := context.TODO()
	if len(resourceChanges) == 0 {
		return
	}

	var updates []mongo.WriteModel
	for _, doc := range resourceChanges {
		filter := bson.D{{"_id", doc.ID}}
		update := bson.D{{"$set", doc}}
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
		// updates = append(updates, mongo.NewInsertOneModel().SetDocument(doc))

	}
	results, err := coll.BulkWrite(ctx, updates, nil)
	// lib.CheckFatalError(err)
	if err != nil {
		if !strings.Contains(err.Error(), "E11000 duplicate key") && !strings.Contains(err.Error(), "Duplicate key violation") {
			lib.CheckFatalError(err)
		}
	}

	return
}

//
//

func UpsertIdentityChanges(identityChanges []azure.IdentityChange, coll *mongo.Collection) (results *mongo.BulkWriteResult) {
	ctx := context.TODO()
	if len(identityChanges) == 0 {
		return
	}

	var updates []mongo.WriteModel
	for _, doc := range identityChanges {
		filter := bson.D{{"_id", doc.ReportID}}
		update := bson.D{{"$set", doc}}
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
		// updates = append(updates, mongo.NewInsertOneModel().SetDocument(doc))

	}
	results, err := coll.BulkWrite(ctx, updates, nil)
	// lib.CheckFatalError(err)
	if err != nil {
		if !strings.Contains(err.Error(), "E11000 duplicate key") && !strings.Contains(err.Error(), "Duplicate key violation") {
			lib.CheckFatalError(err)
		}
	}

	return
}
