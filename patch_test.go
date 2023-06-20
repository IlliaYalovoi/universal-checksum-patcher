package main

import (
	"os"
	"testing"
)

func TestApplyPatchWindows(t *testing.T) {
	t.Parallel()
	err := backupFile("./test_files/eu4_original.exe", "./test_files/eu4_test.exe")

	if err != nil {
		t.Error(err)
	}
	err = applyPatch(true, "./test_files/eu4_test", "./test_files/eu4_test", "windows")

	if err != nil {
		t.Error(err)
	}

	if !compareExes("./test_files/eu4_patched.exe", "./test_files/eu4_test.exe") {
		if err != nil {
			t.Error("Executables doesnt match")
		}
	}
	os.Remove("./test_files/eu4_test.exe")
}

func TestApplyPatchLinux(t *testing.T) {
	t.Parallel()
	err := backupFile("./test_files/eu4_original", "./test_files/eu4_test")

	if err != nil {
		t.Error(err)
	}
	err = applyPatch(true, "./test_files/eu4_test", "./test_files/eu4_test", "linux")
	if err != nil {
		t.Error(err)
	}

	if !compareExes("./test_files/eu4_patched", "./test_files/eu4_test") {
		if err != nil {
			t.Error("Executables doesnt match")
		}
	}
	os.Remove("./test_files/eu4_test")

}
