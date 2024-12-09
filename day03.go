package main

import (
	"AdventOfCode2024/utils"
	"fmt"
	"regexp"
)

func Day03Part1(lines []string) int {
	result := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, 100000)
		for _, match := range matches {
			num1 := utils.MustAtoi(match[1])
			num2 := utils.MustAtoi(match[2])
			result += num1 * num2
		}
	}
	return result
}

func Day03Part2(lines []string) int {
	result := 0
	stateEnabled := true
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			fmt.Println("found:", match[0])
			if match[0] == "do()" {
				stateEnabled = true
			} else if match[0] == "don't()" {
				stateEnabled = false
			} else {
				if stateEnabled {
					num1 := utils.MustAtoi(match[1])
					num2 := utils.MustAtoi(match[2])
					result += num1 * num2
				}
			}
		}
	}
	return result
}
