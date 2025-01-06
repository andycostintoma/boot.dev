package main

import (
	"errors"
	"fmt"
	"log"
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

func enrichUser(userID string) (User, error) {
	user, err := getUser(userID)
	if err != nil {
		return User{}, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func main() {
	user, err := enrichUser("")
	if err != nil {
		log.Fatal("Error: ", err)
	} else {
		fmt.Println("User: ", user)
	}
}
