package main

import "testing"

func TestSwapNumbers(t *testing.T) {
	test := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, -2}, []int{-2, 1}},
		{[]int{0, 0}, []int{0, 0}},
		{[]int{-1, 2}, []int{2, -1}},
	}

	for _, tt := range test {
		got := swap_numbers(tt.input[0], tt.input[1])
		if got[0] != tt.expected[0] || got[1] != tt.expected[1] {
			t.Errorf("swapNumbers(%v) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}
