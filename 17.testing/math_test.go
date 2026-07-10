package testing

import (
	"testing"
)

// ============================================
// BASIC TEST
// ============================================
// A basic test function starts with "Test" and takes *testing.T
// Run with: go test -v -run TestAdd

func TestAdd(t *testing.T) {
	// ============================================
	// SIMPLE TEST CASE
	// ============================================
	// Test the basic functionality
	result := Add(2, 3)
	expected := 5

	// Compare result with expected
	if result != expected {
		// t.Errorf reports a failure but continues the test
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

// ============================================
// TABLE-DRIVEN TESTS
// ============================================
// Table-driven tests are a Go best practice
// They let you test many cases with one test function
// Run with: go test -v -run TestAddTableDriven

func TestAddTableDriven(t *testing.T) {
	// Define test cases as a slice of structs
	tests := []struct {
		name     string   // Name of the test case
		a        int      // First input
		b        int      // Second input
		expected int      // Expected result
	}{
		{
			name:     "positive numbers",
			a:        2,
			b:        3,
			expected: 5,
		},
		{
			name:     "zero",
			a:        0,
			b:        5,
			expected: 5,
		},
		{
			name:     "negative numbers",
			a:        -2,
			b:        -3,
			expected: -5,
		},
		{
			name:     "mixed positive and negative",
			a:        10,
			b:        -3,
			expected: 7,
		},
	}

	// Run each test case
	for _, tt := range tests {
		// t.Run creates a subtest
		// This makes each case show up separately in test output
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// ============================================
// TESTING ERRORS
// ============================================
// Test functions that return errors
// Run with: go test -v -run TestDivide

func TestDivide(t *testing.T) {
	// Test successful division
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) returned error: %v", err)
	}
	if result != 5 {
		t.Errorf("Divide(10, 2) = %d; expected 5", result)
	}

	// Test division by zero error
	result, err = Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) expected error, got nil")
	}
	if err != ErrDivisionByZero {
		t.Errorf("Divide(10, 0) returned wrong error: %v", err)
	}
	if result != 0 {
		t.Errorf("Divide(10, 0) = %d; expected 0", result)
	}
}

// ============================================
// TABLE-DRIVEN TEST WITH ERRORS
// ============================================
// Combine table-driven tests with error testing
// Run with: go test -v -run TestDivideTableDriven

func TestDivideTableDriven(t *testing.T) {
	tests := []struct {
		name        string
		a           int
		b           int
		expected    int
		expectError bool
	}{
		{
			name:        "successful division",
			a:           10,
			b:           2,
			expected:    5,
			expectError: false,
		},
		{
			name:        "division by zero",
			a:           10,
			b:           0,
			expected:    0,
			expectError: true,
		},
		{
			name:        "negative result",
			a:           -10,
			b:           2,
			expected:    -5,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)

			if tt.expectError {
				if err == nil {
					t.Error("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("Divide(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
				}
			}
		})
	}
}

// ============================================
// TEST HELPERS
// ============================================
// Helper functions make tests more readable
// Mark them with t.Helper() to exclude from line numbers in errors
// Run with: go test -v -run TestSubtractWithHelper

func assertEqual(t *testing.T, got, expected int, msg string) {
	t.Helper() // This marks it as a helper function
	if got != expected {
		t.Errorf("%s: got %d, expected %d", msg, got, expected)
	}
}

func TestSubtractWithHelper(t *testing.T) {
	assertEqual(t, Subtract(10, 3), 7, "Subtract(10, 3)")
	assertEqual(t, Subtract(5, 5), 0, "Subtract(5, 5)")
	assertEqual(t, Subtract(0, 5), -5, "Subtract(0, 5)")
}
