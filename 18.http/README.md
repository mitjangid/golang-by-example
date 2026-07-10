# HTTP Examples

Run these files from the project root:

```bash
go run http/server.go
```

Then visit these URLs in your browser or use curl:
- http://localhost:8080/
- http://localhost:8080/hello
- http://localhost:8080/users
- http://localhost:8080/users/1

## What To Practice

- Create a basic HTTP server with `net/http`
- Define handler functions with specific signatures
- Handle different HTTP methods (GET, POST, etc.)
- Parse URL parameters and path variables
- Serve JSON responses
- Use middleware for common functionality

## Useful Predefined Functions

```go
http.HandleFunc("/", handler) // Register a handler for a path pattern
http.ListenAndServe(":8080", nil) // Start the server on a port
http.ResponseWriter // Interface for writing HTTP responses
http.Request // Represents an HTTP request
http.StatusNotFound // 404 status code constant
http.StatusInternalServerError // 500 status code constant
```

Use this folder when you want to practice building web servers and APIs.
