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
	config := lib.GetCldConfig(nil)
	_ = config
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	// token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{})
	tokenReq, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{GetWriteToken: true})
	lib.CheckFatalError(err)
	token, err := tokenReq.SelectTenant("REDDTQ")
	lib.CheckFatalError(err)
	// _ = tokens
	_ = token

	// subscriptionId := "2ff9367c-2183-4ef6-9ba2-102c2b014d94"
	// resourceGroupName := "rg-apcdtqshared-automon"
	// vmName := "packerbuilder2"

	var (
		wg sync.WaitGroup
	)

	// resourceId2 :=

	resources := []string{
		"/subscriptions/2ff9367c-2183-4ef6-9ba2-102c2b014d94/resourceGroups/rg-apcdtqshared-automon/providers/microsoft.compute/virtualMachines/automon-host",
		// "/subscriptions/2ff9367c-2183-4ef6-9ba2-102c2b014d94/resourceGroups/rg-apcdtqshared-automon/providers/Microsoft.Compute/virtualMachines/packerbuilder2",
	}

	for _, resource := range resources {
		wg.Add(1)
		go func() {
			defer wg.Done()
			StartAzureVm(resource, token)
		}()
	}
	wg.Wait()

	// StopAzureVm(subscriptionId, resourceGroupName, vmName, token)

	// fmt.Println(subId[1])
	// jsonStr, _ := json.MarshalIndent(resourceIdSplit, "", "  ")

	// fmt.Println(string(jsonStr))
	// fmt.Println(string(res))

	elapsed := time.Since(startTime).Truncate(time.Second).String()

	fmt.Println("Completed requests after " + elapsed)
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

	_, resHeader, err := azure.HttpPost(urlString, *mat)
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

	_, resHeader, err := azure.HttpPost(urlString, *mat)
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
