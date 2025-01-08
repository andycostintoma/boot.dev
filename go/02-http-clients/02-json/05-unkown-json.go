package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func getResources(url string) ([]map[string]any, error) {
	var resources []map[string]any

	res, err := http.Get(url)
	if err != nil {
		return resources, err
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&resources); err != nil {
		return resources, err
	}

	return resources, nil

}

func logResources(resources []map[string]any) {
	var formattedStrings []string

	for _, resource := range resources {
		for k, v := range resource {
			formattedStrings = append(formattedStrings, fmt.Sprintf("Key: %s - Value: %v", k, v))
		}
	}

	sort.Strings(formattedStrings)

	for _, str := range formattedStrings {
		fmt.Println(str)
	}
}

const baseUrl = "https://api.boot.dev"

func main() {
	issues, err := getResources(baseUrl + "/v1/courses_rest_api/learn-http/issues?limit=1")
	if err != nil {
		fmt.Println("Error getting issues:", err)
		return
	}
	fmt.Println("Issue:")
	logResources(issues)
	fmt.Println("---")

	projects, err := getResources(baseUrl + "/v1/courses_rest_api/learn-http/projects?limit=1")
	if err != nil {
		fmt.Println("Error getting projects:", err)
		return
	}
	fmt.Println("Project:")
	logResources(projects)
	fmt.Println("---")

	users, err := getResources(baseUrl + "/v1/courses_rest_api/learn-http/users?limit=1")
	if err != nil {
		fmt.Println("Error getting users:", err)
		return
	}
	fmt.Println("User:")
	logResources(users)
}
