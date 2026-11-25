package lib

import (
	json "encoding/json/v2"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

func TestFormatErrorStringIncludesCallerAndMessage(t *testing.T) {
	err := FormatErrorString("something went wrong")
	if err == nil {
		t.Fatal("expected an error")
	}
	msg := err.Error()
	if !strings.Contains(msg, "something went wrong") {
		t.Fatalf("expected original message in %q", msg)
	}
	if !strings.Contains(msg, "wrapperFuncs_test.go:") {
		t.Fatalf("expected the calling file and line in %q", msg)
	}
	if !strings.Contains(msg, "Original caller: ") {
		t.Fatalf("expected the caller prefix in %q", msg)
	}
}

func TestFormatErrorIncludesCallerAndMessage(t *testing.T) {
	err := FormatError(errors.New("underlying failure"))
	if err == nil {
		t.Fatal("expected an error")
	}
	msg := err.Error()
	if !strings.Contains(msg, "underlying failure") {
		t.Fatalf("expected original message in %q", msg)
	}
	if !strings.Contains(msg, "wrapperFuncs_test.go:") {
		t.Fatalf("expected the calling file and line in %q", msg)
	}
}

func TestCheckFatalErrorNilDoesNotExit(t *testing.T) {
	// Would end the test binary if it treated nil as an error
	CheckFatalError(nil)
	CheckErrorNonFatal(nil)
	CheckErrorNonFatal(errors.New("printed, not fatal"))
}

func TestWaitGroupCount(t *testing.T) {
	var wg WaitGroupCount

	if got := wg.GetCount(); got != 0 {
		t.Fatalf("expected a new wait group to count 0, got %d", got)
	}

	wg.Add(3)
	if got := wg.GetCount(); got != 3 {
		t.Fatalf("expected 3 after Add(3), got %d", got)
	}
	wg.Add(2)
	if got := wg.GetCount(); got != 5 {
		t.Fatalf("expected 5 after Add(2), got %d", got)
	}

	wg.Done()
	wg.Done()
	if got := wg.GetCount(); got != 3 {
		t.Fatalf("expected 3 after two Done calls, got %d", got)
	}

	wg.Done()
	wg.Done()
	wg.Done()
	wg.Wait() // returns straight away once everything is done
	if got := wg.GetCount(); got != 0 {
		t.Fatalf("expected 0 once everything is done, got %d", got)
	}
}

func TestWaitGroupCountConcurrentUse(t *testing.T) {
	const workers = 200
	var (
		wg      WaitGroupCount
		started sync.WaitGroup
		release = make(chan struct{})
	)

	started.Add(workers)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			started.Done()
			<-release
		}()
	}

	started.Wait()
	if got := wg.GetCount(); got != workers {
		t.Fatalf("expected %d in flight, got %d", workers, got)
	}

	close(release)
	wg.Wait()
	if got := wg.GetCount(); got != 0 {
		t.Fatalf("expected 0 after Wait, got %d", got)
	}
}

func TestJsonMarshalAndWriteFile(t *testing.T) {
	type payload struct {
		Name  string   `json:"name"`
		Count int      `json:"count"`
		Tags  []string `json:"tags"`
	}
	want := payload{Name: "cloudini", Count: 2, Tags: []string{"a", "b"}}
	path := filepath.Join(t.TempDir(), "out.json")

	JsonMarshalAndWriteFile(want, path)

	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("expected file to be written: %v", err)
	}
	var got payload
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("expected valid JSON in file, got %q: %v", raw, err)
	}
	if got.Name != want.Name || got.Count != want.Count || len(got.Tags) != 2 || got.Tags[0] != "a" || got.Tags[1] != "b" {
		t.Fatalf("got %+v, want %+v", got, want)
	}
	if !strings.Contains(string(raw), "\n  \"name\"") {
		t.Fatalf("expected two space indented output, got %q", raw)
	}
}
