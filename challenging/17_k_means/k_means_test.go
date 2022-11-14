package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTournamentWinner(t *testing.T) {
	user_feature_map := map[string][]float64{
		"user1":  {1, 1, 1},
		"user2":  {2, 2, 2},
		"user3":  {3, 3, 3},
		"user4":  {4, 4, 4},
		"user5":  {5, 5, 5},
		"user6":  {6, 6, 6},
		"user7":  {7, 7, 7},
		"user8":  {8, 8, 8},
		"user9":  {9, 9, 9},
		"user10": {10, 10, 10},
	}
	num_features_per_user := 3
	k := 2
	expected := [][]float64{{8.0, 8.0, 8.0}, {3.0, 3.0, 3.0}}
	actual := get_k_means(user_feature_map, num_features_per_user, k)
	require.Equal(t, expected, actual)
}
