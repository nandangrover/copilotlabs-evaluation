package main

import "fmt"

// O(n) time | O(n) space - where n is the length of the input array
func SortedSquaredArray(array []int) []int {
	sortedSquares := make([]int, len(array))

	smallerValueIdx := 0
	largerValueIdx := len(array) - 1

	for idx := len(array) - 1; idx >= 0; idx-- {
		smallerValue := array[smallerValueIdx]
		largerValue := array[largerValueIdx]

		if abs(smallerValue) > abs(largerValue) {
			sortedSquares[idx] = smallerValue * smallerValue
			smallerValueIdx += 1
		} else {
			sortedSquares[idx] = largerValue * largerValue
			largerValueIdx -= 1
		}
	}

	return sortedSquares
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	array := []int{-5, -4, 1, 2, 3, 5}
	fmt.Println(SortedSquaredArray(array))
}
