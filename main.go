package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

const (
	eu4  = "eu4.exe"
	hoi4 = "hoi4.exe"
)

func main() {
	prompt := promptui.Select{
		Label: "Select game to patch",
		Items: []string{
			"Europa Universalis IV",
			"Hearts of Iron IV",
		},
		HideHelp: true,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "Europa Universalis IV":
		err = applyPatch(eu4)
	case "Hearts of Iron IV":
		err = applyPatch(hoi4)
	}

	if err != nil {
		fmt.Println("ERROR:", err)
		fmt.Println("Wasn't not installed, file hasn't changed")
	} else {
		fmt.Println("Patch successfully installed, your original executable has been backuped in [original name].backup")
	}

	fmt.Println("Press enter to exit")
	fmt.Scanln()
}
