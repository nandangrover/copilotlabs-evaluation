package main

import (
	"fmt"
)

func numberOfWaysToTraverseGraph(width int, height int) int {
	numberOfWays := make([][]int, height+1)
	for i := 0; i < height+1; i++ {
		numberOfWays[i] = make([]int, width+1)
	}
	for widthIdx := 1; widthIdx < width+1; widthIdx++ {
		for heightIdx := 1; heightIdx < height+1; heightIdx++ {
			if widthIdx == 1 || heightIdx == 1 {
				numberOfWays[heightIdx][widthIdx] = 1
			} else {
				waysLeft := numberOfWays[heightIdx][widthIdx-1]
				waysUp := numberOfWays[heightIdx-1][widthIdx]
				numberOfWays[heightIdx][widthIdx] = waysLeft + waysUp
			}
		}
	}
	return numberOfWays[height][width]
}

func main() {
	width := 4
	height := 3
	fmt.Println(numberOfWaysToTraverseGraph(width, height))
}
