package exercise

import (
	"bytes"
	"context"
	"os/exec"
)

type VerifyResult struct {
	Passed bool
	Output string
	Errors string
}

type TestRunner interface {
	Run(ctx context.Context, workDir string) (*VerifyResult, error)
}

type GoTestRunner struct{}

func (r *GoTestRunner) Run(ctx context.Context, workDir string) (*VerifyResult, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.CommandContext(ctx, "go", "test", "-v", "./...")
	cmd.Dir = workDir
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if ctx.Err() != nil {
		return &VerifyResult{
			Passed: false,
			Output: stdout.String(),
			Errors: "検証がタイムアウトしました（制限時間: 60秒）",
		}, ctx.Err()
	}
	return &VerifyResult{
		Passed: err == nil,
		Output: stdout.String(),
		Errors: stderr.String(),
	}, nil
}

var defaultRunner TestRunner = &GoTestRunner{}

func Verify(ctx context.Context, workDir string) (*VerifyResult, error) {
	return defaultRunner.Run(ctx, workDir)
}
