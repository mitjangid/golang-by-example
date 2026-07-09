package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// Print values and simple expressions.
	fmt.Println("Hello world")
	fmt.Println(1)
	fmt.Println(1 + 1)
	fmt.Println(700 / 2)
	fmt.Println(3 * 6)
	fmt.Println(1.0 + 3.5)
	fmt.Println(true)
	fmt.Println(nil)

	// Explicit variable declarations include the type.
	var learnerName string = "Amit"
	fmt.Println(learnerName)

	var isLearning bool = true
	fmt.Println(isLearning)

	var lessonCount int = 5
	fmt.Println(lessonCount)

	var completionRatio float32 = 1.5
	fmt.Println(completionRatio)

	// Short declaration syntax infers the type and can be used inside functions.
	language := "Go"
	fmt.Println(language)

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}
