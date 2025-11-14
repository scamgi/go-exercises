package main

import "fmt"

func main() {
	var s []int
	oldCap := cap(s)

	fmt.Println("Appending up to 2048 elements...")

	for i := 0; i < 2048; i++ {
		s = append(s, i)
		newCap := cap(s)
		if newCap > oldCap {
			fmt.Printf("len: %-4d | old cap: %-4d | new cap: %-4d\n", len(s), oldCap, newCap)
			oldCap = newCap
		}
	}
}