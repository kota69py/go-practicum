package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
)

func exportHTML(all []exercise.Exercise, prog *progress.Data, outPath string) {
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

	out := outPath
	if out == "" {
		home, _ := os.UserHomeDir()
		out = filepath.Join(home, "go-practicum-progress.html")
	}
	if err := os.WriteFile(out, []byte(sb.String()), 0644); err != nil { //nolint:gosec
		fmt.Fprintf(os.Stderr, "エラー: 書き込み失敗: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✅ エクスポート完了: %s\n", colorGreen(out))
}
