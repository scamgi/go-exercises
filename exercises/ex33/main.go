package main

import (
	"fmt"
	"reflect"
)

// inspect examines a variable `v` of any type at runtime and prints information about its type and value.
func inspect(v any) {
	fmt.Printf("--- Inspecting: %+v ---\n", v)

	// reflect.TypeOf returns the static type of the variable.
	t := reflect.TypeOf(v)
	// reflect.ValueOf returns the dynamic value of the variable.
	val := reflect.ValueOf(v)

	// The 'Kind' is the underlying type, e.g., struct, int, ptr.
	// This is different from the 'Type', which can be a named type like 'User'.
	fmt.Printf("Type: %s, Kind: %s\n", t.Name(), t.Kind())

	// Use a switch on the Kind to handle different types of data.
	switch t.Kind() {
	case reflect.Int:
		// For basic types, we can extract the value directly.
		fmt.Printf("Integer value: %d\n", val.Int())

	case reflect.String:
		fmt.Printf("String value: \"%s\"\n", val.String())

	case reflect.Struct:
		// If it's a struct, we can iterate over its fields.
		fmt.Println("Fields of the struct:")
		// t.NumField() gives the number of fields in the struct.
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)   // Get the field's type information.
			value := val.Field(i) // Get the field's value.
			fmt.Printf("  - Name: %s, Type: %s, Value: %v\n", field.Name, field.Type, value.Interface())
		}

	case reflect.Ptr:
		// If it's a pointer, we can inspect the element it points to.
		elem := val.Elem()
		fmt.Printf("This is a pointer to a '%s'\n", elem.Kind())

		// We can also check if the value can be modified.
		// For a value to be settable, it must be addressable, which usually means
		// it was obtained from a pointer.
		if elem.CanSet() {
			fmt.Println("The value it points to can be changed.")
			// Let's modify it if it's an integer.
			if elem.Kind() == reflect.Int {
				fmt.Printf("Original value: %d\n", elem.Int())
				elem.SetInt(1000) // Use reflection to set the value.
				fmt.Printf("Modified value: %d\n", elem.Int())
			}
		}

	default:
		fmt.Printf("Unhandled kind: %s\n", t.Kind())
	}

	fmt.Println("---------------------------------")
}

// User is a simple struct for demonstration purposes.
type User struct {
	ID   int
	Name string
}

func main() {
	// --- Case 1: Inspecting a Struct ---
	user := User{ID: 1, Name: "Ada Lovelace"}
	inspect(user)

	// --- Case 2: Inspecting a simple type (string) ---
	message := "Hello, Reflection!"
	inspect(message)

	// --- Case 3: Inspecting and Modifying via a Pointer ---
	number := 42
	fmt.Printf("Original 'number' variable before inspection: %d\n", number)
	// We pass a pointer to the 'number' variable.
	inspect(&number)
	fmt.Printf("Final 'number' variable after inspection: %d\n", number)
}
