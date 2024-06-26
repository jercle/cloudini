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
	// req.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IlhSdmtvOFA3QTNVYVdTblU3Yk05blQwTWpoQSIsImtpZCI6IlhSdmtvOFA3QTNVYVdTblU3Yk05blQwTWpoQSJ9.eyJhdWQiOiJodHRwczovL21hbmFnZW1lbnQuY29yZS53aW5kb3dzLm5ldCIsImlzcyI6Imh0dHBzOi8vc3RzLndpbmRvd3MubmV0LzY0OGE1ZWQ3LWM1YWMtNDViNy1iNDA2LTc5MWYwZDQzMzRhNi8iLCJpYXQiOjE3MTAyMDU0MzUsIm5iZiI6MTcxMDIwNTQzNSwiZXhwIjoxNzEwMjEwMTA0LCJhY3IiOiIxIiwiYWlvIjoiQVZRQXEvOFdBQUFBdkJQbTU4eVVOMzBvUjRwUWdpN3pEMlZyam9VT0xYVEJuS0tVTDA4SllyY25tejg1L0tYb3IvOEZtVTVuaXM3d3YrNzQrbXROaEpaUERaUjI5TFhrK2xVWWhFRmhlMzN4SVE3bjFBT0Q3Q2c9IiwiYW1yIjpbInB3ZCIsIm1mYSJdLCJhcHBpZCI6IjE4ZmJjYTE2LTIyMjQtNDVmNi04NWIwLWY3YmYyYjM5YjNmMyIsImFwcGlkYWNyIjoiMCIsImZhbWlseV9uYW1lIjoiYWRtLjExNDQwZSIsImdpdmVuX25hbWUiOiJFIiwiZ3JvdXBzIjpbImNhMTgxYzg5LTEwYjUtNDA4MC1iZTg0LTYzNGRiNmJlOGZmNCIsIjQyMjk4N2E4LTMwNGEtNGY5ZC04ODM1LWZkNzU0YzQ5OWQ5YyIsImFhYWY0MzBiLTRlNTgtNDlhMC1iNjI5LTRiYzIxY2EzYjY1MCIsImQ3Y2ViMjBiLTc4YWUtNDQ0OC04MGFkLWMxMjc2ZGVlZWJlNyIsIjhmZTk0ODBmLWY3ZTItNDhhZi1iZDcxLWQ2Y2UyYmVkOTllNiIsIjFlYTgxMzEwLTRiZTEtNGMwYy1hN2MwLTY5NDZmOWIzNjU0ZSIsImY2MmQ3ZTEyLTAzZWMtNDUxYi05MjU3LWU5OWExNWRkMWZjMiIsImM0NTI2OTE2LTIyYjItNDE4OS05NTM1LTUxMjZiN2VjMGQ2YiIsIjI3MGVhYzM2LTZiYjYtNGQ5Yi1iNDBiLTJjMTY1NWFmM2YwMCIsIjRjYmIyZTM4LTg3ZWUtNDU4NS04N2NhLTIxNzYwZTNkMzYxZCIsIjgwMzcxMjQzLTQxMDktNDU2ZC05N2QzLTA0ZTliZGY1MWY5OCIsIjU2OWMyYTQ2LWM5YmItNDc3ZC1hODhkLWFmMTgyOWI3NmIxNyIsImJkYjRhYzQ5LWQwY2QtNDdjNy1iMWQxLWJkZmExN2VlZjdjYyIsIjdiZmY5MjU3LThhMWYtNDA5Ny05NzkwLTBmYzRkYTY5MGJhYyIsIjUzZGIwZTU5LTU3OWMtNDFmYy1hNmYzLTU1YTA4ZjBmMmJlNSIsImQ4ODU5NzViLWU3NTEtNDYzYy04YzAxLWFiYWZkMmI4OTAwYyIsIjdkMDMwNDYzLWUwZGUtNDg5Ny1hZjUyLWVmODQ4MWFhNjI3YiIsIjQ3NTJlZTc0LTA4MjEtNGEwZC1hZWNiLTczZjFmOWE1MmEwYiIsIjY3YjVmYTc0LTlhMmItNGU2ZS04NjZlLWIyNDZhN2QxNjRhMCIsIjUxMjQyNjg0LTU0MjYtNGZhZi05NzFhLTE1ZjVjYThiZWRmZSIsImVmY2Y0MTg3LTE4YzctNGE3OC1hMjM3LTVhYWJkOGIzNDYwZiIsIjBlZTVhYzhiLWYwNDktNDYxZi05ZjMwLTBiM2U0YmIyNWIyMCIsImVlMDA1ODkzLWE3ODMtNGNjMC1iNTgxLWYxZTJlNWEyNzM0MyIsImMzNzBkNjliLTA2MTItNGI3My1hOWRlLTNlMTg0NGQ1MjM5NyIsImE0N2VmNTljLWY3YzQtNDhkMi05NmIyLWZmOGJkNGI5NmJkZiIsIjRjMTMxYTllLWRkZTYtNGMwMC1hYmE0LTEwYWE0ZTU3ZGRmZCIsImVhMWE4MDllLTc3ZDgtNGJlYS05NGVlLTI5YjYwYzk3MjhmNiIsIjE4NjkwMDlmLWZjZDctNGU5Ny05YmM1LTliNTUxMjU0MTExMSIsImRlYmYxZGEwLWFkOTUtNDFlYi1iYTQ3LTQ0ODBiM2Y0ODBkYSIsImUwMGYxNmFkLTUwMjYtNDY1OC05Y2E5LTBmOTI0NGNlNmI1MCIsImY0ZjdkZGIyLTBhYTItNDBkNi05Mzg0LTc4MmRjYWE0ZjFlMiIsIjBiZjQ2MmI4LTI0NDgtNGIwMS1iZTRjLTIzOTI1M2UxNDEyMSIsIjA3ZDlmY2NjLWMzMTYtNGY3Mi04NTU3LTE5NzExYTkxZTYzZCIsIjg3ZDdhM2Q2LTFkZmItNDBiNC04NTM3LTZhMDc2MjZjMzc5MSIsIjlhNDVkYmRkLTE3ZWItNDgyMS1iNmMxLWMwOTI1YzM1MmM3YSIsIjI1NzdkMGVlLTU4ZTktNGNlNi04NzkxLWMwMWZiMWJlMWQ0YSIsIjJhOTY5NGZiLTExN2MtNDNmNC05YzQxLTQ5ZDgzZGMwYTgyYiIsImFkMmUwMWZlLTg5MGQtNDNhZC1hMWY3LTQxZWVhMmQ3NTVlYyIsImZmNjRiMGZlLWRlY2QtNDJhYy04OWIxLTc1NTA5ODc1NTEyNSJdLCJpZHR5cCI6InVzZXIiLCJpcGFkZHIiOiIxNDQuMTQwLjE1MC41IiwibmFtZSI6ImFkbS4xMTQ0MGUiLCJvaWQiOiJjMWYwZTNiNy00OTVlLTQ2NjEtOTAyNS1hZjQxNDUzYzNjODUiLCJvbnByZW1fc2lkIjoiUy0xLTUtMjEtMTA5Mzk4NjEwMC0xNzU3NDc3MjUyLTI4NDc3NTk5NzUtODYwNyIsInB1aWQiOiIxMDAzMjAwMzNCOTYxNjNBIiwicmgiOiIwLkFVSUExMTZLWkt6RnQwVzBCbmtmRFVNMHBrWklmM2tBdXRkUHVrUGF3ZmoyTUJOQ0FMUS4iLCJzY3AiOiJ1c2VyX2ltcGVyc29uYXRpb24iLCJzdWIiOiJ0N1Vvb2VNZVM4ZS16Ynl3ZlgxZTZ1UGVGT3I0Vl83VUtnQ3pqVkZTamdNIiwidGlkIjoiNjQ4YTVlZDctYzVhYy00NWI3LWI0MDYtNzkxZjBkNDMzNGE2IiwidW5pcXVlX25hbWUiOiJhZG0uMTE0NDBlQGFzaW8uZ292LmF1IiwidXBuIjoiYWRtLjExNDQwZUBhc2lvLmdvdi5hdSIsInV0aSI6Il9jLWdQTGVZQkVDbjNaYTJJbXNGQUEiLCJ2ZXIiOiIxLjAiLCJ3aWRzIjpbImIwZjU0NjYxLTJkNzQtNGM1MC1hZmEzLTFlYzgwM2YxMmVmZSIsIjYyZTkwMzk0LTY5ZjUtNDIzNy05MTkwLTAxMjE3NzE0NWUxMCIsImI3OWZiZjRkLTNlZjktNDY4OS04MTQzLTc2YjE5NGU4NTUwOSJdLCJ4bXNfY2FlIjoiMSIsInhtc190Y2R0IjoxNjA2ODYxNDUwfQ.UDjV7WWeRELnw6EDOtbiUStSDexz3DjAuJNFEx_gQpl9daUFwNuE8J4sf40FHHTlh_OvCEMP6llDmFFElP0i9hDZZdvXrDRTnialdfMYk6KRTzeflES8wVguSk3Tf_J80Njt8ymkmJUPf-7212rdfGPVywG7qGiYl-VSIUu3wB1lLpAjF9HQ0Y1yyjJTUVrgOcYiX4UsJRMJFca76m437KRao2qb2DsdJuQ_oCk1MhAXNcI20wzpwitEAQz6aOlmGFLQLn0xd8dcD4sh4XVsJeZgBxrGemij9otuZ-8TvQ8ytqGWK52fOYe52WNlhnqCVys0jhcSzNfLvJ5tkmBk_Q")

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
