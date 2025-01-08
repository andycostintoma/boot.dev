package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getIssueData() ([]byte, error) {
	res, err := http.Get("https://api.boot.dev/v1/courses_rest_api/learn-http/issues")

	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	return data, err
}

func main() {
	issues, err := getIssueData()

	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}

	stringifiedData := string(issues)
	fmt.Println(stringifiedData)
}
