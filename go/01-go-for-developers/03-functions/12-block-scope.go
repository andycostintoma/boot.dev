package main

import (
	"fmt"
)

// Package (global) scoped variable
var age = 25

func sendEmail() {
	// Function scoped variable
	name := "Jon Snow"

	for i := 0; i < 3; i++ {
		// Loop scoped variable
		email := "snow@winterfell.net"
		fmt.Println(i, name, email) // All these variables are in scope here
	}

	// Trying to access loop scoped variable outside the loop will result in an error
	// fmt.Println(email) // This is not okay, email is out of scope here
}

func main() {
	// Accessing package scoped variable
	fmt.Println("Package scoped age:", age)

	// Explicit block to create a new scope
	{
		// This age variable shadows the package scoped one within this block
		age := 19
		fmt.Println("Block scoped age:", age) // This is okay, prints 19
	}

	// Accessing the package scoped variable again
	fmt.Println("Package scoped age:", age) // This is okay, prints 25

	// Calling the sendEmail function
	sendEmail()

	// Trying to access function scoped variable outside the function will result in an error
	// fmt.Println(name) // This is not okay, name is out of scope here
}
