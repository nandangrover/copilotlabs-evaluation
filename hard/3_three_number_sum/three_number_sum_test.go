package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestThreeNumberSum(t *testing.T) {
	queries := []int{12, 3, 1, 2, -6, 5, -8, 6}
	target := 0
	expected := [][]int{{-8, 2, 6}, {-8, 3, 5}, {-6, 1, 5}}

	actual := ThreeNumberSum(queries, target)
	require.Equal(t, expected, actual)

	queries = []int{1, 2, 3}
	target = 6
	expected = [][]int{{1, 2, 3}}
	actual = ThreeNumberSum(queries, target)
	require.Equal(t, expected, actual)
}
