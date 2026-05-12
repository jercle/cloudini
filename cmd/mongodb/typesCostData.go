package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type UpsertMonthlyTenantSubResGrpCostsResults struct {
	UpdateTenants                   *mongo.BulkWriteResult
	UpdateTenantsCostData           *mongo.BulkWriteResult
	UpdateTenantsProcessedUpdates   *mongo.BulkWriteResult
	UpdateSubs                      *mongo.BulkWriteResult
	UpdateSubsCostData              *mongo.BulkWriteResult
	UpdateSubsProcessedUpdates      *mongo.BulkWriteResult
	UpdateResGrps                   *mongo.BulkWriteResult
	UpdateResGrpsCostData           *mongo.BulkWriteResult
	UpdateResGrpsProcessedUpdates   *mongo.BulkWriteResult
	UpdateResources                 *mongo.BulkWriteResult
	UpdateResourcesCostData         *mongo.BulkWriteResult
	UpdateResourcesProcessedUpdates *mongo.BulkWriteResult
	UpdateMeters                    *mongo.BulkWriteResult
	UpdateMetersCostData            *mongo.BulkWriteResult
	UpdateMetersProcessedUpdates    *mongo.BulkWriteResult
}

type MongoDbAzureSubscription struct {
	DisplayName    string `json:"displayName" bson:"displayName"`
	ResourceID     string `json:"id" bson:"resourceId"`
	SubscriptionID string `json:"subscriptionId" bson:"subscriptionId"`
	TenantId       string `json:"tenantId" bson:"-"`
}

type MongoDbAzureTenant struct {
	TenantId            string                              `json:"tenantId" bson:"_id"`
	TenantName          string                              `json:"tenantName" bson:"tenantName"`
	CostExportsLocation string                              `json:"costExportsLocation" bson:"costExportsLocation"`
	Subscriptions       map[string]MongoDbAzureSubscription `json:"subscriptions" bson:"subscriptions"`
	Aliases             []string                            `json:"aliases" bson:"aliases"`
}
