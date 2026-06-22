package cmd

import (
	"fmt"
	"os"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "演習一覧を表示",
	Run: func(cmd *cobra.Command, args []string) {
		if exercFS == nil {
			fmt.Fprintln(os.Stderr, "エラー: 演習データが見つかりません")
			os.Exit(1)
		}
		exs, err := exercise.ListFromFS(exercFS)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
			os.Exit(1)
		}

		prog, _ := progress.Load()

		if len(exs) == 0 {
			fmt.Println("演習が見つかりませんでした。")
			return
		}

		fmt.Println("利用可能な演習:")
		fmt.Println()
		for _, ex := range exs {
			status := " "
			if prog.IsCompleted(ex.Name) {
				status = "✅"
			}
			difficulty := ""
			for i := 0; i < ex.Difficulty; i++ {
				difficulty += "★"
			}
			fmt.Printf("  %s  %s  [%s]  %s\n", status, difficulty, ex.Category, ex.Title)
			fmt.Printf("      go-practicum start %s\n", ex.Name)
			fmt.Println()
		}
	},
}
