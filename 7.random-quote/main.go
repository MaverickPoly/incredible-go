package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const URL = "https://api.api-ninjas.com/v2/randomquotes?categories=success,wisdom"
const API_KEY = ""

type Quote struct {
	Quote      string   `json:"quote"`
	Author     string   `json:"author"`
	Work       string   `json:"work"`
	Categories []string `json:"categories"`
}

func main() {
	client := http.Client{Timeout: 10 * time.Second}

	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatal("Invalid URL!")
	}

	request.Header.Add("X-Api-Key", API_KEY)

	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Error sending request: ", err.Error())
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error parsing body: ", err.Error())
	}

	var quotes []Quote
	if err := json.Unmarshal(body, &quotes); err != nil {
		log.Fatal("Error parsing json: ", err.Error())
	}

	quote := quotes[0]
	fmt.Println("Quote:", quote.Quote)
	fmt.Println("Author:", quote.Author)
	fmt.Println("Work:", quote.Work)

	fmt.Println("Categories:")
	for _, category := range quote.Categories {
		fmt.Println("- ", category)
	}
}
