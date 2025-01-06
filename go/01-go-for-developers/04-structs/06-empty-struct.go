package main

import (
	"fmt"
	"unsafe"
)

// Define a named empty struct
type emptyStruct struct{}

func main() {
	// Create an instance of an anonymous empty struct
	emptyAnonymous := struct{}{}

	// Create an instance of a named empty struct
	emptyNamed := emptyStruct{}

	// Print the zero memory usage of empty structs
	fmt.Printf("Size of anonymous empty struct: %d bytes\n", unsafe.Sizeof(emptyAnonymous))
	fmt.Printf("Size of named empty struct: %d bytes\n", unsafe.Sizeof(emptyNamed))
}
