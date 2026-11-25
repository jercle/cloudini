package lib

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"testing"
)

func mustWrite(t *testing.T, path string, data []byte) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatal(err)
	}
}

func TestGetFullFilePaths(t *testing.T) {
	dir := t.TempDir()
	mustWrite(t, filepath.Join(dir, "a.txt"), []byte("a"))
	mustWrite(t, filepath.Join(dir, "sub", "b.txt"), []byte("b"))
	mustWrite(t, filepath.Join(dir, "sub", "deeper", "c.txt"), []byte("c"))
	if err := os.MkdirAll(filepath.Join(dir, "empty"), 0755); err != nil {
		t.Fatal(err)
	}

	got := GetFullFilePaths(dir)
	sort.Strings(got)
	want := []string{
		filepath.Join(dir, "a.txt"),
		filepath.Join(dir, "sub", "b.txt"),
		filepath.Join(dir, "sub", "deeper", "c.txt"),
	}
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestGetFullFilePathsEmptyDir(t *testing.T) {
	if got := GetFullFilePaths(t.TempDir()); len(got) != 0 {
		t.Fatalf("expected no files, got %v", got)
	}
}

func TestRemoveSubdirectoriesOfPath(t *testing.T) {
	dir := t.TempDir()
	mustWrite(t, filepath.Join(dir, "one", "a.txt"), []byte("a"))
	mustWrite(t, filepath.Join(dir, "two", "nested", "b.txt"), []byte("b"))
	mustWrite(t, filepath.Join(dir, "top.txt"), []byte("top"))

	RemoveSubdirectoriesOfPath(dir)

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	// Everything directly under the path is removed, top level files included
	if len(entries) != 0 {
		t.Fatalf("expected directory to be emptied, still has %d entries", len(entries))
	}
	if !CheckDirExists(dir) {
		t.Fatal("the path itself should still exist")
	}
}

func TestCheckDirExists(t *testing.T) {
	dir := t.TempDir()
	if !CheckDirExists(dir) {
		t.Fatal("expected existing directory to return true")
	}
	if CheckDirExists(filepath.Join(dir, "nope")) {
		t.Fatal("expected missing directory to return false")
	}

	// Despite the name, this is an existence check and is true for files too
	file := filepath.Join(dir, "file.txt")
	mustWrite(t, file, []byte("x"))
	if !CheckDirExists(file) {
		t.Fatal("expected existing file to return true")
	}
}

func TestDeleteFilesInDirMatchingString(t *testing.T) {
	dir := t.TempDir()
	mustWrite(t, filepath.Join(dir, "a.log"), []byte("a"))
	mustWrite(t, filepath.Join(dir, "b.log"), []byte("b"))
	mustWrite(t, filepath.Join(dir, "keep.txt"), []byte("k"))

	// The pattern is appended directly to dir, so dir needs its trailing separator
	DeleteFilesInDirMatchingString(dir+string(filepath.Separator), "*.log")

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 || entries[0].Name() != "keep.txt" {
		var names []string
		for _, e := range entries {
			names = append(names, e.Name())
		}
		t.Fatalf("expected only keep.txt to remain, got %v", names)
	}
}

func TestRemoveJsonByteOrderMark(t *testing.T) {
	withBom := append([]byte("\xef\xbb\xbf"), []byte(`{"a":1}`)...)
	if got := RemoveJsonByteOrderMark(withBom); string(got) != `{"a":1}` {
		t.Fatalf("expected BOM to be stripped, got %q", got)
	}
	if got := RemoveJsonByteOrderMark([]byte(`{"a":1}`)); string(got) != `{"a":1}` {
		t.Fatalf("expected input without BOM to be unchanged, got %q", got)
	}
	if got := RemoveJsonByteOrderMark(nil); len(got) != 0 {
		t.Fatalf("expected empty output for nil input, got %q", got)
	}

	// Only a leading BOM is removed
	inner := []byte("a\xef\xbb\xbfb")
	if got := RemoveJsonByteOrderMark(inner); string(got) != string(inner) {
		t.Fatalf("expected non leading BOM to be kept, got %q", got)
	}
}

func TestSplitPath(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("expectations use unix style paths")
	}
	tests := []struct {
		in   string
		want []string
	}{
		{"/usr/local/bin", []string{"/", "usr", "local", "bin"}},
		{"a/b/c", []string{"a", "b", "c"}},
		{"a/b/c/", []string{"a", "b", "c"}},
		{"a//b/./c", []string{"a", "b", "c"}},
		{"file.txt", []string{"file.txt"}},
		{"/", []string{"/"}},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			if got := SplitPath(tc.in); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

// "hi\n" style helper: encodes ASCII text as UTF-16
func utf16Bytes(s string, littleEndian bool, bom bool) []byte {
	var out []byte
	if bom {
		if littleEndian {
			out = append(out, 0xFF, 0xFE)
		} else {
			out = append(out, 0xFE, 0xFF)
		}
	}
	for i := 0; i < len(s); i++ {
		if littleEndian {
			out = append(out, s[i], 0x00)
		} else {
			out = append(out, 0x00, s[i])
		}
	}
	return out
}

func TestReadFileUTF16(t *testing.T) {
	const text = "{\"name\": \"cloudini\"}"
	tests := []struct {
		name string
		data []byte
	}{
		{"little endian with BOM", utf16Bytes(text, true, true)},
		{"big endian with BOM", utf16Bytes(text, false, true)},
		{"big endian without BOM", utf16Bytes(text, false, false)},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "utf16.json")
			mustWrite(t, path, tc.data)

			got, err := ReadFileUTF16(path)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if string(got) != text {
				t.Fatalf("got %q, want %q", got, text)
			}
		})
	}
}

func TestReadFileUTF16MissingFile(t *testing.T) {
	if _, err := ReadFileUTF16(filepath.Join(t.TempDir(), "missing")); err == nil {
		t.Fatal("expected an error for a missing file")
	}
}
