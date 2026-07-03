package cmd

import (
	"fmt"
	"strings"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/spf13/cobra"
)

func (r *Runner) newInfoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "info <name>",
		Short: "演習の詳細を表示",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			if r.exercFS == nil {
				return fmt.Errorf("演習データが見つかりません")
			}
			ex, err := exercise.LoadFromFS(r.exercFS, name)
			if err != nil {
				if hints := exercise.SuggestNames(r.exercFS, name, 3); len(hints) > 0 {
					return fmt.Errorf("演習 %q が見つかりません\n  もしかして: %s", name, strings.Join(hints, ", "))
				}
				return fmt.Errorf("演習 %q が見つかりません", name)
			}

			cmd.Printf("  %s\n", colorCyan(ex.Title))
			cmd.Println()
			cmd.Printf("  名前:     %s\n", ex.Name)
			cmd.Printf("  カテゴリ: %s\n", ex.Category)
			cmd.Printf("  難易度:   %s\n", stars(ex.Difficulty))
			if len(ex.Topics) > 0 {
				cmd.Printf("  トピック: %s\n", strings.Join(ex.Topics, ", "))
			}
			if len(ex.Files) > 0 {
				var files []string
				for _, f := range ex.Files {
					files = append(files, strings.TrimSuffix(f, ".txt"))
				}
				cmd.Printf("  ファイル: %s\n", strings.Join(files, ", "))
			}
			if len(ex.Hints) > 0 {
				cmd.Println()
				cmd.Println("  ヒント:")
				for i, h := range ex.Hints {
					cmd.Printf("    %d. %s\n", i+1, h)
				}
			}
			cmd.Println()
			cmd.Printf("  開始: %s\n", colorCyan("go-practicum start "+ex.Name))
			return nil
		},
	}
}
