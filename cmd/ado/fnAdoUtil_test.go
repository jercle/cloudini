package ado

import (
	"encoding/base64"
	json "encoding/json/v2"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestCreateBasicAuthHeaderValue(t *testing.T) {
	got := CreateBasicAuthHeaderValue("", "my-pat")

	if !strings.HasPrefix(got, "Basic ") {
		t.Fatalf("expected a Basic auth value, got %q", got)
	}
	decoded, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(got, "Basic "))
	if err != nil {
		t.Fatalf("expected base64 credentials: %v", err)
	}
	if string(decoded) != ":my-pat" {
		t.Fatalf("expected an empty username with the PAT as password, got %q", decoded)
	}
}

func TestCreateBasicAuthHeaderValueIgnoresUsername(t *testing.T) {
	// Azure DevOps PAT auth uses an empty username, so the username is not sent
	withUser := CreateBasicAuthHeaderValue("someone@example.com", "my-pat")
	withoutUser := CreateBasicAuthHeaderValue("", "my-pat")
	if withUser != withoutUser {
		t.Fatalf("expected the username to be ignored, got %q and %q", withUser, withoutUser)
	}
}

func TestCreateBasicAuthHeaderValueMatchesHttpBasicAuth(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "https://dev.azure.com/org", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", CreateBasicAuthHeaderValue("", "p@t:with/odd+chars="))

	user, pass, ok := req.BasicAuth()
	if !ok {
		t.Fatal("expected the header to parse as basic auth")
	}
	if user != "" || pass != "p@t:with/odd+chars=" {
		t.Fatalf("got user %q and password %q", user, pass)
	}
}

func TestAzureDevOpsRestGet(t *testing.T) {
	var gotMethod, gotContentType, gotPat string
	var gotBasicAuth bool
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotContentType = r.Header.Get("Content-Type")
		_, gotPat, gotBasicAuth = r.BasicAuth()
		fmt.Fprint(w, `{"count": 2, "value": [{"id": 1, "name": "build"}, {"id": 2, "name": "release"}]}`)
	}))
	defer srv.Close()

	got := azureDevOpsRestGet(srv.URL+"/org/project/_apis/pipelines?"+AzureDevopsApiVersion, "my-pat")

	if gotMethod != http.MethodGet {
		t.Fatalf("unexpected method %q", gotMethod)
	}
	if gotContentType != "application/json" {
		t.Fatalf("unexpected Content-Type %q", gotContentType)
	}
	if !gotBasicAuth || gotPat != "my-pat" {
		t.Fatalf("expected the PAT to be sent as the basic auth password, got %q (basic auth: %v)", gotPat, gotBasicAuth)
	}

	var decoded, want map[string]interface{}
	if err := json.Unmarshal(got, &decoded); err != nil {
		t.Fatalf("expected JSON to be returned, got %q: %v", got, err)
	}
	json.Unmarshal([]byte(`{"count": 2, "value": [{"id": 1, "name": "build"}, {"id": 2, "name": "release"}]}`), &want)
	if !reflect.DeepEqual(decoded, want) {
		t.Fatalf("got %v, want %v", decoded, want)
	}
}

func TestAzureDevOpsDefaults(t *testing.T) {
	if !strings.HasPrefix(AzureDevOpsBaseUrl, "https://") || !strings.HasSuffix(AzureDevOpsBaseUrl, "/") {
		t.Fatalf("base URL should be https and end with a slash so org names can be appended, got %q", AzureDevOpsBaseUrl)
	}
	if !strings.HasPrefix(AzureDevopsApiVersion, "api-version=") {
		t.Fatalf("API version should be a ready to append query parameter, got %q", AzureDevopsApiVersion)
	}
}
