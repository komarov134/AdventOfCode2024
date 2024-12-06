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

func mapSlice[T any, U any](arr []T, transform func(T) U) []U {
	var result []U
	for _, v := range arr {
		result = append(result, transform(v))
	}
	return result
}

func contains[T comparable](slice []T, elem T) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}
