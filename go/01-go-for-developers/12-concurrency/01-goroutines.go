package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Number: %d\n", i)
		time.Sleep(700 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "ABCDE" {
		fmt.Printf("Letter: %c\n", letter)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	go printNumbers()

	go printLetters()

	time.Sleep(4 * time.Second)

	fmt.Println("Finished")
}
