package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
)

const FILENAME = "data.json"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Category  string `json:"category"`
	Completed bool   `json:"completed"`
}

var todos []Todo

// Utility functions
func SaveTodos() {
	file, err := os.Create(FILENAME)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Could not open file to write!")
		os.Exit(1)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(todos); err != nil {
		fmt.Fprintln(os.Stderr, "Error: Could not encode todos!")
		os.Exit(1)
	}
}

func LoadTodos() {
	file, err := os.Open(FILENAME)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Could not open file to read!")
		os.Exit(1)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&todos); err != nil {
		fmt.Fprintln(os.Stderr, "Error: Could not decode todos!")
		os.Exit(1)
	}
}

// Utils
func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("   add <title> <category>    -- Add a new Todo")
	fmt.Println("   list                      -- List all Todos")
	fmt.Println("   delete <id>               -- Delete Todo")
	fmt.Println("   mark <id>                 -- Mark Todo as Completed")
	fmt.Println("   unmark <id>               -- Mark Todo as Uncompleted")
}

func GetIndex() int {
	if len(todos) == 0 {
		return 1
	}

	index := 0
	for _, todo := range todos {
		if todo.ID > index {
			index = todo.ID
		}
	}

	return index + 1
}

// Main Functionalities
func ListTodos() {
	fmt.Println("-------- Todos --------")
	for _, todo := range todos {
		if todo.Completed {
			text := fmt.Sprintf("%d. %s - %s", todo.ID, todo.Title, todo.Category)
			fmt.Println(Green + text + Reset)
		} else {
			text := fmt.Sprintf("%d. %s - %s", todo.ID, todo.Title, todo.Category)
			fmt.Println(Red + text + Reset)
		}
	}
}

func AddTodo(title string, category string) {
	id := GetIndex()
	todo := Todo{id, title, category, false}

	todos = append(todos, todo)
	SaveTodos()
}

func MarkDone(id int) {
	foundIndex := -1
	for index, currentTodo := range todos {
		if currentTodo.ID == id {
			foundIndex = index
			break
		}
	}

	if foundIndex == -1 {
		fmt.Fprintln(os.Stderr, "Error: Invalid todo id!")
		os.Exit(1)
	}

	todos[foundIndex].Completed = true
	SaveTodos()
}

func UnmarkDone(id int) {
	foundIndex := -1
	for index, currentTodo := range todos {
		if currentTodo.ID == id {
			foundIndex = index
		}
	}

	if foundIndex == -1 {
		fmt.Fprintln(os.Stderr, "Error: Invalid todo id!")
		os.Exit(1)
	}

	todos[foundIndex].Completed = false
	SaveTodos()
}

func DeleteTodo(id int) {
	foundIndex := -1
	for index, currentTodo := range todos {
		if currentTodo.ID == id {
			foundIndex = index
		}
	}

	if foundIndex == -1 {
		fmt.Fprintln(os.Stderr, "Error: Invalid todo id!")
		os.Exit(1)
	}

	todos = slices.Delete(todos, foundIndex, foundIndex+1)
	SaveTodos()
}

// Main Function
func main() {
	if len(os.Args) < 2 {
		PrintUsage()
		os.Exit(1)
	}

	LoadTodos()

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Error: Missing todo title!")
			os.Exit(1)
		}
		if len(os.Args) < 4 {
			fmt.Fprintln(os.Stderr, "Error: Missing todo category!")
			os.Exit(1)
		}

		title := os.Args[2]
		category := os.Args[3]

		AddTodo(title, category)
	case "list":
		ListTodos()
	case "delete":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Error: Missing todo id to delete!")
			os.Exit(1)
		}

		idStr := os.Args[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: Invalid todo id!")
			os.Exit(1)
		}

		DeleteTodo(id)
	case "done":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Error: Missing todo id to mark as done!")
			os.Exit(1)
		}

		idStr := os.Args[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: Invalid todo id!")
			os.Exit(1)
		}

		MarkDone(id)
	case "undone":
		if len(os.Args) < 3 {
			fmt.Fprintln(os.Stderr, "Error: Missing todo id to mark as undone!")
			os.Exit(1)
		}

		idStr := os.Args[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: Invalid todo id!")
			os.Exit(1)
		}

		UnmarkDone(id)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		PrintUsage()
		os.Exit(1)
	}
}
