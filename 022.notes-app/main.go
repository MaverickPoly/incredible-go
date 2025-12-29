package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"
)

type Note struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

const FILENAME = "data.json"

var notes []Note

func saveNotes() {
	file, err := os.Create(FILENAME)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Failed to open file to write!")
		os.Exit(1)
	}

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(notes); err != nil {
		fmt.Fprintln(os.Stderr, "Error: Failed to encode json!")
		os.Exit(1)
	}
}

func loadNotes() {
	file, err := os.Open(FILENAME)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Failed to open file to read!")
		os.Exit(1)
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&notes); err != nil {
		fmt.Fprintln(os.Stderr, "Error: Failed to decode json!")
		os.Exit(1)
	}
}

func getId() int {
	id := 0
	for _, note := range notes {
		if note.ID > id {
			id = note.ID
		}
	}

	return id + 1
}

func printUsage() {
	fmt.Println("Usage:")

	fmt.Println("   list                       # List all notes")
	fmt.Println("   add \"Title\" \"Content\"  # Add new note")
	fmt.Println("   view <id>                  # View note")
	fmt.Println("   delete <id>                # Delete note")
}

func printNote(note Note) {
	created_at := note.CreatedAt.Format("02 15-04-05")
	fmt.Printf("%d. %s - %s\n", note.ID, note.Title, created_at)
	fmt.Println(note.Content)
}

func listNotes() {
	fmt.Println("======== Notes ========")
	for _, note := range notes {
		printNote(note)
	}
}

func addNote(title string, content string) {
	id := getId()
	note := Note{id, title, content, time.Now()}
	notes = append(notes, note)
	saveNotes()
}

func viewNote(id int) {
	for _, note := range notes {
		if note.ID == id {
			printNote(note)
			return
		}
	}

	fmt.Fprintln(os.Stderr, "Error: Invalid note id!")
	os.Exit(1)
}

func deleteNote(id int) {
	for index, note := range notes {
		if note.ID == id {
			notes = slices.Delete(notes, index, index+1)
			saveNotes()
			return
		}
	}

	fmt.Fprintln(os.Stderr, "Error: Invalid note id!")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	loadNotes()
	command := os.Args[1]

	switch command {
	case "list":
		listNotes()
	case "add":
		if len(os.Args) < 4 {
			fmt.Fprintln(os.Stderr, "   add \"Title\" \"Content\"  # Add new note")
			os.Exit(1)
		}

		title := os.Args[2]
		content := os.Args[3]
		addNote(title, content)
	case "view":
		if len(os.Args) < 3 {

			fmt.Println("   view <id>                  # View note")
			os.Exit(1)
		}

		idStr := os.Args[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid note id!")
			os.Exit(1)
		}

		viewNote(id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("   delete <id>                # Delete note")
			os.Exit(1)
		}
		idStr := os.Args[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid note id!")
			os.Exit(1)
		}

		deleteNote(id)
	default:
		printUsage()
		os.Exit(1)
	}
}
