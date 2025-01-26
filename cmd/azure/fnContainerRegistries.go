package azure

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/jercle/cloudini/lib"
)

type ContainerRegistry struct {
	ID         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		AdminUserEnabled      bool      `json:"adminUserEnabled"`
		AnonymousPullEnabled  bool      `json:"anonymousPullEnabled"`
		CreationDate          time.Time `json:"creationDate"`
		DataEndpointEnabled   bool      `json:"dataEndpointEnabled"`
		DataEndpointHostNames []string  `json:"dataEndpointHostNames"`
		Encryption            struct {
			Status string `json:"status"`
		} `json:"encryption"`
		LoginServer              string `json:"loginServer"`
		NetworkRuleBypassOptions string `json:"networkRuleBypassOptions"`
		NetworkRuleSet           struct {
			DefaultAction string `json:"defaultAction"`
			IpRules       []any  `json:"ipRules"`
		} `json:"networkRuleSet"`
		Policies struct {
			AzureAdAuthenticationAsArmPolicy struct {
				Status string `json:"status"`
			} `json:"azureADAuthenticationAsArmPolicy"`
			ExportPolicy struct {
				Status string `json:"status"`
			} `json:"exportPolicy"`
			QuarantinePolicy struct {
				Status string `json:"status"`
			} `json:"quarantinePolicy"`
			RetentionPolicy struct {
				Days            int       `json:"days"`
				LastUpdatedTime time.Time `json:"lastUpdatedTime"`
				Status          string    `json:"status"`
			} `json:"retentionPolicy"`
			SoftDeletePolicy struct {
				LastUpdatedTime time.Time `json:"lastUpdatedTime"`
				RetentionDays   int       `json:"retentionDays"`
				Status          string    `json:"status"`
			} `json:"softDeletePolicy"`
			TrustPolicy struct {
				Status string `json:"status"`
				Type   string `json:"type"`
			} `json:"trustPolicy"`
		} `json:"policies"`
		PrivateEndpointConnections []struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			Properties struct {
				PrivateEndpoint struct {
					ID string `json:"id"`
				} `json:"privateEndpoint"`
				PrivateLinkServiceConnectionState struct {
					Description string `json:"description"`
					Status      string `json:"status"`
				} `json:"privateLinkServiceConnectionState"`
				ProvisioningState string `json:"provisioningState"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"privateEndpointConnections"`
		ProvisioningState   string `json:"provisioningState"`
		PublicNetworkAccess string `json:"publicNetworkAccess"`
		ZoneRedundancy      string `json:"zoneRedundancy"`
	} `json:"properties"`
	Sku struct {
		Name string `json:"name"`
		Tier string `json:"tier"`
	} `json:"sku"`
	SystemData struct {
		CreatedAt          time.Time `json:"createdAt"`
		CreatedBy          string    `json:"createdBy"`
		CreatedByType      string    `json:"createdByType"`
		LastModifiedAt     time.Time `json:"lastModifiedAt"`
		LastModifiedBy     string    `json:"lastModifiedBy"`
		LastModifiedByType string    `json:"lastModifiedByType"`
	} `json:"systemData"`
	Tags struct{} `json:"tags"`
	Type string   `json:"type"`
}

type ContainerRegisryResponse struct {
	Value []ContainerRegistry
}

type ContainerRegistriesPerSub struct {
	SubscriptionName    string
	SubscriptionID      string
	TenantName          string
	TenantID            string
	NumACRs             int
	ContainerRegistries []ContainerRegistry
}

func listAllTenantACRs(tenantName string, tenantId string, token *lib.AzureTokenData) []ContainerRegistry {
	var (
		allTenantACRs []ContainerRegistry
		wg            sync.WaitGroup
		mutex         sync.Mutex
	)
	subsToken := lib.AzureMultiAuthToken{
		TenantName: tenantName,
		TenantId:   tenantId,
		TokenData:  *token,
	}

	subsList, err := listSubscriptions(lib.AzureMultiAuthToken(subsToken))
	lib.CheckFatalError(err)

	for _, sub := range subsList {
		wg.Add(1)
		go func() {
			defer wg.Done()
			subACRs := listSubscriptionACRs(tenantId, sub.SubscriptionID, token)
			mutex.Lock()
			allTenantACRs = append(allTenantACRs, subACRs...)
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	return allTenantACRs
}

func listAllTenantACRsBySub(tenantName string, tenantId string, token *lib.AzureTokenData) []ContainerRegistriesPerSub {
	var (
		allTenantACRs []ContainerRegistriesPerSub
		wg            sync.WaitGroup
		mutex         sync.Mutex
	)
	subsToken := lib.AzureMultiAuthToken{
		TenantName: tenantName,
		TenantId:   tenantId,
		TokenData:  *token,
	}

	subsList, err := listSubscriptions(lib.AzureMultiAuthToken(subsToken))
	lib.CheckFatalError(err)

	for _, sub := range subsList {
		wg.Add(1)
		go func() {
			subACRs := listSubscriptionACRs(tenantId, sub.SubscriptionID, token)
			var currentSub ContainerRegistriesPerSub
			currentSub.ContainerRegistries = make([]ContainerRegistry, 0)
			currentSub.SubscriptionName = sub.DisplayName
			currentSub.SubscriptionID = sub.SubscriptionID
			currentSub.TenantName = sub.TenantName
			currentSub.TenantID = sub.TenantID
			currentSub.NumACRs = len(subACRs)

			if len(subACRs) != 0 {
				currentSub.ContainerRegistries = subACRs
			}

			mutex.Lock()
			allTenantACRs = append(allTenantACRs, currentSub)
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	return allTenantACRs
}

// Get all Azure Container Registries for subscription
func listSubscriptionACRs(tenantId string, subscriptionId string, token *lib.AzureTokenData) []ContainerRegistry {
	urlString := "https://management.com/subscriptions/" +
		subscriptionId +
		"/providers/Microsoft.ContainerRegistry/registries?api-version=2023-01-01-preview"

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.Token)

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)

	responseBody, err := io.ReadAll(res.Body)
	lib.CheckFatalError(err)

	var regReqBody ContainerRegisryResponse
	err = json.Unmarshal(responseBody, &regReqBody)
	lib.CheckFatalError(err)

	var allSubscriptionContainerRegistries []ContainerRegistry

	for _, registry := range regReqBody.Value {
		allSubscriptionContainerRegistries = append(allSubscriptionContainerRegistries, registry)
	}

	return allSubscriptionContainerRegistries
}

func listSubscriptions(token lib.AzureMultiAuthToken) ([]lib.FetchedSubscription, error) {
	urlString := "https://management.azure.com/subscriptions?api-version=2022-12-01"
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.TokenData.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// log.Fatal("Error fetching list of Subscriptions")
		return nil, err
	}

	responseBody, err := io.ReadAll(res.Body)
	if res.StatusCode == 400 {
		// log.Fatal("Error fetching list of Subscriptions: ", string(responseBody))
		return nil, err
	}
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}
	defer res.Body.Close()

	var subsList lib.SubsReqResBody
	json.Unmarshal(responseBody, &subsList)
	subsList.UpdateTenantName(token.TenantName)

	return subsList.Value, nil
}
