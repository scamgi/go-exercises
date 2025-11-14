package main

import "fmt"

type Person struct {
	name string
	age  uint
}

func (person Person) greetingsFromMe() {
	fmt.Printf("Hi from %v! I am %v years old.\n", person.name, person.age)
}

func main() {
	var name string
	var age uint
	fmt.Print("Hi! Tell me your name: ")
	fmt.Scan(&name)
	fmt.Print("What's your age?: ")
	fmt.Scan(&age)

	var person Person = Person{name, age}

	person.greetingsFromMe()
}
