package main

import "testing"

func TestProgram(t *testing.T) {
	input := 4
	expected := 2
	actual := nonAttackingQueens(input)
	if actual != expected {
		t.Errorf("nonAttackingQueens(%d) = %d; expected %d", input, actual, expected)
	}
}
