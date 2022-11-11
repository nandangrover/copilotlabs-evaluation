package main

import (
	"fmt"
	"sort"
)

// O(nlogn) time | O(1) space - where n is the number of queries
func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)

	totalWaitingTime := 0
	for idx, duration := range queries {
		queriesLeft := len(queries) - (idx + 1)
		totalWaitingTime += duration * queriesLeft
	}

	return totalWaitingTime
}

func main() {
	queries := []int{3, 2, 1, 2, 6}
	fmt.Println(MinimumWaitingTime(queries))
}
