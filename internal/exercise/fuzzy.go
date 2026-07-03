package exercise

import (
	"io/fs"
	"sort"
)

func Levenshtein(a, b string) int {
	la, lb := len(a), len(b)
	if la == 0 {
		return lb
	}
	if lb == 0 {
		return la
	}
	prev := make([]int, lb+1)
	cur := make([]int, lb+1)
	for j := range prev {
		prev[j] = j
	}
	for i := 0; i < la; i++ {
		cur[0] = i + 1
		for j := 0; j < lb; j++ {
			cost := 1
			if a[i] == b[j] {
				cost = 0
			}
			cur[j+1] = min(cur[j]+1, min(prev[j+1]+1, prev[j]+cost))
		}
		prev, cur = cur, prev
	}
	return prev[lb]
}

type match struct {
	name string
	dist int
}

func SuggestNames(fsys fs.FS, query string, max int) []string {
	all, err := ListFromFS(fsys)
	if err != nil || len(all) == 0 {
		return nil
	}
	var matches []match
	for _, ex := range all {
		d := Levenshtein(query, ex.Name)
		if d <= 3 || d <= len(query)/2 {
			matches = append(matches, match{ex.Name, d})
		}
	}
	sort.Slice(matches, func(i, j int) bool {
		if matches[i].dist != matches[j].dist {
			return matches[i].dist < matches[j].dist
		}
		return matches[i].name < matches[j].name
	})
	if len(matches) > max {
		matches = matches[:max]
	}
	res := make([]string, len(matches))
	for i, m := range matches {
		res[i] = m.name
	}
	return res
}
