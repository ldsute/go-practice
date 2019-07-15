package main

/**
 * No exceptions in go
 * All errors implement the error interface
 * Functions should return errors - can use multiple return values
 * To instantiate error, use errors.New() or fmt.Errorf() to create errors on the fly ...
 * or create own type
 */

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// "class" sampleError implements type "error" interface
// To implement an interface in Go, we just need to implement all the methods in the interface, which is just one in this case
type sampleError struct {
	Message string
}

// Member function "Error" for the sampleError "class"
// Implements error interface
func (e sampleError) Error() string {
	return e.Message
}

func createAnError(errorMessage string) error {
	number := rand.Intn(5)

	if number == 0 {
		// Use a custom error class we've created
		return sampleError{"Here's an error with message (using custom class): " + errorMessage}
	} else if number == 1 {
		// use the fmt.Errorf call
		return fmt.Errorf("Here's an error with message (using fmt.Errorf): %s", errorMessage)
	} else if number == 2 {
		// Use the errors.New() call
		return errors.New("Here's an error with message (using errors.New): " + errorMessage)
	} else if number == 3 {
		// Return no error at all
		return nil
	} else {
		panic("Panic now!")
	}
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// This function call is deferred until main() is exited. If recover() actually recovers
	// from a panic() call, the err will have content we can act upon.
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered: ", err)
		}
	}() // These parens make this a function call, not just a function declaration

	// Call the createAnError message - and only enter the body if err != nil
	if err := createAnError("test message"); err != nil {
		// Do a type check - is err of type sampleError?
		if e, ok := err.(sampleError); ok {
			fmt.Println("Got sampleError:", e.Message)
		} else {
			fmt.Println("Got canned error:", err)
		}
	} else {
		fmt.Println("No error here.")
	}
}
