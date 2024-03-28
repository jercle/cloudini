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
	"github.com/jercle/azg/lib"
	"github.com/jercle/azg/lib/cldTypes"
)

// var usrHomeDir, err = os.UserHomeDir()
var configPath = lib.InitConfig(cldTypes.CldConfigOptions{})
var tCacheFile = configPath + "/tCache"

func (tokens *AllTenantTokens) SaveToFile() {

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

func (tokens AllTenantTokens) SelectTenant(tenantName string) (*MultiAuthToken, error) {
	// var tenantToken MultiAuthToken
	// fmt.Println(tenantName)
	var tenantToken *MultiAuthToken

	for _, token := range tokens {
		if token.TenantName == tenantName {
			tenantToken = &token
			break
		}
	}

	if tenantToken != nil {
		return tenantToken, nil
	} else {
		return nil, fmt.Errorf("Token not found for supplied tenant name")
	}
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
	var tokens AllTenantTokens
	if _, err := os.Stat(configPath); err != nil {
		os.MkdirAll(configPath, os.ModePerm)
	}
	if _, err := os.Stat(tCacheFile); err != nil {
		os.Create(tCacheFile)
	}
	fileData, err := os.ReadFile(tCacheFile)
	lib.CheckFatalError(err)
	byteData, err := b64.StdEncoding.DecodeString(string(fileData))
	lib.CheckFatalError(err)
	json.Unmarshal(byteData, &tokens)
	if len(tokens) == 0 {
		fmt.Println("Fetching new tokens")
		tokens, err = GetAllTenantSPTokens(MultiAuthTokenRequestOptions{})
		lib.CheckFatalError(err)
	}
	fmt.Println(tokens)
	return tokens
}

func GetServicePrincipalToken(tenant string, spDetails MultiAuthTokenRequestOptions) (*TokenData, error) {
	ctx := context.Background()
	var tokenRequestOptions policy.TokenRequestOptions

	switch spDetails.Scope {
	case "graph":
		tokenRequestOptions.Scopes = []string{"https://graph.microsoft.com/.default"}
	default:
		tokenRequestOptions.Scopes = []string{"https://management.core.windows.net/.default"}
	}
	tokenRequestOptions.EnableCAE = true

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
	lib.CheckFatalError(err)

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

func GetToken(tenantName string) MultiAuthToken {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	lib.CheckFatalError(err)

	sub, err := GetActiveSub()
	lib.CheckFatalError(err)
	// fmt.Println(sub.TenantID)
	// os.Exit(0)

	ctx := context.Background()
	tokenRequestOptions := policy.TokenRequestOptions{
		Scopes: []string{
			"https://management.core.windows.net/.default",
		},
	}

	tokenResponse, err := cred.GetToken(ctx, tokenRequestOptions)
	lib.CheckFatalError(err)

	multiAuthToken := MultiAuthToken{
		TenantId:   sub.TenantID,
		TenantName: tenantName,
		TokenData: TokenData{
			Token:     tokenResponse.Token,
			ExpiresOn: tokenResponse.ExpiresOn.Local().String(),
		},
	}

	return multiAuthToken
}

// Gets a token for each tenant configured in the cld config file
//
// Default path is ~/.config/cld/config.json
//
// First parameter passed into this function overwrites the config file path
func GetAllTenantSPTokens(options MultiAuthTokenRequestOptions) (AllTenantTokens, error) {
	var (
		config       = lib.GetCldConfig(lib.CldConfigOptions{})
		tenantTokens []MultiAuthToken
		wg           sync.WaitGroup
		mut          sync.Mutex
		err          error
	)

	for _, tenant := range config.Azure.MultiTenantAuth.Tenants {
		wg.Add(1)
		go func() {
			var (
				tokenData *TokenData
				authOpts  MultiAuthTokenRequestOptions
			)

			// fmt.Println("Getting token for " + tenant.TenantName)

			jsonTok, _ := json.Marshal(options)

			options.ClientID = tenant.Writer.ClientID
			options.ClientSecret = tenant.Writer.ClientSecret

			switch options.GetWriteToken {
			case true:
				options.ClientID = tenant.Writer.ClientID
				options.ClientSecret = tenant.Writer.ClientSecret
			default:
				options.ClientID = tenant.Reader.ClientID
				options.ClientSecret = tenant.Reader.ClientSecret
			}

			switch options.Scope {
			case "graph":

			default:
			}

			if options.GetWriteToken {
				// var authOpts MultiAuthTokenRequestOptions
				// authOpts := MultiAuthTokenRequestOptions(tenant.Writer)

				// if options.Scope == "graph" {
				// 	// authOpts.AzureGraph = true

				// }
				tokenData, err = GetServicePrincipalToken(tenant.TenantID, authOpts)
				lib.CheckFatalError(err)
			} else if !options.GetWriteToken {
				authOpts := tenant.Reader
				// if options.AzureGraph {
				// 	authOpts.AzureGraph = true
				// }
				tokenData, err = GetServicePrincipalToken(tenant.TenantID, authOpts)
				lib.CheckFatalError(err)
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

func GetTenantSPToken(tenantName string, options MultiAuthTokenRequestOptions) (*MultiAuthToken, error) {
	var (
		config      = lib.GetCldConfig(lib.CldConfigOptions{})
		tenantToken MultiAuthToken
		err         error
	)

	// if tenant.TenantName == tenantName {
	// 	return nil, fmt.Errorf("Tenant not in config")
	// }

	tenant := config.Azure.MultiTenantAuth.Tenants[tenantName]

	var tokenData *TokenData
	// fmt.Println("Getting token for " + tenant.TenantName)
	if options.GetWriteToken {
		authOpts := tenant.Writer
		if options.AzureGraph {
			authOpts.AzureGraph = true
		}
		tokenData, err = GetServicePrincipalToken(tenant.TenantID, authOpts)
		lib.CheckFatalError(err)
	} else if !options.GetWriteToken {
		authOpts := tenant.Reader
		if options.AzureGraph {
			authOpts.AzureGraph = true
		}
		tokenData, err = GetServicePrincipalToken(tenant.TenantID, authOpts)
		lib.CheckFatalError(err)
	}
	mat := MultiAuthToken{
		TenantId:   tenant.TenantID,
		TenantName: tenant.TenantName,
		TokenData:  *tokenData,
	}

	tenantToken = mat

	return &tenantToken, nil
}
