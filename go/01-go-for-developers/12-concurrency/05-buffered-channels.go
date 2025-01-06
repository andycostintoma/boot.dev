package main

import "fmt"

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))
	for _, email := range emails {
		ch <- email
	}
	return ch
}

func main() {
	ch := addEmailsToQueue([]string{"a", "b", "c"})
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
