package main

import "fmt"

func main() {
	fmt.Println(countInversions([]int{8, 5, 2, 9, 5, 6, 3}))
	fmt.Println(countInversions([]int{1, 2, 3, 4, 5, 6, 7, 8}))
}

func countInversions(array []int) int {
	return countSubArrayInversions(array, 0, len(array))
}

func countSubArrayInversions(array []int, start int, end int) int {
	if end-start <= 1 {
		return 0
	}
	middle := start + (end-start)/2
	leftInversions := countSubArrayInversions(array, start, middle)
	rightInversions := countSubArrayInversions(array, middle, end)
	mergedArrayInversions := mergeSortAndCountInversions(array, start, middle, end)
	return leftInversions + rightInversions + mergedArrayInversions
}

func mergeSortAndCountInversions(array []int, start int, middle int, end int) int {
	sortedArray := []int{}
	left := start
	right := middle
	inversions := 0

	for left < middle && right < end {
		if array[left] <= array[right] {
			sortedArray = append(sortedArray, array[left])
			left++
		} else {
			inversions += middle - left
			sortedArray = append(sortedArray, array[right])
			right++
		}
	}
	sortedArray = append(sortedArray, array[left:middle]...)
	sortedArray = append(sortedArray, array[right:end]...)
	for idx, num := range sortedArray {
		array[start+idx] = num
	}
	return inversions
}
