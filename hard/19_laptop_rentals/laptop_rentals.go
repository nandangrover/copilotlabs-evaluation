package main

import (
	"fmt"
	"sort"
)

func laptopRentals(times [][]int) int {
	if len(times) == 0 {
		return 0
	}

	usedLaptops := 0
	startTimes := []int{}
	endTimes := []int{}

	for _, interval := range times {
		startTimes = append(startTimes, interval[0])
		endTimes = append(endTimes, interval[1])
	}

	sort.Ints(startTimes)
	sort.Ints(endTimes)

	startIterator := 0
	endIterator := 0

	for startIterator < len(times) {
		if startTimes[startIterator] >= endTimes[endIterator] {
			usedLaptops -= 1
			endIterator += 1
		}

		usedLaptops += 1
		startIterator += 1
	}

	return usedLaptops
}

func main() {
	fmt.Println(laptopRentals([][]int{{0, 2}, {1, 4}, {4, 6}, {0, 4}, {7, 8}, {9, 11}, {3, 10}}))
}
