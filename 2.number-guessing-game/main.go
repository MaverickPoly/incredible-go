package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

func main() {
	fmt.Println("Welcome to random number guessing game!")
	fmt.Println("You need to guess a random number between 0 and 100!")

	running := true
	currentRound := 0
	var again string

	for running {
		currentRound++
		fmt.Printf("\n-------- Round %d --------\n", currentRound)

		randomNumber := rand.IntN(100)
		var guess int
		attempts := 0

		for true {
			attempts++
			fmt.Print("Take a guess: ")

			if _, err := fmt.Scanln(&guess); err != nil {
				fmt.Println("Invalid guess!")
				continue
			}

			if guess == randomNumber {
				fmt.Println("Correct!")
				fmt.Printf("The random number was %d\n", randomNumber)
				fmt.Printf("It took you %d attempts to find it\n", attempts)
				break
			} else if guess > randomNumber {
				fmt.Println("Too high!")
			} else {
				fmt.Println("Too low!")
			}
		}

		fmt.Print("Do you want to play again(y/n): ")
		fmt.Scanln(&again)

		if strings.ToLower(again) == "n" {
			running = false
		}
	}

	fmt.Println("Adios")
}
