package main

import "fmt"

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) (int, error) {
	if costMultiplier <= 1 {
		return 0, fmt.Errorf("cost multiplier must be greater than 1")
	}

	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(maxCostInPennies) - actualCostInPennies

	for balance >= actualCostInPennies*costMultiplier {
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	return maxMessagesToSend, nil
}

func main() {
	maxMessagesToSend, err := getMaxMessagesToSend(1.5, 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Max messages to send: %v\n", maxMessagesToSend)
}
