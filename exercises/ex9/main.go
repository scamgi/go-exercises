package main

import "fmt"

func main() {
	fmt.Println(division(5, 3))
}

func division(first int, second int) (int, int) {
	var div = first / second
	var rem = first % second
	return div, rem
}
