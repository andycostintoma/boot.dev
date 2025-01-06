package main

import (
	"fmt"
)

func main() {
	// Declare an array of 10 integers
	var myInts [10]int
	// Initialize the array with values
	for i := 0; i < len(myInts); i++ {
		myInts[i] = i * i
	}

	// Print the elements of the array
	fmt.Println("Array myInts:", myInts)

	// Declare and initialize an array literal of prime numbers
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Print the elements of the prime numbers array
	fmt.Println("Array primes:", primes)
}
