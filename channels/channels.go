package main

// `go run channels.go`

// Import libraries
import (
	"fmt"
)

func main() {
	// Make a channel "done" that communicates boolean values
	c := make(chan bool)

	// Asynchronously call waitAndSay (on a different thread)
	// Returns immediately
	go newWaitAndSay(c, "World")

	fmt.Println("Hello")

	// Send a signal to c
	c <- true

	// Wait/block until a value comes from channel "c"
	<-c
}

func newWaitAndSay(c chan bool, s string) {
	// Wait/block until we receive a boolean over channel "c" - assigned to variable "b"
	// Execute body if "b" == true
	// Note the syntax of variable initialization and condition in the "if" statement
	if b := <-c; b {
		fmt.Println(s)
	}

	// Send true back over channel "c"
	c <- true
}
