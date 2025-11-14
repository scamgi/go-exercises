package main

import "fmt"

func main() {
	underlyingArray := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	mySlice := underlyingArray[2:5] // len=3, cap=6, elements [c, d, e]

	fmt.Printf("Initial Slice -> len: %d, cap: %d, data: %v\n", len(mySlice), cap(mySlice), mySlice)
	fmt.Printf("Underlying Array -> %v\n\n", underlyingArray)

	// Append one element. Still fits within capacity.
	mySlice = append(mySlice, "X")
	fmt.Printf("After 1st Append -> len: %d, cap: %d, data: %v\n", len(mySlice), cap(mySlice), mySlice)
	// The underlying array is MODIFIED
	fmt.Printf("Underlying Array -> %v\n\n", underlyingArray)

	// Append three more elements. This will exceed the original capacity.
	mySlice = append(mySlice, "Y", "Z", "W")
	fmt.Printf("After 2nd Append -> len: %d, cap: %d, data: %v\n", len(mySlice), cap(mySlice), mySlice)
	// The original underlying array is NO LONGER AFFECTED
	fmt.Printf("Underlying Array -> %v\n", underlyingArray)
}
