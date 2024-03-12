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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type TokenRequestResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	ExpiresOn    string `json:"expires_on"`
	ExtExpiresIn string `json:"ext_expires_in"`
	NotBefore    string `json:"not_before"`
	Resource     string `json:"resource"`
	TokenType    string `json:"token_type"`
}

func main() {
	getLogAnalyticsToken()
}

type azureAuthDetails struct {
	tenantId       string
	subscriptionId string
	clientId       string
	clientSecret   string
}

func getLogAnalyticsToken() (TokenRequestResponse, error) {
	var (
		authDetails         azureAuthDetails
		authRequestResponse TokenRequestResponse
	)
	// ctx := context.Background()

	authDetails.tenantId = os.Getenv("ARM_TENANT_ID")
	authDetails.subscriptionId = os.Getenv("ARM_SUBSCRIPTION_ID")
	authDetails.clientId = os.Getenv("ARM_CLIENT_ID")
	authDetails.clientSecret = os.Getenv("ARM_CLIENT_SECRET")
	urlString := "https://login.microsoftonline.com/" + authDetails.tenantId + "/oauth2/token"
	tokenReqStr := "grant_type=client_credentials&client_id=" + authDetails.clientId + "&resource=https://api.loganalytics.io&client_secret=" + authDetails.clientSecret

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

func armGetRequest() {
	urlString := "https://api.loganalytics.azure.com/v1/subscriptions/fd7b1915-0f99-4d2d-8154-6dd848a03f4a/resourceGroups/rg-apcdesktop-w10ms/providers/Microsoft.Compute/virtualMachines/vm-apdt-w10m601/query"
	// urlString := "https://management.azure.com/subscriptions/fd7b1915-0f99-4d2d-8154-6dd848a03f4a/resourceGroups/rg-apcdesktop-w10ms/providers/Microsoft.Compute/virtualMachines/vm-apdt-w10m601/providers/microsoft.insights/logs?api-version=2018-08-01-preview"

	cred, err := azidentity.NewAzureCLICredential(nil)
	// cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(token)

	ctx := context.Background()
	tokenRequestOptions := policy.TokenRequestOptions{
		Scopes: []string{
			"https://management.core.windows.net/.default",
		},
	}

	token, err := cred.GetToken(ctx, tokenRequestOptions)
	// _ = token
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(token.Token)
	// os.Exit(0)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.Token)
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
