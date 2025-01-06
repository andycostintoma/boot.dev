package main

import "unicode"

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}

	atLeastOneUpper := false
	atLeastOneDigit := false

	for _, c := range password {
		if unicode.IsUpper(c) {
			atLeastOneUpper = true
		}
		if unicode.IsDigit(c) {
			atLeastOneDigit = true
		}
		if atLeastOneUpper && atLeastOneDigit {
			return true
		}
	}

	return false
}
