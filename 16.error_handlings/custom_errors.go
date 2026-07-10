package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// ============================================
	// CUSTOM ERROR TYPES
	// ============================================
	// Sometimes you need more than just an error message
	// You can create your own error types with extra information
	// Just implement the Error() method to satisfy the error interface

	fmt.Println("=== Custom Error Types ===")

	// Successful withdrawal
	err := withdraw(100, 50)
	if err != nil {
		fmt.Println("Withdrawal error:", err)
	} else {
		fmt.Println("Withdrawal successful")
	}

	// Failed withdrawal - insufficient funds
	err = withdraw(50, 100)
	if err != nil {
		fmt.Println("Withdrawal error:", err)
		// Type assertion: check if error is our custom type
		// This lets us access the extra fields (Balance, Requested)
		if insErr, ok := err.(*InsufficientFundsError); ok {
			fmt.Printf("  Current balance: $%.2f\n", insErr.Balance)
			fmt.Printf("  Amount requested: $%.2f\n", insErr.Requested)
		}
	}

	// ============================================
	// ERROR TYPE CHECKING WITH errors.As
	// ============================================
	// errors.As checks if an error can be converted to a specific type
	// This is safer than manual type assertion

	fmt.Println("\n=== Error Type Checking ===")

	err = processPayment(150)
	var paymentErr *PaymentError
	if errors.As(err, &paymentErr) {
		fmt.Println("Payment error detected:", paymentErr)
		fmt.Println("  Error code:", paymentErr.Code)
		fmt.Println("  Can retry:", paymentErr.Retryable)
	}

	// ============================================
	// SENTINEL ERRORS
	// ============================================
	// Sentinel errors are predefined error values
	// Useful for comparing errors with errors.Is()
	// Common in standard library (e.g., io.EOF, os.ErrNotExist)

	fmt.Println("\n=== Sentinel Errors ===")

	// Try to connect to an invalid server
	err = connectToServer("invalid-server")
	if errors.Is(err, ErrConnectionFailed) {
		fmt.Println("Connection failed - server is unreachable")
	}

	// Try to validate a user with invalid age
	err = validateUser(-5)
	if errors.Is(err, ErrInvalidAge) {
		fmt.Println("Validation failed - age must be between 0 and 120")
	}

	// ============================================
	// ERROR WRAPPING WITH CUSTOM TYPES
	// ============================================
	// You can wrap errors with custom error types
	// Implement Unwrap() to allow error unwrapping

	fmt.Println("\n=== Error Wrapping with Custom Types ===")

	err = performOperation("fail")
	if err != nil {
		fmt.Println("Operation error:", err)
		// Extract the custom error type
		var opErr *OperationError
		if errors.As(err, &opErr) {
			fmt.Println("  Operation name:", opErr.Operation)
			fmt.Println("  Error code:", opErr.Code)
			fmt.Println("  Underlying error:", opErr.Err)
		}
	}

	duration := time.Since(start)
	fmt.Printf("\nCode execution took: %s\n", duration)
}

// ============================================
// CUSTOM ERROR TYPE: InsufficientFundsError
// ============================================
// This error type holds extra information about the failed transaction
type InsufficientFundsError struct {
	Balance   float64 // Current account balance
	Requested float64 // Amount the user tried to withdraw
}

// Error is required to implement the error interface
// This method returns the error message as a string
func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("insufficient funds: balance $%.2f, requested $%.2f", e.Balance, e.Requested)
}

// withdraw attempts to withdraw money from an account
// Returns nil on success, or InsufficientFundsError on failure
func withdraw(balance, amount float64) error {
	if amount > balance {
		// Return our custom error type with the relevant details
		return &InsufficientFundsError{
			Balance:   balance,
			Requested: amount,
		}
	}
	fmt.Printf("Withdrew $%.2f, new balance: $%.2f\n", amount, balance-amount)
	return nil
}

// ============================================
// CUSTOM ERROR TYPE: PaymentError
// ============================================
// This error includes a code and retry information
type PaymentError struct {
	Message   string // Human-readable error message
	Code      string // Machine-readable error code
	Retryable bool   // Whether the operation can be retried
}

func (e *PaymentError) Error() string {
	return fmt.Sprintf("payment error [%s]: %s (retryable: %v)", e.Code, e.Message, e.Retryable)
}

// processPayment processes a payment transaction
// Returns nil on success, or PaymentError if amount exceeds limit
func processPayment(amount float64) error {
	if amount > 100 {
		return &PaymentError{
			Message:   "amount exceeds limit",
			Code:      "LIMIT_EXCEEDED",
			Retryable: false,
		}
	}
	fmt.Printf("Payment of $%.2f processed successfully\n", amount)
	return nil
}

// ============================================
// SENTINEL ERRORS
// ============================================
// These are predefined error values that can be compared
// They're useful when you need to check for specific error conditions
var (
	ErrConnectionFailed = errors.New("connection failed")
	ErrInvalidAge       = errors.New("invalid age")
)

// connectToServer simulates connecting to a server
// Returns wrapped sentinel error if server is invalid
func connectToServer(server string) error {
	if server == "invalid-server" {
		// Wrap the sentinel error with additional context
		return fmt.Errorf("server '%s' unreachable: %w", server, ErrConnectionFailed)
	}
	fmt.Printf("Connected to %s\n", server)
	return nil
}

// validateUser checks if user age is valid
// Returns sentinel error if age is out of range
func validateUser(age int) error {
	if age < 0 || age > 120 {
		return ErrInvalidAge
	}
	fmt.Printf("User age %d is valid\n", age)
	return nil
}

// ============================================
// CUSTOM ERROR TYPE WITH UNWRAP: OperationError
// ============================================
// This error wraps another error and adds operation context
type OperationError struct {
	Operation string // Name of the operation that failed
	Code      int    // HTTP-style error code
	Err       error  // The underlying error that caused this
}

func (e *OperationError) Error() string {
	return fmt.Sprintf("operation '%s' failed with code %d: %v", e.Operation, e.Code, e.Err)
}

// Unwrap returns the underlying error
// This allows errors.Is() and errors.As() to work with wrapped errors
func (e *OperationError) Unwrap() error {
	return e.Err
}

// performOperation executes an operation
// Returns OperationError wrapping the underlying error on failure
func performOperation(input string) error {
	if input == "fail" {
		return &OperationError{
			Operation: "process",
			Code:      500,
			Err:       errors.New("internal processing error"),
		}
	}
	fmt.Printf("Operation '%s' completed\n", input)
	return nil
}
