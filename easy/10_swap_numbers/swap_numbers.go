package main

import "fmt"

func swap_numbers(a, b int) []int {
	return []int{b, a}
}

func main() {
	fmt.Println(swap_numbers(2, 3))
}
