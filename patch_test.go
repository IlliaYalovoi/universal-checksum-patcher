package main

import (
	"testing"
)

func TestApplyPatch(t *testing.T) {

	err := backupFile("eu4_original.exe", "eu4_test.exe")

	if err != nil {
		t.Error(err)
	}

	err = applyPatch("eu4_test.exe", "eu4_test.exe")

	if err != nil {
		t.Error(err)
	}

	if !compareExes("eu4_patched.exe", "eu4_test.exe") {
		if err != nil {
			t.Error("Exes doesnt match")
		}
	}

}
