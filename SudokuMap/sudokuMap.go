package main

import (
	"fmt"
)

func validPuzzle(mp map[int]map[int]int) bool {
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			iEl := i * 3
			jEl := j * 3
			mySlice := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
			for k := 0; k <= 2; k++ {
				for m := 0; m <= 2; m++ {
					bigI := iEl + k
					bigJ := jEl + m
					val := mp[bigI][bigJ]
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
			sum = sum + mp[i][j]
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
			sum = sum + mp[j][i]
		}
		if sum != 45 {
			return false
		}
		sum = 0
	}
	return true
}

func main() {
	myMap := map[int]map[int]int{0: map[int]int{0: 6, 1: 3, 2: 7, 3: 1, 4: 5, 5: 9, 6: 2, 7: 4, 8: 8},
		1: map[int]int{0: 2, 1: 8, 2: 1, 3: 3, 4: 4, 5: 7, 6: 9, 7: 5, 8: 6},
		2: map[int]int{0: 5, 1: 9, 2: 4, 3: 2, 4: 6, 5: 8, 6: 1, 7: 7, 8: 3},
		3: map[int]int{0: 8, 1: 1, 2: 6, 3: 5, 4: 9, 5: 2, 6: 7, 7: 3, 8: 4},
		4: map[int]int{0: 4, 1: 2, 2: 9, 3: 7, 4: 8, 5: 3, 6: 6, 7: 1, 8: 5},
		5: map[int]int{0: 3, 1: 7, 2: 5, 3: 6, 4: 1, 5: 4, 6: 8, 7: 2, 8: 9},
		6: map[int]int{0: 7, 1: 4, 2: 2, 3: 9, 4: 3, 5: 6, 6: 5, 7: 8, 8: 1},
		7: map[int]int{0: 9, 1: 5, 2: 3, 3: 8, 4: 2, 5: 1, 6: 4, 7: 6, 8: 7},
		8: map[int]int{0: 1, 1: 6, 2: 8, 3: 4, 4: 7, 5: 5, 6: 3, 7: 9, 8: 2}}

	if validPuzzle(myMap) {
		fmt.Println("Correct Sudoku puzzle!")
	} else {
		fmt.Println("Incorrect Sudoku puzzle!")
	}
}
