package main

import "fmt"

func alphabetTriangle() {
	for i := 'A'; i <= 'G'; i++ {
		for j := 'A'; j <= i; j++ {
			fmt.Printf("%c ", j)
		}
		fmt.Println()
	}
}

func main() {
	alphabetTriangle()
}
