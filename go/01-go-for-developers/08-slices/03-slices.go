package main

import (
	"fmt"
)

// A simple function that modifies the elements of a slice
func modifySlice(s []int) {
	for i := range s {
		s[i] *= 2
	}
}

func main() {
	// Create an array
	array := [5]int{1, 2, 3, 4, 5}

	// Create a slice that references the array
	slice1 := array[:]

	// Create another slice referencing a portion of the array
	slice2 := array[1:4]

	fmt.Println("Original array:", array)
	fmt.Println("Slice1:", slice1)
	fmt.Println("Slice2:", slice2)

	// Modify the slice1
	modifySlice(slice1)

	fmt.Println("After modifying slice1:")
	fmt.Println("Array:", array)
	fmt.Println("Slice1:", slice1)
	fmt.Println("Slice2:", slice2)
}
