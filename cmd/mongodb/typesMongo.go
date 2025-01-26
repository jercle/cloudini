package mongodb

import (
	"github.com/jercle/cloudini/lib"
	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateAllAzureResourcesAndVcpuCountsOptionsOptions struct {
	AzureResourcesDatabase      string
	CitrixDatabase              string
	AzResImageGalleryImagesColl string
	CitrixMachineCatalogsColl   string
}

type UpdateAllAzureResourcesAndVcpuCountsOptions struct {
	SkuListSubscription        string
	SkuListAuth                lib.CldConfigTenantAuth
	Location                   string
	CostDataMonth              string
	CostDataBlobPrefix         string
	AzResSKUColl               *mongo.Collection
	AzResVcpuCountsColl        *mongo.Collection
	EnvOptCostingTenantsColl   *mongo.Collection
	EnvOptCostingSubsColl      *mongo.Collection
	EnvOptCostingResGrpsColl   *mongo.Collection
	EnvOptCostingResourcesColl *mongo.Collection
	EnvOptCostingMetersColl    *mongo.Collection
	AzResTenantsColl           *mongo.Collection
	AzResResourceListColl      *mongo.Collection
	AzResGrpsListColl          *mongo.Collection
}

type UpdateEntraItemsOptions struct {
	EntraAppRegColl              *mongo.Collection
	EntraAppRegCredsExpiringColl *mongo.Collection
}

type UpdateEntraPimItemsOptions struct {
	EntraRoleEligibilityScheduleInstancesColl *mongo.Collection
	EntraRoleAssignmentScheduleInstancesColl  *mongo.Collection
}
