package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type virtualMachines []struct {
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

type ResponseBody struct {
	Value virtualMachines `json:"value"`
}

type TokenData struct {
	Token     string
	ExpiresOn string
}

func getToken() TokenData {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	tokenRequestOptions := policy.TokenRequestOptions{
		Scopes: []string{
			"https://management.core.windows.net/.default",
		},
	}

	tokenResponse, err := cred.GetToken(ctx, tokenRequestOptions)
	if err != nil {
		log.Fatal(err)
	}

	token := TokenData{
		Token:     tokenResponse.Token,
		ExpiresOn: tokenResponse.ExpiresOn.Local().String(),
	}
	return token
}

func getAllSubscriptionVMs(cred *azidentity.DefaultAzureCredential, subscriptionId string) virtualMachines {
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
	token := getToken()
	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/providers/Microsoft.Compute/virtualMachines?api-version=2023-09-01"
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.Token)

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

	fmt.Println(string(responseBody))

	var responseBodyUnmarshal ResponseBody
	json.Unmarshal(responseBody, &responseBodyUnmarshal)
	allVirtualMachines := responseBodyUnmarshal.Value
	return allVirtualMachines
}

func listSubscriptions(cred *azidentity.DefaultAzureCredential) interface{} {
	urlString := "https://management.azure.com/subscriptions?api-version=2022-12-01"
	token := getToken()
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.Token)

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
}

func main() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(err)
	}
	// subscriptionId := "bae338c7-6098-4d52-b173-e2147e107dfa"

	subsList := listSubscriptions(cred)

	// allVMs := getAllSubscriptionVMs(cred, subscriptionId)
	// _ = allVMs

	// for _, vm := range allVMs {
	// 	_ = vm
	// 	// fmt.Println(vm.Name)
	// }

	// jsonData, err := json.MarshalIndent(allVMs, "", "  ")

}
