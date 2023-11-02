package main

import (
	"os"
	"testing"
)

func TestApplyPatchEU4Windows(t *testing.T) {
	err := backupFile("./test_files/eu4/eu4_windows_original.exe", "./test_files/eu4/eu4_windows_test.exe")

	if err != nil {
		t.Error(err)
	}
	err = applyPatch("./test_files/eu4/eu4_windows_test", "windows")

	if err != nil {
		t.Error(err)
	}

	if !compareExes("./test_files/eu4/eu4_windows_patched.exe", "./test_files/eu4/eu4_windows_test.exe") {
		t.Error("Executables doesnt match")
	}

	err = os.Remove("./test_files/eu4/eu4_windows_test.exe")
	if err != nil {
		t.Error("Cannot delete temp files")
	}
	err = os.Remove("./test_files/eu4/eu4_windows_test_backup.exe")
	if err != nil {
		t.Error("Cannot delete temp files")
	}
	err = os.Remove("./test_files/eu4/eu4_windows.exe")
	if err != nil {
		t.Error("Cannot delete temp files")
	}
}
