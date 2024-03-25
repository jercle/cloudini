package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	"github.com/jercle/azg/lib"
)

type roleAssignmentRequestOptions struct {
	Properties struct {
		Condition                       string `json:"condition"`
		ConditionVersion                string `json:"conditionVersion"`
		LinkedRoleEligibilityScheduleID string `json:"linkedRoleEligibilityScheduleId"`
		PrincipalID                     string `json:"principalId"`
		RequestType                     string `json:"requestType"`
		RoleDefinitionID                string `json:"roleDefinitionId"`
		ScheduleInfo                    struct {
			Expiration struct {
				Duration    string `json:"duration"`
				EndDateTime any    `json:"endDateTime"`
				Type        string `json:"type"`
			} `json:"expiration"`
			StartDateTime time.Time `json:"startDateTime"`
		} `json:"scheduleInfo"`
	} `json:"properties"`
}

type unifiedRoleAssignmentScheduleRequest struct {
	_Odata_Type       string `json:"@odata.type"`
	Action            string `json:"action"`
	AppScopeID        string `json:"appScopeId"`
	ApprovalID        string `json:"approvalId"`
	CompletedDateTime string `json:"completedDateTime"`
	CreatedBy         struct {
		_Odata_Type string `json:"@odata.type"`
	} `json:"createdBy"`
	CreatedDateTime  string `json:"createdDateTime"`
	CustomData       string `json:"customData"`
	DirectoryScopeID string `json:"directoryScopeId"`
	ID               string `json:"id"`
	IsValidationOnly string `json:"isValidationOnly"`
	Justification    string `json:"justification"`
	PrincipalID      string `json:"principalId"`
	RoleDefinitionID string `json:"roleDefinitionId"`
	ScheduleInfo     struct {
		_Odata_Type string `json:"@odata.type"`
	} `json:"scheduleInfo"`
	Status           string `json:"status"`
	TargetScheduleID string `json:"targetScheduleId"`
	TicketInfo       struct {
		_Odata_Type string `json:"@odata.type"`
	} `json:"ticketInfo"`
}

func main() {
	ctx := context.Background()
	client, err := public.New(clientId, public.WithAuthority("https://login.microsoftonline.com/"+tenantId))

	lib.CheckFatalError(err)
	var result public.AuthResult

	scopes := []string{"Directory.ReadAll"}
	// t, _ := client.AuthCodeURL(context.Background(), clientId, "http://localhost", scopes)
	// t, _ := client.AcquireTokenInteractive(context.TODO(), scopes)
	a, _ := client.Accounts(ctx)

	fmt.Println(a)
	os.Exit(0)
	accounts, err := client.Accounts(context.TODO())
	lib.CheckFatalError(err)

	fmt.Println(accounts)

	if len(accounts) > 0 {
		result, err = client.AcquireTokenSilent(context.TODO(), scopes, public.WithSilentAccount(accounts[0]))
		fmt.Println(result)
	}

	// deviceConfig := auth.NewDeviceFlowConfig(applicationId, tenantId)
	// authorizer, err := deviceConfig.Authorizer()
}

// func main() {
// 	ctx := context.Background()
// 	// token, err := azure.GetAllTenantSPTokens(azure.AzureRequestOptions{
// 	// 	TenantId: "e9f4bce2-7308-461a-91ce-3f663d079f47",
// 	// })
// 	// token := azure.GetAzCliToken()
// 	token, err := azidentity.NewAzureCLICredential(&azidentity.AzureCLICredentialOptions{})
// 	lib.CheckFatalError(err)
// 	tokData, err := token.GetToken(ctx, policy.TokenRequestOptions{
// 		Scopes: []string{
// 			"api://4ca69554-1c1c-4eb5-84d6-a7e61551929a/Directory.Read",
// 			// "https://management.core.windows.net//.default",
// 		},
// 	})
// 	lib.CheckFatalError(err)
// 	fmt.Println(tokData)
// 	// token := azure.GetToken("e9f4bce2-7308-461a-91ce-3f663d079f47")
// 	// os.Exit(0)
// 	urlString := "https://graph.microsoft.com/v1.0/roleManagement/directory/roleAssignmentScheduleRequests/filterByCurrentUser(on='principal')"
// 	// lib.CheckFatalError(err)
// 	// roleAssignmentScheduleRequestName := uuid.New()
// 	// urlString := "https://management.azure.com/" +
// 	// 	// scope +
// 	// 	"/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/" +
// 	// 	roleAssignmentScheduleRequestName.String() +
// 	// 	"?api-version=2020-10-01"
// 	// fmt.Println(token)
// 	// // os.Exit(0)
// 	req, err := http.NewRequest(http.MethodGet, urlString, nil)
// 	lib.CheckFatalError(err)
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Authorization", "Bearer "+tokData.Token)

// 	// fmt.Fprintln(os.Stdout, []any{token}...)
// 	// activatePIMRequestOptions := unifiedRoleAssignmentScheduleRequest{

// 	// }

// 	res, err := http.DefaultClient.Do(req)
// 	lib.CheckFatalError(err)
// 	// if err != nil {
// 	// 	// log.Fatal("Error fetching list of Subscriptions")
// 	// 	lib.CheckFatalError(err)
// 	// }

// 	responseBody, err := io.ReadAll(res.Body)
// 	lib.CheckFatalError(err)

// 	fmt.Println(string(responseBody))
// 	// if res.StatusCode == 400 {
// 	// 	// log.Fatal("Error fetching list of Subscriptions: ", string(responseBody))
// 	// 	lib.CheckFatalError(err)
// 	// }
// 	// if err != nil {
// 	// 	// log.Fatal(err)
// 	// 	lib.CheckFatalError(err)
// 	// }
// 	// defer res.Body.Close()

// 	// fmt.Println(string(responseBody))
// 	// var subsList azure.SubsReqResBody
// 	// json.Unmarshal(responseBody, &subsList)
// 	// subsList.UpdateTenantName(token.TenantName)
// 	// lib.MarshalAndPrintJson(subsList.Value)
// }
