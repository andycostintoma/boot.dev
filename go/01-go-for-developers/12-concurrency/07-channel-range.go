package main

import "fmt"

func concurrentFib(n int) []int {
	ch := make(chan int)
	nums := make([]int, 0, n)

	go fibonacci(n, ch)

	for i := range ch {
		nums = append(nums, i)
	}

	return nums
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	n := 10
	result := concurrentFib(n)
	fmt.Println(result)
}
