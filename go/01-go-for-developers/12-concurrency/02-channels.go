package main

import (
	"fmt"
	"time"
)

func sendNumber(ch chan int) {
	ch <- 42
	fmt.Println("Sent: 42")
	close(ch)
}

func receiveNumber(ch chan int) {
	num := <-ch
	fmt.Printf("Received: %d\n", num)
}

func main() {
	ch := make(chan int)

	go sendNumber(ch)
	go receiveNumber(ch)

	time.Sleep(1 * time.Second)

	fmt.Println("Finished")
}
