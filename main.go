package main

import (
	"fmt"
	"runtime"

	"github.com/manifoldco/promptui"
)

func main() {

	prompt := promptui.Select{
		Label: "Select game to patch",
		Items: []string{"Europa Universalis IV", "Hearts of Iron IV"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "Europa Universalis IV":
		err = applyPatch(false, "eu4", runtime.GOOS)
	case "Hearts of Iron IV":
		err = applyPatch(false, "hoi4", runtime.GOOS)
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
