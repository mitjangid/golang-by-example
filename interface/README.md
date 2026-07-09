# Interface Examples

Run this file from the project root:

```bash
go run interface/interface.go
```

## What To Practice

- Define behavior with an interface.
- Let concrete types satisfy an interface by implementing its methods.
- Depend on behavior instead of a specific type.
- Swap implementations without changing the caller.

## Useful Predefined Functions

```go
fmt.Println(value) // Print which implementation is running.
fmt.Printf("%T\n", gateway) // Show the concrete type stored behind an interface.
```

Use this folder when you want flexible code that can work with multiple implementations.
