package calculator

import "testing"

// Exercise 29: Testing (Basics)
// Write a simple unit test for your calculator package.

// TestSum tests the Sum function.
// Go test functions must start with "Test" and take a single argument of type *testing.T.
func TestSum(t *testing.T) {
	// We define a simple test case.
	a := 5
	b := 3
	expected := 8

	// We call the function we want to test.
	result := Sum(a, b)

	// We check if the result is what we expect.
	if result != expected {
		// If the result is not what we expected, we use the t.Errorf function
		// to fail the test and print a formatted error message.
		t.Errorf("Sum(%d, %d) = %d; expected %d", a, b, result, expected)
	}
}

// It's good practice to test various cases, including zero and negative numbers.
func TestSumWithNegativeNumbers(t *testing.T) {
	a := -10
	b := 5
	expected := -5

	result := Sum(a, b)

	if result != expected {
		t.Errorf("Sum(%d, %d) = %d; expected %d", a, b, result, expected)
	}
}

// We can also add a test for our Sub function.
func TestSub(t *testing.T) {
	a := 10
	b := 4
	expected := 6

	result := Sub(a, b)

	if result != expected {
		t.Errorf("Sub(%d, %d) = %d; expected %d", a, b, result, expected)
	}
}

func TestSubWithNegativeNumber(t *testing.T) {
	a := 10
	b := -4
	expected := 14

	result := Sub(a, b)

	if result != expected {
		t.Errorf("Sub(%d, %d) = %d; expected %d", a, b, result, expected)
	}
}
