package utils

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

func MustAtoi64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
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

func MapSlice[T any, U any](arr []T, transform func(T) U) []U {
	var result []U
	for _, v := range arr {
		result = append(result, transform(v))
	}
	return result
}

func Contains[T comparable](slice []T, elem T) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

func Remove[T comparable](slice []T, elem T) []T {
	var result []T
	for _, v := range slice {
		if v != elem {
			result = append(result, v)
		}
	}
	return result
}
