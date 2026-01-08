package main

import (
	"fmt"
	"log"
)

var questions = []string{
	"What is the capital of France?",
	"Which planet is known as the Red Planet?",
	"In which year did World War 2 end?",
	"What is 2 + 2 * 2?",
	"Who is the founder of Python?",
}

var options = []string{
	"Moscow", "Paris", "London", "Delhi",
	"Earth", "Venus", "Mars", "Jupiter",
	"1941", "1944", "1946", "1945",
	"6", "8", "2", "4",
	"Guido Van Rossum", "Bill Gates", "Linus Torvalds", "Dennis Ritchie",
}

var correctAnswers = []int{
	2, 3, 4, 1, 1,
}

func main() {
	fmt.Println("Welcome to quiz app!")
	fmt.Println("======================")
	fmt.Println()

	correct := 0
	for index, question := range questions {
		fmt.Println(question)
		for i := range 4 {
			fmt.Printf("%d) %s\n", i+1, options[index*4+i])
		}
		fmt.Print("Answer: ")
		var answer int
		if _, err := fmt.Scanln(&answer); err != nil {
			log.Fatalln("Invalid option!")
		}

		if answer == correctAnswers[index] {
			correct++
			fmt.Println("Correct!")
		} else {
			fmt.Println("Wrong!")
			fmt.Printf("The answer was %s\n", options[index*4+correctAnswers[index]-1])
		}

		fmt.Println("----------------")
	}

	fmt.Printf("Your result: %d/%d\n", correct, len(correctAnswers))
}
