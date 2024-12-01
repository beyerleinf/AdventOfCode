package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	leftNumbers, rightNumbers, err := getSortedNumbers()
	if err != nil {
		panic(1)
	}

	distance := getDistanceBetweenLists(leftNumbers, rightNumbers)
	fmt.Printf("Total Distance is %d\n", distance)

	similarity := getSimilarityScore(leftNumbers, rightNumbers)
	fmt.Printf("Similarity Score is %d\n", similarity)

}

func getDistanceBetweenLists(leftNumbers, rightNumbers []int) int {
	totalDistance := 0
	for i := range leftNumbers {
		left := leftNumbers[i]
		right := rightNumbers[i]

		distance := left - right
		if distance < 0 {
			distance = -distance
		}

		totalDistance += distance
	}

	return totalDistance
}

func getSimilarityScore(leftNumbers, rightNumbers []int) int {
	similarity := 0

	for _, num := range leftNumbers {
		count := 0
		for _, right := range rightNumbers {
			if right == num {
				count++
			}
		}

		similarity += num * count
	}

	return similarity
}

func getSortedNumbers() ([]int, []int, error) {
	rawInput, err := os.ReadFile("./input")
	if err != nil {
		fmt.Printf("Failed to load input:\n%v\n", err)
		panic(1)
	}

	lines := strings.Split(string(rawInput), "\n")

	leftNumbers := []int{}
	rightNumbers := []int{}

	for _, line := range lines {
		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			fmt.Printf("line is invalid: %v\n", line)
			continue
		}

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Failed to convert left number: %s\n%v\n", parts[0], err)
			return nil, nil, err
		}

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Failed to convert right number: %s\n%v\n", parts[1], err)
			return nil, nil, err
		}

		leftNumbers = append(leftNumbers, left)
		rightNumbers = append(rightNumbers, right)
	}

	slices.Sort(leftNumbers)
	slices.Sort(rightNumbers)

	return leftNumbers, rightNumbers, nil
}
