package lib

import (
	"bytes"
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
	"sync"
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

	CheckFatalError(ensureSecretDir(configPath))
	CheckFatalError(ensureSecretDir(cachePath))

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

	if err != nil || fileStat.Size() == 0 {
		config := CldConfigRoot{
			Azure: &AzureConfig{},
		}
		config.Azure.MultiTenantAuth.Tenants = make(CldConfigTenants)

		jsonBytes, _ := json.MarshalIndent(config, "", "  ")
		err := WriteSecretFile(configFile, jsonBytes)
		CheckFatalError(err)
	}

	return configFile, configPath, cachePath
}

func SaveCldConfig(configFilePath string, config CldConfigRoot, options *CldConfigOptions) {
	if configFilePath == "azureAppConfig" {
		jsonBytes, _ := json.MarshalIndent(config, "", "  ")
		err := WriteSecretFile(configFilePath, jsonBytes)
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
		CheckFatalError(WriteSecretFile(configFilePath, encConfig))
	} else {
		jsonBytes, _ := json.MarshalIndent(config, "", "  ")
		err := WriteSecretFile(configFilePath, jsonBytes)
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

	CheckFatalError(WriteSecretFile(configFile, encryptedConfig))
	if removeUnencryptedFile {
		os.Remove(unencryptedFile)
	}
}

func EncryptDecryptedConfigFile(configFilePath string, outputFileName string) {
	config, err := os.ReadFile(configFilePath)
	CheckFatalError(err)

	encryptedConfig, err := Encrypt(config, nil)
	CheckFatalError(err)

	CheckFatalError(WriteSecretFile(outputFileName, encryptedConfig))
}

func DecryptEncryptedConfigFile(configFilePath string, outputFileName string) {
	if configFilePath == "azureAppConfig" {
		config := GetCldConfig(nil)
		jsonBytes, _ := json.MarshalIndent(config, "", "  ")
		err := WriteSecretFile(outputFileName, jsonBytes)
		CheckFatalError(err)
		return
	}
	config, err := os.ReadFile(configFilePath)
	CheckFatalError(err)

	decryptedConfig, err := Decrypt(config, nil)
	CheckFatalError(err)

	CheckFatalError(WriteSecretFile(outputFileName, decryptedConfig))
}

//
//

func DecryptEncryptedTokenCache(tokenCachePath string, outputFileName string) {
	config, err := os.ReadFile(tokenCachePath)
	CheckFatalError(err)

	decryptedConfig, err := Decrypt(config, nil)
	CheckFatalError(err)

	CheckFatalError(WriteSecretFile(outputFileName, decryptedConfig))
	fmt.Println("Saved to", outputFileName)
}

func GetCldConfig(options *CldConfigOptions) CldConfigRoot {
	var (
		config CldConfigRoot
	)
	encryptedConfig, _ := CheckConfigEncryptionOption()

	azAppConfigUrl := os.Getenv("AZURE_APPCONFIG_ENDPOINT")
	if azAppConfigUrl != "" {
		useAzCliAuth := false
		if options != nil && options.UseAzCliAuth {
			fmt.Println("UseAzCliAuth")
		}
		// useAzCliAuth := options.UseAzCliAuth || false
		// JsonMarshalAndPrint(useAzCliAuth)
		// fmt.Println("PRINTING")
		return getAzureAppConfigData(useAzCliAuth)
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

// legacyCipherIV is the fixed IV used by the original AES-CFB format. It is
// kept only so that files written by older versions can still be read.
var legacyCipherIV = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// encryptedV2Prefix marks data encrypted with AES-GCM and a random nonce. The
// legacy format is plain base64, which can never contain ":".
const encryptedV2Prefix = "cld2:"

// Encrypt encrypts text with AES-GCM using a random nonce, so identical
// plaintexts never produce the same ciphertext and any tampering or use of the
// wrong key is detected on decryption.
func Encrypt(text []byte, opts *CldConfigOptions) ([]byte, error) {
	return encryptWithKey([]byte(GetCipherKey(opts)), text)
}

// Decrypt decrypts data produced by Encrypt. Data written by older versions
// (AES-CFB with a fixed IV) is still accepted and is upgraded to the current
// format the next time it is saved.
func Decrypt(text []byte, opts *CldConfigOptions) ([]byte, error) {
	return decryptWithKey([]byte(GetCipherKey(opts)), text)
}

func encryptWithKey(key []byte, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	sealed := gcm.Seal(nonce, nonce, text, nil)
	return append([]byte(encryptedV2Prefix), Encode(sealed)...), nil
}

func decryptWithKey(key []byte, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	text = bytes.TrimSpace(text)

	if bytes.HasPrefix(text, []byte(encryptedV2Prefix)) {
		sealed, err := Decode(text[len(encryptedV2Prefix):])
		if err != nil {
			return nil, err
		}
		gcm, err := cipher.NewGCM(block)
		if err != nil {
			return nil, err
		}
		if len(sealed) < gcm.NonceSize() {
			return nil, fmt.Errorf("encrypted data is truncated")
		}
		nonce, cipherText := sealed[:gcm.NonceSize()], sealed[gcm.NonceSize():]
		plainText, err := gcm.Open(nil, nonce, cipherText, nil)
		if err != nil {
			return nil, fmt.Errorf("unable to decrypt: wrong key or corrupted data")
		}
		return plainText, nil
	}

	// Legacy format: unauthenticated AES-CFB with a fixed IV.
	cipherText, err := Decode(text)
	if err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBDecrypter(block, legacyCipherIV)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return plainText, nil
}

func GetCipherKey(opts *CldConfigOptions) string {
	_, _, cachePath := InitConfig(opts)

	_, configSecret := CheckConfigEncryptionOption()
	if configSecret != "" {
		return configSecret
	}

	keyFile := cachePath + "/key"
	if key, err := os.ReadFile(keyFile); err == nil && len(key) > 0 {
		return string(key)
	}

	key, err := GenerateRandomString(32, true, true, true)
	CheckFatalError(err)
	CheckFatalError(WriteSecretFile(keyFile, []byte(key)))
	return key
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

// tokenCacheMu serialises read-modify-write cycles on the token cache file so
// concurrent tenant goroutines cannot overwrite each other's entries.
var tokenCacheMu sync.Mutex

// readTokenCache returns the decrypted token cache. A missing, unreadable or
// undecryptable cache is treated as empty rather than as a fatal error.
func readTokenCache(cacheFile string, opts *CldConfigOptions) TokenCache {
	cacheFileData := TokenCache{}

	cacheData, err := os.ReadFile(cacheFile)
	if err != nil || len(cacheData) == 0 {
		return cacheFileData
	}
	decrypted, err := Decrypt(cacheData, opts)
	if err != nil {
		return TokenCache{}
	}
	if err := json.Unmarshal(decrypted, &cacheFileData); err != nil || cacheFileData == nil {
		return TokenCache{}
	}
	return cacheFileData
}

func CacheSaveToken[T AzureMultiAuthToken | AzureTokenData | CitrixTokenData](tokenData T, tokenType string, opts *CldConfigOptions) error {
	_, _, cachePath := InitConfig(opts)

	cacheFile := cachePath + "/" + "tkn"

	tokenDataStr, err := json.Marshal(tokenData)
	if err != nil {
		return err
	}

	tokenCacheMu.Lock()
	defer tokenCacheMu.Unlock()

	cacheFileData := readTokenCache(cacheFile, opts)
	cacheFileData[tokenType] = string(Encode(tokenDataStr))

	jsonStr, err := json.Marshal(cacheFileData)
	if err != nil {
		return err
	}

	encryptedCacheData, err := Encrypt(jsonStr, opts)
	if err != nil {
		return err
	}

	return WriteSecretFile(cacheFile, encryptedCacheData)
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

// GetCachedToken returns the cached token stored under tokenType. It returns
// nil when no token cache exists and a zero value token (which callers treat
// as expired) when the cache has no usable entry for tokenType.
func GetCachedToken[T AzureMultiAuthToken | AzureTokenData | CitrixTokenData](tokenType string, opts *CldConfigOptions) *T {
	_, _, cachePath := InitConfig(opts)
	cacheFile := cachePath + "/" + "tkn"

	if _, err := os.Stat(cacheFile); err != nil {
		return nil
	}

	tokenCacheMu.Lock()
	cacheFileData := readTokenCache(cacheFile, opts)
	tokenCacheMu.Unlock()

	var token T

	encodedTokenStr := cacheFileData[tokenType]
	if encodedTokenStr == "" {
		return &token
	}

	decodedTokenStr, err := Decode([]byte(encodedTokenStr))
	if err != nil {
		return &token
	}
	if err := json.Unmarshal(decodedTokenStr, &token); err != nil {
		var empty T
		return &empty
	}
	return &token
}

// tokenExpirySkew is how long before its real expiry a cached token is treated
// as expired, so a token cannot lapse part way through a long-running command.
const tokenExpirySkew = 2 * time.Minute

func CheckCachedTokenExpired(expiry time.Time) bool {
	return !time.Now().Add(tokenExpirySkew).Before(expiry)
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

func getAzureAppConfigData(useAzCliAuth bool) CldConfigRoot {
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
