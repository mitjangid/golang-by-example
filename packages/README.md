# Packages Examples

Run this example from the project root:

```bash
go run packages/main.go
```

## What To Practice

- Import local packages with the module path from `go.mod`.
- Export names by starting them with a capital letter.
- Keep lowercase names private to their package.
- Use third-party packages through Go modules.
- Run `go mod tidy` after changing imports.

## Useful Predefined Commands And Functions

```go
fmt.Println(value) // Print values returned from another package.
color.Cyan("text") // Call a third-party package function.
time.Now() // Create timestamp values inside package code.
```

```bash
go mod tidy # Add missing module requirements and remove unused ones.
go get github.com/fatih/color # Add a third-party dependency.
```

Use this folder when you want to understand how a real Go project is split into packages.
