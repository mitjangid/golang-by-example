package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	age := 18

	// Use if/else for branching.
	if age >= 18 {
		fmt.Println("Person is an adult")
	} else {
		fmt.Println("Person is not an adult")
	}

	// Add else-if branches when more than two outcomes are possible.
	if age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 12 {
		fmt.Println("Person is a teenager")
	} else {
		fmt.Println("Person is a child")
	}

	// Combine conditions with logical operators.
	role := "admin"
	hasPermission := true
	if role == "admin" && hasPermission {
		fmt.Println("You are allowed")
	}

	// A short statement can initialize a value scoped to the if block.
	if value := 15; value >= 18 {
		fmt.Println("Person is an adult")
	} else {
		fmt.Println("Person is below adult age")
	}

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}
