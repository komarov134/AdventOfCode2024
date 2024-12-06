package main

import (
	"fmt"
)

// Define a generic Set type using a map
type Set[T comparable] map[T]struct{}

// Method to create a new Set
func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

// Method to add an element to the set
func (s Set[T]) Add(element T) {
	s[element] = struct{}{}
}

func (s Set[T]) AddAll(elements []T) {
	for _, element := range elements {
		s.Add(element)
	}
}

// Method to remove an element from the set
func (s Set[T]) Remove(element T) {
	delete(s, element)
}

// Method to check if the set contains an element
func (s Set[T]) Contains(element T) bool {
	_, exists := s[element]
	return exists
}

// Method to get the size of the set
func (s Set[T]) Size() int {
	return len(s)
}

// Method to print the set
func (s Set[T]) Print() {
	for element := range s {
		fmt.Println(element)
	}
}

func howto() {
	// Create a new set of integers
	intSet := NewSet[int]()

	// Add elements to the set
	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(3)

	// Print the set
	fmt.Println("Integer Set:")
	intSet.Print()

	// Check if the set contains a specific element
	fmt.Println("Contains 2?", intSet.Contains(2)) // true
	fmt.Println("Contains 4?", intSet.Contains(4)) // false

	// Remove an element
	intSet.Remove(2)
	fmt.Println("Set after removing 2:")
	intSet.Print()

	// Create a new set of strings
	strSet := NewSet[string]()

	// Add elements to the string set
	strSet.Add("apple")
	strSet.Add("banana")

	// Print the string set
	fmt.Println("String Set:")
	strSet.Print()
}
