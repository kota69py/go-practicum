package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "学習進捗を表示",
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

		prog, _ := progress.Load()

		completed := 0
		catCount := map[string]int{}
		catDone := map[string]int{}
		for _, ex := range all {
			catCount[ex.Category]++
			if prog.IsCompleted(ex.Name) {
				completed++
				catDone[ex.Category]++
			}
		}

		pct := 0
		if len(all) > 0 {
			pct = completed * 100 / len(all)
		}

		fmt.Printf("進捗: %s\n", colorGreen(fmt.Sprintf("%d/%d (%d%%)", completed, len(all), pct)))
		fmt.Println()

		if prog.InProgress != "" {
			ex, err := exercise.LoadFromFS(exercFS, prog.InProgress)
			if err == nil {
				fmt.Printf("進行中: %s  %s [%s] %s\n", stars(ex.Difficulty), ex.Category, colorCyan(ex.Title), prog.InProgress)
				fmt.Printf("         %s\n", colorCyan("go-practicum hint"))
				fmt.Printf("         %s\n", colorCyan("go-practicum solution"))
				fmt.Println()
			}
		}

		// Category breakdown
		var cats []string
		for c := range catCount {
			cats = append(cats, c)
		}
		sort.Strings(cats)
		for _, c := range cats {
			total := catCount[c]
			done := catDone[c]
			bar := progressBar(done, total, 10)
			fmt.Printf("  %s %s %d/%d\n", bar, c, done, total)
		}
	},
}

func progressBar(done, total, width int) string {
	if total == 0 {
		return ""
	}
	filled := done * width / total
	s := "["
	for i := 0; i < width; i++ {
		if i < filled {
			s += "="
		} else {
			s += " "
		}
	}
	s += "]"
	return s
}
