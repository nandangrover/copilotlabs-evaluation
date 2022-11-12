package main

import (
	"fmt"
	"sort"
)

func taskAssignment(k int, tasks []int) [][]int {
	pairedTasks := [][]int{}
	taskDurationToIndices := getTaskDurationToIndices(tasks)
	sortedTasks := append([]int{}, tasks...)
	sort.Ints(sortedTasks)
	for idx := 0; idx < k; idx++ {
		task1Duration := sortedTasks[idx]
		indicesWithTask1Duration := taskDurationToIndices[task1Duration]
		task1Index := indicesWithTask1Duration[len(indicesWithTask1Duration)-1]
		indicesWithTask1Duration = indicesWithTask1Duration[:len(indicesWithTask1Duration)-1]
		taskDurationToIndices[task1Duration] = indicesWithTask1Duration
		task2SortedIndex := len(tasks) - 1 - idx
		task2Duration := sortedTasks[task2SortedIndex]
		indicesWithTask2Duration := taskDurationToIndices[task2Duration]
		task2Index := indicesWithTask2Duration[len(indicesWithTask2Duration)-1]
		indicesWithTask2Duration = indicesWithTask2Duration[:len(indicesWithTask2Duration)-1]
		taskDurationToIndices[task2Duration] = indicesWithTask2Duration
		pairedTasks = append(pairedTasks, []int{task1Index, task2Index})
	}
	return pairedTasks
}

func getTaskDurationToIndices(tasks []int) map[int][]int {
	taskDurationToIndices := map[int][]int{}
	for idx := 0; idx < len(tasks); idx++ {
		taskDuration := tasks[idx]
		if val, ok := taskDurationToIndices[taskDuration]; ok {
			taskDurationToIndices[taskDuration] = append(val, idx)
		} else {
			taskDurationToIndices[taskDuration] = []int{idx}
		}
	}
	return taskDurationToIndices
}

func main() {
	tasks := []int{1, 3, 5, 3, 1, 4}
	k := 3
	fmt.Println(taskAssignment(k, tasks))
}
