# Channels Examples

Run this file from the project root:

```bash
go run channels/channels.go
```

## What To Practice

- Create channels with `make(chan T)`.
- Send values with `channel <- value`.
- Receive values with `<-channel`.
- Close a channel when no more values will be sent.
- Use `range` to receive until a channel is closed.
- Use `select` to wait on multiple channel operations.

## Useful Predefined Functions

```go
make(chan string) // Create an unbuffered string channel.
make(chan string, 3) // Create a buffered channel with capacity 3.
close(channel) // Signal that no more values will be sent.
len(channel) // Count queued values in a buffered channel.
cap(channel) // Check channel buffer capacity.
time.After(time.Second) // Create a timeout channel for select statements.
```

Use this folder when goroutines need to communicate safely.
