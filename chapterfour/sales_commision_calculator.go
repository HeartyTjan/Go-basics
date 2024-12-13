package main

import "fmt"

func main() {
	const BASESALARY float64 = 200.0
	const PERCENT float64 = 0.09

	fmt.Println("Agent name : ")
	var agentName string
	fmt.Scan(&agentName)
	fmt.Println("Enter agent sales for the week :")
	var salesAmount float64
	fmt.Scan(&salesAmount)

	comission := BASESALARY + (salesAmount * PERCENT)
	fmt.Print("Name\tCommission\n")
	fmt.Printf("%s\t %.2f\n", agentName, comission)

}
