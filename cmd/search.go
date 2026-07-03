package cmd

import (
	"fmt"
	"strings"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

var searchCategory string

func (r *Runner) newSearchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search <query>",
		Short: "演習を検索（名前・タイトル・トピック）",
		Args:  cobra.ExactArgs(1),
		RunE: func(c *cobra.Command, args []string) error {
			if r.exercFS == nil {
				return fmt.Errorf("演習データが見つかりません")
			}
			all, err := exercise.ListFromFS(r.exercFS)
			if err != nil {
				return fmt.Errorf("%v", err)
			}

			query := strings.ToLower(args[0])
			prog, _ := progress.Load()
			var matched []exercise.Exercise

			for _, ex := range all {
				if searchCategory != "" && !strings.EqualFold(ex.Category, searchCategory) {
					continue
				}
				if matchExercise(ex, query) {
					matched = append(matched, ex)
				}
			}

			if len(matched) == 0 {
				c.Printf("「%s」に一致する演習は見つかりませんでした。\n", args[0])
				return nil
			}

			c.Printf("「%s」の検索結果: %d 件\n\n", args[0], len(matched))
			for _, ex := range matched {
				status := " "
				if prog.IsCompleted(ex.Name) {
					status = "✅"
				}
				c.Printf("  %s %s [%s] %s\n", status, stars(ex.Difficulty), ex.Category, ex.Title)
				c.Printf("        → %s\n", colorCyan("go-practicum start "+ex.Name))
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&searchCategory, "category", "c", "", "カテゴリで絞り込み")
	return cmd
}

func matchExercise(ex exercise.Exercise, query string) bool {
	if query == "" {
		return false
	}
	q := strings.ToLower(query)
	if strings.Contains(strings.ToLower(ex.Name), q) ||
		strings.Contains(strings.ToLower(ex.Title), q) ||
		strings.Contains(strings.ToLower(ex.Category), q) {
		return true
	}
	for _, t := range ex.Topics {
		if strings.Contains(strings.ToLower(t), q) {
			return true
		}
	}
	return false
}
