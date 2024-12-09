package day07

import (
	"AdventOfCode2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Line struct {
	Target  int64
	Numbers []int64
}

func parseInput(lines []string) []Line {
	parsedLine := []Line{}
	for _, line := range lines {
		tn := strings.Split(line, ": ")
		target := utils.MustAtoi64(tn[0])
		numbersStr := strings.Split(tn[1], " ")
		numbers := utils.MapSlice(numbersStr, func(s string) int64 { return utils.MustAtoi64(s) })
		parsedLine = append(parsedLine, Line{target, numbers})
	}
	return parsedLine
}

func canBeSolved(line Line) bool {
	k := len(line.Numbers) - 1
	k2 := int(math.Pow(2, float64(k)))
	// n encodes operations '+'(0) and '*'(1) in its bits
	// 010 means x1 + x2 * x3 + x4
	for n := 0; n < k2; n++ {
		res := line.Numbers[0]
		nn := n // we want to change it
		for i := 0; i < k; i++ {
			isMult := nn&1 == 1
			if isMult {
				res = res * line.Numbers[i+1]
			} else {
				res = res + line.Numbers[i+1]
			}
			nn >>= 1
		}
		if res == line.Target {
			return true
		}
	}
	return false
}

func Part1(lines []string) int64 {
	parsedLines := parseInput(lines)
	for _, row := range parsedLines {
		fmt.Println(row)
	}
	result := int64(0)
	for _, line := range parsedLines {
		if canBeSolved(line) {
			result += line.Target
		}
	}
	return result
}

// [0,1,2,1] means 4 operations: ['+', '*', '||', '*']
func nextOperation(arr []int) []int {
	overflow := true
	for i := 0; i < len(arr) && overflow; i++ {
		arr[i] += 1
		overflow = arr[i] > 2
		if overflow {
			arr[i] = 0
		}
	}
	return arr
}

func canBeSolved2(line Line) bool {
	k := len(line.Numbers) - 1
	k2 := int(math.Pow(3, float64(k)))
	operation := make([]int, k) // filled with zeros
	for n := 0; n < k2; n++ {
		res := line.Numbers[0]
		for i := 0; i < k; i++ {
			switch operation[i] {
			case 0:
				res = res + line.Numbers[i+1]
			case 1:
				res = res * line.Numbers[i+1]
			case 2:
				concatNumbers := strconv.FormatInt(res, 10) + strconv.FormatInt(line.Numbers[i+1], 10)
				res = utils.MustAtoi64(concatNumbers)
			default:
				panic("operation should contain one of [0,1,2]")
			}
		}
		if res == line.Target {
			return true
		}
		operation = nextOperation(operation)
	}
	return false
}

func Part2(lines []string) int64 {
	parsedLines := parseInput(lines)
	for _, row := range parsedLines {
		fmt.Println(row)
	}
	result := int64(0)
	for _, line := range parsedLines {
		if canBeSolved2(line) {
			result += line.Target
		}
	}
	return result
}
