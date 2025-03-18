package main

import "testing"

func TestHexLength(t *testing.T) {
	if len(start1) != len(start2) ||
		len(start2) != len(start3) {

		t.Errorf("start length not equal")
	}

	if len(end) != len(replacement) {
		t.Errorf("end length not equal")
	}
}
