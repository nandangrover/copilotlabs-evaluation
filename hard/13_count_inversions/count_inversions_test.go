package main

import "testing"

func TestProgram(test *testing.T) {
	input := []int{2, 3, 3, 1, 9, 5, 6}
	expected := 5
	actual := countInversions(input)
	if actual != expected {
		test.Errorf("Test Failed: input=%v, expected=%v, actual=%v", input, expected, actual)
	}
}
