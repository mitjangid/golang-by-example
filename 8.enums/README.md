# Enum-Style Constants Examples

Run this file from the project root:

```bash
go run enums/enum.go
```

## What To Practice

- Create a named type for a limited set of values.
- Use `const` blocks to define allowed values.
- Use `iota` for increasing integer constants.
- Use string constants when printed values should be readable.

## Useful Predefined Functions

```go
fmt.Println(status) // Print enum-style values.
fmt.Printf("%T\n", status) // Print the named type, such as main.OrderStatus.
```

Use this folder when a value should come from a small known list of options.
