package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

func (r *Runner) newSolutionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "solution",
		Short: "解答例を表示",
		Run: func(cmd *cobra.Command, args []string) {
			prog, _ := progress.Load()
			if prog.InProgress == "" {
				fmt.Fprintln(os.Stderr, "エラー: 進行中の演習がありません")
				os.Exit(1)
			}
			if r.exercFS == nil {
				fmt.Fprintln(os.Stderr, "エラー: 解答データが見つかりません")
				os.Exit(1)
			}
			ex, err := exercise.LoadFromFS(r.exercFS, prog.InProgress)
			if err != nil {
				fmt.Fprintf(os.Stderr, "エラー: 演習 %q が見つかりません\n", prog.InProgress)
				os.Exit(1)
			}

			for _, f := range ex.Files {
				data, err := fs.ReadFile(r.exercFS, prog.InProgress+"/solution/"+f)
				if err != nil {
					fmt.Fprintf(os.Stderr, "  解答 %s が見つかりません\n", f)
					continue
				}
				displayName := strings.TrimSuffix(f, ".txt")
				fmt.Printf("=== %s ===\n", colorCyan(displayName))
				fmt.Println(string(data))
				fmt.Println()
			}
		},
	}
}
