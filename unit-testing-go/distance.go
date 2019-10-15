package main

import (
	"fmt"
	"math"
)

var x1, x2, y1, y2 float64

// DistanceInterface allows user to input items
func DistanceInterface() error {

	fmt.Println("")
	fmt.Println("*** Distance Formula ***")
	fmt.Println("Enter value for coordinates (x1, x2), (y1, y2)")

	fmt.Print("Please enter x1: ")
	_, err := fmt.Scanf("%f", &x1)
	if err != nil {
		fmt.Println("***\nInvalid input for distance \n***")
		return err
	}
	fmt.Print("Please enter x2: ")
	fmt.Scanf("%f", &x2)
	fmt.Print("Please enter y1: ")
	fmt.Scanf("%f", &y1)
	fmt.Print("Please enter y2: ")
	fmt.Scanf("%f", &y2)

	fmt.Printf("Distance is: %.2f\n", CalculateDistance(x1, x2, y1, y2))
	fmt.Println("")

	return nil
}

// CalculateDistance computes distance formula
func CalculateDistance(x1, x2, y1, y2 float64) float64 {
	X := (x2 - x1)
	Y := (y2 - y1)
	return math.Sqrt(Squared(X) + Squared(Y))
}

// Squared returns float input squared
func Squared(value float64) float64 {
	return value * value
}

// compare to mathewarehouse.com
// http://www.mathwarehouse.com/calculators/distance-formula-calculator.php
