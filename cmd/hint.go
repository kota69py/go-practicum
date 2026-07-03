package cmd

import (
	"fmt"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

func (r *Runner) newHintCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "hint",
		Short: "ヒントを表示",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			prog, _ := progress.Load()
			if prog.InProgress == "" {
				return fmt.Errorf("進行中の演習がありません")
			}
			if r.exercFS == nil {
				return fmt.Errorf("演習データが見つかりません")
			}
			ex, err := exercise.LoadFromFS(r.exercFS, prog.InProgress)
			if err != nil {
				return fmt.Errorf("演習 %q が見つかりません", prog.InProgress)
			}

			if len(ex.Hints) == 0 {
				cmd.Println("ヒントは用意されていません。")
				return nil
			}

			cmd.Println("💡 " + colorYellow("ヒント:"))
			cmd.Println()
			for i, h := range ex.Hints {
				cmd.Printf("  %d. %s\n", i+1, h)
			}
			return nil
		},
	}
}
