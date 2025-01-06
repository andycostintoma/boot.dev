package main

import (
	"fmt"
)

func getCreator(os string) string {
	var creator string
	switch os {
	case "linux":
		creator = "Linus Torvalds"
	case "windows":
		creator = "Bill Gates"
	case "Mac OS": // fallthrough example
		fallthrough
	case "Mac OS X": // fallthrough example
		fallthrough
	case "mac":
		creator = "A Steve"
	default:
		creator = "Unknown"
	}
	return creator
}

func main() {
	fmt.Println(getCreator("linux"))   // Output: Linus Torvalds
	fmt.Println(getCreator("windows")) // Output: Bill Gates
	fmt.Println(getCreator("mac"))     // Output: A Steve
	fmt.Println(getCreator("Mac OS"))  // Output: A Steve
	fmt.Println(getCreator("unknown")) // Output: Unknown
}
