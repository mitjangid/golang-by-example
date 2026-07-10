# JSON Examples

Run these files from the project root:

```bash
go run json/json.go
```

## What To Practice

- Encode Go structs to JSON with `json.Marshal`
- Decode JSON to Go structs with `json.Unmarshal`
- Use struct tags to control JSON field names
- Handle nested structs and slices
- Work with JSON streams
- Handle JSON errors gracefully

## Useful Predefined Functions

```go
json.Marshal(data) // Convert Go data to JSON bytes
json.Unmarshal(jsonBytes, &data) // Convert JSON bytes to Go data
json.NewEncoder(writer).Encode(data) // Stream JSON to a writer
json.NewDecoder(reader).Decode(&data) // Stream JSON from a reader
```

Use this folder when you want to practice working with JSON data, which is essential for APIs and data storage.
