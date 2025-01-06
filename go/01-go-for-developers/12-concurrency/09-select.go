package main

import (
	"fmt"
	"time"
)

func logMessages(chEmails, chSms chan string) {
	for {
		select {
		case email, ok := <-chEmails:
			if !ok {
				return
			}
			logEmail(email)
		case sms, ok := <-chSms:
			if !ok {
				return
			}
			logSms(sms)
		}
	}
}

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("Email:", email)
}

func main() {
	chEmails := make(chan string)
	chSms := make(chan string)

	go logMessages(chEmails, chSms)

	// Simulate sending emails and SMS messages
	go func() {
		emails := []string{"email1@example.com", "email2@example.com", "email3@example.com"}
		for _, email := range emails {
			chEmails <- email
			time.Sleep(100 * time.Millisecond) // Simulate work
		}
		close(chEmails)
	}()

	go func() {
		smsMessages := []string{"sms1", "sms2", "sms3"}
		for _, sms := range smsMessages {
			chSms <- sms
			time.Sleep(150 * time.Millisecond) // Simulate work
		}
		close(chSms)
	}()

	// Wait for logMessages to process all messages
	time.Sleep(1 * time.Second)
	fmt.Println("Main finished")
}
