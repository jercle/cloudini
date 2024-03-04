package azure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"

	// "log"
	"os"

	"github.com/charmbracelet/log"
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
	InstallationID string         `json:"installationId"`
	Subscriptions  []Subscription `json:"subscriptions"`
}

// type T struct {
// 	InstallationID string `json:"installationId"`
// 	Subscriptions  []struct {
// 		EnvironmentName  string `json:"environmentName"`
// 		HomeTenantID     string `json:"homeTenantId"`
// 		ID               string `json:"id"`
// 		IsDefault        bool   `json:"isDefault"`
// 		ManagedByTenants []struct {
// 			TenantID string `json:"tenantId"`
// 		} `json:"managedByTenants"`
// 		Name     string `json:"name"`
// 		State    string `json:"state"`
// 		TenantID string `json:"tenantId"`
// 		User     struct {
// 			Name string `json:"name"`
// 			Type string `json:"type"`
// 		} `json:"user"`
// 	} `json:"subscriptions"`
// }

type Subscription struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	User struct {
		Name string `json:"name"`
	} `json:"user"`
	IsDefault bool   `json:"isDefault"`
	TenantID  string `json:"tenantId"`
}

func GetActiveSub() (string, error) {
	subs, _ := getSubs()

	for _, sub := range subs.Subscriptions {
		if sub.IsDefault {
			return sub.ID, nil
		}
	}

	return "", fmt.Errorf("no default subscription")
}

func getSubs() (AzureProfile, []byte) {
	userHomeDir, _ := os.UserHomeDir()
	content, readError := os.ReadFile(userHomeDir + "/.azure/azureProfile.json")
	// content, readError := os.ReadFile("/home/jercle/git/azg/testData/azCliProfile.json")

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
