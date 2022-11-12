package main

import (
	"fmt"
	"math"
)

func knapsackProblem(items [][]int, capacity int) []int {
	knapsackValues := make([][]int, len(items)+1)
	for i := range knapsackValues {
		knapsackValues[i] = make([]int, capacity+1)
	}
	for i := 1; i < len(items)+1; i++ {
		currentWeight := items[i-1][1]
		currentValue := items[i-1][0]
		for c := 0; c < capacity+1; c++ {
			if currentWeight > c {
				knapsackValues[i][c] = knapsackValues[i-1][c]
			} else {
				knapsackValues[i][c] = int(math.Max(
					float64(knapsackValues[i-1][c]),
					float64(knapsackValues[i-1][c-currentWeight]+currentValue),
				))
			}
		}
	}
	return []int{knapsackValues[len(knapsackValues)-1][len(knapsackValues[0])-1], getKnapsackItems(knapsackValues, items)}
}

func getKnapsackItems(knapsackValues [][]int, items [][]int) int {
	sequence := 0
	i := len(knapsackValues) - 1
	c := len(knapsackValues[0]) - 1
	for i > 0 {
		if knapsackValues[i][c] == knapsackValues[i-1][c] {
			i -= 1
		} else {
			sequence += i - 1
			c -= items[i-1][1]
			i -= 1
		}
		if c == 0 {
			break
		}
	}
	return sequence
}

func main() {
	fmt.Println(knapsackProblem([][]int{{1, 2}, {4, 3}, {5, 6}, {6, 7}}, 10))
}
