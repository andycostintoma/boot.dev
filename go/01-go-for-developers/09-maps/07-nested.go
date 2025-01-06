package main

import (
	"fmt"
)

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		if len(name) == 0 {
			continue
		}
		firstChar := []rune(name)[0]
		if _, ok := counts[firstChar]; !ok {
			counts[firstChar] = make(map[string]int)
		}
		counts[firstChar][name]++
	}
	return counts
}

func main() {
	names := []string{"billy", "billy", "bob", "joe"}
	counts := getNameCounts(names)

	for firstChar, namesMap := range counts {
		fmt.Printf("%c: {\n", firstChar)
		for name, count := range namesMap {
			fmt.Printf("    %s: %d\n", name, count)
		}
		fmt.Println("}")
	}
}
