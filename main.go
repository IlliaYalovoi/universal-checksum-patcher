package main

import (
	"errors"
	"fmt"
	"os"
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

func main() {

	originalFile := "eu4.exe"
	finalFile := "eu4.exe"

	if _, err := os.Stat("./eu4.exe"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("ERROR: Cannot locate eu4.exe in current folder")
		fmt.Println("Press enter to exit")
		fmt.Scanln()
		return
	}

	backupData, err := os.ReadFile("./eu4.exe")
	if err != nil {
		panic(err.Error())
	}
	backupOutFile, err := os.Create("./eu4_backup.exe")
	if err != nil {
		panic(err.Error())
	}

	backupOutFile.Write(backupData)
	backupOutFile.Close()

	originalByte, err := os.ReadFile(originalFile)
	if err != nil {
		panic(err.Error())
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
		fmt.Println("ERROR: Unsupported version of eu4.exe or it's patched already. Patch has not been applied")
		os.Remove("./eu4_backup.exe")
		fmt.Println("Press enter to exit")
		fmt.Scanln()
		return
	}

	for i, h := range finalHex {
		value, err := strconv.ParseInt(h, 16, 16)
		if err != nil {
			panic(err.Error())
		}
		finalByte[i] = byte(value)
	}

	out, err := os.Create(finalFile)
	if err != nil {
		panic(err.Error())
	}

	out.Write(finalByte)
	out.Close()

	fmt.Println("eu4.exe successfully patched")
	fmt.Println("Press enter to exit")
	fmt.Scanln()
}
