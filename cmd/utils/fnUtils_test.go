package utils

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Runs fn and returns everything it wrote to os.Stdout
func captureStdout(t *testing.T, fn func()) string {
	t.Helper()
	original := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	defer func() { os.Stdout = original }()

	done := make(chan string)
	go func() {
		out, _ := io.ReadAll(r)
		done <- string(out)
	}()

	fn()

	w.Close()
	os.Stdout = original
	return <-done
}

func TestGeneratePassword(t *testing.T) {
	const (
		lower   = "abcdefghijklmnopqrstuvwxyz"
		upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		numbers = "0123456789"
		special = "-_~`!@#$%^&*()"
	)
	tests := []struct {
		name                    string
		length                  int
		upper, numbers, special bool
		allowed                 string
	}{
		{"lower only", 40, false, false, false, lower},
		{"upper and numbers", 40, true, true, false, lower + upper + numbers},
		{"everything", 200, true, true, true, lower + upper + numbers + special},
		{"single character", 1, true, true, true, lower + upper + numbers + special},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GeneratePassword(tc.length, tc.upper, tc.numbers, tc.special)
			if len(got) != tc.length {
				t.Fatalf("expected %d characters, got %d", tc.length, len(got))
			}
			for _, c := range got {
				if !strings.ContainsRune(tc.allowed, c) {
					t.Fatalf("unexpected character %q in %q", c, got)
				}
			}
		})
	}

	if GeneratePassword(32, true, true, true) == GeneratePassword(32, true, true, true) {
		t.Fatal("two generated passwords should not be the same")
	}
}

func TestGenerateGpoBkupInfoToStdOut(t *testing.T) {
	out := captureStdout(t, func() {
		GenerateGpoBkupInfoToStdOut(
			"Workstation Baseline",
			"11111111-2222-3333-4444-555555555555",
			"{AAAAAAAA-BBBB-CCCC-DDDD-EEEEEEEEEEEE}",
			"corp.example.com",
			"dc01.corp.example.com",
		)
	})

	for _, want := range []string{
		`<BackupInst xmlns="http://www.microsoft.com/GroupPolicy/GPOOperations/Manifest">`,
		"<GPOGuid>",
		"11111111-2222-3333-4444-555555555555",
		"<GPODomain>",
		"corp.example.com",
		"<GPODomainController>",
		"dc01.corp.example.com",
		"<ID>",
		"AAAAAAAA-BBBB-CCCC-DDDD-EEEEEEEEEEEE",
		"<GPODisplayName>",
		"Workstation Baseline",
		"</BackupInst>",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("expected output to contain %q, got:\n%s", want, out)
		}
	}
}

func TestGenerateGpoBkupInfoFiles(t *testing.T) {
	backupPath := t.TempDir()

	// <backup path>/<GPO name>/<GPO guid>/<backup id>/
	gpos := []struct{ name, guid, backupId string }{
		{"Workstation Baseline", "{11111111-2222-3333-4444-555555555555}", "{AAAAAAAA-0000-0000-0000-000000000001}"},
		{"Server Hardening", "{66666666-7777-8888-9999-000000000000}", "{BBBBBBBB-0000-0000-0000-000000000002}"},
	}
	for _, gpo := range gpos {
		if err := os.MkdirAll(filepath.Join(backupPath, gpo.name, gpo.guid, gpo.backupId), 0755); err != nil {
			t.Fatal(err)
		}
	}
	// WMI filters live alongside the GPO backups and are skipped
	if err := os.MkdirAll(filepath.Join(backupPath, "WMI_Filters"), 0755); err != nil {
		t.Fatal(err)
	}

	captureStdout(t, func() {
		GenerateGpoBkupInfoFiles(backupPath)
	})

	for _, gpo := range gpos {
		infoFile := filepath.Join(backupPath, gpo.name, gpo.guid, gpo.backupId, "bkupInfo.xml")
		raw, err := os.ReadFile(infoFile)
		if err != nil {
			t.Fatalf("expected %s to be created: %v", infoFile, err)
		}
		content := string(raw)
		for _, want := range []string{
			"<BackupInst ",
			gpo.name,
			strings.Trim(gpo.guid, "{}"),
			strings.Trim(gpo.backupId, "{}"),
			"</BackupInst>",
		} {
			if !strings.Contains(content, want) {
				t.Fatalf("expected %s to contain %q, got:\n%s", infoFile, want, content)
			}
		}
	}

	entries, err := os.ReadDir(filepath.Join(backupPath, "WMI_Filters"))
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 0 {
		t.Fatalf("expected nothing to be written under WMI_Filters, found %d entries", len(entries))
	}
}
