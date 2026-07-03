package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

func (r *Runner) newCheckCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "check",
		Short: "現在の演習コードを静的解析 (go vet / gofmt)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			prog, _ := progress.Load()
			if prog.InProgress == "" {
				return fmt.Errorf("進行中の演習がありません")
			}

			ex, err := exercise.LoadFromFS(r.exercFS, prog.InProgress)
			if err != nil {
				return fmt.Errorf("演習 %q のデータが見つかりません", prog.InProgress)
			}

			cwd, _ := os.Getwd()
			if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err != nil {
				cmd.PrintErrf("エラー: カレントディレクトリに go.mod が見つかりません。\n")
				cmd.PrintErrf("ヒント: 演習 %q のディレクトリで実行していますか？\n", ex.Title)
				return fmt.Errorf("go.mod not found")
			}

			cmd.Printf("🔍 %s をチェック中...\n\n", colorCyan(ex.Title))
			hasIssues := false

			// gofmt check
			cmd.Println("🔍 " + colorCyan("gofmt チェック..."))
			var fmtOut bytes.Buffer
			fmtCmd := exec.Command("gofmt", "-l", ".")
			fmtCmd.Dir = cwd
			fmtCmd.Stdout = &fmtOut
			fmtCmd.Stderr = &fmtOut
			fmtCmd.Run()
			if strings.TrimSpace(fmtOut.String()) != "" {
				cmd.Printf("  ❌ フォーマットが必要なファイル:\n")
				for _, f := range strings.Split(strings.TrimSpace(fmtOut.String()), "\n") {
					cmd.Printf("    - %s\n", f)
				}
				hasIssues = true
			} else {
				cmd.Println("  ✅ フォーマットは適切です")
			}

			// go vet check
			cmd.Println()
			cmd.Println("🔍 " + colorCyan("go vet 静的解析..."))
			var vetOut, vetErr bytes.Buffer
			vetCmd := exec.Command("go", "vet", "./...")
			vetCmd.Dir = cwd
			vetCmd.Stdout = &vetOut
			vetCmd.Stderr = &vetErr
			if err := vetCmd.Run(); err != nil {
				cmd.Println("  ❌ 問題が見つかりました:")
				if s := strings.TrimSpace(vetOut.String()); s != "" {
					cmd.Println(s)
				}
				if s := strings.TrimSpace(vetErr.String()); s != "" {
					cmd.Println(s)
				}
				hasIssues = true
			} else {
				cmd.Println("  ✅ 静的解析を通過しました")
			}

			// Summary
			cmd.Println()
			if hasIssues {
				return fmt.Errorf("%s に修正が必要です", colorRed(ex.Title))
			}
			cmd.Println("✅ " + colorGreen("すべてのチェックを通過しました"))
			return nil
		},
	}
}
