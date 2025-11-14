package main

import "fmt"

// Structs: Define a Person struct with Name and Age fields. Create an instance of this struct and print its fields.

type Person struct {
	name string
	age  int
}

func newPerson(name string, age int) *Person {
	p := Person{name: name}
	p.age = age

	return &p
}

func main() {
	fmt.Println(Person{"Bob", 20})

	fmt.Println(Person{name: "Alice", age: 30})

	fmt.Println(Person{name: "Fred"})

	fmt.Println(&Person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon", 12))

	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

}
