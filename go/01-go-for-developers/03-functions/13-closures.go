package main

import (
	"fmt"
)

func concatter() func(string) string {
	doc := "" // 'doc' is captured by the closure
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func main() {
	// Create a new concatter instance
	harryPotterAggregator := concatter()

	// Add words to the document.
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	// Print the concatenated document.
	fmt.Println(harryPotterAggregator("Drive"))
	// Output: Mr. and Mrs. Dursley of number four, Privet Drive
}
