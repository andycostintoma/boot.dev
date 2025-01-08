package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Role       string `json:"role"`
	ID         string `json:"id"`
	Experience int    `json:"experience"`
	Remote     bool   `json:"remote"`
	User       struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Age      int    `json:"age"`
	} `json:"user"`
}

func logUsers(users []User) {
	for _, user := range users {
		fmt.Printf("User Name: %s, Role: %s, Experience: %d, Remote: %v\n", user.User.Name, user.Role, user.Experience, user.Remote)
	}
}

func getUsers(url string) ([]User, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var users []User
	if err = json.NewDecoder(res.Body).Decode(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func main() {
	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	users, err := getUsers(url)
	if err != nil {
		log.Fatal("Error getting users:", err)
	}
	logUsers(users)
}
