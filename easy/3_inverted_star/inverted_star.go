package main

import "fmt"

func invertedStar(n int) {
	for i := 0; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Printf(" ")
		}
		for k := 0; k < i; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func main() {
	invertedStar(5)
}
