package datatransforms

import (
	json "encoding/json/v2"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func writeTempFile(t *testing.T, name string, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	return path
}

func decodeEntries(t *testing.T, jsonStr string) []map[string]interface{} {
	t.Helper()
	var entries []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &entries); err != nil {
		t.Fatalf("output is not a JSON array of objects: %v\n%s", err, jsonStr)
	}
	return entries
}

func TestArrayContentMatch(t *testing.T) {
	tests := []struct {
		in        string
		wantKey   string
		wantIndex int
	}{
		{"name", "name", -1},
		{"tags[0]", "tags", 0},
		{"tags[12]", "tags", 12},
		{"tags[]", "tags", 0},
		{"tags[abc]", "tags", 0},
		{"tags[1", "tags[1", -1},
		{"", "", -1},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			key, index := arrayContentMatch(tc.in)
			if key != tc.wantKey || index != tc.wantIndex {
				t.Fatalf("got (%q, %d), want (%q, %d)", key, index, tc.wantKey, tc.wantIndex)
			}
		})
	}
}

func TestCsvToJsonFlatRows(t *testing.T) {
	out, err := csvToJson([][]string{
		{"name", "size", "location"},
		{"vm01", "Standard_D2s_v5", "australiaeast"},
		{"vm02", "Standard_B2ms", "australiasoutheast"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []map[string]interface{}{
		{"name": "vm01", "size": "Standard_D2s_v5", "location": "australiaeast"},
		{"name": "vm02", "size": "Standard_B2ms", "location": "australiasoutheast"},
	}
	if got := decodeEntries(t, out); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestCsvToJsonValuesStayStrings(t *testing.T) {
	out, err := csvToJson([][]string{
		{"count", "enabled", "empty"},
		{"42", "true", ""},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []map[string]interface{}{{"count": "42", "enabled": "true", "empty": ""}}
	if got := decodeEntries(t, out); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestCsvToJsonNestedObjects(t *testing.T) {
	out, err := csvToJson([][]string{
		{"name", "address.city", "address.geo.lat", "address.geo.long"},
		{"office", "Sydney", "-33.86", "151.20"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []map[string]interface{}{{
		"name": "office",
		"address": map[string]interface{}{
			"city": "Sydney",
			"geo":  map[string]interface{}{"lat": "-33.86", "long": "151.20"},
		},
	}}
	if got := decodeEntries(t, out); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestCsvToJsonArrays(t *testing.T) {
	out, err := csvToJson([][]string{
		{"name", "tags[0]", "tags[1]", "tags[2]"},
		{"vm01", "prod", "web", "au"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []map[string]interface{}{{
		"name": "vm01",
		"tags": []interface{}{"prod", "web", "au"},
	}}
	if got := decodeEntries(t, out); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestCsvToJsonArraysOfObjects(t *testing.T) {
	out, err := csvToJson([][]string{
		{"name", "disks[0].name", "disks[0].size", "disks[1].name", "disks[1].size"},
		{"vm01", "os", "128", "data", "512"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []map[string]interface{}{{
		"name": "vm01",
		"disks": []interface{}{
			map[string]interface{}{"name": "os", "size": "128"},
			map[string]interface{}{"name": "data", "size": "512"},
		},
	}}
	if got := decodeEntries(t, out); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestCsvToJsonHeaderOnly(t *testing.T) {
	out, err := csvToJson([][]string{{"name", "size"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// With no data rows there are no entries. json v2 writes a nil slice as []
	if out != "[]" {
		t.Fatalf("got %q, want %q", out, "[]")
	}
}

func TestCsvToJsonOutputIsTabIndented(t *testing.T) {
	out, err := csvToJson([][]string{{"name"}, {"vm01"}})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "\n\t{\n\t\t\"name\": \"vm01\"\n\t}\n") {
		t.Fatalf("expected tab indented output, got %q", out)
	}
}

func TestReadAndParseCsvIsTabDelimited(t *testing.T) {
	path := writeTempFile(t, "data.tsv", "name\tdescription\nvm01\thas, a comma\nvm02\t\"quoted\tvalue\"\n")

	rows, err := readAndParseCsv(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := [][]string{
		{"name", "description"},
		{"vm01", "has, a comma"},
		{"vm02", "quoted\tvalue"},
	}
	if !reflect.DeepEqual(rows, want) {
		t.Fatalf("got %q, want %q", rows, want)
	}
}

func TestReadAndParseCsvMissingFile(t *testing.T) {
	_, err := readAndParseCsv(filepath.Join(t.TempDir(), "missing.csv"))
	if err == nil {
		t.Fatal("expected an error for a missing file")
	}
	if !strings.Contains(err.Error(), "error opening") {
		t.Fatalf("unexpected error %q", err.Error())
	}
}

func TestReadAndParseCsvInconsistentColumns(t *testing.T) {
	path := writeTempFile(t, "bad.tsv", "a\tb\n1\t2\t3\n")

	rows, err := readAndParseCsv(path)
	if err == nil {
		t.Fatal("expected an error for a row with too many fields")
	}
	if !strings.Contains(err.Error(), "failed to parse csv") {
		t.Fatalf("unexpected error %q", err.Error())
	}
	// Rows read before the bad one are still returned
	if len(rows) != 1 {
		t.Fatalf("expected the header row to be returned, got %q", rows)
	}
}

func TestCsvFileToJson(t *testing.T) {
	path := writeTempFile(t, "vms.tsv", "name\tnetwork.vnet\tips[0]\tips[1]\nvm01\tvnet-prod\t10.0.0.4\t10.0.0.5\n")

	out := CsvFileToJson(path)

	want := []map[string]interface{}{{
		"name":    "vm01",
		"network": map[string]interface{}{"vnet": "vnet-prod"},
		"ips":     []interface{}{"10.0.0.4", "10.0.0.5"},
	}}
	if got := decodeEntries(t, out); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestCsvFileToJsonPanicsOnMissingFile(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected a panic for a missing file")
		}
	}()
	CsvFileToJson(filepath.Join(t.TempDir(), "missing.csv"))
}
