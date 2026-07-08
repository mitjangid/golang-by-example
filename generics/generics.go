package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	numbers := []int{1, 2, 3}
	printIntSlice(numbers)

	languages := []string{"Go", "TypeScript", "Node.js", "Laravel"}
	printStringSlice(languages)

	printSlice(languages)
	printSlice(numbers)

	numberStack := stack[int]{
		elements: []int{1, 2, 3, 4},
	}
	numberStack.push(5)
	fmt.Println(numberStack.pop())

	printComparableSlice(languages)

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}

// Generics were introduced in Go 1.18. They let one function or type work with
// multiple concrete types while keeping compile-time type safety.
func printIntSlice(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printStringSlice(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printSlice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

type stack[T any] struct {
	elements []T
}

func (s *stack[T]) push(item T) {
	s.elements = append(s.elements, item)
}

func (s *stack[T]) pop() (T, bool) {
	var zeroValue T
	if len(s.elements) == 0 {
		return zeroValue, false
	}

	lastIndex := len(s.elements) - 1
	item := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return item, true
}

func printComparableSlice[T comparable](items []T) {
	// comparable allows values that can be checked with == and !=.
	for _, item := range items {
		fmt.Println(item)
	}
}
