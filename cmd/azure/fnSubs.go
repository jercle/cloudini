package azure

import (
	"bytes"
	"cmp"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"sort"
	"strings"

	// "log"
	"os"

	"github.com/jercle/cloudini/lib"
	"github.com/charmbracelet/log"
	"github.com/manifoldco/promptui"
)

// func GetTenant() {
// options := arm.ClientOptions{
// 	ClientOptions: azcore.ClientOptions{
// 		Cloud: cloud.AzurePublic
// 	},
// }
// clientFactory, err := arm.NewClient()

// if err != nil {
// 	log.Error("", err, err)
// }

// client, err := arm.NewClient()

// client.
// }

type AzureProfile struct {
	InstallationID string                     `json:"installationId"`
	Subscriptions  []AzureProfileSubscription `json:"subscriptions"`
}

type AzureProfileSubscription struct {
	EnvironmentName  string `json:"environmentName"`
	HomeTenantID     string `json:"homeTenantId,omitempty"`
	ID               string `json:"id"`
	IsDefault        bool   `json:"isDefault"`
	ManagedByTenants []struct {
		TenantID string `json:"tenantId"`
	} `json:"managedByTenants"`
	Name                string `json:"name"`
	State               string `json:"state"`
	TenantDefaultDomain string `json:"tenantDefaultDomain,omitempty"`
	TenantDisplayName   string `json:"tenantDisplayName,omitempty"`
	TenantMappedName    string `json:"tenantMappedName,omitempty"`
	TenantID            string `json:"tenantId"`
	User                struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"user"`
}

type AzureProfileTenant struct {
	TenantDefaultDomain string `json:"tenantDefaultDomain,omitempty"`
	TenantDisplayName   string `json:"tenantDisplayName,omitempty"`
	TenantID            string `json:"tenantId"`
	TenantMappedName    string `json:"tenantMappedName,omitempty"`
}

func GetActiveCliSub() (*AzureProfileSubscription, error) {
	subs, _ := GetCliSubs()

	for _, sub := range subs.Subscriptions {
		if sub.IsDefault {
			return &sub, nil
		}
	}

	return nil, fmt.Errorf("no default subscription")
}

func GetCliSubs() (AzureProfile, []byte) {
	userHomeDir, _ := os.UserHomeDir()
	content, readError := os.ReadFile(userHomeDir + "/.azure/azureProfile.json")
	// content, readError := os.ReadFile(userHomeDir + "/git/azg/testData/azCliProfile.json")
	// fmt.Println(string(content))
	if readError != nil {
		log.Fatal("Error when opening Azure Profile. Have you logged into az-cli?", readError)
	}
	content = bytes.TrimPrefix(content, []byte("\xef\xbb\xbf"))

	var userAzureProfile AzureProfile
	// var payload map[string]interface{}
	unmarshalError := json.Unmarshal(content, &userAzureProfile)

	// fmt.Println(unmarshalError)

	if unmarshalError != nil {
		log.Error("Error during Unmarshal():", "err", unmarshalError)
	}

	// fmt.Println(userAzureProfile)

	return userAzureProfile, content
}

func (s *AzureProfile) PrintSubs() {
	for _, sub := range s.Subscriptions {
		var subString string = sub.ID + " - " + sub.Name

		if sub.IsDefault {
			subString += " - Current active"
		}
		fmt.Println(subString)
	}
}

func (s *AzureProfile) Sort() {
	keys := make([]string, 0, len(s.Subscriptions))

	for _, k := range s.Subscriptions {
		keys = append(keys, k.TenantID)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	for k := range keys {
		fmt.Println(s.Subscriptions[k])
	}
}

// Lists Azure subscriptions availabe to a given auth token
func ListSubscriptions(token lib.AzureMultiAuthToken) ([]lib.FetchedSubscription, error) {
	urlString := "https://management.azure.com/subscriptions?api-version=2022-12-01"
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.TokenData.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// log.Fatal("Error fetching list of Subscriptions")
		return nil, err
	}

	// jsonStr, _ := json.MarshalIndent(res.Status, "", "  ")
	// fmt.Println(string(jsonStr))

	responseBody, err := io.ReadAll(res.Body)
	if res.StatusCode == 400 {
		// log.Fatal("Error fetching list of Subscriptions: ", string(responseBody))
		return nil, err
	}

	if err != nil {
		// log.Fatal(err)
		return nil, err
	}
	defer res.Body.Close()

	var subsList lib.SubsReqResBody
	json.Unmarshal(responseBody, &subsList)
	subsList.UpdateTenantName(token.TenantName)
	// lib.MarshalAndPrintJson(subsList.Value)

	return subsList.Value, nil
}

func TenantMapByDomain() map[string]string {
	config := lib.GetCldConfig(nil)

	tenantMapByDomain := make(map[string]string)
	for name, domain := range config.Azure.TenantMap {
		tenantMapByDomain[domain] = name
	}

	return tenantMapByDomain
}

func PromptSelectTenant() (AzureProfileTenant, []AzureProfileSubscription) {
	var (
		azProfileData AzureProfile
	)
	usrHomeDir, err := os.UserHomeDir()
	azureProfile := usrHomeDir + "/.azure/azureProfile.json"
	lib.CheckFatalError(err)
	azProfileFile, err := os.ReadFile(azureProfile)
	lib.CheckFatalError(err)
	azProfileBomRemoved := lib.RemoveJsonByteOrderMark(azProfileFile)
	err = json.Unmarshal(azProfileBomRemoved, &azProfileData)
	lib.CheckFatalError(err)

	tenantMapByDomain := TenantMapByDomain()

	byTenant := make(map[string][]AzureProfileSubscription)
	tenants := make(map[string]AzureProfileTenant)
	var tenantsSlice []AzureProfileTenant

	for _, sub := range azProfileData.Subscriptions {
		if sub.TenantDisplayName == "" {
			continue
		}
		sub.TenantMappedName = tenantMapByDomain[sub.TenantDefaultDomain]
		byTenant[sub.TenantMappedName] = append(byTenant[sub.TenantMappedName], sub)
		var t AzureProfileTenant
		t.TenantDefaultDomain = sub.TenantDefaultDomain
		t.TenantDisplayName = sub.TenantDisplayName
		t.TenantID = sub.TenantID
		t.TenantMappedName = sub.TenantMappedName
		tenants[sub.TenantMappedName] = t
	}

	for _, t := range tenants {
		tenantsSlice = append(tenantsSlice, t)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "* {{ .TenantMappedName | white | underline }}",
		Inactive: "  {{ .TenantMappedName | cyan }}",
		Selected: "Selected tenant: {{ .TenantMappedName | faint | cyan }}",
		Details: `
--------- Selected Tenant Details ----------
{{ "Name:" | faint }}	{{ .TenantMappedName }}
{{ "ID:" | faint }}	{{ .TenantID }}
{{ "Default Domain:" | faint }}	{{ .TenantDefaultDomain }}
{{ "Display Name:" | faint }}	{{ .TenantDisplayName }}`,
	}

	searcher := func(input string, index int) bool {
		tenant := tenantsSlice[index]
		name := strings.Replace(strings.ToLower(tenant.TenantMappedName), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Select Tenant",
		Items:     tenantsSlice,
		Templates: templates,
		Size:      8,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()
	selectedTenant := tenantsSlice[i]
	subs := byTenant[selectedTenant.TenantMappedName]

	// updatedAzureProfile.InstallationID = azProfileData.InstallationID

	return selectedTenant, subs
}

func ChangeActiveSubscription(subs []AzureProfileSubscription) {
	usrHomeDir, err := os.UserHomeDir()
	azureProfile := usrHomeDir + "/.azure/azureProfile.json"
	lib.CheckFatalError(err)
	azProfileFile, err := os.ReadFile(azureProfile)
	lib.CheckFatalError(err)
	azProfileBomRemoved := lib.RemoveJsonByteOrderMark(azProfileFile)
	var azProfileData AzureProfile
	err = json.Unmarshal(azProfileBomRemoved, &azProfileData)
	lib.CheckFatalError(err)

	var currentActive AzureProfileSubscription

	for _, sub := range azProfileData.Subscriptions {
		if sub.IsDefault {
			currentActive = sub
		}
	}

	slices.SortFunc(subs, func(a, b AzureProfileSubscription) int {
		return cmp.Or(
			cmp.Compare(a.Name, b.Name),
		)
	})

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "* {{ .Name | white | underline }} ({{ .ID | faint }})",
		Inactive: "  {{ .Name | cyan }} ({{ .ID | faint }})",
		Selected: "New active subscription: {{ .Name | faint | cyan }}",
		Details: `
Current Active: ` + currentActive.Name + ` - ` + currentActive.TenantDefaultDomain + `
--------- Selected Sub Details ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "ID:" | faint }}	{{ .ID }}
{{ "Username:" | faint }}	{{ .User.Name }}
{{ "Tenant:" | faint }}	{{ .TenantDisplayName }}
{{ "Tenant ID:" | faint }}	{{ .TenantID }}`,
	}

	searcher := func(input string, index int) bool {
		sub := subs[index]
		name := strings.Replace(strings.ToLower(sub.Name), " ", "", -1)
		tenantName := strings.Replace(strings.ToLower(sub.TenantDisplayName), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input) || strings.Contains(tenantName, input)
	}

	prompt := promptui.Select{
		Label:     "Select Subscription",
		Items:     subs,
		Templates: templates,
		Size:      8,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	var updatedAzureProfile AzureProfile
	updatedAzureProfile.InstallationID = azProfileData.InstallationID

	for _, sub := range azProfileData.Subscriptions {
		currSub := sub
		if subs[i].ID == sub.ID {
			currSub.IsDefault = true
			updatedAzureProfile.Subscriptions = append(updatedAzureProfile.Subscriptions, currSub)
		} else {
			currSub.IsDefault = false
			updatedAzureProfile.Subscriptions = append(updatedAzureProfile.Subscriptions, currSub)
		}
	}
	jsonStr, _ := json.Marshal(updatedAzureProfile)
	os.WriteFile(azureProfile, jsonStr, 0644)
}

func ListAllAuthenticatedSubscriptions(tokens *lib.AllTenantTokens) TenantList {
	// allSubscriptions := make(map[string]string)
	// allTenantSubs := TenantList{}
	var allTenants TenantList

	for _, token := range *tokens {
		subs, err := ListSubscriptions(token)
		lib.CheckFatalError(err)
		var currTenant TenantDetails
		currTenant.Subscriptions = make(map[string]string)
		currTenant.TenantId = token.TenantId
		currTenant.TenantName = token.TenantName

		for _, sub := range subs {
			currTenant.Subscriptions[sub.DisplayName] = sub.SubscriptionID
		}
		allTenants = append(allTenants, currTenant)
	}

	// jsonStr, _ := json.MarshalIndent(allTenants, "", "  ")
	// fmt.Println(string(jsonStr))
	// os.Exit(0)
	return allTenants
}
