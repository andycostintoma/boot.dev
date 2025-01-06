package main

import "fmt"

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(c *customer, t transaction) error {
	switch t.transactionType {
	case transactionDeposit:
		c.balance += t.amount
	case transactionWithdrawal:
		if t.amount > c.balance {
			return fmt.Errorf("insufficient funds")
		}
		c.balance -= t.amount
	default:
		return fmt.Errorf("unknown transaction type")
	}
	return nil
}
