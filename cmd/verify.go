package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

func (r *Runner) newVerifyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "verify",
		Short: "演習を検証（go test を実行）",
		Run: func(cmd *cobra.Command, args []string) {
			prog, _ := progress.Load()
			if prog.InProgress == "" {
				fmt.Fprintln(os.Stderr, "エラー: 進行中の演習がありません（go-practicum start <name> で開始してください）")
				os.Exit(1)
			}

			ex, _ := exercise.LoadFromFS(r.exercFS, prog.InProgress)

			cwd, _ := os.Getwd()
			if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err != nil {
				fmt.Fprintf(os.Stderr, "エラー: カレントディレクトリに go.mod が見つかりません。\n")
				if ex != nil {
					fmt.Fprintf(os.Stderr, "ヒント: 演習 %q のディレクトリで実行していますか？\n", ex.Title)
				}
				fmt.Fprintln(os.Stderr, "ヒント: go-practicum start で展開されたディレクトリで実行してください。")
				os.Exit(1)
			}

			if ex != nil {
				fmt.Printf("🔍 %s を検証中...\n", colorCyan(ex.Title))
			} else {
				fmt.Printf("🔍 %s を検証中...\n", colorCyan(prog.InProgress))
			}

			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()

			result, err := exercise.Verify(ctx, cwd)
			if err != nil {
				fmt.Fprintf(os.Stderr, "エラー: 検証失敗: %v\n", err)
				os.Exit(1)
			}

			if result.Passed {
				fmt.Println("✅ " + colorGreen("全テスト通過！"))
				prog.Complete(prog.InProgress)
				prog.InProgress = ""
				prog.Save()
			} else {
				fmt.Println("❌ " + colorRed("テスト失敗"))
			}

			if result.Output != "" {
				fmt.Println()
				fmt.Println(result.Output)
			}
			if result.Errors != "" {
				fmt.Fprintln(os.Stderr, result.Errors)
			}
		},
	}
}
