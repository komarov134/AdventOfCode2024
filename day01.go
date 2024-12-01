package main

import (
	"fmt"
	"sort"
	"strings"
)

func Day01Part1() {
	fmt.Println("Solution:", solvePart1(ReadFileLines("day01_input.txt")))
}

func Day01Part2() {
	fmt.Println("Solution:", solvePart2(ReadFileLines("day01_input.txt")))
}

func solvePart1(lines []string) int64 {
	arr1 := []int{}
	arr2 := []int{}
	for _, line := range lines {
		numbers := strings.Fields(line)
		// fmt.Printf("Line %d: %s\n", i, strings.Join(numbers, "|"))
		arr1 = append(arr1, MustAtoi(numbers[0]))
		arr2 = append(arr2, MustAtoi(numbers[1]))
	}

	sort.Ints(arr1)
	sort.Ints(arr2)
	sum := int64(0)
	for i, e1 := range arr1 {
		e2 := arr2[i]
		sum = sum + int64(Abs(e2-e1))
	}
	return sum
}

func solvePart2(lines []string) int64 {
	arr1 := []int{}
	map2 := make(map[int]int)
	for _, line := range lines {
		numbers := strings.Fields(line)
		// fmt.Printf("Line %d: %s\n", i, strings.Join(numbers, "|"))
		arr1 = append(arr1, MustAtoi(numbers[0]))
		e2 := MustAtoi(numbers[1])
		map2[e2]++
	}
	similarityScore := 0
	for _, e1 := range arr1 {
		times, exists := map2[e1]
		if exists {
			similarityScore += times * e1
		}
	}
	return int64(similarityScore)
}
