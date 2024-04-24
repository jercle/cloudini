package main

import (
	"encoding/json"
	"sync"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func main() {
	tokens, err := azure.GetAllTenantSPTokens(azure.MultiAuthTokenRequestOptions{})
	lib.CheckFatalError(err)
	// redToken, _ := tokens.SelectTenant("REDDTQ")
	// _ = redToken

	// allSubs, err :=
	// lib.CheckFatalError(err)

	var (
	// allVnets []azure.Vnet
	// emptySubnets []ProcessedSubnet

	)

	// allTenantEmptySubnets := map[string][]ProcessedSubnet{}

	// for _, token := range tokens {
	// 	allSubs, _ := azure.ListSubscriptions(token)
	// 	for _, sub := range allSubs {
	// 		subVnets := azure.ListAllSubscriptionVnets(sub.SubscriptionID, token)
	// 		// allVnets = append(allVnets, subVnets...)
	// 		for _, vnet := range subVnets {
	// 			vnetSubnets := ListVnetSubnets(sub.SubscriptionID, vnet.ResourceGroup, vnet.Name, token)
	// 			for _, subnet := range vnetSubnets {
	// 				if len(subnet.Properties.IpConfigurations) == 0 && len(subnet.Properties.Delegations) == 0 {
	// 					var currentSubnet ProcessedSubnet
	// 					jsonStr, _ := json.Marshal(subnet)
	// 					json.Unmarshal(jsonStr, &currentSubnet)
	// 					currentSubnet.VnetName = vnet.Name
	// 					allTenantEmptySubnets[token.TenantName] = append(allTenantEmptySubnets[token.TenantName], currentSubnet)
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	// jsonStr, _ := json.MarshalIndent(allTenantEmptySubnets, "", "  ")
	// fmt.Println(string(jsonStr))

}

func ListAllTenantEmptySubnets(token azure.MultiAuthToken) azure.IPAddressList {

	var (
		wg           lib.WaitGroupCount
		mutex        sync.Mutex
		vnets        = make(chan azure.Vnet, 100)
		vnetSubnets  = make(chan azure.ProcessedSubnet, 10000)
		emptySubnets []azure.ProcessedSubnet
		done         = make(chan bool, 1)
	)

	allSubs, err := azure.ListSubscriptions(token)
	lib.CheckFatalError(err)

	for _, sub := range allSubs {
		wg.Add(1)
		go azure.ListAllSubscriptionVnetsWithChan(sub.SubscriptionID, token, vnets, &wg)
	}

	go func() {
		// chanLoop:
		for {
			select {
			case vnet := <-vnets:
				wg.Add(1)
				go ListVnetSubnetsToChan(token, vnet, vnetSubnets, &wg)
			case subnet := <-vnetSubnets:
				if len(subnet.Properties.Delegations) == 0 && len(subnet.Properties.IpConfigurations) == 0 {
					this
					mutex.Lock()
					allIpAddresses.PublicAddresses = append(allIpAddresses.PublicAddresses, publicIp)
					mutex.Unlock()
				}
			// case privateIp := <-privateIpAddresses:
			case <-done:
				break
				// break chanLoop
			}
		}
	}()

	wg.Wait()
	done <- true
	return allIpAddresses
}

// mat MultiAuthToken, vnet Vnet, publicIps chan<- IPAddressItem, privateIps chan<- IPAddressItem, wg *lib.WaitGroupCount) {
func ListVnetSubnetsToChan(mat azure.MultiAuthToken, vnet azure.Vnet, subnets chan<- azure.ProcessedSubnet, wg *lib.WaitGroupCount) {
	var listSubnetResponse azure.ListSubnetsResponse

	urlString := "https://management.azure.com/subscriptions/" +
		vnet.SubscriptionID +
		"/resourceGroups/" +
		vnet.ResourceGroup +
		"/providers/Microsoft.Network/virtualNetworks/" +
		vnet.Name +
		"/subnets?api-version=2023-09-01"

	response := azure.HttpGet(urlString, mat)
	json.Unmarshal(response, &listSubnetResponse)

	for _, subnet := range listSubnetResponse.Value {
		var currentSubnet
		subnets <- subnet
	}

	wg.Done()
}
