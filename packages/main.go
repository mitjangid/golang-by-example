package main

import (
	"fmt"
	"time"
)

func main() { // entry package file is default main.go
	start := time.Now()

	testPackage

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}
