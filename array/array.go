package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// Arrays have a fixed length that is part of their type.
	var numbers [3]int
	numbers[0] = 2
	numbers[1] = 3
	numbers[2] = 4
	fmt.Println(len(numbers))
	fmt.Println(numbers)

	var flags [4]bool
	flags[2] = true
	fmt.Println(flags)

	var messages [2]string
	messages[0] = "The user"
	messages[1] = "is learning Go"
	fmt.Println(messages)

	// Values can be assigned when the array is declared.
	// fixedNumbers := [3]int{1, 2, 3, 4} // Compile error: too many values.
	fixedNumbers := [3]int{1, 2, 3}
	fmt.Println(fixedNumbers)

	// Multidimensional arrays are arrays whose elements are arrays.
	profileTable := [...][2]string{{"name", "age"}, {"Amit", "37"}}
	fmt.Println(profileTable)

	// Use arrays when the size is known and should not change.

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}
