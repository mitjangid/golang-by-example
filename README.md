# Go Quick Overview

A practical, beginner-friendly Go walkthrough created by **Amit K. Jangir**.

This repository is a compact learning map for Go fundamentals. Each folder focuses on one concept, with small runnable examples and comments written to explain the idea without overwhelming the reader.

Each concept folder also has its own `README.md` with the files to run, the main ideas to practice, and useful predefined functions or commands for that topic.

## What You Will Learn

- Variables, constants, and basic printing
- Conditions, switch statements, and type switches
- Arrays, slices, maps, and range loops
- Functions, multiple returns, variadic functions, and closures
- Structs, embedded structs, methods, and pointer receivers
- Interfaces and behavior-based design
- Pointers and pass-by-value behavior
- File and folder handling with the `os`, `io`, and `path/filepath` packages
- Generics with typed functions and a generic stack
- Goroutines, channels, WaitGroups, and mutexes
- Enum-style typed constants with `iota`
- Local packages, modules, `go.mod`, `go.sum`, and third-party dependencies
- Error handling, custom errors, panic, and recover
- Testing with table-driven tests and benchmarks
- HTTP servers and handlers
- JSON encoding and decoding
- Context for cancellations, timeouts, and request-scoped values
- Defer for resource cleanup and timing

## Project Structure

```text
.
|-- array/            # Arrays, slices, maps, and range
|-- channels/         # Channel basics, buffering, done signals, and select
|-- conditions/       # if/else and switch examples
|-- context/          # Context for cancellations, timeouts, and request-scoped values
|-- defer/            # Deferred function execution for cleanup and timing
|-- enums/            # Enum-style constants with iota
|-- error_handlings/  # Error handling, custom errors, panic, and recover
|-- files/            # File reading, writing, renaming, deleting, and folders
|-- functions/        # Functions, closures, variadic params, function values
|-- generics/         # Generic functions and generic types
|-- goroutines/       # Concurrent execution and WaitGroups
|-- http/             # HTTP servers and handlers
|-- interface/        # Interface-driven payment gateway example
|-- json/             # JSON encoding and decoding
|-- loops/            # for loops, break, and continue
|-- mutex/            # Race-condition-safe shared state
|-- packages/         # Local packages, exported names, and dependencies
|-- pointer/          # Pointers and value updates
|-- struct/           # Structs, methods, embedding, constructors
|-- testing/          # Unit testing, table-driven tests, and benchmarks
|-- variables/        # Variables, constants, and simple expressions
|-- go.mod            # Module name and dependency requirements
|-- go.sum            # Dependency checksums
`-- main.go           # Small root entry point
```

## How To Run

Install Go, then clone the repository and run any sample file directly:

```bash
go run variables/variables.go
go run array/slice.go
go run goroutines/goroutines.go
go run packages/main.go
go run error_handlings/basic_error_handling.go
go run http/server.go
go run json/json.go
go run context/context.go
go run defer/defer.go
```

Most concept files are standalone `package main` examples. Run files one at a time instead of running an entire folder, because some folders contain multiple standalone examples.

For quick notes about a topic, open the `README.md` inside that folder before running its examples.

The `packages/` example imports local packages and the third-party `github.com/fatih/color` package, so run it from the repository root:

```bash
go run packages/main.go
```

The `files/` example can be run from the root or from inside the `files` folder:

```bash
go run files/files.go
cd files && go run files.go
```

The `testing/` folder contains test files that are run with `go test`:

```bash
go test -v ./testing/
go test -bench=. ./testing/
```

The `http/` example starts a web server that you can access in your browser:

```bash
go run http/server.go
# Then visit http://localhost:8080/
```

## Recommended Go Version

Use **Go 1.21 or newer**.

Some examples use modern standard-library helpers such as `slices`, `maps`, and the `clear` built-in.

## Learning Path

1. Start with `variables/` and `conditions/`.
2. Move to `loops/`, `array/`, and `functions/`.
3. Learn `pointer/`, `struct/`, `interface/`, and `files/`.
4. Explore `packages/` to understand modules, imports, exported names, and dependencies.
5. Learn `error_handlings/` for Go's explicit error handling approach.
6. Practice `defer/` for resource cleanup and timing.
7. Explore `context/` for cancellations, timeouts, and request-scoped values.
8. Learn `json/` for encoding and decoding JSON data.
9. Build web servers with `http/`.
10. Write tests with `testing/`.
11. Finish with `generics/`, `goroutines/`, `channels/`, and `mutex/`.

## Module And Dependencies

This project uses Go modules:

- `go.mod` defines the module path and required dependencies.
- `go.sum` stores checksums for downloaded dependencies.
- `go mod tidy` cleans unused dependencies and adds missing ones after import changes.

The `packages/` example demonstrates both local imports and a small third-party package.

## Why This Repository Exists

The goal is to help new Go learners get a quick but useful overview of the language. The examples are intentionally small, readable, and commented so readers can focus on one concept at a time.

## Author

Created and maintained by **Amit K. Jangir**.

If this repository helps you, consider starring it and sharing it with someone starting their Go journey.
