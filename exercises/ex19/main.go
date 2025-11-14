package main

import (
	"fmt"
	"os"
)

func main() {
	// file, err := os.Open("secret.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(file.Stat())
	// defer file.Close()

	content, err := os.ReadFile("secret.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(content))
}
