package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawInput, err := os.ReadFile("./input")
	if err != nil {
		fmt.Printf("Failed to load input:\n%v\n", err)
		panic(1)
	}

	lines := strings.Split(string(rawInput), "\n")

	safeReports := 0
	for _, line := range lines {
		if isValidReport(line) {
			safeReports++
		}
	}

	fmt.Printf("There are %d safe reports\n", safeReports)
}

func isValidReport(report string) bool {
	numbers := []int{}
	for _, num := range strings.Fields(report) {
		converted, err := strconv.Atoi(num)
		if err != nil {
			continue
		}

		numbers = append(numbers, converted)
	}

	if len(numbers) < 2 {
		return false
	}

	isIncreasing := numbers[1] > numbers[0]
	maxDelta := 3

	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]

		if isIncreasing {
			if diff > maxDelta || diff <= 0 {
				return false
			}
		} else {
			if -diff > maxDelta || diff >= 0 {
				return false
			}
		}
	}

	return true
}
