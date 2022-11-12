package main

import (
	"fmt"
	"math"
	"sort"
)

func get_statistics(input_list []int) map[string]float64 {
	sorted_input := sort.IntSlice(input_list)
	sort.Sort(sorted_input)
	input_length := len(sorted_input)

	mean := sum(sorted_input) / input_length

	middle_idx := (len(sorted_input) - 1) // 2
	median := sorted_input[middle_idx]

	if input_length%2 == 0 {
		middle_number_1 := sorted_input[middle_idx]
		middle_number_2 := sorted_input[middle_idx+1]
		median = (middle_number_1 + middle_number_2) / 2
	}

	number_counts := map[int]int{}
	for _, x := range sorted_input {
		number_counts[x] += 1
	}
	mode := 0
	for unique_number, count := range number_counts {
		if count > number_counts[mode] {
			mode = unique_number
		}
	}

	sample_variance := 0.0
	for _, number := range sorted_input {
		sample_variance += math.Pow(number-mean, 2) / (input_length - 1)
	}

	sample_standard_deviation := math.Pow(sample_variance, 0.5)

	mean_standard_error := sample_standard_deviation / math.Pow(float64(input_length), 0.5)
	z_score_standard_error := 1.96 * mean_standard_error
	mean_confidence_interval := []float64{mean - z_score_standard_error, mean + z_score_standard_error}

	return map[string]float64{
		"mean":                      mean,
		"median":                    median,
		"mode":                      mode,
		"sample_variance":           sample_variance,
		"sample_standard_deviation": sample_standard_deviation,
		"mean_confidence_interval":  mean_confidence_interval,
	}
}

func main() {
	fmt.Println(get_statistics([]int{2, 1, 3, 4, 4, 5, 6, 7}))
}
