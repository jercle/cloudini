package lib

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

const testConfigSecret = "0123456789abcdef0123456789abcdef" // 32 chars, AES-256

// Points the config and cache directories at a temp home directory and clears
// every environment variable that changes how config is located or encrypted.
//
// A config file is created up front, as InitConfig expects one to exist.
func setupTestConfigHome(t *testing.T) (home string, configFile string, cachePath string) {
	t.Helper()
	home = t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)

	for _, name := range []string{"AZURE_APPCONFIG_ENDPOINT", "CLD_CONFIG_PATH", "CLD_CONFIG_ENCRYPT", "CLD_CONFIG_SECRET"} {
		// t.Setenv registers the restore of the original value, Unsetenv then clears it
		t.Setenv(name, "")
		os.Unsetenv(name)
	}

	configPath := filepath.Join(home, ".config", "cld")
	cachePath = filepath.Join(configPath, "cache")
	if err := os.MkdirAll(cachePath, 0755); err != nil {
		t.Fatal(err)
	}
	configFile = filepath.Join(configPath, "cldConf.json")
	if err := os.WriteFile(configFile, []byte(`{"azure":{"multiTenantAuth":{"tenants":{}}}}`), 0644); err != nil {
		t.Fatal(err)
	}
	return home, configFile, cachePath
}

func samePath(t *testing.T, got string, want string) {
	t.Helper()
	if filepath.Clean(got) != filepath.Clean(want) {
		t.Fatalf("got path %q, want %q", got, want)
	}
}

//
// InitConfig
//

func TestInitConfigDefaultLocation(t *testing.T) {
	home, wantConfigFile, wantCachePath := setupTestConfigHome(t)

	configFile, configPath, cachePath := InitConfig(nil)

	samePath(t, configFile, wantConfigFile)
	samePath(t, configPath, filepath.Join(home, ".config", "cld"))
	samePath(t, cachePath, wantCachePath)
}

func TestInitConfigCreatesConfigAndCacheDirectories(t *testing.T) {
	home, _, _ := setupTestConfigHome(t)

	// Config file lives elsewhere, so the default directories can be removed
	external := filepath.Join(t.TempDir(), "elsewhere.json")
	if err := os.WriteFile(external, []byte(`{}`), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.RemoveAll(filepath.Join(home, ".config")); err != nil {
		t.Fatal(err)
	}

	_, configPath, cachePath := InitConfig(&CldConfigOptions{ConfigFile: external})

	for _, dir := range []string{configPath, cachePath} {
		info, err := os.Stat(dir)
		if err != nil || !info.IsDir() {
			t.Fatalf("expected %q to be created as a directory (err: %v)", dir, err)
		}
	}
}

func TestInitConfigEnvVarOverridesDefault(t *testing.T) {
	setupTestConfigHome(t)

	fromEnv := filepath.Join(t.TempDir(), "from-env.json")
	if err := os.WriteFile(fromEnv, []byte(`{}`), 0644); err != nil {
		t.Fatal(err)
	}
	t.Setenv("CLD_CONFIG_PATH", fromEnv)

	configFile, _, _ := InitConfig(nil)
	samePath(t, configFile, fromEnv)

	// Options with no config file set fall through to the env var as well
	configFile, _, _ = InitConfig(&CldConfigOptions{})
	samePath(t, configFile, fromEnv)
}

func TestInitConfigOptionsOverrideEnvVar(t *testing.T) {
	setupTestConfigHome(t)

	dir := t.TempDir()
	fromEnv := filepath.Join(dir, "from-env.json")
	fromOpts := filepath.Join(dir, "from-opts.json")
	for _, f := range []string{fromEnv, fromOpts} {
		if err := os.WriteFile(f, []byte(`{}`), 0644); err != nil {
			t.Fatal(err)
		}
	}
	t.Setenv("CLD_CONFIG_PATH", fromEnv)

	configFile, _, _ := InitConfig(&CldConfigOptions{ConfigFile: fromOpts})
	samePath(t, configFile, fromOpts)
}

func TestInitConfigAzureAppConfig(t *testing.T) {
	_, _, wantCachePath := setupTestConfigHome(t)
	t.Setenv("AZURE_APPCONFIG_ENDPOINT", "https://example.azconfig.io")
	// App Configuration wins over an explicit config file
	t.Setenv("CLD_CONFIG_PATH", "/some/other/config.json")

	configFile, configPath, cachePath := InitConfig(&CldConfigOptions{ConfigFile: "/ignored.json"})

	if configFile != "azureAppConfig" || configPath != "azureAppConfig" {
		t.Fatalf("expected azureAppConfig markers, got %q and %q", configFile, configPath)
	}
	samePath(t, cachePath, wantCachePath)
}

//
// Encryption option env vars
//

func TestCheckConfigEncryptionOption(t *testing.T) {
	tests := []struct {
		name        string
		encrypt     *string
		secret      string
		wantEncrypt bool
	}{
		{"nothing set", nil, "", false},
		{"secret only", nil, testConfigSecret, false},
		{"encrypt true with secret", strPtr("true"), testConfigSecret, true},
		{"encrypt 1 with secret", strPtr("1"), testConfigSecret, true},
		{"encrypt TRUE with secret", strPtr("TRUE"), testConfigSecret, true},
		{"encrypt false with secret", strPtr("false"), testConfigSecret, false},
		{"encrypt 0 without secret", strPtr("0"), "", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			setupTestConfigHome(t)
			if tc.encrypt != nil {
				t.Setenv("CLD_CONFIG_ENCRYPT", *tc.encrypt)
			}
			if tc.secret != "" {
				t.Setenv("CLD_CONFIG_SECRET", tc.secret)
			}

			gotEncrypt, gotSecret := CheckConfigEncryptionOption()
			if gotEncrypt != tc.wantEncrypt {
				t.Fatalf("encrypt: got %v, want %v", gotEncrypt, tc.wantEncrypt)
			}
			if gotSecret != tc.secret {
				t.Fatalf("secret: got %q, want %q", gotSecret, tc.secret)
			}
		})
	}
}

func strPtr(s string) *string { return &s }

//
// Encode / Decode
//

func TestEncodeDecodeRoundTrip(t *testing.T) {
	inputs := [][]byte{
		[]byte("hello world"),
		[]byte(`{"token":"abc","expiresOn":"2025-01-01T00:00:00Z"}`),
		{0x00, 0xff, 0x10, 0x80},
		{},
	}
	for _, in := range inputs {
		encoded := Encode(in)
		if string(encoded) != base64.StdEncoding.EncodeToString(in) {
			t.Fatalf("expected standard base64 for %q, got %q", in, encoded)
		}
		decoded, err := Decode(encoded)
		if err != nil {
			t.Fatalf("unexpected decode error for %q: %v", in, err)
		}
		if string(decoded) != string(in) {
			t.Fatalf("round trip mismatch: got %q, want %q", decoded, in)
		}
	}
}

func TestDecodeInvalidInput(t *testing.T) {
	if _, err := Decode([]byte("not valid base64!!")); err == nil {
		t.Fatal("expected an error for invalid base64")
	}
}

//
// Cipher key
//

func TestGetCipherKeyUsesSecretEnvVar(t *testing.T) {
	_, _, cachePath := setupTestConfigHome(t)
	t.Setenv("CLD_CONFIG_SECRET", testConfigSecret)

	if got := GetCipherKey(nil); got != testConfigSecret {
		t.Fatalf("expected secret from env var, got %q", got)
	}
	if _, err := os.Stat(filepath.Join(cachePath, "key")); err == nil {
		t.Fatal("no key file should be written when the secret comes from the env var")
	}
}

func TestGetCipherKeyGeneratesAndReusesKeyFile(t *testing.T) {
	_, _, cachePath := setupTestConfigHome(t)

	first := GetCipherKey(nil)
	if len(first) != 32 {
		t.Fatalf("expected a 32 character generated key, got %d characters", len(first))
	}

	onDisk, err := os.ReadFile(filepath.Join(cachePath, "key"))
	if err != nil {
		t.Fatalf("expected key file to be written: %v", err)
	}
	if string(onDisk) != first {
		t.Fatal("key file content should match the returned key")
	}

	if second := GetCipherKey(nil); second != first {
		t.Fatal("expected the stored key to be reused on later calls")
	}
}

//
// Encrypt / Decrypt
//

func TestEncryptDecryptRoundTrip(t *testing.T) {
	setupTestConfigHome(t)
	t.Setenv("CLD_CONFIG_SECRET", testConfigSecret)

	inputs := []string{
		`{"azure":{"multiTenantAuth":{"tenants":{}}}}`,
		"short",
		strings.Repeat("long text with unicode ✓ ", 500),
		"",
	}
	for _, in := range inputs {
		encrypted, err := Encrypt([]byte(in), nil)
		if err != nil {
			t.Fatalf("unexpected encrypt error: %v", err)
		}
		if in != "" && strings.Contains(string(encrypted), in) {
			t.Fatal("encrypted output should not contain the plain text")
		}
		if _, err := base64.StdEncoding.DecodeString(string(encrypted)); err != nil {
			t.Fatalf("expected encrypted output to be base64: %v", err)
		}

		decrypted, err := Decrypt(encrypted, nil)
		if err != nil {
			t.Fatalf("unexpected decrypt error: %v", err)
		}
		if string(decrypted) != in {
			t.Fatalf("round trip mismatch: got %q, want %q", decrypted, in)
		}
	}
}

func TestDecryptWithDifferentSecretDoesNotReturnPlainText(t *testing.T) {
	setupTestConfigHome(t)
	const plain = "super secret config data"

	t.Setenv("CLD_CONFIG_SECRET", testConfigSecret)
	encrypted, err := Encrypt([]byte(plain), nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Setenv("CLD_CONFIG_SECRET", "ffffffffffffffffffffffffffffffff")
	decrypted, err := Decrypt(encrypted, nil)
	if err != nil {
		t.Fatalf("unexpected decrypt error: %v", err)
	}
	if string(decrypted) == plain {
		t.Fatal("decrypting with a different secret should not give back the plain text")
	}
}

func TestEncryptDecryptRoundTripWithGeneratedKey(t *testing.T) {
	setupTestConfigHome(t)

	encrypted, err := Encrypt([]byte("uses the key file"), nil)
	if err != nil {
		t.Fatal(err)
	}
	decrypted, err := Decrypt(encrypted, nil)
	if err != nil {
		t.Fatal(err)
	}
	if string(decrypted) != "uses the key file" {
		t.Fatalf("round trip mismatch, got %q", decrypted)
	}
}

//
// GenerateRandomString
//

func TestGenerateRandomStringCharacterSets(t *testing.T) {
	const (
		lower   = "abcdefghijklmnopqrstuvwxyz"
		upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numbers = "0123456789"
		special = "-_~`!@#$%^&*()"
	)
	tests := []struct {
		name                    string
		upper, numbers, special bool
		allowed                 string
	}{
		{"lower only", false, false, false, lower},
		{"with upper", true, false, false, lower + upper},
		{"with numbers", false, true, false, lower + numbers},
		{"with special", false, false, true, lower + special},
		{"everything", true, true, true, lower + upper + numbers + special},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GenerateRandomString(500, tc.upper, tc.numbers, tc.special)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(got) != 500 {
				t.Fatalf("expected 500 characters, got %d", len(got))
			}
			for _, c := range got {
				if !strings.ContainsRune(tc.allowed, c) {
					t.Fatalf("unexpected character %q in %q", c, got)
				}
			}
		})
	}
}

func TestGenerateRandomStringLengths(t *testing.T) {
	for _, n := range []int{0, 1, 16, 32, 64} {
		got, err := GenerateRandomString(n, true, true, true)
		if err != nil {
			t.Fatalf("unexpected error for length %d: %v", n, err)
		}
		if len(got) != n {
			t.Fatalf("expected length %d, got %d", n, len(got))
		}
	}
}

func TestGenerateRandomStringIsNotRepeated(t *testing.T) {
	seen := map[string]bool{}
	for i := 0; i < 50; i++ {
		got, err := GenerateRandomString(32, true, true, true)
		if err != nil {
			t.Fatal(err)
		}
		if seen[got] {
			t.Fatalf("generated the same 32 character string twice: %q", got)
		}
		seen[got] = true
	}
}

//
// Token cache
//

func TestGetCachedTokenWithoutCacheFile(t *testing.T) {
	setupTestConfigHome(t)
	t.Setenv("CLD_CONFIG_SECRET", testConfigSecret)

	if got := GetCachedToken[AzureTokenData]("azure", nil); got != nil {
		t.Fatalf("expected nil when there is no cache file, got %+v", got)
	}
}

func TestCacheSaveTokenAndGetCachedToken(t *testing.T) {
	_, _, cachePath := setupTestConfigHome(t)
	t.Setenv("CLD_CONFIG_SECRET", testConfigSecret)

	expiry := time.Date(2030, time.June, 1, 12, 0, 0, 0, time.UTC)
	azToken := AzureTokenData{Token: "azure-token-value", ExpiresOn: expiry}

	if err := CacheSaveToken(azToken, "azure", nil); err != nil {
		t.Fatalf("unexpected error saving token: %v", err)
	}

	raw, err := os.ReadFile(filepath.Join(cachePath, "tkn"))
	if err != nil {
		t.Fatalf("expected cache file to be written: %v", err)
	}
	if strings.Contains(string(raw), "azure-token-value") {
		t.Fatal("token should not be stored in plain text")
	}

	got := GetCachedToken[AzureTokenData]("azure", nil)
	if got == nil {
		t.Fatal("expected cached token")
	}
	if got.Token != azToken.Token || !got.ExpiresOn.Equal(expiry) {
		t.Fatalf("got %+v, want %+v", *got, azToken)
	}
}

func TestCacheSaveTokenKeepsOtherTokenTypes(t *testing.T) {
	setupTestConfigHome(t)
	t.Setenv("CLD_CONFIG_SECRET", testConfigSecret)

	expiry := time.Date(2030, time.June, 1, 12, 0, 0, 0, time.UTC)
	mat := AzureMultiAuthToken{
		TenantId:   "tenant-id",
		TenantName: "PROD",
		TokenData:  AzureTokenData{Token: "multi-auth-token", ExpiresOn: expiry},
	}
	citrix := CitrixTokenData{AccessToken: "citrix-token", CustomerId: "cust", Expiry: expiry}

	if err := CacheSaveToken(mat, "PROD_mat", nil); err != nil {
		t.Fatal(err)
	}
	if err := CacheSaveToken(citrix, "citrix", nil); err != nil {
		t.Fatal(err)
	}
	// Saving again under an existing type replaces that entry only
	mat.TokenData.Token = "multi-auth-token-v2"
	if err := CacheSaveToken(mat, "PROD_mat", nil); err != nil {
		t.Fatal(err)
	}

	gotMat := GetCachedToken[AzureMultiAuthToken]("PROD_mat", nil)
	if gotMat == nil || gotMat.TokenData.Token != "multi-auth-token-v2" || gotMat.TenantName != "PROD" || gotMat.TenantId != "tenant-id" {
		t.Fatalf("unexpected multi auth token: %+v", gotMat)
	}
	gotCitrix := GetCachedToken[CitrixTokenData]("citrix", nil)
	if gotCitrix == nil || gotCitrix.AccessToken != "citrix-token" || gotCitrix.CustomerId != "cust" || !gotCitrix.Expiry.Equal(expiry) {
		t.Fatalf("unexpected citrix token: %+v", gotCitrix)
	}
}

func TestGetCachedTokenUnknownTypeReturnsZeroValue(t *testing.T) {
	setupTestConfigHome(t)
	t.Setenv("CLD_CONFIG_SECRET", testConfigSecret)

	if err := CacheSaveToken(AzureTokenData{Token: "x", ExpiresOn: time.Now().Add(time.Hour)}, "azure", nil); err != nil {
		t.Fatal(err)
	}

	got := GetCachedToken[AzureTokenData]("never-saved", nil)
	if got == nil {
		t.Fatal("expected a zero value token rather than nil when the cache file exists")
	}
	if got.Token != "" || !got.ExpiresOn.IsZero() {
		t.Fatalf("expected zero value token, got %+v", *got)
	}
}

func TestGetCachedTokenUnreadableCacheReturnsZeroValue(t *testing.T) {
	_, _, cachePath := setupTestConfigHome(t)
	t.Setenv("CLD_CONFIG_SECRET", testConfigSecret)

	// Valid base64 that does not decrypt to JSON, eg a cache written with another key
	garbage := base64.StdEncoding.EncodeToString([]byte("this is not an encrypted token cache"))
	if err := os.WriteFile(filepath.Join(cachePath, "tkn"), []byte(garbage), 0644); err != nil {
		t.Fatal(err)
	}

	got := GetCachedToken[AzureTokenData]("azure", nil)
	if got == nil || got.Token != "" {
		t.Fatalf("expected a zero value token, got %+v", got)
	}

	// Saving over an unreadable cache replaces it
	if err := CacheSaveToken(AzureTokenData{Token: "fresh"}, "azure", nil); err != nil {
		t.Fatal(err)
	}
	if got := GetCachedToken[AzureTokenData]("azure", nil); got == nil || got.Token != "fresh" {
		t.Fatalf("expected the fresh token after overwriting, got %+v", got)
	}
}

func TestClearTokenCache(t *testing.T) {
	_, _, cachePath := setupTestConfigHome(t)
	t.Setenv("CLD_CONFIG_SECRET", testConfigSecret)

	if err := CacheSaveToken(AzureTokenData{Token: "x"}, "azure", nil); err != nil {
		t.Fatal(err)
	}

	ClearTokenCache(nil)

	if _, err := os.Stat(filepath.Join(cachePath, "tkn")); !os.IsNotExist(err) {
		t.Fatalf("expected cache file to be removed, stat error: %v", err)
	}
	if got := GetCachedToken[AzureTokenData]("azure", nil); got != nil {
		t.Fatalf("expected nil after clearing the cache, got %+v", got)
	}
}

func TestCheckCachedTokenExpired(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name   string
		expiry time.Time
		want   bool
	}{
		{"expired an hour ago", now.Add(-time.Hour), true},
		{"expired a moment ago", now.Add(-time.Millisecond), true},
		{"expires in an hour", now.Add(time.Hour), false},
		{"zero time", time.Time{}, true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := CheckCachedTokenExpired(tc.expiry); got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

//
// Config lookups
//

func TestMapAzureSubscriptionToCustomTenantName(t *testing.T) {
	config := AzureConfig{
		CustomSubIdToTenantNameMap: map[string][]string{
			"SHARED":  {"sub-1", "sub-2"},
			"SANDBOX": {"sub-3"},
		},
	}

	tests := []struct{ subId, want string }{
		{"sub-1", "SHARED"},
		{"sub-2", "SHARED"},
		{"sub-3", "SANDBOX"},
		{"sub-unknown", ""},
		{"", ""},
		{"SUB-1", ""}, // case sensitive
	}
	for _, tc := range tests {
		if got := MapAzureSubscriptionToCustomTenantName(tc.subId, config); got != tc.want {
			t.Fatalf("%q: got %q, want %q", tc.subId, got, tc.want)
		}
	}

	if got := MapAzureSubscriptionToCustomTenantName("sub-1", AzureConfig{}); got != "" {
		t.Fatalf("expected empty name when no map is configured, got %q", got)
	}
}

func TestMapTenantIdToConfiguredTenantName(t *testing.T) {
	var config AzureConfig
	config.MultiTenantAuth.Tenants = CldConfigTenants{
		"prod": {TenantName: "PROD", TenantID: "11111111-1111-1111-1111-111111111111"},
		"dev":  {TenantName: "DEV", TenantID: "22222222-2222-2222-2222-222222222222"},
	}

	if got := MapTenantIdToConfiguredTenantName("22222222-2222-2222-2222-222222222222", config); got != "DEV" {
		t.Fatalf("got %q, want DEV", got)
	}
	if got := MapTenantIdToConfiguredTenantName("11111111-1111-1111-1111-111111111111", config); got != "PROD" {
		t.Fatalf("got %q, want PROD", got)
	}
	if got := MapTenantIdToConfiguredTenantName("unknown", config); got != "" {
		t.Fatalf("expected empty name for an unknown tenant id, got %q", got)
	}
	if got := MapTenantIdToConfiguredTenantName("anything", AzureConfig{}); got != "" {
		t.Fatalf("expected empty name with no tenants configured, got %q", got)
	}
}

func TestIsValidJson(t *testing.T) {
	valid := []string{`{}`, `[]`, `{"a":[1,2,{"b":null}]}`, `"just a string"`, `42`, `true`, `null`, "  {\"a\": 1}\n"}
	for _, s := range valid {
		if !IsValidJson(s) {
			t.Fatalf("expected %q to be valid JSON", s)
		}
	}
	invalid := []string{``, `{`, `{"a":}`, `{'a': 1}`, `{"a":1,}`, `not json`, `{"a":1} trailing`}
	for _, s := range invalid {
		if IsValidJson(s) {
			t.Fatalf("expected %q to be invalid JSON", s)
		}
	}
}
