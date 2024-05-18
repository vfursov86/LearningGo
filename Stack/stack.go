package main

import (
	"fmt"
)

type Node struct {
	Value, Number, Seed int
	Next                *Node
}

var size = 0
var stack = new(Node)

func Push(v, n, s int) bool {
	if stack == nil {
		stack = &Node{v, n, s, nil}
		size = 1
		return true
	}

	temp := &Node{v, n, s, nil}
	temp.Next = stack
	stack = temp
	size++
	return true
}

func Pop(t *Node) (int, int, int, bool) {
	if size == 0 {
		return 0, 0, 0, false
	}

	if size == 1 {
		size = 0
		stack = nil
		return t.Value, t.Number, t.Seed, true
	}

	stack = stack.Next
	size--
	return t.Value, t.Number, t.Seed, true
}

func traverse(t *Node) {
	if size == 0 {
		fmt.Println("Empty stack!")
		return
	}

	for t != nil {
		fmt.Printf("%d, %d, %d -> ", t.Value, t.Number, t.Seed)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	stack = nil
	v, n, s, b := Pop(stack)
	if b {
		fmt.Print(v, n, s, " ")
	} else {
		fmt.Println("Pop() failed!")
	}

	Push(100, 34, 435)
	traverse(stack)
	Push(200, 354, 2)
	traverse(stack)

	for i := 0; i < 10; i++ {
		Push(i, i*2, i+45)
	}

	for i := 0; i < 15; i++ {
		v, n, s, b := Pop(stack)
		if b {
			fmt.Print(v, n, s, " ")
		} else {
			break
		}
	}

	fmt.Println()
	traverse(stack)
}
