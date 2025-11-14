package main

import (
	"context"
	"fmt"
	"time"
)

// longRunningTask simulates a task that takes a significant amount of time to complete.
// It accepts a context.Context, which allows the task to listen for cancellation signals.
func longRunningTask(ctx context.Context) {
	fmt.Println("Task started...")

	// We use a `for` loop to simulate continuous work.
	for i := 1; ; i++ {
		select {
		// Case 1: The context has been canceled.
		// `ctx.Done()` returns a channel that is closed when the context is canceled
		// or its deadline expires.
		case <-ctx.Done():
			// `ctx.Err()` returns the reason for the cancellation.
			fmt.Printf("Task canceled. Reason: %v\n", ctx.Err())
			return // Exit the function gracefully.

		// Case 2: The context is not canceled, so we continue our work.
		default:
			fmt.Printf("Doing work, step %d...\n", i)
			// Simulate a piece of work that takes 1 second.
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	fmt.Println("Starting the main application.")

	// We create a parent context. `context.Background()` is typically used at the
	// highest level of an application.
	parentCtx := context.Background()

	// Now, we create a new context that will be automatically canceled after 3 seconds.
	// `context.WithTimeout` returns the new context and a cancel function.
	ctx, cancel := context.WithTimeout(parentCtx, 3*time.Second)

	// It's a best practice to call the `cancel` function using `defer`.
	// This ensures that any resources associated with the context are cleaned up,
	// even if the function exits early for other reasons.
	defer cancel()

	// We run our long-running task in a separate goroutine.
	// We pass the cancellable context to it.
	go longRunningTask(ctx)

	// The main goroutine will wait here.
	// We could use a `sync.WaitGroup` or simply wait for the context to be done.
	// Waiting on `ctx.Done()` demonstrates how the main flow can also react to the cancellation.
	<-ctx.Done()

	fmt.Println("Main application: The context has been canceled. The program will now exit.")

	// We add a small delay to ensure the cancellation message from the goroutine is printed before the main function exits.
	time.Sleep(100 * time.Millisecond)
}
