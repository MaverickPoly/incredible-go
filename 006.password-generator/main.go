package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
)

func generatePassword(size int, combs string) string {
	var password []byte

	for range size {
		password = append(password, combs[rand.Intn(len(combs))])
	}

	return string(password)
}

func main() {
	size := flag.Int("size", 8, "Size of the password!")
	uppercase := flag.Bool("upper", false, "Include Uppercase Letters")
	digits := flag.Bool("digits", false, "Include Digits")
	punc := flag.Bool("punc", false, "Include Punctuation")

	flag.Parse()

	if *size > 100 {
		log.Fatal("Size is too large!")
	}

	combs := "abcdefghijklmnopqrstuvwxyz"
	if *uppercase {
		combs += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if *digits {
		combs += "0123456789"
	}
	if *punc {
		combs += "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	}

	password := generatePassword(*size, combs)

	fmt.Println("Password:")
	fmt.Println(password)
}
