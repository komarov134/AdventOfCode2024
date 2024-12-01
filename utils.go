package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFileString(name string) string {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Error reading file", name, "got error", err)
	}
	return string(data)
}

func ReadFileLines(name string) []string {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	return strings.Split(string(data), "\n")
}

func MustAtoi(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err) // You can change this to return a default value if preferred
	}
	return num
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
