package main

import "fmt"

// Pointers: Write a function that takes a pointer to an integer and modifies the original value

func pointer(number *int) {
	*number = 3
}

func main() {
	a := 18

	fmt.Println("Prima:", a)

	pointer(&a)

	fmt.Println("Dopo:", a)
}
