package cmd

import (
	"encoding/json"
	"sort"
	"strings"
	"time"

	"github.com/kota69py/go-practicum/internal/exercise"
	"github.com/kota69py/go-practicum/internal/progress"
)

func exportJSON(all []exercise.Exercise, prog *progress.Data, outPath string) {
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
	var buf strings.Builder
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	enc.Encode(data)
	writeOutput(buf.String(), outPath)
}
