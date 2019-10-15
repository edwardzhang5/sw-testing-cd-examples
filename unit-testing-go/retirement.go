package main

import "fmt"

var aSalary, pSaved, sGoal float64
var cAge int

// RetirementInterface gets user input and provides output
func RetirementInterface() {

	fmt.Println("")
	fmt.Println("*** Retirement Savings ***")
	fmt.Println("Enter values below to calculate retirement")
	fmt.Print("Current Age: ")
	fmt.Scanf("%d", &cAge)
	fmt.Print("Annual Salary: ")
	fmt.Scanf("%f", &aSalary)
	fmt.Print("Percentage Saved (e.g., 20): ")
	fmt.Scanf("%f", &pSaved)
	fmt.Print("Desired Savings Goal: ")
	fmt.Scanf("%f", &sGoal)

	yearlySavings := YearlySavings(pSaved, aSalary)
	fmt.Println("Yearly Savings:", yearlySavings)

	retire, rAge := CalculateRetirement(cAge, yearlySavings, sGoal)

	if retire {
		fmt.Printf("Goal of %v reached at age: %v\n", sGoal, rAge)
		fmt.Println("")
	} else {
		fmt.Println("Sorry, according to our calculations (dead > 100), you'll be dead before you reach your goal")
		fmt.Println("")
	}
}

// CalculateRetirement determines age of desired savings goal whether
// it's reached given a persons age, yearly savings and savings goal
func CalculateRetirement(age int, ySavings, sGoal float64) (bool, int) {
	var savings float64
	for age <= 100 {
		savings += ySavings
		age++
		if savings >= sGoal {
			return true, age
		}
	}
	return false, 0
}

// YearlySavings returns projected yearly savings given percent saved and annual salary
func YearlySavings(pSaved, annSalary float64) float64 {
	pSaved = pSaved / 100
	return (annSalary * pSaved) * 2
}
