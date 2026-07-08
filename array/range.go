package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	languageDetails := map[string]string{
		"name":    "Go",
		"type":    "Programming language",
		"level":   "Intermediate",
		"version": "1.21+",
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// range works with arrays, slices, maps, strings, and channels.
	for key, value := range languageDetails {
		fmt.Println(key, "-", value)
	}

	for index, number := range numbers {
		fmt.Println(index, number)
	}

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}
