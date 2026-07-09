# Debugging & Tools — Quick Guide

Formatting & basic checks
- `gofmt` / `go fmt ./...` to format code
- `go vet ./...` to catch suspicious constructs

Lint
- `golangci-lint run` (if installed) for a battery of linters

Delve (debugger)
- Install: https://github.com/go-delve/delve
- Start debugging: `dlv debug` (for package main)
- Attach to running process or run tests with `dlv test`.

Profiling (pprof)
- CPU: `go test -cpuprofile cpu.out` or instrument main
- Memory: `go test -memprofile mem.out`
- View: `go tool pprof -http=:8080 cpu.out`

Race detector
- `go test -race ./...` to find data races in tests

Common workflows
- Reproduce failing test locally: `go test -run TestName -v ./pkg`
- Quick prints: `fmt.Printf("%+v\n", x)` for structs

Try it
- Run `go test -race ./...` on packages that use goroutines (goroutines/, mutex/, channels/).
