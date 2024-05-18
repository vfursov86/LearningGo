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

func main() {
	MIN := 0
	MAX := 94
	var LENGTH int64 = 18
	numPasswords := []string{}

	arguments := os.Args

	switch len(arguments) {
	case 2:
		LENGTH, _ = strconv.ParseInt(os.Args[1], 10, 64)
	default:
		fmt.Println("Using default values!")
	}

	SEED := time.Now().Unix()
	rand.Seed(SEED)

	startChar := "!"
	var i int64 = 1
	var newChar string

	// Generate 16 passwords.
	for j := 0; j < 16; j++ {
		// Create an accumulator string.
		tempPass := []byte{}

		for {
			myRand := random(MIN, MAX)
			newChar = string(startChar[0] + byte(myRand))
			tempPass = fmt.Append(tempPass, newChar)
			if i == LENGTH {
				numPasswords = append(numPasswords, string(tempPass))
				tempPass = nil
				i = 1
				break
			}
			i++
		}
	}

	randomChosenPassword := numPasswords[random(0, len(numPasswords)-1)]
	fmt.Println(numPasswords)
	// Choose a random password from the list and combine it with a current system time.
	fmt.Println(randomChosenPassword + strconv.Itoa(int(SEED)))
}
