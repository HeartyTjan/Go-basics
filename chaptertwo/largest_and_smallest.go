package main

import "fmt"

func largestAndSmallest() {
	firstInteger := 1
	secondInteger := 2
	thirdInteger := 3
	fourthInteger := 4
	fifthIntegers := 5

	largest := firstInteger
	smallest := firstInteger

	if secondInteger > largest {
		largest = secondInteger
	}
	if thirdInteger > largest {
		largest = thirdInteger
	}
	if fourthInteger > largest {
		largest = fourthInteger
	}
	if fifthIntegers > largest {
		largest = fifthIntegers
	}
	fmt.Println(largest)

	if secondInteger < smallest {
		smallest = secondInteger
	}
	if thirdInteger < smallest {
		smallest = thirdInteger
	}
	if fourthInteger < smallest {
		smallest = fourthInteger
	}
	if fifthIntegers < smallest {
		smallest = fifthIntegers
	}
	fmt.Println(smallest)
}

func main() {
	largestAndSmallest()
}
