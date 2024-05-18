package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func importFile(file string) ([][]int, error) {
	var err error
	var mySlice = make([][]int, 0)

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		fields := strings.Fields(line)
		temp := make([]int, 0)
		for _, v := range fields {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			temp = append(temp, n)
		}
		if len(temp) != 0 {
			mySlice = append(mySlice, temp)
		}

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if len(temp) != len(mySlice[0]) {
			return nil, errors.New("Wrong number of elements")
		}
	}
	return mySlice, nil
}

func validPuzzle(sl [][]int) bool {
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			iEl := i * 3
			jEl := j * 3
			mySlice := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
			for k := 0; k <= 2; k++ {
				for m := 0; m <= 2; m++ {
					bigI := iEl + k
					bigJ := jEl + m
					val := sl[bigI][bigJ]
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
			sum = sum + sl[i][j]
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
			sum = sum + sl[j][i]
		}
		if sum != 45 {
			return false
		}
		sum = 0
	}
	return true
}

func main() {
	mySlice := [][]int{{6, 3, 7, 1, 5, 9, 2, 4, 8},
		{2, 8, 1, 3, 4, 7, 9, 5, 6},
		{5, 9, 4, 2, 6, 8, 1, 7, 3},
		{8, 1, 6, 5, 9, 2, 7, 3, 4},
		{4, 2, 9, 7, 8, 3, 6, 1, 5},
		{3, 7, 5, 6, 1, 4, 8, 2, 9},
		{7, 4, 2, 9, 3, 6, 5, 8, 1},
		{9, 5, 3, 8, 2, 1, 4, 6, 7},
		{1, 6, 8, 4, 7, 5, 3, 9, 2}}

	if validPuzzle(mySlice) {
		fmt.Println("Correct Sudoku puzzle!")
	} else {
		fmt.Println("Incorrect Sudoku puzzle!")
	}
}
