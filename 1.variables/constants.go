package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// Constants are assigned at compile time and cannot be reassigned.
	const languageName string = "Go"
	// languageName = "Golang" // Uncommenting this line causes a compile error.

	fmt.Println(languageName)

	// Group related constants in a const block.
	const (
		starterPrice    = 500
		professionalFee = 2000
		gstRate         = "18%"
	)

	fmt.Println(
		starterPrice,
		professionalFee,
		gstRate,
	)

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}
