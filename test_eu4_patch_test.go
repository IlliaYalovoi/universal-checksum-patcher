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
	err = applyPatch(true, "./test_files/eu4/eu4_windows_test", "windows")

	if err != nil {
		t.Error(err)
	}

	if !compareExes("./test_files/eu4/eu4_windows_patched.exe", "./test_files/eu4/eu4_windows_test.exe") {
		if err != nil {
			t.Error("Executables doesnt match")
		}
	}
	os.Remove("./test_files/eu4/eu4_windows_test.exe")
	os.Remove("./test_files/eu4/eu4_windows_test_backup.exe")
	os.Remove("./test_files/eu4/eu4_windows.exe")
}

func TestApplyPatchEU4Linux(t *testing.T) {
	err := backupFile("./test_files/eu4/eu4_linux_original", "./test_files/eu4/eu4_linux_test")

	if err != nil {
		t.Error(err)
	}
	err = applyPatch(true, "./test_files/eu4/eu4_linux_test", "linux")
	if err != nil {
		t.Error(err)
	}

	if !compareExes("./test_files/eu4/eu4_linux_patched", "./test_files/eu4/eu4_linux_test") {
		if err != nil {
			t.Error("Executables doesnt match")
		}
	}
	os.Remove("./test_files/eu4/eu4_linux_test")
	os.Remove("./test_files/eu4/eu4_linux_test_backup")
	os.Remove("./test_files/eu4/eu4_linux")
}

func TestApplyPatchEU4Darwin(t *testing.T) {
	err := backupFile("./test_files/eu4/eu4_darwin_original", "./test_files/eu4/eu4_darwin_test")

	if err != nil {
		t.Error(err)
	}
	err = applyPatch(true, "./test_files/eu4/eu4_darwin_test", "darwin")
	if err != nil {
		t.Error(err)
	}

	if !compareExes("./test_files/eu4/eu4_darwin_patched", "./test_files/eu4/eu4_darwin_test") {
		if err != nil {
			t.Error("Executables doesnt match")
		}
	}
	os.Remove("./test_files/eu4/eu4_darwin_test")
	os.Remove("./test_files/eu4/eu4_darwin_test_backup")
	os.Remove("./test_files/eu4/eu4_darwin")
}
