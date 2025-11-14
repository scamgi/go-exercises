package main

import (
	"fmt"
)

func main() {
	var name string

	fmt.Print("Hi! Please eneter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Hi, %v!\n", name)
}
