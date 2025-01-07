package main

import "fmt"

type Permission interface {
	isPermission()
}
type ReadPermission struct{}
type WritePermission struct{}
type ExecPermission struct{}

func (ReadPermission) isPermission()  {}
func (WritePermission) isPermission() {}
func (ExecPermission) isPermission()  {}

func checkPermission(p Permission) {
	switch p.(type) {
	case ReadPermission, WritePermission, ExecPermission:
		fmt.Println("Valid Permission")
	default:
		fmt.Println("Not valid Permission")
	}
}

func main() {
	checkPermission(ReadPermission{})
	checkPermission(WritePermission{})
	checkPermission(ExecPermission{})

	//Uncommenting the next line will cause a compilation error
	//because it doesn't implement the Permission interface.
	//checkPermission(struct{}{})
}
