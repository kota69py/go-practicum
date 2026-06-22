package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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
		if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err != nil {
			fmt.Fprintln(os.Stderr, "エラー: カレントディレクトリに go.mod が見つかりません。演習ディレクトリで実行してください。")
			os.Exit(1)
		}

		result, err := exercise.Verify(cwd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラー: 検証失敗: %v\n", err)
			os.Exit(1)
		}

		if result.Passed {
			fmt.Println("✅ " + colorGreen("全テスト通過！"))
			prog.Complete(prog.InProgress)
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
