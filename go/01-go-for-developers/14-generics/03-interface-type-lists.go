package main

import (
	"fmt"
)

// Ordered is a type constraint that matches any ordered type.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | string
}

// Generic function to find the maximum value in a slice of ordered types
func findMax[T Ordered](vals []T) T {
	if len(vals) == 0 {
		var zero T
		return zero
	}
	maximum := vals[0]
	for _, val := range vals[1:] {
		if val > maximum {
			maximum = val
		}
	}
	return maximum
}

func main() {
	ints := []int{1, 5, 3, 9, 2}
	floats := []float64{2.4, 3.6, 1.8, 4.5}
	strings := []string{"apple", "banana", "cherry"}

	maxInt := findMax(ints)
	maxFloat := findMax(floats)
	maxString := findMax(strings)

	fmt.Println("Max int:", maxInt)
	fmt.Println("Max float:", maxFloat)
	fmt.Println("Max string:", maxString)
}
