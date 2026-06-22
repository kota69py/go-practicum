package exercise

import (
	"bytes"
	"os/exec"
)

type VerifyResult struct {
	Passed bool
	Output string
	Errors string
}

func Verify(workDir string) (*VerifyResult, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("go", "test", "-v", "./...")
	cmd.Dir = workDir
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	return &VerifyResult{
		Passed: err == nil,
		Output: stdout.String(),
		Errors: stderr.String(),
	}, nil
}
