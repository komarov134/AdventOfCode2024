package day12

import (
	"AdventOfCode2024/utils"
)

// import "fmt"

func parseInput(lines []string) [][]Cell {
	field := [][]Cell{}
	for _, line := range lines {
		row := []Cell{}
		for _, c := range line {
			row = append(row, Cell{Plant: byte(c), RegionNumber: 0, Fences: nil})
		}
		field = append(field, row)
	}
	return field
}

type Cell struct {
	Plant        byte
	RegionNumber int
	Fences       []int
}

func isPlant(plant byte, i int, j int, field [][]Cell) bool {
	return i >= 0 && i < len(field) && j >= 0 && j < len(field[0]) && field[i][j].Plant == plant
}

var DxDy = [4][2]int{
	{0, -1}, // left
	{-1, 0}, // top
	{0, +1}, // right
	{+1, 0}, // down
}

// enumerate and build fences
func bfs(plant byte, regionNumber int, i int, j int, field [][]Cell) {
	// out of field OR another plant OR visited
	if i < 0 || i >= len(field) || j < 0 || j >= len(field[0]) || field[i][j].Plant != plant || field[i][j].RegionNumber == regionNumber {
		return
	}
	field[i][j].RegionNumber = regionNumber // set visited
	for d, dxdy := range DxDy {
		x, y := i+dxdy[0], j+dxdy[1]
		if !isPlant(plant, x, y, field) {
			field[i][j].Fences = append(field[i][j].Fences, d)
		}
		bfs(plant, regionNumber, x, y, field)
	}
}

// price = area * perimeter
func price(regionNumber int, field [][]Cell) int {
	fences := 0
	area := 0
	for _, row := range field {
		for _, cell := range row {
			if cell.RegionNumber == regionNumber {
				fences += len(cell.Fences)
				area++
			}
		}
	}
	return area * fences
}

func Part1(lines []string) int {
	field := parseInput(lines)

	regionNumber := 1
	for i, row := range field {
		for j, cell := range row {
			if cell.RegionNumber == 0 { // not visited
				bfs(cell.Plant, regionNumber, i, j, field)
				regionNumber++
			}
		}
	}

	totalPrice := 0
	for r := 1; r < regionNumber; r++ {
		totalPrice += price(r, field)
	}

	return totalPrice
}

func dropAdjacentFences(plant byte, field [][]Cell, fence int, x int, y int, direction int) {
	for isPlant(plant, x, y, field) && utils.Contains(field[x][y].Fences, fence) {
		// fmt.Println("dropping", fence, "for", x, y)
		field[x][y].Fences = utils.Remove(field[x][y].Fences, fence)
		x, y = x+DxDy[direction][0], y+DxDy[direction][1]
	}
}

func dropAllAdjacentFences(field [][]Cell, i int, j int) {
	// fmt.Println("dropAllAdjacentFences", i, j, "Fences", field[i][j].Fences)
	plant := field[i][j].Plant
	for _, fence := range field[i][j].Fences {
		// if fence is left or right then go up, go down
		if fence == 0 || fence == 2 {
			dropAdjacentFences(plant, field, fence, i+DxDy[1][0], j+DxDy[1][1], 1)
			dropAdjacentFences(plant, field, fence, i+DxDy[3][0], j+DxDy[3][1], 3)
		} else { // if d is top or down then go left, go right
			dropAdjacentFences(plant, field, fence, i+DxDy[0][0], j+DxDy[0][1], 0)
			dropAdjacentFences(plant, field, fence, i+DxDy[2][0], j+DxDy[2][1], 2)
		}
	}
}

// price = area * (number of sides)
func price2(regionNumber int, field [][]Cell) int {
	fences := 0
	area := 0
	for _, row := range field {
		for _, cell := range row {
			if cell.RegionNumber == regionNumber {
				area++
			}
		}
	}
	for i, row := range field {
		for j, cell := range row {
			if cell.RegionNumber == regionNumber {
				dropAllAdjacentFences(field, i, j)
			}
		}
	}
	for _, row := range field {
		for _, cell := range row {
			if cell.RegionNumber == regionNumber {
				fences += len(cell.Fences)
			}
		}
	}
	return area * fences
}

func Part2(lines []string) int {
	field := parseInput(lines)

	regionNumber := 1
	for i, row := range field {
		for j, cell := range row {
			if cell.RegionNumber == 0 { // not visited
				bfs(cell.Plant, regionNumber, i, j, field)
				regionNumber++
			}
		}
	}
	// for i, row := range field {
	// 	for j, cell := range row {
	// 		fmt.Println("Fences for", i, j, "are", cell.Fences)
	// 	}
	// }

	totalPrice := 0
	for r := 1; r < regionNumber; r++ {
		totalPrice += price2(r, field)
	}

	// for i, row := range field {
	// 	for j, cell := range row {
	// 		fmt.Println("Fences for", i, j, "are", cell.Fences)
	// 	}
	// }

	return totalPrice
}
