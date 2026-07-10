package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// ============================================
	// BASIC PANIC
	// ============================================
	// panic() stops normal execution immediately
	// It unwinds the stack (runs deferred functions in reverse order)
	// If not recovered, the program crashes with a stack trace
	// Use panic ONLY for truly unrecoverable situations
	// For normal errors, return error values instead

	fmt.Println("=== Basic Panic ===")
	// This function will panic - but we've commented out the actual panic
	// Uncomment the panic line to see the program crash
	demonstrateBasicPanic()

	// ============================================
	// PANIC WITH RECOVER
	// ============================================
	// recover() regains control after a panic
	// IMPORTANT: recover() only works when called directly from a deferred function
	// If recover() is not called, the panic continues and the program crashes

	fmt.Println("\n=== Panic with Recover ===")
	demonstratePanicWithRecover()

	// ============================================
	// RECOVER WITH NAMED RETURN VALUES
	// ============================================
	// You can use named return values with recover
	// This lets you return error values even after a panic
	// This is a common pattern for converting panics to errors

	fmt.Println("\n=== Recover with Named Returns ===")

	// This will panic and be converted to an error
	result, err := safeDivision(10, 0)
	fmt.Printf("Result: %d, Error: %v\n", result, err)

	// This will succeed normally
	result, err = safeDivision(10, 2)
	fmt.Printf("Result: %d, Error: %v\n", result, err)

	// ============================================
	// RECOVER IN CLEANUP
	// ============================================
	// defer is often used for cleanup (closing files, releasing resources)
	// If a panic occurs, deferred functions still run
	// This ensures cleanup happens even when things go wrong

	fmt.Println("\n=== Recover in Cleanup ===")
	demonstrateCleanup()

	// ============================================
	// PANIC WITH VALUES
	// ============================================
	// panic() can accept any value, not just strings
	// You can panic with integers, structs, errors, etc.
	// Use type assertion in recover() to handle different types

	fmt.Println("\n=== Panic with Values ===")
	demonstrateValuePanic()

	duration := time.Since(start)
	fmt.Printf("\nCode execution took: %s\n", duration)
}

// demonstrateBasicPanic shows what happens when panic is called
// NOTE: The panic is commented out so this example doesn't crash
func demonstrateBasicPanic() {
	fmt.Println("About to panic...")
	// Uncomment the next line to see the program crash
	// panic("something went terribly wrong!")
	// This line will never execute if panic is called
	fmt.Println("This will not print if panic occurs")
}

// demonstratePanicWithRecover shows how to recover from a panic
// The defer function runs even when panic occurs
func demonstratePanicWithRecover() {
	// defer schedules this function to run when the current function exits
	// It runs whether the function exits normally or panics
	defer func() {
		// recover() returns the panic value if a panic occurred
		// It returns nil if there was no panic
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	fmt.Println("Starting operation...")
	// This panic will be caught by the deferred recover
	panic("operation failed unexpectedly")
	// This line will never execute
	fmt.Println("This will not print")
}

// safeDivision performs division safely
// Uses named return values to modify them in the deferred recover
func safeDivision(a, b int) (result int, err error) {
	// Named return values (result, err) are available in the entire function
	// The defer function can modify them even after a panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in safeDivision: %v\n", r)
			// Convert the panic to an error
			err = fmt.Errorf("division by zero: %v", r)
			result = 0
		}
	}()

	if b == 0 {
		// Panic instead of returning an error
		// The deferred function will convert this to an error
		panic("cannot divide by zero")
	}
	return a / b, nil
}

// demonstrateCleanup shows the common pattern of cleanup with recover
func demonstrateCleanup() {
	// This defer ensures cleanup happens even if panic occurs
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered during cleanup: %v\n", r)
		}
		// This cleanup code runs regardless of panic
		fmt.Println("Cleanup completed")
	}()

	fmt.Println("Acquiring resources...")
	// Simulate resource acquisition that might fail
	// If panic occurs, the deferred cleanup still runs
	panic("resource acquisition failed")
	// This line never executes
	fmt.Println("Resources acquired")
}

// demonstrateValuePanic shows that panic can accept any type
func demonstrateValuePanic() {
	defer func() {
		if r := recover(); r != nil {
			// Type assertion (switch) to handle different panic types
			switch v := r.(type) {
			case string:
				fmt.Printf("Recovered string panic: %s\n", v)
			case int:
				fmt.Printf("Recovered int panic: %d\n", v)
			case error:
				fmt.Printf("Recovered error panic: %v\n", v)
			default:
				fmt.Printf("Recovered unknown panic type: %v\n", v)
			}
		}
	}()

	// Try different panic types by uncommenting:
	panic(42) // Integer
	// panic("custom error message") // String
	// panic(fmt.Errorf("wrapped error")) // Error type
	// panic(struct{ Field string }{Field: "custom struct"}) // Struct
}
