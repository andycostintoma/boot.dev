package main

import (
	"fmt"
)

func main() {
	x := 5
	increment(x)

	fmt.Println(x)
	// Still prints 5,
	// because the increment function
	// received a copy of x
}

func increment(x int) {
	x++
}
