package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Random number from 1 - 10
	rand.Seed(time.Now().UnixNano())
	numberToGuess := rand.Intn(10) + 1

	// This variable is used to determine if the app should continue prompting the user for input
	var foundCorrectNumber bool

	for !foundCorrectNumber {
		// Get user input
		var guess int
		fmt.Print("Guess a number from 1 to 10: ")
		fmt.Scan(&guess)

		// Compare the guess to the secret answer and let the user know.
		if guess == numberToGuess {
			fmt.Println("Congrats, you got it!")
			foundCorrectNumber = true
		} else {
			fmt.Println("Sorry, guess again!")
		}
	}
}
