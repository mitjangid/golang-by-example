# Concurrency Cheatsheet — Goroutines, Channels, Select, WaitGroup, Mutex, Context

Short summary
- Goroutines run concurrently; use channels to communicate, WaitGroup to wait for completions, mutexes to protect shared state, and context for cancellation/timeouts.

Start goroutine
- `go func() { doWork() }()`

Channels
- Unbuffered: `ch := make(chan T)` // sender blocks until receiver ready
- Buffered: `ch := make(chan T, 3)` // capacity 3, sender blocks when full
- Close to signal end: `close(ch)`
- Receive loop: `for v := range ch { ... }` // ends when closed

Select
- Wait on multiple operations:
```
select {
case v := <-ch1:
  // handle v
case ch2 <- msg:
  // sent
case <-time.After(time.Second):
  // timeout
}
```

WaitGroup
- Pattern:
```
var wg sync.WaitGroup
wg.Add(n)
for ... { go func(){ defer wg.Done(); ... }() }
wg.Wait()
```

Mutex
- `var mu sync.Mutex; mu.Lock(); defer mu.Unlock()`
- For read-heavy use `sync.RWMutex`.

Context cancellation
- Use `context.WithCancel` / `WithTimeout` and check `select { case <-ctx.Done(): return }` inside goroutines.

Anti-patterns
- Using `time.Sleep` to wait for goroutines — prefer WaitGroup or channels.
- Leaving channels unclosed when receivers expect closure.
- Sharing memory without synchronization — run `go test -race` to detect races.

Try it
1) Worker pool: start 3 workers reading from jobs chan, send 10 jobs, close the channel, wait with WaitGroup.
2) Cancellation: start a long-running goroutine that returns when `ctx.Done()` is closed.
