package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	// ============================================
	// BASIC ERROR CHECKING
	// ============================================
	// In Go, functions that can fail return (result, error)
	// Always check if error is nil before using the result
	// If error is not nil, something went wrong

	fmt.Println("=== Basic Error Checking ===")

	// This will succeed - no error
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	// This will fail - division by zero
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // This will print
	} else {
		fmt.Println("Result:", result) // This won't run
	}

	// ============================================
	// MULTIPLE RETURN VALUES
	// ============================================
	// Go functions can return multiple values
	// Common pattern: (result, error)
	// Use _ to ignore a value (but don't ignore errors in real code!)

	fmt.Println("\n=== Multiple Return Values ===")

	// Try to convert a string to a number
	number, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("Conversion failed:", err)
	} else {
		fmt.Println("Converted string '42' to number:", number)
	}

	// This conversion will fail because "hello" is not a number
	number, err = strconv.Atoi("hello")
	if err != nil {
		fmt.Println("Conversion failed:", err) // This will print
	} else {
		fmt.Println("Converted number:", number) // This won't run
	}

	// ============================================
	// ERROR WRAPPING
	// ============================================
	// Wrap errors to add more context about where/why they happened
	// Use %w verb to preserve the original error for checking later
	// Use %v if you just want to format the error message

	fmt.Println("\n=== Error Wrapping ===")

	// Valid input - this will work
	validNumber, err := parseAndValidate("42")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Valid number:", validNumber)
	}

	// Invalid input - this will fail with a wrapped error
	invalidNumber, err := parseAndValidate("abc")
	if err != nil {
		fmt.Println("Error:", err)
		// Unwrap reveals the original error underneath
		fmt.Println("Original error:", errors.Unwrap(err))
	} else {
		fmt.Println("Valid number:", invalidNumber)
	}

	// ============================================
	// SHORT IF STATEMENT
	// ============================================
	// You can declare variables in the if statement itself
	// The variable is only visible inside the if/else block

	fmt.Println("\n=== Short If Statement ===")

	if value, err := safeOperation("valid"); err != nil {
		fmt.Println("Operation failed:", err)
	} else {
		fmt.Println("Operation succeeded:", value)
	}

	duration := time.Since(start)
	fmt.Printf("\nCode execution took: %s\n", duration)
}

// divide divides two numbers
// Returns the result and an error if dividing by zero
func divide(a, b int) (int, error) {
	// Check for the error condition first
	if b == 0 {
		// Return zero value and an error
		return 0, errors.New("cannot divide by zero")
	}
	// Return the result and nil (no error)
	return a / b, nil
}

// parseAndValidate converts a string to an integer and validates it
// Wraps errors with additional context using %w
func parseAndValidate(input string) (int, error) {
	// Try to convert string to integer
	value, err := strconv.Atoi(input)
	if err != nil {
		// Wrap the error with more context
		// %w preserves the original error for error checking
		return 0, fmt.Errorf("failed to parse '%s': %w", input, err)
	}

	// Validate the value
	if value < 0 {
		return 0, fmt.Errorf("validation failed: %d is negative", value)
	}

	// Return the valid result
	return value, nil
}

// safeOperation performs a simple operation that might fail
func safeOperation(input string) (string, error) {
	if input == "" {
		return "", errors.New("input cannot be empty")
	}
	return "processed: " + input, nil
}
