package day10

import (
	"AdventOfCode2024/utils"
	"fmt"
)

func parseInput(lines []string) [][]uint8 {
	field := [][]uint8{}
	for _, line := range lines {
		row := utils.MapSlice([]uint8(line), func(n uint8) uint8 {
			return n - uint8('0')
		})
		field = append(field, row)
	}
	return field
}

const Visited = 99

func resetVisited(field [][]uint8) {
	for i, row := range field {
		for j, cell := range row {
			if cell == Visited {
				field[i][j] = 9
			}
		}
	}
}

func countTrails(i int, j int, expectedHeight uint8, field [][]uint8) int {
	if i < 0 || i >= len(field) || j < 0 || j >= len(field[0]) || field[i][j] != expectedHeight {
		return 0
	}
	if field[i][j] == 9 {
		// field[i][j] = Visited // comment to solve Part2
		return 1
	}
	res := 0
	res += countTrails(i, j-1, expectedHeight+1, field)
	res += countTrails(i, j+1, expectedHeight+1, field)
	res += countTrails(i-1, j, expectedHeight+1, field)
	res += countTrails(i+1, j, expectedHeight+1, field)
	return res
}

func Part1(lines []string) int {
	field := parseInput(lines)
	for _, row := range field {
		fmt.Println(row)
	}
	totalTrails := 0
	for i, row := range field {
		for j, cell := range row {
			if cell == 0 {
				totalTrails += countTrails(i, j, 0, field)
				// resetVisited(field) // comment to solve Part2
				fmt.Println("totalTrails", totalTrails)
			}
		}
	}

	return totalTrails
}
