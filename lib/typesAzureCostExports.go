package lib

import "time"

type RowData struct {
	DepartmentName   string  `csv:"-"`
	AccountName      string  `csv:"-"`
	AccountOwnerId   string  `csv:"-"`
	SubscriptionGuid string  `csv:"SubscriptionGuid" json:"SubscriptionGuid"`
	SubscriptionName string  `csv:"SubscriptionName" json:"SubscriptionName"`
	ResourceGroup    string  `csv:"ResourceGroup" json:"ResourceGroup"`
	ResourceLocation string  `csv:"-"`
	AvailabilityZone string  `csv:"-"`
	UsageDateTime    string  `csv:"UsageDateTime" json:"UsageDateTime"`
	ProductName      string  `csv:"ProductName" json:"ProductName"`
	MeterCategory    string  `csv:"-"`
	MeterSubcategory string  `csv:"-"`
	MeterId          string  `csv:"-"`
	MeterName        string  `csv:"MeterName" json:"MeterName"`
	MeterRegion      string  `csv:"-"`
	UnitOfMeasure    string  `csv:"UnitOfMeasure" json:"UnitOfMeasure"`
	UsageQuantity    float64 `csv:"UsageQuantity" json:"UsageQuantity"`
	ResourceRate     float64 `csv:"ResourceRate" json:"ResourceRate"`
	PreTaxCost       float64 `csv:"PreTaxCost" json:"PreTaxCost"`
	CostCenter       string  `csv:"-"`
	ConsumedService  string  `csv:"ConsumedService" json:"ConsumedService"`
	ResourceType     string  `csv:"ResourceType" json:"ResourceType"`
	InstanceId       string  `csv:"InstanceId" json:"InstanceId"`
	Tags             string  `csv:"Tags" json:"Tags"`
	OfferId          string  `csv:"-"`
	AdditionalInfo   string  `csv:"AdditionalInfo" json:"AdditionalInfo"`
	ServiceInfo1     string  `csv:"-"`
	ServiceInfo2     string  `csv:"-"`
	Currency         string  `csv:"Currency" json:"Currency"`
	Datafile         string
}

type CostExportData []RowData

type FieldMismatch struct {
	Expected, Found int
}

type UnsupportedType struct {
	Type string
}

type TransformedCostItem struct {
	SubscriptionName string `csv:"SubscriptionName" json:"SubscriptionName"`
	SubscriptionId   string `csv:"SubscriptionId" json:"SubscriptionId"`
	ResourceGroup    string `csv:"ResourceGroup" json:"ResourceGroup"`
	UsageDateTime    string `csv:"UsageDateTime" json:"UsageDateTime"`
	ProductName      string `csv:"ProductName" json:"ProductName"`
	MeterCategory    string `csv:"MeterCategory" json:"MeterCategory"`
	MeterSubcategory string `csv:"MeterSubcategory" json:"MeterSubcategory"`
	// MeterId          string  `csv:"-"`
	MeterName string `csv:"MeterName" json:"MeterName"`
	// MeterRegion      string  `csv:"-"`
	UnitOfMeasure string  `csv:"UnitOfMeasure" json:"UnitOfMeasure"`
	UsageQuantity float64 `csv:"UsageQuantity" json:"UsageQuantity"`
	ResourceRate  float64 `csv:"ResourceRate" json:"ResourceRate"`
	PreTaxCost    float64 `csv:"PreTaxCost" json:"PreTaxCost"`
	// CostCenter      string  `csv:"CostCenter" json:"CostCenter"`
	ConsumedService string            `csv:"ConsumedService" json:"ConsumedService"`
	ResourceType    string            `csv:"ResourceType" json:"ResourceType"`
	InstanceId      string            `csv:"InstanceId" json:"InstanceId"`
	Tags            map[string]string `csv:"Tags,omitempty" json:"Tags,omitempty"`
	// Tags string `csv:"Tags,omitempty" json:"Tags,omitempty"`
	// OfferId          string  `csv:"-"`
	AdditionalInfo interface{} `csv:"AdditionalInfo,omitempty" json:"AdditionalInfo,omitempty"`
	// AdditionalInfo map[string]string `csv:"AdditionalInfo,omitempty" json:"AdditionalInfo,omitempty"`
	// AdditionalInfo string `csv:"AdditionalInfo,omitempty" json:"AdditionalInfo,omitempty"`
	// ServiceInfo1     string  `csv:"-"`
	// ServiceInfo2     string  `csv:"-"`
	// Currency string `csv:"Currency" json:"Currency"`

	ResourceMeterIdentifier string `csv:"ResourceMeterIdentifier" json:"ResourceMeterIdentifier"`
	ResourceName            string `csv:"ResourceName" json:"ResourceName"`
	Tenant                  string `csv:"Tenant" json:"Tenant"`
	Datafile                string `csv:"Datafile" json:"Datafile"`
}

type TransformedTenantData struct {
	PreTaxCost float64
	ResGroups  []TransformedCostItem
	// MeterData  map[string]TransformedCostItem
}

type AggregatedCostTenant struct {
	CostPerDay     CostPerDay                            `json:"costPerDay,omitempty"  bson:"costPerDay,omitempty"`
	MonthTotalCost float64                               `json:"monthTotalCost,omitempty"  bson:"monthTotalCost,omitempty"`
	Subscriptions  map[string]AggregatedCostSubscription `json:"subscriptions,omitempty"  bson:"subscriptions,omitempty"`
	CostGroups     map[string]string                     `json:"costGroups,omitempty"  bson:"costGroups,omitempty"`
}

type MongoDbCostTenant struct {
	TenantName        string                             `json:"tenantName,omitempty"  bson:"tenantName,omitempty"`
	TenantId          string                             `json:"tenantId,omitempty"  bson:"tenantId,omitempty"`
	CostData          map[string]MongoDbCostData         `json:"costData,omitempty"  bson:"costData,omitempty"`
	LifetimeTotalCost float64                            `json:"lifetimeTotalCost,omitempty"  bson:"lifetimeTotalCost,omitempty"`
	Subscriptions     map[string]MongoDbCostSubscription `json:"subscriptions,omitempty"  bson:"subscriptions,omitempty"`
	CostGroups        map[string]string                  `json:"costGroups,omitempty"  bson:"costGroups,omitempty"`
	RelatedCostMeters []string                           `json:"relatedCostMeters,omitempty"  bson:"relatedCostMeters,omitempty"`
	LastDBSync        time.Time                          `json:"lastDatabaseSync,omitempty"  bson:"lastDatabaseSync,omitempty"`
}

type MongoDbCostSubscription struct {
	TenantId          string                              `json:"tenantId,omitempty"  bson:"tenantId,omitempty"`
	TenantName        string                              `json:"tenantName,omitempty"  bson:"tenantName,omitempty"`
	SubscriptionName  string                              `json:"subscriptionName,omitempty"  bson:"subscriptionName,omitempty"`
	SubscriptionId    string                              `json:"subscriptionId,omitempty"  bson:"subscriptionId,omitempty"`
	CostData          map[string]MongoDbCostData          `json:"costData,omitempty"  bson:"costData,omitempty"`
	LifetimeTotalCost float64                             `json:"lifetimeTotalCost,omitempty"  bson:"lifetimeTotalCost,omitempty"`
	ResourceGroups    map[string]MongoDbCostResourceGroup `json:"resourceGroups,omitempty"  bson:"resourceGroups,omitempty"`
	CostGroups        map[string]string                   `json:"costGroups,omitempty"  bson:"costGroups,omitempty"`
	RelatedCostMeters []string                            `json:"relatedCostMeters,omitempty"  bson:"relatedCostMeters,omitempty"`
	LastDBSync        time.Time                           `json:"lastDatabaseSync,omitempty"  bson:"lastDatabaseSync,omitempty"`
}

type MongoDbCostResourceGroup struct {
	MongoId           string                     `json:"mongoId,omitempty"  bson:"_id,omitempty"`
	TenantId          string                     `json:"tenantId,omitempty"  bson:"tenantId,omitempty"`
	TenantName        string                     `json:"tenantName,omitempty"  bson:"tenantName,omitempty"`
	SubscriptionId    string                     `json:"subscriptionId,omitempty"  bson:"subscriptionId,omitempty"`
	SubscriptionName  string                     `json:"subscriptionName,omitempty"  bson:"subscriptionName,omitempty"`
	Name              string                     `json:"name,omitempty"  bson:"name,omitempty"`
	CostData          map[string]MongoDbCostData `json:"costData,omitempty"  bson:"costData,omitempty"`
	LifetimeTotalCost float64                    `json:"lifetimeTotalCost,omitempty"  bson:"lifetimeTotalCost,omitempty"`
	CostGroups        map[string]string          `json:"costGroups,omitempty"  bson:"costGroups,omitempty"`
	RelatedCostMeters []string                   `json:"relatedCostMeters,omitempty"  bson:"relatedCostMeters,omitempty"`
	LastDBSync        time.Time                  `json:"lastDatabaseSync,omitempty"  bson:"lastDatabaseSync,omitempty"`
}

type MongoDbCostResource struct {
	MongoId           string                     `json:"mongoId,omitempty"  bson:"_id,omitempty"`
	TenantId          string                     `json:"tenantId,omitempty"  bson:"tenantId,omitempty"`
	TenantName        string                     `json:"tenantName,omitempty"  bson:"tenantName,omitempty"`
	SubscriptionId    string                     `json:"subscriptionId,omitempty"  bson:"subscriptionId,omitempty"`
	SubscriptionName  string                     `json:"subscriptionName,omitempty"  bson:"subscriptionName,omitempty"`
	ResourceGroupName string                     `json:"resourceGroupName,omitempty"  bson:"resourceGroupName,omitempty"`
	ResGrpMongoId     string                     `json:"resGrpmongoId,omitempty"  bson:"resGrpMongoId,omitempty"`
	Name              string                     `json:"name,omitempty"  bson:"name,omitempty"`
	CostData          map[string]MongoDbCostData `json:"costData,omitempty"  bson:"costData,omitempty"`
	LifetimeTotalCost float64                    `json:"lifetimeTotalCost,omitempty"  bson:"lifetimeTotalCost,omitempty"`
	CostGroups        map[string]string          `json:"costGroups,omitempty"  bson:"costGroups,omitempty"`
	ResourceType      string                     `json:"resourceType,omitempty"  bson:"resourceType,omitempty"`
	RelatedCostMeters []string                   `json:"relatedCostMeters,omitempty"  bson:"relatedCostMeters,omitempty"`
	LastDBSync        time.Time                  `json:"lastDatabaseSync,omitempty"  bson:"lastDatabaseSync,omitempty"`
}

type MongoDbCostMeter struct {
	ResourceMeterIdentifier string                     `json:"mongoId,omitempty"  bson:"_id,omitempty"`
	TenantId                string                     `json:"tenantId,omitempty"  bson:"tenantId,omitempty"`
	TenantName              string                     `json:"tenantName,omitempty"  bson:"tenantName,omitempty"`
	SubscriptionId          string                     `json:"subscriptionId,omitempty"  bson:"subscriptionId,omitempty"`
	SubscriptionName        string                     `json:"subscriptionName,omitempty"  bson:"subscriptionName,omitempty"`
	ResourceGroupName       string                     `json:"resourceGroupName,omitempty"  bson:"resourceGroupName,omitempty"`
	ResGrpMongoId           string                     `json:"resGrpmongoId,omitempty"  bson:"resGrpMongoId,omitempty"`
	ResourceId              string                     `json:"resourceId,omitempty"  bson:"resourceId,omitempty"`
	ResourceName            string                     `json:"resourceName,omitempty"  bson:"resourceName,omitempty"`
	ResourceMongoId         string                     `json:"resourceMongoId,omitempty"  bson:"resourceMongoId,omitempty"`
	CostData                map[string]MongoDbCostData `json:"costData,omitempty"  bson:"costData,omitempty"`
	LifetimeTotalCost       float64                    `json:"lifetimeTotalCost,omitempty"  bson:"lifetimeTotalCost,omitempty"`
	ProductName             string                     `csv:"ProductName" json:"productName,omitempty"  bson:"productName,omitempty"`
	MeterCategory           string                     `csv:"MeterCategory" json:"meterCategory,omitempty"  bson:"meterCategory,omitempty"`
	MeterSubcategory        string                     `csv:"MeterSubcategory" json:"meterSubcategory,omitempty"  bson:"meterSubcategory,omitempty"`
	MeterName               string                     `csv:"MeterName" json:"meterName,omitempty"  bson:"meterName,omitempty"`
	UnitOfMeasure           string                     `csv:"UnitOfMeasure" json:"unitOfMeasure,omitempty"  bson:"unitOfMeasure,omitempty"`
	ResourceRate            float64                    `csv:"ResourceRate" json:"resourceRate,omitempty"  bson:"resourceRate,omitempty"`
	ConsumedService         string                     `csv:"ConsumedService" json:"consumedService,omitempty"  bson:"consumedService,omitempty"`
	ResourceType            string                     `csv:"ResourceType" json:"resourceType,omitempty"  bson:"resourceType,omitempty"`
	InstanceId              string                     `csv:"InstanceId" json:"instanceId,omitempty"  bson:"instanceId,omitempty"`
	AdditionalInfo          interface{}                `csv:"AdditionalInfo,omitempty" json:"additionalInfo,omitempty"  bson:"additionalInfo,omitempty"`
	LastDBSync              time.Time                  `json:"lastDatabaseSync,omitempty"  bson:"lastDatabaseSync,omitempty"`
}

type CostPerDay map[string]float64

// MeterData           map[string]AggregatedCostItem

type AggregatedCostSubscription struct {
	CostPerDay     CostPerDay                             `json:"costPerDay,omitempty"  bson:"costPerDay,omitempty"`
	MonthTotalCost float64                                `json:"monthTotalCost,omitempty"  bson:"monthTotalCost,omitempty"`
	ResourceGroups map[string]AggregatedCostResourceGroup `json:"resourceGroups,omitempty"  bson:"resourceGroups,omitempty"`
	CostGroups     map[string]string                      `json:"costGroups,omitempty"  bson:"costGroups,omitempty"`
}

type AggregatedCostResourceGroup struct {
	CostPerDay     CostPerDay                        `json:"costPerDay,omitempty"  bson:"costPerDay,omitempty"`
	MonthTotalCost float64                           `json:"monthTotalCost,omitempty"  bson:"monthTotalCost,omitempty"`
	Resources      map[string]AggregatedCostResource `json:"resources,omitempty"  bson:"resources,omitempty"`
	CostGroups     map[string]string                 `json:"costGroups,omitempty"  bson:"costGroups,omitempty"`
}

type AggregatedCostResource struct {
	CostPerDay     CostPerDay                    `json:"costPerDay,omitempty"  bson:"costPerDay,omitempty"`
	MonthTotalCost float64                       `json:"monthTotalCost,omitempty"  bson:"monthTotalCost,omitempty"`
	MeterData      map[string]AggregatedCostItem `json:"meterData,omitempty"  bson:"meterData,omitempty"`
	CostGroups     map[string]string             `json:"costGroups,omitempty"  bson:"costGroups,omitempty"`
}

type AggregatedCostCostGroups struct {
}

type AggregatedCostItem struct {
	CostPerDay          map[string]float64 `json:"costPerDay,omitempty"  bson:"costPerDay,omitempty"`
	UsageQuantityPerDay map[string]float64 `json:"usageQuantityPerDay,omitempty"  bson:"usageQuantityPerDay,omitempty"`
	MonthTotalCost      float64            `json:"monthTotalCost,omitempty"  bson:"monthTotalCost,omitempty"`
	TenantName          string             `json:"tenantName,omitempty"  bson:"tenantName,omitempty"`
	SubscriptionName    string             `json:"subscriptionName,omitempty"  bson:"subscriptionName,omitempty"`
	ProductName         string             `csv:"ProductName" json:"productName,omitempty"  bson:"productName,omitempty"`
	MeterCategory       string             `csv:"MeterCategory" json:"meterCategory,omitempty"  bson:"meterCategory,omitempty"`
	MeterSubcategory    string             `csv:"MeterSubcategory" json:"meterSubcategory,omitempty"  bson:"meterSubcategory,omitempty"`
	MeterName           string             `csv:"MeterName" json:"meterName,omitempty"  bson:"meterName,omitempty"`
	UnitOfMeasure       string             `csv:"UnitOfMeasure" json:"unitOfMeasure,omitempty"  bson:"unitOfMeasure,omitempty"`
	ResourceRate        float64            `csv:"ResourceRate" json:"resourceRate,omitempty"  bson:"resourceRate,omitempty"`
	ConsumedService     string             `csv:"ConsumedService" json:"consumedService,omitempty"  bson:"consumedService,omitempty"`
	ResourceType        string             `csv:"ResourceType" json:"resourceType,omitempty"  bson:"resourceType,omitempty"`
	InstanceId          string             `csv:"InstanceId" json:"instanceId,omitempty"  bson:"instanceId,omitempty"`
	// UsageQuantity    float64           `csv:"UsageQuantity" json:"usageQuantity" json:"usageQuantity"`
	Tags map[string]string `csv:"Tags" json:"tags"  bson:"tags"`
	// Tags string `csv:"Tags,omitempty" json:"Tags,omitempty"`
	// OfferId          string  `csv:"-"`
	AdditionalInfo          interface{} `csv:"AdditionalInfo,omitempty" json:"additionalInfo,omitempty"  bson:"additionalInfo,omitempty"`
	ResourceMeterIdentifier string      `json:"resourceMeterIdentifier,omitempty"  bson:"resourceMeterIdentifier,omitempty"`
}

//
//

type CostItemFlat struct {
	Cost                    float64   `json:"cost,omitempty"  bson:"cost,omitempty"`
	Date                    time.Time `json:"date,omitempty"  bson:"date,omitempty"`
	TenantName              string    `json:"tenantName,omitempty"  bson:"tenantName,omitempty"`
	SubscriptionName        string    `json:"subscriptionName,omitempty"  bson:"subscriptionName,omitempty"`
	ProductName             string    `csv:"ProductName" json:"productName,omitempty"  bson:"productName,omitempty"`
	InstanceId              string    `csv:"InstanceId" json:"instanceId,omitempty"  bson:"instanceId,omitempty"`
	UsageQuantity           float64   `csv:"UsageQuantity" json:"usageQuantity" bson:"usageQuantity"`
	Tags                    string    `csv:"Tags" json:"tags"  bson:"tags"`
	AdditionalInfo          string    `csv:"AdditionalInfo,omitempty" json:"additionalInfo,omitempty"  bson:"additionalInfo,omitempty"`
	ResourceMeterIdentifier string    `json:"resourceMeterIdentifier,omitempty"  bson:"resourceMeterIdentifier,omitempty"`
}

//
//

type AzureCostMeterFlat struct {
	ProductName             string  `csv:"ProductName" json:"productName,omitempty"  bson:"productName,omitempty"`
	MeterCategory           string  `csv:"MeterCategory" json:"meterCategory,omitempty"  bson:"meterCategory,omitempty"`
	MeterSubcategory        string  `csv:"MeterSubcategory" json:"meterSubcategory,omitempty"  bson:"meterSubcategory,omitempty"`
	MeterName               string  `csv:"MeterName" json:"meterName,omitempty"  bson:"meterName,omitempty"`
	UnitOfMeasure           string  `csv:"UnitOfMeasure" json:"unitOfMeasure,omitempty"  bson:"unitOfMeasure,omitempty"`
	ResourceRate            float64 `csv:"ResourceRate" json:"resourceRate,omitempty"  bson:"resourceRate,omitempty"`
	ConsumedService         string  `csv:"ConsumedService" json:"consumedService,omitempty"  bson:"consumedService,omitempty"`
	ResourceType            string  `csv:"ResourceType" json:"resourceType,omitempty"  bson:"resourceType,omitempty"`
	ResourceMeterIdentifier string  `json:"resourceMeterIdentifier,omitempty"  bson:"resourceMeterIdentifier,omitempty"`
}

type AzureCostMeter struct {
	ProductName             string      `csv:"ProductName" json:"productName,omitempty"  bson:"productName,omitempty"`
	MeterCategory           string      `csv:"MeterCategory" json:"meterCategory,omitempty"  bson:"meterCategory,omitempty"`
	MeterSubcategory        string      `csv:"MeterSubcategory" json:"meterSubcategory,omitempty"  bson:"meterSubcategory,omitempty"`
	MeterName               string      `csv:"MeterName" json:"meterName,omitempty"  bson:"meterName,omitempty"`
	UnitOfMeasure           string      `csv:"UnitOfMeasure" json:"unitOfMeasure,omitempty"  bson:"unitOfMeasure,omitempty"`
	ResourceRate            float64     `csv:"ResourceRate" json:"resourceRate,omitempty"  bson:"resourceRate,omitempty"`
	ConsumedService         string      `csv:"ConsumedService" json:"consumedService,omitempty"  bson:"consumedService,omitempty"`
	ResourceType            string      `csv:"ResourceType" json:"resourceType,omitempty"  bson:"resourceType,omitempty"`
	InstanceId              string      `csv:"InstanceId" json:"instanceId,omitempty"  bson:"instanceId,omitempty"`
	AdditionalInfo          interface{} `csv:"AdditionalInfo,omitempty" json:"additionalInfo,omitempty"  bson:"additionalInfo,omitempty"`
	ResourceMeterIdentifier string      `json:"resourceMeterIdentifier,omitempty"  bson:"resourceMeterIdentifier,omitempty"`
}

type MongoDbCostItem struct {
	TenantName        string                     `json:"tenantName,omitempty"  bson:"tenantName,omitempty"`
	SubscriptionName  string                     `json:"subscriptionName,omitempty"  bson:"subscriptionName,omitempty"`
	LifetimeTotalCost float64                    `json:"lifetimeTotalCost,omitempty"  bson:"lifetimeTotalCost,omitempty"`
	ProductName       string                     `csv:"ProductName" json:"productName,omitempty"  bson:"productName,omitempty"`
	MeterCategory     string                     `csv:"MeterCategory" json:"meterCategory,omitempty"  bson:"meterCategory,omitempty"`
	MeterSubcategory  string                     `csv:"MeterSubcategory" json:"meterSubcategory,omitempty"  bson:"meterSubcategory,omitempty"`
	MeterName         string                     `csv:"MeterName" json:"meterName,omitempty"  bson:"meterName,omitempty"`
	ConsumedService   string                     `csv:"ConsumedService" json:"consumedService,omitempty"  bson:"consumedService,omitempty"`
	ResourceType      string                     `csv:"ResourceType" json:"resourceType,omitempty"  bson:"resourceType,omitempty"`
	InstanceId        string                     `csv:"InstanceId" json:"instanceId,omitempty"  bson:"instanceId,omitempty"`
	AdditionalInfo    interface{}                `csv:"AdditionalInfo,omitempty" json:"additionalInfo,omitempty"  bson:"additionalInfo,omitempty"`
	CostData          map[string]MongoDbCostData `json:"costData,omitempty"  bson:"costData,omitempty"`
	// UsageQuantity           float64                    `csv:"UsageQuantity" json:"usageQuantity" json:"usageQuantity"`
	ResourceMeterIdentifier string `json:"resourceMeterIdentifier,omitempty"  bson:"_id,omitempty"`
}

type MongoDbCostData struct {
	CostPerDay          map[string]float64 `json:"costPerDay,omitempty"  bson:"costPerDay,omitempty"`
	MonthTotalCost      float64            `json:"monthTotalCost,omitempty"  bson:"monthTotalCost,omitempty"`
	UnitOfMeasure       string             `csv:"UnitOfMeasure" json:"unitOfMeasure,omitempty"  bson:"unitOfMeasure,omitempty"`
	ResourceRate        float64            `csv:"ResourceRate" json:"resourceRate,omitempty"  bson:"resourceRate,omitempty"`
	UsageQuantityPerDay map[string]float64 `json:"usageQuantityPerDay,omitempty"  bson:"usageQuantityPerDay,omitempty"`
}

type AggregatedCostData map[string]AggregatedCostTenant

// type TransformedCostItemsByTenantMap map[string]TransformedTenantData

type TransformedCostItemsByTenant map[string]TransformedTenantData

type DownloadAllConfiguredTenantCostExportsForMonthOptions struct {
	// Prefix for blob files
	// Example: "monthly-cost-exports/202404"
	// Example: "daily-month-to-date-exports/202404"
	BlobPrefix string

	// Path for file download
	OutfilePath string

	// filename prefix for file download
	// Tenant Name will be appended to filename
	// Example: "outputs/cost-exports" would become "outputs/cost-exports__TENANTNAME"
	OutfileNamePrefix string

	// YYYYMM format
	// Example: 202401
	CostExportMonth string

	// Used when downloading for multiple months using DownloadAllConfiguredTenantCostExportsForMonth
	SuppressSteps bool
}

type CostQueryResponse struct {
	ID         string `json:"id"`
	Properties struct {
		NextLink string `json:"nextLink"`
		Columns  []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"columns"`
		Rows [][]interface{} `json:"rows"`
	} `json:"properties"`
}

type Post struct {
	DataSet struct {
		Aggregation struct {
			TotalCost struct {
				Function string
				Name     string
			}
		}
		Granularity string
		Grouping    []struct {
			Name string
			Type string
		}
		Sorting []struct {
			Direction string
			Name      string
		}
	}
	TimePeriod struct {
		From string
		To   string
	}
	Timeframe string
	Type      string
}

type BlobItem struct {
	Name               string `json:"Name"`
	ContainerName      string `json:"containerName"`
	TenantName         string `json:"tenantName"`
	StorageAccountName string `json:"storageAccountName"`
	BlobTags           any    `json:"BlobTags"`
	Deleted            any    `json:"Deleted"`
	IsCurrentVersion   any    `json:"IsCurrentVersion"`
	Metadata           any    `json:"Metadata"`
	OrMetadata         any    `json:"OrMetadata"`
	Properties         struct {
		AccessTier             string    `json:"AccessTier"`
		AccessTierChangeTime   any       `json:"AccessTierChangeTime"`
		AccessTierInferred     bool      `json:"AccessTierInferred"`
		BlobType               string    `json:"BlobType"`
		ContentMd5             string    `json:"ContentMD5"`
		ContentType            string    `json:"ContentType"`
		CreationTime           time.Time `json:"CreationTime"`
		DeletedTime            any       `json:"DeletedTime"`
		LastAccessedOn         any       `json:"LastAccessedOn"`
		LastModified           time.Time `json:"LastModified"`
		RemainingRetentionDays any       `json:"RemainingRetentionDays"`
		ServerEncrypted        bool      `json:"ServerEncrypted"`
	} `json:"Properties"`
	Snapshot  any `json:"Snapshot"`
	VersionID any `json:"VersionID"`
}

type BlobList []BlobItem

type StorageAccountRequestOptions struct {
	StorageAccountName   string
	ContainerName        string
	ConfiguredTenantName string
	GetWriteToken        bool
	BlobName             string
	DownloadFileName     string
	OverwriteExisting    bool   // Only used with azure.DownloadAllBlobsInContainer
	ShowDownloadedCount  bool   // Only used with azure.DownloadAllBlobsInContainer
	DownloadPath         string // Only used with azure.DownloadAllBlobsInContainer
}

type BlobListFilterOptions struct {
	FilterPrefix   string
	FilterSuffix   string
	FilterContains string
}
