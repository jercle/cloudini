package azure

import (
	"encoding/json/v2"
	"sync"

	"github.com/jercle/cloudini/lib"
)

func GetAllVnetsWithCapacities(token lib.AzureMultiAuthToken) (allVnets []Vnet) {
	allSubs, err := ListSubscriptions(token)
	lib.CheckFatalError(err)

	var (
		wg  sync.WaitGroup
		mux sync.Mutex
	)

	for _, sub := range allSubs {
		wg.Go(func() {
			subVnets := ListAllSubscriptionVnets(sub.SubscriptionID, token)
			lib.JsonMarshalAndPrint(subVnets)
			for _, vnet := range subVnets {
				vnet.TenantName = token.TenantName
				vnet.SubscriptionName = sub.DisplayName
				mux.Lock()
				allVnets = append(allVnets, vnet)
				mux.Unlock()
			}
		})
	}

	wg.Wait()

	return
}

func GetSubnetCapacitiesForVnet(vnetId string, mat *lib.AzureMultiAuthToken) []SubnetCapacities {

	urlString := "https://management.azure.com" + vnetId + "/usages?api-version=2025-09-01"

	res, err := HttpGet(urlString, *mat)
	lib.CheckFatalError(err)

	var rsp SubnetCapacitiesResponse

	json.Unmarshal(res, &rsp)

	return rsp.Value
}

type SubnetCapacitiesResponse struct {
	Value []SubnetCapacities `json:"value"`
}

type SubnetCapacities struct {
	ParentVnet   string  `json:"parentVnet"`
	CurrentValue float64 `json:"currentValue"`
	ID           string  `json:"id"`
	IsAdjustable bool    `json:"isAdjustable"`
	Limit        float64 `json:"limit"`
	Name         struct {
		LocalizedValue string `json:"localizedValue"`
		Value          string `json:"value"`
	} `json:"name"`
	Unit string `json:"unit"`
}
