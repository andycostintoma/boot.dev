package main

import (
	"fmt"
	"sync"
)

// Use a WaitGroup to wait for both goroutines to finish
var wg sync.WaitGroup

func goroutineA(ch1, ch2 chan int) {
	defer wg.Done() // Signal that this goroutine is done
	fmt.Println("Goroutine A: Sending to ch1")
	ch1 <- 1 // This operation will block
	fmt.Println("Goroutine A: Sent to ch1")

	fmt.Println("Goroutine A: Receiving from ch2")
	<-ch2 // This operation will block
	fmt.Println("Goroutine A: Received from ch2")
}

func goroutineB(ch1, ch2 chan int) {
	defer wg.Done() // Signal that this goroutine is done
	fmt.Println("Goroutine B: Sending to ch2")
	ch2 <- 1 // This operation will block
	fmt.Println("Goroutine B: Sent to ch2")

	fmt.Println("Goroutine B: Receiving from ch1")
	<-ch1 // This operation will block
	fmt.Println("Goroutine B: Received from ch1")
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg.Add(2) // We have two goroutines to wait for
	go goroutineA(ch1, ch2)
	go goroutineB(ch1, ch2)

	// Wait for both goroutines to finish
	wg.Wait()

	// This will not be executed due to deadlock
	fmt.Println("Main finished")
}
