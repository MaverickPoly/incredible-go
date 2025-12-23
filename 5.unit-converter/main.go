package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Unit Converter!")

	for true {
		fmt.Println()
		fmt.Println("1. Celsius to Fahrenheit")
		fmt.Println("2. Fahrenheit to Celsius")
		fmt.Println("3. Exit")

		var option int
		fmt.Print("Select an option: ")
		if _, err := fmt.Scanln(&option); err != nil {
			fmt.Println("Invalid temperature!")
			continue
		}

		if option < 1 || option > 3 {
			fmt.Println("Invalid temperature!")
			continue
		}

		if option == 3 {
			break
		}

		var fromString, toString string
		if option == 1 {
			fromString = "Celsius"
			toString = "Fahrenheit"
		} else {
			fromString = "Fahrenheit"
			toString = "Celsius"
		}

		var temperature, result float64
		fmt.Printf("Enter temperature in %s: ", fromString)
		if _, err := fmt.Scanln(&temperature); err != nil {
			fmt.Println("Invalid temperature!")
			continue
		}

		if option == 1 {
			result = (temperature * 9 / 5) + 32
		} else {
			result = (temperature - 32) * 5 / 9
		}

		fmt.Printf("%.1f %s = %.1f %s\n", temperature, fromString, result, toString)
	}

	fmt.Println("Bye...")
}
