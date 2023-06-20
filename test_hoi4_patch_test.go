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
	err = applyPatch(true, "./test_files/hoi4/hoi4_test", "windows")

	if err != nil {
		t.Error(err)
	}

	if !compareExes("./test_files/hoi4/hoi4_patched.exe", "./test_files/hoi4/hoi4_test.exe") {
		if err != nil {
			t.Error("Executables doesnt match")
		}
	}
	os.Remove("./test_files/hoi4/hoi4_test.exe")
	os.Remove("./test_files/hoi4/hoi4_test_backup.exe")
	os.Remove("./test_files/hoi4/hoi4.exe")
}

// func TestApplyPatchHOILinux(t *testing.T) {
// 	err := backupFile("./test_files/hoi4/hoi4_original", "./test_files/hoi4/hoi4_test")

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	err = applyPatch(true, "./test_files/hoi4/hoi4_test", "linux")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if !compareExes("./test_files/hoi4/hoi4_patched", "./test_files/hoi4/hoi4_test") {
// 		if err != nil {
// 			t.Error("Executables doesnt match")
// 		}
// 	}
// 	os.Remove("./test_files/hoi4/hoi4_test")
// 	os.Remove("./test_files/hoi4/hoi4_test_backup")
// 	os.Remove("./test_files/hoi4/hoi4")

// }
