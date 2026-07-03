package cmd

import (
	"fmt"
	"sort"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

func (r *Runner) newGraphCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "graph",
		Short: "カテゴリ別の学習マップを表示",
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
			byCat := map[string][]exercise.Exercise{}
			var cats []string
			for _, ex := range all {
				if _, ok := byCat[ex.Category]; !ok {
					cats = append(cats, ex.Category)
				}
				byCat[ex.Category] = append(byCat[ex.Category], ex)
			}
			sort.Strings(cats)

			cmd.Println("📚 " + colorCyan("学習マップ (カテゴリ別)"))
			cmd.Println()

			for _, c := range cats {
				exs := byCat[c]
				sort.Slice(exs, func(i, j int) bool {
					if exs[i].Difficulty != exs[j].Difficulty {
						return exs[i].Difficulty < exs[j].Difficulty
					}
					return exs[i].Name < exs[j].Name
				})

				done := 0
				for _, ex := range exs {
					if prog.IsCompleted(ex.Name) {
						done++
					}
				}
				cmd.Printf("  %s (%d/%d)\n", colorYellow(c), done, len(exs))
				for _, ex := range exs {
					status := "  "
					if prog.IsCompleted(ex.Name) {
						status = "✅"
					}
					cmd.Printf("    %s %s %s\n", status, stars(ex.Difficulty), ex.Title)
				}
				cmd.Println()
			}

			total, comp := countProgress(all, prog)
			pct := 0
			if total > 0 {
				pct = comp * 100 / total
			}
			cmd.Printf("  合計: %d/%d (%d%%)\n", comp, total, pct)
			return nil
		},
	}
}

func countProgress(all []exercise.Exercise, prog *progress.Data) (total, completed int) {
	for _, ex := range all {
		total++
		if prog.IsCompleted(ex.Name) {
			completed++
		}
	}
	return
}
