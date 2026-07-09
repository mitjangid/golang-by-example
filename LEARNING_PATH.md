# Learning Path — One Page

Goal: a 1-hour quick tour, or a multi-day deep dive for beginners.

Quick 1-hour run (linear)
1. `go run variables/variables.go` (5 min)
2. `go run conditions/conditions.go` (5 min)
3. `go run loops/loops.go` (5 min)
4. `go run array/slice.go` (10 min)
5. `go run functions/functions.go` (10 min)
6. `go run pointer/pointer.go` (5 min)
7. `go run struct/struct.go` (5 min)
8. `go run interface/interface.go` (10 min)

Deep dive (multi-day)
Day 1: variables, conditions, loops, arrays/slices
Day 2: functions, pointers, structs, interfaces
Day 3: packages, files, error handling
Day 4: generics, goroutines, channels, mutex
Day 5: testing, profiling, tooling

Exercises
- Try the 'Try it' exercises in each cheatsheet. After that, write a small program that reads a file, processes lines into a map, and serves results over `net/http`.

Tips
- Run `go test ./...` occasionally to exercise test files.
- Use `go mod tidy` after changing imports.
