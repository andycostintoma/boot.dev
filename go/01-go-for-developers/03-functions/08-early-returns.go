package main

import (
	"errors"
	"fmt"
)

// Simple divide function demonstrating early return
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return dividend / divisor, nil
}

// Defining a struct to represent insurance status
type insuranceStatus struct {
	hasInsurance bool
	isTotaled    bool
	isDented     bool
	isBigDent    bool
}

// Function to get insurance amount using guard clauses
func getInsuranceAmount(status insuranceStatus) int {
	if !status.hasInsurance {
		return 1
	}
	if status.isTotaled {
		return 10000
	}
	if !status.isDented {
		return 0
	}
	if status.isBigDent {
		return 270
	}
	return 160
}

func main() {
	// Test the divide function
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result of division: %d\n", result)
	}

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: can't divide by zero
	}

	// Test the getInsuranceAmount function
	status1 := insuranceStatus{hasInsurance: true, isTotaled: false, isDented: true, isBigDent: true}
	status2 := insuranceStatus{hasInsurance: true, isTotaled: true, isDented: true, isBigDent: true}
	status3 := insuranceStatus{hasInsurance: true, isTotaled: false, isDented: false, isBigDent: false}
	status4 := insuranceStatus{hasInsurance: false, isTotaled: false, isDented: false, isBigDent: false}

	fmt.Printf("Insurance amount for status1: %d\n", getInsuranceAmount(status1)) // Output: 270
	fmt.Printf("Insurance amount for status2: %d\n", getInsuranceAmount(status2)) // Output: 10000
	fmt.Printf("Insurance amount for status3: %d\n", getInsuranceAmount(status3)) // Output: 0
	fmt.Printf("Insurance amount for status4: %d\n", getInsuranceAmount(status4)) // Output: 1
}
