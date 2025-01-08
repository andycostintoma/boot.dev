package main

import (
	"fmt"
	"net/url"
)

type ParsedURL struct {
	protocol string
	username string
	password string
	hostname string
	port     string
	path     string
	query    string
	fragment string
}

func newParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}

	password := ""
	if pw, hasPassword := parsedUrl.User.Password(); hasPassword {
		password = pw
	}

	return ParsedURL{
		protocol: parsedUrl.Scheme,
		username: parsedUrl.User.Username(),
		password: password,
		hostname: parsedUrl.Hostname(),
		port:     parsedUrl.Port(),
		path:     parsedUrl.Path,
		query:    parsedUrl.RawQuery,
		fragment: parsedUrl.Fragment,
	}
}

func main() {
	myUrl := "https://waynelagner:pwn3d@jello.app:8080/boards?sort=createdAt#id"
	parsed := newParsedURL(myUrl)

	fmt.Printf("Protocol: %s\n", parsed.protocol)
	fmt.Printf("Username: %s\n", parsed.username)
	fmt.Printf("Password: %s\n", parsed.password)
	fmt.Printf("Hostname: %s\n", parsed.hostname)
	fmt.Printf("Port:     %s\n", parsed.port)
	fmt.Printf("Path: 	 %s\n", parsed.path)
	fmt.Printf("Query:    %s\n", parsed.query)
	fmt.Printf("Fragment: %s\n", parsed.fragment)
}
