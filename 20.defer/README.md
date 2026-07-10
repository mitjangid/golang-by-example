# Defer Examples

Run these files from the project root:

```bash
go run defer/defer.go
```

## What To Practice

- Use `defer` to schedule function calls to run when a function exits
- Understand that deferred functions run in LIFO (last-in-first-out) order
- Use defer for resource cleanup (closing files, releasing locks)
- Use defer with named return values
- Understand how defer interacts with panic and recover
- Measure function execution time with defer

## Useful Predefined Functions

```go
defer function() // Schedules function to run when surrounding function exits
file.Close() // Close a file (commonly used with defer)
mutex.Unlock() // Release a mutex lock (commonly used with defer)
recover() // Recover from a panic (must be called in deferred function)
```

Use this folder when you want to practice defer, which is essential for clean resource management in Go.
