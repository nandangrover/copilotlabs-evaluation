package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKnapsackProblem(t *testing.T) {
	input := [][]int{
		{0, 2},
		{1, 4},
		{4, 6},
		{0, 4},
		{7, 8},
		{9, 11},
		{3, 10},
	}
	expected := 3
	actual := laptopRentals(input)
	require.Equal(t, expected, actual)
}
