package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID   string
	Name string
}

func getUser(userID string) (User, error) {
	if userID == "" {
		return User{}, errors.New("user not found")
	}
	return User{ID: userID, Name: "John Doe"}, nil
}

func enrichUser(userID string) User {
	user, err := getUser(userID)
	if err != nil {
		panic(err)
	}
	return user
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	user := enrichUser("")
	fmt.Println("User:", user)
}
