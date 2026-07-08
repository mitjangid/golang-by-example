package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// Go passes arguments by value. This function receives a copy.
	number := 3
	changeNumberByValue(number)
	fmt.Println("Value after pass by value:", number)

	// A pointer stores the address of a value. Passing a pointer lets a function
	// update the original value.
	fmt.Println("Memory address:", &number)
	changeNumberByPointer(&number)
	fmt.Println("Value after pointer update:", number)

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}

func changeNumberByValue(number int) int {
	number = 5
	return number
}

func changeNumberByPointer(number *int) {
	*number = 5 // * dereferences the pointer so the original value can be changed.
}
