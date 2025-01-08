package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getIssueData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	return data, nil
}

func prettify(data string) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(data), "", "  ")
	if err != nil {
		return "", fmt.Errorf("error indenting JSON: %w", err)
	}
	return prettyJSON.String(), nil
}

const issueURL = "https://api.boot.dev/v1/courses_rest_api/learn-http/issues"

func main() {
	issues, err := getIssueData(issueURL)
	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}
	prettyData, err := prettify(string(issues))
	if err != nil {
		log.Fatalf("error prettifying data: %v", err)
	}
	fmt.Println(prettyData)
}
