package main

import (
	"fmt"
)

func main() {
	// Create a slice using make with specified length and capacity
	sliceWithMake := make([]int, 5, 10)
	fmt.Println("Slice created with make, length:", len(sliceWithMake), ", capacity:", cap(sliceWithMake))

	// Print the slice to see the zero values
	fmt.Println("Slice with make:", sliceWithMake)

	// Create a slice using make with only length (capacity defaults to length)
	sliceDefaultCap := make([]int, 5)
	fmt.Println("Slice with default capacity, length:", len(sliceDefaultCap), ", capacity:", cap(sliceDefaultCap))

	// Print the slice to see the zero values
	fmt.Println("Slice with default capacity:", sliceDefaultCap)

	// Create a slice using a slice literal
	sliceLiteral := []string{"I", "love", "Go"}
	fmt.Println("Slice created with literal, length:", len(sliceLiteral), ", capacity:", cap(sliceLiteral))

	// Print the slice values
	fmt.Println("Slice literal:", sliceLiteral)

	// Demonstrate automatic growth of a slice
	originalSlice := []int{1, 2, 3}
	fmt.Println("Original slice, length:", len(originalSlice), ", capacity:", cap(originalSlice))
	fmt.Println("Original slice:", originalSlice)

	// Append elements to the slice
	originalSlice = append(originalSlice, 4, 5, 6)
	fmt.Println("Modified slice after append, length:", len(originalSlice), ", capacity:", cap(originalSlice))
	fmt.Println("Modified slice:", originalSlice)

	// Show how capacity changes as slice grows
	for i := 0; i < 10; i++ {
		originalSlice = append(originalSlice, i)
		fmt.Printf("After appending %d elements, length: %d, capacity: %d\n", i+4, len(originalSlice), cap(originalSlice))
	}
}
