# Pointer Examples

Run this file from the project root:

```bash
go run pointer/pointer.go
```

## What To Practice

- Use `&value` to get a memory address.
- Use `*pointer` to read or update the value at that address.
- Compare pass-by-value with passing a pointer.
- Use pointers when a function must update the original value.

## Useful Predefined Functions

```go
fmt.Println(&number) // Print a value's address for learning.
fmt.Printf("%p\n", &number) // Print an address with pointer formatting.
new(int) // Allocate a zero-value int and return *int.
```

Use this folder when you want to understand when a function changes a copy and when it changes the original.
