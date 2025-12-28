package main

import (
	"fmt"
	"regexp"
)

func isValidEmail(email string) bool {
	if email == "" {
		return false
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,24}$`)
	return emailRegex.MatchString(email)
}

func main() {
	testEmails := []string{
		"user@example.com",
		"test.email+tag@domain.co.uk",
		"user@localhost",
		"plainaddress",
		"@missingdomain.com",
		"missing@.com",
		"spaces @example.com",
		"user@domain",
		"valid@sub.domain.com",
		"123@456.789",
		"user_name@domain-hyphen.org",
	}

	fmt.Println("Email Validator Results:")
	fmt.Println("========================")
	for _, email := range testEmails {
		valid := isValidEmail(email)
		status := "✅"
		if !valid {
			status = "❌"
		}
		fmt.Printf("%s %-35s -> %t\n", status, email, valid)
	}
}
