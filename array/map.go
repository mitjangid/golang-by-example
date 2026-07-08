package main

import (
	"fmt"
	"maps"
	"time"
)

func main() {
	start := time.Now()

	// Maps store key/value pairs. They are similar to dictionaries or hash maps.
	languageDetails := make(map[string]string)
	languageDetails["name"] = "Go"
	languageDetails["type"] = "Programming language"
	fmt.Println(languageDetails)
	fmt.Println(languageDetails["type"])

	missingValue := languageDetails["missing"]
	fmt.Printf("Missing key value: %q\n", missingValue) // Missing keys return the zero value.

	// Map literals initialize maps with values.
	learnerProfile := map[string]string{"age": "26", "name": "Amit"}
	fmt.Println(learnerProfile)
	fmt.Println(len(learnerProfile))
	delete(learnerProfile, "name")
	fmt.Println(learnerProfile)

	clear(learnerProfile)
	fmt.Println(learnerProfile)

	// Use the comma-ok idiom to check whether a key exists.
	if _, exists := learnerProfile["name"]; exists {
		fmt.Println("Key exists")
	} else {
		fmt.Println("Key does not exist")
	}

	// The maps package provides helpers for map operations.
	fmt.Println(maps.Equal(languageDetails, learnerProfile))

	duration := time.Since(start)
	fmt.Printf("Code execution took: %s\n", duration)
}
