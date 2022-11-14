package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrint(t *testing.T) {
	input := [][]int{
		{1, 2, 5},
		{0, 2},
		{0, 1, 3},
		{2, 4, 5},
		{3, 5},
		{0, 3, 4},
	}
	expected := true
	actual := twoEdgeConnectedGraph(input)
	require.Equal(t, expected, actual)
}
