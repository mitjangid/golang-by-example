# Exercises — Short Challenges

This folder contains short exercises for learners. Try them before peeking at solutions in exercises/solutions.md.

1) Merge maps
- Implement `MergeMaps(map[string]int, map[string]int) map[string]int` that returns a combined map summing values for keys that appear in both.

2) Worker pool
- Implement a worker pool that starts N workers and processes M jobs from a channel. Use WaitGroup to wait for completion.

3) Safe counter
- Implement a struct with `Inc()` and `Value()` methods that is safe for concurrent use. Use `sync.Mutex`.

4) Generic IndexOf
- Implement `IndexOf[T comparable]([]T, T) int`.

5) Table-driven tests
- Write tests for MergeMaps and IndexOf using table-driven style.
