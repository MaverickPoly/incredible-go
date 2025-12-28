package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	metersToFeet = 3.28084
	feetToMeters = 1 / metersToFeet
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("    go run . m2ft <meters>    # Meters to Feet")
	fmt.Println("    go run . ft2m <feets>     # Feet to Meters")
}

func main() {
	if len(os.Args) != 3 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	amountStr := os.Args[2]

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("")
		os.Exit(1)
	}

	switch command {
	case "m2ft":
		result := amount * metersToFeet
		fmt.Printf("%g m = %g ft\n", amount, result)
	case "ft2m":
		result := amount * feetToMeters
		fmt.Printf("%g ft = %g m\n", amount, result)
	default:
		printUsage()
		os.Exit(1)
	}
}
