package progress

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Data struct {
	Completed       []string `json:"completed"`
	InProgress      string   `json:"in_progress,omitempty"`
	LastCompletedAt string   `json:"last_completed_at,omitempty"`
}

func (d *Data) Complete(name string) {
	for _, c := range d.Completed {
		if c == name {
			return
		}
	}
	d.Completed = append(d.Completed, name)
	d.LastCompletedAt = time.Now().Format(time.RFC3339)
}

func (d *Data) IsCompleted(name string) bool {
	for _, c := range d.Completed {
		if c == name {
			return true
		}
	}
	return false
}

type Store interface {
	Load() (*Data, error)
	Save(*Data) error
}

type FileStore struct{}

func NewFileStore() *FileStore {
	return &FileStore{}
}

func (s *FileStore) Load() (*Data, error) {
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

func (s *FileStore) Save(d *Data) error {
	data, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dir := filepath.Join(home, ".go-practicum")
	os.MkdirAll(dir, 0755)
	path := filepath.Join(dir, "progress.json")
	return os.WriteFile(path, data, 0644)
}

var defaultStore Store = NewFileStore()

func Load() (*Data, error) {
	return defaultStore.Load()
}

func Save(d *Data) error {
	return defaultStore.Save(d)
}
