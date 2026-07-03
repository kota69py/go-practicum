package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
)

func exportMarkdown(all []exercise.Exercise, prog *progress.Data, outPath string) {
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
	sb.WriteString("# go-practicum 学習進捗\n\n")
	fmt.Fprintf(&sb, "**%d / %d (%d%%) 完了**\n\n", completed, len(all), pct)

	sb.WriteString("## カテゴリ別進捗\n\n")
	sb.WriteString("| カテゴリ | 進捗 |\n|---------|------|\n")
	for _, c := range cats {
		total := catTotal[c]
		done := catDone[c]
		bar := progressBar(done, total, 10)
		fmt.Fprintf(&sb, "| %s | %s %d/%d |\n", c, bar, done, total)
	}

	sb.WriteString("\n## 演習一覧\n\n")
	sb.WriteString("| 名前 | タイトル | カテゴリ | 難易度 | 状態 |\n")
	sb.WriteString("|------|---------|---------|--------|------|\n")
	for _, ex := range all {
		status := "未"
		if prog.IsCompleted(ex.Name) {
			status = "✓"
		}
		diff := strings.Repeat("★", ex.Difficulty) + strings.Repeat("☆", 5-ex.Difficulty)
		fmt.Fprintf(&sb, "| %s | %s | %s | %s | %s |\n", ex.Name, ex.Title, ex.Category, diff, status)
	}
	writeOutput(sb.String(), outPath)
}
