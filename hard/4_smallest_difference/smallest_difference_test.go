package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSmallestDifference(t *testing.T) {
	expected := []int{28, 26}
	output := SmallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17})
	require.Equal(t, expected, output)

	expected = []int{25, 1005}
	output = SmallestDifference([]int{10, 0, 20, 25}, []int{1005, 1006, 1014, 1032, 1031})
	require.Equal(t, expected, output)
}
