package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func multiplyMatrices(m1 [][]int, m2 [][]int, pow int) [][]int {

	result := make([][]int, len(m1))
	for i := 0; i < len(m1); i++ {
		result[i] = make([]int, len(m2[0]))
		for j := 0; j < len(m2[0]); j++ {
			for k := 0; k < len(m2); k++ {
				result[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}
	if pow == 2 {
		return result
	} else {
		result = multiplyMatrices(result, m1, pow-1)
	}
	return result
}

func createMatrix(row, col int) [][]int {
	r := make([][]int, row)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			r[i] = append(r[i], random(-5, i*j))
		}
	}

	return r
}

func main() {
	rand.Seed(time.Now().Unix())
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("Wrong number of arguments!")
	}

	var dimension, power int
	dimension, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println("Need an integer: ", arguments[1])
		return
	}

	power, err = strconv.Atoi(arguments[2])
	if err != nil {
		fmt.Println("Need an integer: ", arguments[2])
		return
	}

	// Initialize matrix. Only square matrices.
	matrix := createMatrix(dimension, dimension)
	fmt.Println(matrix)

	// Muliply.
	r1 := multiplyMatrices(matrix, matrix, power)
	fmt.Println("r1:", r1)
}
