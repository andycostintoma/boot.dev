package main

import (
	"fmt"
)

func main() {
	// Creating a map using a literal
	m := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}

	// Inserting an element
	m["Charlie"] = 29

	// Getting and printing an element
	age := m["Alice"]
	fmt.Printf("Alice's age: %d\n", age)

	// Checking if a key exists
	age, ok := m["Bob"]
	if ok {
		fmt.Printf("Bob's age: %d\n", age)
	} else {
		fmt.Println("Bob is not in the map")
	}

	// Trying to get an element for a key that doesn't exist
	age, ok = m["David"]
	if ok {
		fmt.Printf("David's age: %d\n", age)
	} else {
		fmt.Printf("David is not in the map; default age: %d\n", age)
	}

	// Deleting an element
	delete(m, "Alice")

	// Checking if the key "Alice" still exists after deletion
	age, ok = m["Alice"]
	if ok {
		fmt.Printf("Alice's age: %d\n", age)
	} else {
		fmt.Println("Alice is not in the map")
	}

	// Printing the entire map to show the final state
	fmt.Println("Final map:")
	for name, age := range m {
		fmt.Printf("%s: %d\n", name, age)
	}
}
