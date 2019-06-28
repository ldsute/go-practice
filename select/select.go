package main

/*
 * Select allows code to wait on multiple channels at the same time
 * Select blocks until one channel is ready - if multiple ready, one chosen at random
 * Handles channel sending and receiving cases
 */

import (
	"fmt"
	"math/rand"
	"time"
)

// Create a mapping (hash) from string to int
var scMapping = map[string]int{
	"James": 5,
	"Kevin": 10,
	"Rahul": 9,
}

// Simulate a database search against the scMapping (including synthetic delay)
func findSC(name string, c chan int) {
	// Simulate a database search by sleeping for a time
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	// Return the value
	c <- scMapping[name]
}

func main() {

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create 2 "databases" - simulate them using channels
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Our "subject"
	name := "James"

	// Kick off the 2 "database queries" in separate goroutines
	go findSC(name, ch1)
	go findSC(name, ch2)

	// Use "select" to wait on multiple channels - note: the results of only the first
	// responding channel will be acted upon
	select {
	// Wait on ch1 and print the results
	case sc := <-ch1:
		fmt.Println(name, "has a security clearance of", sc, "as found in server 1")
	// Wait on ch2
	case sc := <-ch2:
		fmt.Println(name, "has a security clearance of", sc, "as found in server 2")
	// Abort blocking if 5 seconds has passed without hearing from one of the other channels
	case <-time.After(5 * time.Second):
		fmt.Println("Search timed out")
	}

}
