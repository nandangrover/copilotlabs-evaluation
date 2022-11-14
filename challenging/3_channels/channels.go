package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(values chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	values <- value
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Go Channel Tutorial")

	values := make(chan int)
	defer close(values)

	go CalculateValue(values)

	value := <-values
	fmt.Println(value)
}
