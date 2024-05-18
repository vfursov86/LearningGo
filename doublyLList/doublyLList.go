package main

import (
	"fmt"
)

var root = new(Node)

type Node struct {
	Value    int
	Previous *Node
	Next     *Node
}

func addNode(t *Node, v int) int {
	// 1. List does not exists.
	if root == nil {
		t = &Node{v, nil, nil}
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
		temp := &Node{v, nil, root}
		root = temp
		return -2
	}
	// 4. Common case. Add new Node somewhere in list.
	if t.Next != nil && v < t.Next.Value {
		temp := &Node{v, t, t.Next}
		t.Next = temp
		return -2

	}
	// 5. Add new Node at the end of list.
	if t.Next == nil {
		t.Next = &Node{v, t, nil}
		return -2
	}
	// 6. Keep scanning for eligible place in the list.
	return addNode(t.Next, v)

}

func deleteNode(t *Node, v int) int {
	founded, node := lookupNode(t, v)
	if founded {
		node.Previous.Next = node.Next
		return node.Value
	}
	return -1

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

func reverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	temp := t
	for t != nil {
		temp = t
		t = t.Next
	}

	for temp.Previous != nil {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Previous
	}
	fmt.Printf("%d -> ", temp.Value)
	fmt.Println()
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	n := 0
	for t != nil {
		n++
		t = t.Next
	}
	return n
}

func lookupNode(t *Node, v int) (bool, *Node) {
	if root == nil {
		return false, nil
	}

	if v == t.Value {
		return true, t
	}

	if t.Next == nil {
		return false, nil
	}

	return lookupNode(t.Next, v)
}

func main() {
	fmt.Println(root)
	root = nil
	traverse(root)
	addNode(root, 1)
	addNode(root, 1)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 0)
	addNode(root, -1)
	addNode(root, -10)
	addNode(root, -5)
	addNode(root, 100)
	addNode(root, 45)
	addNode(root, 123)
	traverse(root)

	deleteNode(root, -5)
	deleteNode(root, 45)
	deleteNode(root, 123)
	traverse(root)
	reverse(root)
}
