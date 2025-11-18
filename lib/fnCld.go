package lib

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"slices"
	"sort"
	"strconv"
	"time"

	"github.com/Azure/AppConfiguration-GoProvider/azureappconfiguration"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// usrHomeDir, err := os.UserHomeDir()
// var ConfigPath, _, _ = InitConfig(nil)

func InitConfig(options *CldConfigOptions) (configFile string, configPath string, cachePath string) {

	// if options != nil && options.EncryptUnencryptedFile {
	// 	encryptConfig, configSecret := CheckConfigEncryptionOption()

	// }

	usrHomeDir, err := os.UserHomeDir()
	CheckFatalError(err)
	configPath = usrHomeDir + "/.config/cld"
	configFilePath := configPath + "/cldConf.json"
	cachePath = configPath + "/cache"

	if _, err := os.Stat(configPath); err != nil {
		os.MkdirAll(configPath, os.ModePerm)
	}
	if _, err := os.Stat(cachePath); err != nil {
		os.MkdirAll(cachePath, os.ModePerm)
	}

	if os.Getenv("AZURE_APPCONFIG_ENDPOINT") != "" {
		// useAzAppConfig = true
		// azAppConfigTenantId = os.Getenv("AZURE_APPCONFIG_TENANT_ID")
		// azAppConfigClientId = os.Getenv("AZURE_APPCONFIG_CLIENT_ID")
		// azAppConfigClientSecret = os.Getenv("AZURE_APPCONFIG_CLIENT_SECRET")
		return "azureAppConfig", "azureAppConfig", cachePath
	}

	CLD_CONFIG_PATH := os.Getenv("CLD_CONFIG_PATH")

	if options != nil && options.ConfigFile != "" {
		configFile = options.ConfigFile
	} else if CLD_CONFIG_PATH != "" {
		configFile = CLD_CONFIG_PATH
	} else {
		configFile = configFilePath
	}

	fileStat, err := os.Stat(configFile)

	if err != nil {
		config := CldConfigRoot{}
		config.Azure.MultiTenantAuth.Tenants = make(CldConfigTenants)

		jsonBytes, _ := json.MarshalIndent(config, "", "  ")
		err := os.WriteFile(configFilePath, jsonBytes, os.ModePerm)
		CheckFatalError(err)
	}

	if fileStat.Size() == 0 {
		config := CldConfigRoot{}
		config.Azure.MultiTenantAuth.Tenants = make(CldConfigTenants)

		jsonBytes, _ := json.MarshalIndent(config, "", "  ")
		err := os.WriteFile(configFilePath, jsonBytes, os.ModePerm)
		CheckFatalError(err)
	}

	return configFile, configPath, cachePath
}

func SaveCldConfig(configFilePath string, config CldConfigRoot, options *CldConfigOptions) {
	if configFilePath == "azureAppConfig" {
		jsonBytes, _ := json.MarshalIndent(config, "", "  ")
		err := os.WriteFile(configFilePath, jsonBytes, os.ModePerm)
		CheckFatalError(err)
		return
	}
	encryptConfig, _ := CheckConfigEncryptionOption()
	_ = encryptConfig

	// jsonBytes, _ := json.Marshal(config)
	// fmt.Println(string(jsonBytes))

	// os.Exit(0)
	if encryptConfig {
		jsonBytes, _ := json.Marshal(config)
		encConfig, err := Encrypt(jsonBytes, options)
		CheckFatalError(err)
		os.WriteFile(configFilePath, encConfig, os.ModePerm)
	} else {
		jsonBytes, _ := json.MarshalIndent(config, "", "  ")
		err := os.WriteFile(configFilePath, jsonBytes, os.ModePerm)
		CheckFatalError(err)
	}
}

func EncryptUnencryptedConfigFile(unencryptedFile string, removeUnencryptedFile bool) {
	configFile, _, _ := InitConfig(nil)

	config, err := os.ReadFile(unencryptedFile)
	CheckFatalError(err)

	isValidJson := IsValidJson(string(config))
	if !isValidJson {
		err := fmt.Errorf("invalid json file provided")
		CheckFatalError(err)
	}

	encryptedConfig, err := Encrypt(config, nil)
	CheckFatalError(err)

	os.WriteFile(configFile, encryptedConfig, os.ModePerm)
	if removeUnencryptedFile {
		os.Remove(unencryptedFile)
	}
}

func EncryptDecryptedConfigFile(configFilePath string, outputFileName string) {
	config, err := os.ReadFile(configFilePath)
	CheckFatalError(err)

	encryptedConfig, err := Encrypt(config, nil)
	CheckFatalError(err)

	os.WriteFile(outputFileName, encryptedConfig, os.ModePerm)
}

func DecryptEncryptedConfigFile(configFilePath string, outputFileName string) {
	if configFilePath == "azureAppConfig" {
		config := GetCldConfig(nil)
		jsonBytes, _ := json.MarshalIndent(config, "", "  ")
		err := os.WriteFile(outputFileName, jsonBytes, os.ModePerm)
		CheckFatalError(err)
		return
	}
	config, err := os.ReadFile(configFilePath)
	CheckFatalError(err)

	decryptedConfig, err := Decrypt(config, nil)
	CheckFatalError(err)

	os.WriteFile(outputFileName, decryptedConfig, os.ModePerm)
}

//
//

func DecryptEncryptedTokenCache(tokenCachePath string, outputFileName string) {
	config, err := os.ReadFile(tokenCachePath)
	CheckFatalError(err)

	decryptedConfig, err := Decrypt(config, nil)
	CheckFatalError(err)

	os.WriteFile(outputFileName, decryptedConfig, os.ModePerm)
	fmt.Println("Saved to", outputFileName)
}

func GetCldConfig(options *CldConfigOptions) CldConfigRoot {
	var (
		config CldConfigRoot
	)
	encryptedConfig, _ := CheckConfigEncryptionOption()

	azAppConfigUrl := os.Getenv("AZURE_APPCONFIG_ENDPOINT")
	if azAppConfigUrl != "" {
		return getAzureAppConfigData()
	}

	// fmt.Println(encryptedConfig)
	// os.Exit(0)

	// fmt.Println(options)
	// fmt.Println(encryptedConfig)
	// os.Exit(0)

	configFile, _, _ := InitConfig(options)

	// fmt.Println("Getting config from", configFile)
	byteValue, err := os.ReadFile(configFile)
	CheckFatalError(err)

	// fmt.Println(string(byteValue))
	// os.Exit(0)
	if encryptedConfig {
		decryptedConfig, err := Decrypt(byteValue, options)
		CheckFatalError(err)
		// jsonStr, _ := json.MarshalIndent(decryptedConfig, "", "  ")
		// fmt.Println(string(jsonStr))
		// os.Exit(0)
		err = json.Unmarshal(decryptedConfig, &config)
		CheckFatalError(err)
	} else {
		err = json.Unmarshal(byteValue, &config)
		CheckFatalError(err)
	}

	return config
}

// Possible values for CLD_CONFIG_ENCRYPT environment variables are:
// "1"
// "t"
// "T"
// "TRUE"
// "true"
// "True"
// "0"
// "f"
// "F"
// "FALSE"
// "false"
// "False"

func CheckConfigEncryptionOption() (bool, string) {
	encryptConfig := false
	configSecret := os.Getenv("CLD_CONFIG_SECRET")
	encryptConfigEnvVar, encryptConfigEnvVarExists := os.LookupEnv("CLD_CONFIG_ENCRYPT")
	if encryptConfigEnvVarExists {
		v, err := strconv.ParseBool(encryptConfigEnvVar)
		CheckFatalError(err)
		encryptConfig = v
	}

	if encryptConfig && configSecret == "" {
		fmt.Fprintln(os.Stderr, "CLD_CONFIG_SECRET environment variable required to encrypt and decrypt config file. Must be 32 characters in length")
		fmt.Fprintln(os.Stderr, "IMPORTANT: You must keep this secret safe so you do not lose your config data.")
		os.Exit(1)
	}
	return encryptConfig, configSecret
}

func Encode(b []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(b))
}

func Decode(s []byte) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(string(s))
	return data, err
}

var cipherBytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Encrypt(text []byte, opts *CldConfigOptions) ([]byte, error) {
	secret := GetCipherKey(opts)

	block, err := aes.NewCipher([]byte(secret))
	CheckFatalError(err)

	cfb := cipher.NewCFBEncrypter(block, cipherBytes)
	cipherText := make([]byte, len(text))
	cfb.XORKeyStream(cipherText, text)
	return Encode(cipherText), nil
}

func Decrypt(text []byte, opts *CldConfigOptions) ([]byte, error) {

	// fmt.Println("decrypting")
	secret := GetCipherKey(opts)

	block, err := aes.NewCipher([]byte(secret))
	CheckFatalError(err)
	if err != nil {
		return []byte{}, err
	}
	// fmt.Println("decoding")
	cipherText, err := Decode(text)
	// if err != nil {
	// jsonStr, _ := json.MarshalIndent(secret, "", "  ")
	// fmt.Println(string(text))
	// fmt.Println(err.Error())
	// os.WriteFile("main-test-text", text, 0644)
	// os.Exit(0)
	// return []byte{}, err
	// }
	CheckFatalError(err)

	cfb := cipher.NewCFBDecrypter(block, cipherBytes)

	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return plainText, nil
}

func GetCipherKey(opts *CldConfigOptions) string {
	_, _, cachePath := InitConfig(opts)

	_, configSecret := CheckConfigEncryptionOption()
	if configSecret != "" {
		return configSecret
	} else {
		if _, err := os.Stat(cachePath + "/key"); err != nil {
			key, err := GenerateRandomString(32, true, true, true)
			CheckFatalError(err)
			f, err := os.Create(cachePath + "/key")
			CheckFatalError(err)
			_, err = f.WriteString(key)
			CheckFatalError(err)
			return key
		} else {
			key, err := os.ReadFile(cachePath + "/key")
			CheckFatalError(err)
			return string(key)
		}
	}
}

func GenerateRandomString(n int, includeUpper bool, includeNumbers bool, includeSpecial bool) (string, error) {
	lowChars := "abcdefghijklmnopqrstuvwxyz"
	uppChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numChars := "0123456789"
	spcChars := "-_~`!@#$%^&*()"

	chars := lowChars

	if includeUpper {
		chars += uppChars
	}
	if includeNumbers {
		chars += numChars
	}
	if includeSpecial {
		chars += spcChars
	}

	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		ret[i] = chars[num.Int64()]
	}

	return string(ret), nil
}

func CacheSaveToken[T AzureMultiAuthToken | AzureTokenData | CitrixTokenData](tokenData T, tokenType string, opts *CldConfigOptions) error {
	_, _, cachePath := InitConfig(opts)

	cacheFile := cachePath + "/" + "tkn"

	cacheFileData := TokenCache{}
	tokenDataStr, err := json.Marshal(tokenData)
	tokenDataStrEncoded := Encode(tokenDataStr)
	CheckFatalError(err)

	if _, err := os.Stat(cacheFile); err != nil {
		cacheFileData[tokenType] = string(tokenDataStrEncoded)
		jsonStr, _ := json.MarshalIndent(cacheFileData, "", "  ")
		encCacheStr, err := Encrypt(jsonStr, opts)
		CheckFatalError(err)
		f, err := os.Create(cacheFile)
		CheckFatalError(err)
		_, err = f.WriteString(string(encCacheStr))
		CheckFatalError(err)
	} else {
		cacheData, err := os.ReadFile(cacheFile)
		CheckFatalError(err)
		decrypted, err := Decrypt(cacheData, opts)
		CheckFatalError(err)
		err = json.Unmarshal(decrypted, &cacheFileData)
		if err != nil {
			cacheFileData[tokenType] = string(tokenDataStrEncoded)
			jsonStr, _ := json.MarshalIndent(cacheFileData, "", "  ")
			encCacheStr, err := Encrypt(jsonStr, opts)
			CheckFatalError(err)
			os.WriteFile(cacheFile, encCacheStr, 0644)
			return nil
		}

		cacheFileData[tokenType] = string(tokenDataStrEncoded)

		jsonStr, err := json.Marshal(cacheFileData)
		CheckFatalError(err)

		encryptedCacheData, err := Encrypt(jsonStr, opts)
		CheckFatalError(err)

		os.WriteFile(cacheFile, encryptedCacheData, 0644)
	}
	return nil
}

func ClearTokenCache(opts *CldConfigOptions) {
	_, _, cachePath := InitConfig(opts)

	cacheFile := cachePath + "/" + "tkn"

	// fmt.Println(cacheFile)
	// os.Exit(0)
	err := os.Remove(cacheFile)
	CheckFatalError(err)
	fmt.Println("Deleted cached token data: " + cacheFile)
}

func GetCachedToken[T AzureMultiAuthToken | AzureTokenData | CitrixTokenData](tokenType string, opts *CldConfigOptions) *T {
	_, _, cachePath := InitConfig(opts)
	// cacheFile := cachePath + "/" + tokenType + "Tok"
	cacheFile := cachePath + "/" + "tkn"
	cacheFileData := TokenCache{}

	if _, err := os.Stat(cacheFile); err != nil {
		return nil
	}

	// cachedEncryptedToken, err := os.ReadFile(cacheFile)
	// CheckFatalError(err)
	// // fmt.Println(string(cachedEncryptedToken))

	// cachedDecryptedToken, err := Decrypt(cachedEncryptedToken, opts)
	// CheckFatalError(err)
	// // fmt.Println(string(cachedDecryptedToken))

	// var token T
	// json.Unmarshal(cachedDecryptedToken, &token)

	// return &token

	cacheData, err := os.ReadFile(cacheFile)
	CheckFatalError(err)

	// for i, b := range cacheData {
	// 	if i > 10934 && i < 10938 {
	// 		fmt.Println(i)
	// 		fmt.Println(string(b))
	// 	}
	// }
	// os.Exit(0)
	decrypted, err := Decrypt(cacheData, opts)
	CheckFatalError(err)
	// fmt.Println(string(cacheData))
	// os.Exit(0)

	// os.Exit(0)

	// fmt.Println(string(decrypted))
	// JsonMarshalAndPrint(opts)
	// fmt.Println(tokenType)
	// os.Exit(0)

	err = json.Unmarshal(decrypted, &cacheFileData)
	if err != nil {
		var token T

		return &token
	}

	encodedTokenStr := cacheFileData[tokenType]
	if encodedTokenStr == "" {
		var token T

		return &token
	} else {
		// fmt.Println(encodedTokenStr)
		decodedTokenStr, err := Decode([]byte(encodedTokenStr))
		CheckFatalError(err)
		// fmt.Println(string(decodedTokenStr))
		var token T
		json.Unmarshal(decodedTokenStr, &token)
		// fmt.Println(token)
		// _ = decodedTokenStr
		return &token
	}

	// var token T
	// json.Unmarshal()

	// fmt.Println(cacheFileData)

	// os.Exit(0)
	// return nil
}

// type GenericTokenData interface {
// 	AzureMultiAuthToken | AzureTokenData | CitrixTokenData
// 	// IsValid IsValid
// }

// Returns true if expired
func CheckCachedTokenExpired(expiry time.Time) bool {
	currentTime := time.Now()
	// fmt.Println("Expiry: " + expiry.Format(time.RFC3339))
	// fmt.Println("Current Time: " + currentTime.Format(time.RFC3339))

	if currentTime.After(expiry) || currentTime.Equal(expiry) {
		return true
	}
	return false
}

// func IsValid() bool {
// 	var token CitrixTokenData

// 	return true
// }

func SortMapByKey(mapData map[string]interface{}) {
	keys := make([]string, 0, len(mapData))

	for k := range mapData {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, mapData[k])
	}
}

func MapAzureSubscriptionToCustomTenantName(subscriptionId string, config AzureConfig) string {
	customTenantName := ""
	for name, subIds := range config.CustomSubIdToTenantNameMap {
		if slices.Contains(subIds, subscriptionId) {
			return name
		}
	}
	return customTenantName
}

func IsValidJson(s string) bool {
	var js interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

//
//

func MapTenantIdToConfiguredTenantName(tenantId string, config AzureConfig) (tenantName string) {
	for _, tConf := range config.MultiTenantAuth.Tenants {
		if tConf.TenantID == tenantId {
			tenantName = tConf.TenantName
			continue
		}
	}
	return
}

//
//

func getAzureAppConfigData() CldConfigRoot {
	azAppConfigUrl := os.Getenv("AZURE_APPCONFIG_ENDPOINT")
	azAppConfigTenantId := os.Getenv("AZURE_APPCONFIG_TENANT_ID")
	azAppConfigClientId := os.Getenv("AZURE_APPCONFIG_CLIENT_ID")
	azAppConfigClientSecret := os.Getenv("AZURE_APPCONFIG_CLIENT_SECRET")
	azAppConfigManagedIdentity := os.Getenv("AZURE_APPCONFIG_MANAGED_IDENTITY")

	azAppConfigLabel := os.Getenv("AZURE_APPCONFIG_LABEL")

	clientOptions := policy.ClientOptions{
		Telemetry: policy.TelemetryOptions{
			Disabled: true,
		},
	}
	credOptions := &azidentity.ClientSecretCredentialOptions{
		ClientOptions: clientOptions,
	}

	authOptions := azureappconfiguration.AuthenticationOptions{
		Endpoint: azAppConfigUrl,
		// Credential: credential,
	}

	kvOptions := &azureappconfiguration.KeyVaultOptions{}

	if azAppConfigManagedIdentity != "" {
		mgIdentClientId := azidentity.ClientID(azAppConfigManagedIdentity)
		mgdIdentOpts := azidentity.ManagedIdentityCredentialOptions{
			ID: mgIdentClientId,
		}

		credential, err := azidentity.NewManagedIdentityCredential(&mgdIdentOpts)
		CheckFatalError(err)
		authOptions.Credential = credential
		kvOptions.Credential = credential
	} else {
		if azAppConfigClientSecret == "" || azAppConfigClientId == "" || azAppConfigTenantId == "" {
			fmt.Println("Ensure AZURE_APPCONFIG_TENANT_ID, AZURE_APPCONFIG_CLIENT_ID, and AZURE_APPCONFIG_CLIENT_SECRET are set")
			os.Exit(1)
		}
		credential, err := azidentity.NewClientSecretCredential(azAppConfigTenantId, azAppConfigClientId, azAppConfigClientSecret, credOptions)
		CheckFatalError(err)
		authOptions.Credential = credential
		kvOptions.Credential = credential
	}

	options := &azureappconfiguration.Options{
		KeyVaultOptions: *kvOptions,
		// Selectors: []azureappconfiguration.Selector{
		// 	{
		// 		KeyFilter:   "*",
		// 		LabelFilter: azAppConfigLabel,
		// 	},
		// },
	}

	appConfig, err := azureappconfiguration.Load(context.TODO(), authOptions, options)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	var cldConfig CldConfigRoot

	appConfigBytes, err := appConfig.GetBytes(nil)
	CheckFatalError(err)
	json.Unmarshal(appConfigBytes, &cldConfig)

	if azAppConfigLabel != "" {
		options := &azureappconfiguration.Options{
			KeyVaultOptions: *kvOptions,
			Selectors: []azureappconfiguration.Selector{
				{
					KeyFilter:   "*",
					LabelFilter: azAppConfigLabel,
				},
			},
		}

		appConfig, err := azureappconfiguration.Load(context.TODO(), authOptions, options)
		if err != nil {
			log.Fatalf("Failed to load configuration: %v", err)
		}

		appConfigBytes, err := appConfig.GetBytes(nil)
		CheckFatalError(err)
		json.Unmarshal(appConfigBytes, &cldConfig)
	}

	return cldConfig
}
