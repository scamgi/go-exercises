package main

import (
	"errors"
	"fmt"
)

func division(dividend int, divisor int) (int, int, error) {
	if divisor == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	var quotient = dividend / divisor
	var remainder = dividend % divisor

	return quotient, remainder, nil
}

func performDivision(dividend int, divisor int) error {
	quotient, remainder, err := division(dividend, divisor)
	if err != nil {
		return fmt.Errorf("an error occurred during division of %d by %d: %w", dividend, divisor, err)
	}

	fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)
	return nil
}

func main() {
	fmt.Println("--- Successful Case ---")
	err := performDivision(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n--- Error Case ---")
	err = performDivision(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
