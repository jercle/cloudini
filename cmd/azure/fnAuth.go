package azure

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/charmbracelet/log"
	"github.com/jercle/cloudini/lib"
)

// var usrHomeDir, err = os.UserHomeDir()

// func GetCachedTokens(cldConfOpts *lib.CldConfigOptions) lib.AllTenantTokens {
// 	var tokens lib.AllTenantTokens
// 	_, _, cachePath := lib.InitConfig(cldConfOpts)
// 	cacheFile := cachePath + "/azTok"

// 	// if _, err := os.Stat(lib.ConfigPath); err != nil {
// 	// 	os.MkdirAll(lib.ConfigPath, os.ModePerm)
// 	// }
// 	// if _, err := os.Stat(cachePath); err != nil {
// 	// 	os.Create(cachePath)
// 	// 	return lib.AllTenantTokens{}
// 	// }
// 	fileData, err := os.ReadFile(cacheFile)
// 	lib.CheckFatalError(err)
// 	byteData, err := b64.StdEncoding.DecodeString(string(fileData))
// 	lib.CheckFatalError(err)
// 	json.Unmarshal(byteData, &tokens)
// 	if len(tokens) == 0 {
// 		fmt.Println("Fetching new tokens")
// 		tokens, err = GetAllTenantSPTokens(lib.AzureMultiAuthTokenRequestOptions{}, cldConfOpts)
// 		lib.CheckFatalError(err)
// 	}
// 	// fmt.Println(tokens)
// 	return tokens
// }

func GetServicePrincipalToken(tenantId string, matOptions lib.AzureMultiAuthTokenRequestOptions, cldConfigOpts *lib.CldConfigOptions, mut *sync.Mutex) (*lib.AzureTokenData, error) {
	ctx := context.Background()
	options := matOptions

	// lib.JsonMarshalAndPrint(matOptions)
	// os.Exit(0)

	var tokenRequestOptions policy.TokenRequestOptions
	var cachedToken *lib.AzureTokenData

	if options.Scope == "" {
		options.Scope = "default"
	}

	if mut != nil {
		mut.Lock()
	}

	readOrWrite := "read"

	if options.GetWriteToken {
		readOrWrite = "write"
	}

	if !options.NoCache {
		cachedToken = lib.GetCachedToken[lib.AzureTokenData]("az"+strings.ToLower(options.TenantName)+options.Scope+readOrWrite, cldConfigOpts)
		if mut != nil {
			mut.Unlock()
		}

		if cachedToken != nil {
			isExpired := lib.CheckCachedTokenExpired(cachedToken.ExpiresOn)
			if !isExpired {
				return cachedToken, nil
			}
		}
	}

	// jsonBytes, _ := json.MarshalIndent(options, "", "  ")
	// fmt.Println(string(jsonBytes))

	switch options.Scope {
	case "defender":
		tokenRequestOptions.Scopes = []string{"https://api.securitycenter.microsoft.com/.default"}
	case "graph":
		tokenRequestOptions.Scopes = []string{"https://graph.microsoft.com/.default"}
	case "storage":
		tokenRequestOptions.Scopes = []string{"https://storage.azure.com/.default"}
	case "monitor":
		tokenRequestOptions.Scopes = []string{"https://monitor.azure.com/.default"}
	case "loganalytics":
		tokenRequestOptions.Scopes = []string{"https://api.loganalytics.io/.default"}
	case "keyvault":
		tokenRequestOptions.Scopes = []string{"cfa8b339-82a2-471a-a3c9-0fc0be7a4093/.default"}
	// case "acr":
	// tokenRequestOptions.Scopes = []string{}

	// encodedData := b64.StdEncoding.EncodeToString([]byte(options.ClientID + ":" + options.ClientSecret))
	// urlString := "https://" +
	// 	options.AzureContainerRepositoryName +
	// 	".azurecr.io/oauth2/token?service=" +
	// 	options.AzureContainerRepositoryName +
	// 	".azurecr.io&scope=repository:*:*"
	// req, err := http.NewRequest(http.MethodGet, urlString, nil)
	// lib.CheckFatalError(err)

	// req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Authorization", "Basic "+encodedData)

	// res, err := http.DefaultClient.Do(req)
	// lib.CheckFatalError(err)

	// responseBody, err := io.ReadAll(res.Body)
	// lib.CheckFatalError(err)
	// defer res.Body.Close()

	// var token lib.AcrAccessToken
	// json.Unmarshal(responseBody, &token)

	// tokenData := lib.AzureTokenData{
	// 	Token: token.AccessToken,
	// }

	// if !options.NoCache {
	// 	if mut != nil {
	// 		mut.Lock()
	// 	}
	// 	lib.CacheSaveToken(tokenData, "az"+strings.ToLower(options.TenantName)+options.Scope, cldConfigOpts)
	// 	if mut != nil {
	// 		mut.Unlock()
	// 	}
	// }
	// return &tokenData, nil

	default:
		tokenRequestOptions.Scopes = []string{"https://management.core.windows.net/.default"}
	}
	tokenRequestOptions.EnableCAE = true

	var tokenResponse azcore.AccessToken
	if strings.HasPrefix(options.ClientSecret, "certPath:") {
		opts := azidentity.ClientCertificateCredentialOptions{
			SendCertificateChain: true,
		}
		// tokenRequestOptions.Claims = "CN=automon-automation"
		// certSplit := strings.Split(options.ClientSecret, ":")
		// certPath := certSplit[1]
		// certPwd := certSplit[2]
		// certData, err := os.ReadFile(certPath)
		// lib.CheckFatalError(err)
		// cert, key, err := azidentity.ParseCertificates(certData, []byte(certPwd))
		// lib.CheckFatalError(err)
		keyFile, err := os.ReadFile("/home/jercle/.config/cld/key.pem")
		lib.CheckFatalError(err)
		block, _ := pem.Decode(keyFile)
		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		lib.CheckFatalError(err)
		certFile, err := os.ReadFile("/home/jercle/.config/cld/cert.pem")
		certBlock, _ := pem.Decode(certFile)
		certData, err := x509.ParseCertificate(certBlock.Bytes)
		cert := []*x509.Certificate{certData}
		lib.CheckFatalError(err)
		cred, err := azidentity.NewClientCertificateCredential(tenantId, options.ClientID, cert, privateKey, &opts)
		lib.CheckFatalError(err)
		tokenResponse, err = cred.GetToken(ctx, tokenRequestOptions)
		lib.CheckFatalError(err)
	} else {
		cred, err := azidentity.NewClientSecretCredential(tenantId, options.ClientID, options.ClientSecret, nil)
		lib.CheckFatalError(err)
		tokenResponse, err = cred.GetToken(ctx, tokenRequestOptions)
		lib.CheckFatalError(err)
	}

	token := lib.AzureTokenData{
		Token:     tokenResponse.Token,
		ExpiresOn: tokenResponse.ExpiresOn,
	}

	// fmt.Println(token)
	if !options.NoCache {
		if mut != nil {
			mut.Lock()
		}
		lib.CacheSaveToken(token, "az"+strings.ToLower(options.TenantName)+options.Scope, cldConfigOpts)
		if mut != nil {
			mut.Unlock()
		}
	}

	return &token, nil
}

func GetServicePrincipalMultiAuthToken(spDetails lib.AzureMultiAuthTokenRequestOptions) (*lib.AzureMultiAuthToken, error) {
	ctx := context.Background()
	var tokenRequestOptions policy.TokenRequestOptions

	switch spDetails.Scope {
	case "graph":
		tokenRequestOptions.Scopes = []string{"https://graph.microsoft.com/.default"}
	case "storage":
		tokenRequestOptions.Scopes = []string{"https://storage.azure.com/.default"}
	case "monitor":
		tokenRequestOptions.Scopes = []string{"https://monitor.azure.com//.default"}
	default:
		tokenRequestOptions.Scopes = []string{"https://management.core.windows.net/.default"}
	}
	tokenRequestOptions.EnableCAE = true

	cred, err := azidentity.NewClientSecretCredential(spDetails.TenantID, spDetails.ClientID, spDetails.ClientSecret, nil)
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

	token := lib.AzureMultiAuthToken{
		TenantId:   spDetails.TenantID,
		TenantName: spDetails.TenantName,
		TokenData: lib.AzureTokenData{
			Token:     tokenResponse.Token,
			ExpiresOn: tokenResponse.ExpiresOn,
		},
	}

	// fmt.Println(token)
	return &token, nil
}

func GetAzCliToken() lib.AzureMultiAuthToken {
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

	token := lib.AzureMultiAuthToken{
		// tokenre
		TokenData: lib.AzureTokenData{
			Token:     tokenResponse.Token,
			ExpiresOn: tokenResponse.ExpiresOn,
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

func GetToken(tenantName string, cldConfigOpts *lib.CldConfigOptions) lib.AzureMultiAuthToken {
	cachedToken := lib.GetCachedToken[lib.AzureMultiAuthToken]("azMA-"+tenantName, cldConfigOpts)
	// fmt.Println(cachedToken)
	isExpired := lib.CheckCachedTokenExpired(cachedToken.TokenData.ExpiresOn)
	if !isExpired {
		return *cachedToken
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	lib.CheckFatalError(err)

	sub, err := GetActiveCliSub()
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

	multiAuthToken := lib.AzureMultiAuthToken{
		TenantId:   sub.TenantID,
		TenantName: tenantName,
		TokenData: lib.AzureTokenData{
			Token:     tokenResponse.Token,
			ExpiresOn: tokenResponse.ExpiresOn,
		},
	}

	return multiAuthToken
}

// Gets a token for each tenant configured in the cld config file
//
// Default path for config file is ~/.config/cld/cldConfig.json
//
// First parameter passed into this function overwrites the config file path
func GetAllTenantSPTokens(options lib.AzureMultiAuthTokenRequestOptions, cldConfOpts *lib.CldConfigOptions) (lib.AllTenantTokens, error) {
	var (
		config       lib.CldConfigRoot
		tenantTokens []lib.AzureMultiAuthToken
		wg           sync.WaitGroup
		mut          sync.Mutex
	)

	config = lib.GetCldConfig(cldConfOpts)

	azConfig := *config.Azure
	tenants := azConfig.MultiTenantAuth.Tenants

	for _, tenant := range tenants {
		wg.Add(1)
		go func() {
			var (
				writerConfig lib.CldConfigClientAuthDetails
				readerConfig lib.CldConfigClientAuthDetails
			)
			configExists := false
			options.TenantName = tenant.TenantName
			opts := options

			switch options.GetWriteToken {
			case true:
				if tenant.Writer != nil {
					writerConfig = *tenant.Writer
					opts.ClientID = writerConfig.ClientID
					opts.ClientSecret = writerConfig.ClientSecret
					configExists = true
				}
			default:
				if tenant.Reader != nil {
					readerConfig = *tenant.Reader
					opts.ClientID = readerConfig.ClientID
					opts.ClientSecret = readerConfig.ClientSecret
					configExists = true
				}
			}

			if configExists {
				tokenData, err := GetServicePrincipalToken(tenant.TenantID, opts, cldConfOpts, &mut)
				lib.CheckFatalError(err)

				tenantToken := lib.AzureMultiAuthToken{
					TenantId:   tenant.TenantID,
					ClientId:   opts.ClientID,
					TenantName: tenant.TenantName,
					TokenData:  *tokenData,
				}
				mut.Lock()
				tenantTokens = append(tenantTokens, tenantToken)
				mut.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return tenantTokens, nil
}

func GetTenantSPToken(options lib.AzureMultiAuthTokenRequestOptions, cldConfOpts *lib.CldConfigOptions) (*lib.AzureMultiAuthToken, error) {
	var (
		config      lib.CldConfigRoot
		tenantToken lib.AzureMultiAuthToken
		tenant      lib.CldConfigTenantAuth
	)

	config = lib.GetCldConfig(cldConfOpts)

	if options.TenantName == "" {
		tn, err := config.Azure.GetDefaultTenant()
		lib.CheckFatalError(err)
		tenant = *tn
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

	tokenData, err := GetServicePrincipalToken(tenant.TenantID, options, cldConfOpts, nil)
	lib.CheckFatalError(err)

	mat := lib.AzureMultiAuthToken{
		TenantId:   tenant.TenantID,
		TenantName: tenant.TenantName,
		TokenData:  *tokenData,
	}

	tenantToken = mat

	return &tenantToken, nil
}

func GetTenantAzCred(tenantName string, getWriteToken bool, cldConfOpts *lib.CldConfigOptions) (*azidentity.ClientSecretCredential, error) {
	var (
		cred *azidentity.ClientSecretCredential
		err  error
	)
	config := lib.GetCldConfig(cldConfOpts)
	tenant := config.Azure.MultiTenantAuth.Tenants[tenantName]

	if getWriteToken {
		cred, err = azidentity.NewClientSecretCredential(tenant.TenantID, tenant.Writer.ClientID, tenant.Writer.ClientSecret, nil)
		// lib.CheckFatalError(err)
		if err != nil {
			return nil, err
		}
	} else {
		cred, err = azidentity.NewClientSecretCredential(tenant.TenantID, tenant.Reader.ClientID, tenant.Reader.ClientSecret, nil)
		// lib.CheckFatalError(err)
		if err != nil {
			return nil, err
		}
	}

	return cred, nil
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
