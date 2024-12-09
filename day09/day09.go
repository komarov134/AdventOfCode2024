package day09

import (
	"AdventOfCode2024/utils"
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

const FreeBlock = -1

func findNextFreeBlockIndex(from int, blocks []int) int {
	i := from
	for i = from; i < len(blocks) && blocks[i] != FreeBlock; i++ {
	}
	return i
}

func findNextFileBlockIndex(from int, blocks []int) int {
	i := from
	for i = from; i >= 0 && blocks[i] == FreeBlock; i-- {
	}
	return i
}

func checksum(blocks []int) int64 {
	sum := int64(0)
	for i, b := range blocks {
		if b != FreeBlock {
			sum += int64(i * b)
		}
	}
	return sum
}

func Part1(lines []string) int64 {
	line := lines[0]
	blocks := []int{}
	readFile := true
	fileId := 0
	for _, c := range line {
		digit := utils.MustAtoi(string(c))
		if readFile {
			// append 'digit' blocks for fileId
			for i := 0; i < digit; i++ {
				blocks = append(blocks, fileId)
			}
			fileId++
		} else {
			// append 'digit' free blocks
			for i := 0; i < digit; i++ {
				blocks = append(blocks, FreeBlock)
			}
		}
		readFile = !readFile
	}
	// fmt.Println(blocks)
	nextFree := findNextFreeBlockIndex(0, blocks)
	nextFile := findNextFileBlockIndex(len(blocks)-1, blocks)
	for nextFree < nextFile {
		blocks[nextFree], blocks[nextFile] = blocks[nextFile], blocks[nextFree]
		nextFree = findNextFreeBlockIndex(nextFree+1, blocks)
		nextFile = findNextFileBlockIndex(nextFile-1, blocks)
	}
	// fmt.Println(blocks)

	return checksum(blocks)
}

type Block struct {
	Start int
	End   int
}

func (b Block) Size() int {
	return b.End - b.Start + 1
}

func nextFreeBlock(from int, blocks []int) Block {
	start := findNextFreeBlockIndex(from, blocks)
	end := start
	for end < len(blocks) && blocks[end] == FreeBlock {
		end++
	}
	return Block{start, end - 1}
}

func nextFileBlock(from int, blocks []int) Block {
	end := findNextFileBlockIndex(from, blocks)
	fileId := blocks[end]
	start := end
	for start >= 0 && blocks[start] == fileId {
		start--
	}
	return Block{start + 1, end}
}

func buildFileBlocks(blocks []int) []Block {
	fileBlocks := []Block{}
	i := 0
	for i < len(blocks) {
		fileId := blocks[i]
		start := i
		for i < len(blocks) && blocks[i] == fileId {
			i++
		}
		end := i - 1
		fileBlocks = append(fileBlocks, Block{start, end})
		for i < len(blocks) && blocks[i] == FreeBlock {
			i++
		}
	}
	return fileBlocks
}

func Part2(lines []string) int64 {
	line := lines[0]
	blocks := []int{}
	readFile := true
	fileId := 0
	for _, c := range line {
		digit := utils.MustAtoi(string(c))
		if readFile {
			// append 'digit' blocks for fileId
			for i := 0; i < digit; i++ {
				blocks = append(blocks, fileId)
			}
			fileId++
		} else {
			// append 'digit' free blocks
			for i := 0; i < digit; i++ {
				blocks = append(blocks, FreeBlock)
			}
		}
		readFile = !readFile
	}
	fmt.Println(blocks)
	fileBlocks := buildFileBlocks(blocks)
	fmt.Println(fileBlocks)

	for i := len(fileBlocks) - 1; i >= 0; i-- {
		b := fileBlocks[i]
		nextFree := nextFreeBlock(0, blocks)
		moved := false
		for nextFree.Start < b.Start && !moved {
			if nextFree.Size() >= b.Size() {
				// move the whole block
				for i := 0; i < b.Size(); i++ {
					blocks[nextFree.Start+i], blocks[b.Start+i] = blocks[b.Start+i], blocks[nextFree.Start+i]
				}
				moved = true
			} else {
				nextFree = nextFreeBlock(nextFree.End+1, blocks) // skip the rest of free block
			}
		}
	}
	fmt.Println(blocks)

	return checksum(blocks)
}
