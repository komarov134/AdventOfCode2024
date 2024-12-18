package day14

import (
	"AdventOfCode2024/utils"
	"fmt"
	"strings"
	// "time"
)

type Robot struct {
	X  int
	Y  int
	Dx int
	Dy int
}

func parseInput(lines []string) []Robot {
	robots := []Robot{}
	for _, line := range lines {
		pv := strings.Split(line, " ")
		p := strings.Split(pv[0][2:], ",")
		v := strings.Split(pv[1][2:], ",")
		robots = append(robots, Robot{X: utils.MustAtoi(p[1]), Y: utils.MustAtoi(p[0]), Dx: utils.MustAtoi(v[1]), Dy: utils.MustAtoi(v[0])})
	}
	return robots
}

const Wide = 101
const Tall = 103
const Seconds = 100

func Part1(lines []string) int {
	robots := parseInput(lines)
	field := make([][]int, Tall)
	for i := range field {
		field[i] = make([]int, Wide)
	}
	for _, r := range robots {
		fmt.Println(r)
		field[r.X][r.Y]++
	}
	for _, row := range field {
		fmt.Println(row)
	}
	for i := range field {
		field[i] = make([]int, Wide)
	}

	qs := [4]int{}
	for _, r := range robots {
		newX := (r.X + r.Dx*Seconds + 1000*Tall) % Tall
		newY := (r.Y + r.Dy*Seconds + 1000*Wide) % Wide
		field[newX][newY]++
		if newX < Tall/2 && newY < Wide/2 {
			qs[0]++
		} else if newX < Tall/2 && newY > Wide/2 {
			qs[1]++
		} else if newX > Tall/2 && newY < Wide/2 {
			qs[2]++
		} else if newX > Tall/2 && newY > Wide/2 {
			qs[3]++
		}
	}
	fmt.Println("=====================================")
	for _, row := range field {
		fmt.Println(row)
	}
	return qs[0] * qs[1] * qs[2] * qs[3]
}

func printField(field [][]int) {
	for _, row := range field {
		for _, c := range row {
			if c == 0 {
				fmt.Print("  ")
			} else {
				fmt.Print(" *")
			}
		}
		fmt.Println()
	}
}

// exists 10 rows such that
func looksLikeTree(field [][]int) bool {
	count := 0
	for _, row := range field {
		maxNonEmptyLength := 0
		curNonEmptyLength := 0
		for _, c := range row {
			if c == 0 {
				curNonEmptyLength = 0
			} else {
				curNonEmptyLength++
			}
			maxNonEmptyLength = max(maxNonEmptyLength, curNonEmptyLength)
		}
		if maxNonEmptyLength > 5 {
			count++
		}
	}
	return count > 10
}

func updateField(seconds int, robots []Robot, field [][]int) {
	for i := range field {
		field[i] = make([]int, Wide)
	}
	for _, r := range robots {
		newX := (r.X + r.Dx*seconds + seconds*Tall) % Tall
		newY := (r.Y + r.Dy*seconds + seconds*Wide) % Wide
		field[newX][newY]++
	}
}

func Part2(lines []string) int {
	robots := parseInput(lines)
	field := make([][]int, Tall)

	seconds := 0

	printCount := 10

	for printCount > 0 {
		seconds++
		updateField(seconds, robots, field)
		if looksLikeTree(field) {
			fmt.Println("=====================================>", seconds)
			printField(field)
			printCount--
		}
	}

	return 0
}
