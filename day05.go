package main

import (
	"fmt"
	"strings"
)

type Rule struct {
	First  int
	Second int
}

func parseInput(lines []string) ([]Rule, [][]int) {
	rules := []Rule{}
	manuals := [][]int{}
	readRules := true
	for _, line := range lines {
		if line == "" {
			readRules = false
		} else if readRules {
			ruleStr := strings.Split(line, "|")
			rule := Rule{MustAtoi(ruleStr[0]), MustAtoi(ruleStr[1])}
			rules = append(rules, rule)
		} else {
			manual := mapSlice(strings.Split(line, ","), func(s string) int {
				return MustAtoi(s)
			})
			manuals = append(manuals, manual)
		}
	}
	return rules, manuals
}

func Day05Part1(lines []string) int {
	rules, manuals := parseInput(lines)

	result := 0
	for _, manual := range manuals {
		if checkManual(manual, rules) {
			result += getMiddle(manual)
		}
	}
	return result
}

func checkManual(manual []int, rules []Rule) bool {
	for i := range manual {
		for j := i + 1; j < len(manual); j++ {
			if (!contains(rules, Rule{manual[i], manual[j]})) {
				return false
			}
		}
	}
	return true
}

func getMiddle(manual []int) int {
	return manual[len(manual)/2]
}

func fixManual(manual []int, rules []Rule) []int {
	return sortManual(manual, rules)
}

// This is full ordering. Rule X|Y means X<Y. No rule X|Y means X>Y.
// 97 before [13, 61, 47, 29, 53, 75]
// 75 before [29, 47, 61, 13, 53]
// 47 before [53, 13, 61, 29]
// 61 before [13, 53, 29]
// 53 before [29, 13]
// 29 before [13]
//
// 75,97,47,61,53 becomes 97,75,47,61,53.
// 61,13,29 becomes 61,29,13.
// 97,13,75,29,47 becomes 97,75,47,29,13
func Day05Part2(lines []string) int {
	rules, manuals := parseInput(lines)

	result := 0
	for _, manual := range manuals {
		if !checkManual(manual, rules) {
			fixedManual := fixManual(manual, rules)
			result += getMiddle(fixedManual)
		}
	}
	return result
}

// This is full ordering. Rule X|Y means X<Y. No rule X|Y means X>Y.
// This means we can build full sorted array from rules
func buildFullManual(rules []Rule) []int {
	allElements := []int{}
	for _, rule := range rules {
		if !contains(allElements, rule.First) {
			allElements = append(allElements, rule.First)
		}
		if !contains(allElements, rule.Second) {
			allElements = append(allElements, rule.Second)
		}
	}
	return allElements
}

// This is full ordering. Rule X|Y means X<Y. No rule X|Y means X>Y.
// This means we can build full sorted array from manuals
func buildFullManual2(manuals [][]int) []int {
	allElements := []int{}
	for _, manual := range manuals {
		for _, el := range manual {
			if !contains(allElements, el) {
				allElements = append(allElements, el)
			}
		}
	}
	return allElements
}

func sortManual(manual []int, rules []Rule) []int {
	for i := range manual {
		for j := i + 1; j < len(manual); j++ {
			isLess := contains(rules, Rule{manual[i], manual[j]})
			if !isLess {
				manual[i], manual[j] = manual[j], manual[i]
			}
		}
	}
	return manual
}

func fixManualOptimized(manual []int, fullManual []int) []int {
	fixedManual := []int{}
	for _, el := range fullManual {
		if contains(manual, el) {
			fixedManual = append(fixedManual, el)
		}
	}
	return fixedManual
}

// doesn't work
func Day05Part2Optimized(lines []string) int {
	rules, manuals := parseInput(lines)
	result := 0
	fullManual := buildFullManual2(manuals)
	fmt.Println(fullManual)
	fullManual = sortManual(fullManual, rules)
	fmt.Println(fullManual)
	for _, manual := range manuals {
		if !checkManual(manual, rules) {
			fixedManual := fixManualOptimized(manual, fullManual)
			result += getMiddle(fixedManual)
		}
	}
	return result
}
