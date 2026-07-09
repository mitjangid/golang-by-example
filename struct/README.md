# Struct Examples

Run this file from the project root:

```bash
go run struct/struct.go
```

## What To Practice

- Group related data in a struct.
- Create values with struct literals.
- Add methods with value and pointer receivers.
- Embed one struct inside another.
- Use constructor-style functions such as `newOrder`.

## Useful Predefined Functions

```go
fmt.Println(orderValue) // Print a struct with default formatting.
fmt.Printf("%+v\n", orderValue) // Print field names and values.
new(order) // Allocate a zero-value order and return *order.
time.Now() // Store creation or update timestamps in struct fields.
```

Use this folder when you want to model real-world data and attach behavior to it.
