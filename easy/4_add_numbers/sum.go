package main

import "fmt"

// Add two numbers and print the result
func addNumbers(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(addNumbers(2, 3))
}
