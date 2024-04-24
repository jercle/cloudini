package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"
)

type providerListResult struct {
	Value []struct {
		Namespace     string `json:"namespace"`
		ResourceTypes []struct {
			APIProfiles []struct {
				APIVersion     string `json:"apiVersion"`
				ProfileVersion string `json:"profileVersion"`
			} `json:"apiProfiles"`
			APIVersions       []string `json:"apiVersions"`
			Capabilities      string   `json:"capabilities"`
			DefaultAPIVersion string   `json:"defaultApiVersion,omitempty"`
			Locations         []string `json:"locations"`
			ResourceType      string   `json:"resourceType"`
			ZoneMappings      []any    `json:"zoneMappings"`
		} `json:"resourceTypes"`
	} `json:"value"`
}

type azureResourceTypes struct {
	ProviderNamespaces []string
	ResourceTypes      []string
}

func main() {
	var (
		tenantId  = os.Getenv("AZURE_TENANT_ID")
		spDetails lib.CldConfigClientAuthDetails
	)
	spDetails.ClientID = os.Getenv("AZURE_CLIENT_ID")
	spDetails.ClientSecret = os.Getenv("AZURE_CLIENT_SECRET")

	token, err := azure.GetServicePrincipalToken(tenantId, spDetails)
	lib.CheckFatalError(err)

	urlString := ""
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	lib.CheckFatalError(err)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.Token)

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)

	responseBody, err := io.ReadAll(res.Body)
	lib.CheckFatalError(err)

	fmt.Println(string(responseBody))
}
