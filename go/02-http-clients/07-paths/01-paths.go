package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getResources(path string) []map[string]any {
	fullURL := "https://api.boot.dev" + path

	res, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	defer res.Body.Close()

	var resources []map[string]any
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resources)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return nil
	}

	return resources
}

func main() {
	projects := getResources("/v1/courses_rest_api/learn-http/projects")
	fmt.Println("Projects:")
	logResources(projects)
	fmt.Println(" --- ")

	issues := getResources("/v1/courses_rest_api/learn-http/issues")
	fmt.Println("Issues:")
	logResources(issues)
	fmt.Println(" --- ")

	users := getResources("/v1/courses_rest_api/learn-http/users")
	fmt.Println("Users:")
	logResources(users)
}

func logResources(resources []map[string]any) {
	for _, resource := range resources {
		jsonResource, err := json.Marshal(resource)
		if err != nil {
			fmt.Println("Error marshalling resource:", err)
			continue
		}
		fmt.Printf(" - %s\n", jsonResource)
	}
}
