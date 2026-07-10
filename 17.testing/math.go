package testing

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference between two integers
func Subtract(a, b int) int {
	return a - b
}

// Divide returns the result of dividing two integers
// Returns an error if dividing by zero
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

// ErrDivisionByZero is returned when attempting to divide by zero
var ErrDivisionByZero = Error("division by zero")

// Error is a custom error type for testing
type Error string

func (e Error) Error() string {
	return string(e)
}
