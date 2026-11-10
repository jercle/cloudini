package lib

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

var testKey = []byte("0123456789abcdef0123456789abcdef")

func TestEncryptRoundTrip(t *testing.T) {
	plain := []byte(`{"secret":"value"}`)

	enc, err := encryptWithKey(testKey, plain)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.HasPrefix(enc, []byte(encryptedV2Prefix)) {
		t.Fatalf("expected %q prefix, got %q", encryptedV2Prefix, enc)
	}

	dec, err := decryptWithKey(testKey, enc)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(dec, plain) {
		t.Fatalf("got %q, want %q", dec, plain)
	}
}

func TestEncryptUsesRandomNonce(t *testing.T) {
	plain := []byte("same plaintext")
	a, _ := encryptWithKey(testKey, plain)
	b, _ := encryptWithKey(testKey, plain)
	if bytes.Equal(a, b) {
		t.Fatal("encrypting the same plaintext twice produced identical output")
	}
}

func TestDecryptDetectsTamperingAndWrongKey(t *testing.T) {
	enc, _ := encryptWithKey(testKey, []byte("payload"))

	tampered := bytes.Clone(enc)
	i := len(tampered) - 6
	if tampered[i] == 'A' {
		tampered[i] = 'B'
	} else {
		tampered[i] = 'A'
	}
	if _, err := decryptWithKey(testKey, tampered); err == nil {
		t.Fatal("expected error for tampered ciphertext")
	}

	wrongKey := []byte("abcdef0123456789abcdef0123456789")
	if _, err := decryptWithKey(wrongKey, enc); err == nil {
		t.Fatal("expected error for wrong key")
	}
}

func TestDecryptLegacyFormat(t *testing.T) {
	plain := []byte(`{"legacy":true}`)

	block, err := aes.NewCipher(testKey)
	if err != nil {
		t.Fatal(err)
	}
	cipherText := make([]byte, len(plain))
	cipher.NewCFBEncrypter(block, legacyCipherIV).XORKeyStream(cipherText, plain)
	legacy := Encode(cipherText)

	dec, err := decryptWithKey(testKey, legacy)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(dec, plain) {
		t.Fatalf("got %q, want %q", dec, plain)
	}
}

func TestTokenCacheRoundTrip(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)
	t.Setenv("CLD_CONFIG_PATH", "")
	t.Setenv("CLD_CONFIG_SECRET", "")
	t.Setenv("CLD_CONFIG_ENCRYPT", "")
	t.Setenv("AZURE_APPCONFIG_ENDPOINT", "")
	os.Unsetenv("CLD_CONFIG_ENCRYPT")

	if got := GetCachedToken[AzureTokenData]("aztestdefaultread", nil); got != nil {
		t.Fatalf("expected nil before any token is cached, got %+v", got)
	}

	read := AzureTokenData{Token: "read-token", ExpiresOn: time.Now().Add(time.Hour).UTC().Truncate(time.Second)}
	write := AzureTokenData{Token: "write-token", ExpiresOn: time.Now().Add(time.Hour).UTC().Truncate(time.Second)}

	if err := CacheSaveToken(read, "aztestdefaultread", nil); err != nil {
		t.Fatal(err)
	}
	if err := CacheSaveToken(write, "aztestdefaultwrite", nil); err != nil {
		t.Fatal(err)
	}

	gotRead := GetCachedToken[AzureTokenData]("aztestdefaultread", nil)
	if gotRead == nil || gotRead.Token != read.Token || !gotRead.ExpiresOn.Equal(read.ExpiresOn) {
		t.Fatalf("read token mismatch: %+v", gotRead)
	}
	gotWrite := GetCachedToken[AzureTokenData]("aztestdefaultwrite", nil)
	if gotWrite == nil || gotWrite.Token != write.Token {
		t.Fatalf("write token mismatch: %+v", gotWrite)
	}

	missing := GetCachedToken[AzureTokenData]("does-not-exist", nil)
	if missing == nil || missing.Token != "" || !CheckCachedTokenExpired(missing.ExpiresOn) {
		t.Fatalf("expected empty, expired token for unknown key, got %+v", missing)
	}

	// A corrupted cache must behave like an empty cache, not abort the program.
	cacheFile := filepath.Join(home, ".config", "cld", "cache", "tkn")
	if err := os.WriteFile(cacheFile, []byte(encryptedV2Prefix+"AAAA"), 0o600); err != nil {
		t.Fatal(err)
	}
	if got := GetCachedToken[AzureTokenData]("aztestdefaultread", nil); got == nil || got.Token != "" {
		t.Fatalf("expected empty token from corrupted cache, got %+v", got)
	}

	if runtime.GOOS != "windows" {
		for _, p := range []string{
			filepath.Join(home, ".config", "cld"),
			filepath.Join(home, ".config", "cld", "cache"),
			filepath.Join(home, ".config", "cld", "cldConf.json"),
			filepath.Join(home, ".config", "cld", "cache", "key"),
			cacheFile,
		} {
			info, err := os.Stat(p)
			if err != nil {
				t.Fatal(err)
			}
			if info.Mode().Perm()&0o077 != 0 {
				t.Errorf("%s is accessible to group/other: %v", p, info.Mode().Perm())
			}
		}
	}
}

func TestCheckCachedTokenExpired(t *testing.T) {
	if !CheckCachedTokenExpired(time.Now().Add(-time.Minute)) {
		t.Error("past expiry should be expired")
	}
	if !CheckCachedTokenExpired(time.Now().Add(30 * time.Second)) {
		t.Error("expiry inside the skew window should be treated as expired")
	}
	if CheckCachedTokenExpired(time.Now().Add(time.Hour)) {
		t.Error("expiry an hour away should not be expired")
	}
}
