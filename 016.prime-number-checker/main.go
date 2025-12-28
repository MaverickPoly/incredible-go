package main

import "fmt"

func IsPrime1(number int) bool {
	if number < 2 {
		return false
	}

	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func IsPrime2(number int) bool {
	if number < 2 {
		return false
	}
	if number <= 3 {
		return true
	}
	if number%2 == 0 || number%3 == 0 {
		return false
	}

	for i := 5; i*i <= number; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	for i := range 15 {
		fmt.Printf("%-2d - %t\n", i, IsPrime1(i))
	}

	fmt.Println()

	for i := range 15 {
		fmt.Printf("%-2d - %t\n", i, IsPrime2(i))
	}
}
