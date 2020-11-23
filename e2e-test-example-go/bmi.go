package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
)

//BCat enum values for BMI Category
type BCat int

// category constants
const (
	Unknown     BCat = iota
	Underweight
	Normal
	Overweight
	Obese
)

// BMIHandler calculates BMI Function
func BMIHandler(c *gin.Context) {
	var err error
	var hFeet, hInches, weight float64
	hFeet, err = strconv.ParseFloat(c.Query("hFeet"), 64)
	if err != nil {
		c.Error(err)
		return
	}
	hInches, err = strconv.ParseFloat(c.Query("hInches"), 64)
	if err != nil {
		c.Error(err)
		return
	}
	weight, err = strconv.ParseFloat(c.Query("weight"), 64)
	if err != nil {
		c.Error(err)
		return
	}

	bmi, err := CalculateBMI(hFeet, hInches, weight)
	if err != nil {
		c.Error(err)
		return
	}

	bmiCat := BMICategory(bmi)
	var message string
	switch bmiCat {
	case Underweight:
		message = "Underweight: BMW < 18.5"
	case Normal:
		message = "Normal weight: BMI 18.5–24.99 "
	case Overweight:
		message = "Overweight: BMI 25.0-29.99 "
	case Obese:
		message = "Obese: BMI of 30 or greater"
	case Unknown:
		message = "Unable to determine BMI Category"
	}
	c.JSON(http.StatusOK, gin.H{"message": message, "BMI": fmt.Sprintf("%.2f", bmi)})
}

// BMIInterface gets user input and runs function
func BMIInterface() error {
	var hFeet, hInches, pWeight float64

	fmt.Println("")
	fmt.Println("*** BMI Formula ***")
	fmt.Println("Enter value for Height in Feet & Inches and Weight in Pounds")
	fmt.Print("Height in feet: ")
	_, err := fmt.Scanf("%f", &hFeet)
	if err != nil {
		fmt.Println("***\nInvalid Input for feet, need a single numeric value...returning to main menu\n\n***")
		return err
	}

	fmt.Print("Height in inches: ")
	_, err = fmt.Scanf("%f", &hInches)
	if err != nil {
		fmt.Println("***\nInvalid Input for inches, need a single numeric value...returning to main menu\n\n***")
		return err
	}

	fmt.Print("Weight in pounds: ")
	_, err = fmt.Scanf("%f", &pWeight)
	if err != nil {
		fmt.Println("***\nInvalid Input for weight, need a single numeric value...returning to main menu\n\n***")
		return err
	}

	bmi, err := CalculateBMI(hFeet, hInches, pWeight)
	if err != nil {
		fmt.Println(err)
		return err
	}

	bmiCat := BMICategory(bmi)

	fmt.Printf("BMI = %.2f\n", bmi)
	switch bmiCat {
	case Underweight:
		fmt.Println("Underweight: BMW < 18.5")
	case Normal:
		fmt.Println("Normal weight: BMI 18.5–24.99 ")
	case Overweight:
		fmt.Println("Overweight: BMI 25.0-29.99 ")
	case Obese:
		fmt.Println("Obese: BMI of 30 or greater")
	case Unknown:
		fmt.Println("Unable to determine BMI Category")
	}
	fmt.Println("")

	fmt.Println("hello")
	return nil
}

// CalculateBMI returns body mass index
func CalculateBMI(feet, inches, pounds float64) (float64, error) {
	if feet < 0 || inches < 0 || pounds < 0 {
		// return 0, errors.New("Negative number not acceptable for any input")
		return 0, fmt.Errorf("***\n negative number not acceptable for any input\n***")
	}
	mWeight := pounds * 0.45
	mHeight := MetersHeightFromInches(InchesHeight(feet, inches))
	mHSquared := Squared(mHeight)
	return mWeight / mHSquared, nil
}

// InchesHeight returns height in inches given ft + in
func InchesHeight(f float64, i float64) float64 {
	return (f * 12) + i
}

// MetersHeightFromInches converts inches to meters
func MetersHeightFromInches(inches float64) float64 {
	return inches * 0.025
}

// BMICategory returns the category of BMI status
func BMICategory(bmi float64) BCat {
	switch {
	case bmi < 18.5:
		return Underweight
	case bmi >= 18.5 && bmi < 25:
		return Normal
	case bmi >= 25.0 && bmi < 30:
		return Overweight
	case bmi >= 30:
		return Obese
	}
	return Unknown
}

// https://www.nhlbi.nih.gov/health/educational/lose_wt/BMI/bmicalc.htm

// http://extoxnet.orst.edu/faqs/dietcancer/web2/twohowto.html
