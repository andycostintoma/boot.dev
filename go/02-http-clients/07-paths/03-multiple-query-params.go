package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Issue struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Estimate int    `json:"estimate"`
	Status   string `json:"status"`
}

func getIssues(url string) []Issue {
	res, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	var issues []Issue
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&issues)
	if err != nil {
		return nil
	}

	return issues
}

func logIssues(issues []Issue) string {
	log := ""
	for _, issue := range issues {
		log += fmt.Sprintf("- Issue: %s - Estimate: %d\n", issue.Title, issue.Estimate)
	}
	return log
}

func fetchTasks(baseURL, availability string) []Issue {
	amountOfIssues := map[string]int{
		"Low":    1,
		"Medium": 3,
		"High":   5,
	}

	fullURL := fmt.Sprintf("%s?sort=estimate&limit=%d", baseURL, amountOfIssues[availability])
	return getIssues(fullURL)
}

func main() {
	issues := fetchTasks("https://api.boot.dev/v1/courses_rest_api/learn-http/issues", "Low")
	logIssues(issues)
}
