package main

import (
	"fmt"
	"sync"
)

// Exercise 28: Mutexes
// Write a program where multiple goroutines increment a shared counter
// safely using a sync.Mutex.

func main() {
	// A WaitGroup to wait for all goroutines to finish.
	var wg sync.WaitGroup

	// A Mutex to protect access to the shared counter.
	// A mutex (mutual exclusion lock) ensures that only one goroutine
	// can access the code between Lock() and Unlock() at any given time.
	var mu sync.Mutex

	// The shared counter that multiple goroutines will increment.
	var counter int = 0

	// The number of goroutines to launch.
	numGoroutines := 1000

	// We add the total number of goroutines to the WaitGroup counter.
	wg.Add(numGoroutines)

	// Launch 1000 goroutines.
	for i := 0; i < numGoroutines; i++ {
		go func() {
			// Defer wg.Done() to ensure it is called when the goroutine exits.
			defer wg.Done()

			// Lock the mutex before accessing the shared counter.
			// This blocks any other goroutine from acquiring the lock
			// until it is released.
			mu.Lock()

			// --- Critical Section ---
			// The code here is now safe from race conditions. Only one
			// goroutine can be executing this part at a time.
			counter++
			// --- End Critical Section ---

			// Unlock the mutex to allow other goroutines to acquire it.
			mu.Unlock()
		}()
	}

	// Wait for all goroutines to call wg.Done().
	wg.Wait()

	// Print the final result. If the mutex is used correctly,
	// the result will always be 1000.
	fmt.Printf("Final counter value: %d\n", counter)
}
