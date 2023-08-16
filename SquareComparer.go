package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// This program will operate on numbers and compare the difference between them
// e.g. 2^2 = 4, 3^3=9, dif of 5. 3*3=9 4*4=16, dif of 7, 4*4=16, 5*5=25, dif of 9
// so the dif grows by 2 each time, does this pattern hold?
// What other patterns can we find with other operations?
// e.g. primes and various sets

// open file for starting data

//

// Get how many numbers to square
func GetNumber() int64 {
	var userInput string
	if len(os.Args) < 2 {
		fmt.Println("Please enter how many numbers to square and compare")
		fmt.Scanln(&userInput)
		userInput = strings.TrimSuffix(userInput, "\n")
	} else {
		userInput = os.Args[1]
	}

	output, _ := strconv.ParseInt(userInput, 0, 64)

	return output
}

func main() {

	startingNumber := int(GetNumber())

	for i := 0; i < startingNumber; i++ {
		x := i * i
		y := (i - 1) * (i - 1)
		z := x - y
		fmt.Println(i, x, z)

	}
}

//

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
