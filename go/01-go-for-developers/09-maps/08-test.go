package main

import (
	"fmt"
)

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	// Create a set of the user's direct friends
	directFriends := make(map[string]bool)
	for _, friend := range friendships[username] {
		directFriends[friend] = true
	}

	// Create a map for mutual friends to ensure uniqueness
	suggestedFriends := make(map[string]bool)
	for _, friend := range friendships[username] {
		for _, friendOfFriend := range friendships[friend] {
			if friendOfFriend != username && !directFriends[friendOfFriend] {
				suggestedFriends[friendOfFriend] = true
			}
		}
	}

	// Convert the map keys to a slice
	suggestions := make([]string, 0, len(suggestedFriends))
	for mutualFriend := range suggestedFriends {
		suggestions = append(suggestions, mutualFriend)
	}

	return suggestions
}

func main() {
	friendships := map[string][]string{
		"Alice":   {"Bob", "Charlie"},
		"Bob":     {"Alice", "Charlie", "David"},
		"Charlie": {"Alice", "Bob", "David", "Eve"},
		"David":   {"Bob", "Charlie"},
		"Eve":     {"Charlie"},
	}

	suggestedFriends := findSuggestedFriends("Alice", friendships)
	fmt.Printf("Suggested friends for Alice: %v\n", suggestedFriends)
}
