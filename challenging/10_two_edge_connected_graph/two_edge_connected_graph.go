package main

import "fmt"

func twoEdgeConnectedGraph(edges [][]int) bool {
	if len(edges) == 0 {
		return true
	}

	arrivalTimes := make([]int, len(edges))
	for i := range arrivalTimes {
		arrivalTimes[i] = -1
	}
	startVertex := 0

	if getMinimumArrivalTimeOfAncestors(startVertex, -1, 0, arrivalTimes, edges) == -1 {
		return false
	}

	return areAllVerticesVisited(arrivalTimes)
}

func areAllVerticesVisited(arrivalTimes []int) bool {
	for _, time := range arrivalTimes {
		if time == -1 {
			return false
		}
	}

	return true
}

func getMinimumArrivalTimeOfAncestors(currentVertex int, parent int, currentTime int, arrivalTimes []int, edges [][]int) int {
	arrivalTimes[currentVertex] = currentTime

	minimumArrivalTime := currentTime

	for _, destination := range edges[currentVertex] {
		if arrivalTimes[destination] == -1 {
			minimumArrivalTime = min(minimumArrivalTime, getMinimumArrivalTimeOfAncestors(destination, currentVertex, currentTime+1, arrivalTimes, edges))
		} else if destination != parent {
			minimumArrivalTime = min(minimumArrivalTime, arrivalTimes[destination])
		}
	}

	// A bridge was detected, which means the graph isn't two-edge-connected.
	if minimumArrivalTime == currentTime && parent != -1 {
		return -1
	}

	return minimumArrivalTime
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	input1 := [][]int{
		{1, 2, 5},
		{0, 2},
		{0, 1, 3},
		{2, 4, 5},
		{3, 5},
		{0, 3, 4},
	}
	input2 := [][]int{
		{1, 2, 5},
		{0, 2},
		{0, 1, 3},
		{2, 4, 5},
		{3, 5},
		{0, 3, 4, 6},
		{5},
	}
	fmt.Println(twoEdgeConnectedGraph(input1))
	fmt.Println(twoEdgeConnectedGraph(input2))
}
