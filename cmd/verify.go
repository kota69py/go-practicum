package cmd

import (
	"fmt"
	"os"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(verifyCmd)
}

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "演習を検証（go test を実行）",
	Run: func(cmd *cobra.Command, args []string) {
		prog, _ := progress.Load()
		if prog.InProgress == "" {
			fmt.Fprintln(os.Stderr, "エラー: 進行中の演習がありません（go-practicum start <name> で開始してください）")
			os.Exit(1)
		}

		cwd, _ := os.Getwd()
		result, err := exercise.Verify(cwd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラー: 検証失敗: %v\n", err)
			os.Exit(1)
		}

		if result.Passed {
			fmt.Println("✅ 全テスト通過！")
			prog.Complete(prog.InProgress)
			prog.Save()
		} else {
			fmt.Println("❌ テスト失敗")
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
