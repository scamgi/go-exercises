package main

import "fmt"

func main() {
	var arr = []int{0, 4, 4, 6, 7}

	for i, v := range arr {
		fmt.Println(i, v)
	}

	for _, v := range arr {
		fmt.Println(v)
	}
}
