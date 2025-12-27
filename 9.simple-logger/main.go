package main

import (
	"fmt"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Bold   = "\033[1m"
)

func success(message string) {
	text := fmt.Sprintf("SUCCESS - %v: %s", time.Now().Format("15:04:05"), message)
	fmt.Println(Green + text + Reset)
}

func warning(message string) {
	text := fmt.Sprintf("WARNING - %v: %s", time.Now().Format("15:04:05"), message)
	fmt.Println(Yellow + text + Reset)
}

func info(message string) {
	text := fmt.Sprintf("INFO - %v: %s", time.Now().Format("15:04:05"), message)
	fmt.Println(Blue + text + Reset)
}

func error(message string) {
	text := fmt.Sprintf("ERROR - %v: %s", time.Now().Format("15:04:05"), message)
	fmt.Println(Red + text + Reset)
}

func main() {
	success("Fetched data!")
	warning("You are running out of memory!")
	info("New User joined chat!")
	error("Your computer has a virus!")
}
