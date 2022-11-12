package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKnapsackProblem(t *testing.T) {
	expected := []interface{}{10, []int{1, 3}}
	output := knapsackProblem([][]int{{1, 2}, {4, 3}, {5, 6}, {6, 7}}, 10)
	require.Equal(t, expected, output)
}
