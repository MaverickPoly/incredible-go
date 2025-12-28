package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:")
		fmt.Println("  go run . <filepath>")
		os.Exit(1)
	}

	filepath := os.Args[1]
	stat, err := os.Stat(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v\n", err.Error())
		os.Exit(1)
	}

	bytes := stat.Size()

	var size string
	if bytes < 1024 {
		size = fmt.Sprintf("%d B", bytes)
	} else if bytes < 1024*1024 {
		size = fmt.Sprintf("%.1f KB", float64(bytes)/1024)
	} else if bytes < 1024*1024*1024 {
		size = fmt.Sprintf("%.1f MB", float64(bytes)/(1024*1024))
	} else {
		size = fmt.Sprintf("%.1f GB", float64(bytes)/(1024*1024*1024))
	}

	fmt.Printf("File: %s\n", filepath)
	fmt.Printf("Size: %s\n", size)
}
