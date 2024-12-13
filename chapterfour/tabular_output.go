package main

import "fmt"

func main() {

	fmt.Printf("%s\t%s\t%s\t%s\n", "N", "N2", "N3", "N4")
	for index := 1; index <= 5; index++ {
		fmt.Printf("%d\t%d\t%d\t%d\n", index, pow(index, 2), pow(index, 3), pow(index, 4))
	}
}

func pow(base int, exponent int) int {
	product := 1
	for index := 0; index < exponent; index++ {
		product *= base
	}
	return product
}
