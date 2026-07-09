# Variables Examples

Run these files from the project root:

```bash
go run variables/variables.go
go run variables/constants.go
```

## What To Practice

- Use `fmt.Println` to print values with a newline.
- Use `fmt.Printf` when you want formatted output.
- Use `var` when you want to show the type clearly.
- Use `:=` inside functions when Go can infer the type.
- Use `const` for values that should not change.

## Useful Predefined Functions

```go
fmt.Println("Hello", "Go") // Print values separated by spaces and add a newline.
fmt.Printf("Age: %d\n", 37) // Print formatted text; %d is for integers.
fmt.Sprintf("Name: %s", "Amit") // Build a formatted string without printing it.
len("Golang") // Count bytes in a string; useful before slicing or validating input.
```

Use this folder when you want to practice basic values, types, constants, and simple output.
