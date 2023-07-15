package main

import (
	"fmt"
	"runtime"

	"github.com/manifoldco/promptui"
)

var debugMode = false

func main() {

	OS := runtime.GOOS

	if debugMode {

		promptOSChoice := promptui.Select{
			Label:    fmt.Sprintf("Choose desired OS"),
			Items:    []string{"Windows", "Linux", "Darwin"},
			HideHelp: true,
		}
		_, desiredOS, _ := promptOSChoice.Run()

		switch desiredOS {
		case "windows":
			OS = "windows"
		case "linux":
			OS = "linux"
		case "darwin":
			OS = "darwin"
		}
	}
	promptGame := promptui.Select{
		Label:    "Select game to patch",
		Items:    []string{"Europa Universalis IV", "Hearts of Iron IV"},
		HideHelp: true,
	}

	_, result, err := promptGame.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "Europa Universalis IV":
		err = applyPatch(debugMode, "eu4", OS)
	case "Hearts of Iron IV":
		err = applyPatch(debugMode, "hoi4", OS)
	}

	if err != nil {
		fmt.Println("ERROR:", err)
		fmt.Println("Patch not installed, no file has been changed")
	} else {
		fmt.Println("Patch successfully installed, your original executable has been backuped in [original name]_backup")
	}
	fmt.Println("Press enter to exit")
	fmt.Scanln()

}
