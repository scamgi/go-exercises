package main

import "fmt"

// Exercise 26: Buffered Channels
// Create a buffered channel and demonstrate sending multiple values
// without a receiver immediately ready.

func main() {
	// Create a buffered channel of strings with a capacity of 3.
	// This means up to 3 values can be sent to the channel without
	// a corresponding receiver being ready.
	messages := make(chan string, 3)

	fmt.Println("Sending three messages to the buffered channel...")

	// Send three values to the channel.
	// These sends are non-blocking because the channel's buffer is not yet full.
	messages <- "first message"
	messages <- "second message"
	messages <- "third message"
	// messages <- "fourth message" // this line causes deadlock

	fmt.Println("All three messages were sent without blocking.")

	// If you were to send a fourth message here, the program would block
	// until a receiver reads a value, freeing up space in the buffer.
	// For example: `messages <- "fourth message"` would cause a deadlock in this program.

	fmt.Println("\nReceiving messages from the channel...")

	// Now, we can receive the three values that are held in the buffer.
	fmt.Println("Received:", <-messages)
	fmt.Println("Received:", <-messages)
	fmt.Println("Received:", <-messages)
	// fmt.Println("Received:", <-messages) // this line causes deadlock too

	fmt.Println("\nAll buffered messages have been received.")
}
