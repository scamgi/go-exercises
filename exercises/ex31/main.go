package main

import "fmt"

// printSlice is a generic function that can print elements of a slice of any type.
// [T any] declares a type parameter 'T' with the 'any' constraint, meaning T can be any type.
func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

func main() {
	// Create slices of different types.
	intSlice := []int{10, 20, 30}
	stringSlice := []string{"hello", "ciao", "hola"}
	floatSlice := []float64{3.14, 1.618, 2.718}

	fmt.Print("Integer Slice: ")
	printSlice(intSlice)

	fmt.Print("String Slice: ")
	printSlice(stringSlice)

	fmt.Print("Float Slice: ")
	printSlice(floatSlice)
}
