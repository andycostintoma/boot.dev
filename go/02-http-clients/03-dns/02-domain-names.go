package main

import (
	"fmt"
	"log"
	"net/url"
)

func getDomainNameFromURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %w", err)
	}
	return parsedURL.Hostname(), nil
}

func main() {

	domainName, err := getDomainNameFromURL("https://www.google.com")
	if err != nil {
		log.Fatalf("error getting domain name: %v", err)
	}
	fmt.Println(domainName)
}
