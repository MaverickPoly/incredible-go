package main

import (
	"fmt"
	"log"
	"math/rand/v2"
)

func main() {
	var diceNum, sides int

	fmt.Println("Dice rolling simulator!")

	fmt.Print("How many dice: ")
	if _, err := fmt.Scanln(&diceNum); err != nil {
		log.Fatal("Invalid number of dice...")
	}

	fmt.Print("How many sides: ")
	if _, err := fmt.Scanln(&sides); err != nil {
		log.Fatal("Invalid number of sides...")
	}

	var dices []int
	for range diceNum {
		dices = append(dices, rand.IntN(sides)+1)
	}

	total := 0
	for i, die := range dices {
		total += die
		fmt.Printf("Die %d: %d\n", i+1, die)
	}

	fmt.Println("Total: ", total)
}
