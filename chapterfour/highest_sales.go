package main

import "fmt"

func main() {
	highestSales := 0
	for index := 0; index < 10; index++ {
		fmt.Printf("Enter sales by salesperson %d\n", index+1)
		var sale int
		fmt.Scan(&sale)
		if sale > highestSales {
			highestSales = sale
		}
	}
	fmt.Printf("The highest sales is %d", highestSales)

}
