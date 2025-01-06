package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// UTF-8 encoded string with zany and multi-byte characters
	str := "Hello, 世界! 😊"
	fmt.Printf("Original String: %s\n", str)

	// Length of the string in bytes
	fmt.Printf("Length in bytes: %d\n", len(str))

	// Length of the string in runes (characters)
	fmt.Printf("Length in runes: %d\n", utf8.RuneCountInString(str))

	// Iterating over the string by runes
	fmt.Println("Characters (runes) in the string:")
	for i, runeValue := range str {
		fmt.Printf("Character %d: %c (Unicode: %U)\n", i, runeValue, runeValue)
	}

	// Explicitly converting string to runes
	runes := []rune(str)
	fmt.Println("Rune Slice:", runes)
	fmt.Printf("Length of rune slice: %d\n", len(runes))

	// Accessing individual runes
	fmt.Printf("First rune: %c (Unicode: %U)\n", runes[0], runes[0])
	fmt.Printf("Third rune: %c (Unicode: %U)\n", runes[2], runes[2])

	// Substring using runes
	substring := string(runes[7:10])
	fmt.Printf("Substring (runes 7 to 9): %s\n", substring)
}
