package main

import (
	"fmt"
)

type Value struct {
	Val int
}

type Row struct {
	Val1 Value
	Val2 Value
	Val3 Value
	Val4 Value
}

type Matrix struct {
	Row1 Row
	Row2 Row
	Row3 Row
	Row4 Row
}

func main() {
	matrix := Matrix{Row1: Row{Val1: Value{34}, Val2: Value{546}, Val3: Value{768}, Val4: Value{23}},
		Row2: Row{Val1: Value{678}, Val2: Value{5}, Val3: Value{-34}, Val4: Value{2546}},
		Row3: Row{Val1: Value{768768}, Val2: Value{-234}, Val3: Value{0}, Val4: Value{-89}},
		Row4: Row{Val1: Value{345}, Val2: Value{-23}, Val3: Value{-56}, Val4: Value{-879}}}

	fmt.Println(matrix.Row4.Val1)
	fmt.Println(matrix.Row2.Val2)
}
