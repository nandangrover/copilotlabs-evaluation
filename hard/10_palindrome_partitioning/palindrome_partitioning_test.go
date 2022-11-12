package main

import (
	"testing"
)

func TestNumberOfWaysToTraverseGraph(t *testing.T) {
	testCases := []struct {
		width  int
		height int
		want   int
	}{
		{width: 4, height: 3, want: 10},
	}

	for _, tc := range testCases {
		got := numberOfWaysToTraverseGraph(tc.width, tc.height)
		if got != tc.want {
			t.Errorf("numberOfWaysToTraverseGraph(%v, %v) = %v, want: %v", tc.width, tc.height, got, tc.want)
		}
	}
}
