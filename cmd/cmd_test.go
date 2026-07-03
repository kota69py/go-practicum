package cmd

import (
	"strings"
	"testing"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
)
func TestStars(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{0, "☆☆☆☆☆"},
		{1, "★☆☆☆☆"},
		{3, "★★★☆☆"},
		{5, "★★★★★"},
		{7, "★★★★★"},
		{-1, "☆☆☆☆☆"},
	}
	for _, tt := range tests {
		got := stars(tt.n)
		if got != tt.want {
			t.Errorf("stars(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}

func TestColorGreen(t *testing.T) {
	useColor = true
	got := colorGreen("ok")
	if !strings.HasPrefix(got, "\033[32m") || !strings.HasSuffix(got, "\033[0m") {
		t.Errorf("colorGreen = %q, want ANSI wrapped", got)
	}
	useColor = false
	got = colorGreen("ok")
	if got != "ok" {
		t.Errorf("colorGreen(no color) = %q, want %q", got, "ok")
	}
	useColor = true
}

func TestColorRed(t *testing.T) {
	useColor = true
	got := colorRed("err")
	if !strings.HasPrefix(got, "\033[31m") {
		t.Errorf("colorRed = %q, wants red ANSI", got)
	}
}

func TestConfirm(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"y\n", true},
		{"Y\n", true},
		{"n\n", false},
		{"N\n", false},
		{"\n", false},
		{"", false},
	}
	for _, tt := range tests {
		r := strings.NewReader(tt.input)
		got := confirm("test?", r)
		if got != tt.want {
			t.Errorf("confirm(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestCountCategories(t *testing.T) {
	exercises := []exercise.Exercise{
		{Name: "01-a", Category: "concurrency"},
		{Name: "02-b", Category: "concurrency"},
		{Name: "03-c", Category: "testing"},
		{Name: "04-d", Category: "io"},
	}
	if n := countCategories(exercises); n != 3 {
		t.Errorf("countCategories = %d, want 3", n)
	}
}

func TestCountProgress(t *testing.T) {
	all := []exercise.Exercise{
		{Name: "01-a"},
		{Name: "02-b"},
		{Name: "03-c"},
	}
	prog := &progress.Data{Completed: []string{"01-a", "03-c"}}
	total, done := countProgress(all, prog)
	if total != 3 {
		t.Errorf("total = %d, want 3", total)
	}
	if done != 2 {
		t.Errorf("done = %d, want 2", done)
	}
}

func TestMatchExercise(t *testing.T) {
	ex := exercise.Exercise{
		Name:     "01-goroutine-basics",
		Title:    "Goroutine 基礎",
		Category: "concurrency",
		Topics:   []string{"goroutine", "channel", "sync"},
	}
	tests := []struct {
		query    string
		expected bool
	}{
		{"goroutine", true},
		{"GOROUTINE", true},
		{"concurrency", true},
		{"Goroutine 基礎", true},
		{"sync", true},
		{"grpc", false},
		{"", false},
	}
	for _, tt := range tests {
		got := matchExercise(ex, tt.query)
		if got != tt.expected {
			t.Errorf("match(%q, %q) = %v, want %v", ex.Name, tt.query, got, tt.expected)
		}
	}
}

func BenchmarkMatchExercise(b *testing.B) {
	ex := exercise.Exercise{
		Name:     "01-goroutine-basics",
		Title:    "Goroutine 基礎",
		Category: "concurrency",
		Topics:   []string{"goroutine", "channel", "sync"},
	}
	b.ResetTimer()
	for range b.N {
		matchExercise(ex, "goroutine")
	}
}

func BenchmarkStars(b *testing.B) {
	for range b.N {
		stars(3)
	}
}

func TestLevenshtein(t *testing.T) {
	tests := []struct {
		a, b string
		want int
	}{
		{"", "", 0},
		{"abc", "", 3},
		{"", "abc", 3},
		{"abc", "abc", 0},
		{"abc", "abd", 1},
		{"interfce-design", "interface-design", 1},
		{"01-interfce-design", "01-interface-design", 1},
		{"gorutine", "goroutine", 1},
	}
	for _, tt := range tests {
		got := exercise.Levenshtein(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("levenshtein(%q, %q) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
