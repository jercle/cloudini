package main

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/jercle/cloudini/cmd/azure"
	"github.com/jercle/cloudini/lib"

	graph "github.com/microsoftgraph/msgraph-sdk-go"
	graphidentitygovernance "github.com/microsoftgraph/msgraph-sdk-go/identitygovernance"
)

func main() {
	// startTime := time.Now()
	config := lib.GetCldConfig(nil)
	_ = config
	// tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	// lib.CheckFatalError(err)
	// token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{})
	// tokenReq, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
	tokenReq, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{
		Scope:         "graph",
		GetWriteToken: true,
	}, nil)
	lib.CheckFatalError(err)
	token, err := tokenReq.SelectTenant("REDDTQ")
	lib.CheckFatalError(err)
	// _ = tokens
	_ = token

	// urlString := ""
	// urlString := "https://graph.microsoft.com/v1.0/devices"
	// urlString := "https://graph.microsoft.com/beta/identityGovernance/privilegedAccess/group/eligibilitySchedules"
	// urlString := "https://graph.microsoft.com/v1.0/identityGovernance/privilegedAccess/group/eligibilitySchedules?$filter=groupId eq '2a9694fb-117c-43f4-9c41-49d83dc0a82b' and principalId eq 'f24915f7-3c26-4e24-a18d-97a43a1ba103'"
	// urlString := "https://graph.microsoft.com/beta/identityGovernance/privilegedAccess/group/eligibilitySchedules?$filter=groupId eq '2a9694fb-117c-43f4-9c41-49d83dc0a82b'"
	// urlString := "https://graph.microsoft.com/beta/identityGovernance/privilegedAccess/group/eligibilitySchedules?$filter=groupId eq 'f24915f7-3c26-4e24-a18d-97a43a1ba103'"
	// https://graph.microsoft.com/v1.0/identityGovernance/privilegedAccess/group/eligibilitySchedules?$filter=groupId eq '2a9694fb-117c-43f4-9c41-49d83dc0a82b'
	// var opts azidentity.ClientSecretCredentialOptions

	// cred, _ := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
	// 	TenantID: config.Azure.MultiTenantAuth.Tenants["RED"].TenantID,
	// 	ClientID: config.Azure.MultiTenantAuth.Tenants["RED"].Writer.ClientID,
	// 	// ClientSecret: config.Azure.MultiTenantAuth.Tenants["RED"].Writer.ClientSecret,
	// 	UserPrompt: func(ctx context.Context, message azidentity.DeviceCodeMessage) error {
	// 		fmt.Println(message.Message)
	// 		return nil
	// 	},
	// })

	fmt.Println("start")

	cred, _ := azidentity.NewClientSecretCredential(
		config.Azure.MultiTenantAuth.Tenants["RED"].TenantID,
		config.Azure.MultiTenantAuth.Tenants["RED"].Writer.ClientID,
		config.Azure.MultiTenantAuth.Tenants["RED"].Writer.ClientSecret,
		nil,
	)

	fmt.Println("cred")
	fmt.Println(cred)

	// graphClient, _ := graph.NewGraphServiceClientWithCredentials(
	// 	cred, []string{"User.Read"})

	graphClient, _ := graph.NewGraphServiceClientWithCredentials(
		cred, []string{"https://graph.microsoft.com/.default"})

	// requestFilter := "groupId eq '2a9694fb-117c-43f4-9c41-49d83dc0a82b'"
	requestFilter := "groupId eq 'f24915f7-3c26-4e24-a18d-97a43a1ba103'"
	// requestFilter := "groupId eq '2a9694fb-117c-43f4-9c41-49d83dc0a82b' and principalId eq 'f24915f7-3c26-4e24-a18d-97a43a1ba103'"

	requestParameters := &graphidentitygovernance.PrivilegedAccessGroupEligibilitySchedulesRequestBuilderGetQueryParameters{
		Filter: &requestFilter,
		Select: []string{"accessId", "principalId", "groupId"},
	}
	configuration := &graphidentitygovernance.PrivilegedAccessGroupEligibilitySchedulesRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	// var eligibilitySchedulestst models.PrivilegedAccessGroupEligibilityScheduleCollectionResponseable
	// To initialize your graphClient, see https://learn.microsoft.com/en-us/graph/sdks/create-client?from=snippets&tabs=go
	eligibilitySchedules, err := graphClient.IdentityGovernance().PrivilegedAccess().Group().EligibilitySchedules().Get(context.Background(), configuration)
	lib.CheckFatalError(err)
	// val := eligibilitySchedules.GetAdditionalData()
	val := eligibilitySchedules.GetValue()

	for _, v := range val {
		d := v.GetAdditionalData()
		fmt.Println(d)
		// s := d.GetAdditionalData()
		// fmt.Println(s)
		// ser := d.GetFieldDeserializers()
		// for s, _ := range ser {
		// 	fmt.Println(s)
		// }
		// fmt.Println(d)
		// fmt.Println(ser)

	}
	// mgc-beta identity-governance privileged-access group eligibility-schedules list --filter "groupId eq '2a9694fb-117c-43f4-9c41-49d83dc0a82b' and principalId eq '3cce9d87-3986-4f19-8335-7ed075408ca2'"

	// res, err := azure.HttpGet(urlString, *token)
	// lib.CheckFatalError(err)

	// fmt.Println(string(res))
	fmt.Println(val)
	// jsonStr, _ := json.MarshalIndent(eligibilitySchedules, "", "  ")
	// fmt.Println(string(jsonStr))

	// elapsed := time.Since(startTime)
	// _ = elapsed
	// fmt.Println(elapsed)
}

// startTime := time.Now()
// config := lib.GetCldConfig(nil)
// _ = config
// // tokens, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{})
// // lib.CheckFatalError(err)
// // token, err := azure.GetTenantSPToken(lib.MultiAuthTokenRequestOptions{})
// tokenReq, err := azure.GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{
// 	Scope:         "graph",
// 	GetWriteToken: true,
// }, nil)
// lib.CheckFatalError(err)
// token, err := tokenReq.SelectTenant("RED")
// lib.CheckFatalError(err)
// // _ = tokens
// _ = token
