package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Creating a new request to a free public API
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Setting headers on the new request
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("User-Agent", "Go-http-client/1.1")

	// Making the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()

	// Printing all response headers
	fmt.Println("\nResponse Headers:")
	for key, values := range res.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}
