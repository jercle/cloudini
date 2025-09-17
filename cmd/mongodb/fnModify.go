package mongodb

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/jercle/cloudini/cmd/citrix"
	"github.com/jercle/cloudini/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func MarkImageGalleryImagesUsedByCitrix(machineCatalogs map[string]citrix.MachineCatalogCurrentImage, collection *mongo.Collection) *mongo.BulkWriteResult {

	imgVersionsUsed := make(map[string]bool)
	var imgVersionsUsedSlice []string
	mapImageToMachineCat := make(map[string][]string)

	for mcName, mcImg := range machineCatalogs {
		if mcImg.IsPreparedImage {
			continue
		}
		matchString := strings.ToLower(mcImg.ImageGallery + "_" + mcImg.ImageDefinitionName)
		matchStringWithVersion := strings.ToLower(mcImg.ImageGallery + "_" + mcImg.ImageDefinitionName + "_" + mcImg.Version)
		imgVersionsUsed[matchStringWithVersion] = true
		mapImageToMachineCat[matchString] = append(mapImageToMachineCat[matchString], mcName)
		mapImageToMachineCat[matchStringWithVersion] = append(mapImageToMachineCat[matchStringWithVersion], mcName)
	}
	for key, _ := range imgVersionsUsed {
		imgVersionsUsedSlice = append(imgVersionsUsedSlice, key)
	}

	ctx := context.TODO()

	var updates []mongo.WriteModel

	for _, img := range machineCatalogs {
		if img.IsPreparedImage {
			continue
		}

		ctx := context.TODO()

		filter := bson.D{
			{"name", img.ImageDefinitionName},
			{"galleryName", img.ImageGallery},
		}

		var galleryImage lib.GalleryImage

		err := collection.FindOne(ctx, filter).Decode(&galleryImage)
		if err != nil {
			lib.JsonMarshalAndPrint(filter)
		}
		lib.CheckFatalError(err)
		currImg := galleryImage

		currImg.UsedByCitrix = false
		for vName, vData := range galleryImage.ImageVersions {
			currVersion := vData
			currVersion.LastDBSync = time.Now()
			matchString := strings.ToLower(galleryImage.GalleryName + "_" + galleryImage.Name)
			matchStringWithVersion := strings.ToLower(galleryImage.GalleryName + "_" + galleryImage.Name + "_" + vData.Name)
			if slices.Contains(imgVersionsUsedSlice, matchStringWithVersion) {
				currVersion.UsedByCitrix = true
				currImg.UsedByCitrix = true
				currImg.MachineCatalogsUsingImage = mapImageToMachineCat[matchString]
				currVersion.MachineCatalogsUsingImage = mapImageToMachineCat[matchStringWithVersion]
			} else {
				currVersion.UsedByCitrix = false
			}
			currImg.ImageVersions[vName] = currVersion
		}
		// jsonStr, _ := json.MarshalIndent(galleryImage, "", "  ")
		// // fmt.Println(string(jsonStr))
		// os.WriteFile("outputs/images/singleImageWithVersions-updated.json", jsonStr, 0644)
		// os.Exit(0)
		update := bson.D{{"$set", currImg}}
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	results, err := collection.BulkWrite(ctx, updates)
	lib.CheckFatalError(err)

	// fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
	// fmt.Printf("Number of documents matched: %d\n", results.MatchedCount)
	// // fmt.Printf("Number of documents matched: %d\n", )
	// fmt.Printf("Number of documents inserted: %d\n", results.UpsertedCount)
	// fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
	// fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
	// fmt.Println("Upserted IDs:")
	return results
}

// Gets all documents from a collection then
// calculates monthly totals from each costData item and saves
// the sum back to the cost item's "lifetimeTotalCost" field.
// Only updates documents if there is a difference
func ResetLifetimeCostValues(collection *mongo.Collection) *mongo.BulkWriteResult {
	var (
		allCostData []lib.MongoDbCostItem
		updates     []mongo.WriteModel
	)

	ctx := context.TODO()

	filter := bson.D{{}}

	rsp, err := collection.Find(ctx, filter)
	lib.CheckFatalError(err)

	err = rsp.All(ctx, &allCostData)
	lib.CheckFatalError(err)

	// fmt.Println(allCostData)
	// os.Exit(0)

	for _, costItem := range allCostData {
		// if costItem.LifetimeTotalCost == 0 {
		// 	continue
		// }
		var lifetimeTotalCost float64
		for _, monthData := range costItem.CostData {
			lifetimeTotalCost += monthData.MonthTotalCost
		}

		// fmt.Println(costItem.LifetimeTotalCost)
		// fmt.Println(lifetimeTotalCost)
		// os.Exit(0)

		if costItem.LifetimeTotalCost != lifetimeTotalCost {
			filter := bson.D{{"_id", costItem.ResourceMeterIdentifier}}
			update := bson.D{{"$set", bson.D{{"lifetimeTotalCost", lifetimeTotalCost}}}}
			updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update))
		}
	}

	results, err := collection.BulkWrite(ctx, updates)
	if err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, "must provide at least one element") {
			fmt.Println("No data to update")
			return nil
		} else {
			lib.CheckFatalError(err)
		}
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

func ResetRelatedResourcesAndCostItems(collection *mongo.Collection) *mongo.BulkWriteResult {
	var (
		allResources []lib.AzureResourceDetails
		updates      []mongo.WriteModel
	)

	ctx := context.TODO()

	filter := bson.D{{}}

	rsp, err := collection.Find(ctx, filter)
	lib.CheckFatalError(err)

	err = rsp.All(ctx, &allResources)
	lib.CheckFatalError(err)

	for _, res := range allResources {
		curr := res

		relCm := make(map[string]bool)
		for _, id := range res.RelatedCostMeters {
			relCm[id] = true
		}
		var relCmSlice []string
		for id, _ := range relCm {
			relCmSlice = append(relCmSlice, id)
		}
		curr.RelatedCostMeters = relCmSlice

		relRes := make(map[string]bool)
		for _, id := range res.RelatedResources {
			relRes[id] = true
		}
		var relResSlice []string
		for id, _ := range relCm {
			relResSlice = append(relResSlice, id)
		}
		curr.RelatedResources = relResSlice
		filter := bson.D{{"_id", curr.ID}}
		update := bson.D{{"$set", curr}}
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update))
	}

	// fmt.Println(allCostData)
	// os.Exit(0)

	// for _, costItem := range allCostData {
	// 	// if costItem.LifetimeTotalCost == 0 {
	// 	// 	continue
	// 	// }
	// 	var lifetimeTotalCost float64
	// 	for _, monthData := range costItem.CostData {
	// 		lifetimeTotalCost += monthData.MonthTotalCost
	// 	}

	// 	// fmt.Println(costItem.LifetimeTotalCost)
	// 	// fmt.Println(lifetimeTotalCost)
	// 	// os.Exit(0)

	// 	if costItem.LifetimeTotalCost != lifetimeTotalCost {
	// 		filter := bson.D{{"_id", costItem.ResourceMeterIdentifier}}
	// 		update := bson.D{{"$set", bson.D{{"lifetimeTotalCost", lifetimeTotalCost}}}}
	// 		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update))
	// 	}
	// }

	results, err := collection.BulkWrite(ctx, updates)
	if err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, "must provide at least one element") {
			fmt.Println("No data to update")
			return nil
		} else {
			lib.CheckFatalError(err)
		}
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

func UpdateImageDataWithBuildHostLogs(buildData []lib.PackerLogBuildData, collection *mongo.Collection) *mongo.BulkWriteResult {
	var (
		images  []lib.GalleryImage
		updates []mongo.WriteModel
	)
	ctx := context.TODO()
	// ctx := context.TODO()
	filter := bson.D{{}}

	rsp, err := collection.Find(ctx, filter)
	lib.CheckFatalError(err)

	err = rsp.All(ctx, &images)
	lib.CheckFatalError(err)

	imagesById := make(map[string]lib.GalleryImage)
	for _, img := range images {
		imagesById[img.ID] = img
	}

	// jsonStr, _ := json.MarshalIndent(imagesById, "", "  ")
	// fmt.Println(string(jsonStr))
	// os.Exit(0)

	for _, data := range buildData {

		imageId := strings.Split(data.OutputImgId, "/versions")[0]

		if _, ok := imagesById[imageId].ImageVersions[data.OutputImgVersion]; ok {
			currVersion := imagesById[imageId].ImageVersions[data.OutputImgVersion]
			currVersion.AzDoBuildData = data
			imagesById[imageId].ImageVersions[data.OutputImgVersion] = currVersion
		}

		// if _, ok := imagesById[data.OutputImgId]; !ok {
		// 	continue
		// // }
		// if _, ok := imagesById[data.OutputImgId].ImageVersions[data.OutputImgVersion]; !ok {
		// 	continue
		// }

	}

	for _, img := range imagesById {
		filter := bson.D{{"_id", img.ID}}
		update := bson.D{{"$set", img}}
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))

	}

	results, err := collection.BulkWrite(ctx, updates)
	lib.CheckFatalError(err)

	// fmt.Printf("Number of documents inserted: %d\n", results.InsertedCount)
	// fmt.Printf("Number of documents matched: %d\n", results.MatchedCount)
	// // fmt.Printf("Number of documents matched: %d\n", )
	// fmt.Printf("Number of documents inserted: %d\n", results.UpsertedCount)
	// fmt.Printf("Number of documents replaced or updated: %d\n", results.ModifiedCount)
	// fmt.Printf("Number of documents deleted: %d\n", results.DeletedCount)
	// fmt.Println("Upserted IDs:")
	return results
}

//
//

func UpdateResourcesNotExistInAzure(azureResources []lib.AzureResourceDetails, collection *mongo.Collection) *mongo.BulkWriteResult {
	var (
		allDbRes []lib.AzureResourceDetails
		updates  []mongo.WriteModel
	)
	ctx := context.TODO()
	filter := bson.D{{}}
	rsp, err := collection.Find(ctx, filter)
	lib.CheckFatalError(err)

	err = rsp.All(ctx, &allDbRes)
	lib.CheckFatalError(err)
	currentDbResMap := make(map[string]bool)
	for _, res := range azureResources {
		currentDbResMap[res.ID] = true
	}

	for _, res := range allDbRes {
		curr := res
		if _, ok := currentDbResMap[res.ID]; !ok {
			curr.ExistsInAzure = false
		} else {
			curr.ExistsInAzure = true
		}

		// curr.Properties.Other = ""

		filter := bson.D{{"_id", res.ID}}
		update := bson.D{{"$set", curr}}
		updates = append(updates, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true))
	}

	results, err := collection.BulkWrite(ctx, updates)
	lib.CheckFatalError(err)

	return results
}

func DeleteAllDocumentsInCollection(collection *mongo.Collection) (results *mongo.DeleteResult) {
	filter := bson.D{{}}
	results, err := collection.DeleteMany(context.TODO(), filter)
	lib.CheckFatalError(err)

	return
}

func UnsetField(fieldName string, collection *mongo.Collection) *mongo.UpdateResult {
	filter := bson.D{{}}
	update := bson.D{{"$unset", bson.D{{fieldName, 1}}}}
	results, err := collection.UpdateMany(context.TODO(), filter, update)
	lib.CheckFatalError(err)
	lib.JsonMarshalAndPrint(results)
	return results
}
