package main

import "fmt"

func main() {
	height := 5

	// Simple if statement without parentheses around the condition
	if height > 4 {
		fmt.Println("You are tall enough!")
	}

	// Using else if and else for additional conditions
	if height > 6 {
		fmt.Println("You are super tall!")
	} else if height > 4 {
		fmt.Println("You are tall enough!")
	} else {
		fmt.Println("You are not tall enough!")
	}
}
