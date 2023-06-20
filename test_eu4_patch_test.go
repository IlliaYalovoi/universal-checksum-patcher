package main

import (
	"os"
	"testing"
)

func TestApplyPatchEU4Windows(t *testing.T) {
	err := backupFile("./test_files/eu4/eu4_original.exe", "./test_files/eu4/eu4_test.exe")

	if err != nil {
		t.Error(err)
	}
	err = applyPatch(true, "./test_files/eu4/eu4_test", "windows")

	if err != nil {
		t.Error(err)
	}

	if !compareExes("./test_files/eu4/eu4_patched.exe", "./test_files/eu4/eu4_test.exe") {
		if err != nil {
			t.Error("Executables doesnt match")
		}
	}
	os.Remove("./test_files/eu4/eu4_test.exe")
	os.Remove("./test_files/eu4/eu4_test_backup.exe")
	os.Remove("./test_files/eu4/eu4.exe")
}

func TestApplyPatchEU4Linux(t *testing.T) {
	err := backupFile("./test_files/eu4/eu4_original", "./test_files/eu4/eu4_test")

	if err != nil {
		t.Error(err)
	}
	err = applyPatch(true, "./test_files/eu4/eu4_test", "linux")
	if err != nil {
		t.Error(err)
	}

	if !compareExes("./test_files/eu4/eu4_patched", "./test_files/eu4/eu4_test") {
		if err != nil {
			t.Error("Executables doesnt match")
		}
	}
	os.Remove("./test_files/eu4/eu4_test")
	os.Remove("./test_files/eu4/eu4_test_backup")
	os.Remove("./test_files/eu4/eu4")

}
