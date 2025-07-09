package azure

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

type ResponseBody struct {
	Value LogAnalyticsTables `json:"value"`
}

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

func (tables *LogAnalyticsTables) filterByRetentionInDaysAsDefault(filter bool) {
	var filteredTables LogAnalyticsTables
	for _, table := range *tables {
		if table.Properties.RetentionInDaysAsDefault == filter {
			filteredTables = append(filteredTables, table)
		}
	}
	*tables = filteredTables
}

func (tables *LogAnalyticsTables) printJSON() {
	jsonData, err := json.MarshalIndent(tables, "", "  ")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}

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

	// fmt.Println(string(responseBody))

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

func GetAzureWorkbookAlerts(token *lib.AzureMultiAuthToken) (alerts []AzureAlertProcessed) {
	logAnalyticsToken, err := GetTenantSPToken(lib.AzureMultiAuthTokenRequestOptions{
		TenantName: token.TenantName,
		Scope:      "loganalytics",
	}, nil)
	lib.CheckFatalError(err)

	urlString := "https://management.azure.com/providers/Microsoft.ResourceGraph/resources?api-version=2021-03-01"

	graphQuery := `alertsmanagementresources
	| where type == 'microsoft.alertsmanagement/alerts'
	| extend AlertCreated = todatetime(properties[\"essentials\"].[\"startDateTime\"])
	| where AlertCreated > ago(7d)
	| extend Severity = tostring(properties[\"essentials\"][\"severity\"])
	| extend Results = tostring(properties.[\"context\"].[\"context\"].[\"condition\"].[\"allOf\"].[0].linkToFilteredSearchResultsUI)
	| extend ResourceID1 = tostring(properties.[\"context\"].[\"linkedResourceId\"])
	| extend ResourceID2 = tostring(properties.[\"context\"].[\"context\"].[\"resourceId\"])
	| extend AffectedResource = coalesce(ResourceID1, ResourceID2)
	| extend AffectedResource = iff(AffectedResource contains \"la-\", \"\", AffectedResource)
	| extend Description =  tostring(properties[\"essentials\"].[\"description\"])
	| extend AlertLastModified = todatetime(properties[\"essentials\"].[\"lastModifiedDateTime\"])
	| extend AlertLastModifiedBy = tostring(properties[\"essentials\"].[\"lastModifiedUserName\"])
	| extend AlertState = tostring(properties[\"essentials\"].[\"alertState\"])
	| extend AlertLastModifiedBy = iff(AlertLastModifiedBy == \"System\", \"Not triaged\", AlertLastModifiedBy)
	| where properties[\"essentials\"][\"monitorCondition\"] in~ ('Fired')
	| extend TriageAlert = \"Acknowledge\"
	| project AlertCreated, Severity, Name = name, AffectedResource, Description, Results, AlertState, TriageAlert, AlertLastModifiedBy, AlertLastModified, properties, id
	| extend AlertCreated1 = todatetime(AlertCreated)
	| order by AlertCreated1 desc
	| extend AlertCreated = datetime_utc_to_local(AlertCreated, 'Australia/Canberra')
	| extend AlertLastModified = datetime_utc_to_local(AlertLastModified, 'Australia/Canberra')
	| extend AlertCreated =  format_datetime( AlertCreated, \"HH:mm tt dd-MM-yy\")
	| extend AlertLastModified = format_datetime( AlertLastModified, \"HH:mm tt dd-MM-yy\")
	| order by AlertCreated1 desc
	| project-away AlertCreated1`

	jsonBody := `{"query": "` + graphQuery + `"}`

	res, _, err := HttpPost(urlString, jsonBody, *token)
	lib.CheckFatalError(err)

	var alertsResponse GetAzureAlertsResponse
	err = json.Unmarshal(res, &alertsResponse)

	currentTime := time.Now()

	for _, alert := range alertsResponse.Data {
		jsonStr, _ := json.Marshal(alert)
		var curr AzureAlertProcessed
		err = json.Unmarshal(jsonStr, &curr)
		// lib.CheckFatalError(err)
		curr.TenantName = token.TenantName
		alertCreated, err := time.Parse("15:04 PM 01-02-06", alert.AlertCreated)
		lib.CheckFatalError(err)
		curr.AlertCreated = alertCreated
		alertLastModified, err := time.Parse("15:04 PM 01-02-06", alert.AlertLastModified)
		lib.CheckFatalError(err)
		curr.AlertLastModified = alertLastModified

		// 15:57 PM 07-07-25

		curr.LastAzureSync = currentTime
		if alert.Properties.Context.Context != nil {
			curr.LinkToFilteredSearchResultsAPI = alert.Properties.Context.Context.Condition.AllOf[0].LinkToFilteredSearchResultsAPI
			curr.AlertData = GetAlertDataFromSearchResultsLink(curr.LinkToFilteredSearchResultsAPI, logAnalyticsToken)
		}
		alerts = append(alerts, curr)
	}

	return
}

func GetAlertDataFromSearchResultsLink(linkToFilteredSearchResultsAPI string, token *lib.AzureMultiAuthToken) (alertData []map[string]any) {
	res, err := HttpGet(linkToFilteredSearchResultsAPI, *token)
	lib.CheckFatalError(err)

	var resData GetAlertDataFromSearchResultsLinkResult
	err = json.Unmarshal(res, &resData)

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
