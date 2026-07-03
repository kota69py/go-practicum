package cmd

import (
	"fmt"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/spf13/cobra"
)

func (r *Runner) newDoctorCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "doctor",
		Short: "演習データを検証",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if r.exercFS == nil {
				return fmt.Errorf("演習データが見つかりません")
			}

			cmd.Println("🔍 " + colorCyan("演習データを検証中..."))
			cmd.Println()

			errs := exercise.Validate(r.exercFS)
			if len(errs) == 0 {
				cmd.Println("✅ " + colorGreen("すべての演習データは正常です"))
				return nil
			}

			cmd.Printf("❌ %d 件の問題が見つかりました:\n\n", len(errs))
			for _, ve := range errs {
				cmd.Printf("  %s\n", colorRed(ve.Error()))
			}
			return fmt.Errorf("検証エラー")
		},
	}
}
