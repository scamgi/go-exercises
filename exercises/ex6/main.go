package main

import "fmt"

func main() {
	var numb int

	fmt.Print("Please enter a number: ")
	fmt.Scan(&numb)

	if numb%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}
}
