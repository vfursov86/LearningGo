package main

import (
	"fmt"
)

type Node struct {
	Value float64
	Next  *Node
}

var size = 0
var queue = new(Node)

func Push(t *Node, v float64) bool {
	if queue == nil {
		queue = &Node{v, nil}
		size++
		return true
	}

	t = &Node{v, nil}
	t.Next = queue
	queue = t
	size++
	return true
}

func Pop(t *Node) (float64, bool) {
	if size == 0 {
		return 0, false
	}

	if size == 1 {
		queue = nil
		size--
		return t.Value, true
	}

	temp := t
	for (t.Next) != nil {
		temp = t
		t = t.Next
	}

	v := (temp.Next).Value
	temp.Next = nil

	size--
	return v, true

}

func traverse(t *Node) {
	if size == 0 {
		fmt.Println("Empty Queue!")
		return
	}
	for t != nil {
		fmt.Printf("%.2f -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	queue = nil
	Push(queue, 10)
	fmt.Println("Size:", size)
	traverse(queue)

	v, b := Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)

	Push(queue, 0.3456)
	Push(queue, 43563.456)
	Push(queue, 0.7)
	Push(queue, 3.14)

	traverse(queue)
	fmt.Println("Size:", size)

	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)

	v, b = Pop(queue)
	if b {
		fmt.Println("Pop:", v)
	}
	fmt.Println("Size:", size)
	traverse(queue)
}
