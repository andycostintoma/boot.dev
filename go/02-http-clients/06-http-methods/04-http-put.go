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

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id
	requestBody, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}
	request, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return User{}, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", apiKey)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return User{}, err
	}
	if response.StatusCode == 404 {
		return User{}, fmt.Errorf("user not found")
	}
	defer response.Body.Close()
	var responseData User
	if err = json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		return User{}, err
	}
	return responseData, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id
	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return User{}, err
	}
	request.Header.Set("X-API-Key", apiKey)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return User{}, err
	}
	if response.StatusCode == 404 {
		return User{}, fmt.Errorf("user not found")
	}
	defer response.Body.Close()
	var responseData User
	if err = json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		return User{}, err
	}
	return responseData, nil
}

func logUser(user User) {
	fmt.Printf("User Name: %s, Role: %s, Experience: %d, Remote: %v\n",
		user.User.Name, user.Role, user.Experience, user.Remote)
}

func main() {
	userId := "2f8282cb-e2f9-496f-b144-c0aa4ced56db"
	baseURL := "https://api.boot.dev/v1/courses_rest_api/learn-http/users"
	apiKey := generateKey()

	userData, err := getUserById(baseURL, userId, apiKey)
	if err != nil {
		fmt.Println(err)
	}
	logUser(userData)

	fmt.Printf("Updating user with id: %s\n", userData.ID)
	userData.Role = "Senior Backend Developer"
	userData.Experience = 7
	userData.Remote = true
	userData.User.Name = "Allan"

	updatedUser, err := updateUser(baseURL, userId, apiKey, userData)
	if err != nil {
		fmt.Println(err)
		return
	}
	logUser(updatedUser)
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
