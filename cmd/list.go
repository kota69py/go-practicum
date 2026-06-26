package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

var (
	listCategory   string
	listDifficulty int
)

func init() {
	listCmd.Flags().StringVarP(&listCategory, "category", "c", "", "カテゴリでフィルタ (例: concurrency, testing)")
	listCmd.Flags().IntVarP(&listDifficulty, "difficulty", "d", 0, "難易度でフィルタ (1-5)")
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
		all, err := exercise.ListFromFS(exercFS)
		if err != nil {
			fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
			os.Exit(1)
		}

		// Filter
		var exs []exercise.Exercise
		for _, ex := range all {
			if listCategory != "" && !strings.EqualFold(ex.Category, listCategory) {
				continue
			}
			if listDifficulty > 0 && ex.Difficulty != listDifficulty {
				continue
			}
			exs = append(exs, ex)
		}

		prog, _ := progress.Load()

		if len(exs) == 0 {
			fmt.Println("条件に一致する演習がありません。")
			return
		}

		// Progress summary
		completed := 0
		for _, ex := range all {
			if prog.IsCompleted(ex.Name) {
				completed++
			}
		}
		pct := 0
		if len(all) > 0 {
			pct = completed * 100 / len(all)
		}
		fmt.Printf("進捗: %d/%d (%d%%)  カテゴリ数: %d\n", completed, len(all), pct, countCategories(all))
		if listCategory != "" || listDifficulty > 0 {
			fmt.Printf("表示: %d 件\n", len(exs))
		}
		fmt.Println()

		for _, ex := range exs {
			status := " "
			if prog.IsCompleted(ex.Name) {
				status = "✅"
			}
			fmt.Printf("  %s %s [%s] %s\n", status, stars(ex.Difficulty), ex.Category, ex.Title)
			fmt.Printf("       %s\n", colorCyan("go-practicum start "+ex.Name))
		}
	},
}

func countCategories(exs []exercise.Exercise) int {
	seen := map[string]bool{}
	for _, ex := range exs {
		seen[ex.Category] = true
	}
	return len(seen)
}
