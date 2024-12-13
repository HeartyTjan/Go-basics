package main

import "fmt"

func main() {
	fmt.Println("Enter your weight : ")
	var weight float64
	fmt.Scanf("%f", &weight)

	fmt.Println("Enter your height: ")
	var height float64
	fmt.Scanf("%f", &height)
	fmt.Printf("Body Max is : %.2f ", calculateBMI(weight, height))

}

func calculateBMI(weight float64, height float64) float64 {
	result := weight * 703 / height * height
	return result
}
