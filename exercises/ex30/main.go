package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Exercise 30: JSON Marshalling and Unmarshalling
// Convert a Go struct to a JSON string (marshalling) and back (unmarshalling).

// User represents a user with a name, age, and a list of programming languages.
type User struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Languages []string `json:"languages"`
}

func main() {
	// --- Part 1: Marshalling (Go Struct -> JSON) ---

	fmt.Println("--- Marshalling ---")

	// Create an instance of the User struct.
	user1 := User{
		Name:      "Gopher",
		Age:       15,
		Languages: []string{"Go", "Python", "JavaScript"},
	}

	// The json.Marshal function converts a Go data structure into a JSON byte slice.
	// This process is called "marshalling".
	jsonData, err := json.Marshal(user1)
	if err != nil {
		// If an error occurs during marshalling, we log it and exit.
		log.Fatalf("Error marshalling JSON: %s", err)
	}

	// Print the resulting JSON. We convert the byte slice to a string for readability.
	fmt.Println("Go Struct:", user1)
	fmt.Println("JSON Output:", string(jsonData))

	fmt.Println("\n--------------------\n")

	// --- Part 2: Unmarshalling (JSON -> Go Struct) ---

	fmt.Println("--- Unmarshalling ---")

	// This is our JSON data, represented as a raw string literal.
	jsonString := `{"name":"Ada Lovelace","age":36,"languages":["Assembly","C++","Lisp"]}`

	// Create a new User variable to store the unmarshalled data.
	var user2 User

	// The json.Unmarshal function parses a JSON byte slice and populates the
	// fields of the struct pointed to by the second argument.
	// This process is called "unmarshalling".
	err = json.Unmarshal([]byte(jsonString), &user2)
	if err != nil {
		// If the JSON is malformed or doesn't match the struct, an error is returned.
		log.Fatalf("Error unmarshalling JSON: %s", err)
	}

	// Print the resulting struct to verify that the data was parsed correctly.
	fmt.Println("JSON Input:", jsonString)
	fmt.Println("Go Struct:", user2)
	fmt.Printf("Parsed Name: %s\n", user2.Name)
	fmt.Printf("Parsed Languages: %v\n", user2.Languages)
}
