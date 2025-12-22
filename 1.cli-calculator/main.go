package main

import "fmt"

func main() {
	fmt.Println("Welcome to a simple CLI Calculator!")

	var option int
	result := 0.0
	var number1, number2 float64

	for true {
		fmt.Println("\n1. Addition\n2. Subtraction\n3. Multiplication\n4. Division\n5. Exit")
		fmt.Print("Select an option: ")

		if _, err := fmt.Scanln(&option); err != nil {
			fmt.Println("Invalid option!")
			continue
		}

		if option <= 0 || option > 5 {
			fmt.Println("Invalid option!")
			continue
		}

		if option == 5 {
			break
		}

		fmt.Print("Enter the first number: ")
		if _, err := fmt.Scanln(&number1); err != nil {
			fmt.Println("Invalid number!")
			continue
		}
		fmt.Print("Enter the second number: ")
		if _, err := fmt.Scanln(&number2); err != nil {
			fmt.Println("Invalid number!")
			continue
		}

		switch option {
		case 1:
			result = number1 + number2
			fmt.Printf("%g + %g = %g", number1, number2, result)
		case 2:
			result = number1 - number2
			fmt.Printf("%g - %g = %g", number1, number2, result)
		case 3:
			result = number1 * number2
			fmt.Printf("%g * %g = %g", number1, number2, result)
		case 4:
			result = number1 / number2
			fmt.Printf("%g / %g = %g", number1, number2, result)
		}
	}

	fmt.Println("Take care..")
}
