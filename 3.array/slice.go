package main

import (
	"fmt"
	"slices"
	"time"
)

func main() {
	start := time.Now()

	// Slices are flexible views over arrays. A nil slice has no backing array yet.
	profile := [][]string{{"name", "Amit"}, {"age", "37"}}
	fmt.Println(profile)

	// make creates a slice with length 2 and capacity 5.
	numbers := make([]int, 2, 5)
	fmt.Println("capacity:", cap(numbers))
	numbers[0] = 34
	numbers[1] = 44
	fmt.Println(numbers)

	// append adds values and may grow the underlying array.
	numbers = append(numbers, 54, 64, 74, 84, 94)
	fmt.Println(numbers)
	fmt.Println("capacity after append:", cap(numbers))

	// copy copies elements into another slice.
	copiedNumbers := make([]int, len(numbers))
	copy(copiedNumbers, numbers)
	fmt.Println(copiedNumbers)

	// Slice expressions use half-open ranges: start is included, end is excluded.
	fmt.Println(numbers[0:2])
	fmt.Println(numbers[:2])
	fmt.Println(numbers[len(numbers)-2:])

	// The slices package provides helpers for common slice operations.
	fmt.Println(slices.Equal(numbers, copiedNumbers))
	numbers = append(numbers, 104)
	fmt.Println(slices.Equal(numbers, copiedNumbers))

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}
