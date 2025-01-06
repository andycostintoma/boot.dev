package main

import (
	"fmt"
	"reflect"
)

// Optimized struct with fields ordered by size (largest to smallest)
type statsOptimized struct {
	Reach    uint16 // 2 bytes
	NumPosts uint8  // 1 byte
	NumLikes uint8  // 1 byte
}

// Non-optimized struct with inefficient field ordering
type statsNonOptimized struct {
	NumPosts uint8  // 1 byte
	Reach    uint16 // 2 bytes (padding will be added here)
	NumLikes uint8  // 1 byte
}

func main() {
	// Get the memory size of the optimized struct
	typOptimized := reflect.TypeOf(statsOptimized{})
	fmt.Printf("Optimized struct is %d bytes\n", typOptimized.Size())

	// Get the memory size of the non-optimized struct
	typNonOptimized := reflect.TypeOf(statsNonOptimized{})
	fmt.Printf("Non-optimized struct is %d bytes\n", typNonOptimized.Size())
}
