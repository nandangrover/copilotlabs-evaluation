package main

import "fmt"

// Area of triangle
func areaOfTriangle(base float64, height float64) float64 {
	return base * height / 2.0
}

func main() {
	fmt.Println(areaOfTriangle(2.0, 3.0))
	fmt.Println(areaOfTriangle(3.0, 4.0))
}
