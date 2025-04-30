package azure

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/jercle/cloudini/lib"
)

func GetAzureVMStatus(resourceId string, mat *lib.AzureMultiAuthToken) {
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

	res, err := HttpGet(urlString, *mat)
	// _, resHeader, err := HttpPost(urlString, nil, *mat)
	if err != nil {

		if strings.Contains(err.Error(), "ARMResourceNotFoundFix") {
			fmt.Println(vmName + ": Does not exist")
			return
		} else {
			lib.CheckFatalError(err)
		}
	}

	// var reqRes AzureAsyncRequestResponse

	var vmData VirtualMachineInstanceView

	json.Unmarshal(res, &vmData)

	// updRes, err := HttpGet(reqRes.AzureAsyncoperation[0], *mat)
	// lib.CheckFatalError(err)

	// var asyncOpStatus AzureAsyncOpUpdateResponse
	// json.Unmarshal(res, &asyncOpStatus)

	// fmt.Println("Getting status of VM: " + vmName)

	// for asyncOpStatus.Status == "InProgress" {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("In progress - Start " + vmName)
	// 	updRes, err = HttpGet(reqRes.AzureAsyncoperation[0], *mat)
	// 	lib.CheckFatalError(err)
	// 	json.Unmarshal(updRes, &asyncOpStatus)
	// }

	// fmt.Println("Start command " + asyncOpStatus.Status + " for " + vmName)
	status := ""
	for _, s := range vmData.Statuses {
		if strings.Contains(s.Code, "PowerState") {
			status = s.DisplayStatus
		}
	}

	fmt.Println(vmName+":", status)
	// jsonStr, _ := json.MarshalIndent(vmData.Statuses, "", "  ")
	// fmt.Println(string(jsonStr))
	// fmt.Println(string(res))
}

//
//

func StopAzureVms(resourceIds []string, mat *lib.AzureMultiAuthToken) {
}

//
//

func StartAzureVm(resourceId string, mat *lib.AzureMultiAuthToken) {
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

	_, resHeader, err := HttpPost(urlString, "", *mat)
	lib.CheckFatalError(err)

	var reqRes AzureAsyncRequestResponse

	json.Unmarshal(resHeader, &reqRes)

	// lib.JsonMarshalAndPrint(reqRes)
	// fmt.Println(string(resBody))

	updRes, err := HttpGet(reqRes.AzureAsyncoperation[0], *mat)
	lib.CheckFatalError(err)

	var asyncOpStatus AzureAsyncOpUpdateResponse
	json.Unmarshal(updRes, &asyncOpStatus)

	fmt.Println("Start command started for " + vmName)

	for asyncOpStatus.Status == "InProgress" {
		time.Sleep(2 * time.Second)
		fmt.Println("In progress - Start " + vmName)
		updRes, err = HttpGet(reqRes.AzureAsyncoperation[0], *mat)
		lib.CheckFatalError(err)
		json.Unmarshal(updRes, &asyncOpStatus)
	}

	fmt.Println("Start command " + asyncOpStatus.Status + " for " + vmName)
}

//
//

func GetAzureVm(subscriptionId string, resourceGroupName string, vmName string, mat *lib.AzureMultiAuthToken) {
	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resourceGroups/" +
		resourceGroupName +
		"/providers/Microsoft.Compute/virtualMachines/" +
		vmName +
		"?api-version=2024-03-01"

	res, err := HttpGet(urlString, *mat)
	lib.CheckFatalError(err)

	fmt.Println(string(res))
}

//
//

func StopAzureVm(resourceId string, mat *lib.AzureMultiAuthToken) {
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

	_, resHeader, err := HttpPost(urlString, "", *mat)
	lib.CheckFatalError(err)

	var reqRes AzureAsyncRequestResponse

	json.Unmarshal(resHeader, &reqRes)

	updRes, err := HttpGet(reqRes.AzureAsyncoperation[0], *mat)
	lib.CheckFatalError(err)

	var asyncOpStatus AzureAsyncOpUpdateResponse
	json.Unmarshal(updRes, &asyncOpStatus)

	fmt.Println("Stop command started for " + vmName)

	for asyncOpStatus.Status == "InProgress" {
		time.Sleep(2 * time.Second)
		fmt.Println("In progress - Stop " + vmName)
		updRes, err = HttpGet(reqRes.AzureAsyncoperation[0], *mat)
		lib.CheckFatalError(err)
		json.Unmarshal(updRes, &asyncOpStatus)
	}

	fmt.Println("Stop command " + asyncOpStatus.Status + " for " + vmName)
}

//
//

func DeallocateAzureVm(resourceId string, mat *lib.AzureMultiAuthToken) {
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
		"/deallocate?" +
		apiVersion

		// export http_proxy=http://127.0.0.1:3128
		// export https_proxy=http://127.0.0.1:3128
		// export ftp_proxy=http://127.0.0.1:3128

	_, resHeader, err := HttpPost(urlString, "", *mat)
	lib.CheckFatalError(err)

	var reqRes AzureAsyncRequestResponse

	json.Unmarshal(resHeader, &reqRes)

	updRes, err := HttpGet(reqRes.AzureAsyncoperation[0], *mat)
	lib.CheckFatalError(err)

	var asyncOpStatus AzureAsyncOpUpdateResponse
	json.Unmarshal(updRes, &asyncOpStatus)

	fmt.Println("Deallocate command started for " + vmName)

	for asyncOpStatus.Status == "InProgress" {
		time.Sleep(2 * time.Second)
		fmt.Println("In progress - Deallocate " + vmName)
		updRes, err = HttpGet(reqRes.AzureAsyncoperation[0], *mat)
		lib.CheckFatalError(err)
		json.Unmarshal(updRes, &asyncOpStatus)
	}

	fmt.Println("Deallocate command " + asyncOpStatus.Status + " for " + vmName)
}

//
//

func GetMultipleAzureVMStatuses(resourceIds []string, token *lib.AzureMultiAuthToken) {
	var wg sync.WaitGroup
	for _, resource := range resourceIds {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetAzureVMStatus(resource, token)
		}()
	}
	wg.Wait()
}

//
//

func StartMultipleAzureVms(resourceIds []string, token *lib.AzureMultiAuthToken) {
	var wg sync.WaitGroup
	// lib.JsonMarshalAndPrint(resourceIds)
	// os.Exit(0)
	for _, resource := range resourceIds {
		wg.Add(1)
		go func() {
			defer wg.Done()
			StartAzureVm(resource, token)
			// DeallocateAzureVm(resource, token)
		}()
	}
	wg.Wait()
}

//
//

func StopMultipleAzureVms(resourceIds []string, token *lib.AzureMultiAuthToken) {
	var wg sync.WaitGroup
	for _, resource := range resourceIds {
		wg.Add(1)
		go func() {
			defer wg.Done()
			StopAzureVm(resource, token)
			// DeallocateAzureVm(resource, token)
		}()
	}
	wg.Wait()
}

//
//

func DeallocateMultipleAzureVms(resourceIds []string, token *lib.AzureMultiAuthToken) {
	var wg sync.WaitGroup
	for _, resource := range resourceIds {
		wg.Add(1)
		go func() {
			defer wg.Done()
			DeallocateAzureVm(resource, token)
			// DeallocateAzureVm(resource, token)
		}()
	}
	wg.Wait()
}

//
//
