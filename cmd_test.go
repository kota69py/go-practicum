package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/exercdata"
)

func buildTestBinary(t *testing.T) string {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "go-practicum-test.exe")
	cmd := exec.Command("go", "build", "-o", bin, ".")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build failed: %v\n%s", err, out)
	}
	return bin
}

func TestBinaryList(t *testing.T) {
	bin := buildTestBinary(t)
	out, err := exec.Command(bin, "list").CombinedOutput()
	if err != nil {
		t.Fatalf("list failed: %v\n%s", err, out)
	}
	output := string(out)
	if !strings.Contains(output, "01-interface-design") {
		t.Errorf("output missing '01-interface-design':\n%s", output)
	}
	if !strings.Contains(output, "go-practicum start") {
		t.Errorf("output missing 'go-practicum start':\n%s", output)
	}
}

func TestBinaryStartUnknown(t *testing.T) {
	bin := buildTestBinary(t)
	cmd := exec.Command(bin, "start", "nonexistent")
	out, _ := cmd.CombinedOutput()
	output := string(out)
	if !strings.Contains(output, "見つかりません") {
		t.Errorf("expected error message, got:\n%s", output)
	}
}

func TestBinaryVerifyNoProgress(t *testing.T) {
	bin := buildTestBinary(t)
	// Override USERPROFILE to a clean temp dir (no progress.json)
	dir := t.TempDir()
	cmd := exec.Command(bin, "verify")
	cmd.Env = append(os.Environ(), "USERPROFILE="+dir)
	out, _ := cmd.CombinedOutput()
	output := string(out)
	if !strings.Contains(output, "進行中") {
		t.Errorf("expected '進行中' error, got:\n%s", output)
	}
}

func TestBinaryStartAndVerify(t *testing.T) {
	bin := buildTestBinary(t)
	homeDir := t.TempDir()
	workdir := t.TempDir()

	// start 01-interface-design
	cmd := exec.Command(bin, "start", "01-interface-design")
	cmd.Dir = workdir
	cmd.Env = append(os.Environ(), "USERPROFILE="+homeDir)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("start failed: %v\n%s", err, out)
	}
	if !strings.Contains(string(out), "インターフェース設計") {
		t.Errorf("start output missing exercise title:\n%s", out)
	}

	// Apply solution over starter so verify passes
	fsys := exercdata.FS()
	if err := exercise.CopyFromFS(fsys, "01-interface-design/solution", workdir); err != nil {
		t.Fatalf("copy solution failed: %v", err)
	}

	// verify should pass with solution applied
	cmd = exec.Command(bin, "verify")
	cmd.Dir = workdir
	cmd.Env = append(os.Environ(), "USERPROFILE="+homeDir)
	out, err = cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("verify failed: %v\n%s", err, out)
	}
	if !strings.Contains(string(out), "全テスト通過") {
		t.Errorf("verify output missing success:\n%s", out)
	}
}

func TestBinaryHintNoProgress(t *testing.T) {
	bin := buildTestBinary(t)
	dir := t.TempDir()
	cmd := exec.Command(bin, "hint")
	cmd.Env = append(os.Environ(), "USERPROFILE="+dir)
	out, _ := cmd.CombinedOutput()
	output := string(out)
	if !strings.Contains(output, "進行中") {
		t.Errorf("expected '進行中' error, got:\n%s", output)
	}
}

func TestBinarySolutionNoProgress(t *testing.T) {
	bin := buildTestBinary(t)
	dir := t.TempDir()
	cmd := exec.Command(bin, "solution")
	cmd.Env = append(os.Environ(), "USERPROFILE="+dir)
	out, _ := cmd.CombinedOutput()
	output := string(out)
	if !strings.Contains(output, "進行中") {
		t.Errorf("expected '進行中' error, got:\n%s", output)
	}
}
