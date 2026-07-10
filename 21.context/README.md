# Context Examples

Run these files from the project root:

```bash
go run context/context.go
```

## What To Practice

- Create contexts with `context.Background()` and `context.TODO()`
- Set timeouts and deadlines with `context.WithTimeout` and `context.WithDeadline`
- Cancel operations with `context.WithCancel`
- Pass request-scoped values with `context.WithValue`
- Check for context cancellation in long-running operations
- Handle context errors gracefully

## Useful Predefined Functions

```go
context.Background() // Creates an empty context (root context)
context.TODO() // Creates a context when unsure which to use
context.WithTimeout(parent, duration) // Creates a context that cancels after duration
context.WithCancel(parent) // Creates a context that can be manually cancelled
context.WithValue(parent, key, value) // Creates a context with a key-value pair
ctx.Done() // Returns a channel that closes when context is cancelled
ctx.Err() // Returns the error that caused context cancellation
```

Use this folder when you want to practice managing cancellations, timeouts, and request-scoped values.
