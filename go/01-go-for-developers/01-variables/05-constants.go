package main

import "fmt"

// Declaring constants
const (
	pi       = 3.14159
	greeting = "Hello, World!"
	luckyNum = 7
	isGoFun  = true
)

func main() {
	// Using constants in the program
	fmt.Println("Pi:", pi)
	fmt.Println("Greeting:", greeting)
	fmt.Println("Lucky Number:", luckyNum)
	fmt.Println("Is Go programming fun?", isGoFun)

	// Trying to change a constant would result in a compilation error
	// Uncommenting the following line will cause an error
	// pi = 3.14
}
