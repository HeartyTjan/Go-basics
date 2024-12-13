package main

import "fmt"

func main() {
	fmt.Println("Number of trips: ")
	var numberOfTrips int
	fmt.Scan(&numberOfTrips)

	var totalMilesDriven int
	var totalGallonUsed int

	for i := 0; i < numberOfTrips; i++ {
		fmt.Println("Enter miles driven")
		var milesDriven int
		fmt.Scanf("%d", &milesDriven)
		totalMilesDriven += milesDriven

		fmt.Println("Enter gallon used")
		var gallon int
		fmt.Scanf("%d", &gallon)
		totalGallonUsed += gallon

		var milesPerGallon = milesDriven / gallon
		fmt.Println("Miles per gallon: ", milesPerGallon)
	}
	totalMilesPerGallon := totalMilesDriven / totalGallonUsed
	fmt.Println("Total Miles per gallon: ", totalMilesPerGallon)

}
