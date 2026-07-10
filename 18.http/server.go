package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// ============================================
	// REGISTER HANDLERS
	// ============================================
	// http.HandleFunc registers a handler function for a specific path pattern
	// The handler function must have the signature: func(w http.ResponseWriter, r *http.Request)

	// Handle the root path
	http.HandleFunc("/", homeHandler)

	// Handle a simple greeting
	http.HandleFunc("/hello", helloHandler)

	// Handle user-related endpoints
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/", userByIDHandler)

	// ============================================
	// START THE SERVER
	// ============================================
	// ListenAndServe starts the HTTP server
	// First argument is the address (port)
	// Second argument is the handler (nil uses DefaultServeMux)

	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	fmt.Println("Try these URLs:")
	fmt.Println("  http://localhost:8080/")
	fmt.Println("  http://localhost:8080/hello")
	fmt.Println("  http://localhost:8080/users")
	fmt.Println("  http://localhost:8080/users/1")
	fmt.Println("\nPress Ctrl+C to stop the server")

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}

// ============================================
// BASIC HANDLER
// ============================================
// homeHandler handles requests to the root path "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Write the HTTP status code
	w.WriteHeader(http.StatusOK)

	// Write the response body
	fmt.Fprintf(w, "Welcome to the Go HTTP Server!\n")
	fmt.Fprintf(w, "Current time: %s\n", time.Now().Format(time.RFC3339))
}

// ============================================
// SIMPLE RESPONSE HANDLER
// ============================================
// helloHandler returns a simple greeting
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check the HTTP method
	if r.Method != http.MethodGet {
		// Return 405 Method Not Allowed for non-GET requests
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Write a simple response
	fmt.Fprintf(w, "Hello, World!\n")
	fmt.Fprintf(w, "You accessed: %s\n", r.URL.Path)
}

// ============================================
// JSON RESPONSE HANDLER
// ============================================
// User represents a user in our system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// usersHandler returns a list of users as JSON
func usersHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")

	// Create sample users
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com"},
	}

	// Encode the users slice to JSON
	// json.NewEncoder creates an encoder that writes to the response writer
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		// If encoding fails, return an error
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
		return
	}
}

// ============================================
// PATH PARAMETER HANDLER
// ============================================
// userByIDHandler handles requests like /users/1, /users/2, etc.
func userByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the URL path
	// r.URL.Path gives the full path like "/users/1"
	// We need to remove the "/users/" prefix to get the ID

	path := r.URL.Path
	// Check if the path starts with "/users/"
	if len(path) <= len("/users/") {
		http.Error(w, "User ID required", http.StatusBadRequest)
		return
	}

	userID := path[len("/users/"):]

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Create a sample user based on the ID
	user := User{
		ID:    1,
		Name:  "User " + userID,
		Email: "user" + userID + "@example.com",
	}

	// Encode and return the user
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Failed to encode user", http.StatusInternalServerError)
		return
	}
}
