package main

import (
	"fmt"
)

func main() {
	// Creating a map using the make() function
	ages := make(map[string]int)
	ages["John"] = 37
	ages["Mary"] = 24
	// Overwriting the value associated with the key "Mary"
	ages["Mary"] = 21

	// Printing the map after using make()
	fmt.Println("Map created using make():")
	for name, age := range ages {
		fmt.Printf("%s: %d\n", name, age)
	}

	// Creating a map using a literal
	ages = map[string]int{
		"John": 37,
		"Mary": 21,
	}

	// Printing the map created using a literal
	fmt.Println("\nMap created using a literal:")
	for name, age := range ages {
		fmt.Printf("%s: %d\n", name, age)
	}

	// Using the len() function to get the number of key/value pairs in the map
	length := len(ages)
	fmt.Printf("\nNumber of key/value pairs in the map: %d\n", length)

	// Accessing map values and checking if a key exists
	age, exists := ages["John"]
	if exists {
		fmt.Printf("\nJohn's age is: %d\n", age)
	} else {
		fmt.Println("\nJohn's age is not found in the map")
	}

	// Trying to access a key that doesn't exist
	age, exists = ages["Paul"]
	if exists {
		fmt.Printf("\nPaul's age is: %d\n", age)
	} else {
		fmt.Println("\nPaul's age is not found in the map")
	}
}
