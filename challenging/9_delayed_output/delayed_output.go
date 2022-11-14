package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleep(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func log(s string, delay int, randomized bool) {
	for _, c := range s {
		fmt.Printf("%c", c)
		if randomized {
			sleep(rand.Intn(delay))
		} else {
			sleep(delay)
		}
	}
	fmt.Printf("\n")
}

func main() {
	log("Hello, world!", 200, true)
}
