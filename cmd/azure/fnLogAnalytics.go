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

	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/fatih/color"
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
