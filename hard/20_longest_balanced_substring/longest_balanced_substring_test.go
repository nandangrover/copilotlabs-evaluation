package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKnapsackProblem(t *testing.T) {
	input := "(()))("
	expected := 4
	actual := LongestBalancedSubstring(input)
	require.Equal(t, expected, actual)
}
