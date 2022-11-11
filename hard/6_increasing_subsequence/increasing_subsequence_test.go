package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIncreasingSubsequence(t *testing.T) {
	outputSum, outputSequence := MaxSumIncreasingSubsequence([]int{10, 70, 20, 30, 50, 11, 30})
	require.Equal(t, 110, outputSum)
	require.Equal(t, []int{10, 20, 30, 50}, outputSequence)
}
