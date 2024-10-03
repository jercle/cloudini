package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

func main() {
	startTime := time.Now()
	_ = startTime
	config := lib.GetCldConfig(nil)
	_ = config
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	// token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{})
	tokenReq, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{GetWriteToken: true}, nil)
	lib.CheckFatalError(err)
	token, err := tokenReq.SelectTenant("REDDTQ")
	lib.CheckFatalError(err)
	// _ = tokens
	_ = token

	var (
		wg sync.WaitGroup
	)

	// subscriptionId := ""
	// resourceGroup := ""
	resourceString := "/subscriptions/" + subscriptionId + "/resourceGroups/" + resourceGroup + "/providers/Microsoft.Compute/virtualMachines/"
	_ = resourceString

	startResources := []string{
		// resourceString + "test-runner01",
		// resourceString + "test-runner02",
		// resourceString + "test-runner03",
		// resourceString + "test-runner04",
		// resourceString + "test-crt",
	}

	deallocateResources := []string{
		// resourceString + "test-runner01",
		// resourceString + "test-runner02",
		// resourceString + "test-runner03",
		// resourceString + "test-runner04",
		// resourceString + "test-crt",
		// resourceString + "test-runner-tmp",
	}

	getStatusOfResources := []string{
		// resourceString + "test-runner01",
		// resourceString + "test-runner02",
		// resourceString + "test-runner03",
		// resourceString + "test-runner04",
		// resourceString + "test-crt",
		// resourceString + "test-runner-tmp",
	}

	for _, resource := range startResources {
		wg.Add(1)
		go func() {
			defer wg.Done()
			StartAzureVm(resource, token)
			// DeallocateAzureVm(resource, token)
		}()
	}
	for _, resource := range deallocateResources {
		wg.Add(1)
		go func() {
			defer wg.Done()
			DeallocateAzureVm(resource, token)
		}()

	}
	for _, resource := range getStatusOfResources {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetAzureVMStatus(resource, token)
		}()

	}
	wg.Wait()

	// StopAzureVm(subscriptionId, resourceGroupName, vmName, token)

	// fmt.Println(subId[1])
	// jsonStr, _ := json.MarshalIndent(resourceIdSplit, "", "  ")

	// fmt.Println(string(jsonStr))
	// fmt.Println(string(res))

	// elapsed := time.Since(startTime).Truncate(time.Second).String()

	// fmt.Println("Completed requests after " + elapsed)
}

func GetAzureVMStatus(resourceId string, mat *lib.MultiAuthToken) {
	resourceIdSplit := strings.Split(resourceId, "/")
	subscriptionId := resourceIdSplit[2]
	resourceGroupName := resourceIdSplit[4]
	vmName := resourceIdSplit[8]

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resourceGroups/" +
		resourceGroupName +
		"/providers/Microsoft.Compute/virtualMachines/" +
		vmName +
		"/instanceView?api-version=2024-03-01"

	res, err := azure.HttpGet(urlString, *mat)
	// _, resHeader, err := azure.HttpPost(urlString, nil, *mat)
	lib.CheckFatalError(err)

	// var reqRes AzureAsyncRequestResponse

	var vmData VirtualMachineInstanceView

	json.Unmarshal(res, &vmData)

	// updRes, err := azure.HttpGet(reqRes.AzureAsyncoperation[0], *mat)
	// lib.CheckFatalError(err)

	// var asyncOpStatus AzureAsyncOpUpdateResponse
	// json.Unmarshal(res, &asyncOpStatus)

	// fmt.Println("Getting status of VM: " + vmName)

	// for asyncOpStatus.Status == "InProgress" {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("In progress - Start " + vmName)
	// 	updRes, err = azure.HttpGet(reqRes.AzureAsyncoperation[0], *mat)
	// 	lib.CheckFatalError(err)
	// 	json.Unmarshal(updRes, &asyncOpStatus)
	// }

	// fmt.Println("Start command " + asyncOpStatus.Status + " for " + vmName)

	fmt.Println(vmName)
	jsonStr, _ := json.MarshalIndent(vmData.Statuses, "", "  ")
	fmt.Println(string(jsonStr))
	// fmt.Println(string(res))
}

func StartAzureVms(resourceIds []string, mat *lib.MultiAuthToken) {

}
func StopAzureVms(resourceIds []string, mat *lib.MultiAuthToken) {

}

func StartAzureVm(resourceId string, mat *lib.MultiAuthToken) {
	resourceIdSplit := strings.Split(resourceId, "/")
	subscriptionId := resourceIdSplit[2]
	resourceGroupName := resourceIdSplit[4]
	vmName := resourceIdSplit[8]

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resourceGroups/" +
		resourceGroupName +
		"/providers/Microsoft.Compute/virtualMachines/" +
		vmName +
		"/start?api-version=2024-03-01"

	_, resHeader, err := azure.HttpPost(urlString, "", *mat)
	lib.CheckFatalError(err)

	var reqRes AzureAsyncRequestResponse

	json.Unmarshal(resHeader, &reqRes)

	updRes, err := azure.HttpGet(reqRes.AzureAsyncoperation[0], *mat)
	lib.CheckFatalError(err)

	var asyncOpStatus AzureAsyncOpUpdateResponse
	json.Unmarshal(updRes, &asyncOpStatus)

	fmt.Println("Start command started for " + vmName)

	for asyncOpStatus.Status == "InProgress" {
		time.Sleep(2 * time.Second)
		fmt.Println("In progress - Start " + vmName)
		updRes, err = azure.HttpGet(reqRes.AzureAsyncoperation[0], *mat)
		lib.CheckFatalError(err)
		json.Unmarshal(updRes, &asyncOpStatus)
	}

	fmt.Println("Start command " + asyncOpStatus.Status + " for " + vmName)
}

func GetAzureVm(subscriptionId string, resourceGroupName string, vmName string, mat *lib.MultiAuthToken) {
	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resourceGroups/" +
		resourceGroupName +
		"/providers/Microsoft.Compute/virtualMachines/" +
		vmName +
		"?api-version=2024-03-01"

	res, err := azure.HttpGet(urlString, *mat)
	lib.CheckFatalError(err)

	fmt.Println(string(res))
}

func StopAzureVm(resourceId string, mat *lib.MultiAuthToken) {
	resourceIdSplit := strings.Split(resourceId, "/")
	subscriptionId := resourceIdSplit[2]
	resourceGroupName := resourceIdSplit[4]
	vmName := resourceIdSplit[8]

	// https://www.google.com/search?q=golang+url+param&oq=golang+url+param&gs_lcrp=EgZjaHJvbWUyBggAEEUYOTIHCAEQABiABDIHCAIQABiABDIHCAMQABiABDIHCAQQABiABDIHCAUQABiABDIHCAYQABiABDIHCAcQABiABDIHCAgQABiABNIBCDIwMzhqMGoxqAIAsAIA&sourceid=chrome&ie=UTF-8
	apiVersion := "api-version=2024-03-01"

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resourceGroups/" +
		resourceGroupName +
		"/providers/Microsoft.Compute/virtualMachines/" +
		vmName +
		"/powerOff?" +
		apiVersion

	_, resHeader, err := azure.HttpPost(urlString, "", *mat)
	lib.CheckFatalError(err)

	var reqRes AzureAsyncRequestResponse

	json.Unmarshal(resHeader, &reqRes)

	updRes, err := azure.HttpGet(reqRes.AzureAsyncoperation[0], *mat)
	lib.CheckFatalError(err)

	var asyncOpStatus AzureAsyncOpUpdateResponse
	json.Unmarshal(updRes, &asyncOpStatus)

	fmt.Println("Stop command started for " + vmName)

	for asyncOpStatus.Status == "InProgress" {
		time.Sleep(2 * time.Second)
		fmt.Println("In progress - Stop " + vmName)
		updRes, err = azure.HttpGet(reqRes.AzureAsyncoperation[0], *mat)
		lib.CheckFatalError(err)
		json.Unmarshal(updRes, &asyncOpStatus)
	}

	fmt.Println("Stop command " + asyncOpStatus.Status + " for " + vmName)
}

func DeallocateAzureVm(resourceId string, mat *lib.MultiAuthToken) {
	resourceIdSplit := strings.Split(resourceId, "/")
	subscriptionId := resourceIdSplit[2]
	resourceGroupName := resourceIdSplit[4]
	vmName := resourceIdSplit[8]

	apiVersion := "api-version=2024-03-01"

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resourceGroups/" +
		resourceGroupName +
		"/providers/Microsoft.Compute/virtualMachines/" +
		vmName +
		"/deallocate?" +
		apiVersion

	_, resHeader, err := azure.HttpPost(urlString, "", *mat)
	lib.CheckFatalError(err)

	var reqRes AzureAsyncRequestResponse

	json.Unmarshal(resHeader, &reqRes)

	updRes, err := azure.HttpGet(reqRes.AzureAsyncoperation[0], *mat)
	lib.CheckFatalError(err)

	var asyncOpStatus AzureAsyncOpUpdateResponse
	json.Unmarshal(updRes, &asyncOpStatus)

	fmt.Println("Deallocate command started for " + vmName)

	for asyncOpStatus.Status == "InProgress" {
		time.Sleep(2 * time.Second)
		fmt.Println("In progress - Deallocate " + vmName)
		updRes, err = azure.HttpGet(reqRes.AzureAsyncoperation[0], *mat)
		lib.CheckFatalError(err)
		json.Unmarshal(updRes, &asyncOpStatus)
	}

	fmt.Println("Deallocate command " + asyncOpStatus.Status + " for " + vmName)
}

type AzureAsyncRequestResponse struct {
	AzureAsyncnotification                        []string `json:"Azure-Asyncnotification"`
	AzureAsyncoperation                           []string `json:"Azure-Asyncoperation"`
	CacheControl                                  []string `json:"Cache-Control"`
	ContentLength                                 []string `json:"Content-Length"`
	Date                                          []string `json:"Date"`
	Expires                                       []string `json:"Expires"`
	Location                                      []string `json:"Location"`
	Pragma                                        []string `json:"Pragma"`
	StrictTransportSecurity                       []string `json:"Strict-Transport-Security"`
	XCache                                        []string `json:"X-Cache"`
	XContentTypeOptions                           []string `json:"X-Content-Type-Options"`
	XMsCorrelationRequestID                       []string `json:"X-Ms-Correlation-Request-Id"`
	XMsRatelimitRemainingResource                 []string `json:"X-Ms-Ratelimit-Remaining-Resource"`
	XMsRatelimitRemainingSubscriptionGlobalWrites []string `json:"X-Ms-Ratelimit-Remaining-Subscription-Global-Writes"`
	XMsRatelimitRemainingSubscriptionWrites       []string `json:"X-Ms-Ratelimit-Remaining-Subscription-Writes"`
	XMsRequestID                                  []string `json:"X-Ms-Request-Id"`
	XMsRoutingRequestID                           []string `json:"X-Ms-Routing-Request-Id"`
	XMsedgeRef                                    []string `json:"X-Msedge-Ref"`
}

type AzureAsyncOpUpdateResponse struct {
	EndTime   time.Time `json:"endTime"`
	Name      string    `json:"name"`
	StartTime time.Time `json:"startTime"`
	Status    string    `json:"status"`
}

type VirtualMachine struct {
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

type VirtualMachineInstanceView struct {
	BootDiagnostics struct{} `json:"bootDiagnostics"`
	ComputerName    string   `json:"computerName"`
	Disks           []struct {
		Name     string `json:"name"`
		Statuses []struct {
			Code          string    `json:"code"`
			DisplayStatus string    `json:"displayStatus"`
			Level         string    `json:"level"`
			Time          time.Time `json:"time"`
		} `json:"statuses"`
	} `json:"disks"`
	Extensions []struct {
		Name     string `json:"name"`
		Statuses []struct {
			Code          string `json:"code"`
			DisplayStatus string `json:"displayStatus"`
			Level         string `json:"level"`
			Message       string `json:"message"`
		} `json:"statuses"`
		Type               string `json:"type"`
		TypeHandlerVersion string `json:"typeHandlerVersion"`
	} `json:"extensions"`
	HyperVGeneration string `json:"hyperVGeneration"`
	OSName           string `json:"osName"`
	OSVersion        string `json:"osVersion"`
	PatchStatus      struct {
		AvailablePatchSummary struct {
			AssessmentActivityID          string  `json:"assessmentActivityId"`
			CriticalAndSecurityPatchCount float64 `json:"criticalAndSecurityPatchCount"`
			Error                         struct {
				Code    string `json:"code"`
				Details []struct {
					Code    string `json:"code"`
					Message string `json:"message"`
				} `json:"details"`
				Message string `json:"message"`
			} `json:"error"`
			LastModifiedTime time.Time `json:"lastModifiedTime"`
			OtherPatchCount  float64   `json:"otherPatchCount"`
			RebootPending    bool      `json:"rebootPending"`
			StartTime        time.Time `json:"startTime"`
			Status           string    `json:"status"`
		} `json:"availablePatchSummary"`
		ConfigurationStatuses []struct {
			Code          string    `json:"code"`
			DisplayStatus string    `json:"displayStatus"`
			Level         string    `json:"level"`
			Time          time.Time `json:"time"`
		} `json:"configurationStatuses"`
	} `json:"patchStatus"`
	Statuses []struct {
		Code          string    `json:"code"`
		DisplayStatus string    `json:"displayStatus"`
		Level         string    `json:"level"`
		Time          time.Time `json:"time,omitempty"`
	} `json:"statuses"`
	VmAgent struct {
		ExtensionHandlers []struct {
			Status struct {
				Code          string `json:"code"`
				DisplayStatus string `json:"displayStatus"`
				Level         string `json:"level"`
				Message       string `json:"message"`
			} `json:"status"`
			Type               string `json:"type"`
			TypeHandlerVersion string `json:"typeHandlerVersion"`
		} `json:"extensionHandlers"`
		Statuses []struct {
			Code          string    `json:"code"`
			DisplayStatus string    `json:"displayStatus"`
			Level         string    `json:"level"`
			Message       string    `json:"message"`
			Time          time.Time `json:"time"`
		} `json:"statuses"`
		VmAgentVersion string `json:"vmAgentVersion"`
	} `json:"vmAgent"`
}
