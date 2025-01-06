package main

import (
	"fmt"
	"time"
)

func writeChannel(ch chan<- time.Time) {
	ticker := time.Tick(1 * time.Second)
	for t := range ticker {
		select {
		case ch <- t:
		default:
			// If writing would block, do nothing (we'll use a default case)
			fmt.Println("Channel full, skipping timestamp:", t)
		}
	}
}

func readChannel(ch <-chan time.Time) {
	for {
		select {
		case t := <-ch:
			fmt.Println("Received timestamp:", t)
		default:
			fmt.Println("No new timestamp, doing other work...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ch := make(chan time.Time, 2) // Buffered channel with capacity 2

	go writeChannel(ch) // Start writing timestamps
	go readChannel(ch)  // Start reading timestamps

	// Run for a period to demonstrate functionality
	time.Sleep(10 * time.Second)
	fmt.Println("Main finished")
}
