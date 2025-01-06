package main

import (
	"fmt"
)

func main() {
	// Define a variable to test.go
	number := 15

	if number < 10 {
		fmt.Println("The number is less than 10")
	} else if number == 10 {
		fmt.Println("The number is exactly 10")
	} else {
		fmt.Println("The number is greater than 10")
	}

}
