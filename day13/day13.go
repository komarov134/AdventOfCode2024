package day13

import (
	"AdventOfCode2024/utils"
	"fmt"
	"strings"
)

// Button A: X+94, Y+34
// Button B: X+22, Y+67
func parseButton(line string) ButtonConf {
	s := strings.Split(strings.Split(line, ": ")[1], ", ")
	return ButtonConf{DeltaX: utils.MustAtoi64(s[0][2:]), DeltaY: utils.MustAtoi64(s[1][2:])}
}

// Prize: X=8400, Y=5400
func parsePrize(line string) Coord {
	s := strings.Split(strings.Split(line, ": ")[1], ", ")
	return Coord{X: 10000000000000 + utils.MustAtoi64(s[0][2:]), Y: 10000000000000 + utils.MustAtoi64(s[1][2:])}
	// return Coord{X: utils.MustAtoi64(s[0][2:]), Y: utils.MustAtoi64(s[1][2:])}
}

func parseInput(lines []string) []Conf {
	configs := []Conf{}
	for i := 0; i < (len(lines)+1)/4; i++ {
		a := parseButton(lines[4*i+0])
		b := parseButton(lines[4*i+1])
		prize := parsePrize(lines[4*i+2])
		configs = append(configs, Conf{a, b, prize})
	}
	return configs
}

type Coord struct {
	X int64
	Y int64
}

type ButtonConf struct {
	DeltaX int64
	DeltaY int64
}

type Conf struct {
	ButtonA    ButtonConf
	ButtonB    ButtonConf
	PrizeCoord Coord
}

const ButtonAPrice = 3
const ButtonBPrice = 1

func minPrice(conf Conf) int64 {
	minPrice := int64(-1)
	for i := int64(0); i < 101; i++ {
		for j := int64(0); j < 101; j++ {
			currentX := conf.ButtonA.DeltaX*i + conf.ButtonB.DeltaX*j
			currentY := conf.ButtonA.DeltaY*i + conf.ButtonB.DeltaY*j
			currentCoord := Coord{currentX, currentY}
			currentPrice := ButtonAPrice*i + ButtonBPrice*j
			if currentCoord == conf.PrizeCoord && (minPrice == -1 || currentPrice < minPrice) {
				minPrice = currentPrice
			}
		}
	}
	return minPrice
}

func Part1(lines []string) int64 {
	configs := parseInput(lines)
	totalPrice := int64(0)
	for _, conf := range configs {
		fmt.Println(conf)
		mp := minPrice(conf)
		if mp != -1 {
			fmt.Println("min price:", mp)
			totalPrice += mp
		} else {
			fmt.Println("could not find min price")
		}
	}
	return totalPrice
}

// k*A1 + m*B1 = C1
// k*A2 + m*B2 = C2
func minPrice2(conf Conf) int64 {
	A1 := conf.ButtonA.DeltaX
	B1 := conf.ButtonB.DeltaX
	C1 := conf.PrizeCoord.X
	A2 := conf.ButtonA.DeltaY
	B2 := conf.ButtonB.DeltaY
	C2 := conf.PrizeCoord.Y
	// solve equation
	// k*A1 + m*B1 = C1  // multiply by A2
	// k*A2 + m*B2 = C2  // multiply b2 A1
	// and then we get
	// k*A1*A2 + m*B1*A2 = C1*A2
	// k*A2*A1 + m*B2*A1 = C2*A1
	// then we have:
	// m*(B1*A2 - B2*A1) = C1*A2 - C2*A1	// it's good because C1 and C2 are big numbers
	// we can even check if it's divisible
	if B1*A2-B2*A1 == 0 {
		// can't divide by 0
		return -1
	}
	if (C1*A2-C2*A1)%(B1*A2-B2*A1) != 0 {
		// we can't press a button 2.75 times
		return -1
	}
	m := (C1*A2 - C2*A1) / (B1*A2 - B2*A1)
	if m < 0 {
		// doesn't make sence
		return -1
	}
	// now we can find k
	if (C1-m*B1)%A1 != 0 {
		// we can't press a button 2.75 times
		return -1
	}
	k := (C1 - m*B1) / A1

	return k*ButtonAPrice + m*ButtonBPrice
}

func Part2(lines []string) int64 {
	configs := parseInput(lines)
	totalPrice := int64(0)
	for _, conf := range configs {
		fmt.Println(conf)
		mp := minPrice2(conf)
		if mp != -1 {
			fmt.Println("min price:", mp)
			totalPrice += mp
		} else {
			fmt.Println("could not find min price")
		}
	}
	return totalPrice
}
