package main

import "fmt"

func main() {
	var m map[string]int = make(map[string]int)

	m["Tom"] = 20
	m["Giacomo"] = 24
	m["Tiziano"] = 12

	fmt.Printf("%v %v %v\n", m["Tom"], m["Giacomo"], m["Tiziano"])
	fmt.Printf("%v %v %v\n", m["Tom"], m["Giacomo"], m["Elio"])

	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
}
