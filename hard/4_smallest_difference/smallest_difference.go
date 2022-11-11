package main

import (
	"fmt"
	"math"
	"sort"
)

// O(nlog(n) + mlog(m)) time | O(1) space
func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	idxOne, idxTwo := 0, 0
	smallest, current := math.MaxInt32, math.MaxInt32
	smallestPair := []int{}
	for idxOne < len(array1) && idxTwo < len(array2) {
		first, second := array1[idxOne], array2[idxTwo]
		if first < second {
			current = second - first
			idxOne += 1
		} else if second < first {
			current = first - second
			idxTwo += 1
		} else {
			return []int{first, second}
		}
		if smallest > current {
			smallest = current
			smallestPair = []int{first, second}
		}
	}
	return smallestPair
}

func main() {
	array1 := []int{-1, 5, 10, 20, 28, 3}
	array2 := []int{26, 134, 135, 15, 17}
	fmt.Println(SmallestDifference(array1, array2))
}
