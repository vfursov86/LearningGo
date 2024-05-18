package main

import (
	"container/heap"
	"fmt"
)

type ComplexStruct struct {
	Val1 int
	Val2 float32
	Val3 string
}

type heapComplex []ComplexStruct

func (n *heapComplex) Pop() interface{} {
	old := *n
	x := old[len(old)-1]
	new := old[0 : len(old)-1]
	*n = new
	return x
}

func (n *heapComplex) Push(x interface{}) {
	*n = append(*n, x.(ComplexStruct))
}

func (n heapComplex) Len() int {
	return len(n)
}

func (n heapComplex) Less(a, b int) bool {
	return n[a].Val1 < n[b].Val1
}

func (n heapComplex) Swap(a, b int) {
	n[a], n[b] = n[b], n[a]
}

func main() {
	myHeap := &heapComplex{ComplexStruct{23, 7.12, "some1"}, ComplexStruct{-34, 0.56, "test"}, ComplexStruct{-1000, -0.23, "abc"}}
	heap.Init(myHeap)
	size := len(*myHeap)
	fmt.Printf("Heap size: %d\n", size)
	fmt.Printf("%v\n", myHeap)

	myHeap.Push(ComplexStruct{56, 56.4, "newItem"})
	myHeap.Push(ComplexStruct{198435345, 0.1, "newItem2"})
	fmt.Printf("Heap size: %d\n", len(*myHeap))
	fmt.Printf("%v\n", myHeap)
	heap.Init(myHeap)
	fmt.Printf("%v\n", myHeap)
}
