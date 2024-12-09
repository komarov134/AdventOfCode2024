package main

import (
	"AdventOfCode2024/utils"
	"fmt"
	"strings"
)

func Day02Part1(lines []string) int {
	safeCount := 0
	for _, line := range lines {
		numbers := strings.Fields(line)
		if isSafe(numbers) {
			safeCount += 1
		}
	}
	return safeCount
}

func isSafeDampener(numbers []string) bool {
	if isSafe(numbers) {
		return true
	}
	for i := range numbers {
		numbersWithoutElement := []string{}
		for j, e := range numbers {
			if i != j {
				numbersWithoutElement = append(numbersWithoutElement, e)
			}
		}
		// fmt.Printf("Building %s\n", strings.Join(numbersWithoutElement, "|"))
		if isSafe(numbersWithoutElement) {
			return true
		}
	}
	return false
}

func isSafe(numbers []string) bool {
	goodChanges := 0
	for i := 0; i < len(numbers)-1; i++ {
		n1 := utils.MustAtoi(numbers[i])
		n2 := utils.MustAtoi(numbers[i+1])
		diff := n2 - n1
		if -3 <= diff && diff <= -1 {
			goodChanges -= 1
		} else if 1 <= diff && diff <= 3 {
			goodChanges += 1
		}
	}
	return utils.Abs(goodChanges) == len(numbers)-1
}

func Day02Part2(lines []string) int {
	safeCount := 0
	for _, line := range lines {
		numbers := strings.Fields(line)
		if isSafeDampener(numbers) {
			fmt.Printf("Safe %s\n", strings.Join(numbers, "|"))
			safeCount += 1
		}
	}
	return safeCount
}
