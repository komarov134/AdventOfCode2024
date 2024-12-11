package day11

import (
	"AdventOfCode2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func applyRules(s int64) []int64 {
	if s == 0 {
		return []int64{1}
	} else {
		str := strconv.FormatInt(s, 10)
		length := len(str)
		if length%2 == 0 {
			s1 := str[0 : length/2]
			s2 := str[length/2 : length]
			return []int64{utils.MustAtoi64(s1), utils.MustAtoi64(s2)}
		} else {
			return []int64{2024 * s}
		}
	}
}

func Part1(lines []string) int {
	stones := utils.MapSlice(strings.Split(lines[0], " "), func(s string) int64 {
		return utils.MustAtoi64(s)
	})
	fmt.Println(stones)

	for i := 0; i < 75; i++ {
		newStones := []int64{}
		for _, s := range stones {
			newStones = append(newStones, applyRules(s)...)
		}
		fmt.Println("step", i, "length", len(newStones))
		stones = newStones
	}

	return len(stones)
}

// cache for each layer
type LeveledCache struct {
	Caches []map[int64]int64
}

func (c LeveledCache) Get(layer int, key int64) (int64, bool) {
	value, exists := c.Caches[layer][key]
	return value, exists
}

func (c LeveledCache) Put(layer int, key int64, value int64) {
	c.Caches[layer][key] = value
}

func count(level int, s int64, cache LeveledCache) int64 {
	c, exists := cache.Get(level, s)
	if exists {
		return c
	} else if level == 0 {
		return 1
	} else {
		totalCount := int64(0)
		for _, newS := range applyRules(s) {
			totalCount += count(level-1, newS, cache)
		}
		cache.Put(level, s, totalCount)
		return totalCount
	}
}

func Part2(lines []string) int64 {
	stones := utils.MapSlice(strings.Split(lines[0], " "), func(s string) int64 {
		return utils.MustAtoi64(s)
	})
	fmt.Println(stones)

	blinkCount := 75
	totalCount := int64(0)
	cache := LeveledCache{[]map[int64]int64{}}
	for i := 0; i <= blinkCount; i++ {
		cache.Caches = append(cache.Caches, map[int64]int64{})
	}
	for _, s := range stones {
		totalCount += count(blinkCount, s, cache)
	}

	return totalCount
}
