package day08

import (
	// "AdventOfCode2024/utils"
	"fmt"
)

type Coord struct {
	X int
	Y int
}

func parseInput(lines []string) [][]byte {
	field := [][]byte{}
	for _, line := range lines {
		row := []byte{}
		for _, c := range line {
			row = append(row, byte(c))
		}
		field = append(field, row)
	}
	return field
}

func Part1(lines []string) int {
	field := parseInput(lines)
	for _, row := range field {
		fmt.Println(string(row))
	}
	m := make(map[byte][]Coord)
	for i, row := range field {
		for j, c := range row {
			if c != '.' {
				m[c] = append(m[c], Coord{i, j})
			}
		}
	}
	antinodeCount := 0
	for c, coords := range m {
		fmt.Println(c, "has", coords)
		for _, c1 := range coords {
			for _, c2 := range coords {
				if c1 != c2 {
					dx := c2.X - c1.X
					dy := c2.Y - c1.Y
					x, y := c2.X+dx, c2.Y+dy
					if x >= 0 && x < len(field) && y >= 0 && y < len(field[0]) && field[x][y] != '#' {
						antinodeCount++
						field[x][y] = '#'
					}
				}
			}
		}
	}
	for _, row := range field {
		fmt.Println(string(row))
	}

	return antinodeCount
}

func Part2(lines []string) int {
	field := parseInput(lines)
	for _, row := range field {
		fmt.Println(string(row))
	}
	m := make(map[byte][]Coord)
	for i, row := range field {
		for j, c := range row {
			if c != '.' {
				m[c] = append(m[c], Coord{i, j})
			}
		}
	}
	for c, coords := range m {
		fmt.Println(c, "has", coords)
		for _, c1 := range coords {
			for _, c2 := range coords {
				if c1 != c2 {
					dx := c2.X - c1.X
					dy := c2.Y - c1.Y
					for x, y := c2.X+dx, c2.Y+dy; x >= 0 && x < len(field) && y >= 0 && y < len(field[0]); x, y = x+dx, y+dy {
						field[x][y] = '#'
					}
				}
			}
		}
	}
	for _, row := range field {
		fmt.Println(string(row))
	}
	antinodeCount := 0
	for _, row := range field {
		for _, c := range row {
			if c != '.' {
				antinodeCount++
			}
		}
	}

	return antinodeCount
}
