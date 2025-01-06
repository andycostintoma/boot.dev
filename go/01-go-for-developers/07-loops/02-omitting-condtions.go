package main

import "fmt"

func maxMessages(thresh int) int {
	cost := 0
	for i := 0; ; i++ {
		cost += 100 + i
		if cost > thresh {
			return i
		}
	}
}

func main() {

	m := maxMessages(1000)
	fmt.Printf("Max messages: %d\n", m)
}
