package mongodb

import (
	"context"

	"github.com/jercle/cloudini/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllResources(coll *mongo.Collection) (documentData []lib.AzureResourceDetails) {

	ctx := context.Background()

	filter := bson.D{{}}

	rsp, err := coll.Find(ctx, filter)
	lib.CheckFatalError(err)

	err = rsp.All(ctx, &documentData)
	lib.CheckFatalError(err)

	return documentData
}

//
//

func GetResource(name, resourceGroupName, subscriptionName, tenantName string, coll *mongo.Collection) (resource lib.AzureResourceDetails) {

	filter := bson.D{
		{
			"name", name,
		}, {
			"tenantName", tenantName,
		}, {
			"resourceGroupName", resourceGroupName,
		}, {
			"subscriptionName", subscriptionName,
		},
	}

	err := coll.FindOne(context.TODO(), filter, nil).Decode(&resource)
	lib.CheckFatalError(err)

	// filter := bson.D{
	// 	{
	// 		"date_ordered", bson.D{{"$lte", time.Date(2009, 11, 30, 0, 0, 0, 0, time.Local)}}}}

	// rsp, err := coll.Find(ctx, filter)
	// lib.CheckFatalError(err)

	// err = rsp.All(ctx, &documentData)
	// lib.CheckFatalError(err)

	return
}

//
//

func GetResourceSKUs(collection *mongo.Collection) (skus []lib.AzureResourceSku) {
	ctx := context.Background()

	filter := bson.D{{}}

	rsp, err := collection.Find(ctx, filter)
	lib.CheckFatalError(err)

	err = rsp.All(ctx, &skus)
	lib.CheckFatalError(err)
	// if len(skus) == 0 {
	// 	fmt.Println("No apps in slice")
	// 	return nil
	// }
	// ctx := context.TODO()

	// var updates []mongo.WriteModel

	// for _, sku := range skus {
	// 	curr := sku
	// 	curr.LastDBSync = time.Now()
	// 	filter := bson.D{{"_id", sku.Name}}
	// 	update := bson.D{{"$set", curr}}

	// 	// .SetUpsert(true)
	// 	updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	// }

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

	// for _, chunk := range chunks {
	// 	res, err := collection.BulkWrite(ctx, chunk, &opts)
	// 	results = append(results, *res)
	// 	lib.CheckFatalError(err)
	// }

	// // fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
	// // fmt.Printf("Number of documents matched: %d\n", results.MatchedCount)
	// // fmt.Printf("Number of documents matched: %d\n", )
	// // fmt.Printf("Number of documents inserted: %d\n", results.UpsertedCount)
	// // fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
	// // fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
	// // fmt.Println("Upserted IDs:")
	// // jsonStr, _ := json.MarshalIndent(results.UpsertedIDs, "", "  ")
	// // fmt.Println(string(jsonStr))
	// return results
	return
}
