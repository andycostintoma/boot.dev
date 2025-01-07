package main

import "fmt"

func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func getLast[T any](s []T) T {
	if len(s) == 0 {
		var zeroValue T
		return zeroValue
	}
	return s[len(s)-1]
}

func main() {

	s1, s2 := splitAnySlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(s1, s2)
	fmt.Println(getLast(s1))

	s3, s4 := splitAnySlice([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"})
	fmt.Println(s3, s4)
	fmt.Println(getLast(s3))

}
