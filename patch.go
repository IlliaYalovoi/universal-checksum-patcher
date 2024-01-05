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

func modifyBytes(bytes []byte) error {
	for i := 0; i < len(bytes); i++ {
		if !isByteSlicesEqual(bytes[i:i+len(start1)], start1) &&
			!isByteSlicesEqual(bytes[i:i+len(start2)], start2) &&
			!isByteSlicesEqual(bytes[i:i+len(start3)], start3) {
			continue
		}

		for j := i + len(start1); j < i+len(start1)+limit && j < len(bytes); j++ {
			if !isByteSlicesEqual(bytes[j:j+len(end)], end) {
				continue
			}

			copy(bytes[j:j+len(end)], replacement)
			return nil
		}
	}

	return ErrNoMatch
}
