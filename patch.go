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

	// hexA := make([]string, len(byteA))
	// for i := range byteA {
	// 	hexA[i] = fmt.Sprintf("%X", byteA[i])
	// }

	// hexB := make([]string, len(byteB))
	// for i := range byteB {
	// 	hexB[i] = fmt.Sprintf("%X", byteB[i])
	// }

	return reflect.DeepEqual(byteA, byteB)
}

func applyPatch(originalFileName, finalFileName string) error {

	if _, err := os.Stat("eu4.exe"); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("ERROR: Cannot locate %s in current folder\n", originalFileName)
		fmt.Println("Press enter to exit")
		fmt.Scanln()
		return errors.New("Cant find exe")
	}

	err := backupFile("eu4.exe", "eu4_backup.exe")
	if err != nil {
		return err
	}

	originalByte, err := os.ReadFile(originalFileName)
	if err != nil {
		return err
	}
	originalHex := make([]string, len(originalByte))
	for i := range originalByte {
		originalHex[i] = fmt.Sprintf("%X", originalByte[i])
	}

	hexExists := []string{"48", "8D", "0D", "??", "??", "??", "01", "E8", "??", "??", "??", "01", "85", "C0", "0F", "94", "C3", "E8", "97", "79", "E9"}
	hexWanted := []string{"48", "8D", "0D", "??", "??", "??", "01", "E8", "??", "??", "??", "01", "31", "C0", "0F", "94", "C3", "E8", "97", "79", "E9"}

	hexExists = normalizeHex(hexExists)
	hexWanted = normalizeHex(hexWanted)

	finalHex := originalHex
	finalByte := make([]byte, len(originalByte))

	matchesNeeded := len(hexExists)
	matches := 0
	status := false
	for i := range finalHex {

		if finalHex[i] == hexExists[0] {
			matches++
			// fmt.Println(matches, finalHex[i])
			for j := range hexExists {
				if (finalHex[i+j] == hexExists[j]) || (hexExists[j] == "??") {
					matches++
				} else {
					matches = 0
					break
				}
			}
			if matches >= matchesNeeded {
				for k := range hexExists {
					if hexExists[k] != "??" {
						finalHex[i+k] = hexWanted[k]
					}
				}
				status = true
				break
			}
		}
	}

	if !status {
		fmt.Printf("ERROR: Unsupported version of %s or it's patched already. Patch has not been applied\n", originalFileName)
		os.Remove("./eu4_backup.exe")
		fmt.Println("Press enter to exit")
		fmt.Scanln()

		return errors.New("Cant apply patch")
	}

	for i, h := range finalHex {
		value, err := strconv.ParseInt(h, 16, 16)
		if err != nil {
			return err
		}
		finalByte[i] = byte(value)
	}

	out, err := os.Create(finalFileName)
	if err != nil {
		return err
	}

	out.Write(finalByte)
	out.Close()

	fmt.Printf("%s successfully patched\n", originalFileName)
	fmt.Println("Press enter to exit")
	fmt.Scanln()

	return nil
}
