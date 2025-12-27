package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filepath := flag.String("filepath", "", "Path of the file to count its words..")

	flag.Parse()

	if *filepath == "" {
		log.Fatal("Please provide a filepath...")
	}

	file, err := os.Open(*filepath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Fatal("File not found!")
		}

		log.Fatal("Error opening file...")
	}

	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		count += len(strings.Split(strings.TrimSpace(line), " "))
	}

	fmt.Println("Number of words: ")
	fmt.Println(count)
}
