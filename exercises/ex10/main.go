package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("--- Successful case ---")

	div, rem, err := division(5, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Quotient: %d, Remainder: %d\n", div, rem)
	}

	fmt.Println("\n--- Error case ---")

	div, rem, err = division(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Quotient: %d, Remainder: %d\n", div, rem)
	}
}

func division(dividend int, divisor int) (int, int, error) {
	if divisor == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	var quotient = dividend / divisor
	var remainder = dividend % divisor

	return quotient, remainder, nil
}
