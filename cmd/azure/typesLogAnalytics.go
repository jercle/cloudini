package azure

import "time"

type GetAzureAlertsResponse struct {
	Count           float64      `json:"count,omitempty,omitzero" bson:"count,omitempty,omitzero"`
	Data            []AzureAlert `json:"data,omitempty,omitzero" bson:"data,omitempty,omitzero"`
	Facets          []any        `json:"facets,omitempty,omitzero" bson:"facets,omitempty,omitzero"`
	ResultTruncated string       `json:"resultTruncated,omitempty,omitzero" bson:"resultTruncated,omitempty,omitzero"`
	TotalRecords    float64      `json:"totalRecords,omitempty,omitzero" bson:"totalRecords,omitempty,omitzero"`
}

type AzureAlert struct {
	AffectedResource    string `json:"affectedResource,omitempty,omitzero" bson:"affectedResource,omitempty,omitzero"`
	AlertCreated        string `json:"alertCreated,omitempty,omitzero" bson:"alertCreated,omitempty,omitzero"`
	AlertLastModified   string `json:"alertLastModified,omitempty,omitzero" bson:"alertLastModified,omitempty,omitzero"`
	AlertLastModifiedBy string `json:"alertLastModifiedBy,omitempty,omitzero" bson:"alertLastModifiedBy,omitempty,omitzero"`
	AlertState          string `json:"alertState,omitempty,omitzero" bson:"alertState,omitempty,omitzero"`
	Description         string `json:"description,omitempty,omitzero" bson:"description,omitempty,omitzero"`
	Name                string `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Results             string `json:"results,omitempty,omitzero" bson:"results,omitempty,omitzero"`
	Severity            string `json:"severity,omitempty,omitzero" bson:"severity,omitempty,omitzero"`
	TriageAlert         string `json:"triageAlert,omitempty,omitzero" bson:"triageAlert,omitempty,omitzero"`
	ID                  string `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
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

type AzureAlertProcessed struct {
	AffectedResource               string    `json:"affectedResource,omitempty,omitzero" bson:"affectedResource,omitempty,omitzero"`
	AlertCreated                   time.Time `json:"alertCreated,omitempty,omitzero" bson:"alertCreated,omitempty,omitzero"`
	AlertLastModified              time.Time `json:"alertLastModified,omitempty,omitzero" bson:"alertLastModified,omitempty,omitzero"`
	AlertLastModifiedBy            string    `json:"alertLastModifiedBy,omitempty,omitzero" bson:"alertLastModifiedBy,omitempty,omitzero"`
	AlertState                     string    `json:"alertState,omitempty,omitzero" bson:"alertState,omitempty,omitzero"`
	Description                    string    `json:"description,omitempty,omitzero" bson:"description,omitempty,omitzero"`
	Name                           string    `json:"name,omitempty,omitzero" bson:"name,omitempty,omitzero"`
	Results                        string    `json:"results,omitempty,omitzero" bson:"results,omitempty,omitzero"`
	Severity                       string    `json:"severity,omitempty,omitzero" bson:"severity,omitempty,omitzero"`
	TriageAlert                    string    `json:"triageAlert,omitempty,omitzero" bson:"triageAlert,omitempty,omitzero"`
	ID                             string    `json:"id,omitempty,omitzero" bson:"id,omitempty,omitzero"`
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
