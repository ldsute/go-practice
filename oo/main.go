package main

// `go run github.com/ldsute/go-practice/oo`

import (
	"fmt"

	"github.com/ldsute/go-practice/oo/pkg1"
	"github.com/ldsute/go-practice/oo/pkg2"
	"github.com/ldsute/go-practice/oo/shared"
)

func createNode(v int) shared.LLNode {
	node := pkg1.NewSLLNode()
	node.SetValue(v)
	return node
}

// Create a pointer to type, don't give it anything to point to,
// and dereference methods - go doesn't go NPE in this case
func test1() {
	fmt.Println("================== Running test1 ==================")
	var node *pkg1.SLLNode
	// Look Ma, no NPE! Just an error is returned - check out the impl.
	fmt.Println(node.SetValue(55))
}

// Trying to call a method on an interface which is uninitialized will
// cause a panic
func test2() {
	// This function call is deferred until test2() is exited. If recover() actually recovers
	// from a panic() call, the err will have content we can act upon.
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}() // These parens make this a function call, not just a function declaration

	fmt.Println("================== Running test2 ==================")
	var n shared.LLNode
	fmt.Println(n.GetValue())
}

func test3() {
	fmt.Println("================== Running test3 ==================")
	n := createNode(18)

	switch actualType := n.(type) {
	case *pkg1.SLLNode:
		fmt.Println(actualType.SNodeMessage, n.GetValue())
	case *pkg2.PowerNode:
		fmt.Println(actualType.PNodeMessage, n.GetValue())
	}
}

func main() {
	test1()
	test2()
	test3()
}
