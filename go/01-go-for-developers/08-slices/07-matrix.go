package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			matrix[i][j] = i * j
		}
	}
	return matrix
}

func main() {
	matrix := createMatrix(5, 5)
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
