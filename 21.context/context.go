package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// ============================================
	// BASIC CONTEXT
	// ============================================
	// context.Background() is the root context for all other contexts
	// context.TODO() is used when you're not sure which context to use
	// Both are empty contexts that are never cancelled

	fmt.Println("=== Basic Context ===")

	bgContext := context.Background()
	todoContext := context.TODO()

	fmt.Printf("Background context: %v\n", bgContext)
	fmt.Printf("TODO context: %v\n", todoContext)

	// ============================================
	// CONTEXT WITH TIMEOUT
	// ============================================
	// context.WithTimeout creates a context that automatically cancels
	// after a specified duration
	// Useful for operations that should not run indefinitely

	fmt.Println("\n=== Context With Timeout ===")

	// Create a context that cancels after 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Always call cancel to release resources

	// Simulate a long-running operation
	done := make(chan bool)
	go func() {
		worker(ctx, "Task with timeout")
		done <- true
	}()

	// Wait for the worker to finish or timeout
	select {
	case <-done:
		fmt.Println("Task completed successfully")
	case <-ctx.Done():
		fmt.Println("Task timed out:", ctx.Err())
	}

	// ============================================
	// CONTEXT WITH CANCEL
	// ============================================
	// context.WithCancel creates a context that can be manually cancelled
	// Useful for cancelling operations based on external conditions

	fmt.Println("\n=== Context With Cancel ===")

	ctx, cancel = context.WithCancel(context.Background())

	done = make(chan bool)
	go func() {
		worker(ctx, "Task with manual cancel")
		done <- true
	}()

	// Cancel the task after 1 second
	time.Sleep(1 * time.Second)
	cancel() // Manually cancel the context

	select {
	case <-done:
		fmt.Println("Task completed successfully")
	case <-ctx.Done():
		fmt.Println("Task was cancelled:", ctx.Err())
	}

	// ============================================
	// CONTEXT WITH DEADLINE
	// ============================================
	// context.WithDeadline creates a context that cancels at a specific time
	// Similar to WithTimeout but uses an absolute time instead of duration

	fmt.Println("\n=== Context With Deadline ===")

	// Set deadline to 1 second from now
	deadline := time.Now().Add(1 * time.Second)
	ctx, cancel = context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go func() {
		worker(ctx, "Task with deadline")
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Task completed successfully")
	case <-ctx.Done():
		fmt.Println("Task missed deadline:", ctx.Err())
	}

	// ============================================
	// CONTEXT WITH VALUE
	// ============================================
	// context.WithValue stores key-value pairs in the context
	// Useful for passing request-scoped data (user ID, trace ID, etc.)
	// Keys should be custom types to avoid collisions

	fmt.Println("\n=== Context With Value ===")

	// Define a custom type for context keys
	type contextKey string

	// Store values in context
	ctx = context.WithValue(context.Background(), contextKey("userID"), 12345)
	ctx = context.WithValue(ctx, contextKey("requestID"), "abc-123")

	// Retrieve values from context
	userID := ctx.Value(contextKey("userID"))
	requestID := ctx.Value(contextKey("requestID"))

	fmt.Printf("User ID: %v\n", userID)
	fmt.Printf("Request ID: %v\n", requestID)

	// Missing key returns nil
	missingValue := ctx.Value(contextKey("missing"))
	fmt.Printf("Missing key value: %v\n", missingValue)

	// ============================================
	// CHECKING CONTEXT IN OPERATIONS
	// ============================================
	// Long-running operations should periodically check ctx.Done()
	// This allows them to respond to cancellation quickly

	fmt.Println("\n=== Checking Context in Operations ===")

	ctx, cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	result := processWithCancellation(ctx)
	fmt.Printf("Processing result: %s\n", result)

	duration := time.Since(start)
	fmt.Printf("\nCode execution took: %s\n", duration)
}

// ============================================
// WORKER FUNCTION
// ============================================
// worker simulates a long-running task that respects context cancellation
func worker(ctx context.Context, taskName string) {
	fmt.Printf("Starting %s\n", taskName)

	// Simulate work with periodic context checks
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			// Context was cancelled, stop working
			fmt.Printf("%s was cancelled: %v\n", taskName, ctx.Err())
			return
		default:
			// Continue working
			fmt.Printf("%s working... step %d\n", taskName, i+1)
			time.Sleep(500 * time.Millisecond)
		}
	}

	fmt.Printf("%s completed\n", taskName)
}

// ============================================
// PROCESSING WITH CANCELLATION
// ============================================
// processWithCancellation shows how to check context during processing
func processWithCancellation(ctx context.Context) string {
	for i := 0; i < 10; i++ {
		// Check if context is cancelled
		select {
		case <-ctx.Done():
			return fmt.Sprintf("cancelled at step %d: %v", i+1, ctx.Err())
		default:
			// Continue processing
			time.Sleep(200 * time.Millisecond)
		}
	}
	return "completed successfully"
}
