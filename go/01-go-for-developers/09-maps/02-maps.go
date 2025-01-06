package main

import (
	"errors"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	users := make(map[string]user)
	for i, name := range names {
		users[name] = user{
			name:        name,
			phoneNumber: phoneNumbers[i],
		}
	}
	return users, nil
}

type user struct {
	name        string
	phoneNumber int
}
