package main

import (
	"flag"
	"fmt"
)

func generateFibonacci(n int) []int {
	res := []int{1, 1}

	a := 1
	b := 1
	for len(res) != n {
		temp := a + b
		a = b
		b = temp
		res = append(res, b)
	}

	return res
}

func main() {
	nFlag := flag.Int("n", 2, "The number of fibonacci numbers to generate!")

	flag.Parse()

	numbers := generateFibonacci(*nFlag)

	fmt.Printf("%+v\n", numbers)
}
