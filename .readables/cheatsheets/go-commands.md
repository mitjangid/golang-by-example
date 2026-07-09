# Go CLI Commands — Quick Reference

When to use this
- Short reference for common `go` commands used while learning and experimenting.

Run & build
- go run <file.go>
  - Run a single file or package main: `go run variables/variables.go`
- go build
  - Compile the package in the current directory: `go build ./...`

Module & dependencies
- go mod init <module/path>
  - Initialize a module (already present in this repo)
- go get <module>@<ver>
  - Add or upgrade a dependency: `go get github.com/fatih/color@latest`
- go mod tidy
  - Add missing and remove unused dependencies

Testing & vet
- go test ./...
  - Run all tests in the module
- go test -run TestName ./pkg
  - Run a single test or pattern
- go test -bench . -benchmem
  - Run benchmarks
- go test -cover
  - Run tests and show coverage
- go vet ./...
  - Static checks; helpful to find suspicious code

Formatting & linting
- go fmt ./...
  - Format Go code in-place
- go list -deps
  - List package dependencies

Debug & profiling
- dlv debug
  - Start Delve debugger for package main (install Delve separately)
- go test -run TestX -cpuprofile cpu.out
  - Capture CPU profile from tests
- go tool pprof cpu.out
  - Inspect profile (web, top, list)

Environment & info
- go env
  - Show Go environment
- go version
  - Show installed go version

Short examples
- Run a folder example: `go run packages/main.go`
- Clean dependencies: `go mod tidy`