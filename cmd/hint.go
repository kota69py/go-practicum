package cmd

import (
	"fmt"
	"os"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(hintCmd)
}

var hintCmd = &cobra.Command{
	Use:   "hint",
	Short: "ヒントを表示",
	Run: func(cmd *cobra.Command, args []string) {
		prog, _ := progress.Load()
		if prog.InProgress == "" {
			fmt.Fprintln(os.Stderr, "エラー: 進行中の演習がありません")
			os.Exit(1)
		}
		if exercFS == nil {
			fmt.Fprintln(os.Stderr, "エラー: 演習データが見つかりません")
			os.Exit(1)
		}
		ex, err := exercise.LoadFromFS(exercFS, prog.InProgress)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラー: 演習 %q が見つかりません\n", prog.InProgress)
			os.Exit(1)
		}

		if len(ex.Hints) == 0 {
			fmt.Println("ヒントは用意されていません。")
			return
		}

		fmt.Println("💡 " + colorYellow("ヒント:"))
		fmt.Println()
		for i, h := range ex.Hints {
			fmt.Printf("  %d. %s\n", i+1, h)
		}
	},
}
