package main

import "fmt"

func main() {
	var numero1 int
	var numero2 int

	fmt.Print("Please enter the first number: ")
	fmt.Scan(&numero1)
	fmt.Print("Please enter the second number: ")
	fmt.Scan(&numero2)
	fmt.Printf("The sum is: %v\n", numero1+numero2)
}
