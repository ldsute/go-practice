package main

/*
 * Buffered channels specify a number of values that can exist in the channel before blocking occurs
 * Sender only blocks when buffer is full
 * Receiver only blocks when buffer is empty
 *
 * Use "range" to receive values from a channel inside a for loop. Note: range will block until close is called
 * Use "close" to retire a channel
 * test commit!
 */

// Import libraries
import (
	"fmt"
)

func main() {
	// Make a string channel with a buffer size of 2
	ch := make(chan string, 2)

	// Add 2 items to the channel
	ch <- "Hello"
	ch <- "World"

	// Pull 2 items off the channel and print
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Start a goroutine to say Hello 10 times
	// If this isn't a separate goroutine, we have deadlock - producer blocks and consumer hasn't started
	go SayHelloMultipleTimes(ch, 10)

	// Now receive channel values using range - until the channel is closed
	for s := range ch {
		fmt.Println(s)
	}
}

// SayHelloMultipleTimes pushes "Hello" on to the given channel the specified number of times
func SayHelloMultipleTimes(c chan string, numberOfTimes int) {
	for i := 0; i <= numberOfTimes; i++ {
		c <- "Hello"
	}
	close(c)
}
