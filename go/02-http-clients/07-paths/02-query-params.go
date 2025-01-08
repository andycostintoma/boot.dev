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
		fmt.Printf("User Name: %s, Role: %s, Experience: %d, Remote: %v\n",
			user.User.Name, user.Role, user.Experience, user.Remote)
	}
}

func getUsers(url string) ([]User, error) {
	fullURL := url + "?sort=experience"
	res, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var users []User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func main() {
	baseURL := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"

	users, err := getUsers(baseURL)
	if err != nil {
		log.Fatal(err)
	}
	logUsers(users)

}
