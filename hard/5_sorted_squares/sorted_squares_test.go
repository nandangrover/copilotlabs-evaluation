package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	input := []int{1, 2, 3, 5, 6, 8, 9}
	expected := []int{1, 4, 9, 25, 36, 64, 81}
	actual := SortedSquaredArray(input)
	require.Equal(t, expected, actual)
}
