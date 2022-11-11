package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTotalWaitingTime(t *testing.T) {
	queries := []int{3, 2, 1, 2, 6}
	expected := 17
	actual := MinimumWaitingTime(queries)

	require.Equal(t, expected, actual)

	queries = []int{1, 4, 5}
	expected = 6
	actual = MinimumWaitingTime(queries)
	require.Equal(t, expected, actual)
}
