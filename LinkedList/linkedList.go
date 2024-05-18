package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Value int
	Next  *Node
}

var root = new(Node)

func addNode(t *Node, v int) int {
	// 1. List does not exists.
	if root == nil {
		t = &Node{v, nil}
		root = t
		return 0
	}
	// 2. Handle duplicates.
	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}
	// 3. Add Node before the root.
	if t == root && v < t.Value {
		temp := &Node{v, root}
		root = temp
		return -2
	}
	// 4. Common case. Add new Node somewhere in list.
	if t.Next != nil && v < t.Next.Value {
		temp := &Node{v, t.Next}
		t.Next = temp
		return -2

	}
	// 5. Add new Node at the end of list.
	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}
	// 6. Keep scanning for eligible place in the list.
	return addNode(t.Next, v)
}

func deleteNode(t *Node, v int) int {
	// 1. Return -1 if empty list.
	if root == nil {
		fmt.Println("Empty list now.")
		return -1
	}
	// 2. Handle first elemetn deletion.
	if root.Value == v {
		root = root.Next
	}
	// 3. Handle the latest node deletion.
	if t.Next == nil {
		fmt.Println("Nothing to delete.")
		return -1
	}
	// 4. Main logic.
	if v == t.Next.Value {
		t.Next = t.Next.Next
		return 1
	}

	// 4. Keep scanning for eligible place in the list.
	return deleteNode(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}

	fmt.Println()
}

func lookupNode(t *Node, v int) bool {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

func main() {
	fmt.Println(root)
	root = nil
	for i := 0; i < 10; i++ {
		addNode(root, i)
	}
	traverse(root)

	if lookupNode(root, 100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}

	deleteNode(root, 3)
	traverse(root)
	deleteNode(root, 9)
	traverse(root)
	deleteNode(root, 1)
	traverse(root)
	deleteNode(root, 15)
	traverse(root)
	deleteNode(root, 0)
	traverse(root)
}

func random() int {
	return rand.Intn(1000)
}
