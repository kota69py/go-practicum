package exercise

import (
	"fmt"
	"io/fs"
	"path"

	"github.com/kota69py/go-practicum/internal/categories"
)

type ValidationError struct {
	Exercise string
	Message  string
}

func (ve ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", ve.Exercise, ve.Message)
}

func Validate(fsys fs.FS) []ValidationError {
	var errs []ValidationError

	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return []ValidationError{{Message: fmt.Sprintf("FS読み取り失敗: %v", err)}}
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		name := e.Name()

		ex, loadErr := LoadFromFS(fsys, name)
		if loadErr != nil {
			errs = append(errs, ValidationError{
				Exercise: name,
				Message:  fmt.Sprintf("exercise.json 読み込み失敗: %v", loadErr),
			})
			continue
		}

		// name一致
		if ex.Name != name {
			errs = append(errs, ValidationError{
				Exercise: name,
				Message:  fmt.Sprintf("name %q がディレクトリ名と一致しません", ex.Name),
			})
		}

		// difficulty 1-5
		if ex.Difficulty < 1 || ex.Difficulty > 5 {
			errs = append(errs, ValidationError{
				Exercise: name,
				Message:  fmt.Sprintf("難易度 %d が範囲外 (1-5)", ex.Difficulty),
			})
		}

		// category既知
		if ex.Category != "" && !categories.IsKnown(ex.Category) {
			errs = append(errs, ValidationError{
				Exercise: name,
				Message:  fmt.Sprintf("未知のカテゴリ %q", ex.Category),
			})
		}

		// files存在確認
		for _, f := range ex.Files {
			starterPath := path.Join(name, "starter", f)
			if _, err := fs.Stat(fsys, starterPath); err != nil {
				errs = append(errs, ValidationError{
					Exercise: name,
					Message:  fmt.Sprintf("starter に %s がありません", f),
				})
			}
			solutionPath := path.Join(name, "solution", f)
			if _, err := fs.Stat(fsys, solutionPath); err != nil {
				errs = append(errs, ValidationError{
					Exercise: name,
					Message:  fmt.Sprintf("solution に %s がありません", f),
				})
			}
		}
	}

	return errs
}
