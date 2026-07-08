package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// A switch compares one value against multiple cases.
	dayNumber := 1
	switch dayNumber {
	case 1:
		fmt.Println("Monday")
		fallthrough //this will continue to next case even case 1 is true
	case 2:
		fmt.Println("Tuesday")
		fallthrough 
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	default:
		fmt.Println("Unknown day")
	}

	// A case can match more than one value.
	switch time.Now().Weekday() {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("It's a working day")
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("Invalid weekday")
	}

	// A type switch checks the dynamic type stored in an interface value.
	printType := func(value any) {
		switch typedValue := value.(type) {
		case int:
			fmt.Println("Integer type:", typedValue)
		case bool:
			fmt.Println("Boolean type:", typedValue)
		case string:
			fmt.Println("String type:", typedValue)
		default:
			fmt.Println("Unknown type:", typedValue)
		}
	}
	printType(50)

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}
