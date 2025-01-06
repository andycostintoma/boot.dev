package main

import (
	"fmt"
)

// Mock function to get the length of an email
func getLength(email string) int {
	return len(email)
}

func main() {
	email := "test.go@example.com"

	// Using the initial statement in the if block
	if length := getLength(email); length < 1 {
		fmt.Println("Email is invalid: cannot be empty.")
	} else {
		fmt.Println("Email is valid!")
	}
}
