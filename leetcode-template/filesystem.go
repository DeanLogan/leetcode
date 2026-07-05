package main

import (
	"errors"
	"os"
	"path/filepath"
)

func findRepoRoot() (string, error) {
	current, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(current, "go.mod")); err == nil {
			return current, nil
		}

		parent := filepath.Dir(current)
		if parent == current {
			return "", errors.New("go.mod not found")
		}
		current = parent
	}
}

func writeFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0o644)
}
