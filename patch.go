package main

import (
	"errors"
	"os"
)

func applyPatch(filename string) error {
	_, err := os.Stat(filename)
	if errors.Is(err, os.ErrNotExist) {
		return ErrCantLocate
	}

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = modifyBytes(bytes)
	if err != nil {
		return err
	}

	err = backupFile(filename)
	if err != nil {
		return err
	}

	out, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = out.Write(bytes)
	if err != nil {
		return err
	}
	err = out.Close()
	if err != nil {
		return err
	}

	return nil
}

func isStartCandidate(bytes []byte) bool {
	return isByteSlicesEqual(bytes, start1) || isByteSlicesEqual(bytes, start2) || isByteSlicesEqual(bytes, start3)
}

func isEndCandidate(bytes []byte) bool {
	return isByteSlicesEqual(bytes, end)
}

func modifyBytes(bytes []byte) error {
	atLeastOnePatched := false

	for i := 0; i < len(bytes); i++ {
		if i > len(bytes)-limit {
			break
		}

		if !isStartCandidate(bytes[i : i+startLength]) {
			continue
		}

		for j := i + startLength; j < i+startLength+limit && j < len(bytes)-endLength; j++ {
			if !isEndCandidate(bytes[j : j+endLength]) {
				continue
			}

			for k := 0; k < len(replacement); k++ {
				bytes[j+k] = replacement[k]
			}

			atLeastOnePatched = true
		}
	}

	if atLeastOnePatched {
		return nil
	}
	return ErrNoMatch
}
