package main

import (
	"fmt"
	"os"
)

func backupFile(original string) error {
	data, err := os.ReadFile(original)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	backup, err := os.Create(fmt.Sprintf("%s.backup", original))
	if err != nil {
		return fmt.Errorf("failed to create backup file: %w", err)
	}
	_, err = backup.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write to backup file: %w", err)
	}
	err = backup.Close()
	if err != nil {
		return fmt.Errorf("failed to close backup file: %w", err)
	}

	return nil
}

func isSlicesEqual[T comparable](first, second []T) bool {
	if len(first) != len(second) {
		return false
	}

	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}

func getFilesInCurrentDir() ([]string, error) {
	files, err := os.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var result []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		result = append(result, file.Name())
	}

	return result, nil
}
