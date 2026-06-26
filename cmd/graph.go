package cmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

func (r *Runner) newGraphCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "graph",
		Short: "カテゴリ別の学習マップを表示",
		Run: func(cmd *cobra.Command, args []string) {
			if r.exercFS == nil {
				fmt.Fprintln(os.Stderr, "エラー: 演習データが見つかりません")
				os.Exit(1)
			}
			all, err := exercise.ListFromFS(r.exercFS)
			if err != nil {
				fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
				os.Exit(1)
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

			fmt.Println("📚 " + colorCyan("学習マップ (カテゴリ別)"))
			fmt.Println()

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
				fmt.Printf("  %s (%d/%d)\n", colorYellow(c), done, len(exs))
				for _, ex := range exs {
					status := "  "
					if prog.IsCompleted(ex.Name) {
						status = "✅"
					}
					fmt.Printf("    %s %s %s\n", status, stars(ex.Difficulty), ex.Title)
				}
				fmt.Println()
			}

			total, comp := countProgress(all, prog)
			pct := 0
			if total > 0 {
				pct = comp * 100 / total
			}
			fmt.Printf("  合計: %d/%d (%d%%)\n", comp, total, pct)
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
