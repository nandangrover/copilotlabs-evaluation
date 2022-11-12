package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func validIPAddresses(string string) []string {
	ipAddressesFound := []string{}

	for i := 0; i < int(math.Min(float64(len(string)), 4.0)); i++ {
		currentIPAddressParts := []string{"", "", "", ""}

		currentIPAddressParts[0] = string[:i]
		if !isValidPart(currentIPAddressParts[0]) {
			continue
		}

		for j := i + 1; j < int(math.Min(float64(len(string)-i), 4.0)); j++ {
			currentIPAddressParts[1] = string[i:j]
			if !isValidPart(currentIPAddressParts[1]) {
				continue
			}

			for k := j + 1; k < int(math.Min(float64(len(string)-j), 4.0)); k++ {
				currentIPAddressParts[2] = string[j:k]
				currentIPAddressParts[3] = string[k:]
				if isValidPart(currentIPAddressParts[2]) && isValidPart(currentIPAddressParts[3]) {
					ipAddressesFound = append(ipAddressesFound, strings.Join(currentIPAddressParts, "."))
				}
			}
		}
	}

	return ipAddressesFound
}

func isValidPart(string string) bool {
	stringAsInt, _ := strconv.Atoi(string)
	if stringAsInt > 255 {
		return false
	}

	return len(string) == len(strconv.Itoa(stringAsInt))
}

func main() {
	string := "1921680"
	fmt.Println(validIPAddresses(string))
}
