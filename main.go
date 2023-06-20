package main

import "runtime"

func main() {

	originalFileName := "eu4"
	finalFileName := "eu4"

	applyPatch(false, originalFileName, finalFileName, runtime.GOOS)

}
