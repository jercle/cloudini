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
	Count           float64                    `json:"count"`
	Data            []lib.AzureResourceDetails `json:"data"`
	Facets          []interface{}              `json:"facets"`
	ResultTruncated string                     `json:"resultTruncated"`
	SkipToken       string                     `json:"$skipToken"`
	TotalRecords    float64                    `json:"totalRecords"`
}

type ResourceGraphResponseDataInterface struct {
	Count           float64       `json:"count"`
	Data            []interface{} `json:"data"`
	Facets          []interface{} `json:"facets"`
	ResultTruncated string        `json:"resultTruncated"`
	SkipToken       string        `json:"$skipToken"`
	TotalRecords    float64       `json:"totalRecords"`
}

//
//

type ListAllResourceGroupsResponse struct {
	Count           float64         `json:"count"`
	Data            []ResourceGroup `json:"data"`
	Facets          []any           `json:"facets"`
	ResultTruncated string          `json:"resultTruncated"`
	TotalRecords    float64         `json:"totalRecords"`
	SkipToken       string          `json:"$skipToken"`
}

//
//

type ResourceGroup struct {
	ExtendedLocation any    `json:"extendedLocation"`
	ID               string `json:"id"`
	Identity         any    `json:"identity"`
	Kind             string `json:"kind"`
	Location         string `json:"location"`
	ManagedBy        string `json:"managedBy"`
	Name             string `json:"name"`
	Plan             any    `json:"plan"`
	Properties       struct {
		ProvisioningState string `json:"provisioningState"`
	} `json:"properties"`
	ResourceGroup  string            `json:"resourceGroup"`
	Sku            any               `json:"sku"`
	SubscriptionID string            `json:"subscriptionId"`
	Tags           map[string]string `json:"tags"`
	TenantID       string            `json:"tenantId"`
	Type           string            `json:"type"`
	Zones          any               `json:"zones"`
	TenantName     string            `json:"tenantName"`
	LastAzureSync  time.Time         `json:"lastAzureSync,omitempty,omitzero" bson:"lastAzureSync,omitempty,omitzero"`
	LastDBSync     time.Time         `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
}
