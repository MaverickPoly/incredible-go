package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const FILENAME = "data.txt"

func WriteFile(lines []string) {
	file, err := os.Create(FILENAME)
	if err != nil {
		log.Fatal("Error creating file!")
	}

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			log.Fatal("Error writing file!")
		}
	}
}

func ReadFile() string {
	content, err := os.ReadFile(FILENAME)
	if err != nil {
		log.Fatal("Error reading file...")
	}

	return string(content)
}

func main() {
	// Reading Initial Data
	fmt.Println("Welcome to file reader/writer!")

	initialContent := ReadFile()
	fmt.Println("Initial Content:")
	fmt.Println(initialContent)

	fmt.Println("\n----------")

	// Writing Data
	fmt.Println("Enter a text to write to a file:")
	scanner := bufio.NewScanner(os.Stdin)
	var allLines []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" && len(allLines) > 0 {
			break
		}

		allLines = append(allLines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input...")
		return
	}
	WriteFile(allLines)

	// Reading AfterWards data
	fmt.Println("\n----------")

	endContent := ReadFile()
	fmt.Println("End Content:")
	fmt.Println(endContent)
}
