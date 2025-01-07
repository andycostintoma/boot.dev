package main

import "fmt"

type Permission string

const (
	ReadPermission  Permission = "READ"
	WritePermission Permission = "WRITE"
	ExecPermission  Permission = "EXEC"
)

func checkPermission(p Permission) {
	switch p {
	case ReadPermission, WritePermission, ExecPermission:
		fmt.Println("Valid Permission:", p)
	default:
		fmt.Println("Not valid Permission:", p)
	}
}

func main() {
	checkPermission(ReadPermission)
	checkPermission("INVALID") // Invalid but allowed at compile time
}
