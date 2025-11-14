package main

import (
	"errors"
	"fmt"
)

type UserRegistrationError struct {
	Username string
}

func (e *UserRegistrationError) Error() string {
	return fmt.Sprintf("username '%s' already exists", e.Username)
}

var existingUsers = map[string]bool{
	"admin":    true,
	"gopher":   true,
	"testuser": true,
}

func registerUser(username string) error {
	if _, exists := existingUsers[username]; exists {
		return &UserRegistrationError{Username: username}
	}

	fmt.Printf("User '%s' registered successfully.\n", username)
	return nil
}

func main() {
	// --- Case 1: Attempt to register an existing user ---
	fmt.Println("Attempting to register 'gopher'...")
	err := registerUser("gopher")
	if err != nil {
		var registrationErr *UserRegistrationError
		if errors.As(err, &registrationErr) {
			fmt.Printf("Could not register user. Specific reason: %s\n", registrationErr)
		} else {
			fmt.Printf("An unexpected error occurred: %v\n", err)
		}
	}

	fmt.Println("\n--------------------\n")

	// --- Case 2: Attempt to register a new user ---
	fmt.Println("Attempting to register 'newdeveloper'...")
	err = registerUser("newdeveloper")
	if err != nil {
		var registrationErr *UserRegistrationError
		if errors.As(err, &registrationErr) {
			fmt.Printf("Could not register user. Specific reason: %s\n", registrationErr)
		} else {
			fmt.Printf("An unexpected error occurred: %v\n", err)
		}
	}
}
