# Generics Examples

Run this file from the project root:

```bash
go run generics/generics.go
```

## What To Practice

- Use type parameters such as `[T any]`.
- Restrict accepted types with constraints such as `int | string`.
- Use `comparable` when values must support `==` and `!=`.
- Create generic types such as `stack[T]`.
- Return a zero value when a generic container is empty.

## Useful Predefined Functions

```go
fmt.Println(item) // Print values without caring about their concrete type.
len(items) // Check how many values are in a generic slice.
append(items, item) // Add a typed value to a generic slice.
```

Use this folder when you want one implementation to work with several types safely.
