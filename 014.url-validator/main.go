package main

import (
	"fmt"
	"net/url"
)

func isValidUrl(rawUrl string) bool {
	if rawUrl == "" {
		return false
	}

	parsed, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		return false
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return false
	}

	if parsed.Host == "" {
		return false
	}

	return true
}

func main() {
	testURLs := []string{
		"https://example.com",
		"http://localhost:8080/path",
		"https://sub.domain.co.uk/page?query=1",
		"ftp://invalid.com",
		"example.com",
		"http://",
		"https://",
		"not a url",
		"http://192.168.1.1:3000",
		"https://example.com:80",
		"http:// example.com",
	}

	for _, u := range testURLs {
		valid := isValidUrl(u)
		status := "✅"
		if !valid {
			status = "❌"
		}
		fmt.Printf("%s %-40s - %t\n", status, u, valid)
	}
}
