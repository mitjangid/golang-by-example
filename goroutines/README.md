# Goroutines Examples

Run this file from the project root:

```bash
go run goroutines/goroutines.go
```

## What To Practice

- Start concurrent work with `go functionName()`.
- Pass loop values into goroutines safely.
- Use `sync.WaitGroup` to wait for work to finish.
- Use `defer wg.Done()` so completion is reported even if the function grows later.

## Useful Predefined Functions

```go
go printTask("task", 1) // Start a function concurrently.
wg.Add(1) // Tell the WaitGroup one goroutine is starting.
wg.Done() // Tell the WaitGroup one goroutine is finished.
wg.Wait() // Block until all registered goroutines finish.
time.Sleep(50 * time.Millisecond) // Pause only for simple demos; prefer WaitGroup or channels.
```

Use this folder when you want to run independent work at the same time.
