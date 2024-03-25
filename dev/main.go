// Get Azure ACRs

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jercle/azg/cmd/azure"
	"github.com/jercle/azg/lib"
)

func main() {
	listSubscriptionContainerRegistries()
}

func listSubscriptionContainerRegistries() {
	// Get all Azure Container Registries for subscription
	// fmt.Println("ACRs")
	tenantId := os.Getenv("AZURE_TENANT_ID")
	subscriptionId := os.Getenv("AZURE_SUBSCRIPTION_ID")
	clientId := os.Getenv("AZURE_CLIENT_ID")
	clientSecret := os.Getenv("AZURE_CLIENT_SECRET")
	token, err := azure.GetServicePrincipalToken(tenantId, lib.CldConfigClientAuthDetails{ClientID: clientId, ClientSecret: clientSecret})
	lib.CheckFatalError(err)

	// cred, err := azidentity.NewClientSecretCredential(tenantId, clientId, clientSecret, nil)
	// lib.CheckFatalError(err)

	// clientFactory, err := armcontainerregistry.NewClientFactory(subscriptionId, cred, nil)

	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/providers/Microsoft.ContainerRegistry/registries?api-version=2023-01-01-preview"

	// client := clientFactory.NewRegistriesClient()

	// pager := client.NewListPager(nil)

	// pager.

	// fmt.Println(token)

	// ctx := context.Background()
	// tokenRequestOptions := policy.TokenRequestOptions{
	// 	Scopes: []string{
	// 		"https://management.core.windows.net/.default",
	// 	},
	// }

	// token, err := cred.GetToken(ctx, tokenRequestOptions)
	// _ = token
	// if err != nil {
	// 	log.Fatal(err)
	// }

	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		log.Fatal(err)
	}

	// POST https://management.azure.com/providers/Microsoft.Billing/billingAccounts/4f53d9ee-159b-562a-dbdf-0569086af295%3A59377d87-4d23-4a91-8051-d6819401ef72_2019-05-31/billingProfiles/ZBRW-RA37-BG7-PGB/providers/Microsoft.CostManagement/pricesheets/default/download?api-version=2023-11-01
	// Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IlhSdmtvOFA3QTNVYVdTblU3Yk05blQwTWpoQSIsImtpZCI6IlhSdmtvOFA3QTNVYVdTblU3Yk05blQwTWpoQSJ9.eyJhdWQiOiJodHRwczovL21hbmFnZW1lbnQuY29yZS53aW5kb3dzLm5ldCIsImlzcyI6Imh0dHBzOi8vc3RzLndpbmRvd3MubmV0LzY0OGE1ZWQ3LWM1YWMtNDViNy1iNDA2LTc5MWYwZDQzMzRhNi8iLCJpYXQiOjE3MTAyMDU0MzUsIm5iZiI6MTcxMDIwNTQzNSwiZXhwIjoxNzEwMjEwMTA0LCJhY3IiOiIxIiwiYWlvIjoiQVZRQXEvOFdBQUFBdkJQbTU4eVVOMzBvUjRwUWdpN3pEMlZyam9VT0xYVEJuS0tVTDA4SllyY25tejg1L0tYb3IvOEZtVTVuaXM3d3YrNzQrbXROaEpaUERaUjI5TFhrK2xVWWhFRmhlMzN4SVE3bjFBT0Q3Q2c9IiwiYW1yIjpbInB3ZCIsIm1mYSJdLCJhcHBpZCI6IjE4ZmJjYTE2LTIyMjQtNDVmNi04NWIwLWY3YmYyYjM5YjNmMyIsImFwcGlkYWNyIjoiMCIsImZhbWlseV9uYW1lIjoiYWRtLjExNDQwZSIsImdpdmVuX25hbWUiOiJFIiwiZ3JvdXBzIjpbImNhMTgxYzg5LTEwYjUtNDA4MC1iZTg0LTYzNGRiNmJlOGZmNCIsIjQyMjk4N2E4LTMwNGEtNGY5ZC04ODM1LWZkNzU0YzQ5OWQ5YyIsImFhYWY0MzBiLTRlNTgtNDlhMC1iNjI5LTRiYzIxY2EzYjY1MCIsImQ3Y2ViMjBiLTc4YWUtNDQ0OC04MGFkLWMxMjc2ZGVlZWJlNyIsIjhmZTk0ODBmLWY3ZTItNDhhZi1iZDcxLWQ2Y2UyYmVkOTllNiIsIjFlYTgxMzEwLTRiZTEtNGMwYy1hN2MwLTY5NDZmOWIzNjU0ZSIsImY2MmQ3ZTEyLTAzZWMtNDUxYi05MjU3LWU5OWExNWRkMWZjMiIsImM0NTI2OTE2LTIyYjItNDE4OS05NTM1LTUxMjZiN2VjMGQ2YiIsIjI3MGVhYzM2LTZiYjYtNGQ5Yi1iNDBiLTJjMTY1NWFmM2YwMCIsIjRjYmIyZTM4LTg3ZWUtNDU4NS04N2NhLTIxNzYwZTNkMzYxZCIsIjgwMzcxMjQzLTQxMDktNDU2ZC05N2QzLTA0ZTliZGY1MWY5OCIsIjU2OWMyYTQ2LWM5YmItNDc3ZC1hODhkLWFmMTgyOWI3NmIxNyIsImJkYjRhYzQ5LWQwY2QtNDdjNy1iMWQxLWJkZmExN2VlZjdjYyIsIjdiZmY5MjU3LThhMWYtNDA5Ny05NzkwLTBmYzRkYTY5MGJhYyIsIjUzZGIwZTU5LTU3OWMtNDFmYy1hNmYzLTU1YTA4ZjBmMmJlNSIsImQ4ODU5NzViLWU3NTEtNDYzYy04YzAxLWFiYWZkMmI4OTAwYyIsIjdkMDMwNDYzLWUwZGUtNDg5Ny1hZjUyLWVmODQ4MWFhNjI3YiIsIjQ3NTJlZTc0LTA4MjEtNGEwZC1hZWNiLTczZjFmOWE1MmEwYiIsIjY3YjVmYTc0LTlhMmItNGU2ZS04NjZlLWIyNDZhN2QxNjRhMCIsIjUxMjQyNjg0LTU0MjYtNGZhZi05NzFhLTE1ZjVjYThiZWRmZSIsImVmY2Y0MTg3LTE4YzctNGE3OC1hMjM3LTVhYWJkOGIzNDYwZiIsIjBlZTVhYzhiLWYwNDktNDYxZi05ZjMwLTBiM2U0YmIyNWIyMCIsImVlMDA1ODkzLWE3ODMtNGNjMC1iNTgxLWYxZTJlNWEyNzM0MyIsImMzNzBkNjliLTA2MTItNGI3My1hOWRlLTNlMTg0NGQ1MjM5NyIsImE0N2VmNTljLWY3YzQtNDhkMi05NmIyLWZmOGJkNGI5NmJkZiIsIjRjMTMxYTllLWRkZTYtNGMwMC1hYmE0LTEwYWE0ZTU3ZGRmZCIsImVhMWE4MDllLTc3ZDgtNGJlYS05NGVlLTI5YjYwYzk3MjhmNiIsIjE4NjkwMDlmLWZjZDctNGU5Ny05YmM1LTliNTUxMjU0MTExMSIsImRlYmYxZGEwLWFkOTUtNDFlYi1iYTQ3LTQ0ODBiM2Y0ODBkYSIsImUwMGYxNmFkLTUwMjYtNDY1OC05Y2E5LTBmOTI0NGNlNmI1MCIsImY0ZjdkZGIyLTBhYTItNDBkNi05Mzg0LTc4MmRjYWE0ZjFlMiIsIjBiZjQ2MmI4LTI0NDgtNGIwMS1iZTRjLTIzOTI1M2UxNDEyMSIsIjA3ZDlmY2NjLWMzMTYtNGY3Mi04NTU3LTE5NzExYTkxZTYzZCIsIjg3ZDdhM2Q2LTFkZmItNDBiNC04NTM3LTZhMDc2MjZjMzc5MSIsIjlhNDVkYmRkLTE3ZWItNDgyMS1iNmMxLWMwOTI1YzM1MmM3YSIsIjI1NzdkMGVlLTU4ZTktNGNlNi04NzkxLWMwMWZiMWJlMWQ0YSIsIjJhOTY5NGZiLTExN2MtNDNmNC05YzQxLTQ5ZDgzZGMwYTgyYiIsImFkMmUwMWZlLTg5MGQtNDNhZC1hMWY3LTQxZWVhMmQ3NTVlYyIsImZmNjRiMGZlLWRlY2QtNDJhYy04OWIxLTc1NTA5ODc1NTEyNSJdLCJpZHR5cCI6InVzZXIiLCJpcGFkZHIiOiIxNDQuMTQwLjE1MC41IiwibmFtZSI6ImFkbS4xMTQ0MGUiLCJvaWQiOiJjMWYwZTNiNy00OTVlLTQ2NjEtOTAyNS1hZjQxNDUzYzNjODUiLCJvbnByZW1fc2lkIjoiUy0xLTUtMjEtMTA5Mzk4NjEwMC0xNzU3NDc3MjUyLTI4NDc3NTk5NzUtODYwNyIsInB1aWQiOiIxMDAzMjAwMzNCOTYxNjNBIiwicmgiOiIwLkFVSUExMTZLWkt6RnQwVzBCbmtmRFVNMHBrWklmM2tBdXRkUHVrUGF3ZmoyTUJOQ0FMUS4iLCJzY3AiOiJ1c2VyX2ltcGVyc29uYXRpb24iLCJzdWIiOiJ0N1Vvb2VNZVM4ZS16Ynl3ZlgxZTZ1UGVGT3I0Vl83VUtnQ3pqVkZTamdNIiwidGlkIjoiNjQ4YTVlZDctYzVhYy00NWI3LWI0MDYtNzkxZjBkNDMzNGE2IiwidW5pcXVlX25hbWUiOiJhZG0uMTE0NDBlQGFzaW8uZ292LmF1IiwidXBuIjoiYWRtLjExNDQwZUBhc2lvLmdvdi5hdSIsInV0aSI6Il9jLWdQTGVZQkVDbjNaYTJJbXNGQUEiLCJ2ZXIiOiIxLjAiLCJ3aWRzIjpbImIwZjU0NjYxLTJkNzQtNGM1MC1hZmEzLTFlYzgwM2YxMmVmZSIsIjYyZTkwMzk0LTY5ZjUtNDIzNy05MTkwLTAxMjE3NzE0NWUxMCIsImI3OWZiZjRkLTNlZjktNDY4OS04MTQzLTc2YjE5NGU4NTUwOSJdLCJ4bXNfY2FlIjoiMSIsInhtc190Y2R0IjoxNjA2ODYxNDUwfQ.UDjV7WWeRELnw6EDOtbiUStSDexz3DjAuJNFEx_gQpl9daUFwNuE8J4sf40FHHTlh_OvCEMP6llDmFFElP0i9hDZZdvXrDRTnialdfMYk6KRTzeflES8wVguSk3Tf_J80Njt8ymkmJUPf-7212rdfGPVywG7qGiYl-VSIUu3wB1lLpAjF9HQ0Y1yyjJTUVrgOcYiX4UsJRMJFca76m437KRao2qb2DsdJuQ_oCk1MhAXNcI20wzpwitEAQz6aOlmGFLQLn0xd8dcD4sh4XVsJeZgBxrGemij9otuZ-8TvQ8ytqGWK52fOYe52WNlhnqCVys0jhcSzNfLvJ5tkmBk_Q
	// Content-type: application/json

	// https://management.azure.com/subscriptions/9564369d-5e7c-42c8-bea5-328ea671bbc6/providers/Microsoft.Web/billingMeters?api-version=2022-03-01
	// POST https://management.azure.com/providers/Microsoft.Billing/billingAccounts/D 6462/billingProfiles/D 6462/providers/Microsoft.CostManagement/pricesheets/default/download?api-version=2023-11-01
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.Token)

	res, err := http.DefaultClient.Do(req)
	lib.CheckFatalError(err)

	// fmt.Println(res)

	responseBody, err := io.ReadAll(res.Body)
	lib.CheckFatalError(err)

	fmt.Println(string(responseBody))

}
