package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

const FILENAME = "data.json"

func main() {
	// 1. Write to json
	config := Config{
		"Example App", 42,
	}

	file, err := os.Create(FILENAME)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Failed to create file!")
		os.Exit(1)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		fmt.Fprintln(os.Stderr, "Error: Failed to write json!")
		os.Exit(1)
	}

	fmt.Println("✅ Wrote data to", FILENAME)

	// 2. Read Data From Json
	file, err = os.Open(FILENAME)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Failed to open file!")
		os.Exit(1)
	}

	var loadedConfig Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&loadedConfig); err != nil {
		fmt.Fprintln(os.Stderr, "Error: Failed to read json!")
		os.Exit(1)
	}

	fmt.Printf("✅ Read data: %+v\n", loadedConfig)
}
