package main

import "fmt"

func knuthMorrisPrattAlgorithm(string, substring string) bool {
	pattern := buildPattern(substring)
	return doesMatch(string, substring, pattern)
}
func buildPattern(substring string) []int {
	pattern := make([]int, len(substring))
	for i := range pattern {
		pattern[i] = -1
	}
	j := 0
	i := 1
	for i < len(substring) {
		if substring[i] == substring[j] {
			pattern[i] = j
			i += 1
			j += 1
		} else if j > 0 {
			j = pattern[j-1] + 1
		} else {
			i += 1
		}
	}
	return pattern
}
func doesMatch(string, substring string, pattern []int) bool {
	i := 0
	j := 0
	for i+len(substring)-j <= len(string) {
		if string[i] == substring[j] {
			if j == len(substring)-1 {
				return true
			}
			i += 1
			j += 1
		} else if j > 0 {
			j = pattern[j-1] + 1
		} else {
			i += 1
		}
	}
	return false
}
func main() {
	fmt.Println(knuthMorrisPrattAlgorithm("aefoaefcdaefcdaed", "aefcdaed"))
	fmt.Println(knuthMorrisPrattAlgorithm("aefoaefceaefcdaet", "aefcdaed"))
}
