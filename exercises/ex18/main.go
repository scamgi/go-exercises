package main

import (
	"fmt"
	"mygoproject/calculator"
)

func main() {
	var num1, num2 int

	fmt.Print("Type the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Type the second number: ")
	fmt.Scan(&num2)

	fmt.Printf("The sum is %v\n", calculator.Sum(num1, num2))
	fmt.Printf("The sub is %v\n", calculator.Sub(num1, num2))
}
