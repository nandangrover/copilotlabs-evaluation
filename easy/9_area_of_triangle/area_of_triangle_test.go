package main

import "testing"

func TestAreaOfTriangle(t *testing.T) {
	if areaOfTriangle(2, 3) != 3 {
		t.Error("Failed")
	}
	if areaOfTriangle(3, 4) != 6 {
		t.Error("Failed")
	}
	if areaOfTriangle(4, 5) != 10 {
		t.Error("Failed")
	}
	if areaOfTriangle(5, 6) != 15 {
		t.Error("Failed")
	}
}
