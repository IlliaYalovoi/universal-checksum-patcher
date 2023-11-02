package main

import (
	"os"
	"testing"
)

func TestApplyPatchHOI4Windows(t *testing.T) {
	err := backupFile("./test_files/hoi4/hoi4_original.exe", "./test_files/hoi4/hoi4_test.exe")

	if err != nil {
		t.Error(err)
	}
	err = applyPatch("./test_files/hoi4/hoi4_test", "windows")

	if err != nil {
		t.Error(err)
	}

	if !compareExes("./test_files/hoi4/hoi4_patched.exe", "./test_files/hoi4/hoi4_test.exe") {
		t.Error("Executables doesnt match")
	}

	err = os.Remove("./test_files/hoi4/hoi4_test.exe")
	if err != nil {
		t.Error("Cannot delete temp files")
	}
	err = os.Remove("./test_files/hoi4/hoi4_test_backup.exe")
	if err != nil {
		t.Error("Cannot delete temp files")
	}
	err = os.Remove("./test_files/hoi4/hoi4.exe")
	if err != nil {
		t.Error("Cannot delete temp files")
	}
}
