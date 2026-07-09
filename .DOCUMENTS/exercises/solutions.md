# Solutions — Exercises

Solutions are intentionally concise. Try to solve exercises before reading.

1) MergeMaps
```go
func MergeMaps(a, b map[string]int) map[string]int {
  out := make(map[string]int)
  for k, v := range a { out[k] = v }
  for k, v := range b { out[k] += v }
  return out
}
```

2) Worker pool (outline)
```go
func StartWorkers(n int, jobs <-chan int, wg *sync.WaitGroup) {
  for i := 0; i < n; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      for job := range jobs { /* process job */ }
    }()
  }
}
```

3) Safe counter
```go
type Counter struct{ mu sync.Mutex; n int }
func (c *Counter) Inc() { c.mu.Lock(); c.n++; c.mu.Unlock() }
func (c *Counter) Value() int { c.mu.Lock(); defer c.mu.Unlock(); return c.n }
```

4) IndexOf
```go
func IndexOf[T comparable](s []T, v T) int {
  for i, x := range s { if x == v { return i } }
  return -1
}
```

5) Table-driven tests
- See cheatsheets/testing-and-benchmarks.md for a test template.
