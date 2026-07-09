package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// main is the program entry point. Other functions are called from here.
	values, learnerName, learnerAge := buildProfile(1, 2, "Amit")
	// Use _ when you intentionally ignore a returned value.
	// values, learnerName, _ := buildProfile(1, 2, "Amit")
	fmt.Println(values, learnerName, learnerAge)

	double := func(number int) int {
		return number * 2
	}
	fmt.Println(applyToNumber(double, 5))

	demonstrateVariadicFunctions()
	demonstrateClosures()

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}

func buildProfile(firstNumber, secondNumber int, name string) ([3]int, string, int) {
	// Go functions can return multiple values.
	fmt.Println("Input:", firstNumber, secondNumber, name)
	return [3]int{1, 2, 3}, "Amit", 37
}

func applyToNumber(transform func(number int) int, value int) int {
	// Functions can be passed as values.
	return transform(value)
}

func demonstrateVariadicFunctions() {
	numbers := []int{1, 2, 3, 4}
	printNumbers(numbers...) // Use ... to pass a slice to a variadic function.
}

func printNumbers(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

func demonstrateClosures() {
	increment := newCounter()

	fmt.Println(increment())
	fmt.Println(increment())
}

func newCounter() func() int {
	count := 0
	return func() int {
		// The returned function keeps access to count after newCounter returns.
		count++
		return count
	}
}
