package main

import (
	"fmt"
)

const (
	eu4  = "eu4.exe"
	eu5  = "eu5.exe"
	hoi4 = "hoi4.exe"
)

var (
	l    *logger
	exes = map[string]bool{
		eu4:  true,
		eu5:  true,
		hoi4: true,
	}
)

func main() {
	l = newLogger()

	func() {
		filesInDir, err := getFilesInCurrentDir()
		if err != nil {
			l.Error(err)
			return
		}

		filesToPatch := make([]string, 0)
		for _, file := range filesInDir {
			if exes[file] {
				l.Infof("found %s in current directory", file)
				filesToPatch = append(filesToPatch, file)
			}
		}

		if len(filesToPatch) == 0 {
			l.Error(errCantLocate)
			return
		}

		for _, file := range filesToPatch {
			l.Infof("patching %s", file)
			err = applyPatch(file)
			if err != nil {
				l.Error(err)
				l.Info("patch wasn't installed, no file have been changed")
				return
			}
			l.Infof("patch successfully installed, original executable has been backed up in %s.backup", file)
		}

	}()

	l.Info("press enter to exit...")
	_, _ = fmt.Scanln()
}
