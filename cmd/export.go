package cmd

import (
	"fmt"
	"os"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

var exportOutput string

func (r *Runner) newExportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export [format]",
		Short: "学習進捗をエクスポート (json / html / csv / md)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(c *cobra.Command, args []string) error {
			format := "json"
			if len(args) > 0 {
				format = args[0]
			}
			if r.exercFS == nil {
				return fmt.Errorf("演習データが見つかりません")
			}
			all, err := exercise.ListFromFS(r.exercFS)
			if err != nil {
				return fmt.Errorf("%v", err)
			}
			prog, _ := progress.Load()

			switch format {
			case "json":
				exportJSON(all, prog, exportOutput)
			case "html":
				exportHTML(all, prog, exportOutput)
			case "csv":
				exportCSV(all, prog, exportOutput)
			case "md":
				exportMarkdown(all, prog, exportOutput)
			default:
				return fmt.Errorf("未対応の形式 %q (json / html / csv / md)", format)
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&exportOutput, "output", "o", "", "出力先ファイルパス（省略時は標準出力）")
	return cmd
}

type exportExercise struct {
	Name       string   `json:"name"`
	Title      string   `json:"title"`
	Category   string   `json:"category"`
	Difficulty int      `json:"difficulty"`
	Topics     []string `json:"topics"`
	Completed  bool     `json:"completed"`
}

type exportData struct {
	ExportedAt string           `json:"exported_at"`
	Total      int              `json:"total"`
	Completed  int              `json:"completed"`
	Percent    int              `json:"percent"`
	Categories []exportCategory `json:"categories"`
	Exercises  []exportExercise `json:"exercises"`
}

type exportCategory struct {
	Name     string `json:"name"`
	Total    int    `json:"total"`
	Complete int    `json:"complete"`
}

func writeOutput(data string, path string) {
	if path == "" {
		fmt.Println(data)
		return
	}
	if err := os.WriteFile(path, []byte(data), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "エラー: 書き込み失敗: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✅ エクスポート完了: %s\n", colorGreen(path))
}
