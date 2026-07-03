package cmd

import (
	"fmt"
	"sort"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

func (r *Runner) newStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "学習進捗を表示",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if r.exercFS == nil {
				return fmt.Errorf("演習データが見つかりません")
			}
			all, err := exercise.ListFromFS(r.exercFS)
			if err != nil {
				return fmt.Errorf("%v", err)
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

			cmd.Printf("進捗: %s\n", colorGreen(fmt.Sprintf("%d/%d (%d%%)", completed, len(all), pct)))
			cmd.Println()

			if prog.InProgress != "" {
				ex, err := exercise.LoadFromFS(r.exercFS, prog.InProgress)
				if err == nil {
					cmd.Printf("進行中: %s  %s [%s] %s\n", stars(ex.Difficulty), ex.Category, colorCyan(ex.Title), prog.InProgress)
					cmd.Printf("         %s\n", colorCyan("go-practicum hint"))
					cmd.Printf("         %s\n", colorCyan("go-practicum solution"))
					cmd.Println()
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
				cmd.Printf("  %s %s %d/%d\n", bar, c, done, total)
			}
			return nil
		},
	}
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
