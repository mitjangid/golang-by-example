# Auth Package Example

This folder is used by `packages/main.go`.

Run the package demo from the project root:

```bash
go run packages/main.go
```

## What To Practice

- Keep helper functions private with lowercase names, such as `loginWithCredentials`.
- Export functions with capitalized names, such as `LoginWithCredentials`.
- Split one package across multiple `.go` files.
- Return package-owned types from constructor-style functions.

## Useful Predefined Functions

```go
fmt.Println("Login request") // Print package behavior while learning.
len(password) // Count characters or bytes for simple validation demos.
time.Now() // Store the creation time for a session value.
```

Use this folder to understand package visibility and how files in the same package work together.
