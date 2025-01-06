package main

import (
	"fmt"
	"strings"
)

func countDistinctWords(messages []string) int {
	words := make(map[string]bool)
	for _, message := range messages {
		for _, word := range strings.Fields(message) {
			words[strings.ToLower(word)] = true
		}
	}

	return len(words)
}

func main() {
	messages := []string{"hello", "Hello", "HELLO", "world", "World", "WORLD"}
	distinctWords := countDistinctWords(messages)
	fmt.Println(distinctWords)
}
