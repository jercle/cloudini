// Get vm images
// https://learn.microsoft.com/en-us/rest/api/compute/virtual-machine-images/list?view=rest-compute-2024-03-01&tabs=HTTP

package main

import (
	"encoding/json"
	// "fmt"
	"os"
	"slices"
	"time"

	"github.com/3th1nk/cidr"
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

type ProxyConfigTypes struct {
	HostEndpoints []string
	IpEndpoints   []CustomIPEndpoint
}

type ProxyEndpointConfig struct {
	Direct  ProxyConfigTypes
	Proxied ProxyConfigTypes
}

type CustomEndpoint struct {
	Comment       string `json:"Comment,omitempty"`
	Hostname      string `json:"Hostname,omitempty"`
	Network       string `json:"Network,omitempty"`
	ProxyVariable string `json:"ProxyVariable,omitempty"`
	SubnetMask    string `json:"SubnetMask,omitempty"`
	Type          string `json:"Type"`
}

type CustomHostnameEndpoint struct {
	Comment       string `json:"comment,omitempty"`
	ProxyVariable string `json:"proxyVariable,omitempty"`
	Hostname      string `json:"hostname"`
}

type CustomIPEndpoint struct {
	Comment       string `json:"comment,omitempty"`
	ProxyVariable string `json:"proxyVariable,omitempty"`
	Network       string `json:"network"`
	SubnetMask    string `json:"subnetMask"`
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

	categories := []string{
		"Optimize",
		"Allow",
	}

	webserviceUrl := "https://endpoints.office.com"
	clientRequestId := "b10c5ed1-bad1-445f-b386-b919946339a7"
	tenantName := "REDDTQ"

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
		endpointList []Microsoft365Endpoint
		// requiredEndpoints []Microsoft365Endpoint
		// directEndpoints ProxyConfigEndpoints
		// proxyEndpoints  ProxyConfigEndpoints
		// customEndpoints []CustomEndpoint
		proxyConfig ProxyEndpointConfig
	)

	// networkEndpoints := []Microsoft365Endpoint{}

	json.Unmarshal(res, &endpointList)

	for _, item := range endpointList {
		if slices.Contains(categories, item.Category) {
			for _, url := range item.Urls {
				proxyConfig.Direct.HostEndpoints = append(proxyConfig.Direct.HostEndpoints, url)
			}
			for _, ip := range item.Ips {
				var ep CustomIPEndpoint
				c, _ := cidr.Parse(ip)
				ep.Network = c.IP().String()
				ep.SubnetMask = c.Mask().String()
				proxyConfig.Direct.IpEndpoints = append(proxyConfig.Direct.IpEndpoints, ep)

			}
		} else {
			for _, url := range item.Urls {
				proxyConfig.Proxied.HostEndpoints = append(proxyConfig.Proxied.HostEndpoints, url)
			}
			for _, ip := range item.Ips {
				var ep CustomIPEndpoint
				c, _ := cidr.Parse(ip)
				ep.Network = c.IP().String()
				ep.SubnetMask = c.Mask().String()
				proxyConfig.Proxied.IpEndpoints = append(proxyConfig.Proxied.IpEndpoints, ep)
			}
		}
	}

	customHostnameEndpoints, customIpEndpoints := GetCustomRules("proxyfile.params.json")
	for _, item := range customHostnameEndpoints {
		if item.ProxyVariable == "direct" {
			proxyConfig.Direct.HostEndpoints = append(proxyConfig.Direct.HostEndpoints, item.Hostname)
		} else {
			proxyConfig.Proxied.HostEndpoints = append(proxyConfig.Proxied.HostEndpoints, item.Hostname)
		}
	}

	for _, item := range customIpEndpoints {
		if item.ProxyVariable == "direct" {
			var ep CustomIPEndpoint
			ep.Network = item.Network
			ep.SubnetMask = item.SubnetMask
			proxyConfig.Direct.IpEndpoints = append(proxyConfig.Direct.IpEndpoints, ep)
		} else {
			var ep CustomIPEndpoint
			ep.Network = item.Network
			ep.SubnetMask = item.SubnetMask
			proxyConfig.Proxied.IpEndpoints = append(proxyConfig.Proxied.IpEndpoints, ep)
		}
	}

	// fmt.Println(proxyConfig)
	// jsonStr, _ := json.MarshalIndent(proxyConfig, "", "  ")
	// fmt.Println(string(jsonStr))
	// for

	// templatePath := "proxy.tmpl"
	// t, err := template.New("proxy.tmpl").ParseFiles(templatePath)
	// _ = t
	// lib.CheckFatalError(err)
	// t.Execute(os.Stdout, proxyConfig)
	// lib.err

	// fmt.Println(customEndpoints)
	// _ = requiredEndpoints
	// fmt.Println(directEndpoints)

	// jsonStr, err := json.MarshalIndent(processedEndpoints, "", "  ")
	// lib.CheckFatalError(err)

	// for _, ep := range appEndpoints {
	// 	var _ Microsoft365Endpoint
	// 	// var rule FirewallPolicyRuleCollectionRule
	// 	// json.Unmarshal(endpoint, &)
	// 	if slices.Contains(strings.Split(ep.TcpPorts, ","), "80") {
	// 		fmt.Println(ep)
	// 	}

	// }

	// fmt.Println(string(jsonStr))
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

func GetCustomRules(filename string) ([]CustomHostnameEndpoint, []CustomIPEndpoint) {
	var (
		endpoints         []CustomEndpoint
		hostnameEndpoints []CustomHostnameEndpoint
		ipEndpoints       []CustomIPEndpoint
	)

	byteValue, err := os.ReadFile(filename)
	lib.CheckFatalError(err)

	// decodedBytes, err := b64.StdEncoding.DecodeString(string(byteValue))
	// CheckFatalError(err)

	// fmt.Println(string(byteValue))
	// os.Exit(0)

	// err = json.Unmarshal(decodedBytes, &config)
	err = json.Unmarshal(byteValue, &endpoints)
	lib.CheckFatalError(err)

	for _, item := range endpoints {
		if item.Type == "host" {
			var ep CustomHostnameEndpoint
			ep.Comment = item.Comment
			ep.ProxyVariable = item.ProxyVariable
			ep.Hostname = item.Hostname
			hostnameEndpoints = append(hostnameEndpoints, ep)
		}
		if item.Type == "hostIp" {
			var ep CustomIPEndpoint
			ep.Comment = item.Comment
			ep.ProxyVariable = item.ProxyVariable
			ep.Network = item.Network
			ep.SubnetMask = item.SubnetMask
			ipEndpoints = append(ipEndpoints, ep)
		}
	}

	return hostnameEndpoints, ipEndpoints
}
