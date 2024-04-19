package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
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

type allVmReqResponseBody struct {
	Value []virtualMachine `json:"value"`
}

// type singleVm

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

	var responseBodyUnmarshal allVmReqResponseBody
	json.Unmarshal(responseBody, &responseBodyUnmarshal)
	allVirtualMachines := responseBodyUnmarshal.Value
	return allVirtualMachines
}

func main() {

	// config := lib.GetCldConfig(lib.CldConfigOptions{})
	// sub, err := azure.GetActiveSub()
	// lib.CheckFatalError(err)

	// var allSecrets

	// config.AddAzureTenant("e9f4bce2-7308-461a-91ce-3f663d079f47", "FakeTest")
	tokens, err := azure.GetAllTenantTokens(azure.AzureRequestOptions{})
	// fmt.Println(tokens)
	lib.CheckFatalError(err)

	for _, tenant := range tokens {
		subs, err := listSubscriptions(tenant)
		lib.CheckFatalError(err)
		for _, sub := range subs {
			lib.PrintSrcLoc("getting kvs")
			GetAllSubscriptionKeyvaults(tenant, sub.ID, sub.DisplayName)
		}
	}
}

type KeyVault struct {
	ID         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		AccessPolicies               []any  `json:"accessPolicies"`
		EnableRbacAuthorization      bool   `json:"enableRbacAuthorization"`
		EnableSoftDelete             bool   `json:"enableSoftDelete"`
		EnabledForDeployment         bool   `json:"enabledForDeployment"`
		EnabledForDiskEncryption     bool   `json:"enabledForDiskEncryption"`
		EnabledForTemplateDeployment bool   `json:"enabledForTemplateDeployment"`
		ProvisioningState            string `json:"provisioningState"`
		PublicNetworkAccess          string `json:"publicNetworkAccess"`
		Sku                          struct {
			Family string `json:"family"`
			Name   string `json:"name"`
		} `json:"sku"`
		SoftDeleteRetentionInDays float64 `json:"softDeleteRetentionInDays"`
		TenantID                  string  `json:"tenantId"`
		VaultURI                  string  `json:"vaultUri"`
	} `json:"properties"`
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

type GetSubscriptionKeyvaultsResponse struct {
	NextLink string        `json:"nextLink"`
	Value    []interface{} `json:"value"`
}

func GetAllSubscriptionKeyvaults(token azure.MultiAuthToken, subscriptionId string, subscriptionName string) {
	var allKeyvaults interface{}
	// urlString := "https://management.azure.com" + subscriptionId + "/resources?$filter=resourceType eq 'Microsoft.KeyVault/vaults'&api-version=2015-11-01" // resources query
	urlString := "https://management.azure.com" + subscriptionId + "/providers/Microsoft.KeyVault/vaults?api-version=2022-07-01" // keyvault api
	// fmt.Println(subscriptionId)
	// fmt.Println(urlString)

	lib.PrintSrcLoc(subscriptionName)
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.TokenData.Token)

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)

	responseBody, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	fmt.Println(string(responseBody))

	var unmarshedBody GetSubscriptionKeyvaultsResponse

	json.Unmarshal(responseBody, unmarshedBody)

	for _, kv := range unmarshedBody.Value {
		allKeyvaults = append(allKeyvaults, kv)
		fmt.Println(kv)
	}

	// lib.MarshalAndPrintJson(allKeyvaults)

	fmt.Println(allKeyvaults)

}

func GetVmDetails(options azure.AzureRequestOptions, token string) {
	var urlString string

	if options.ResourceId != "" {
		urlString = "https://management.azure.com" + options.ResourceId + "?api-version=2023-09-01"
	} else {
		urlString = "https://management.azure.com/subscriptions/" +
			options.SubscriptionId +
			"resourceGroups/" +
			options.ResourceGroupName +
			"/providers/Microsoft.Compute/virtualMachines/" +
			options.ResourceName +
			"?api-version=2023-09-01"
	}

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error fetching LA Workspace Tables")
	}

	responseBody, err := io.ReadAll(res.Body)
	if res.StatusCode == 400 {
		log.Fatal("Error fetching LA Workspace Tables: ", string(responseBody))
	}
	if err != nil {
		lib.CheckFatalError(err)
	}
	defer res.Body.Close()

	fmt.Println(string(responseBody))
	// var responseBodyUnmarshal vmReqResponseBody
	// json.Unmarshal(responseBody, &responseBodyUnmarshal)
	// vmDetails := responseBodyUnmarshal.Value
	// return vmDetails
}
