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

func applyPatch(test bool, originalFileName, OS string) error {

	var fileExtension string
	var hexExists []string
	var hexWanted []string
	var hexExistsWindows []string
	var hexWantedWindows []string
	var hexExistsLinux []string
	var hexWantedLinux []string
	// EU4
	hexExistsEU4Windows := []string{"48", "8D", "0D", "??", "??", "??", "01", "E8", "??", "??", "??", "01", "85", "C0", "0F", "94", "C3", "E8"}
	hexWantedEU4Windows := []string{"48", "8D", "0D", "??", "??", "??", "01", "E8", "??", "??", "??", "01", "31", "C0", "0F", "94", "C3", "E8"}
	hexExistsEU4Linux := []string{"E8", "65", "95", "E5", "FF", "89", "C3", "E8", "38", "08", "EC", "FF", "31", "F6", "85", "DB", "40", "0F", "94", "C6", "48", "89", "C7"}
	hexWantedEU4Linux := []string{"E8", "65", "95", "E5", "FF", "89", "C3", "E8", "38", "08", "EC", "FF", "31", "F6", "31", "DB", "40", "0F", "94", "C6", "48", "89", "C7"}
	// HOI4
	hexExistsHOI4Windows := []string{"48", "8D", "0D", "77", "B6", "C9", "01", "E8", "CA", "86", "B3", "01", "85", "C0", "0F", "94", "C3", "E8", "90"}
	hexWantedHOI4Windows := []string{"48", "8D", "0D", "77", "B6", "C9", "01", "E8", "CA", "86", "B3", "01", "31", "C0", "0F", "94", "C3", "E8", "90"}
	hexExistsHOI4Linux := []string{}
	hexWantedHOI4Linux := []string{}

	if strings.Contains(originalFileName, "eu4") {
		hexExistsWindows = hexExistsEU4Windows
		hexWantedWindows = hexWantedEU4Windows

		hexExistsLinux = hexExistsEU4Linux
		hexWantedLinux = hexWantedEU4Linux
	} else if strings.Contains(originalFileName, "hoi4") {
		if OS == "linux" {

			return errors.New("patching linux version of Hearts of Iron IV is not currently supported")
		}

		hexExistsWindows = hexExistsHOI4Windows
		hexWantedWindows = hexWantedHOI4Windows

		hexExistsLinux = hexExistsHOI4Linux
		hexWantedLinux = hexWantedHOI4Linux

	} else {
		return errors.New("not supported executable")
	}

	switch OS {
	case "windows":
		fileExtension = ".exe"
		hexExists = hexExistsWindows
		hexWanted = hexWantedWindows
	case "linux":
		fileExtension = ""
		hexExists = hexExistsLinux
		hexWanted = hexWantedLinux
	default:
		fileExtension = ""
		return fmt.Errorf("this OS (%s) is not supported", OS)

	}

	byteExists := make([]byte, len(hexExists))
	byteWanted := make([]byte, len(hexWanted))
	for i, h := range hexExists {
		var value int64
		if h == "??" {
			value = 0
		} else {
			value, _ = strconv.ParseInt(h, 16, 16)
		}
		byteExists[i] = byte(value)
	}
	for i, h := range hexWanted {
		var value int64
		if h == "??" {
			value = 0
		} else {
			value, _ = strconv.ParseInt(h, 16, 16)
		}
		byteWanted[i] = byte(value)
	}

	go func() {
		backupFile(originalFileName+fileExtension, originalFileName+"_backup"+fileExtension)
	}()

	if _, err := os.Stat(originalFileName + fileExtension); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("cannot locate %s in current folder", originalFileName+fileExtension)
	}

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

				break
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

	out.Write(finalByte)
	out.Close()

	return nil
}
