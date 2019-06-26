package main

/*
 * goroutines share the same memory as main thread
 */

// `go run goroutines1.go`

// Import libraries
import (
	"fmt"
	"time"
)

func main() {
	word := "Hello"

	// Kick off anonymous function in a different thread
	// Remember, this call immediately returns since the method will be executed in a separate goroutine
	go func() {
		// Sleep for 2 seconds
		time.Sleep(2 * time.Second)

		// Print the "word" defined in the parent (main thread) scope
		fmt.Println(word)
	}()

	fmt.Println("Hello")

	// Re-assign "word" variable. Since this happens before the anonymous function executes,
	// "word" will contain "World" instead of "Hello" when it executes.
	word = "World"

	// Wait 3 seconds before exiting the program.
	// If we don't wait, the program exits, regardless of whether waitAndSay() actually executed
	time.Sleep(3 * time.Second)
}
