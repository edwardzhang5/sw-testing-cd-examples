package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	// mgo "github.com/globalsign/mgo"
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

// BMI contains bmi entry data
type BMI struct {
	HFeet     float64 `json:"height_feet"`
	HInches   float64 `json:"height_inches"`
	Weight    float64 `json:"weight"`
	BMI       float64 `json:"bmi"`
	BMIResult string  `json:"bmi_result"`
	TimeStamp string  `json:"time_stamp"`
}

// BMIInterface gets user input and runs function
func BMIInterface(c Session) error {
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

	fmt.Printf("BMI = %.2f\n", bmi)
	fmt.Println(getBMIMessage(bmiCat))
	// write to DB
	bmiData := BMI{
		HFeet:     hFeet,
		HInches:   hInches,
		Weight:    pWeight,
		BMI:       bmi,
		BMIResult: getBMIMessage(bmiCat),
		TimeStamp: BuildTimeStamp(time.Now()),
	}

	err = WriteBMIData(c, bmiData)

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

func getBMIMessage(bmiCat BCat) string {
	bmiResult := ""
	switch bmiCat {
	case Underweight:
		bmiResult = "Underweight: BMW < 18.5"
	case Normal:
		bmiResult = "Normal weight: BMI 18.5–24.99 "
	case Overweight:
		bmiResult = "Overweight: BMI 25.0-29.99 "
	case Obese:
		bmiResult = "Obese: BMI of 30 or greater"
	case Unknown:
		bmiResult = "Unable to determine BMI Category"
	}
	return bmiResult

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

// WriteBMIData writes the data to the database
func WriteBMIData(c Session, bmi BMI) error {
	var err error
	collection := c.DB("swtest").C("bmi")
	err = collection.Insert(bmi)

	return err
}

// GetBMIEntries returns a slice of the retire data entries from the DB
func GetBMIEntries(c Session) []BMI {
	var be []BMI
	collection := c.DB("swtest").C("bmi")
	err := collection.Find(nil).All(&be)
	if err != nil {
		log.Fatal("Unable to Parse BMI Collection:  ", err)
	}
	return be
}

// BMIEndpoint returns all BMI entries as JSON
func (dbh *DBHandler) BMIEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, GetBMIEntries(dbh.Session))
}

// BMIHandler handles the BMI Request
func (dbh *DBHandler) BMIHandler(c *gin.Context) {
	var err error
	feet, err := strconv.ParseFloat(c.Param("feet"), 64)
	if err != nil {
		c.Error(err)
	}
  inches, err := strconv.ParseFloat(c.Param("inches"), 64)
  if err != nil {
		c.Error(err)
	}
  weight, err := strconv.ParseFloat(c.Param("weight"), 64)
  if err != nil {
		c.Error(err)
	}

	bmi, _ := CalculateBMI(feet, inches, weight)

	bmiData := BMI{
		HFeet:     feet,
		HInches:   inches,
		Weight:    weight,
		BMI:       bmi,
		BMIResult: getBMIMessage(BMICategory(bmi)),
		TimeStamp: BuildTimeStamp(time.Now()),
	}

	err = WriteBMIData(dbh.Session, bmiData)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, bmiData)
}

// https://www.nhlbi.nih.gov/health/educational/lose_wt/BMI/bmicalc.htm

// http://extoxnet.orst.edu/faqs/dietcancer/web2/twohowto.html
