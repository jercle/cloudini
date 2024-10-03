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
	"github.com/jercle/cloudini/lib"
)

// var usrHomeDir, err = os.UserHomeDir()

func GetCachedTokens(cldConfOpts *lib.CldConfigOptions) lib.AllTenantTokens {
	var tokens lib.AllTenantTokens
	if _, err := os.Stat(lib.ConfigPath); err != nil {
		os.MkdirAll(lib.ConfigPath, os.ModePerm)
	}
	if _, err := os.Stat(lib.TokenCacheFile); err != nil {
		os.Create(lib.TokenCacheFile)
	}
	fileData, err := os.ReadFile(lib.TokenCacheFile)
	lib.CheckFatalError(err)
	byteData, err := b64.StdEncoding.DecodeString(string(fileData))
	lib.CheckFatalError(err)
	json.Unmarshal(byteData, &tokens)
	if len(tokens) == 0 {
		fmt.Println("Fetching new tokens")
		tokens, err = GetAllTenantSPTokens(lib.MultiAuthTokenRequestOptions{}, cldConfOpts)
		lib.CheckFatalError(err)
	}
	// fmt.Println(tokens)
	return tokens
}

func GetServicePrincipalToken(tenant string, spDetails lib.MultiAuthTokenRequestOptions) (*lib.TokenData, error) {
	ctx := context.Background()
	var tokenRequestOptions policy.TokenRequestOptions

	// jsonBytes, _ := json.MarshalIndent(spDetails, "", "  ")
	// fmt.Println(string(jsonBytes))

	switch spDetails.Scope {
	case "graph":
		tokenRequestOptions.Scopes = []string{"https://graph.microsoft.com/.default"}
	case "storage":
		tokenRequestOptions.Scopes = []string{"https://storage.azure.com/.default"}
	case "acr":
		tokenRequestOptions.Scopes = []string{}
		encodedData := b64.StdEncoding.EncodeToString([]byte(spDetails.ClientID + ":" + spDetails.ClientSecret))
		urlString := "https://" +
			spDetails.AzureContainerRepositoryName +
			".azurecr.io/oauth2/token?service=" +
			spDetails.AzureContainerRepositoryName +
			".azurecr.io&scope=repository:*:*"
		req, err := http.NewRequest(http.MethodGet, urlString, nil)
		lib.CheckFatalError(err)

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Basic "+encodedData)

		res, err := http.DefaultClient.Do(req)
		lib.CheckFatalError(err)

		responseBody, err := io.ReadAll(res.Body)
		lib.CheckFatalError(err)
		defer res.Body.Close()

		var token lib.AcrAccessToken
		json.Unmarshal(responseBody, &token)

		jsonBytes, _ := json.MarshalIndent(token, "", "  ")
		fmt.Println(string(jsonBytes))

		tokenData := lib.TokenData{
			Token: token.AccessToken,
		}

		return &tokenData, nil

	default:
		tokenRequestOptions.Scopes = []string{"https://management.core.windows.net/.default"}
	}
	tokenRequestOptions.EnableCAE = true

	cred, err := azidentity.NewClientSecretCredential(tenant, spDetails.ClientID, spDetails.ClientSecret, nil)
	if err != nil {
		// log.Error("Unable to obtain Azure token", err, err)
		lib.CheckFatalError(err)
		return nil, err
	}
	// envCred, err := azidentity.NewEnvironmentCredential(nil)
	// if err != nil {
	// 	log.Error("Unable to obtain Azure token", err, err)
	// }

	tokenResponse, err := cred.GetToken(ctx, tokenRequestOptions)
	if err != nil {
		// log.Error("Unable to obtain Azure token", err, err)
		lib.CheckFatalError(err)
		return nil, err
	}

	token := lib.TokenData{
		Token:     tokenResponse.Token,
		ExpiresOn: tokenResponse.ExpiresOn.Local().String(),
	}

	// fmt.Println(token)
	return &token, nil
}

func GetServicePrincipalMultiAuthToken(tenantId string, spDetails lib.MultiAuthTokenRequestOptions) (*lib.MultiAuthToken, error) {
	ctx := context.Background()
	var tokenRequestOptions policy.TokenRequestOptions

	switch spDetails.Scope {
	case "graph":
		tokenRequestOptions.Scopes = []string{"https://graph.microsoft.com/.default"}
	// case "apc-sharepoint":
	// 	tokenRequestOptions.Scopes = []string{"https://asiogovau.sharepoint.com/.default"}
	case "storage":
		tokenRequestOptions.Scopes = []string{"https://storage.azure.com/.default"}
	default:
		tokenRequestOptions.Scopes = []string{"https://management.core.windows.net/.default"}
	}
	tokenRequestOptions.EnableCAE = true

	cred, err := azidentity.NewClientSecretCredential(tenantId, spDetails.ClientID, spDetails.ClientSecret, nil)
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

	token := lib.MultiAuthToken{
		TenantId: tenantId,
		TokenData: lib.TokenData{
			Token:     tokenResponse.Token,
			ExpiresOn: tokenResponse.ExpiresOn.Local().String(),
		},
	}

	// fmt.Println(token)
	return &token, nil
}

func GetAzCliToken() lib.MultiAuthToken {
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

	token := lib.MultiAuthToken{
		// tokenre
		TokenData: lib.TokenData{
			Token:     tokenResponse.Token,
			ExpiresOn: tokenResponse.ExpiresOn.Local().String(),
		},

		// Token:     tokenResponse.Token,
		// ExpiresOn: tokenResponse.ExpiresOn.Local().String(),
	}

	// fmt.Println(token)
	return token
}

func GetLogAnalyticsToken() (*lib.TokenRequestResponse, error) {
	var (
		authDetails         lib.AzureAuthDetails
		authRequestResponse *lib.TokenRequestResponse
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

func GetAzureEnvVariables(requiredEnvVars lib.AzureAuthRequirements) *lib.AzureAuthDetails {
	envs := lib.AzureAuthDetails{
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

func GetToken(tenantName string) lib.MultiAuthToken {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	lib.CheckFatalError(err)

	sub, err := GetActiveSub()
	lib.CheckFatalError(err)
	// fmt.Println(sub.TenantID)

	ctx := context.Background()
	tokenRequestOptions := policy.TokenRequestOptions{
		Scopes: []string{
			"https://management.core.windows.net/.default",
		},
	}

	tokenResponse, err := cred.GetToken(ctx, tokenRequestOptions)
	lib.CheckFatalError(err)

	multiAuthToken := lib.MultiAuthToken{
		TenantId:   sub.TenantID,
		TenantName: tenantName,
		TokenData: lib.TokenData{
			Token:     tokenResponse.Token,
			ExpiresOn: tokenResponse.ExpiresOn.Local().String(),
		},
	}

	return multiAuthToken
}

// Gets a token for each tenant configured in the cld config file
//
// Default path for config file is ~/.config/cld/cldConfig.json
//
// First parameter passed into this function overwrites the config file path
func GetAllTenantSPTokens(options lib.MultiAuthTokenRequestOptions, cldConfOpts *lib.CldConfigOptions) (lib.AllTenantTokens, error) {
	var (
		config       lib.CldConfigRoot
		tenantTokens []lib.MultiAuthToken
		wg           sync.WaitGroup
		mut          sync.Mutex
	)

	config = lib.GetCldConfig(cldConfOpts)

	for _, tenant := range config.Azure.MultiTenantAuth.Tenants {
		wg.Add(1)
		go func() {

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

			tokenData, err := GetServicePrincipalToken(tenant.TenantID, options)
			lib.CheckFatalError(err)

			tenantToken := lib.MultiAuthToken{
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

func GetTenantSPToken(options lib.MultiAuthTokenRequestOptions, cldConfOpts *lib.CldConfigOptions) (*lib.MultiAuthToken, error) {
	var (
		config      lib.CldConfigRoot
		tenantToken lib.MultiAuthToken
		tenant      lib.CldConfigTenantAuth
	)

	config = lib.GetCldConfig(cldConfOpts)

	if options.TenantName == "" {
		tenant = config.Azure.GetDefaultTenant()
	} else {
		t, tenantExists := config.Azure.MultiTenantAuth.Tenants[options.TenantName]
		if !tenantExists {
			return nil, fmt.Errorf("Tenant not found in config")
		}
		tenant = t
	}

	switch options.GetWriteToken {
	case true:
		options.ClientID = tenant.Writer.ClientID
		options.ClientSecret = tenant.Writer.ClientSecret
	default:
		options.ClientID = tenant.Reader.ClientID
		options.ClientSecret = tenant.Reader.ClientSecret
	}

	tokenData, err := GetServicePrincipalToken(tenant.TenantID, options)
	lib.CheckFatalError(err)

	mat := lib.MultiAuthToken{
		TenantId:   tenant.TenantID,
		TenantName: tenant.TenantName,
		TokenData:  *tokenData,
	}

	tenantToken = mat

	return &tenantToken, nil
}

func GetTenantAzCred(tenantName string, getWriteToken bool, cldConfOpts *lib.CldConfigOptions) *azidentity.ClientSecretCredential {
	var (
		cred *azidentity.ClientSecretCredential
		err  error
	)
	config := lib.GetCldConfig(cldConfOpts)
	tenant := config.Azure.MultiTenantAuth.Tenants[tenantName]

	if getWriteToken {
		cred, err = azidentity.NewClientSecretCredential(tenant.TenantID, tenant.Writer.ClientID, tenant.Writer.ClientSecret, nil)
		lib.CheckFatalError(err)
	} else {
		cred, err = azidentity.NewClientSecretCredential(tenant.TenantID, tenant.Reader.ClientID, tenant.Reader.ClientSecret, nil)
		lib.CheckFatalError(err)
	}

	return cred
}

// envUpp := strings.ToUpper(env)
// tenantID := os.Getenv(envUpp + "-TENANT-ID")
// clientID := os.Getenv(envUpp + "-CLIENT-ID-AUTOMON-READ")
// serviceConnectionID := os.Getenv(envUpp + "-SERVICE-CONN-ID")
// systemAccessToken := os.Getenv("SYSTEM_ACCESSTOKEN")

type GetAzureDevOpsPipelineTokensOptionsTenantConfig struct {
	TenantName          string
	TenantID            string
	ClientID            string
	ServiceConnectionID string
}
