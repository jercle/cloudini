package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jercle/azg/cmd/azure"
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

func main() {
	token, err := azure.GetSingleTenantSPToken(azure.AzureRequestOptions{})
	lib.CheckFatalError(err)
	roleAssignmentScheduleRequestName := uuid.New()
	urlString := "https://management.azure.com/" +
		// scope +
		"/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/" +
		roleAssignmentScheduleRequestName.String() +
		"?api-version=2020-10-01"
	fmt.Println(token)
	os.Exit(0)
	req, err := http.NewRequest(http.MethodPut, urlString, nil)
	lib.CheckFatalError(err)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.TokenData.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// log.Fatal("Error fetching list of Subscriptions")
		lib.CheckFatalError(err)
	}

	responseBody, err := io.ReadAll(res.Body)
	if res.StatusCode == 400 {
		// log.Fatal("Error fetching list of Subscriptions: ", string(responseBody))
		lib.CheckFatalError(err)
	}
	if err != nil {
		// log.Fatal(err)
		lib.CheckFatalError(err)
	}
	defer res.Body.Close()

	fmt.Println(string(responseBody))
	// var subsList azure.SubsReqResBody
	// json.Unmarshal(responseBody, &subsList)
	// subsList.UpdateTenantName(token.TenantName)
	// lib.MarshalAndPrintJson(subsList.Value)
}
