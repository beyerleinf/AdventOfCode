package main

import (
	"fmt"
	"os"
	"slices"
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

	leftNums := []int{}
	rightNums := []int{}

	for _, line := range lines {
		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			fmt.Printf("line is invalid: %v\n", line)
			continue
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Failed to convert left number: %s\n%v\n", parts[0], err)
			panic(1)
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Failed to convert right number: %s\n%v\n", parts[1], err)
			panic(1)
		}

		leftNums = append(leftNums, left)
		rightNums = append(rightNums, right)
	}

	slices.Sort(leftNums)
	slices.Sort(rightNums)

	totalDistance := 0
	for i := range leftNums {
		left := leftNums[i]
		right := rightNums[i]

		distance := left - right
		if distance < 0 {
			distance = -distance
		}

		totalDistance += distance
	}

	fmt.Printf("Total distance is %d\n", totalDistance)
}
