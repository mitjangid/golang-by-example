# Mutex Examples

Run this file from the project root:

```bash
go run mutex/mutex.go
```

## What To Practice

- Protect shared data with `sync.Mutex`.
- Lock before reading or writing shared state.
- Unlock with `defer` so the lock is released reliably.
- Combine a mutex with `sync.WaitGroup` when many goroutines update one value.

## Useful Predefined Functions

```go
mu.Lock() // Enter the protected section.
mu.Unlock() // Leave the protected section.
wg.Add(1) // Register a goroutine.
wg.Done() // Mark one goroutine as finished.
wg.Wait() // Wait until all goroutines finish.
```

Use this folder when many goroutines share and update the same data.
