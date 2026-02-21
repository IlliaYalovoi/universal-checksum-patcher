package main

import (
	"fmt"
	"os"
	"path/filepath"
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

func findFilesToPatch(overrideDir string) ([]string, error) {
	roots := []string{"."}
	if overrideDir != "" {
		roots = []string{overrideDir}
	} else {
		home, err := os.UserHomeDir()
		if err == nil {
			roots = append(
				roots,
				filepath.Join(home, ".local", "share", "Steam", "steamapps", "common"),
				filepath.Join(home, ".steam", "steam", "steamapps", "common"),
				filepath.Join(home, ".var", "app", "com.valvesoftware.Steam", ".local", "share", "Steam", "steamapps", "common"),
			)
		}

		roots = append(roots, findMountedSteamCommonDirs()...)
	}

	found := make([]string, 0)
	seen := make(map[string]bool)

	for _, root := range roots {
		_, err := os.Stat(root)
		if err != nil {
			continue
		}

		err = filepath.WalkDir(root, func(path string, d os.DirEntry, walkErr error) error {
			if walkErr != nil {
				return nil
			}
			if d.IsDir() {
				return nil
			}
			if !exes[d.Name()] {
				return nil
			}

			absPath, err := filepath.Abs(path)
			if err != nil {
				absPath = path
			}
			if seen[absPath] {
				return nil
			}

			seen[absPath] = true
			found = append(found, absPath)
			l.Infof("found %s", absPath)
			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("failed to walk %s: %w", root, err)
		}
	}

	return found, nil
}

func findMountedSteamCommonDirs() []string {
	patterns := []string{
		"/run/media/*/*/SteamLibrary/steamapps/common",
		"/media/*/*/SteamLibrary/steamapps/common",
		"/mnt/*/SteamLibrary/steamapps/common",
	}

	roots := make([]string, 0)
	seen := make(map[string]bool)
	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			continue
		}

		for _, path := range matches {
			if seen[path] {
				continue
			}
			if _, err := os.Stat(path); err != nil {
				continue
			}
			seen[path] = true
			roots = append(roots, path)
		}
	}

	return roots
}
