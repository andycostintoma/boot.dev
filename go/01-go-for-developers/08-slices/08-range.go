package main

import "fmt"

func main() {
	// Create a slice of strings
	fruits := []string{"apple", "banana", "grape"}

	// Iterate over the slice using range
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	// If you only need the element and not the index, you can discard the index using an underscore
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}
