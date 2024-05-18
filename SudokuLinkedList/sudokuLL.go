package main

import (
	"fmt"
)

type Node struct {
	Index int
	Value int
	Next  *Node
}

var root = new(Node)

func addNode(t *Node, v int, idx int) int {
	// 1. List does not exists.
	if root == nil {
		t = &Node{idx, v, nil}
		root = t
		return 0
	}
	// 5. Add new Node at the end of list.
	if t.Next == nil {
		t.Next = &Node{idx, v, nil}
		return -2
	}
	// 6. Keep scanning for eligible place in the list.
	return addNode(t.Next, v, idx)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		fmt.Printf("%d -> ", t.Index)
		t = t.Next
	}

	fmt.Println()
}

func lookupValueByIndex(t *Node, idx int) int {

	if idx == t.Index {
		return t.Value
	}

	if t.Next == nil {
		return -1
	}

	return lookupValueByIndex(t.Next, idx)
}

func validPuzzle(sm [9]*Node) bool {
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			iEl := i * 3
			jEl := j * 3
			mySlice := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
			for k := 0; k <= 2; k++ {
				for m := 0; m <= 2; m++ {
					bigI := iEl + k
					bigJ := jEl + m
					val := lookupValueByIndex(sm[bigI], bigJ)
					if val > 0 && val < 10 {
						if mySlice[val-1] == 1 {
							fmt.Println("Appeared 2 times:", val)
							return false
						} else {
							mySlice[val-1] = 1
						}
					} else {
						fmt.Println("Invalid value:", val)
						return false
					}
				}
			}
		}
	}

	// Testing columns.
	for i := 0; i <= 8; i++ {
		sum := 0
		for j := 0; j <= 8; j++ {
			sum = sum + lookupValueByIndex(sm[i], j)
		}
		if sum != 45 {
			return false
		}
		sum = 0
	}

	// Testing raws.
	for i := 0; i <= 8; i++ {
		sum := 0
		for j := 0; j <= 8; j++ {
			sum = sum + lookupValueByIndex(sm[j], i)
		}
		if sum != 45 {
			return false
		}
		sum = 0
	}
	return true
}

func main() {

	sudokuMatrix := [9]*Node{}

	vsr1 := [9]int{6, 3, 7, 1, 5, 9, 2, 4, 8}
	vsr2 := [9]int{2, 8, 1, 3, 4, 7, 9, 5, 6}
	vsr3 := [9]int{5, 9, 4, 2, 6, 8, 1, 7, 3}
	vsr4 := [9]int{8, 1, 6, 5, 9, 2, 7, 3, 4}
	vsr5 := [9]int{4, 2, 9, 7, 8, 3, 6, 1, 5}
	vsr6 := [9]int{3, 7, 5, 6, 1, 4, 8, 2, 9}
	vsr7 := [9]int{7, 4, 2, 9, 3, 6, 5, 8, 1}
	vsr8 := [9]int{9, 5, 3, 8, 2, 1, 4, 6, 7}
	vsr9 := [9]int{1, 6, 8, 4, 7, 5, 3, 9, 2}
	init := [9][9]int{vsr1, vsr2, vsr3, vsr4, vsr5, vsr6, vsr7, vsr8, vsr9}

	for i := 0; i < 9; i++ {
		root = nil

		for k := 0; k < 9; k++ {
			addNode(root, init[i][k], k)
		}
		sudokuMatrix[i] = root
	}

	for i := 0; i < 9; i++ {
		root = sudokuMatrix[i]
		traverse(root)

	}

	if validPuzzle(sudokuMatrix) {
		fmt.Println("Correct Sudoku puzzle!")
	} else {
		fmt.Println("Incorrect Sudoku puzzle!")
	}
}
