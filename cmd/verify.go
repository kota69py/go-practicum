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
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			prog, _ := progress.Load()
			if prog.InProgress == "" {
				return fmt.Errorf("進行中の演習がありません（go-practicum start <name> で開始してください）")
			}

			ex, _ := exercise.LoadFromFS(r.exercFS, prog.InProgress)

			cwd, _ := os.Getwd()
			if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err != nil {
				cmd.PrintErrf("エラー: カレントディレクトリに go.mod が見つかりません。\n")
				if ex != nil {
					cmd.PrintErrf("ヒント: 演習 %q のディレクトリで実行していますか？\n", ex.Title)
				}
				cmd.PrintErrln("ヒント: go-practicum start で展開されたディレクトリで実行してください。")
				return fmt.Errorf("go.mod not found")
			}

			if ex != nil {
				cmd.Printf("🔍 %s を検証中...\n", colorCyan(ex.Title))
			} else {
				cmd.Printf("🔍 %s を検証中...\n", colorCyan(prog.InProgress))
			}

			ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
			defer cancel()

			result, err := exercise.Verify(ctx, cwd)
			if err != nil {
				return fmt.Errorf("検証失敗: %v", err)
			}

			if result.Passed {
				cmd.Println("✅ " + colorGreen("全テスト通過！"))
				prog.Complete(prog.InProgress)
				prog.InProgress = ""
				_ = progress.Save(prog)
			} else {
				cmd.Println("❌ " + colorRed("テスト失敗"))
			}

			if result.Output != "" {
				cmd.Println()
				cmd.Println(result.Output)
			}
			if result.Errors != "" {
				cmd.PrintErrln(result.Errors)
			}
			return nil
		},
	}
}
