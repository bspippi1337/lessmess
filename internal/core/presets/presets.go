package presets

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Preset struct {
	Name    string            `json:"name"`
	BaseURL string            `json:"baseUrl"`
	Headers map[string]string `json:"headers,omitempty"`
	Auth    map[string]string `json:"auth,omitempty"`
}

type Store struct {
	Path string
}

func DefaultStore(appName string) (*Store, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	p := filepath.Join(dir, appName, "presets.json")
	return &Store{Path: p}, nil
}

func (s *Store) Load() ([]Preset, error) {
	b, err := os.ReadFile(s.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Preset{}, nil
		}
		return nil, err
	}
	var ps []Preset
	if err := json.Unmarshal(b, &ps); err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *Store) Save(ps []Preset) error {
	if err := os.MkdirAll(filepath.Dir(s.Path), 0o755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(ps, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.Path, b, 0o600)
}
