package main

import "fmt"

func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Printf("Sent: %d\n", i)
	}
	close(ch) // Close the channel after sending all values
}

func receiveNumbers(ch chan int) {
	for {
		if num, ok := <-ch; ok {
			fmt.Printf("Received: %d\n", num)
		} else {
			fmt.Println("Channel closed, no more data.")
			break
		}
	}
}

func main() {
	ch := make(chan int)

	go sendNumbers(ch)

	receiveNumbers(ch)

	fmt.Println("Main finished")
}
