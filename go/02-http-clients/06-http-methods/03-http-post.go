package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
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

func getUsers(url, apiKey string) ([]User, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-API-Key", apiKey)

	res, err := http.DefaultClient.Do(req)
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

func logUsers(users []User) {
	for _, user := range users {
		fmt.Printf("User Name: %s, Role: %s, Experience: %d, Remote: %v\n", user.User.Name, user.Role, user.Experience, user.Remote)
	}
}

func createUser(url, apiKey string, data User) (User, error) {
	reqBody, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return User{}, err
	}

	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	var response User
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return User{}, err
	}

	return response, nil
}

func main() {
	userToCreate := User{
		Role:       "Junior Developer",
		Experience: 2,
		Remote:     true,
		User: struct {
			Name     string `json:"name"`
			Location string `json:"location"`
			Age      int    `json:"age"`
		}{
			Name:     "Dan",
			Location: "NOR",
			Age:      29,
		},
	}

	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	apiKey := generateKey()

	fmt.Println("Retrieving user data...")
	userDataFirst, err := getUsers(url, apiKey)
	if err != nil {
		fmt.Println("Error retrieving users:", err)
		return
	}
	logUsers(userDataFirst)
	fmt.Println("---")

	fmt.Println("Creating new character...")
	creationResponse, err := createUser(url, apiKey, userToCreate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	jsonData, _ := json.Marshal(creationResponse)
	fmt.Printf("Creation response body: %s\n", string(jsonData))
	fmt.Println("---")

	fmt.Println("Retrieving user data...")
	userDataSecond, err := getUsers(url, apiKey)
	if err != nil {
		fmt.Println("Error retrieving users:", err)
		return
	}
	logUsers(userDataSecond)
	fmt.Println("---")
}

func generateKey() string {
	const characters = "ABCDEF0123456789"
	result := ""
	rand.New(rand.NewSource(0))
	for i := 0; i < 16; i++ {
		result += string(characters[rand.Intn(len(characters))])
	}
	return result
}
