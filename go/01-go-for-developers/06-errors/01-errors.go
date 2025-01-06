package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	cost1, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}
	cost2, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}
	return cost1 + cost2, nil
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

func main() {
	cost, err := sendSMSToCouple("Hello!", "Hi!")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Total cost: %v\n", cost)
}
