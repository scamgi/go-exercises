package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Print("Hello ")
	}()

	go func() {
		defer wg.Done()
		fmt.Print("World ")
	}()

	wg.Wait()
}
