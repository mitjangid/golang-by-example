# User Package Example

This folder is used by `packages/main.go`.

Run the package demo from the project root:

```bash
go run packages/main.go
```

## What To Practice

- Export a struct type with a capitalized name.
- Export fields that should be accessible from another package.
- Keep package data simple when the folder is only responsible for one model.

## Useful Predefined Functions

```go
fmt.Println(currentUser) // Print the struct from the importing package.
fmt.Printf("%+v\n", currentUser) // Print exported field names and values.
```

Use this folder to practice defining reusable data types for other packages.
