package main

import "fmt"

// Check prime number
func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPrime(1))  // False
	fmt.Println(isPrime(17)) // True
}
