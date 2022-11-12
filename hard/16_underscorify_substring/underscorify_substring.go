package main

import (
	"fmt"
	"strings"
)

func GetLocations(inputString string, substring string) [][]int {
	locations := [][]int{}
	startIdx := 0
	for startIdx < len(inputString) {
		nextIdx := strings.Index(inputString[startIdx:], substring)
		if nextIdx != -1 {
			locations = append(locations, []int{startIdx + nextIdx, startIdx + nextIdx + len(substring)})
			startIdx += nextIdx + 1
		} else {
			break
		}
	}
	return locations
}

func Collapse(locations [][]int) [][]int {
	if len(locations) == 0 {
		return locations
	}
	newLocations := [][]int{locations[0]}
	previous := newLocations[0]
	for i := 1; i < len(locations); i++ {
		current := locations[i]
		if current[0] <= previous[1] {
			previous[1] = current[1]
		} else {
			newLocations = append(newLocations, current)
			previous = current
		}
	}
	return newLocations
}

func Underscorify(inputString string, locations [][]int) string {
	locationsIdx := 0
	stringIdx := 0
	inBetweenUnderscores := false
	finalChars := []string{}
	i := 0
	for stringIdx < len(inputString) && locationsIdx < len(locations) {
		if stringIdx == locations[locationsIdx][i] {
			finalChars = append(finalChars, "_")
			inBetweenUnderscores = !inBetweenUnderscores
			if !inBetweenUnderscores {
				locationsIdx += 1
			}
			i = 0
			if i == 1 {
				i = 0
			} else {
				i = 1
			}
		}
		finalChars = append(finalChars, string(inputString[stringIdx]))
		stringIdx += 1
	}
	if locationsIdx < len(locations) {
		finalChars = append(finalChars, "_")
	} else if stringIdx < len(inputString) {
		finalChars = append(finalChars, inputString[stringIdx:]...)
	}
	return strings.Join(finalChars, "")
}

func UnderscorifySubstring(inputString string, substring string) string {
	locations := Collapse(GetLocations(inputString, substring))
	return Underscorify(inputString, locations)
}

func main() {
	fmt.Println(UnderscorifySubstring("testthis is a testtest to see if testestest it works", "test"))
}
