package exercise

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func LoadFromFS(fsys fs.FS, name string) (*Exercise, error) {
	data, err := fs.ReadFile(fsys, name+"/exercise.json")
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", name, err)
	}
	var ex Exercise
	if err := json.Unmarshal(data, &ex); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", name, err)
	}
	ex.Name = name
	return &ex, nil
}

func ListFromFS(fsys fs.FS) ([]Exercise, error) {
	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return nil, fmt.Errorf("reading exercises: %w", err)
	}
	var result []Exercise
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		ex, err := LoadFromFS(fsys, e.Name())
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				continue
			}
			return nil, err
		}
		result = append(result, *ex)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result, nil
}

func CopyFromFS(fsys fs.FS, name, dst string) error {
	return copyDirFromFS(fsys, name, dst)
}

func copyDirFromFS(fsys fs.FS, src, dst string) error {
	entries, err := fs.ReadDir(fsys, src)
	if err != nil {
		return err
	}
	for _, e := range entries {
		srcPath := src + "/" + e.Name()
		dstName := strings.TrimSuffix(e.Name(), ".txt")
		dstPath := filepath.Join(dst, dstName)
		if e.IsDir() {
			_ = os.MkdirAll(dstPath, 0755)
			if err := copyDirFromFS(fsys, srcPath, dstPath); err != nil {
				return err
			}
		} else {
			data, err := fs.ReadFile(fsys, srcPath)
			if err != nil {
				return err
			}
			if err := os.WriteFile(dstPath, data, 0644); err != nil { //nolint:gosec
				return err
			}
		}
	}
	return nil
}
