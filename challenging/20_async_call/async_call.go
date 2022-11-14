package main

import (
	"fmt"
)

func ResolveAfter2Seconds() {
	return new Promise((resolve) => {
		setTimeout(() => {
			resolve("resolved");
		}, 2000);
	})
}

func AsyncCall() {
	fmt.Println("calling")
	result := ResolveAfter2Seconds()
	fmt.Println(result)
	// expected output: "resolved"
}

func main() {
	AsyncCall()
}