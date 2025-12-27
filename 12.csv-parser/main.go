package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type User struct {
	Name string
	Age  int
	Job  string
}

const FILENAME = "data.csv"

var defaultUsers = []User{
	{"Bob", 39, "Accountant"},
	{"Alex", 21, "Sportsman"},
	{"Steve", 24, "Programmer"},
}

func UserToList(u User) []string {
	return []string{
		u.Name,
		strconv.Itoa(u.Age),
		u.Job,
	}
}

func WriteUsers(users []User) {
	var usersList [][]string

	for _, user := range users {
		usersList = append(usersList, UserToList(user))
	}

	file, err := os.Create(FILENAME)
	if err != nil {
		log.Fatalln("Error opening file!")
	}

	writer := csv.NewWriter(file)
	err = writer.WriteAll(usersList)
	if err != nil {
		log.Fatalln("Error writing file!")
	}

	fmt.Println("Successfully wrote users!")
}

func ReadUsers() []User {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatalln("Error opening file...")
	}

	reader := csv.NewReader(file)
	usersList, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Error reading csv!")
	}

	var users []User

	for _, user := range usersList {
		ageInt, err := strconv.Atoi(user[1])
		if err != nil {
			log.Fatalln("Invalid age!")
		}
		users = append(users, User{user[0], ageInt, user[2]})
	}

	return users
}

func main() {
	fmt.Println("Csv Parser program!")
	fmt.Println()

	fmt.Println("Writing CSV...")
	WriteUsers(defaultUsers)

	fmt.Println("Reading CSV...")
	users := ReadUsers()
	for _, user := range users {
		fmt.Println(user)
	}
}
