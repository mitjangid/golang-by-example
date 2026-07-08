package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// This root file is intentionally small. Run the examples inside each folder
	// to explore one Go concept at a time.
	fmt.Println("Go quick overview examples")

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}
