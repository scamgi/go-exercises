package main

import (
	"fmt"
	"time"
)

// Exercise 24: Channels (Basics)
// Use a channel to send a message from one goroutine to another.

func main() {
	// Create a channel to transport messages of type string.
	// Channels are the pipes that connect concurrent goroutines.
	messages := make(chan string)

	// Launch a new goroutine (the sender) as an anonymous function.
	go func() {
		fmt.Println("Goroutine: Preparing to send a message...")
		time.Sleep(1 * time.Second) // Simulate some work.
		// Send a message into the channel using the <- operator.
		messages <- "Hello from the other side!"
		fmt.Println("Goroutine: Message has been sent.")
	}()

	fmt.Println("Main: Waiting to receive a message...")

	// The main goroutine blocks here, waiting to receive a value
	// from the 'messages' channel. This synchronizes the two goroutines.
	receivedMessage := <-messages

	fmt.Println("Main: Received message ->", receivedMessage)
}
