package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// func main() {
// token, err := GetLogAnalyticsToken()
// if err != nil {
// 	panic(err)
// }

// armGetRequest(*token)

// reqVars := azureAuthRequirements{
// 	AZURE_TENANT_ID:       true,
// 	AZURE_SUBSCRIPTION_ID: true,
// 	AZURE_CLIENT_ID:       true,
// 	AZURE_CLIENT_SECRET:   true,
// 	AZURE_RESOURCE_GROUP:  false,
// 	AZURE_RESOURCE_NAME:   false,
// }
// envVars := GetAzureEnvVariables(reqVars)

// fmt.Println(envVars)
// }

type TokenRequestResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	ExpiresOn    string `json:"expires_on"`
	ExtExpiresIn string `json:"ext_expires_in"`
	NotBefore    string `json:"not_before"`
	Resource     string `json:"resource"`
	TokenType    string `json:"token_type"`
}

type azureAuthDetails struct {
	AZURE_TENANT_ID       string
	AZURE_SUBSCRIPTION_ID string
	AZURE_CLIENT_ID       string
	AZURE_CLIENT_SECRET   string
	AZURE_RESOURCE_GROUP  string
	AZURE_RESOURCE_NAME   string
}

type azureAuthRequirements struct {
	AZURE_TENANT_ID       bool
	AZURE_SUBSCRIPTION_ID bool
	AZURE_CLIENT_ID       bool
	AZURE_CLIENT_SECRET   bool
	AZURE_RESOURCE_GROUP  bool
	AZURE_RESOURCE_NAME   bool
}

type TokenData struct {
	Token     string
	ExpiresOn string
}

func GetLogAnalyticsToken() (*TokenRequestResponse, error) {
	var (
		authDetails         azureAuthDetails
		authRequestResponse *TokenRequestResponse
	)
	// ctx := context.Background()

	authDetails.AZURE_TENANT_ID = os.Getenv("AZURE_TENANT_ID")
	authDetails.AZURE_SUBSCRIPTION_ID = os.Getenv("AZURE_SUBSCRIPTION_ID")
	authDetails.AZURE_CLIENT_ID = os.Getenv("AZURE_CLIENT_ID")
	authDetails.AZURE_CLIENT_SECRET = os.Getenv("AZURE_CLIENT_SECRET")
	urlString := "https://login.microsoftonline.com/" + authDetails.AZURE_TENANT_ID + "/oauth2/token"
	tokenReqStr := "grant_type=client_credentials&client_id=" + authDetails.AZURE_CLIENT_ID + "&resource=https://api.loganalytics.io&client_secret=" + authDetails.AZURE_CLIENT_SECRET

	req, err := http.NewRequest(http.MethodPost, urlString, bytes.NewBufferString(tokenReqStr))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(res.Body)
	if res.StatusCode == 400 {
		log.Fatal("Error fetching LA Workspace Tables: ", string(responseBody))
	}
	// fmt.Println(string(responseBody))
	err = json.Unmarshal(responseBody, &authRequestResponse)
	if err != nil {
		return nil, err
	}
	return authRequestResponse, nil
}

func GetAzureEnvVariables(requiredEnvVars azureAuthRequirements) *azureAuthDetails {
	envs := azureAuthDetails{
		AZURE_TENANT_ID:       os.Getenv("AZURE_TENANT_ID"),
		AZURE_SUBSCRIPTION_ID: os.Getenv("AZURE_SUBSCRIPTION_ID"),
		AZURE_CLIENT_ID:       os.Getenv("AZURE_CLIENT_ID"),
		AZURE_CLIENT_SECRET:   os.Getenv("AZURE_CLIENT_SECRET"),
		AZURE_RESOURCE_GROUP:  os.Getenv("AZURE_RESOURCE_GROUP"),
		AZURE_RESOURCE_NAME:   os.Getenv("AZURE_RESOURCE_NAME"),
	}
	envVarValues := reflect.ValueOf(envs)
	envVarTypes := envVarValues.Type()
	requiredValues := reflect.ValueOf(requiredEnvVars)
	for i := 0; i < envVarValues.NumField(); i++ {
		if envVarValues.Field(i).String() == "" && requiredValues.Field(i).Bool() {
			log.Fatal(envVarTypes.Field(i).Name + " has not been assigned")
		}
	}
	return &envs
}

func GetToken() TokenData {
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

func armGetRequest(token TokenRequestResponse, env azureAuthDetails) {

	// urlString := "https://api.loganalytics.azure.com/v1/subscriptions/" + env.AZURE_SUBSCRIPTION_ID + "/resourceGroups/" + env.AZURE_RESOURCE_GROUP + "/providers/Microsoft.Compute/virtualMachines/" + env.AZURE_RESOURCE_NAME + "/query"
	urlString := "https://api.loganalytics.azure.com/v1/subscriptions/" + env.AZURE_SUBSCRIPTION_ID + "/resourceGroups/" + env.AZURE_RESOURCE_GROUP + "/providers/Microsoft.Compute/virtualMachines/query"
	// urlString := "https://management.azure.com/subscriptions/" + env.AZURE_SUBSCRIPTION_ID + "/resourceGroups/" + env.AZURE_RESOURCE_GROUP + "/providers/Microsoft.Compute/virtualMachines/" + env.AZURE_RESOURCE_NAME + 1/providers/microsoft.insights/logs?api-version=2018-08-01-preview"

	// cred, err := azidentity.NewAzureCLICredential(nil)
	// // cred, err := azidentity.NewDefaultAzureCredential(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(token)

	// ctx := context.Background()
	// tokenRequestOptions := policy.TokenRequestOptions{
	// 	Scopes: []string{
	// 		"https://management.core.windows.net/.default",
	// 	},
	// }

	// token, err := cred.getToken(ctx, tokenRequestOptions)
	// // _ = token
	// if err != nil {
	// 	log.Fatal(err)
	// }

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(token.Token)
	// os.Exit(0)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	// req.Header.Add("Authorization", "Bearer ")

	res, err := http.DefaultClient.Do(req)
	if err != nil {

		log.Fatal("Error fetching LA Workspace Tables")
	}

	// fmt.Println(res)

	responseBody, err := io.ReadAll(res.Body)
	if res.StatusCode == 400 {
		log.Fatal("Error fetching LA Workspace Tables: ", string(responseBody))
	}

	fmt.Println(string(responseBody))
}

// func multipleTenantAuth() {
// 	os.
// }
