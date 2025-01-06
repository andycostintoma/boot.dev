package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	var costsByDay []float64
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value
	}
	return costsByDay
}

func main() {
	// Test case 1
	costs1 := []cost{
		{0, 4.0},
		{1, 2.1},
		{5, 2.5},
		{1, 3.1},
	}
	result1 := getCostsByDay(costs1)
	expected1 := []float64{4.0, 5.2, 0.0, 0.0, 0.0, 2.5}
	fmt.Printf("Result 1: %v, Expected: %v\n", result1, expected1)
}
