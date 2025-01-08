package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users", nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Adding custom headers
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("User-Agent", "my-custom-agent")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	// Reading the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	fmt.Println("Response Body:", string(body))
}
