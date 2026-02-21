package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type binaryKind int

const (
	binaryKindUnknown binaryKind = iota
	binaryKindPE
	binaryKindELF
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

	kind := detectBinaryKind(bytes)
	switch kind {
	case binaryKindPE:
		err = patchPE(filename, bytes)
	case binaryKindELF:
		err = patchELF(filename, bytes)
	default:
		err = errUnsupportedBinaryFormat
	}
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

func detectBinaryKind(bytes []byte) binaryKind {
	if len(bytes) >= 4 &&
		bytes[0] == 0x7f &&
		bytes[1] == 'E' &&
		bytes[2] == 'L' &&
		bytes[3] == 'F' {
		return binaryKindELF
	}

	if len(bytes) >= 2 &&
		bytes[0] == 'M' &&
		bytes[1] == 'Z' {
		return binaryKindPE
	}

	return binaryKindUnknown
}

func patchPE(filename string, bytes []byte) error {
	return modifyBytes(filename, bytes, replacementMapPE)
}

func patchELF(filename string, bytes []byte) error {
	exeName := filepath.Base(filename)
	if exeName == hoi4bin {
		return patchELFHoi4(bytes)
	}

	return modifyBytes(filename, bytes, replacementMapELF)
}

func patchELFHoi4(data []byte) error {
	matchesCount := 0
	searchFrom := 0

	for {
		idx := bytes.Index(data[searchFrom:], elfHOI4Needle)
		if idx < 0 {
			break
		}
		idx += searchFrom
		patchOffset := idx + 2
		copy(data[patchOffset:patchOffset+replacementLength], replacement)
		matchesCount++
		searchFrom = idx + len(elfHOI4Needle)
	}

	if matchesCount == 0 {
		return errNoMatch
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

func modifyBytes(filename string, bytes []byte, replacementMap map[string][]byte) error {
	exeName := filepath.Base(filename)
	replacementBytes, ok := replacementMap[exeName]
	if !ok {
		return errUnsupportedExecutable
	}

	matchesCount := 0
	bytesLength := len(bytes)

	for i := 0; i <= bytesLength-startLength; i++ {
		if isStartCandidate(bytes[i : i+startLength]) {
			for j := i + startLength; j <= i+startLength+limit && j <= bytesLength-endLength; j++ {
				if isEndCandidate(bytes[j : j+endLength]) {
					l.Tracef("found match #%d", matchesCount+1)
					copy(bytes[j:j+replacementLength], replacementBytes)
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
