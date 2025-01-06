package main

import (
	"fmt"
)

// increment by value
func incrementByValue(x int) {
	x++
	fmt.Println("Inside incrementByValue:", x) // This will print 6
}

// increment by reference
func incrementByReference(x *int) {
	*x++
	fmt.Println("Inside incrementByReference:", *x) // This will print 6
}

type Analytics struct {
	MessagesTotal int
}

func main() {
	// Passing by value
	x := 5
	incrementByValue(x)
	fmt.Println("After incrementByValue:", x) // This will print 5

	// Passing by reference
	incrementByReference(&x)
	fmt.Println("After incrementByReference:", x) // This will print 6

	// Using pointers with structs
	analytics := &Analytics{MessagesTotal: 10}
	fmt.Println("\nOriginal MessagesTotal:", analytics.MessagesTotal) // This will print 10

	// Increment MessagesTotal using pointer
	analytics.MessagesTotal++
	fmt.Println("Updated MessagesTotal via pointer:", analytics.MessagesTotal) // This will print 11

	// As shorthand for:
	(*analytics).MessagesTotal++
	fmt.Println("Updated MessagesTotal using dereferencing:", (*analytics).MessagesTotal) // This will print 12
}
