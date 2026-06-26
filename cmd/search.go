package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

var searchCategory string

func init() {
	searchCmd.Flags().StringVarP(&searchCategory, "category", "c", "", "カテゴリで絞り込み")
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "演習を検索（名前・タイトル・トピック）",
	Args:  cobra.ExactArgs(1),
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
			fmt.Printf("「%s」に一致する演習は見つかりませんでした。\n", args[0])
			return
		}

		fmt.Printf("「%s」の検索結果: %d 件\n\n", args[0], len(matched))
		for _, ex := range matched {
			status := " "
			if prog.IsCompleted(ex.Name) {
				status = "✅"
			}
			fmt.Printf("  %s %s [%s] %s\n", status, stars(ex.Difficulty), ex.Category, ex.Title)
			fmt.Printf("        → %s\n", colorCyan("go-practicum start "+ex.Name))
		}
	},
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
