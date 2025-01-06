package main

import (
	"fmt"
	"time"
)

func sendInts(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
	close(ch)
}

func sendStrings(ch chan string) {
	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("str%d", i)
		time.Sleep(150 * time.Millisecond) // Simulate work
	}
	close(ch)
}

func main() {
	intCh := make(chan int)
	strCh := make(chan string)

	go sendInts(intCh)
	go sendStrings(strCh)

	for intCh != nil || strCh != nil {
		select {
		case i, ok := <-intCh:
			if !ok {
				fmt.Println("Int channel closed")
				intCh = nil
			} else {
				fmt.Printf("Received int: %d\n", i)
			}
		case s, ok := <-strCh:
			if !ok {
				fmt.Println("String channel closed")
				strCh = nil
			} else {
				fmt.Printf("Received string: %s\n", s)
			}
		}
	}

	fmt.Println("Main finished")
}
