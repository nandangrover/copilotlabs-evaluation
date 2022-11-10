package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		x, y, want int
	}{
		{1, 2, 3},
		{1, -2, -1},
		{0, 0, 0},
		{-1, 2, 1},
		{-1, -2, -3},
	}
	for _, tt := range tests {
		got := addNumbers(tt.x, tt.y)
		if got != tt.want {
			t.Errorf("addNumbers(%v, %v) = %v; want %v", tt.x, tt.y, got, tt.want)
		}
	}
}
