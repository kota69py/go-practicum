package progress

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Data struct {
	Completed  []string `json:"completed"`
	InProgress string   `json:"in_progress,omitempty"`
}

func Load() (*Data, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("home dir: %w", err)
	}
	dir := filepath.Join(home, ".go-practicum")
	os.MkdirAll(dir, 0755)
	path := filepath.Join(dir, "progress.json")

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Data{}, nil
		}
		return nil, err
	}

	var p Data
	if err := json.Unmarshal(data, &p); err != nil {
		return &Data{}, nil
	}
	return &p, nil
}

func (d *Data) Save() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dir := filepath.Join(home, ".go-practicum")
	os.MkdirAll(dir, 0755)
	path := filepath.Join(dir, "progress.json")
	data, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func (d *Data) Complete(name string) {
	for _, c := range d.Completed {
		if c == name {
			return
		}
	}
	d.Completed = append(d.Completed, name)
}

func (d *Data) IsCompleted(name string) bool {
	for _, c := range d.Completed {
		if c == name {
			return true
		}
	}
	return false
}
