package main

import (
	"flag"
	"fmt"
	"runtime"
	"github.com/manifoldco/promptui"
)

const (
	eu4exe  = "eu4.exe"
	eu5exe  = "eu5.exe"
	hoi4exe = "hoi4.exe"
	eu4bin  = "eu4"
	eu5bin  = "eu5"
	hoi4bin = "hoi4"
)

var (
	l    *logger
	exes = map[string]bool{
		eu4exe:  true,
		eu5exe:  true,
		hoi4exe: true,
		eu4bin:  true,
		eu5bin:  true,
		hoi4bin: true,
	}
)

func main() {
	l = newLogger()
	searchDir := flag.String("dir", "", "directory to search for game executable(s)")
	flag.Parse()

	if *searchDir != "" {
		if err := runNonInteractive(*searchDir); err != nil {
			l.Error(err)
		}
	} else {
		if err := runInteractive(); err != nil {
			l.Error(err)
		}
	}

	if runtime.GOOS != "windows" {
		return
	}

	l.Info("press enter to exit...")
	_, _ = fmt.Scanln()
}

func runNonInteractive(searchDir string) error {
	filesToPatch, err := findFilesToPatch(searchDir)
	if err != nil {
		return err
	}
	if len(filesToPatch) == 0 {
		return errCantLocate
	}

	for _, file := range filesToPatch {
		l.Infof("patching %s", file)
		err = applyPatch(file)
		if err != nil {
			l.Info("patch wasn't installed, no file have been changed")
			return err
		}
		l.Infof("patch successfully installed, original executable has been backed up in %s.backup", file)
	}

	return nil
}

func runInteractive() error {
	filesToPatch, err := findFilesToPatch("")
	if err != nil {
		return err
	}
	if len(filesToPatch) == 0 {
		return errCantLocate
	}

	selector := promptui.Select{
		Label: "Select executable to patch",
		Items: filesToPatch,
		Size:  10,
	}

	index, _, err := selector.Run()
	if err != nil {
		return fmt.Errorf("interactive selection cancelled: %w", err)
	}

	selected := filesToPatch[index]
	confirm := promptui.Prompt{
		Label:     fmt.Sprintf("Patch %s", selected),
		IsConfirm: true,
		Default:   "n",
	}

	if _, err := confirm.Run(); err != nil {
		return fmt.Errorf("patch cancelled")
	}

	l.Infof("patching %s", selected)
	if err := applyPatch(selected); err != nil {
		l.Info("patch wasn't installed, no file have been changed")
		return err
	}
	l.Infof("patch successfully installed, original executable has been backed up in %s.backup", selected)

	return nil
}
