package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// ============================================
	// BASIC DEFER
	// ============================================
	// defer schedules a function call to run when the surrounding function exits
	// The deferred call runs even if the function panics
	// Deferred functions run in LIFO order (last deferred, first executed)

	fmt.Println("=== Basic Defer ===")

	fmt.Println("Before defer")
	defer fmt.Println("This runs last (deferred)")
	fmt.Println("After defer")

	// ============================================
	// MULTIPLE DEFERRED CALLS
	// ============================================
	// Multiple deferred calls run in reverse order
	// Think of it as a stack - last in, first out

	fmt.Println("\n=== Multiple Deferred Calls ===")

	defer fmt.Println("First deferred - runs last")
	defer fmt.Println("Second deferred - runs second")
	defer fmt.Println("Third deferred - runs first")
	fmt.Println("All defers registered")

	// ============================================
	// DEFER WITH PARAMETERS
	// ============================================
	// Parameters are evaluated when defer is called, not when it executes
	// This is important to understand!

	fmt.Println("\n=== Defer with Parameters ===")

	count := 1
	defer fmt.Printf("Count at defer time: %d\n", count)
	count = 100
	fmt.Printf("Count now: %d\n", count)
	// The deferred function will print 1, not 100

	// ============================================
	// DEFER FOR RESOURCE CLEANUP
	// ============================================
	// defer is commonly used for cleanup: closing files, releasing locks, etc.
	// This ensures cleanup happens even if an error occurs

	fmt.Println("\n=== Defer for Resource Cleanup ===")

	processFile()
	processWithLock()

	// ============================================
	// DEFER WITH NAMED RETURN VALUES
	// ============================================
	// Deferred functions can modify named return values
	// This is useful for error handling and logging

	fmt.Println("\n=== Defer with Named Return Values ===")

	result := calculateWithDefer()
	fmt.Printf("Result from calculateWithDefer: %d\n", result)

	// ============================================
	// DEFER FOR TIMING
	// ============================================
	// defer is great for measuring function execution time

	fmt.Println("\n=== Defer for Timing ===")

	measureExecutionTime()

	// ============================================
	// DEFER IN LOOPS
	// ============================================
	// Be careful with defer in loops - it can consume a lot of resources
	// Each iteration adds a deferred call to the stack

	fmt.Println("\n=== Defer in Loops ===")

	deferInLoops()

	duration := time.Since(start)
	fmt.Printf("\nCode execution took: %s\n", duration)
}

// ============================================
// RESOURCE CLEANUP EXAMPLES
// ============================================

// processFile simulates opening and closing a file
func processFile() {
	fmt.Println("Opening file...")
	// In real code: file, err := os.Open("file.txt")
	// defer file.Close()

	defer fmt.Println("Closing file...")
	fmt.Println("Processing file content")
	fmt.Println("File processing complete")
}

// processWithLock simulates acquiring and releasing a lock
func processWithLock() {
	fmt.Println("Acquiring lock...")
	// In real code: mutex.Lock()
	// defer mutex.Unlock()

	defer fmt.Println("Releasing lock...")
	fmt.Println("Critical section processing")
	fmt.Println("Critical section complete")
}

// ============================================
// NAMED RETURN VALUES WITH DEFER
// ============================================

func calculateWithDefer() (result int) {
	// result is a named return value
	defer func() {
		fmt.Println("Defer function running")
		result = result * 2 // Modify the return value
	}()

	result = 10
	fmt.Println("Before return, result =", result)
	return // Returns 20 (10 * 2)
}

// ============================================
// TIMING WITH DEFER
// ============================================

func measureExecutionTime() {
	start := time.Now()

	defer func() {
		duration := time.Since(start)
		fmt.Printf("Function took: %s\n", duration)
	}()

	fmt.Println("Starting work...")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Work completed")
}

// ============================================
// DEFER IN LOOPS
// ============================================

func deferInLoops() {
	fmt.Println("Processing items...")

	// Each iteration adds a deferred call
	for i := 0; i < 3; i++ {
		defer fmt.Printf("Cleanup for item %d\n", i)
		fmt.Printf("Processing item %d\n", i)
	}

	// All deferred calls run here in reverse order
	fmt.Println("All items processed")
}

// ============================================
// DEFER WITH ANONYMOUS FUNCTIONS
// ============================================

func demonstrateAnonymousDefer() {
	// You can defer anonymous functions
	defer func() {
		fmt.Println("Anonymous deferred function")
	}()

	fmt.Println("Main function body")
}
