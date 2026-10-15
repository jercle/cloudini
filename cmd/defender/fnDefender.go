package defender

import (
	"encoding/json/v2"
	"sync"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func GetMachinesForTenant(token *lib.AzureMultiAuthToken) (allMachines []Machine) {
	urlString := "https://api.security.microsoft.com/api/machines"

	res, err := azure.HttpGet(urlString, *token)
	// _, resHeader, err := HttpPost(urlString, nil, *mat)
	lib.CheckFatalError(err)

	var resData ArmResponse[Machine]

	err = json.Unmarshal(res, &resData)
	lib.CheckFatalError(err)

	nextLink := resData.NextLink
	// lib.JsonMarshalAndPrint(resData.Value)

	rawMachinesMap := make(map[string]Machine)

	for _, m := range resData.Value {
		curr := m
		curr.TenantName = token.TenantName
		// curr.MongoId = token.TenantName + "-" + m.ComputerDnsName
		curr.MongoId = m.ComputerDnsName

		// allMachines = append(allMachines, curr)
		if val, ok := rawMachinesMap[curr.MongoId]; !ok {
			// Key does not exist
			// fmt.Println("Key not found")
			rawMachinesMap[curr.MongoId] = curr
		} else {
			if m.LastSeen.After(val.LastSeen) {
				rawMachinesMap[curr.MongoId] = curr
			}
		}
	}

	count := 1

	for nextLink != "" {
		count++
		var currentSet ArmResponse[Machine]
		response, _ := azure.HttpGet(nextLink, *token)
		// os.WriteFile("outputs/entraApps-"+strconv.Itoa(count)+".json", response, 0644)
		json.Unmarshal(response, &currentSet)
		nextLink = currentSet.NextLink
		// entraApps = append(entraApps, currentSet.Value...)

		for _, m := range currentSet.Value {
			curr := m
			curr.TenantName = token.TenantName
			// curr.MongoId = token.TenantName + "-" + m.ComputerDnsName
			curr.MongoId = m.ComputerDnsName
			if val, ok := rawMachinesMap[curr.MongoId]; !ok {
				// Key does not exist
				// fmt.Println("Key not found")
				rawMachinesMap[curr.MongoId] = curr
			} else {
				if m.LastSeen.After(val.LastSeen) {
					rawMachinesMap[curr.MongoId] = curr
				}
			}
			// allMachines = append(allMachines, curr)
		}
		// fmt.Println(count)
	}

	for _, m := range rawMachinesMap {
		allMachines = append(allMachines, m)
	}

	return
}

func GetMachinesForAllConfiguredTenants() (allMachines []Machine) {
	var (
		wg  sync.WaitGroup
		mut sync.Mutex
	)

	cldConfig := lib.GetCldConfig(nil)

	azTenants := cldConfig.Azure.MultiTenantAuth.Tenants

	for tName, tConfig := range azTenants {
		if !tConfig.CheckDefender {
			continue
		}

		wg.Go(func() {
			token, err := azure.GetTenantSPToken(lib.AzureMultiAuthTokenRequestOptions{
				Scope: "defender",
				// GetWriteToken: true,
				TenantName: tName,
			}, nil)
			lib.CheckFatalError(err)

			machines := GetMachinesForTenant(token)

			mut.Lock()
			allMachines = append(allMachines, machines...)
			mut.Unlock()
		})
	}

	wg.Wait()

	return
}
