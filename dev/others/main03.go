package main

import (
	"context"
	"fmt"

	// "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armcostmanagement"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
	"github.com/charmbracelet/log"
)

// GetActiveSub(), GetAzureIdentity()

func GetAzureIdentity() *azidentity.DefaultAzureCredential {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		// handle error
		log.Error("", err, err)
	}

	return cred
}

func main() {
	// activeSub, err := azure.GetActiveSub()
	// if err != nil {
	// 	log.Error(err)
	// }

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	client := clientFactory.NewQueryClient()

	query := armcostmanagement.QueryDefinition{
		Type: to.Ptr(armcostmanagement.ExportTypeUsage),
		Dataset: &armcostmanagement.QueryDataset{
			// Filter: &armcostmanagement.QueryFilter{
			// 	And: []*armcostmanagement.QueryFilter{
			// 		{
			// 			Or: []*armcostmanagement.QueryFilter{
			// 				{
			// 					Dimensions: &armcostmanagement.QueryComparisonExpression{
			// 						Name:     to.Ptr("ResourceLocation"),
			// 						Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
			// 						Values: []*string{
			// 							to.Ptr("East US"),
			// 							to.Ptr("West Europe")},
			// 					},
			// 				},
			// 				{
			// 					Tags: &armcostmanagement.QueryComparisonExpression{
			// 						Name:     to.Ptr("Environment"),
			// 						Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
			// 						Values: []*string{
			// 							to.Ptr("UAT"),
			// 							to.Ptr("Prod")},
			// 					},
			// 				}},
			// 		},
			// 		{
			// 			Dimensions: &armcostmanagement.QueryComparisonExpression{
			// 				Name:     to.Ptr("ResourceGroup"),
			// 				Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
			// 				Values: []*string{
			// 					to.Ptr("API")},
			// 			},
			// 		}},
			// },
			Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
		},
		Timeframe: to.Ptr(armcostmanagement.TimeframeTypeMonthToDate),
	}

	managementGroup := "ea508bc7-b43c-4b96-8470-489756e59a14"

	res, err := client.Usage(
		ctx,
		"/providers/Microsoft.Management/managementGroups/"+managementGroup,
		query,
		nil,
	)
	if err != nil {
		log.Error(err)
	}

	result, err := res.QueryResult.MarshalJSON()
	if err != nil {
		log.Error(err)
	}

	// var responseObject interface{}

	// json.Unmarshal(res, &responseObject)

	fmt.Println(string(result))
	// armBillingOptions := armbilling.SubscriptionsClientGetOptions{}
	// clientFactory, err :=
	// armsubscriptions.NewClientFactory(azure.GetActiveSub(), GetAzureIdentity(), &options)

	// client, _ := armbilling.NewSubscriptionsClient(activeSub, GetAzureIdentity(), &options)

	// fmt.Println(client)

	// res, err := client.Get(ctx, "D 6462", &armBillingOptions)

	// client.

	// fmt.Println(res)

}
