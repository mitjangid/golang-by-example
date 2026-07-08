# Go Quick Overview

A practical, beginner-friendly Go walkthrough created by **Amit K. Jangir**.

This repository is a compact learning map for Go fundamentals. Each folder focuses on one concept, with small runnable examples and comments written to explain the idea without overwhelming the reader.

## What You Will Learn

- Variables, constants, and basic printing
- Conditions, switch statements, and type switches
- Arrays, slices, maps, and range loops
- Functions, multiple returns, variadic functions, and closures
- Structs, embedded structs, methods, and pointer receivers
- Interfaces and behavior-based design
- Pointers and pass-by-value behavior
- Generics with typed functions and a generic stack
- Goroutines, channels, WaitGroups, and mutexes
- Enum-style typed constants with `iota`

## Project Structure

```text
.
|-- array/        # Arrays, slices, maps, and range
|-- channels/     # Channel basics, buffering, done signals, and select
|-- conditions/   # if/else and switch examples
|-- enums/        # Enum-style constants with iota
|-- functions/    # Functions, closures, variadic params, function values
|-- generics/     # Generic functions and generic types
|-- goroutines/   # Concurrent execution and WaitGroups
|-- interface/    # Interface-driven payment gateway example
|-- loops/        # for loops, break, and continue
|-- mutex/        # Race-condition-safe shared state
|-- pointer/      # Pointers and value updates
|-- struct/       # Structs, methods, embedding, constructors
|-- variables/    # Variables, constants, and simple expressions
`-- main.go       # Small root entry point
```

## How To Run

Install Go, then clone the repository and run any sample file directly:

```bash
go run variables/variables.go
go run array/slice.go
go run goroutines/goroutines.go
```

Each file is designed as a separate `package main` example. Run files one at a time instead of running an entire folder, because some folders contain multiple standalone examples.

## Recommended Go Version

Use **Go 1.21 or newer**.

Some examples use modern standard-library helpers such as `slices`, `maps`, and the `clear` built-in.

## Learning Path

1. Start with `variables/` and `conditions/`.
2. Move to `loops/`, `array/`, and `functions/`.
3. Learn `pointer/`, `struct/`, and `interface/`.
4. Finish with `generics/`, `goroutines/`, `channels/`, and `mutex/`.

## Why This Repository Exists

The goal is to help new Go learners get a quick but useful overview of the language. The examples are intentionally small, readable, and commented so readers can focus on one concept at a time.

## Author

Created and maintained by **Amit K. Jangir**.

If this repository helps you, consider starring it and sharing it with someone starting their Go journey.
