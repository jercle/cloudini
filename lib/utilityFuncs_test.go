package lib

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestUniqueNonEmptyElementsOf(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want []string
	}{
		{"removes duplicates and keeps first-seen order", []string{"b", "a", "b", "c", "a"}, []string{"b", "a", "c"}},
		{"drops empty strings", []string{"", "a", "", "b"}, []string{"a", "b"}},
		{"case sensitive", []string{"a", "A"}, []string{"a", "A"}},
		{"only empty strings", []string{"", ""}, []string{}},
		{"nil input", nil, []string{}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := UniqueNonEmptyElementsOf(tc.in)
			if len(got) != len(tc.want) {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Fatalf("got %q, want %q", got, tc.want)
				}
			}
		})
	}
}

func TestSliceOfStringsToUnique(t *testing.T) {
	got := SliceOfStringsToUnique([]string{"x", "", "y", "x", "", "z"})
	want := []string{"x", "", "y", "z"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}

	// Unlike UniqueNonEmptyElementsOf, an empty input gives a non-nil empty slice
	if got := SliceOfStringsToUnique(nil); got == nil || len(got) != 0 {
		t.Fatalf("expected empty non-nil slice, got %#v", got)
	}
}

func TestStructToMap(t *testing.T) {
	type inner struct {
		City string
		Zip  int
	}
	type outer struct {
		Name         string
		Count        int
		Tags         []string
		Address      inner
		Capabilities []string
	}
	in := outer{
		Name:         "vm01",
		Count:        3,
		Tags:         []string{"a", "b"},
		Address:      inner{City: "Sydney", Zip: 2000},
		Capabilities: []string{"should", "be", "skipped"},
	}

	want := map[string]interface{}{
		"Name":  "vm01",
		"Count": 3,
		"Tags":  []string{"a", "b"},
		"Address": map[string]interface{}{
			"City": "Sydney",
			"Zip":  2000,
		},
	}

	if got := StructToMap(in); !reflect.DeepEqual(got, want) {
		t.Fatalf("value input: got %#v, want %#v", got, want)
	}
	if got := StructToMap(&in); !reflect.DeepEqual(got, want) {
		t.Fatalf("pointer input: got %#v, want %#v", got, want)
	}
}

func TestStructToMapSkipsCapabilitiesCaseInsensitively(t *testing.T) {
	type s struct {
		CAPABILITIES string
		Other        string
	}
	got := StructToMap(s{CAPABILITIES: "x", Other: "y"})
	if _, exists := got["CAPABILITIES"]; exists {
		t.Fatalf("expected CAPABILITIES to be skipped, got %#v", got)
	}
	if got["Other"] != "y" {
		t.Fatalf("expected Other to be kept, got %#v", got)
	}
}

func TestGetStringInBetween(t *testing.T) {
	tests := []struct {
		name, str, start, end string
		want                  string
		found                 bool
	}{
		{"simple", "hello [world] bye", "[", "]", "world", true},
		{"multi character markers", "a<<mid>>b", "<<", ">>", "mid", true},
		{"uses first start and first following end", "x(a)(b)", "(", ")", "a", true},
		{"empty result between adjacent markers", "()", "(", ")", "", true},
		{"start missing", "no markers", "[", "]", "", false},
		{"end missing", "open [only", "[", "]", "", false},
		{"end only before start", "] then [", "[", "]", "", false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, found := GetStringInBetween(tc.str, tc.start, tc.end)
			if got != tc.want || found != tc.found {
				t.Fatalf("got (%q, %v), want (%q, %v)", got, found, tc.want, tc.found)
			}
		})
	}
}

func TestIsSpecificDay(t *testing.T) {
	ts := time.Date(2025, time.March, 14, 23, 59, 59, 0, time.UTC)
	if !IsSpecificDay(ts, 2025, time.March, 14) {
		t.Fatal("expected matching day to return true")
	}
	if IsSpecificDay(ts, 2025, time.March, 15) {
		t.Fatal("expected different day to return false")
	}
	if IsSpecificDay(ts, 2025, time.April, 14) {
		t.Fatal("expected different month to return false")
	}
	if IsSpecificDay(ts, 2024, time.March, 14) {
		t.Fatal("expected different year to return false")
	}
}

func TestSelectMapStringFieldsFromArrayOfKeys(t *testing.T) {
	data := map[string]string{"a": "1", "b": "2", "c": ""}

	got := SelectMapStringFieldsFromArrayOfKeys(data, []string{"a", "c", "missing"})
	want := map[string]string{"a": "1", "c": ""}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}

	if got := SelectMapStringFieldsFromArrayOfKeys(data, nil); len(got) != 0 {
		t.Fatalf("expected empty map for no keys, got %v", got)
	}
	if len(data) != 3 {
		t.Fatalf("source map should not be modified, got %v", data)
	}
}

func TestHttpBearerGet(t *testing.T) {
	var gotAuth, gotContentType, gotMethod string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		gotContentType = r.Header.Get("Content-Type")
		gotMethod = r.Method
		switch r.URL.Path {
		case "/missing":
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, `{"error":"nope"}`)
		case "/forbidden":
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, `denied`)
		default:
			fmt.Fprint(w, `{"ok":true}`)
		}
	}))
	defer srv.Close()

	body, err := HttpBearerGet(srv.URL+"/thing", "tok123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(body) != `{"ok":true}` {
		t.Fatalf("unexpected body %q", body)
	}
	if gotAuth != "Bearer tok123" {
		t.Fatalf("unexpected Authorization header %q", gotAuth)
	}
	if gotContentType != "application/json" {
		t.Fatalf("unexpected Content-Type header %q", gotContentType)
	}
	if gotMethod != http.MethodGet {
		t.Fatalf("unexpected method %q", gotMethod)
	}

	body, err = HttpBearerGet(srv.URL+"/missing", "tok123")
	if err == nil {
		t.Fatal("expected an error for a 404 response")
	}
	if body != nil {
		t.Fatalf("expected nil body for a 404 response, got %q", body)
	}
	if !strings.Contains(err.Error(), "404") || !strings.Contains(err.Error(), `"error":"nope"`) {
		t.Fatalf("expected error to carry status and response body, got %q", err.Error())
	}

	// Only 404 is turned into an error. Other statuses hand back the body
	body, err = HttpBearerGet(srv.URL+"/forbidden", "tok123")
	if err != nil {
		t.Fatalf("unexpected error for 403: %v", err)
	}
	if string(body) != "denied" {
		t.Fatalf("unexpected body for 403: %q", body)
	}
}

// Builds an unsigned JWT with the given claims JSON. JwtIsExpiredUnverified
// does not check the signature, so the last segment only needs to decode.
func testJwt(claimsJson string) string {
	enc := base64.RawURLEncoding
	return enc.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`)) + "." +
		enc.EncodeToString([]byte(claimsJson)) + "." +
		enc.EncodeToString([]byte("signature"))
}

func TestJwtIsExpiredUnverified(t *testing.T) {
	future := time.Now().Add(time.Hour).Unix()
	past := time.Now().Add(-time.Hour).Unix()

	tests := []struct {
		name  string
		token string
		want  bool
	}{
		{"valid token expiring in the future", testJwt(fmt.Sprintf(`{"sub":"a","exp":%d}`, future)), false},
		{"token expired in the past", testJwt(fmt.Sprintf(`{"sub":"a","exp":%d}`, past)), true},
		{"token without exp claim", testJwt(`{"sub":"a"}`), true},
		{"token with non numeric exp claim", testJwt(`{"exp":"soon"}`), true},
		{"not a jwt", "definitely-not-a-jwt", true},
		{"empty string", "", true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := JwtIsExpiredUnverified(tc.token); got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestLoggerJson(t *testing.T) {
	if LoggerJson() == nil {
		t.Fatal("expected a logger")
	}
}

//
// Struct file tag rewriters
//

func writeStructFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "types.go")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	return path
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

const structWithJsonTags = "package x\n\ntype Thing struct {\n\tName string `json:\"name\"`\n\tSize int\n}\n"

func TestAddJsonOmitemptyTagsToStructFile(t *testing.T) {
	path := writeStructFile(t, structWithJsonTags)

	got := AddJsonOmitemptyTagsToStructFile(path, false)
	if !strings.Contains(got, "`json:\"name,omitempty,omitzero\"`") {
		t.Fatalf("expected omitempty and omitzero to be added, got:\n%s", got)
	}
	if !strings.Contains(got, "Size int") {
		t.Fatalf("expected untagged field to be left alone, got:\n%s", got)
	}
	if readFile(t, path) != structWithJsonTags {
		t.Fatal("file should not change when overwriteFile is false")
	}

	// Running over the output again is a no-op
	path2 := writeStructFile(t, got)
	if again := AddJsonOmitemptyTagsToStructFile(path2, false); again != got {
		t.Fatalf("expected already processed file to be unchanged, got:\n%s", again)
	}
}

func TestAddJsonOmitemptyTagsToStructFileStripsAttr(t *testing.T) {
	path := writeStructFile(t, "package x\n\ntype Thing struct {\n\tID string `json:\"id,attr\"`\n}\n")
	got := AddJsonOmitemptyTagsToStructFile(path, false)
	if !strings.Contains(got, "`json:\"id,omitempty,omitzero\"`") {
		t.Fatalf("expected ,attr to be removed, got:\n%s", got)
	}
}

func TestAddJsonOmitemptyTagsToStructFileOverwrite(t *testing.T) {
	path := writeStructFile(t, structWithJsonTags)
	got := AddJsonOmitemptyTagsToStructFile(path, true)
	if onDisk := readFile(t, path); onDisk != got {
		t.Fatalf("expected file to hold the processed struct.\nfile:\n%s\nreturned:\n%s", onDisk, got)
	}
}

func TestAddBsonTagsFromJsonTagsToStructFile(t *testing.T) {
	path := writeStructFile(t, "package x\n\ntype Thing struct {\n\tName string `json:\"name,omitempty\"`\n\tKept string `json:\"kept\" bson:\"other\"`\n}\n")

	got := AddBsonTagsFromJsonTagsToStructFile(path, false)
	if !strings.Contains(got, "`json:\"name,omitempty\" bson:\"name,omitempty\"`") {
		t.Fatalf("expected bson tag copied from json tag, got:\n%s", got)
	}
	if !strings.Contains(got, "`json:\"kept\" bson:\"other\"`") {
		t.Fatalf("expected existing bson tag to be left alone, got:\n%s", got)
	}

	AddBsonTagsFromJsonTagsToStructFile(path, true)
	if onDisk := readFile(t, path); onDisk != got {
		t.Fatalf("expected overwrite to write the processed struct, got:\n%s", onDisk)
	}
}

func TestAddXmlTagsFromJsonTagsToStructFile(t *testing.T) {
	path := writeStructFile(t, "package x\n\ntype Thing struct {\n\tName string `json:\"name\"`\n\tKept string `json:\"kept\" xml:\"other\"`\n}\n")

	got := AddXmlTagsFromJsonTagsToStructFile(path, false)
	if !strings.Contains(got, "`json:\"name\" xml:\"name\"`") {
		t.Fatalf("expected xml tag copied from json tag, got:\n%s", got)
	}
	if !strings.Contains(got, "`json:\"kept\" xml:\"other\"`") {
		t.Fatalf("expected existing xml tag to be left alone, got:\n%s", got)
	}
}

func TestAddTagsFromXmlTagsToStructFile(t *testing.T) {
	path := writeStructFile(t, "package x\n\ntype Thing struct {\n\tName string `xml:\"name\"`\n\tKept string `xml:\"kept\" json:\"other\"`\n}\n")

	got := AddTagsFromXmlTagsToStructFile(path, false)
	if !strings.Contains(got, "`xml:\"name\" json:\"name\"`") {
		t.Fatalf("expected json tag copied from xml tag, got:\n%s", got)
	}
	if !strings.Contains(got, "`xml:\"kept\" json:\"other\"`") {
		t.Fatalf("expected existing json tag to be left alone, got:\n%s", got)
	}
}

func TestChangeJsonIdToMongoIdInFile(t *testing.T) {
	src := "package x\n\ntype Thing struct {\n\tID string `json:\"id\" bson:\"id\"`\n\tGuid string `json:\"guid\"`\n}"
	path := writeStructFile(t, src)

	got := ChangeJsonIdToMongoIdInFile(path, false)
	if !strings.Contains(got, "`json:\"_id\" bson:\"_id\"`") {
		t.Fatalf("expected \"id\" tags to become \"_id\", got:\n%s", got)
	}
	if !strings.Contains(got, "`json:\"guid\"`") {
		t.Fatalf("expected other tags to be left alone, got:\n%s", got)
	}
	if readFile(t, path) != src {
		t.Fatal("file should not change when overwriteFile is false")
	}

	ChangeJsonIdToMongoIdInFile(path, true)
	if onDisk := readFile(t, path); onDisk != got {
		t.Fatalf("expected overwrite to write the processed struct, got:\n%s", onDisk)
	}
}

func TestChangeCaseOfJsonAndBsonTags(t *testing.T) {
	path := writeStructFile(t, "package x\n\ntype Thing struct {\n\tName string `json:\"Name\" bson:\"Name\"`\n\tSize int\n}\n")

	got := ChangeCaseOfJsonAndBsonTags(path, false)
	if !strings.Contains(got, "`json:\"name\" bson:\"name\"`") {
		t.Fatalf("expected first letter of both tags to be lower cased, got:\n%s", got)
	}
	if !strings.Contains(got, "Size int") {
		t.Fatalf("expected untagged field to be left alone, got:\n%s", got)
	}
}
