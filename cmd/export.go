package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
	"github.com/spf13/cobra"
)

func (r *Runner) newExportCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "export [format]",
		Short: "学習進捗をエクスポート (json / html)",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			format := "json"
			if len(args) > 0 {
				format = args[0]
			}
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

			switch format {
			case "json":
				exportJSON(all, prog)
			case "html":
				exportHTML(all, prog)
			default:
				fmt.Fprintf(os.Stderr, "エラー: 未対応の形式 %q (json / html)\n", format)
				os.Exit(1)
			}
		},
	}
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

func exportJSON(all []exercise.Exercise, prog *progress.Data) {
	catTotal := map[string]int{}
	catDone := map[string]int{}
	var exs []exportExercise
	for _, ex := range all {
		c := prog.IsCompleted(ex.Name)
		exs = append(exs, exportExercise{
			Name: ex.Name, Title: ex.Title, Category: ex.Category,
			Difficulty: ex.Difficulty, Topics: ex.Topics, Completed: c,
		})
		catTotal[ex.Category]++
		if c {
			catDone[ex.Category]++
		}
	}
	var cats []exportCategory
	for name, total := range catTotal {
		cats = append(cats, exportCategory{Name: name, Total: total, Complete: catDone[name]})
	}
	sort.Slice(cats, func(i, j int) bool { return cats[i].Name < cats[j].Name })

	completed := 0
	for _, ex := range exs {
		if ex.Completed {
			completed++
		}
	}
	pct := 0
	if len(exs) > 0 {
		pct = completed * 100 / len(exs)
	}

	data := exportData{
		ExportedAt: time.Now().Format(time.RFC3339),
		Total:      len(exs),
		Completed:  completed,
		Percent:    pct,
		Categories: cats,
		Exercises:  exs,
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(data)
}

func exportHTML(all []exercise.Exercise, prog *progress.Data) {
	catTotal := map[string]int{}
	catDone := map[string]int{}
	completed := 0
	for _, ex := range all {
		catTotal[ex.Category]++
		if prog.IsCompleted(ex.Name) {
			completed++
			catDone[ex.Category]++
		}
	}
	pct := 0
	if len(all) > 0 {
		pct = completed * 100 / len(all)
	}

	var cats []string
	for c := range catTotal {
		cats = append(cats, c)
	}
	sort.Strings(cats)

	var sb strings.Builder
	sb.WriteString("<!DOCTYPE html><html lang=\"ja\"><head><meta charset=\"UTF-8\">")
	sb.WriteString("<title>go-practicum 学習進捗</title>")
	sb.WriteString("<style>body{font-family:sans-serif;max-width:800px;margin:2rem auto;padding:0 1rem}")
	sb.WriteString("h1{color:#333}.bar{height:20px;background:#eee;border-radius:10px;overflow:hidden;margin:4px 0}")
	sb.WriteString(".fill{height:100%;background:#4caf50;transition:width .3s}")
	sb.WriteString(".cat{margin:1rem 0}.cat-name{font-weight:bold}.stat{color:#666;font-size:.9rem}")
	sb.WriteString("table{width:100%;border-collapse:collapse;margin-top:2rem}")
	sb.WriteString("th,td{text-align:left;padding:6px 8px;border-bottom:1px solid #ddd}")
	sb.WriteString(".done{color:#4caf50}.todo{color:#999}</style></head><body>")
	fmt.Fprintf(&sb, "<h1>go-practicum 学習進捗</h1>")
	fmt.Fprintf(&sb, "<p class=\"stat\">%d / %d (%d%%) 完了</p>", completed, len(all), pct)

	for _, c := range cats {
		total := catTotal[c]
		done := catDone[c]
		cpct := 0
		if total > 0 {
			cpct = done * 100 / total
		}
		sb.WriteString("<div class=\"cat\">")
		fmt.Fprintf(&sb, "<span class=\"cat-name\">%s</span> <span class=\"stat\">%d/%d</span>", c, done, total)
		sb.WriteString("<div class=\"bar\"><div class=\"fill\" style=\"width:")
		fmt.Fprintf(&sb, "%d%%", cpct)
		sb.WriteString("\"></div></div></div>")
	}

	sb.WriteString("<table><tr><th>#</th><th>演習</th><th>カテゴリ</th><th>難易度</th><th>状態</th></tr>")
	for _, ex := range all {
		done := prog.IsCompleted(ex.Name)
		status := "<span class=\"todo\">未</span>"
		if done {
			status = "<span class=\"done\">✓</span>"
		}
		fmt.Fprintf(&sb, "<tr><td>%s</td><td>%s</td><td>%s</td><td>",
			ex.Name, ex.Title, ex.Category)
		for range ex.Difficulty {
			sb.WriteString("★")
		}
		sb.WriteString("</td><td>" + status + "</td></tr>")
	}
	sb.WriteString("</table></body></html>")

	home, _ := os.UserHomeDir()
	outPath := filepath.Join(home, "go-practicum-progress.html")
	if err := os.WriteFile(outPath, []byte(sb.String()), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "エラー: 書き込み失敗: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✅ エクスポート完了: %s\n", colorGreen(outPath))
}
