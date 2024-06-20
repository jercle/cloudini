// Get vm images
// https://learn.microsoft.com/en-us/rest/api/compute/virtual-machine-images/list?view=rest-compute-2024-03-01&tabs=HTTP

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

type Microsoft365Endpoint struct {
	Category               string   `json:"category"`
	ExpressRoute           bool     `json:"expressRoute"`
	ID                     float64  `json:"id"`
	Ips                    []string `json:"ips,omitempty"`
	Notes                  string   `json:"notes,omitempty"`
	Required               bool     `json:"required"`
	ServiceArea            string   `json:"serviceArea"`
	ServiceAreaDisplayName string   `json:"serviceAreaDisplayName"`
	TcpPorts               string   `json:"tcpPorts,omitempty"`
	UdpPorts               string   `json:"udpPorts,omitempty"`
	Urls                   []string `json:"urls,omitempty"`
	Protocols              []string `json:"protocols,omitempty"`
}

func main() {
	startTime := time.Now()
	config := lib.GetCldConfig(nil)
	_ = config
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	// token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{})
	tokenReq, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	lib.CheckFatalError(err)
	token, err := tokenReq.SelectTenant("REDDTQ")
	lib.CheckFatalError(err)
	// _ = tokens
	_ = token

	webserviceUrl := "https://endpoints.office.com"
	clientRequestId := ""
	tenantName := ""

	// Get the latest versioning data from the Office 365 IP Address and URL web service
	// urlString := webserviceUrl +
	// 	"/version/Worldwide?ClientRequestId=" +
	// 	clientRequestId

	// Query the Office 365 IP Address and URL web service for new data
	urlString := webserviceUrl + "/endpoints/Worldwide?NoIPv6=true&ClientRequestId=" +
		clientRequestId +
		"&TenantName=" +
		tenantName

	res, err := azure.HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	var (
		endpointList       []Microsoft365Endpoint
		processedEndpoints []Microsoft365Endpoint
		networkEndpoints   []Microsoft365Endpoint
		appEndpoints       []Microsoft365Endpoint
	)

	// networkEndpoints := []Microsoft365Endpoint{}

	json.Unmarshal(res, &endpointList)

	for _, item := range endpointList {
		if item.TcpPorts != "" {
			item.Protocols = append(item.Protocols, "TCP")
		}

		if item.UdpPorts != "" {
			item.Protocols = append(item.Protocols, "UDP")
		}

		// fmt.Println(len(item.Ips))

		if len(item.Ips) != 0 {
			// fmt.Println(len(item.Ips))
			networkEndpoints = append(networkEndpoints, item)
			// jsonstr, _ := json.MarshalIndent(item, "", "  ")
			// fmt.Println(string(jsonstr))
		}

		if len(item.Urls) != 0 {
			// fmt.Println(len(item.Urls))
			appEndpoints = append(appEndpoints, item)
		}

		processedEndpoints = append(processedEndpoints, item)
	}

	jsonStr, err := json.MarshalIndent(processedEndpoints, "", "  ")
	lib.CheckFatalError(err)

	fmt.Println(string(jsonStr))
	// fmt.Println(string(res))

	// var imageList ListVirtualMachineImagesResponse
	// json.Unmarshal(res, &imageList)

	// for _, image := range imageList {
	// 	fmt.Println(image.Name)
	// }

	elapsed := time.Since(startTime)
	_ = elapsed
	// fmt.Println(elapsed)
}
