package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/spf13/cobra"
)

func (r *Runner) newInfoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "info <name>",
		Short: "演習の詳細を表示",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			if r.exercFS == nil {
				fmt.Fprintln(os.Stderr, "エラー: 演習データが見つかりません")
				os.Exit(1)
			}
			ex, err := exercise.LoadFromFS(r.exercFS, name)
			if err != nil {
				fmt.Fprintf(os.Stderr, "エラー: 演習 %q が見つかりません\n", name)
				os.Exit(1)
			}

			fmt.Printf("  %s\n", colorCyan(ex.Title))
			fmt.Println()
			fmt.Printf("  名前:     %s\n", ex.Name)
			fmt.Printf("  カテゴリ: %s\n", ex.Category)
			fmt.Printf("  難易度:   %s\n", stars(ex.Difficulty))
			if len(ex.Topics) > 0 {
				fmt.Printf("  トピック: %s\n", strings.Join(ex.Topics, ", "))
			}
			if len(ex.Files) > 0 {
				var files []string
				for _, f := range ex.Files {
					files = append(files, strings.TrimSuffix(f, ".txt"))
				}
				fmt.Printf("  ファイル: %s\n", strings.Join(files, ", "))
			}
			if len(ex.Hints) > 0 {
				fmt.Println()
				fmt.Println("  ヒント:")
				for i, h := range ex.Hints {
					fmt.Printf("    %d. %s\n", i+1, h)
				}
			}
			fmt.Println()
			fmt.Printf("  開始: %s\n", colorCyan("go-practicum start "+ex.Name))
		},
	}
}
