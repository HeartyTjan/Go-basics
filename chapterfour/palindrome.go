package main

import (
	"fmt"
)

func main() {
	palindrome := ""
	fmt.Print("Enter digit : ")
	var digit string
	fmt.Scan(&digit)
	for index := len(digit) - 1; index >= 0; index-- {
		palindrome += string(digit[index])

	}
	if digit == palindrome {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
