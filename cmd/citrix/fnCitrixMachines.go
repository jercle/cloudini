package citrix

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jercle/cloudini/lib"
)

func GetAllMachines(creds lib.CitrixCloudAccountConfig, tokenData lib.CitrixTokenData) []CitrixMachine {
	// fmt.Println(time.Now().Add(-11 * time.Hour).Format("2006-01-02T15:04:05.000Z"))
	// fmt.Println(strings.Split(time.Now().Add(-11*time.Hour).Format(time.RFC3339), "+")[0])
	// os.Exit(0)

	// formattedDate := time.Now().Add(-11 * time.Hour).Format("2006-01-02T15:04:05.000Z")
	// formattedDate := time.Now().Add(-11 * time.Hour).Format("2006-01-02T15:04:05.000Z")
	// filterString := "$filter=CollectedDate gt " + formattedDate + "&$orderby=CollectedDate desc"
	fields := "DeliveryGroup,DnsName,Id,MachineCatalog,Name,OSType,SessionSupport,SummaryState,Hosting"
	urlString := "https://api.cloud.com/cvad/manage/Machines?limit=500&fields=" + fields

	res, err := HttpGet(urlString, creds.CustomerId, creds.SiteId, tokenData)
	lib.CheckFatalError(err)

	// fmt.Println(string(res))
	// os.Exit(0)

	var resData GetAllMachinesResponse
	err = json.Unmarshal(res, &resData)
	lib.CheckFatalError(err)

	machines := resData.Items

	nextLink := resData.ContinuationToken

	iterations := 1

	for nextLink != "" {
		iterations++
		if iterations > 10 {
			break
		}
		fmt.Println(iterations)
		var curr GetAllMachinesResponse
		// fmt.Println(nextLink)

		resNew, _ := HttpGet(urlString+"&continuationToken="+nextLink, creds.CustomerId, creds.SiteId, tokenData)
		json.Unmarshal(resNew, &curr)
		// fmt.Println(string(res))

		nextLink = curr.ContinuationToken
		machines = append(machines, curr.Items...)

	}

	var processedMachines []CitrixMachine

	for _, m := range machines {
		curr := m
		nameSpl := strings.Split(m.Name, "\\")
		curr.Name = nameSpl[1]
		hostedMachineSpl := strings.Split(m.Hosting.HostedMachineID, "/")
		curr.AzureResourceGroup = hostedMachineSpl[0]
		curr.Hosting = nil
		processedMachines = append(processedMachines, curr)
	}

	// jsonStr, _ := json.Marshal(processedMachines, jsontext.WithIndent("  "))
	// os.WriteFile("main-citrix-machines.json", jsonStr, 0644)
	return processedMachines
}
