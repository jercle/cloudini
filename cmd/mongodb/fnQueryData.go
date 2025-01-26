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
