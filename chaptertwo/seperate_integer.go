package main

import (
	"fmt"
	"strconv"
)

func main() {
	var digit int = 5342617
	splitDigit(digit)

}

func splitDigit(digit int) {
	stringDigit := strconv.Itoa(digit)
	size := len(stringDigit)
	for index := 0; index < size; index++ {
		digit, _ := strconv.Atoi(string(stringDigit[index]))
		fmt.Printf("%d ", digit)
	}

}
