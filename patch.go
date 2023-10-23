package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func printBytesArrInHex(arr []byte) {
	for _, b := range arr {
		fmt.Printf("%X", b)
	}
	fmt.Println()
}

func normalizeHex(arr []string) []string {
	for i, h := range arr {
		arr[i] = strings.TrimLeft(h, "0")
	}

	return arr
}

func backupFile(originalFileName, backupFileName string) error {
	backupData, err := os.ReadFile(originalFileName)
	if err != nil {
		return err
	}
	backupOutFile, err := os.Create(backupFileName)
	if err != nil {
		return err
	}
	backupOutFile.Write(backupData)
	backupOutFile.Close()

	return nil
}

func compareExes(exeAName, exeBName string) bool {
	byteA, err := os.ReadFile(exeAName)
	if err != nil {
		return false
	}
	byteB, err := os.ReadFile(exeBName)
	if err != nil {
		return false
	}

	return reflect.DeepEqual(byteA, byteB)
}

var (
	hexExistsEU4Windows = []string{"48", "8D", "0D", "??", "??", "??", "01", "E8", "??", "??", "??", "01", "85", "C0", "0F", "94", "C3", "E8"}
	hexWantedEU4Windows = []string{"48", "8D", "0D", "??", "??", "??", "01", "E8", "??", "??", "??", "01", "31", "C0", "0F", "94", "C3", "E8"}

	hexExistsHOI4Windows = []string{"48", "??", "??", "??", "??", "??", "??", "E8", "??", "??", "??", "01", "85", "C0", "0F", "94", "C3", "E8"}
	hexWantedHOI4Windows = []string{"48", "??", "??", "??", "??", "??", "??", "E8", "??", "??", "??", "01", "31", "C0", "0F", "94", "C3", "E8"}
)

func applyPatch(originalFileName, OS string) error {
	// SUPPORT DROPPED
	//hexExistsEU4Linux := []string{"E8", "??", "??", "E5", "FF", "89", "C3", "E8", "??", "??", "EC", "FF", "31", "F6", "85", "DB", "40", "0F", "94", "C6", "48", "89", "C7"}
	//hexWantedEU4Linux := []string{"E8", "??", "??", "E5", "FF", "89", "C3", "E8", "??", "??", "EC", "FF", "31", "F6", "31", "DB", "40", "0F", "94", "C6", "48", "89", "C7"}
	//hexExistsEU4Darwin := []string{"E8", "7A", "C5", "76", "01", "89", "C3", "E8", "93", "A6", "EC", "FF", "31", "F6", "85", "DB", "40", "0F", "94", "C6", "48", "89", "C7"}
	//hexWantedEU4Darwin := []string{"E8", "7A", "C5", "76", "01", "89", "C3", "E8", "93", "A6", "EC", "FF", "31", "F6", "31", "DB", "40", "0F", "94", "C6", "48", "89", "C7"}
	//hexExistsHOI4Linux := []string{"E8", "20", "70", "F8", "FE", "41", "89", "C7", "31", "DB", "85", "C0", "0F", "94", "C3", "E8", "51", "87", "B8", "FF"}
	//hexWantedHOI4Linux := []string{"E8", "20", "70", "F8", "FE", "41", "89", "C7", "31", "DB", "31", "C0", "0F", "94", "C3", "E8", "51", "87", "B8", "FF"}

	var hexExists, hexWanted []string
	var fileExtension string

	if strings.Contains(originalFileName, "eu4") {
		hexExists = hexExistsEU4Windows
		hexWanted = hexWantedEU4Windows

	} else if strings.Contains(originalFileName, "hoi4") {

		hexExists = hexExistsHOI4Windows
		hexWanted = hexWantedHOI4Windows

	} else {
		return fmt.Errorf("not supported executable")
	}

	switch OS {
	case "windows":
		fileExtension = ".exe"
	default:
		return fmt.Errorf("this OS (%s) is not supported", OS)

	}

	byteExists := make([]byte, len(hexExists))
	byteWanted := make([]byte, len(hexWanted))
	for i := range hexExists {
		var value int64
		if hexExists[i] == "??" {
			value = 0
		} else {
			value, _ = strconv.ParseInt(hexExists[i], 16, 16)
		}
		byteExists[i] = byte(value)
	}
	for i := range hexWanted {
		var value int64
		if hexWanted[i] == "??" {
			value = 0
		} else {
			value, _ = strconv.ParseInt(hexWanted[i], 16, 16)
		}
		byteWanted[i] = byte(value)
	}

	if _, err := os.Stat(originalFileName + fileExtension); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("cannot locate %s in current folder", originalFileName+fileExtension)
	}

	go func() {
		backupFile(originalFileName+fileExtension, originalFileName+"_backup"+fileExtension)
	}()

	originalByte, err := os.ReadFile(originalFileName + fileExtension)
	if err != nil {
		return err
	}

	finalByte := originalByte

	matchesNeeded := len(byteExists)
	matches := 0
	status := false

	for i := 0; i < len(finalByte); i++ {
		if finalByte[i] == byteExists[0] {
			matches++
			for j := range byteExists {
				if (finalByte[i+j] == byteExists[j]) || (byteExists[j] == 0) {
					matches++
				} else {
					matches = 0
					break
				}
			}
			if matches >= matchesNeeded {
				for k := range byteExists {
					if byteExists[k] != 0 {
						finalByte[i+k] = byteWanted[k]
					}
				}
				status = true
			}
		}
	}

	if !status {
		os.Remove(originalFileName + "_backup" + fileExtension)
		return fmt.Errorf("unsupported version of %s or it's patched already. Patch has not been applied", originalFileName+fileExtension)
	}

	out, err := os.Create(originalFileName + fileExtension)
	if err != nil {
		return err
	}

	_, err = out.Write(finalByte)
	if err != nil {
		return err
	}
	err = out.Close()
	if err != nil {
		return err
	}

	return nil
}
