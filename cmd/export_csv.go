package cmd

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
)

func exportCSV(all []exercise.Exercise, prog *progress.Data, outPath string) {
	var buf strings.Builder
	w := csv.NewWriter(&buf)
	_ = w.Write([]string{"name", "title", "category", "difficulty", "completed"})
	for _, ex := range all {
		done := "false"
		if prog.IsCompleted(ex.Name) {
			done = "true"
		}
		_ = w.Write([]string{ex.Name, ex.Title, ex.Category, fmt.Sprintf("%d", ex.Difficulty), done})
	}
	w.Flush()
	writeOutput(buf.String(), outPath)
}
