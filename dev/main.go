package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/jercle/azg/cmd/azure"
	"github.com/jercle/azg/lib"
)

type AzgConfig struct {
	Azure struct {
		MultiTenantAuth struct {
			Tenants []struct {
				Reader struct {
					ClientID     string `json:"clientId"`
					ClientSecret string `json:"clientSecret"`
				} `json:"reader"`
				TenantID   string `json:"tenantId"`
				TenantName string `json:"tenantName"`
				Writer     struct {
					ClientID     string `json:"clientId"`
					ClientSecret string `json:"clientSecret"`
				} `json:"writer"`
			} `json:"tenants"`
		} `json:"multiTenantAuth"`
	} `json:"azure"`
}

type MultiAuthToken struct {
	TenantId   string
	TenantName string
	TokenData  azure.TokenData
}
type virtualMachine struct {
	ID         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		AdditionalCapabilities struct {
			HibernationEnabled bool `json:"hibernationEnabled"`
		} `json:"additionalCapabilities"`
		DiagnosticsProfile struct {
			BootDiagnostics struct {
				Enabled bool `json:"enabled"`
			} `json:"bootDiagnostics"`
		} `json:"diagnosticsProfile"`
		HardwareProfile struct {
			VmSize string `json:"vmSize"`
		} `json:"hardwareProfile"`
		NetworkProfile struct {
			NetworkInterfaces []struct {
				ID         string `json:"id"`
				Properties struct {
					DeleteOption string `json:"deleteOption"`
				} `json:"properties"`
			} `json:"networkInterfaces"`
		} `json:"networkProfile"`
		OSProfile struct {
			AdminUsername            string `json:"adminUsername"`
			AllowExtensionOperations bool   `json:"allowExtensionOperations"`
			ComputerName             string `json:"computerName"`
			LinuxConfiguration       struct {
				DisablePasswordAuthentication bool `json:"disablePasswordAuthentication"`
				EnableVmAgentPlatformUpdates  bool `json:"enableVMAgentPlatformUpdates"`
				PatchSettings                 struct {
					AssessmentMode              string `json:"assessmentMode"`
					AutomaticByPlatformSettings struct {
						BypassPlatformSafetyChecksOnUserSchedule bool   `json:"bypassPlatformSafetyChecksOnUserSchedule"`
						RebootSetting                            string `json:"rebootSetting"`
					} `json:"automaticByPlatformSettings"`
					PatchMode string `json:"patchMode"`
				} `json:"patchSettings"`
				ProvisionVmAgent bool `json:"provisionVMAgent"`
				SSH              struct {
					PublicKeys []struct {
						KeyData string `json:"keyData"`
						Path    string `json:"path"`
					} `json:"publicKeys"`
				} `json:"ssh"`
			} `json:"linuxConfiguration"`
			RequireGuestProvisionSignal bool  `json:"requireGuestProvisionSignal"`
			Secrets                     []any `json:"secrets"`
		} `json:"osProfile"`
		ProvisioningState string `json:"provisioningState"`
		SecurityProfile   struct {
			SecurityType string `json:"securityType"`
			UefiSettings struct {
				SecureBootEnabled bool `json:"secureBootEnabled"`
				VTpmEnabled       bool `json:"vTpmEnabled"`
			} `json:"uefiSettings"`
		} `json:"securityProfile"`
		StorageProfile struct {
			DataDisks          []any  `json:"dataDisks"`
			DiskControllerType string `json:"diskControllerType"`
			ImageReference     struct {
				ExactVersion string `json:"exactVersion"`
				Offer        string `json:"offer"`
				Publisher    string `json:"publisher"`
				Sku          string `json:"sku"`
				Version      string `json:"version"`
			} `json:"imageReference"`
			OSDisk struct {
				Caching      string  `json:"caching"`
				CreateOption string  `json:"createOption"`
				DeleteOption string  `json:"deleteOption"`
				DiskSizeGb   float64 `json:"diskSizeGB"`
				ManagedDisk  struct {
					StorageAccountType string `json:"storageAccountType"`
				} `json:"managedDisk"`
				OSType string `json:"osType"`
			} `json:"osDisk"`
		} `json:"storageProfile"`
		TimeCreated time.Time `json:"timeCreated"`
		VmID        string    `json:"vmId"`
	} `json:"properties"`
	Type  string   `json:"type"`
	Zones []string `json:"zones"`
}

type vmReqResponseBody struct {
	Value []virtualMachine `json:"value"`
}

// Lists all Virtual Machines in a given subscription
func getAllSubscriptionVMs(token azure.MultiAuthToken, subscriptionId string) []virtualMachine {
	// https://learn.microsoft.com/en-us/rest/api/compute/virtual-machines/list-all?view=rest-compute-2023-10-02&tabs=HTTP
	// GET https://management.azure.com/subscriptions/{subscriptionId}/providers/Microsoft.Compute/virtualMachines?api-version=2023-09-01
	// cred, err := azidentity.NewDefaultAzureCredential(nil)
	// subscriptionId := "bae338c7-6098-4d52-b173-e2147e107dfa"
	// var allVirtualMachines virtualMachines
	// ctx := context.Background()
	// tokenRequestOptions := policy.TokenRequestOptions{
	// 	Scopes: []string{
	// 		"https://management.core.windows.net/.default",
	// 	},
	// }

	// token, err := cred.GetToken(ctx, tokenRequestOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/providers/Microsoft.Compute/virtualMachines?api-version=2023-09-01"
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.TokenData.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error fetching LA Workspace Tables")
	}

	responseBody, err := io.ReadAll(res.Body)
	if res.StatusCode == 400 {
		log.Fatal("Error fetching LA Workspace Tables: ", string(responseBody))
	}
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// fmt.Println(string(responseBody))

	var responseBodyUnmarshal vmReqResponseBody
	json.Unmarshal(responseBody, &responseBodyUnmarshal)
	allVirtualMachines := responseBodyUnmarshal.Value
	return allVirtualMachines
}

// Lists Azure subscriptions availabe to a given auth token
func listSubscriptions(token azure.MultiAuthToken) ([]azure.FetchedSubscription, error) {
	urlString := "https://management.azure.com/subscriptions?api-version=2022-12-01"
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		log.Fatal(err)
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

	var subsList azure.SubsReqResBody
	json.Unmarshal(responseBody, &subsList)
	subsList.UpdateTenantName(token.TenantName)
	lib.MarshalAndPrintJson(subsList.Value)

	return subsList.Value, nil
}

func main() {
	// tenantTokens, err := azure.GetAllTenantTokens(azure.GetAllTenantTokenOptions{})
	// _ = tenantTokens
	// lib.CheckFatalError(err)

	// lib.MarshalAndPrintJson(tenantTokens)
	// azure.GetCachedTokens()

	lib.GetCldConfig(lib.CldConfigOptions{})

}

// func main() {
// 	var (
// 		allSubscriptions []azure.FetchedSubscription
// 		wg               sync.WaitGroup
// 		mut              sync.Mutex
// 	)
// 	_ = allSubscriptions
// 	tenantTokens, err := azure.GetAllTenantTokens(azure.GetAllTenantTokenOptions{})
// 	_ = tenantTokens
// 	lib.CheckFatalError(err)

// 	for _, tt := range tenantTokens {
// 		wg.Add(1)
// 		go func() {
// 			tenantSubscriptions, err := listSubscriptions(tt)
// 			lib.CheckFatalError(err)
// 			mut.Lock()
// 			allSubscriptions = append(allSubscriptions, tenantSubscriptions...)
// 			mut.Unlock()
// 			wg.Done()
// 		}()

// 	}
// 	wg.Wait()
// }
