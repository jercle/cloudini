package azure

import (
	"time"

	"encoding/json/jsontext"
)

type GetAzureAlertsResponse struct {
	Count           float64      `json:"count,omitempty,omitzero" bson:"count,omitempty,omitzero"`
	Data            []AzureAlert `json:"data,omitempty,omitzero" bson:"data,omitempty,omitzero"`
	Facets          []any        `json:"facets,omitempty,omitzero" bson:"facets,omitempty,omitzero"`
	ResultTruncated string       `json:"resultTruncated,omitempty,omitzero" bson:"resultTruncated,omitempty,omitzero"`
	TotalRecords    float64      `json:"totalRecords,omitempty,omitzero" bson:"totalRecords,omitempty,omitzero"`
}

//
//

type AzureAlert struct {
	AffectedResource    string         `json:"AffectedResource,omitempty,omitzero" bson:"AffectedResource,omitempty,omitzero"`
	AlertCreated        string         `json:"AlertCreated,omitempty,omitzero" bson:"AlertCreated,omitempty,omitzero"`
	AlertLastModified   string         `json:"AlertLastModified,omitempty,omitzero" bson:"AlertLastModified,omitempty,omitzero"`
	AlertLastModifiedBy string         `json:"AlertLastModifiedBy,omitempty,omitzero" bson:"AlertLastModifiedBy,omitempty,omitzero"`
	AlertState          string         `json:"AlertState,omitempty,omitzero" bson:"AlertState,omitempty,omitzero"`
	Description         string         `json:"Description,omitempty,omitzero" bson:"Description,omitempty,omitzero"`
	UnknownFields       jsontext.Value `json:",unknown"`
	Name                string         `json:"Name,omitempty,omitzero" bson:"Name,omitempty,omitzero"`
	Results             string         `json:"Results,omitempty,omitzero" bson:"Results,omitempty,omitzero"`
	Severity            string         `json:"Severity,omitempty,omitzero" bson:"Severity,omitempty,omitzero"`
	TriageAlert         string         `json:"TriageAlert,omitempty,omitzero" bson:"TriageAlert,omitempty,omitzero"`
	ID                  string         `json:"ID,omitempty,omitzero" bson:"ID,omitempty,omitzero"`
	Properties          struct {
		Context struct {
			AffectedItems      string `json:"affectedItems,omitempty" bson:"affectedItems,omitempty"`
			AlertVersionNumber string `json:"alertVersionNumber,omitempty" bson:"alertVersionNumber,omitempty"`
			Category           string `json:"category,omitempty" bson:"category,omitempty"`
			Context            *struct {
				Condition struct {
					AllOf []struct {
						Dimensions []struct {
							Name  string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
							Value string `json:"value,omitempty,omitzero" bson:"value,omitempty,omitzero"`
						} `json:"dimensions,omitempty,omitzero" bson:"dimensions,omitempty,omitzero"`
						Event          any `json:"event,omitempty,omitzero" bson:"event,omitempty,omitzero"`
						FailingPeriods struct {
							MinFailingPeriodsToAlert  float64 `json:"minFailingPeriodsToAlert,omitempty,omitzero" bson:"minFailingPeriodsToAlert,omitempty,omitzero"`
							NumberOfEvaluationPeriods float64 `json:"numberOfEvaluationPeriods,omitempty,omitzero" bson:"numberOfEvaluationPeriods,omitempty,omitzero"`
						} `json:"failingPeriods,omitempty,omitzero" bson:"failingPeriods,omitempty,omitzero"`
						LinkToFilteredSearchResultsAPI string  `json:"linkToFilteredSearchResultsAPI,omitempty,omitzero" bson:"linkToFilteredSearchResultsAPI,omitempty,omitzero"`
						LinkToFilteredSearchResultsUi  string  `json:"linkToFilteredSearchResultsUI,omitempty,omitzero" bson:"linkToFilteredSearchResultsUI,omitempty,omitzero"`
						LinkToSearchResultsAPI         string  `json:"linkToSearchResultsAPI,omitempty,omitzero" bson:"linkToSearchResultsAPI,omitempty,omitzero"`
						LinkToSearchResultsUi          string  `json:"linkToSearchResultsUI,omitempty,omitzero" bson:"linkToSearchResultsUI,omitempty,omitzero"`
						MetricMeasureColumn            any     `json:"metricMeasureColumn,omitempty,omitzero" bson:"metricMeasureColumn,omitempty,omitzero"`
						MetricValue                    float64 `json:"metricValue,omitempty,omitzero" bson:"metricValue,omitempty,omitzero"`
						Operator                       string  `json:"operator,omitempty,omitzero" bson:"operator,omitempty,omitzero"`
						SearchQuery                    string  `json:"searchQuery,omitempty,omitzero" bson:"searchQuery,omitempty,omitzero"`
						TargetResourceTypes            string  `json:"targetResourceTypes,omitempty,omitzero" bson:"targetResourceTypes,omitempty,omitzero"`
						Threshold                      string  `json:"threshold,omitempty,omitzero" bson:"threshold,omitempty,omitzero"`
						TimeAggregation                string  `json:"timeAggregation,omitempty,omitzero" bson:"timeAggregation,omitempty,omitzero"`
					} `json:"allOf,omitempty,omitzero" bson:"allOf,omitempty,omitzero"`
					WindowEndTime   time.Time `json:"windowEndTime,omitempty,omitzero" bson:"windowEndTime,omitempty,omitzero"`
					WindowSize      string    `json:"windowSize,omitempty,omitzero" bson:"windowSize,omitempty,omitzero"`
					WindowStartTime time.Time `json:"windowStartTime,omitempty,omitzero" bson:"windowStartTime,omitempty,omitzero"`
				} `json:"condition,omitempty,omitzero" bson:"condition,omitempty,omitzero"`
				ConditionType     string    `json:"conditionType,omitempty,omitzero" bson:"conditionType,omitempty,omitzero"`
				Description       string    `json:"description,omitempty,omitzero" bson:"description,omitempty,omitzero"`
				ID                string    `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
				Name              string    `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
				PortalLink        string    `json:"portalLink,omitempty,omitzero" bson:"portalLink,omitempty,omitzero"`
				ResourceGroupName *string   `json:"resourceGroupName,omitempty,omitzero" bson:"resourceGroupName,omitempty,omitzero"`
				ResourceID        string    `json:"resourceId,omitempty,omitzero" bson:"resourceId,omitempty,omitzero"`
				ResourceName      *string   `json:"resourceName,omitempty,omitzero" bson:"resourceName,omitempty,omitzero"`
				ResourceType      *string   `json:"resourceType,omitempty,omitzero" bson:"resourceType,omitempty,omitzero"`
				Severity          string    `json:"severity,omitempty,omitzero" bson:"severity,omitempty,omitzero"`
				SubscriptionID    string    `json:"subscriptionId,omitempty,omitzero" bson:"subscriptionId,omitempty,omitzero"`
				Timestamp         time.Time `json:"timestamp,omitempty,omitzero" bson:"timestamp,omitempty,omitzero"`
			} `json:"context,omitempty" bson:"context,omitempty"`
			ExtendedInfo *struct {
				JobID              string  `json:"jobId,omitempty,omitzero" bson:"jobId,omitempty,omitzero"`
				OccurrenceCount    float64 `json:"occurrenceCount,omitempty,omitzero" bson:"occurrenceCount,omitempty,omitzero"`
				PossibleCauses     any     `json:"possibleCauses,omitempty,omitzero" bson:"possibleCauses,omitempty,omitzero"`
				RecommendedActions string  `json:"recommendedActions,omitempty,omitzero" bson:"recommendedActions,omitempty,omitzero"`
			} `json:"extendedInfo,omitempty" bson:"extendedInfo,omitempty"`
			FirstLevelContainerID  string    `json:"firstLevelContainerId,omitempty" bson:"firstLevelContainerId,omitempty"`
			FormattedSourceType    string    `json:"formattedSourceType,omitempty" bson:"formattedSourceType,omitempty"`
			LinkedResourceID       string    `json:"linkedResourceId,omitempty" bson:"linkedResourceId,omitempty"`
			LinkedResourceName     string    `json:"linkedResourceName,omitempty" bson:"linkedResourceName,omitempty"`
			Properties             *struct{} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
			ProtectedItemID        string    `json:"protectedItemId,omitempty" bson:"protectedItemId,omitempty"`
			SecondLevelContainerID any       `json:"secondLevelContainerId,omitempty" bson:"secondLevelContainerId,omitempty"`
			SourceID               string    `json:"sourceId,omitempty" bson:"sourceId,omitempty"`
			SourceType             string    `json:"sourceType,omitempty" bson:"sourceType,omitempty"`
			SourceVersion          string    `json:"sourceVersion,omitempty" bson:"sourceVersion,omitempty"`
			Status                 string    `json:"status,omitempty" bson:"status,omitempty"`
			Version                string    `json:"version,omitempty" bson:"version,omitempty"`
		} `json:"context,omitempty,omitzero" bson:"context,omitempty,omitzero"`
		CustomProperties *struct{} `json:"customProperties,omitempty" bson:"customProperties,omitempty"`
		Essentials       struct {
			ActionStatus struct {
				IsSuppressed bool `json:"isSuppressed,omitempty,omitzero" bson:"isSuppressed,omitempty,omitzero"`
			} `json:"actionStatus,omitempty,omitzero" bson:"actionStatus,omitempty,omitzero"`
			AlertRule            string    `json:"alertRule,omitempty" bson:"alertRule,omitempty"`
			AlertState           string    `json:"alertState,omitempty,omitzero" bson:"alertState,omitempty,omitzero"`
			Description          string    `json:"description,omitempty,omitzero" bson:"description,omitempty,omitzero"`
			IsStatefulAlert      *bool     `json:"isStatefulAlert,omitempty,omitzero" bson:"isStatefulAlert,omitempty,omitzero"`
			LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty,omitzero" bson:"lastModifiedDateTime,omitempty,omitzero"`
			LastModifiedUserName string    `json:"lastModifiedUserName,omitempty,omitzero" bson:"lastModifiedUserName,omitempty,omitzero"`
			MonitorCondition     string    `json:"monitorCondition,omitempty,omitzero" bson:"monitorCondition,omitempty,omitzero"`
			MonitorService       string    `json:"monitorService,omitempty,omitzero" bson:"monitorService,omitempty,omitzero"`
			Severity             string    `json:"severity,omitempty,omitzero" bson:"severity,omitempty,omitzero"`
			SignalType           string    `json:"signalType,omitempty,omitzero" bson:"signalType,omitempty,omitzero"`
			SourceCreatedID      string    `json:"sourceCreatedId,omitempty,omitzero" bson:"sourceCreatedId,omitempty,omitzero"`
			StartDateTime        time.Time `json:"startDateTime,omitempty,omitzero" bson:"startDateTime,omitempty,omitzero"`
			TargetResource       string    `json:"targetResource,omitempty,omitzero" bson:"targetResource,omitempty,omitzero"`
			TargetResourceGroup  string    `json:"targetResourceGroup,omitempty,omitzero" bson:"targetResourceGroup,omitempty,omitzero"`
			TargetResourceName   string    `json:"targetResourceName,omitempty,omitzero" bson:"targetResourceName,omitempty,omitzero"`
			TargetResourceType   string    `json:"targetResourceType,omitempty,omitzero" bson:"targetResourceType,omitempty,omitzero"`
		} `json:"essentials,omitempty,omitzero" bson:"essentials,omitempty,omitzero"`
	} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
}

//
//

type AzureAlertProcessed struct {
	// UnknownFields                  jsontext.Value `json:",unknown"`
	AzureWorkbookUrl               string    `json:"AzureWorkbookUrl,omitempty,omitzero" bson:"azureWorkbookUrl,omitempty,omitzero"`
	AffectedResource               string    `json:"AffectedResource,omitempty,omitzero" bson:"affectedResource,omitempty,omitzero"`
	AlertCreated                   time.Time `json:"alertCreated,omitempty,omitzero" bson:"alertCreated,omitempty,omitzero"`
	AlertLastModified              time.Time `json:"alertLastModified,omitempty,omitzero" bson:"alertLastModified,omitempty,omitzero"`
	AlertLastModifiedBy            string    `json:"AlertLastModifiedBy,omitempty,omitzero" bson:"alertLastModifiedBy,omitempty,omitzero"`
	AlertState                     string    `json:"AlertState,omitempty,omitzero" bson:"alertState,omitempty,omitzero"`
	Description                    string    `json:"Description,omitempty,omitzero" bson:"description,omitempty,omitzero"`
	Name                           string    `json:"Name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Results                        string    `json:"Results,omitempty,omitzero" bson:"results,omitempty,omitzero"`
	Severity                       string    `json:"Severity,omitempty,omitzero" bson:"severity,omitempty,omitzero"`
	TriageAlert                    string    `json:"triageAlert,omitempty,omitzero" bson:"triageAlert,omitempty,omitzero"`
	ID                             string    `json:"id,omitempty,omitzero" bson:"_id,omitempty,omitzero"`
	LinkToFilteredSearchResultsUi  string    `json:"linkToFilteredSearchResultsUI,omitempty,omitzero" bson:"linkToFilteredSearchResultsUI,omitempty,omitzero"`
	LinkToFilteredSearchResultsAPI string    `json:"linkToFilteredSearchResultsAPI,omitempty,omitzero" bson:"linkToFilteredSearchResultsAPI,omitempty,omitzero"`
	Properties                     struct {
		Context struct {
			AffectedItems      string `json:"affectedItems,omitempty" bson:"affectedItems,omitempty"`
			AlertVersionNumber string `json:"alertVersionNumber,omitempty" bson:"alertVersionNumber,omitempty"`
			Category           string `json:"category,omitempty" bson:"category,omitempty"`
			Context            *struct {
				ConditionType     string    `json:"conditionType,omitempty,omitzero" bson:"conditionType,omitempty,omitzero"`
				Description       string    `json:"description,omitempty,omitzero" bson:"description,omitempty,omitzero"`
				ID                string    `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
				Name              string    `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
				PortalLink        string    `json:"portalLink,omitempty,omitzero" bson:"portalLink,omitempty,omitzero"`
				ResourceGroupName *string   `json:"resourceGroupName,omitempty,omitzero" bson:"resourceGroupName,omitempty,omitzero"`
				ResourceID        string    `json:"resourceId,omitempty,omitzero" bson:"resourceId,omitempty,omitzero"`
				ResourceName      *string   `json:"resourceName,omitempty,omitzero" bson:"resourceName,omitempty,omitzero"`
				ResourceType      *string   `json:"resourceType,omitempty,omitzero" bson:"resourceType,omitempty,omitzero"`
				Severity          string    `json:"severity,omitempty,omitzero" bson:"severity,omitempty,omitzero"`
				SubscriptionID    string    `json:"subscriptionId,omitempty,omitzero" bson:"subscriptionId,omitempty,omitzero"`
				Timestamp         time.Time `json:"timestamp,omitempty,omitzero" bson:"timestamp,omitempty,omitzero"`
			} `json:"context,omitempty" bson:"context,omitempty"`
			ExtendedInfo *struct {
				JobID              string  `json:"jobId,omitempty,omitzero" bson:"jobId,omitempty,omitzero"`
				OccurrenceCount    float64 `json:"occurrenceCount,omitempty,omitzero" bson:"occurrenceCount,omitempty,omitzero"`
				PossibleCauses     any     `json:"possibleCauses,omitempty,omitzero" bson:"possibleCauses,omitempty,omitzero"`
				RecommendedActions string  `json:"recommendedActions,omitempty,omitzero" bson:"recommendedActions,omitempty,omitzero"`
			} `json:"extendedInfo,omitempty" bson:"extendedInfo,omitempty"`
			FirstLevelContainerID  string    `json:"firstLevelContainerId,omitempty" bson:"firstLevelContainerId,omitempty"`
			FormattedSourceType    string    `json:"formattedSourceType,omitempty" bson:"formattedSourceType,omitempty"`
			LinkedResourceID       string    `json:"linkedResourceId,omitempty" bson:"linkedResourceId,omitempty"`
			LinkedResourceName     string    `json:"linkedResourceName,omitempty" bson:"linkedResourceName,omitempty"`
			Properties             *struct{} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
			ProtectedItemID        string    `json:"protectedItemId,omitempty" bson:"protectedItemId,omitempty"`
			SecondLevelContainerID any       `json:"secondLevelContainerId,omitempty" bson:"secondLevelContainerId,omitempty"`
			SourceID               string    `json:"sourceId,omitempty" bson:"sourceId,omitempty"`
			SourceType             string    `json:"sourceType,omitempty" bson:"sourceType,omitempty"`
			SourceVersion          string    `json:"sourceVersion,omitempty" bson:"sourceVersion,omitempty"`
			Status                 string    `json:"status,omitempty" bson:"status,omitempty"`
			Version                string    `json:"version,omitempty" bson:"version,omitempty"`
		} `json:"context,omitempty,omitzero" bson:"context,omitempty,omitzero"`
		Essentials struct {
			ActionStatus struct {
				IsSuppressed bool `json:"isSuppressed,omitempty,omitzero" bson:"isSuppressed,omitempty,omitzero"`
			} `json:"actionStatus,omitempty,omitzero" bson:"actionStatus,omitempty,omitzero"`
			AlertRule            string    `json:"alertRule,omitempty" bson:"alertRule,omitempty"`
			AlertState           string    `json:"alertState,omitempty,omitzero" bson:"alertState,omitempty,omitzero"`
			Description          string    `json:"description,omitempty,omitzero" bson:"description,omitempty,omitzero"`
			IsStatefulAlert      *bool     `json:"isStatefulAlert,omitempty,omitzero" bson:"isStatefulAlert,omitempty,omitzero"`
			LastModifiedDateTime time.Time `json:"lastModifiedDateTime,omitempty,omitzero" bson:"lastModifiedDateTime,omitempty,omitzero"`
			LastModifiedUserName string    `json:"lastModifiedUserName,omitempty,omitzero" bson:"lastModifiedUserName,omitempty,omitzero"`
			MonitorCondition     string    `json:"monitorCondition,omitempty,omitzero" bson:"monitorCondition,omitempty,omitzero"`
			MonitorService       string    `json:"monitorService,omitempty,omitzero" bson:"monitorService,omitempty,omitzero"`
			Severity             string    `json:"severity,omitempty,omitzero" bson:"severity,omitempty,omitzero"`
			SignalType           string    `json:"signalType,omitempty,omitzero" bson:"signalType,omitempty,omitzero"`
			SourceCreatedID      string    `json:"sourceCreatedId,omitempty,omitzero" bson:"sourceCreatedId,omitempty,omitzero"`
			StartDateTime        time.Time `json:"startDateTime,omitempty,omitzero" bson:"startDateTime,omitempty,omitzero"`
			TargetResource       string    `json:"targetResource,omitempty,omitzero" bson:"targetResource,omitempty,omitzero"`
			TargetResourceGroup  string    `json:"targetResourceGroup,omitempty,omitzero" bson:"targetResourceGroup,omitempty,omitzero"`
			TargetResourceName   string    `json:"targetResourceName,omitempty,omitzero" bson:"targetResourceName,omitempty,omitzero"`
			TargetResourceType   string    `json:"targetResourceType,omitempty,omitzero" bson:"targetResourceType,omitempty,omitzero"`
		} `json:"essentials,omitempty,omitzero" bson:"essentials,omitempty,omitzero"`
	} `json:"properties,omitempty,omitzero" bson:"properties,omitempty,omitzero"`
	TenantName       string           `json:"tenantName,omitempty,omitzero" bson:"tenantName,omitempty,omitzero"`
	LastAzureSync    time.Time        `json:"lastAzureSync,omitempty,omitzero" bson:"lastAzureSync,omitempty,omitzero"`
	LastDatabaseSync time.Time        `json:"lastDatabaseSync,omitempty,omitzero" bson:"lastDatabaseSync,omitempty,omitzero"`
	AlertData        []map[string]any `json:"alertData,omitempty,omitzero" bson:"alertData,omitempty,omitzero"`
}

//
//

type GetAlertDataFromSearchResultsLinkResult struct {
	Tables []struct {
		Columns []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"columns"`
		Name string  `json:"name"`
		Rows [][]any `json:"rows"`
	} `json:"tables"`
}

//
//

type LogAnalyticsWorkbook struct {
	Etag     string `json:"etag"`
	ID       string `json:"id"`
	Identity struct {
		Type string `json:"type"`
	} `json:"identity"`
	Kind       string `json:"kind"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		Category       string    `json:"category"`
		DisplayName    string    `json:"displayName"`
		Revision       string    `json:"revision"`
		SerializedData string    `json:"serializedData"`
		SourceID       string    `json:"sourceId"`
		StorageURI     any       `json:"storageUri"`
		Tags           any       `json:"tags"`
		TimeModified   time.Time `json:"timeModified"`
		UserID         string    `json:"userId"`
		Version        string    `json:"version"`
	} `json:"properties"`
	Tags struct {
		HiddenTitle string `json:"hidden-title"`
	} `json:"tags"`
	Type string `json:"type"`
}

//
//

type LogAnalyticsWorkbookSerializedData struct {
	FallbackResourceIds []string `json:"fallbackResourceIds"`
	IsLocked            bool     `json:"isLocked"`
	Items               []struct {
		Content struct {
			CrossComponentResources []string `json:"crossComponentResources"`
			GridSettings            struct {
				Formatters []struct {
					ColumnMatch   string `json:"columnMatch"`
					FormatOptions *struct {
						ArmActionContext *struct {
							Description string `json:"description"`
							Headers     []any  `json:"headers"`
							HTTPMethod  string `json:"httpMethod"`
							Params      []any  `json:"params"`
							Path        string `json:"path"`
							RunLabel    string `json:"runLabel"`
							Title       string `json:"title"`
						} `json:"armActionContext,omitempty"`
						CustomColumnWidthSetting string  `json:"customColumnWidthSetting,omitempty"`
						LinkIsContextBlade       bool    `json:"linkIsContextBlade"`
						LinkLabel                string  `json:"linkLabel,omitempty"`
						LinkTarget               *string `json:"linkTarget"`
						ShowIcon                 bool    `json:"showIcon,omitempty"`
						ThresholdsGrid           []struct {
							Operator       string  `json:"operator"`
							Representation string  `json:"representation"`
							Text           string  `json:"text"`
							ThresholdValue *string `json:"thresholdValue"`
						} `json:"thresholdsGrid,omitempty"`
						ThresholdsOptions string `json:"thresholdsOptions,omitempty"`
					} `json:"formatOptions,omitempty"`
					Formatter float64 `json:"formatter"`
				} `json:"formatters"`
			} `json:"gridSettings"`
			NoDataMessage      string  `json:"noDataMessage"`
			NoDataMessageStyle float64 `json:"noDataMessageStyle"`
			Query              string  `json:"query"`
			QueryType          float64 `json:"queryType"`
			ResourceType       string  `json:"resourceType"`
			ShowRefreshButton  bool    `json:"showRefreshButton"`
			Size               float64 `json:"size"`
			SortBy             []any   `json:"sortBy"`
			Version            string  `json:"version"`
		} `json:"content"`
		Name string  `json:"name"`
		Type float64 `json:"type"`
	} `json:"items"`
	Version string `json:"version"`
}

//
//

type RunLogAnalyticsQueryResponseRaw struct {
	Tables []struct {
		Columns []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"columns"`
		Name string  `json:"name"`
		Rows [][]any `json:"rows"`
	} `json:"tables"`
}

//
//

type LogAnalyticsQueryResponse struct {
	Tables []LogAnalyticsQueryResponseTable `json:"tables"`
}

type LogAnalyticsQueryResponseTable struct {
	Rows []map[string]any `json:"rows"`
	Name string           `json:"name"`
}
