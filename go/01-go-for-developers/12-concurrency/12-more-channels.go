package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("1. Declared but Uninitialized Channel Is Nil:")
	var c chan string
	fmt.Println("Is channel nil?", c == nil)

	c = make(chan string)
	fmt.Println("Is channel nil?", c == nil)

	fmt.Println("\n2. Send to a Nil Channel Blocks Forever:")
	go func() {
		var nilChan chan string
		nilChan <- "let's get started"
		fmt.Println("This will never print")
	}()
	time.Sleep(1 * time.Second)

	fmt.Println("\n3. Receive from a Nil Channel Blocks Forever:")
	go func() {
		var nilChan chan string
		_ = <-nilChan
		fmt.Println("This will also never print")
	}()
	time.Sleep(1 * time.Second)

	fmt.Println("\n4. Send to a Closed Channel Panics:")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		closedChan := make(chan int, 100)
		close(closedChan)
		closedChan <- 1
		fmt.Println("This will never print")
	}()

	fmt.Println("\n5. Receive from a Closed Channel Returns the Zero Value Immediately:")
	closedChan := make(chan int, 100)
	close(closedChan)
	val := <-closedChan
	fmt.Println(val)
	fmt.Println("This will print after zero value")
}
