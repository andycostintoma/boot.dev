package main

import "fmt"

type Permission int

const (
	ReadPermission Permission = iota
	WritePermission
	ExecPermission
)

func checkPermission(p Permission) {
	switch p {
	case ReadPermission, WritePermission, ExecPermission:
		fmt.Println("Valid Permission")
	default:
		fmt.Println("Not valid Permission")
	}
}

func main() {
	checkPermission(ReadPermission)
	checkPermission(100) // Invalid but allowed at compile time
}
