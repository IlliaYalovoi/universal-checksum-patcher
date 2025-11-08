package main

import "testing"

func TestHexLength(t *testing.T) {
	if len(start1) != len(start2) ||
		len(start2) != len(start3) ||
		len(start3) != len(start4) {

		t.Errorf("start length not equal")
	}

	if len(end) != len(endEU5) || len(replacement) != len(replacementEU5) {
		t.Errorf("end length not equal")
	}
}
