package cmd

import (
	"fmt"
	"strings"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

var (
	listCategory   string
	listDifficulty int
)

func (r *Runner) newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "演習一覧を表示",
		Args:  cobra.NoArgs,
		RunE: func(c *cobra.Command, args []string) error {
		if r.exercFS == nil {
			return fmt.Errorf("演習データが見つかりません")
		}
		all, err := exercise.ListFromFS(r.exercFS)
		if err != nil {
			return fmt.Errorf("%v", err)
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
			c.Println("条件に一致する演習がありません。")
			return nil
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
		c.Printf("進捗: %d/%d (%d%%)  カテゴリ数: %d\n", completed, len(all), pct, countCategories(all))
		if listCategory != "" || listDifficulty > 0 {
			c.Printf("表示: %d 件\n", len(exs))
		}
		c.Println()

		for _, ex := range exs {
			status := " "
			if prog.IsCompleted(ex.Name) {
				status = "✅"
			}
			c.Printf("  %s %s [%s] %s\n", status, stars(ex.Difficulty), ex.Category, ex.Title)
			c.Printf("       %s\n", colorCyan("go-practicum start "+ex.Name))
		}
		return nil
	},
	}
	cmd.Flags().StringVarP(&listCategory, "category", "c", "", "カテゴリでフィルタ (例: concurrency, testing)")
	cmd.Flags().IntVarP(&listDifficulty, "difficulty", "d", 0, "難易度でフィルタ (1-5)")
	return cmd
}

func countCategories(exs []exercise.Exercise) int {
	seen := map[string]bool{}
	for _, ex := range exs {
		seen[ex.Category] = true
	}
	return len(seen)
}
