package main

import (
	"strings"
)

func Day04Part1(lines []string) int {
	result := 0

	matrix := [][]byte{}
	for _, line := range lines {
		matrix = append(matrix, []byte(line))
	}

	allLines := []string{}
	allLines = append(allLines, allRows(matrix)...)
	matrix = transpose(matrix)
	allLines = append(allLines, allRows(matrix)...)
	matrix = transpose(matrix)
	allLines = append(allLines, allDiagonals(matrix)...)
	matrix = swapColumns(matrix)
	allLines = append(allLines, allDiagonals(matrix)...)

	for _, line := range allLines {
		result += countXmas(line)
	}
	return result
}

func allRows(matrix [][]byte) []string {
	result := []string{}
	for _, bytes := range matrix {
		result = append(result, string(bytes))
	}
	return result
}

func transpose(matrix [][]byte) [][]byte {
	for i := range len(matrix) {
		for j := i + 1; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func allDiagonals(matrix [][]byte) []string {
	result := []string{}
	for i := 0; i < len(matrix); i++ {
		j := 0
		bytes := []byte{}
		for d := 0; i+d < len(matrix) && j+d < len(matrix[0]); d++ {
			bytes = append(bytes, matrix[i+d][j+d])
		}
		result = append(result, string(bytes))
	}
	for j := 1; j < len(matrix[0]); j++ {
		i := 0
		bytes := []byte{}
		for d := 0; i+d < len(matrix) && j+d < len(matrix[0]); d++ {
			bytes = append(bytes, matrix[i+d][j+d])
		}
		result = append(result, string(bytes))
	}
	return result
}

func swapRows(matrix [][]byte) [][]byte {
	lastIndex := len(matrix) - 1
	for i := 0; i < len(matrix)/2; i++ {
		matrix[i], matrix[lastIndex-i] = matrix[lastIndex-i], matrix[i]
	}
	return matrix
}

func swapColumns(matrix [][]byte) [][]byte {
	matrix = transpose(matrix)
	matrix = swapRows(matrix)
	matrix = transpose(matrix)
	return matrix
}

// count written backwards as well
func countXmas(line string) int {
	return strings.Count(line, "XMAS") + strings.Count(line, "SAMX")
}

var Xmas1 = [3][3]byte{
	{'M', '.', 'M'},
	{'.', 'A', '.'},
	{'S', '.', 'S'},
}
var Xmas2 = [3][3]byte{
	{'M', '.', 'S'},
	{'.', 'A', '.'},
	{'M', '.', 'S'},
}
var Xmas3 = [3][3]byte{
	{'S', '.', 'M'},
	{'.', 'A', '.'},
	{'S', '.', 'M'},
}
var Xmas4 = [3][3]byte{
	{'S', '.', 'S'},
	{'.', 'A', '.'},
	{'M', '.', 'M'},
}
var AllXmas = [4][3][3]byte{Xmas1, Xmas2, Xmas3, Xmas4}

func matchesOneXmas(matrix [][]byte, posX int, posY int, xmas [3][3]byte) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if xmas[i][j] != '.' && xmas[i][j] != matrix[posX+i][posY+j] {
				return false
			}
		}
	}
	return true
}

// (x, y) is the left-top corner of X-MAS
func matchesXmas(matrix [][]byte, posX int, posY int) bool {
	if posX+2 >= len(matrix) || posY+2 >= len(matrix[0]) {
		return false
	}
	for _, xmas := range AllXmas {
		if matchesOneXmas(matrix, posX, posY, xmas) {
			return true
		}
	}
	return false
}

func Day04Part2(lines []string) int {
	result := 0

	matrix := [][]byte{}
	for _, line := range lines {
		matrix = append(matrix, []byte(line))
	}

	for i := range len(matrix) {
		for j := range len(matrix[0]) {
			if matchesXmas(matrix, i, j) {
				result += 1
			}
		}
	}

	return result
}
