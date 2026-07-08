package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	count := 20

	// Go has one loop keyword: for.
	for i := 0; i < count; i++ {
		fmt.Printf("Counting up: %d\n", i)
	}

	// A for loop can also work like a while loop.
	for count > 0 {
		fmt.Printf("Counting down: %d\n", count)
		count--
	}

	// Use break to leave a loop and continue to skip the current iteration.
	for count < 50 {
		count++
		if count%9 == 0 {
			continue
		}

		fmt.Printf("Filtered count: %d\n", count)

		if count == 30 {
			break
		}
	}

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}
