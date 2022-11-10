package main

import "fmt"

func isLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isLeapYear(2000)) // True
	fmt.Println(isLeapYear(2001)) // False
	fmt.Println(isLeapYear(2004)) // True
	fmt.Println(isLeapYear(2100)) // False
}
