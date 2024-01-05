package main

import (
	"fmt"
	"os"
)

func backupFile(original string) error {
	data, err := os.ReadFile(original)
	if err != nil {
		return err
	}
	backup, err := os.Create(fmt.Sprintf("%s.backup", original))
	if err != nil {
		return err
	}
	_, err = backup.Write(data)
	if err != nil {
		return err
	}
	err = backup.Close()
	if err != nil {
		return err
	}

	return nil
}

func isByteSlicesEqual(first, second []byte) bool {
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
