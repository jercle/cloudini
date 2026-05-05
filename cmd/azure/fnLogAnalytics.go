package azure

import (
	"context"
	"encoding/json/jsontext"
	json "encoding/json/v2"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/fatih/color"
	"github.com/jercle/cloudini/lib"
	"github.com/rodaine/table"

	_ "time/tzdata"
)

type LogAnalyticsTables []struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Properties struct {
		ArchiveRetentionInDays   int    `json:"archiveRetentionInDays"`
		Plan                     string `json:"plan"`
		ProvisioningState        string `json:"provisioningState"`
		RetentionInDays          int    `json:"retentionInDays"`
		RetentionInDaysAsDefault bool   `json:"retentionInDaysAsDefault"`
		Schema                   struct {
			Columns []struct {
				IsDefaultDisplay bool   `json:"isDefaultDisplay"`
				IsHidden         bool   `json:"isHidden"`
				Name             string `json:"name"`
				Type             string `json:"type"`
			} `json:"columns,omitempty"`
			IsTroubleshootingAllowed bool     `json:"isTroubleshootingAllowed"`
			Name                     string   `json:"name"`
			Solutions                []string `json:"solutions"`
			StandardColumns          []struct {
				Description      string `json:"description"`
				IsDefaultDisplay bool   `json:"isDefaultDisplay"`
				IsHidden         bool   `json:"isHidden"`
				Name             string `json:"name"`
				Type             string `json:"type"`
			} `json:"standardColumns"`
			TableSubType string `json:"tableSubType"`
			TableType    string `json:"tableType"`
		} `json:"schema"`
		TotalRetentionInDays          int  `json:"totalRetentionInDays"`
		TotalRetentionInDaysAsDefault bool `json:"totalRetentionInDaysAsDefault"`
	} `json:"properties"`
}

//
//

type ResponseBody struct {
	Value LogAnalyticsTables `json:"value"`
}

//
//

func (tables *LogAnalyticsTables) filterByName(filter string, caseInsensitive bool) {
	var filteredTables LogAnalyticsTables
	if caseInsensitive {
		filter = strings.ToLower(filter)
		for _, table := range *tables {
			if strings.Contains(strings.ToLower(table.Name), filter) {
				filteredTables = append(filteredTables, table)
			}
		}
	} else {
		for _, table := range *tables {
			if strings.Contains(table.Name, filter) {
				filteredTables = append(filteredTables, table)
			}
		}
	}
	*tables = filteredTables
}

//
//

func (tables *LogAnalyticsTables) filterByRetentionInDaysAsDefault(filter bool) {
	var filteredTables LogAnalyticsTables
	for _, table := range *tables {
		if table.Properties.RetentionInDaysAsDefault == filter {
			filteredTables = append(filteredTables, table)
		}
	}
	*tables = filteredTables
}

//
//

func (tables *LogAnalyticsTables) printJSON() {
	jsonData, err := json.Marshal(tables, jsontext.WithIndent("  "))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}

//
//

func (tables *LogAnalyticsTables) printTable() {

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Name", "Retention Period", "isDefault")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, record := range *tables {
		retentionPeriod := strconv.Itoa(record.Properties.RetentionInDays)
		var isDefault string
		if record.Properties.RetentionInDaysAsDefault {
			isDefault = "True"
		} else {
			isDefault = "False"
		}
		tbl.AddRow(record.Name, retentionPeriod, isDefault)
	}
	tbl.Print()
}

//
//

func getAllWorkspaceTables(cred *azidentity.DefaultAzureCredential, subscriptionId string, resourceGroup string, workspaceName string) (LogAnalyticsTables, error) {
	ctx := context.Background()
	tokenRequestOptions := policy.TokenRequestOptions{
		Scopes: []string{
			"https://management.core.windows.net/.default",
		},
	}

	token, err := cred.GetToken(ctx, tokenRequestOptions)
	if err != nil {
		log.Fatal(err)
	}
	urlString := "https://management.azure.com/subscriptions/" +
		subscriptionId +
		"/resourceGroups/" +
		resourceGroup +
		"/providers/Microsoft.OperationalInsights/workspaces/" +
		workspaceName +
		"/tables?api-version=2022-10-01"
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {

		log.Fatal("Error fetching LA Workspace Tables")
	}

	responseBody, err := io.ReadAll(res.Body)
	if res.StatusCode == 400 {
		log.Fatal("Error fetching LA Workspace Tables: ", string(responseBody))
	}

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	var responseUnmarshalled ResponseBody
	json.Unmarshal(responseBody, &responseUnmarshalled)

	sort.Slice(responseUnmarshalled.Value, func(i, j int) bool {
		return strings.ToLower(responseUnmarshalled.Value[i].Name) < strings.ToLower(responseUnmarshalled.Value[j].Name)
	})

	return responseUnmarshalled.Value, err
}

//
//

func GetAzureWorkbookAlerts(graphQuery string, token *lib.AzureMultiAuthToken) (alerts []AzureAlertProcessed) {
	// fmt.Println("Fetching alerts for ", token.TenantName)
	logAnalyticsToken, err := GetTenantSPToken(lib.AzureMultiAuthTokenRequestOptions{
		TenantName: token.TenantName,
		Scope:      "loganalytics",
	}, nil)
	lib.CheckFatalError(err)

	urlString := "https://management.azure.com/providers/Microsoft.ResourceGraph/resources?api-version=2021-03-01"

	jsonBody := `{"query": "` + graphQuery + `"}`

	res, _, err := HttpPost(urlString, jsonBody, *token)
	lib.CheckFatalError(err)

	// fmt.Println(string(res))
	// os.WriteFile("/home/jercle/git/cld/cmd/azure/fnLogAnalytics.json", res, 0644)
	// os.Exit(0)

	var alertsResponse GetAzureAlertsResponse
	err = json.Unmarshal(res, &alertsResponse)

	currentTime := time.Now()

	cldConf := lib.GetCldConfig(nil)
	var locale *time.Location
	locale, err = time.LoadLocation("Australia/Sydney")
	// if err != nil {
	// 	file, _ := os.ReadFile("/etc/localtime")
	// 	locale, err = time.LoadLocationFromTZData("Australia/Sydney", file)
	// }
	lib.CheckFatalError(err)

	timeInLocal := false

	if strings.Contains(graphQuery, "datetime_utc_to_local") {
		timeInLocal = true
	}

	for _, alert := range alertsResponse.Data {
		if alert.AlertState == "Closed" {
			continue
		}
		jsonStr, _ := json.Marshal(alert, jsontext.WithIndent("  "))
		// fmt.Println(string(jsonStr))
		// os.Exit(0)
		var curr AzureAlertProcessed
		err = json.Unmarshal(jsonStr, &curr)
		lib.CheckFatalError(err)
		curr.TenantName = token.TenantName
		curr.AzureWorkbookUrl = "https://portal.azure.com/#@" + token.TenantId + "/resource" + cldConf.Azure.SupportAlerts.TenantWorkbookIds[token.TenantName] + "/workbook"
		// lib.JsonMarshalAndPrint(alert.UnknownFields)
		// fmt.Println(string(jsonStr))
		// os.Exit(0)

		var alertCreated time.Time
		if timeInLocal {
			alertCreated, err = time.ParseInLocation("15:04 PM 02-01-06", alert.AlertCreated, locale)
		} else {
			alertCreated, err = time.Parse("15:04 PM 02-01-06", alert.AlertCreated)
		}
		lib.CheckFatalError(err)

		curr.AlertCreated = alertCreated

		var alertLastModified time.Time
		if timeInLocal {
			alertLastModified, err = time.ParseInLocation("15:04 PM 02-01-06", alert.AlertLastModified, locale)
		} else {
			alertLastModified, err = time.Parse("15:04 PM 02-01-06", alert.AlertLastModified)
		}
		lib.CheckFatalError(err)
		curr.AlertLastModified = alertLastModified

		curr.LastAzureSync = currentTime
		if alert.Properties.Context.Context != nil {
			curr.LinkToFilteredSearchResultsAPI = alert.Properties.Context.Context.Condition.AllOf[0].LinkToFilteredSearchResultsAPI
			curr.LinkToFilteredSearchResultsUi = alert.Properties.Context.Context.Condition.AllOf[0].LinkToFilteredSearchResultsUi
			if curr.LinkToFilteredSearchResultsAPI != "" {
				// fmt.Println("Fetching alert data for ", curr.TenantName, curr.Name)
				ad, err := GetAlertDataFromSearchResultsLink(curr.LinkToFilteredSearchResultsAPI, curr.LinkToFilteredSearchResultsUi, logAnalyticsToken)
				// fmt.Println(curr.LinkToFilteredSearchResultsAPI)
				// lib.CheckFatalError(err)
				if err != nil {
					fmt.Println(err)
					lib.JsonMarshalAndPrint(curr)
					os.Exit(1)
				}
				curr.AlertData = ad
			}
		}
		alerts = append(alerts, curr)
	}

	return
}

//
//

func GetAlertDataFromSearchResultsLink(linkToFilteredSearchResultsAPI string, linkToFilteredSearchResultsUi string, token *lib.AzureMultiAuthToken) (alertData []map[string]any, e error) {
	res, err := HttpGet(linkToFilteredSearchResultsAPI, *token)
	if err != nil {
		return nil, err
	}
	// lib.CheckFatalError(err)

	var resData GetAlertDataFromSearchResultsLinkResult
	err = json.Unmarshal(res, &resData)
	if err != nil {
		return alertData, err
	}
	// lib.CheckFatalError(err)

	if len(resData.Tables) == 0 {
		data := make(map[string]any)
		data["queryError"] = string(res)
		data["queryLink"] = linkToFilteredSearchResultsUi
		// fmt.Println(string(res))
		// fmt.Println(linkToFilteredSearchResultsAPI)
		// fmt.Println(token.TenantName)
		// os.Exit(0)
		// err = fmt.Errorf(string(res))
		// return nil, &err
		alertData = append(alertData, data)
		return

	}

	columns := resData.Tables[0].Columns
	rows := resData.Tables[0].Rows

	for _, rowData := range rows {
		rowProcessed := make(map[string]any)
		for i, prop := range rowData {
			rowProcessed[columns[i].Name] = prop
		}
		alertData = append(alertData, rowProcessed)
	}

	return
}

//
//

func GetLogAnalyticsWorkbookQuery(resourceId string, token *lib.AzureMultiAuthToken) string {
	urlString := "https://management.azure.com" + resourceId + "?api-version=2021-08-01&canFetchContent=true"

	res, err := HttpGet(urlString, *token)
	lib.CheckFatalError(err)

	// fmt.Println(string(res))

	var resData LogAnalyticsWorkbook
	err = json.Unmarshal(res, &resData)
	lib.CheckFatalError(err)
	// lib.JsonMarshalAndPrint(resData)

	var serializedData LogAnalyticsWorkbookSerializedData
	err = json.Unmarshal([]byte(resData.Properties.SerializedData), &serializedData)
	lib.CheckFatalError(err)
	query := serializedData.Items[1].Content.Query
	// jsonStr1, _ := json.Marshal(serializedData)
	// os.WriteFile("/home/jercle/git/cld/cmd/azure/fnLogAnalytics.json", jsonStr1, 0644)
	// lib.JsonMarshalAndPrint(serializedData)
	// for i, query := range serializedData.Items {
	// 	fmt.Println(strconv.Itoa(i), query.Name)
	// }
	// os.Exit(0)

	jsonStr, _ := json.Marshal(query)
	queryString := strings.TrimSuffix(string(jsonStr), "\"")
	queryString = strings.TrimPrefix(queryString, "\"")

	return queryString
}

//
//

//
//

func RunLogAnalyticsQuery(workspaceId string, query string, token lib.AzureMultiAuthToken) (results LogAnalyticsQueryResponse) {
	urlString := "https://api.loganalytics.azure.com/v1/workspaces/" + workspaceId + "/query"

	queryStr, _ := json.Marshal(query)
	jsonBody := `{"query":` + string(queryStr) + `, "timespan": "PT24H"}`

	res, _, err := HttpPost(urlString, jsonBody, token)
	lib.CheckFatalError(err)

	var resData RunLogAnalyticsQueryResponseRaw
	err = json.Unmarshal(res, &resData)
	lib.CheckFatalError(err)

	for _, tbl := range resData.Tables {
		var tableData LogAnalyticsQueryResponseTable
		tableData.Name = tbl.Name
		for _, row := range tbl.Rows {
			rowData := make(map[string]any)
			for ri, rowCol := range row {
				rowData[tbl.Columns[ri].Name] = rowCol
			}
			tableData.Rows = append(tableData.Rows, rowData)
		}
		results.Tables = append(results.Tables, tableData)
	}

	return
}
