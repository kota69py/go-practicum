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

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "現在の演習コードを静的解析 (go vet / gofmt)",
	Run: func(cmd *cobra.Command, args []string) {
		prog, _ := progress.Load()
		if prog.InProgress == "" {
			fmt.Fprintln(os.Stderr, "エラー: 進行中の演習がありません")
			os.Exit(1)
		}

		cwd, _ := os.Getwd()
		if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err != nil {
			fmt.Fprintln(os.Stderr, "エラー: カレントディレクトリに go.mod が見つかりません")
			os.Exit(1)
		}

		ex, _ := exercise.LoadFromFS(exercFS, prog.InProgress)
		hasIssues := false

		// gofmt check
		fmt.Println("🔍 " + colorCyan("gofmt チェック..."))
		var fmtOut bytes.Buffer
		fmtCmd := exec.Command("gofmt", "-l", ".")
		fmtCmd.Dir = cwd
		fmtCmd.Stdout = &fmtOut
		fmtCmd.Stderr = &fmtOut
		if err := fmtCmd.Run(); err == nil && strings.TrimSpace(fmtOut.String()) != "" {
			fmt.Printf("  ❌ フォーマットが必要なファイル:\n")
			for _, f := range strings.Split(strings.TrimSpace(fmtOut.String()), "\n") {
				fmt.Printf("    - %s\n", f)
			}
			hasIssues = true
		} else {
			fmt.Println("  ✅ フォーマットは適切です")
		}

		// go vet check
		fmt.Println()
		fmt.Println("🔍 " + colorCyan("go vet 静的解析..."))
		var vetOut, vetErr bytes.Buffer
		vetCmd := exec.Command("go", "vet", "./...")
		vetCmd.Dir = cwd
		vetCmd.Stdout = &vetOut
		vetCmd.Stderr = &vetErr
		if err := vetCmd.Run(); err != nil {
			fmt.Println("  ❌ 問題が見つかりました:")
			if s := strings.TrimSpace(vetOut.String()); s != "" {
				fmt.Println(s)
			}
			if s := strings.TrimSpace(vetErr.String()); s != "" {
				fmt.Println(s)
			}
			hasIssues = true
		} else {
			fmt.Println("  ✅ 静的解析を通過しました")
		}

		// Summary
		fmt.Println()
		if hasIssues {
			if ex != nil {
				fmt.Printf("❌ %s に修正が必要です\n", colorRed(ex.Title))
			}
			os.Exit(1)
		} else {
			fmt.Println("✅ " + colorGreen("すべてのチェックを通過しました"))
		}
	},
}
