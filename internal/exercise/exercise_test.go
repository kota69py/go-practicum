package exercise

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"
	"time"
)

func TestLoadFromFS(t *testing.T) {
	fsys := fstest.MapFS{
		"01-hello/exercise.json": &fstest.MapFile{
			Data: []byte(`{"title":"Hello","category":"basics","difficulty":1,"topics":["fmt"],"hints":["use fmt.Println"],"files":["main.go"]}`),
		},
	}
	ex, err := LoadFromFS(fsys, "01-hello")
	if err != nil {
		t.Fatalf("LoadFromFS error: %v", err)
	}
	if ex.Name != "01-hello" {
		t.Errorf("Name = %q, want %q", ex.Name, "01-hello")
	}
	if ex.Title != "Hello" {
		t.Errorf("Title = %q, want %q", ex.Title, "Hello")
	}
}

func TestLoadFromFS_NotExist(t *testing.T) {
	fsys := fstest.MapFS{}
	_, err := LoadFromFS(fsys, "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent exercise")
	}
}

func TestLoadFromFS_BadJSON(t *testing.T) {
	fsys := fstest.MapFS{
		"bad/exercise.json": &fstest.MapFile{
			Data: []byte(`not json`),
		},
	}
	_, err := LoadFromFS(fsys, "bad")
	if err == nil {
		t.Fatal("expected error for bad JSON")
	}
}

func TestListFromFS_Empty(t *testing.T) {
	fsys := fstest.MapFS{}
	exs, err := ListFromFS(fsys)
	if err != nil {
		t.Fatalf("ListFromFS error: %v", err)
	}
	if len(exs) != 0 {
		t.Errorf("got %d exercises, want 0", len(exs))
	}
}

func TestListFromFS_Multiple(t *testing.T) {
	fsys := fstest.MapFS{
		"02-world/exercise.json": &fstest.MapFile{
			Data: []byte(`{"title":"World","category":"basics","difficulty":1,"topics":[],"hints":[],"files":[]}`),
		},
		"01-hello/exercise.json": &fstest.MapFile{
			Data: []byte(`{"title":"Hello","category":"basics","difficulty":1,"topics":[],"hints":[],"files":[]}`),
		},
	}
	exs, err := ListFromFS(fsys)
	if err != nil {
		t.Fatalf("ListFromFS error: %v", err)
	}
	if len(exs) != 2 {
		t.Fatalf("got %d exercises, want 2", len(exs))
	}
	// should be sorted
	if exs[0].Name != "01-hello" || exs[1].Name != "02-world" {
		t.Errorf("sort order wrong: %v", exs)
	}
}

func TestListFromFS_SkipsNoJSONDir(t *testing.T) {
	fsys := fstest.MapFS{
		"01-hello/exercise.json": &fstest.MapFile{
			Data: []byte(`{"title":"Hello","category":"basics","difficulty":1,"topics":[],"hints":[],"files":[]}`),
		},
		"empty-dir/.keep": &fstest.MapFile{Data: []byte{}},
	}
	exs, err := ListFromFS(fsys)
	if err != nil {
		t.Fatalf("ListFromFS error: %v", err)
	}
	if len(exs) != 1 {
		t.Errorf("got %d exercises, want 1", len(exs))
	}
}

func TestCopyFromFS(t *testing.T) {
	fsys := fstest.MapFS{
		"ex1/starter/main.go.txt": &fstest.MapFile{
			Data: []byte("package main\nfunc main() {}\n"),
		},
		"ex1/starter/helper.txt": &fstest.MapFile{
			Data: []byte("not a go file"),
		},
		"ex1/solution/main.go.txt": &fstest.MapFile{
			Data: []byte("package main\nfunc main() { println() }\n"),
		},
		"ex1/exercise.json": &fstest.MapFile{
			Data: []byte(`{"title":"Ex1","category":"basics","difficulty":1,"topics":[],"hints":[],"files":[]}`),
		},
	}

	dst := t.TempDir()
	if err := CopyFromFS(fsys, "ex1", dst); err != nil {
		t.Fatalf("CopyFromFS error: %v", err)
	}

	// .go.txt should become .go
	if _, err := os.Stat(filepath.Join(dst, "starter", "main.go")); err != nil {
		t.Errorf("main.go not copied: %v", err)
	}
	// .txt suffix is stripped (embed convention), so helper.txt becomes helper
	if _, err := os.Stat(filepath.Join(dst, "starter", "helper")); err != nil {
		t.Errorf("helper not copied: %v", err)
	}
	// solution should also be copied
	if _, err := os.Stat(filepath.Join(dst, "solution", "main.go")); err != nil {
		t.Errorf("solution/main.go not copied: %v", err)
	}
}

func TestVerifyPass(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module test"), 0644)
	os.WriteFile(filepath.Join(dir, "main_test.go"), []byte(`package main
import "testing"
func TestPass(t *testing.T) {}`), 0644)

	result, err := Verify(t.Context(), dir)
	if err != nil {
		t.Fatalf("Verify error: %v", err)
	}
	if !result.Passed {
		t.Errorf("Passed = false, want true. Output: %s", result.Output)
	}
}

func TestVerifyFail(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module test"), 0644)
	os.WriteFile(filepath.Join(dir, "main_test.go"), []byte(`package main
import "testing"
func TestFail(t *testing.T) { t.Error("boom") }`), 0644)

	result, err := Verify(t.Context(), dir)
	if err != nil {
		t.Fatalf("Verify error: %v", err)
	}
	if result.Passed {
		t.Error("Passed = true, want false")
	}
}

func TestVerifyTimeout(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module test"), 0644)
	// A test that sleeps longer than the context deadline
	os.WriteFile(filepath.Join(dir, "slow_test.go"), []byte(`package main
import "testing"
func TestSlow(t *testing.T) { select {} }`), 0644)

	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Millisecond)
	defer cancel()

	result, err := Verify(ctx, dir)
	if err == nil {
		t.Error("expected timeout error, got nil")
	}
	if result.Passed {
		t.Error("Passed = true on timeout, want false")
	}
	if !strings.Contains(result.Errors, "タイムアウト") {
		t.Errorf("Errors missing timeout message: %s", result.Errors)
	}
}

func BenchmarkListFromFS(b *testing.B) {
	fsys := fstest.MapFS{
		"01-hello/exercise.json": &fstest.MapFile{
			Data: []byte(`{"title":"Hello","category":"basics","difficulty":1,"topics":["fmt"],"hints":["use fmt.Println"],"files":["main.go"]}`),
		},
		"02-world/exercise.json": &fstest.MapFile{
			Data: []byte(`{"title":"World","category":"basics","difficulty":2,"topics":["fmt"],"hints":["use fmt.Println"],"files":["main.go"]}`),
		},
		"03-test/exercise.json": &fstest.MapFile{
			Data: []byte(`{"title":"Test","category":"testing","difficulty":3,"topics":["testing"],"hints":[],"files":["main_test.go"]}`),
		},
	}
	b.ResetTimer()
	for range b.N {
		ListFromFS(fsys)
	}
}
