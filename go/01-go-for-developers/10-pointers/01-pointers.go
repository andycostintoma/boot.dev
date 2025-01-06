package main

import (
	"fmt"
)

func main() {
	x := 42
	px := &x // px is a pointer to the variable x

	fmt.Println("Value of x:", x)
	fmt.Println("Value of using pointer px:", *px)
	fmt.Println("Address of x:", px)

	*px = 21 // Change the value at the memory address px points to
	fmt.Println("Value of x:", x)
	fmt.Println("Value of using pointer px:", *px)
	fmt.Println("Address of x:", px)

	// nil pointers
	var p *int
	fmt.Println("Value of p:", p)
}
