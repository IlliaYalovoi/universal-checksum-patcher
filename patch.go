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

func applyPatch(test bool, originalFileName, finalFileName, OS string) error {

	var fileExtension string
	var hexExists []string
	var hexWanted []string
	hexExistsWindows := []string{"48", "8D", "0D", "??", "??", "??", "01", "E8", "??", "??", "??", "01", "85", "C0", "0F", "94", "C3", "E8", "97", "79", "E9"}
	hexWantedWindows := []string{"48", "8D", "0D", "??", "??", "??", "01", "E8", "??", "??", "??", "01", "31", "C0", "0F", "94", "C3", "E8", "97", "79", "E9"}

	hexExistsLinux := []string{"E8", "65", "95", "E5", "FF", "89", "C3", "E8", "38", "08", "EC", "FF", "31", "F6", "85", "DB", "40", "0F", "94", "C6", "48", "89", "C7"}
	hexWantedLinux := []string{"E8", "65", "95", "E5", "FF", "89", "C3", "E8", "38", "08", "EC", "FF", "31", "F6", "31", "DB", "40", "0F", "94", "C6", "48", "89", "C7"}
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
		fmt.Printf("ERROR: This OS is not supported")
		fmt.Println("Press enter to exit")
		fmt.Scanln()
		return errors.New("This OS is not supported")
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
		fmt.Printf("ERROR: Cannot locate %s in current folder\n", originalFileName+fileExtension)
		if !test {
			fmt.Println("Press enter to exit")
			fmt.Scanln()
		}
		return errors.New("Cant find executable file")
	}
	originalByte, err := os.ReadFile(originalFileName + fileExtension)
	if err != nil {
		return err
	}

	finalByte := originalByte

	matchesNeeded := len(byteExists)
	matches := 0
	status := false
	if !test {
		fmt.Println("Patching process started")
	}

	for i := 0; i < len(finalByte); i++ {

		if finalByte[i] == byteExists[0] && finalByte[i+len(byteExists)-1] == byteExists[len(byteExists)-1] {
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
				if !test {
					fmt.Println("Needed byte-combination finded")
				}
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
		fmt.Printf("ERROR: Unsupported version of %s or it's patched already. Patch has not been applied\n", originalFileName+fileExtension)
		os.Remove(originalFileName + "_backup" + fileExtension)
		if !test {
			fmt.Println("Press enter to exit")
			fmt.Scanln()
		}
		return errors.New("Cant apply patch")
	}

	out, err := os.Create(finalFileName + fileExtension)
	if err != nil {
		return err
	}

	out.Write(finalByte)
	out.Close()

	fmt.Printf("%s successfully patched\n", originalFileName+fileExtension)
	if !test {
		fmt.Println("Press enter to exit")
		fmt.Scanln()
	}
	return nil
}
