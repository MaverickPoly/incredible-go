package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  go run . <seconds>  # The number of seconds to run countdown")
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
		os.Exit(1)
	}

	secondsStr := os.Args[1]
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid seconds!")
		os.Exit(1)
	}

	for seconds > 0 {
		fmt.Println(seconds)
		time.Sleep(time.Second)
		seconds--
	}
	fmt.Println()
	fmt.Println("TIME IS UP!!!")
}
