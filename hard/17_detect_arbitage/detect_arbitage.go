package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(DetectArbitrage([][]float64{
		[]float64{1.0, 0.8631, 0.5903},
		[]float64{1.1586, 1.0, 0.6849},
		[]float64{1.6939, 1.46, 1.0},
	})) // True
}

// O(n^3) time | O(n^2) space - where n is the number of currencies
func DetectArbitrage(exchangeRates [][]float64) bool {
	// To use exchange rates as edge weights, we must be able to add them.
	// Since log(a*b) = log(a) + log(b), we can convert all rates to
	// -log10(rate) to use them as edge weights.
	logExchangeRates := ConvertToLogMatrix(exchangeRates)

	// A negative weight cycle indicates an arbitrage.
	return FoundNegativeWeightCycle(logExchangeRates, 0)
}

// Runs the Bellman–Ford Algorithm to detect any negative-weight cycles.
func FoundNegativeWeightCycle(graph [][]float64, start int) bool {
	distancesFromStart := make([]float64, len(graph))
	for i := range distancesFromStart {
		distancesFromStart[i] = math.Inf(1)
	}
	distancesFromStart[start] = 0

	for _ = range graph {
		// If no update occurs, that means there's no negative cycle.
		if !RelaxEdgesAndUpdateDistances(graph, distancesFromStart) {
			return false
		}
	}

	return RelaxEdgesAndUpdateDistances(graph, distancesFromStart)
}

// Returns `True` if any distance was updated
func RelaxEdgesAndUpdateDistances(graph [][]float64, distances []float64) bool {
	updated := false
	for sourceIdx, edges := range graph {
		for destinationIdx, edgeWeight := range edges {
			newDistanceToDestination := distances[sourceIdx] + edgeWeight
			if newDistanceToDestination < distances[destinationIdx] {
				updated = true
				distances[destinationIdx] = newDistanceToDestination
			}
		}
	}

	return updated
}

func ConvertToLogMatrix(matrix [][]float64) [][]float64 {
	newMatrix := make([][]float64, len(matrix))
	for row, rates := range matrix {
		newMatrix[row] = make([]float64, len(rates))
		for rateIdx, rate := range rates {
			newMatrix[row][rateIdx] = -math.Log10(rate)
		}
	}

	return newMatrix
}
