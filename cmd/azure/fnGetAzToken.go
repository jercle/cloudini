package azure

import (
	"bytes"
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/charmbracelet/log"
	"github.com/jercle/azg/cmd"
	"github.com/jercle/azg/lib"
)

var usrHomeDir, err = os.UserHomeDir()
var tCacheFile = usrHomeDir + "/.config/cld/tCache"

type MultiAuthToken struct {
	TenantId   string
	TenantName string
	TokenData  TokenData
}
type Request struct {
	Url     string
	Outfile string
}

type TokenData struct {
	Token     string
	ExpiresOn string
}

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

type GetAllTenantTokenOptions struct {
	GetWriteToken  bool
	ConfigFilePath string
}

type FetchedSubscription struct {
	AuthorizationSource  string   `json:"authorizationSource"`
	DisplayName          string   `json:"displayName"`
	ID                   string   `json:"id"`
	ManagedByTenants     []string `json:"managedByTenants"`
	State                string   `json:"state"`
	SubscriptionID       string   `json:"subscriptionId"`
	SubscriptionPolicies struct {
		LocationPlacementID string `json:"locationPlacementId"`
		QuotaID             string `json:"quotaId"`
		SpendingLimit       string `json:"spendingLimit"`
	} `json:"subscriptionPolicies"`
	TenantID   string `json:"tenantId"`
	TenantName string `json:"tenantName"`
}

type SubsReqResBody struct {
	Count struct {
		Type  string  `json:"type"`
		Value float64 `json:"value"`
	} `json:"count"`
	Value []FetchedSubscription `json:"value"`
}

type AllTenantTokens []MultiAuthToken

func (tokens *AllTenantTokens) SaveToFile() {

	lib.CheckFatalError(err)
	byteData, err := json.Marshal(tokens)
	lib.CheckFatalError(err)
	if _, err := os.Stat(tCacheFile); err != nil {
		os.Create(tCacheFile)
	}
	encodedData := b64.StdEncoding.EncodeToString(byteData)
	os.WriteFile(tCacheFile, []byte(encodedData), os.ModePerm)
	fmt.Println(encodedData)
}

func (tokens *AllTenantTokens) CheckExpiry() {
	fmt.Println(tokens)
}

func (subs *SubsReqResBody) UpdateTenantName(tenantName string) {
	var localSubs SubsReqResBody
	localSubs.Count = subs.Count
	for _, sub := range subs.Value {
		sub.TenantName = tenantName
		localSubs.Value = append(localSubs.Value, sub)
	}
	*subs = localSubs
}

func GetCachedTokens() AllTenantTokens {
	var (
		tokens AllTenantTokens
	)
	fileData, err := os.ReadFile(tCacheFile)
	lib.CheckFatalError(err)
	byteData, err := b64.StdEncoding.DecodeString(string(fileData))
	lib.CheckFatalError(err)
	json.Unmarshal(byteData, &tokens)
	return tokens
}

func GetServicePrincipalToken(tenant string, spDetails cmd.CldConfigClientAuthDetails) (*TokenData, error) {
	ctx := context.Background()

	tokenRequestOptions := policy.TokenRequestOptions{
		Scopes: []string{
			"https://management.core.windows.net/.default",
		},
		EnableCAE: true,
	}

	cred, err := azidentity.NewClientSecretCredential(tenant, spDetails.ClientID, spDetails.ClientSecret, nil)
	if err != nil {
		log.Error("Unable to obtain Azure token", err, err)
		return nil, err
	}
	// envCred, err := azidentity.NewEnvironmentCredential(nil)
	// if err != nil {
	// 	log.Error("Unable to obtain Azure token", err, err)
	// }

	tokenResponse, err := cred.GetToken(ctx, tokenRequestOptions)
	if err != nil {
		log.Error("Unable to obtain Azure token", err, err)
		return nil, err
	}

	token := TokenData{
		Token:     tokenResponse.Token,
		ExpiresOn: tokenResponse.ExpiresOn.Local().String(),
	}

	// fmt.Println(token)
	return &token, nil
}

func GetAzCliToken() TokenData {
	ctx := context.Background()
	tokenRequestOptions := policy.TokenRequestOptions{
		Scopes: []string{
			"https://management.core.windows.net/.default",
		},
		EnableCAE: true,
	}

	cliCred, err := azidentity.NewAzureCLICredential(nil)
	if err != nil {
		log.Error("Unable to obtain Azure token", err, err)
	}

	tokenResponse, err := cliCred.GetToken(ctx, tokenRequestOptions)
	if err != nil {
		log.Error("Unable to obtain Azure token", err, err)
	}

	token := TokenData{
		Token:     tokenResponse.Token,
		ExpiresOn: tokenResponse.ExpiresOn.Local().String(),
	}

	// fmt.Println(token)
	return token
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

// Gets a token for each tenant configured in the cld config file
//
// Default path is ~/.config/cld/config.json
//
// First parameter passed into this function overwrites the config file path
func GetAllTenantTokens(options GetAllTenantTokenOptions) (AllTenantTokens, error) {
	var (
		configPath   string
		config       cmd.CldConfig
		tenantTokens []MultiAuthToken
		wg           sync.WaitGroup
		mut          sync.Mutex
		homeDir, _   = os.UserHomeDir()
	)

	if options.ConfigFilePath != "" {
		configPath = options.ConfigFilePath
	} else {
		configPath = homeDir + "/.config/cld/config.json"
	}

	jsonConfig, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonConfig.Close()

	byteValue, _ := io.ReadAll(jsonConfig)
	json.Unmarshal(byteValue, &config)

	for _, tenant := range config.Azure.TenantAuth.Tenants {
		wg.Add(1)
		go func() {
			var tokenData *TokenData
			// fmt.Println("Getting token for " + tenant.TenantName)
			if options.GetWriteToken {
				tokenData, err = GetServicePrincipalToken(tenant.TenantID, tenant.Writer)
			} else if !options.GetWriteToken {
				tokenData, err = GetServicePrincipalToken(tenant.TenantID, tenant.Reader)
			}
			if err != nil {
				log.Fatal(err)
			}
			tenantToken := MultiAuthToken{
				TenantId:   tenant.TenantID,
				TenantName: tenant.TenantName,
				TokenData:  *tokenData,
			}
			mut.Lock()
			tenantTokens = append(tenantTokens, tenantToken)
			mut.Unlock()
			// fmt.Println("Obtained token for " + tenant.TenantName)
			wg.Done()
		}()
	}
	wg.Wait()
	return tenantTokens, nil
}