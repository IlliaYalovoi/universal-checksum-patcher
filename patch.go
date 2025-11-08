package main

import (
	"errors"
	"fmt"
	"os"
)

func applyPatch(filename string) error {
	_, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return errCantLocate
	}

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	err = modifyBytes(filename, bytes)
	if err != nil {
		return fmt.Errorf("failed to modify bytes: %w", err)
	}

	err = backupFile(filename)
	if err != nil {
		return fmt.Errorf("failed to backup file: %w", err)
	}

	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create/truncate file: %w", err)
	}

	_, err = out.Write(bytes)
	if err != nil {
		return fmt.Errorf("failed to write bytes to file: %w", err)
	}
	err = out.Close()
	if err != nil {
		return fmt.Errorf("failed to close file: %w", err)
	}

	return nil
}

func isStartCandidate(bytes []byte) bool {
	return isSlicesEqual(bytes, start1) ||
		isSlicesEqual(bytes, start2) ||
		isSlicesEqual(bytes, start3) ||
		isSlicesEqual(bytes, start4)
}

func isEndCandidate(bytes []byte) bool {
	return isSlicesEqual(bytes, end) ||
		isSlicesEqual(bytes, endEU5)
}

func modifyBytes(filename string, bytes []byte) error {
	matchesCount := 0
	bytesLength := len(bytes)

	for i := 0; i <= bytesLength-startLength; i++ {
		if isStartCandidate(bytes[i : i+startLength]) {
			for j := i + startLength; j <= i+startLength+limit && j <= bytesLength-endLength; j++ {
				if isEndCandidate(bytes[j : j+endLength]) {
					l.Tracef("found match #%d", matchesCount+1)
					copy(bytes[j:j+replacementLength], replacementMap[filename])
					matchesCount++
					break
				}
			}
		}
	}

	if matchesCount == 0 {
		return errNoMatch
	}

	return nil
}
