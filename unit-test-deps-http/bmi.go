package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//BCat enum values for BMI Category
type BCat int

// category constants
const (
	Unknown BCat = iota
	Underweight
	Normal
	Overweight
	Obese
)

type BMI struct {
	HFeet     float64 `json:"height_feet"`
	HInches   float64 `json:"height_inches"`
	Weight    float64 `json:"weight"`
	BMIResult string  `json:"bmi_result"`
	TimeStamp string  `json:"time_stamp"`
}

// BMIInterface gets user input and runs function
func BMIInterface(c *mongo.Client) error {
	var hFeet, hInches, pWeight float64
	var err error

	fmt.Println("**PRIOR Entries from DB**")

	for i, r := range GetBMIEntries(c) {
		fmt.Printf("ENTRY #%d: %+v\n\n", i, r)
	}

	fmt.Println("*** BMI Formula ***")
	fmt.Println("Enter value for Height in Feet & Inches and Weight in Pounds")
	fmt.Print("Height in feet: ")
	_, err = fmt.Scanf("%f", &hFeet)
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
	bmiResult := ""

	fmt.Printf("BMI = %.2f\n", bmi)
	switch bmiCat {
	case Underweight:
		bmiResult = "Underweight: BMW < 18.5"
		fmt.Println(bmiResult)
	case Normal:
		bmiResult = "Normal weight: BMI 18.5–24.99 "
		fmt.Println(bmiResult)
	case Overweight:
		bmiResult = "Overweight: BMI 25.0-29.99 "
		fmt.Println(bmiResult)
	case Obese:
		bmiResult = "Obese: BMI of 30 or greater"
		fmt.Println(bmiResult)
	case Unknown:
		fmt.Println("Unable to determine BMI Category")
	}

	// write to DB
	bmiData := BMI{
		HFeet:     hFeet,
		HInches:   hInches,
		Weight:    pWeight,
		BMIResult: bmiResult,
		TimeStamp: BuildTimeStamp(time.Now()),
	}

	collection := c.Database("swtest").Collection("bmi")
	_, err = collection.InsertOne(context.Background(), bmiData)
	if err != nil {
		log.Fatalf("Could not write data: %+v\n to DB with err: %v", bmiData, err)
	}

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

// GetBMIEntries returns a slice of the retire data entries from the DB
func GetBMIEntries(c *mongo.Client) []BMI {
	var be []BMI
	collection := c.Database("swtest").Collection("bmi")
	ctx := context.Background()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal("Unable to Parse Collection:  ", err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var bmiRes BMI
		err := cur.Decode(&bmiRes)
		if err != nil {
			log.Fatalf("Could not get BMI data: %v", err)
		}

		be = append(be, bmiRes)
	}
	if err := cur.Err(); err != nil {
		log.Fatalf("Could not get BMI data: %v", err)
	}

	return be
}

// https://www.nhlbi.nih.gov/health/educational/lose_wt/BMI/bmicalc.htm

// http://extoxnet.orst.edu/faqs/dietcancer/web2/twohowto.html
