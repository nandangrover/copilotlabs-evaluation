package main

import "fmt"

func pyramidPattern(n int) {
	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Printf(" ")
		}
		for k := 0; k < i; k++ {
			fmt.Printf("%d", i)
		}
		fmt.Printf("\n")
	}
}

func main() {
	pyramidPattern(5)
}
