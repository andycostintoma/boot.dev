package main

import (
	"fmt"
)

func getMonthlyPrice(tier string) int {
	switch tier {
	case "basic":
		return 10000
	case "premium":
		return 15000
	case "enterprise":
		return 50000
	default:
		return 0
	}
}

// Wrapping test function in a way that can be called manually
func runTests() {
	type testCase struct {
		tier     string
		expected int
	}
	tests := []testCase{
		{"basic", 10000},
		{"premium", 15000},
		{"enterprise", 50000},
	}

	tests = append(tests, []testCase{
		{"invalid", 0},
		{"", 0},
	}...)

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := getMonthlyPrice(test.tier)
		if output != test.expected {
			failCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, test.tier, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, test.tier, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func main() {
	runTests()
}
