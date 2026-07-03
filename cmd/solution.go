package cmd

import (
	"fmt"
	"io/fs"
	"strings"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

func (r *Runner) newSolutionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "solution",
		Short: "解答例を表示",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			prog, _ := progress.Load()
			if prog.InProgress == "" {
				return fmt.Errorf("進行中の演習がありません")
			}
			if r.exercFS == nil {
				return fmt.Errorf("解答データが見つかりません")
			}
			ex, err := exercise.LoadFromFS(r.exercFS, prog.InProgress)
			if err != nil {
				return fmt.Errorf("演習 %q が見つかりません", prog.InProgress)
			}

			for _, f := range ex.Files {
				data, err := fs.ReadFile(r.exercFS, prog.InProgress+"/solution/"+f)
				if err != nil {
					cmd.PrintErrf("  解答 %s が見つかりません\n", f)
					continue
				}
				displayName := strings.TrimSuffix(f, ".txt")
				cmd.Printf("=== %s ===\n", colorCyan(displayName))
				cmd.Println(string(data))
				cmd.Println()
			}
			return nil
		},
	}
}
