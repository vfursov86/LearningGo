package main

import (
	"fmt"
)

var xs = [2][2][2]int{{{1, 45}, {-436, 0}}, {{234, 45}, {-456430, 1}}}
var ys = [2][2][2]int{{{-2, 14}, {45, 54}}, {{11, 0}, {456, 3}}}

func main() {
	fmt.Println(additionArrays(xs, ys))
	fmt.Println(subtractionArrays(xs, ys))
}

func additionArrays(a1 [2][2][2]int, a2 [2][2][2]int) [2][2][2]int {

	var result [2][2][2]int

	for k0, v0 := range a1 {
		for k1, v1 := range v0 {
			for k2, v2 := range v1 {
				result[k0][k1][k2] = v2 + a2[k0][k1][k2]
			}
		}
	}

	return result
}

func subtractionArrays(a1 [2][2][2]int, a2 [2][2][2]int) [2][2][2]int {

	var result [2][2][2]int

	for k0, v0 := range a1 {
		for k1, v1 := range v0 {
			for k2, v2 := range v1 {
				result[k0][k1][k2] = v2 - a2[k0][k1][k2]
			}
		}
	}

	return result
}
