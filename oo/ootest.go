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

func init() {
	fmt.Println("Go is automatically initializing this main package ...")
}

func main() {
	n := createNode(18)

	switch actualType := n.(type) {
	case *pkg1.SLLNode:
		fmt.Println(actualType.SNodeMessage, n.GetValue())
	case *pkg2.PowerNode:
		fmt.Println(actualType.PNodeMessage, n.GetValue())
	}

}
