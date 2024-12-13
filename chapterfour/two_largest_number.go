package main

import (
	"fmt"
)

func main() {
	secondLargestSales := 0
	highestSales := 0

	for index := 0; index < 4; index++ {
		fmt.Printf("Enter sales by salesperson %d\n", index+1)
		var sale int
		fmt.Scan(&sale)
		if sale > highestSales {
			secondLargestSales = highestSales
			highestSales = sale
		}
		if sale > secondLargestSales && sale < highestSales {
			secondLargestSales = sale
		}
	}
	fmt.Printf("The highest sales is %d and second is %d", highestSales, secondLargestSales)
}
