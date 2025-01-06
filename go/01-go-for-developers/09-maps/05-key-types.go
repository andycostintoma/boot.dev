package main

import (
	"fmt"
)

type Key struct {
	Path    string
	Country string
}

func add(m map[string]map[string]int, path, country string) {
	mm, ok := m[path]
	if !ok {
		mm = make(map[string]int)
		m[path] = mm
	}
	mm[country]++
}

func main() {
	// Example using map of maps (complex nested maps)
	hits := make(map[string]map[string]int)

	// Using the add function
	add(hits, "/doc/", "au")
	add(hits, "/doc/", "us")
	add(hits, "/doc/", "au")

	// Retrieving a value from the nested map
	fmt.Printf("Number of times Australians visited /doc/: %d\n", hits["/doc/"]["au"])

	// Example using a map with struct keys
	hitsWithStructKey := make(map[Key]int)

	// Incrementing hits using struct keys
	hitsWithStructKey[Key{"/", "vn"}]++
	hitsWithStructKey[Key{"/", "vn"}]++
	hitsWithStructKey[Key{"/ref/spec", "ch"}]++

	// Retrieving and printing values using struct keys
	fmt.Printf("Number of times Vietnamese visited home page: %d\n", hitsWithStructKey[Key{"/", "vn"}])
	fmt.Printf("Number of times Swiss people read the spec: %d\n", hitsWithStructKey[Key{"/ref/spec", "ch"}])
}
