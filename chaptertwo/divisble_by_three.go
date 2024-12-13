package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var number int
	_, err := fmt.Scan("%d", &number)
	if err != nil {
		return
	}

	if number%3 == 0 {
		fmt.Printf("%d is divisible by 3", number)
	} else {
		fmt.Printf("%d is not divisible by 3", number)
	}
}
