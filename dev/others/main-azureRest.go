package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
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

	// token := azure.GetAzCliToken()
	cred, err := azidentity.NewDefaultAzureCredential(nil)
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
	_ = token
	if err != nil {
		log.Fatal(err)
	}
	// urlString := "https://management.azure.com/providers?api-version=2021-04-01"
	// urlString := "https://management.azure.com/subscriptions/bae338c7-6098-4d52-b173-e2147e107dfa/providers/Microsoft.HDInsight/locations/australiaeast/billingSpecs?api-version=2021-06-01"
	// urlString := "https://management.azure.com/subscriptions/bae338c7-6098-4d52-b173-e2147e107dfa/providers/Microsoft.Consumption/pricesheets/default?api-version=2019-10-01"
	// urlString := "https://prices.azure.com/api/retail/prices?currencyCode='AUD'&armRegionName='australiaeast'"
	// urlString := "https://management.azure.com/subscriptions/9564369d-5e7c-42c8-bea5-328ea671bbc6/providers/Microsoft.Web/billingMeters?api-version=2022-03-01"
	// urlString := "https://management.azure.com/providers/Microsoft.Billing/billingAccounts/4f53d9ee-159b-562a-dbdf-0569086af295%3A59377d87-4d23-4a91-8051-d6819401ef72_2019-05-31/billingProfiles/ZBRW-RA37-BG7-PGB/providers/Microsoft.CostManagement/pricesheets/default/download?api-version=2023-11-01"
	urlString := "https://management.azure.com/providers/Microsoft.Billing/billingAccounts/4f53d9ee-159b-562a-dbdf-0569086af295:59377d87-4d23-4a91-8051-d6819401ef72_2019-05-31/billingProfiles/ZBRW-RA37-BG7-PGB/providers/Microsoft.CostManagement/operationResults/983f886d-3620-469e-a523-bd24e9106a18?sessiontoken=0%3a2521190&api-version=2023-11-01&OperationType=PriceSheet&t=638458028446294165&c=MIIHADCCBeigAwIBAgITHgPrWOVrMb7qufvgEAAAA-tY5TANBgkqhkiG9w0BAQsFADBEMRMwEQYKCZImiZPyLGQBGRYDR0JMMRMwEQYKCZImiZPyLGQBGRYDQU1FMRgwFgYDVQQDEw9BTUUgSW5mcmEgQ0EgMDYwHhcNMjQwMjAxMDQxMjQzWhcNMjUwMTI2MDQxMjQzWjBAMT4wPAYDVQQDEzVhc3luY29wZXJhdGlvbnNpZ25pbmdjZXJ0aWZpY2F0ZS5tYW5hZ2VtZW50LmF6dXJlLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAOil3F8MsZdl8FeicToFLcoRyDn0Zv76EQTwG1IZtUh-z6uwzGIgy23k7GrXNU1gKVyGJp8lO3encPC02rUQkI0lvN-NUJCoJAEnGZPYOLmA9NylSyr1Ik_Qaz1_UueFRAiyVJlo0Lz27ayfzTTSUd82wyh18q-LWdG49N7fSD_fM1rsfxbY7-Eo4Z5CjxDW3OWmAYKpS0tm17o2hEKrmjeNZJQsSqQxUL-1Be4vND7XzGhGI595ogShOZHOzCBueWR2-8fa5VrwlHqtU1AgvjFk3lYmZejl898JrGFMYH-QSC1iWyRweQ_m3289K-aPeRSWqRihXIG9oHEqouTO1xkCAwEAAaOCA-0wggPpMCcGCSsGAQQBgjcVCgQaMBgwCgYIKwYBBQUHAwEwCgYIKwYBBQUHAwIwPQYJKwYBBAGCNxUHBDAwLgYmKwYBBAGCNxUIhpDjDYTVtHiE8Ys-hZvdFs6dEoFggvX2K4Py0SACAWQCAQowggHLBggrBgEFBQcBAQSCAb0wggG5MGMGCCsGAQUFBzAChldodHRwOi8vY3JsLm1pY3Jvc29mdC5jb20vcGtpaW5mcmEvQ2VydHMvQkwyUEtJSU5UQ0EwMi5BTUUuR0JMX0FNRSUyMEluZnJhJTIwQ0ElMjAwNi5jcnQwUwYIKwYBBQUHMAKGR2h0dHA6Ly9jcmwxLmFtZS5nYmwvYWlhL0JMMlBLSUlOVENBMDIuQU1FLkdCTF9BTUUlMjBJbmZyYSUyMENBJTIwMDYuY3J0MFMGCCsGAQUFBzAChkdodHRwOi8vY3JsMi5hbWUuZ2JsL2FpYS9CTDJQS0lJTlRDQTAyLkFNRS5HQkxfQU1FJTIwSW5mcmElMjBDQSUyMDA2LmNydDBTBggrBgEFBQcwAoZHaHR0cDovL2NybDMuYW1lLmdibC9haWEvQkwyUEtJSU5UQ0EwMi5BTUUuR0JMX0FNRSUyMEluZnJhJTIwQ0ElMjAwNi5jcnQwUwYIKwYBBQUHMAKGR2h0dHA6Ly9jcmw0LmFtZS5nYmwvYWlhL0JMMlBLSUlOVENBMDIuQU1FLkdCTF9BTUUlMjBJbmZyYSUyMENBJTIwMDYuY3J0MB0GA1UdDgQWBBTew4tn2tBCof44yM20soU6sfep6zAOBgNVHQ8BAf8EBAMCBaAwggEmBgNVHR8EggEdMIIBGTCCARWgggERoIIBDYY_aHR0cDovL2NybC5taWNyb3NvZnQuY29tL3BraWluZnJhL0NSTC9BTUUlMjBJbmZyYSUyMENBJTIwMDYuY3JshjFodHRwOi8vY3JsMS5hbWUuZ2JsL2NybC9BTUUlMjBJbmZyYSUyMENBJTIwMDYuY3JshjFodHRwOi8vY3JsMi5hbWUuZ2JsL2NybC9BTUUlMjBJbmZyYSUyMENBJTIwMDYuY3JshjFodHRwOi8vY3JsMy5hbWUuZ2JsL2NybC9BTUUlMjBJbmZyYSUyMENBJTIwMDYuY3JshjFodHRwOi8vY3JsNC5hbWUuZ2JsL2NybC9BTUUlMjBJbmZyYSUyMENBJTIwMDYuY3JsMBcGA1UdIAQQMA4wDAYKKwYBBAGCN3sBATAfBgNVHSMEGDAWgBTxRmjG8cPwKy19i2rhsvm-NfzRQTAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwDQYJKoZIhvcNAQELBQADggEBAGBANOAsl8fyImo8IXVQ_ybox-BzCeFiAddg9Ojd61exGTea02-HjvZzuL6GZ2k21c7rEjKSBHqk50RLm2yXrMcUGc1UO7MRchaKQm1WR7osgDiXAU0rDc9-H-i8I245Y_PbyKIBewiK7UQiCjda4b__7Amrw7q4oILOsxshnN_D1NBm2nSjm-dFiDeJ9jpG9X1jDJSBtOFUQ7Ala8ShJvWPMpbCR3gFHOGS9f13ebv6qXblaf7sHxE6T0OpuEZPL10Iu5yh42IOLo8wvp1euqb61U3JWoOmBXEB3mJ0owRfBHEjuAARDUzNVgq28F6FYqgACE3nKhqDKYDEQstlJK0&s=MJYF6Pp_rfHqCDudBy84mSErpJE_aN7YVawF4Hl5sHcAVCm5RxocGgApwciwylbWh-YKvxhsC5E6L23LuoV73G85D5Zij6ltmN0qtUqb4nmPMhP9y1ddHggbFD_IBBMR_BTYPl3A_-P6VTtMn7PL9VavDc3rV8pRo4ny3jy8OPECuvrjmEbmgIT4NoZ-smpLxKXwFWrfdfCligSq5SAPVOnden0ZA0RrRiYJ06X8_3yFJHuyeTbW_T2bUE3VV9r4sGYhyR3JI7TJ48k1asu6IDAacHuyPCmswlof4LoRwUwZx-XMPk1zXan1ewgdw3_QfpbsMs0WZhdK8UWNN5Ikug&h=xIdMrWVIETaBv2fYdd1QLAsKRN1uV9nVnD2vkgRyCWQ"
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
	// req.Header.Add("Authorization", "Bearer "+token.Token)
	req.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IlhSdmtvOFA3QTNVYVdTblU3Yk05blQwTWpoQSIsImtpZCI6IlhSdmtvOFA3QTNVYVdTblU3Yk05blQwTWpoQSJ9.eyJhdWQiOiJodHRwczovL21hbmFnZW1lbnQuY29yZS53aW5kb3dzLm5ldCIsImlzcyI6Imh0dHBzOi8vc3RzLndpbmRvd3MubmV0LzY0OGE1ZWQ3LWM1YWMtNDViNy1iNDA2LTc5MWYwZDQzMzRhNi8iLCJpYXQiOjE3MTAyMDU0MzUsIm5iZiI6MTcxMDIwNTQzNSwiZXhwIjoxNzEwMjEwMTA0LCJhY3IiOiIxIiwiYWlvIjoiQVZRQXEvOFdBQUFBdkJQbTU4eVVOMzBvUjRwUWdpN3pEMlZyam9VT0xYVEJuS0tVTDA4SllyY25tejg1L0tYb3IvOEZtVTVuaXM3d3YrNzQrbXROaEpaUERaUjI5TFhrK2xVWWhFRmhlMzN4SVE3bjFBT0Q3Q2c9IiwiYW1yIjpbInB3ZCIsIm1mYSJdLCJhcHBpZCI6IjE4ZmJjYTE2LTIyMjQtNDVmNi04NWIwLWY3YmYyYjM5YjNmMyIsImFwcGlkYWNyIjoiMCIsImZhbWlseV9uYW1lIjoiYWRtLjExNDQwZSIsImdpdmVuX25hbWUiOiJFIiwiZ3JvdXBzIjpbImNhMTgxYzg5LTEwYjUtNDA4MC1iZTg0LTYzNGRiNmJlOGZmNCIsIjQyMjk4N2E4LTMwNGEtNGY5ZC04ODM1LWZkNzU0YzQ5OWQ5YyIsImFhYWY0MzBiLTRlNTgtNDlhMC1iNjI5LTRiYzIxY2EzYjY1MCIsImQ3Y2ViMjBiLTc4YWUtNDQ0OC04MGFkLWMxMjc2ZGVlZWJlNyIsIjhmZTk0ODBmLWY3ZTItNDhhZi1iZDcxLWQ2Y2UyYmVkOTllNiIsIjFlYTgxMzEwLTRiZTEtNGMwYy1hN2MwLTY5NDZmOWIzNjU0ZSIsImY2MmQ3ZTEyLTAzZWMtNDUxYi05MjU3LWU5OWExNWRkMWZjMiIsImM0NTI2OTE2LTIyYjItNDE4OS05NTM1LTUxMjZiN2VjMGQ2YiIsIjI3MGVhYzM2LTZiYjYtNGQ5Yi1iNDBiLTJjMTY1NWFmM2YwMCIsIjRjYmIyZTM4LTg3ZWUtNDU4NS04N2NhLTIxNzYwZTNkMzYxZCIsIjgwMzcxMjQzLTQxMDktNDU2ZC05N2QzLTA0ZTliZGY1MWY5OCIsIjU2OWMyYTQ2LWM5YmItNDc3ZC1hODhkLWFmMTgyOWI3NmIxNyIsImJkYjRhYzQ5LWQwY2QtNDdjNy1iMWQxLWJkZmExN2VlZjdjYyIsIjdiZmY5MjU3LThhMWYtNDA5Ny05NzkwLTBmYzRkYTY5MGJhYyIsIjUzZGIwZTU5LTU3OWMtNDFmYy1hNmYzLTU1YTA4ZjBmMmJlNSIsImQ4ODU5NzViLWU3NTEtNDYzYy04YzAxLWFiYWZkMmI4OTAwYyIsIjdkMDMwNDYzLWUwZGUtNDg5Ny1hZjUyLWVmODQ4MWFhNjI3YiIsIjQ3NTJlZTc0LTA4MjEtNGEwZC1hZWNiLTczZjFmOWE1MmEwYiIsIjY3YjVmYTc0LTlhMmItNGU2ZS04NjZlLWIyNDZhN2QxNjRhMCIsIjUxMjQyNjg0LTU0MjYtNGZhZi05NzFhLTE1ZjVjYThiZWRmZSIsImVmY2Y0MTg3LTE4YzctNGE3OC1hMjM3LTVhYWJkOGIzNDYwZiIsIjBlZTVhYzhiLWYwNDktNDYxZi05ZjMwLTBiM2U0YmIyNWIyMCIsImVlMDA1ODkzLWE3ODMtNGNjMC1iNTgxLWYxZTJlNWEyNzM0MyIsImMzNzBkNjliLTA2MTItNGI3My1hOWRlLTNlMTg0NGQ1MjM5NyIsImE0N2VmNTljLWY3YzQtNDhkMi05NmIyLWZmOGJkNGI5NmJkZiIsIjRjMTMxYTllLWRkZTYtNGMwMC1hYmE0LTEwYWE0ZTU3ZGRmZCIsImVhMWE4MDllLTc3ZDgtNGJlYS05NGVlLTI5YjYwYzk3MjhmNiIsIjE4NjkwMDlmLWZjZDctNGU5Ny05YmM1LTliNTUxMjU0MTExMSIsImRlYmYxZGEwLWFkOTUtNDFlYi1iYTQ3LTQ0ODBiM2Y0ODBkYSIsImUwMGYxNmFkLTUwMjYtNDY1OC05Y2E5LTBmOTI0NGNlNmI1MCIsImY0ZjdkZGIyLTBhYTItNDBkNi05Mzg0LTc4MmRjYWE0ZjFlMiIsIjBiZjQ2MmI4LTI0NDgtNGIwMS1iZTRjLTIzOTI1M2UxNDEyMSIsIjA3ZDlmY2NjLWMzMTYtNGY3Mi04NTU3LTE5NzExYTkxZTYzZCIsIjg3ZDdhM2Q2LTFkZmItNDBiNC04NTM3LTZhMDc2MjZjMzc5MSIsIjlhNDVkYmRkLTE3ZWItNDgyMS1iNmMxLWMwOTI1YzM1MmM3YSIsIjI1NzdkMGVlLTU4ZTktNGNlNi04NzkxLWMwMWZiMWJlMWQ0YSIsIjJhOTY5NGZiLTExN2MtNDNmNC05YzQxLTQ5ZDgzZGMwYTgyYiIsImFkMmUwMWZlLTg5MGQtNDNhZC1hMWY3LTQxZWVhMmQ3NTVlYyIsImZmNjRiMGZlLWRlY2QtNDJhYy04OWIxLTc1NTA5ODc1NTEyNSJdLCJpZHR5cCI6InVzZXIiLCJpcGFkZHIiOiIxNDQuMTQwLjE1MC41IiwibmFtZSI6ImFkbS4xMTQ0MGUiLCJvaWQiOiJjMWYwZTNiNy00OTVlLTQ2NjEtOTAyNS1hZjQxNDUzYzNjODUiLCJvbnByZW1fc2lkIjoiUy0xLTUtMjEtMTA5Mzk4NjEwMC0xNzU3NDc3MjUyLTI4NDc3NTk5NzUtODYwNyIsInB1aWQiOiIxMDAzMjAwMzNCOTYxNjNBIiwicmgiOiIwLkFVSUExMTZLWkt6RnQwVzBCbmtmRFVNMHBrWklmM2tBdXRkUHVrUGF3ZmoyTUJOQ0FMUS4iLCJzY3AiOiJ1c2VyX2ltcGVyc29uYXRpb24iLCJzdWIiOiJ0N1Vvb2VNZVM4ZS16Ynl3ZlgxZTZ1UGVGT3I0Vl83VUtnQ3pqVkZTamdNIiwidGlkIjoiNjQ4YTVlZDctYzVhYy00NWI3LWI0MDYtNzkxZjBkNDMzNGE2IiwidW5pcXVlX25hbWUiOiJhZG0uMTE0NDBlQGFzaW8uZ292LmF1IiwidXBuIjoiYWRtLjExNDQwZUBhc2lvLmdvdi5hdSIsInV0aSI6Il9jLWdQTGVZQkVDbjNaYTJJbXNGQUEiLCJ2ZXIiOiIxLjAiLCJ3aWRzIjpbImIwZjU0NjYxLTJkNzQtNGM1MC1hZmEzLTFlYzgwM2YxMmVmZSIsIjYyZTkwMzk0LTY5ZjUtNDIzNy05MTkwLTAxMjE3NzE0NWUxMCIsImI3OWZiZjRkLTNlZjktNDY4OS04MTQzLTc2YjE5NGU4NTUwOSJdLCJ4bXNfY2FlIjoiMSIsInhtc190Y2R0IjoxNjA2ODYxNDUwfQ.UDjV7WWeRELnw6EDOtbiUStSDexz3DjAuJNFEx_gQpl9daUFwNuE8J4sf40FHHTlh_OvCEMP6llDmFFElP0i9hDZZdvXrDRTnialdfMYk6KRTzeflES8wVguSk3Tf_J80Njt8ymkmJUPf-7212rdfGPVywG7qGiYl-VSIUu3wB1lLpAjF9HQ0Y1yyjJTUVrgOcYiX4UsJRMJFca76m437KRao2qb2DsdJuQ_oCk1MhAXNcI20wzpwitEAQz6aOlmGFLQLn0xd8dcD4sh4XVsJeZgBxrGemij9otuZ-8TvQ8ytqGWK52fOYe52WNlhnqCVys0jhcSzNfLvJ5tkmBk_Q")

	res, err := http.DefaultClient.Do(req)
	if err != nil {

		log.Fatal("Error fetching LA Workspace Tables")
	}

	fmt.Println(res)

	responseBody, err := io.ReadAll(res.Body)
	if res.StatusCode == 400 {
		log.Fatal("Error fetching LA Workspace Tables: ", string(responseBody))
	}

	fmt.Println(string(responseBody))

	// var resBody providerListResult
	//
	// json.Unmarshal(responseBody, &resBody)

	// fmt.Println(resBody)

	// var azureResourceTypes azureResourceTypes

	// for _, val := range resBody.Value {
	// 	azureResourceTypes.ProviderNamespaces = append(azureResourceTypes.ProviderNamespaces, val.Namespace)

	// 	for _, resType := range val.ResourceTypes {
	// 		azureResourceTypes.ResourceTypes = append(azureResourceTypes.ResourceTypes, val.Namespace+"/"+resType.ResourceType)
	// 	}
	// }

	// jsonData, _ := json.MarshalIndent(resBody.Value, "", "  ")

	// fmt.Println(string(jsonData))
}