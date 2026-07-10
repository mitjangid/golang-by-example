package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	// ============================================
	// ENCODING: STRUCT TO JSON
	// ============================================
	// json.Marshal converts Go data to JSON bytes
	// Struct tags control how fields appear in JSON

	fmt.Println("=== Encoding: Struct to JSON ===")

	// Create a user struct
	user := User{
		ID:       1,
		Name:     "Alice Johnson",
		Email:    "alice@example.com",
		Age:      30,
		IsActive: true,
	}

	// Marshal the struct to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}

	fmt.Println("JSON output:")
	fmt.Println(string(jsonData))

	// ============================================
	// DECODING: JSON TO STRUCT
	// ============================================
	// json.Unmarshal converts JSON bytes to Go data
	// Pass a pointer to the struct to populate it

	fmt.Println("\n=== Decoding: JSON to Struct ===")

	jsonString := `{"id":2,"name":"Bob Smith","email":"bob@example.com","age":25,"is_active":false}`

	var decodedUser User
	err = json.Unmarshal([]byte(jsonString), &decodedUser)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}

	fmt.Printf("Decoded user: %+v\n", decodedUser)
	fmt.Printf("Name: %s, Age: %d\n", decodedUser.Name, decodedUser.Age)

	// ============================================
	// STRUCT TAGS
	// ============================================
	// Struct tags customize JSON field names and behavior
	// `json:"name"` sets the JSON field name
	// `json:"name,omitempty"` omits the field if it has zero value
	// `json:"-"` ignores the field completely

	fmt.Println("\n=== Struct Tags ===")

	product := Product{
		ID:          101,
		Name:        "Laptop",
		Description: "High-performance laptop",
		Price:       999.99,
		InternalID:  "INTERNAL-123", // This won't appear in JSON
	}

	productJSON, _ := json.Marshal(product)
	fmt.Println("Product JSON:")
	fmt.Println(string(productJSON))

	// ============================================
	// NESTED STRUCTS
	// ============================================
	// Structs can contain other structs
	// This represents nested JSON objects

	fmt.Println("\n=== Nested Structs ===")

	order := Order{
		ID: 1001,
		Customer: Customer{
			Name:  "Charlie",
			Email: "charlie@example.com",
		},
		Items: []Item{
			{Name: "Book", Price: 19.99},
			{Name: "Pen", Price: 2.50},
		},
		Total: 22.49,
	}

	orderJSON, _ := json.Marshal(order)
	fmt.Println("Order JSON:")
	fmt.Println(string(orderJSON))

	// ============================================
	// SLICES AND MAPS
	// ============================================
	// You can encode slices and maps to JSON arrays and objects

	fmt.Println("\n=== Slices and Maps ===")

	// Encode a slice
	numbers := []int{1, 2, 3, 4, 5}
	numbersJSON, _ := json.Marshal(numbers)
	fmt.Println("Slice JSON:", string(numbersJSON))

	// Encode a map
	personMap := map[string]interface{}{
		"name":   "David",
		"age":    35,
		"skills": []string{"Go", "Python", "JavaScript"},
	}
	mapJSON, _ := json.Marshal(personMap)
	fmt.Println("Map JSON:")
	fmt.Println(string(mapJSON))

	// ============================================
	// JSON STREAMS
	// ============================================
	// json.Encoder and json.Decoder work with streams
	// Useful for large JSON data or network operations

	fmt.Println("\n=== JSON Streams ===")

	// Encode to a writer (in this case, stdout)
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(map[string]string{"message": "Hello from encoder!"})

	// ============================================
	// HANDLING OPTIONAL FIELDS
	// ============================================
	// Use pointers for optional fields that might be missing

	fmt.Println("\n=== Optional Fields ===")

	// JSON with missing fields
	partialJSON := `{"id":3,"name":"Eve"}`
	var partialUser User
	json.Unmarshal([]byte(partialJSON), &partialUser)

	fmt.Printf("Partial user: %+v\n", partialUser)
	fmt.Printf("Email is zero value: %q\n", partialUser.Email)

	duration := time.Since(start)
	fmt.Printf("\nCode execution took: %s\n", duration)
}

// ============================================
// USER STRUCT WITH TAGS
// ============================================
// The json tag controls the JSON field name
// snake_case in JSON, camelCase in Go

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	IsActive bool   `json:"is_active"`
}

// ============================================
// PRODUCT STRUCT WITH OMITEMPTY
// ============================================
// omitempty omits fields with zero values

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price"`
	InternalID  string  `json:"-"` // This field is ignored
}

// ============================================
// NESTED STRUCTS
// ============================================

type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Item struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	ID       int      `json:"id"`
	Customer Customer `json:"customer"`
	Items    []Item   `json:"items"`
	Total    float64  `json:"total"`
}
