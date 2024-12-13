package main

import "fmt"

func main() {
	count := 0
	for index := 1; index <= 100; index++ {

		if index%3 == 0 {
			fmt.Printf("%d ", index)
			count += 1
			if count%5 == 0 {
				fmt.Printf("\n")
			}
		}
	}
}
