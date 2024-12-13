package main

import "fmt"

func main() {
	fmt.Println("Enter size of triangle: ")
	var size int
	fmt.Scan(&size)
	displayRightAngle(size)
}

func displayRightAngle(size int) {
	for index := 0; index < size; index++ {
		for count := 0; count < index; count++ {
			fmt.Print("*")
		}
		fmt.Println()

	}
}
