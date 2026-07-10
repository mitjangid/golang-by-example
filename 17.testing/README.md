# Testing Examples

Run these files from the project root:

```bash
# Run the tests
go test ./testing/...

# Run with verbose output
go test -v ./testing/

# Run specific test
go test -v ./testing/ -run TestAdd

# Run with coverage
go test -cover ./testing/
```

## What To Practice

- Write basic unit tests with the `testing` package
- Use table-driven tests for multiple test cases
- Compare expected vs actual values with `assert` patterns
- Test error conditions
- Use setup/teardown with `TestMain`
- Run benchmarks to measure performance

## Useful Predefined Functions

```go
t.Error("message") // Report a test failure but continue
t.Fatal("message") // Report a test failure and stop the test
t.Logf("message") // Log information during test execution
t.Helper() // Mark function as a test helper (excludes from line numbers)
testing.Benchmark(func(b *testing.B) { ... }) // Run a benchmark
```

Use this folder when you want to practice writing tests, which is essential for Go development.
