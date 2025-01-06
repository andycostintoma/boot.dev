package main

import (
	"fmt"
)

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func getExpenseReport(e expense) (string, float64) {
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0
	}
}

func main() {
	e := email{isSubscribed: false, body: "Hello, World!", toAddress: "example@example.com"}
	s := sms{isSubscribed: true, body: "Hello!", toPhoneNumber: "+1234567890"}
	i := invalid{}

	eReport, eCost := getExpenseReport(e)
	sReport, sCost := getExpenseReport(s)
	iReport, iCost := getExpenseReport(i)

	fmt.Printf("Email Report: %s, Cost: %.2f\n", eReport, eCost)
	fmt.Printf("SMS Report: %s, Cost: %.2f\n", sReport, sCost)
	fmt.Printf("Invalid Report: %s, Cost: %.2f\n", iReport, iCost)
}
