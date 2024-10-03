// Sample Code from https://learn.microsoft.com/en-us/samples/azure-samples/azure-cxp-developer-support/go-sample-code-to-filter-and-list-the-azure-virtual-machines-in-a-subscription-based-on-resource-tags/

package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

var (
	subscriptionID string
	TenantID       string
)

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

func main() {
	// cred := getToken()
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	// tokenRequestOptions := policy.TokenRequestOptions{
	// 	Scopes: []string{
	// 		"https://management.core.windows.net/.default",
	// 	},
	// }
	// To configure DefaultAzureCredential to authenticate a user-assigned managed identity,
	// set the environment variable AZURE_CLIENT_ID to the identity's client ID.
	//
	// clientID := "XXXXXXXXXXXXXXXX"
	// clientSecret := "XXXXXXXXXXXXXXXX"
	// tenantID := "XXXXXXXXXXXXXXXX"

	// Create and authorize a ResourceGraph client

	client, err := armresourcegraph.NewClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	subscriptionId := ""

	// Create the query request, Run the query and get the results. Update the Tags and subscriptionID details below.
	results, err := client.Resources(ctx,
		armresourcegraph.QueryRequest{
			Query: to.Ptr("Resources | where type =~ 'Microsoft.Compute/virtualMachines' | project name, type, location, properties"),
			// Query: to.Ptr("Resources | where type =~ 'Microsoft.Compute/virtualMachines' and tags.Environment=~ 'Production' | project name, type, location, properties"),
			Subscriptions: []*string{
				to.Ptr(subscriptionId)},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	} else {
		// Print the obtained query results
		fmt.Printf("Resources found: " + strconv.FormatInt(*results.TotalRecords, 10) + "\n")
		fmt.Printf("Results: " + fmt.Sprint(results.Data) + "\n")
	}

}
