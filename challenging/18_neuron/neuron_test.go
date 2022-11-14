package main

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestTournamentWinner(t *testing.T) {
	examples := []Example{
		Example{
			Features: []float64{0.0, 0.0},
			Label:    0.0,
		},
		Example{
			Features: []float64{0.0, 1.0},
			Label:    1.0,
		},
		Example{
			Features: []float64{1.0, 0.0},
			Label:    1.0,
		},
		Example{
			Features: []float64{1.0, 1.0},
			Label:    0.0,
		},
	}
	neuron := Neuron{
		weights:  []float64{0.49671415, 0.64768854, 0.52302986, 0.64589411},
		Examples: examples,
	}
	neuron.Train(0.01, 10, 200)
	actual := math.Round(neuron.Predict([]float64{0.79, 0.89, 0.777}))
	expected := 1.0
	require.Equal(t, expected, actual)
}
