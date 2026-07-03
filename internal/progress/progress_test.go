package progress

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func testHomeDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	t.Setenv("USERPROFILE", dir)
	return dir
}

func TestLoadEmpty(t *testing.T) {
	testHomeDir(t)
	d, err := Load()
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if d == nil {
		t.Fatal("Load() returned nil")
	}
	if len(d.Completed) != 0 {
		t.Errorf("Completed = %v, want empty", d.Completed)
	}
	if d.InProgress != "" {
		t.Errorf("InProgress = %q, want empty", d.InProgress)
	}
}

func TestSaveAndLoad(t *testing.T) {
	testHomeDir(t)

	d := &Data{Completed: []string{"ex1", "ex2"}, InProgress: "ex3"}
	if err := Save(d); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	d2, err := Load()
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if len(d2.Completed) != 2 {
		t.Errorf("Completed count = %d, want 2", len(d2.Completed))
	}
	if d2.InProgress != "ex3" {
		t.Errorf("InProgress = %q, want %q", d2.InProgress, "ex3")
	}
}

func TestLoadCorruptedJSON(t *testing.T) {
	dir := testHomeDir(t)
	os.MkdirAll(filepath.Join(dir, ".go-practicum"), 0755)
	os.WriteFile(filepath.Join(dir, ".go-practicum", "progress.json"), []byte("{invalid"), 0644)

	d, err := Load()
	if err != nil {
		t.Fatalf("Load() on corrupted file should not error: %v", err)
	}
	if d == nil {
		t.Fatal("Load() returned nil")
	}
}

func TestComplete(t *testing.T) {
	d := &Data{}
	d.Complete("ex1")
	if len(d.Completed) != 1 {
		t.Fatalf("Completed count = %d, want 1", len(d.Completed))
	}
	d.Complete("ex1") // duplicate
	if len(d.Completed) != 1 {
		t.Errorf("duplicate Complete added, count = %d", len(d.Completed))
	}
	d.Complete("ex2")
	if len(d.Completed) != 2 {
		t.Errorf("Completed count = %d, want 2", len(d.Completed))
	}
}

func TestIsCompleted(t *testing.T) {
	d := &Data{Completed: []string{"ex1", "ex3"}}
	if !d.IsCompleted("ex1") {
		t.Error("IsCompleted('ex1') = false, want true")
	}
	if d.IsCompleted("ex2") {
		t.Error("IsCompleted('ex2') = true, want false")
	}
	if !d.IsCompleted("ex3") {
		t.Error("IsCompleted('ex3') = false, want true")
	}
}

func TestLoadCorruptedJSON_WithInProgress(t *testing.T) {
	dir := testHomeDir(t)
	os.MkdirAll(filepath.Join(dir, ".go-practicum"), 0755)
	// Corrupted JSON that also has an in_progress field
	os.WriteFile(filepath.Join(dir, ".go-practicum", "progress.json"), []byte(`{"completed":["ex1"],"in_progress":"ex2"`), 0644)

	d, err := Load()
	if err != nil {
		t.Fatalf("Load() on corrupted file should return empty: %v", err)
	}
	if d == nil {
		t.Fatal("Load() returned nil")
	}
	if len(d.Completed) != 0 {
		t.Errorf("Completed = %v on corrupted JSON, want empty", d.Completed)
	}
	if d.InProgress != "" {
		t.Errorf("InProgress = %q on corrupted JSON, want empty", d.InProgress)
	}
}

func BenchmarkIsCompleted(b *testing.B) {
	names := make([]string, 80)
	for i := range names {
		names[i] = fmt.Sprintf("ex-%02d", i)
	}
	d := &Data{Completed: names}
	b.ResetTimer()
	for range b.N {
		d.IsCompleted("ex-79")
	}
}
