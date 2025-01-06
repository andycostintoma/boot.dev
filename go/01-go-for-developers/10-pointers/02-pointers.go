package main

import (
	"strings"
)

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	*message = strings.ReplaceAll(*message, "fubb", "****")
	*message = strings.ReplaceAll(*message, "shiz", "****")
	*message = strings.ReplaceAll(*message, "witch", "*****")
}

func main() {
	message := "I hate fubb and shiz"
	removeProfanity(&message)
	println(message)
}
