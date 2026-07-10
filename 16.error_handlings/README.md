# Error Handling Examples

Run these files from the project root:

```bash
go run error_handlings/basic_error_handling.go
go run error_handlings/custom_errors.go
go run error_handlings/panic_recover.go
```

## What To Practice

- Check if errors are `nil` before proceeding with operations.
- Use multiple return values to handle errors gracefully.
- Create custom error types with `error` interface implementation.
- Wrap errors with context using `fmt.Errorf` and `%w`.
- Use `panic` only for unrecoverable conditions, not normal error handling.
- Use `recover` to handle panics in deferred functions.
- Distinguish between expected errors and exceptional conditions.

## Useful Predefined Functions

```go
errors.New("something went wrong") // Create a simple error with a message.
fmt.Errorf("context: %w", err) // Wrap an error with additional context.
errors.Is(err, targetErr) // Check if an error matches a specific error.
errors.As(err, &target) // Check if an error can be cast到 a specific type.
log.Fatal(err) // Print error and exit the program (use in main only).
log.Panic(err) // Print error and panic (use only in exceptional cases).
```

Use this folder when you want to practice Go's explicit error handling approach and understand how to handle failures gracefully.
