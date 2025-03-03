package mongodb

import (
	"context"
	"log"

	"github.com/jercle/cloudini/lib"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// *mongo.Client

func ClientFromConfig(ctx context.Context, cancel context.CancelFunc) *mongo.Client {

	config := lib.GetCldConfig(nil)
	mongoConf := config.MongoDBConfig
	// mongoPwd := os.Getenv("MONGODB_PWD")
	// connectionString := "mongodb://automonadm:" + mongoPwd + "@localhost:27017/"

	clientOptions := options.Client().ApplyURI(mongoConf.ConnectionString).SetDirect(true)

	err := clientOptions.Validate()
	lib.CheckFatalError(err)
	c, err := mongo.Connect(ctx, clientOptions)
	lib.CheckFatalError(err)

	if err != nil {
		log.Fatalf("unable to initialize connection %v", err)
	}

	err = c.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("unable to connect %v", err)
	}
	return c
}
