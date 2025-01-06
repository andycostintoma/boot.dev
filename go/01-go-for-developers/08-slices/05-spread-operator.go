package main

import (
	"fmt"
)

// Variadic function to concatenate strings
func concat(strs ...string) string {
	final := ""
	// strs is just a slice of strings
	for i := 0; i < len(strs); i++ {
		final += strs[i]
	}
	return final
}

// Variadic function to print each string in separate lines
func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
	// Example of using variadic function directly
	final := concat("Hello ", "there ", "friend!")
	fmt.Println(final)
	// Output: Hello there friend!

	// Example of using the spread operator to pass a slice
	names := []string{"bob", "sue", "alice"}
	printStrings(names...)
	// Output:
	// bob
	// sue
	// alice

	// Using variadic function from standard library
	fmt.Println("This", "is", "a", "variadic", "function")
	// Output: This is a variadic function
}
