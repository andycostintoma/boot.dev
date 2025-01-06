package main

import "fmt"

func bulkSend(numMessages int) float64 {
	cost := 0.0
	for i := 0; i < numMessages; i++ {
		cost += 1.00 + float64(i)*0.01
	}
	return cost
}

func main() {

	cost := bulkSend(100)
	fmt.Printf("Total cost: %.2f\n", cost)
}
