package main

import (
	"fmt"
)

func bootup() {
	defer fmt.Println("TEXTIO BOOTUP DONE")

	if !connectToDB() {
		return
	}

	if !connectToPaymentProvider() {
		return
	}

	fmt.Println("All systems ready!")
}

// Variables to control connection success
var shouldConnectToDB = true
var shouldConnectToPaymentProvider = true

// connectToDB attempts to connect to the database
func connectToDB() bool {
	fmt.Println("Connecting to database...")
	if shouldConnectToDB {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}

// connectToPaymentProvider attempts to connect to the payment provider
func connectToPaymentProvider() bool {
	fmt.Println("Connecting to payment provider...")
	if shouldConnectToPaymentProvider {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}

// test runs the bootup sequence with given connection success parameters
func test(dbSuccess, paymentSuccess bool) {
	shouldConnectToDB = dbSuccess
	shouldConnectToPaymentProvider = paymentSuccess
	bootup()
	fmt.Println("====================================")
}

// main function to run multiple test scenarios
func main() {
	test(true, true)   // Both connections succeed
	test(false, true)  // Database connection fails
	test(true, false)  // Payment provider connection fails
	test(false, false) // Both connections fail
}
