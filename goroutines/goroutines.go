package main

/*
 * Don't communicate by sharing memory - share memory by communicating
 * Rather than threads sharing a variable, one thread sends the value of a variable to another thread
 * goroutines are multiplexed onto multiple OS threads - some goroutines even run on the same thread
 * Can do old-school locking via the sync package - sometimes makes more sense.
 */

// `go run goroutines.go`

// Import libraries
import (
	"fmt"
	"time"
)

func main() {
	// Kick off waitAndSay in a different thread
	// Remember, this call immediately returns since the method will be executed in a separate goroutine
	go waitAndSay("World")

	fmt.Println("Hello")

	// Wait 3 seconds before exiting the program.
	// If we don't wait, the program exits, regardless of whether waitAndSay() actually executed
	time.Sleep(3 * time.Second)
}

func waitAndSay(s string) {
	// Sleep for 2 seconds
	time.Sleep(2 * time.Second)

	// Print the string you received
	fmt.Println(s)
}
