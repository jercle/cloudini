package azure

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
	Tags             string  `csv:"-"`
	OfferId          string  `csv:"-"`
	AdditionalInfo   string  `csv:"-"`
	ServiceInfo1     string  `csv:"-"`
	ServiceInfo2     string  `csv:"-"`
	Currency         string  `csv:"Currency" json:"Currency"`
	Datafile         string
}

type CostExportData []RowData

type FieldMismatch struct {
	expected, found int
}

type UnsupportedType struct {
	Type string
}

type TransformedCostItem struct {
	SubscriptionName string `csv:"SubscriptionName" json:"SubscriptionName"`
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
	CostPerDay     CostPerDay
	MonthTotalCost float64
	Subscriptions  map[string]AggregatedCostSubscription
	CostGroups     map[string]CostPerDay
}

type CostPerDay map[string]float64

// MeterData           map[string]AggregatedCostItem

type AggregatedCostSubscription struct {
	CostPerDay     CostPerDay
	MonthTotalCost float64
	ResourceGroups map[string]AggregatedCostResourceGroup
}

type AggregatedCostResourceGroup struct {
	CostPerDay     CostPerDay
	MonthTotalCost float64
	Resources      map[string]AggregatedCostResource
}

type AggregatedCostResource struct {
	CostPerDay     CostPerDay
	MonthTotalCost float64
	MeterData      map[string]AggregatedCostItem
}

type AggregatedCostCostGroups struct {
}

type AggregatedCostItem struct {
	CostPerDay       map[string]float64
	MonthTotalCost   float64
	ProductName      string            `csv:"ProductName" json:"ProductName"`
	MeterCategory    string            `csv:"MeterCategory" json:"MeterCategory"`
	MeterSubcategory string            `csv:"MeterSubcategory" json:"MeterSubcategory"`
	MeterName        string            `csv:"MeterName" json:"MeterName"`
	UnitOfMeasure    string            `csv:"UnitOfMeasure" json:"UnitOfMeasure"`
	ResourceRate     float64           `csv:"ResourceRate" json:"ResourceRate"`
	ConsumedService  string            `csv:"ConsumedService" json:"ConsumedService"`
	ResourceType     string            `csv:"ResourceType" json:"ResourceType"`
	InstanceId       string            `csv:"InstanceId" json:"InstanceId"`
	Tags             map[string]string `csv:"Tags,omitempty" json:"Tags,omitempty"`
	// Tags string `csv:"Tags,omitempty" json:"Tags,omitempty"`
	// OfferId          string  `csv:"-"`
	AdditionalInfo interface{} `csv:"AdditionalInfo,omitempty" json:"AdditionalInfo,omitempty"`
}

type AggregatedCostData map[string]AggregatedCostTenant

type TransformedCostItemsByTenantMap map[string]TransformedTenantData

// type TransformedCostItemsByTenant map[string]TransformedTenantData
type TransformedCostItemsByTenant struct {
	Blue      TransformedTenantData
	BlueDTQ   TransformedTenantData
	Red       TransformedTenantData
	RedDTQ    TransformedTenantData
	Yellow    TransformedTenantData
	PUD       TransformedTenantData
	PUDDTQ    TransformedTenantData
	Purple    TransformedTenantData
	PurpleDTQ TransformedTenantData
}

type DownloadAllConfiguredTenantLastMonthCostExportsOptions struct {
	// Prefix for blob files
	// Example: "monthly-cost-exports/202404"
	// Example: "daily-month-to-date-exports/202404"
	BlobPrefix string

	// Path and filename prefix for file download
	// Tenant Name will be appended to filename
	// Example: "outputs/cost-exports" would become "outputs/cost-exports__TENANTNAME"
	OutfilePath string
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
}

type BlobListFilterOptions struct {
	FilterPrefix   string
	FilterSuffix   string
	FilterContains string
}
