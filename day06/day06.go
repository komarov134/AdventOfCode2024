package day06

import (
	"fmt"
)

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

var DxDy = [4][2]int{
	{0, -1}, // left
	{-1, 0}, // top
	{0, +1}, // right
	{+1, 0}, // down
}

func isInField(x int, y int, field [][]byte) bool {
	return 0 <= x && x < len(field) && 0 <= y && y < len(field[0])
}

func turnRight(direction int) int {
	return (direction + 1) % 4
}

func nextStep(curX int, curY int, direction int) (int, int) {
	return curX + DxDy[direction][0], curY + DxDy[direction][1]
}

func findStart(field [][]byte) (int, int) {
	for i, row := range field {
		for j, c := range row {
			if c == '^' {
				return i, j
			}
		}
	}
	panic("start position not found")
}

func Part1(lines []string) int {
	field := parseInput(lines)
	for _, row := range field {
		fmt.Println(string(row))
	}
	curX, curY := findStart(field)
	// stop if next is out of field
	// current is always in field
	// current is always marked as visited
	direction := 1 // left, TOP, right, down
	nextX, nextY := nextStep(curX, curY, direction)
	field[curX][curY] = 'X' // set current as visited
	for isInField(curX, curY, field) && isInField(nextX, nextY, field) {
		if field[nextX][nextY] == '#' {
			direction = turnRight(direction)
			nextX, nextY = nextStep(curX, curY, direction)
		} else {
			// it's '.' or 'X' so no obstacle
			// make step, mark as visited, keep direction, update next
			curX, curY = nextX, nextY
			field[curX][curY] = 'X'
			nextX, nextY = nextStep(curX, curY, direction)
		}
	}
	fmt.Println("> guard path is ready:")
	for _, row := range field {
		fmt.Println(string(row))
	}

	return countVisited(field)
}

func countVisited(field [][]byte) int {
	visitedCount := 0
	for _, row := range field {
		for _, c := range row {
			if c == 'X' {
				visitedCount++
			}
		}
	}
	return visitedCount
}

func isStuckInLoop(field [][]byte) bool {
	curX, curY := findStart(field)
	direction := 1 // left, TOP, right, down
	nextX, nextY := nextStep(curX, curY, direction)
	field[curX][curY] = 'X' // set current as visited
	stuckInLoop := false
	prevCountVisited := 0
	prevSteps := 0
	steps := 0
	for isInField(curX, curY, field) && isInField(nextX, nextY, field) && !stuckInLoop {
		if field[nextX][nextY] == '#' {
			direction = turnRight(direction)
			nextX, nextY = nextStep(curX, curY, direction)
		} else {
			// it's '.' or 'X' so no obstacle
			// make step, mark as visited, keep direction, update next
			curX, curY = nextX, nextY
			field[curX][curY] = 'X'
			nextX, nextY = nextStep(curX, curY, direction)
		}
		// simple way to check if it's stuck: it makes no progress.
		steps++
		if steps > prevSteps+300 {
			prevSteps = steps
			curCountVisited := countVisited(field)
			if curCountVisited == prevCountVisited {
				stuckInLoop = true
			}
			prevCountVisited = curCountVisited
		}
	}
	return stuckInLoop
}

func copyField(field [][]byte) [][]byte {
	copy := make([][]byte, len(field))
	for i := range field {
		copy[i] = append([]byte{}, field[i]...)
	}
	return copy
}

func Part2(lines []string) int {
	field := parseInput(lines)
	newObstructionCount := 0
	for i, row := range field {
		for j, c := range row {
			if c == '.' {
				field[i][j] = '#'
				if isStuckInLoop(copyField(field)) {
					newObstructionCount++
				}
				field[i][j] = '.'
			}
		}
	}
	return newObstructionCount
}
