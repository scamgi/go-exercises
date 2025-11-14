package main

import (
	"fmt"
	"time"
)

// Exercise 25: `select` Statement
// Write a program with multiple channels and use a `select` statement
// to receive from whichever channel has a message ready.

func main() {
	// Create two channels.
	channel1 := make(chan string)
	channel2 := make(chan string)

	// Goroutine 1: Sends a message to channel1 after a 2-second delay.
	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "Message from channel 1"
	}()

	// Goroutine 2: Sends a message to channel2 after a 1-second delay.
	go func() {
		time.Sleep(1 * time.Second)
		channel2 <- "Message from channel 2"
	}()

	fmt.Println("Waiting for a message from either channel...")

	// The select statement blocks until one of its cases can run.
	// It chooses one at random if multiple are ready.
	select {
	case msg1 := <-channel1:
		fmt.Println("Received:", msg1)
	case msg2 := <-channel2:
		fmt.Println("Received:", msg2)
	}

	fmt.Println("The first message has been received.")
}
