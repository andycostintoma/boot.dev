package main

import (
	"fmt"
)

func main() {
	// Declare and initialize an array of prime numbers
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Create a slice from the array
	mySlice := primes[1:4] // This slice will contain {3, 5, 7}

	// Print the elements of the slice
	fmt.Println("mySlice:", mySlice)

	// Other ways to create slices from the array
	slice1 := primes[:3] // {2, 3, 5}
	slice2 := primes[2:] // {5, 7, 11, 13}
	slice3 := primes[:]  // {2, 3, 5, 7, 11, 13}

	// Print the elements of the slices
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)

	// Append elements to a slice
	extendedSlice := append(mySlice, 17, 19) // Extends mySlice with 17 and 19

	// Print the extended slice
	fmt.Println("extendedSlice:", extendedSlice)

	// Demonstrate the zero value of a slice
	var zeroSlice []int
	fmt.Println("zeroSlice (is nil):", zeroSlice == nil)
}
