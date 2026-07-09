# Array, Slice, Map, And Range Examples

Run these files from the project root:

```bash
go run array/array.go
go run array/slice.go
go run array/map.go
go run array/range.go
```

## What To Practice

- Use arrays when the size is fixed.
- Use slices for lists that can grow with `append`.
- Use maps for key/value lookup.
- Use `range` to iterate over arrays, slices, maps, strings, and channels.

## Useful Predefined Functions

```go
len(numbers) // Count elements in an array, slice, or map.
cap(numbers) // Check the capacity of a slice before it grows.
append(numbers, 10) // Add values to a slice and return the updated slice.
copy(dst, src) // Copy slice elements into another slice.
delete(profile, "name") // Remove a key from a map.
clear(profile) // Remove all entries from a map or zero all elements in a slice.
slices.Equal(a, b) // Compare two slices from the standard slices package.
maps.Equal(a, b) // Compare two maps from the standard maps package.
```

Use this folder when you want to store many values and practice common collection operations.
