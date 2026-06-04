package azure

import (
	"time"

	"github.com/jercle/cloudini/lib"
)

type SubscriptionResourceList struct {
	ResourceCount int
	Resources     []lib.AzureResourceDetails
}

//
//

type SubscriptionResGrpList []ResourceGroup

//
//

type TenantResourceList struct {
	ResourceCount int
	Subscriptions map[string]SubscriptionResourceList
}

type TenantResGrpList map[string]SubscriptionResGrpList

//
//

type ResourceGraphResponse struct {
	Count           float64                    `json:"count,omitempty,omitzero" bson:"count,omitempty,omitzero"`
	Data            []lib.AzureResourceDetails `json:"data,omitempty,omitzero" bson:"data,omitempty,omitzero"`
	Facets          []interface{}              `json:"facets,omitempty,omitzero" bson:"facets,omitempty,omitzero"`
	ResultTruncated string                     `json:"resultTruncated,omitempty,omitzero" bson:"resultTruncated,omitempty,omitzero"`
	SkipToken       string                     `json:"$skipToken,omitempty,omitzero" bson:"$skipToken,omitempty,omitzero"`
	TotalRecords    float64                    `json:"totalRecords,omitempty,omitzero" bson:"totalRecords,omitempty,omitzero"`
}

type ResourceGraphResponseDataInterface struct {
	Count           float64       `json:"count,omitempty,omitzero" bson:"count,omitempty,omitzero"`
	Data            []interface{} `json:"data,omitempty,omitzero" bson:"data,omitempty,omitzero"`
	Facets          []interface{} `json:"facets,omitempty,omitzero" bson:"facets,omitempty,omitzero"`
	ResultTruncated string        `json:"resultTruncated,omitempty,omitzero" bson:"resultTruncated,omitempty,omitzero"`
	SkipToken       string        `json:"$skipToken,omitempty,omitzero" bson:"$skipToken,omitempty,omitzero"`
	TotalRecords    float64       `json:"totalRecords,omitempty,omitzero" bson:"totalRecords,omitempty,omitzero"`
}

//
//

type ListAllResourceGroupsResponse struct {
	Count           float64         `json:"count,omitempty,omitzero" bson:"count,omitempty,omitzero"`
	Data            []ResourceGroup `json:"data,omitempty,omitzero" bson:"data,omitempty,omitzero"`
	Facets          []any           `json:"facets,omitempty,omitzero" bson:"facets,omitempty,omitzero"`
	ResultTruncated string          `json:"resultTruncated,omitempty,omitzero" bson:"resultTruncated,omitempty,omitzero"`
	TotalRecords    float64         `json:"totalRecords,omitempty,omitzero" bson:"totalRecords,omitempty,omitzero"`
	SkipToken       string          `json:"$skipToken,omitempty,omitzero" bson:"$skipToken,omitempty,omitzero"`
}

//
//

type ResourceGroup struct {
	ExtendedLocation any    `json:"extendedLocation,omitempty,omitzero" bson:"extendedLocation,omitempty,omitzero"`
	ID               string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Identity         any    `json:"identity,omitempty,omitzero" bson:"identity,omitempty,omitzero"`
	Kind             string `json:"kind,omitempty,omitzero" bson:"kind,omitempty,omitzero"`
	Location         string `json:"location,omitempty,omitzero" bson:"location,omitempty,omitzero"`
	ManagedBy        string `json:"managedBy,omitempty,omitzero" bson:"managedBy,omitempty,omitzero"`
	Name             string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Plan             any    `json:"plan,omitempty,omitzero" bson:"plan,omitempty,omitzero"`
	Properties       struct {
		ProvisioningState string `json:"provisioningState,omitempty,omitzero" bson:"provisioningState,omitempty,omitzero"`
	} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
	ResourceGroup  string            `json:"resourceGroup,omitempty,omitzero" bson:"resourceGroup,omitempty,omitzero"`
	Sku            any               `json:"sku,omitempty,omitzero" bson:"sku,omitempty,omitzero"`
	SubscriptionID string            `json:"subscriptionId,omitempty,omitzero" bson:"subscriptionId,omitempty,omitzero"`
	Tags           map[string]string `json:"tags,omitempty,omitzero" bson:"tags,omitempty,omitzero"`
	TenantID       string            `json:"tenantId,omitempty,omitzero" bson:"tenantId,omitempty,omitzero"`
	Type           string            `json:"type,omitempty,omitzero" bson:"type,omitempty,omitzero"`
	Zones          any               `json:"zones,omitempty,omitzero" bson:"zones,omitempty,omitzero"`
	TenantName     string            `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	LastAzureSync  time.Time         `json:"lastAzureSync,omitempty,omitzero" bson:"lastAzureSync,omitempty,omitzero"`
	LastDBSync     time.Time         `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
}

// type GetResourceChangesQuery struct {
// 	SkipToken            string                          `json:"$skipToken,omitempty,omitzero" bson:"$skipToken,omitempty,omitzero"`
// 	Interval             GetResourceChangesQueryInterval `json:"interval,omitempty,omitzero" bson:"interval,omitempty,omitzero"`
// 	SubscriptionId       string                          `json:"subscriptionId,omitempty,omitzero" bson:"subscriptionId,omitempty,omitzero"`
// 	FetchPropertyChanges bool                            `json:"fetchPropertyChanges,omitempty,omitzero" bson:"fetchPropertyChanges,omitempty,omitzero"`
// 	FetchSnapshots       bool                            `json:"fetchSnapshots,omitempty,omitzero" bson:"fetchSnapshots,omitempty,omitzero"`
// }
// type GetResourceChangesQueryInterval struct {
// 	Start time.Time `json:"start,omitempty,omitzero" bson:"start,omitempty,omitzero"`
// 	End   time.Time `json:"end,omitempty,omitzero" bson:"end,omitempty,omitzero"`
// }

// type GetAllResourceChangesResponse struct {
// 	SkipToken string           `json:"$skipToken,omitempty,omitzero" bson:"$skipToken,omitempty,omitzero"`
// 	Changes   []ResourceChange `json:"changes,omitempty,omitzero" bson:"changes,omitempty,omitzero"`
// }

// type ResourceChange struct {
// 	AfterSnapshot struct {
// 		SnapshotID string    `json:"snapshotId,omitempty,omitzero" bson:"snapshotId,omitempty,omitzero"`
// 		Timestamp  time.Time `json:"timestamp,omitempty,omitzero" bson:"timestamp,omitempty,omitzero"`
// 	} `json:"afterSnapshot,omitempty,omitzero" bson:"afterSnapshot,omitempty,omitzero"`
// 	BeforeSnapshot struct {
// 		SnapshotID string    `json:"snapshotId,omitempty,omitzero" bson:"snapshotId,omitempty,omitzero"`
// 		Timestamp  time.Time `json:"timestamp,omitempty,omitzero" bson:"timestamp,omitempty,omitzero"`
// 	} `json:"beforeSnapshot,omitempty,omitzero" bson:"beforeSnapshot,omitempty,omitzero"`
// 	ChangeID         string                     `json:"changeId,omitempty,omitzero" bson:"changeId,omitempty,omitzero"`
// 	ChangeType       string                     `json:"changeType,omitempty,omitzero" bson:"changeType,omitempty,omitzero"`
// 	PropertyChanges  []ResourceChangePropChange `json:"propertyChanges,omitempty" bson:"propertyChanges,omitempty"`
// 	ResourceID       string                     `json:"resourceId,omitempty,omitzero" bson:"resourceId,omitempty,omitzero"`
// 	SubscriptionId   string                     `json:"subscriptionId,omitempty,omitzero" bson:"subscriptionId,omitempty,omitzero"`
// 	SubscriptionName string                     `json:"subscriptionName,omitempty,omitzero" bson:"subscriptionName,omitempty,omitzero"`
// 	TenantId         string                     `json:"tenantId,omitempty,omitzero" bson:"tenantId,omitempty,omitzero"`
// 	TenantName       string                     `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
// }

// type ResourceChangePropChange struct {
// 	AfterValue         string `json:"afterValue,omitempty,omitzero" bson:"afterValue,omitempty,omitzero"`
// 	BeforeValue        string `json:"beforeValue,omitempty,omitzero" bson:"beforeValue,omitempty,omitzero"`
// 	ChangeCategory     string `json:"changeCategory,omitempty,omitzero" bson:"changeCategory,omitempty,omitzero"`
// 	PropertyChangeType string `json:"propertyChangeType,omitempty,omitzero" bson:"propertyChangeType,omitempty,omitzero"`
// 	PropertyName       string `json:"propertyName,omitempty,omitzero" bson:"propertyName,omitempty,omitzero"`
// }

type ResourceChangeRaw struct {
	ChangeOperation    string    `json:"changeOperation,omitempty,omitzero" bson:"changeOperation,omitempty,omitzero"`
	ChangeType         string    `json:"changeType,omitempty,omitzero" bson:"changeType,omitempty,omitzero"`
	ChangedBy          string    `json:"changedBy,omitempty,omitzero" bson:"changedBy,omitempty,omitzero"`
	ChangedByID        string    `json:"changedById,omitempty,omitzero" bson:"changedById,omitempty,omitzero"`
	ChangedByType      string    `json:"changedByType,omitempty,omitzero" bson:"changedByType,omitempty,omitzero"`
	Changes            string    `json:"changes,omitempty,omitzero" bson:"changes,omitempty,omitzero"`
	ChangesCount       uint      `json:"changesCount,omitempty,omitzero" bson:"changesCount,omitempty,omitzero"`
	CorrelationID      string    `json:"correlationId,omitempty,omitzero" bson:"correlationId,omitempty,omitzero"`
	ID                 string    `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
	Name               string    `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	ResourceGroup      string    `json:"resourceGroup,omitempty,omitzero" bson:"resourceGroup,omitempty,omitzero"`
	SubscriptionID     string    `json:"subscriptionId,omitempty,omitzero" bson:"subscriptionId,omitempty,omitzero"`
	SubscriptionName   string    `json:"subscriptionName,omitempty,omitzero" bson:"subscriptionName,omitempty,omitzero"`
	TargetResourceID   string    `json:"targetResourceId,omitempty,omitzero" bson:"targetResourceId,omitempty,omitzero"`
	TargetResourceType string    `json:"targetResourceType,omitempty,omitzero" bson:"targetResourceType,omitempty,omitzero"`
	TenantID           string    `json:"tenantId,omitempty,omitzero" bson:"tenantId,omitempty,omitzero"`
	TenantName         string    `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	Timestamp          time.Time `json:"timestamp,omitempty,omitzero" bson:"timestamp,omitempty,omitzero"`
}
type ResourceChange struct {
	ChangeOperation    string                              `json:"changeOperation,omitempty,omitzero" bson:"changeOperation,omitempty,omitzero"`
	ChangeType         string                              `json:"changeType,omitempty,omitzero" bson:"changeType,omitempty,omitzero"`
	ChangedBy          string                              `json:"changedBy,omitempty,omitzero" bson:"changedBy,omitempty,omitzero"`
	ChangedByID        string                              `json:"changedById,omitempty,omitzero" bson:"changedById,omitempty,omitzero"`
	ChangedByType      string                              `json:"changedByType,omitempty,omitzero" bson:"changedByType,omitempty,omitzero"`
	Changes            map[string]ResourceChangePropChange `json:"changes,omitempty,omitzero" bson:"changes,omitempty,omitzero"`
	ChangesCount       uint                                `json:"changesCount,omitempty,omitzero" bson:"changesCount,omitempty,omitzero"`
	CorrelationID      string                              `json:"correlationId,omitempty,omitzero" bson:"correlationId,omitempty,omitzero"`
	ID                 string                              `json:"id,omitempty,omitzero" bson:"_id,omitempty,omitzero"`
	Name               string                              `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	ResourceGroup      string                              `json:"resourceGroup,omitempty,omitzero" bson:"resourceGroup,omitempty,omitzero"`
	SubscriptionID     string                              `json:"subscriptionId,omitempty,omitzero" bson:"subscriptionId,omitempty,omitzero"`
	SubscriptionName   string                              `json:"subscriptionName,omitempty,omitzero" bson:"subscriptionName,omitempty,omitzero"`
	TargetResourceID   string                              `json:"targetResourceId,omitempty,omitzero" bson:"targetResourceId,omitempty,omitzero"`
	TargetResourceType string                              `json:"targetResourceType,omitempty,omitzero" bson:"targetResourceType,omitempty,omitzero"`
	TenantID           string                              `json:"tenantId,omitempty,omitzero" bson:"tenantId,omitempty,omitzero"`
	TenantName         string                              `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	Timestamp          time.Time                           `json:"timestamp,omitempty,omitzero" bson:"timestamp,omitempty,omitzero"`
}

type ResourceChangePropChange struct {
	NewValue      string `json:"newValue" bson:"newValue"`
	PreviousValue string `json:"previousValue" bson:"previousValue"`
}
