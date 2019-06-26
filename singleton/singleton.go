package main

// `go run singleton.go`

// Import libraries
import (
	"fmt"
	"sync"
)

func main() {
	// Create a "once" variable - it's Do() method will only allow the function
	// passed to it to be run once - no matter how many times Do() is called
	var once sync.Once

	// Create an anonymous function and point to it with the onceBody variable
	onceBody := func() {
		fmt.Println("once and only once")
	}

	// Make a channel "done" that communicates boolean values
	done := make(chan bool)

	for i := 0; i < 10; i++ {

		// Asynchronously call this anonymous function (on a different thread)
		go func() {

			// Wrap the call to onceBody() with the "once" variable - the Do() function
			// will only allow onceBody() to actually be called once
			once.Do(onceBody)

			// Send true to the "done" channel
			done <- true
		}()
	}

	// Wait 10 times for a value to be received on the "done" channel
	for i := 0; i < 10; i++ {
		<-done
	}
}
