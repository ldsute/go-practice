package main

import "fmt"

// This method is automatically called when this package is first used.
// This init method can go anywhere in the package code files.
func init() {
	fmt.Println("Go is automatically initializing this \"main\" package ...")
}
