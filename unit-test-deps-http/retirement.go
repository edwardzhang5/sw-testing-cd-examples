package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var aSalary, pSaved, sGoal float64
var cAge int

// RetData containes retirement data entry
type RetData struct {
	CurrentAge   int     `json:"current_age"`
	RetireAge    int     `json:"retire_age"`
	Salary       float64 `json:"salary"`
	PercentSaved float64 `json:"percent_saved"`
	Goal         float64 `json:"goal"`
	YearlySaved  float64 `json:"yearly_saved"`
	Message      string  `json:"message"`
	TimeStamp    string  `json:"time_stamp"`
}

// RetirementInterface gets user input and provides output
func RetirementInterface(c Session) {
	var err error

	fmt.Println("**PRIOR Entries from DB**")

	for i, r := range GetRetirementEntries(c) {
		fmt.Printf("ENTRY #%d: %+v\n\n", i, r)
	}

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

	rAge, message := CalculateRetirement(cAge, yearlySavings, sGoal)

	retData := RetData{
		Goal:         sGoal,
		PercentSaved: pSaved,
		RetireAge:    rAge,
		CurrentAge:   cAge,
		Salary:       aSalary,
		YearlySaved:  yearlySavings,
		Message:      message,
		TimeStamp:    BuildTimeStamp(time.Now()),
	}

	err = WriteRetireData(c, retData)
	if err != nil {
		log.Fatalf("Could not write data: %+v\n to DB with err: %v", retData, err)
	}
	fmt.Println(message)
}

// CalculateRetirement determines age of desired savings goal whether
// it's reached given a persons age, yearly savings and savings goal
func CalculateRetirement(age int, ySavings, sGoal float64) (int, string) {
	var savings float64
	var message string
	for age < 100 {
		savings += ySavings
		age++
		if savings >= sGoal {
			message = fmt.Sprintf("Goal of %v reached at age: %v\n", sGoal, age)
			return age, message
		}
	}
	message = "Sorry, according to our calculations (dead > 100), you'll be dead before you reach your goal"
	return 0, message

}

// YearlySavings returns projected yearly savings given percent saved and annual salary
func YearlySavings(pSaved, annSalary float64) float64 {
	pSaved = pSaved / 100
	return (annSalary * pSaved) * 2
}

// WriteRetireData writes the data to the database
func WriteRetireData(c Session, ret RetData) error {
	var err error
	collection := c.DB("swtest").C("retire")
	err = collection.Insert(ret)

	return err
}

// RetireEndpoint returns all Retirement entries
func (dbh *DBHandler) RetireEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, GetRetirementEntries(dbh.Session))
}

// RetireHandler handles the Retirement http input request
func (dbh *DBHandler) RetireHandler(c *gin.Context) {
	var err error
	savingsGoal, _ := strconv.ParseFloat(c.Query("sGoal"), 64)
	currentAge, _ := strconv.Atoi(c.Query("cAge"))
	annualSalary, _ := strconv.ParseFloat(c.Query("aSalary"), 64)
	percentSaved, _ := strconv.ParseFloat(c.Query("pSaved"), 64)
	yearlySavings := YearlySavings(percentSaved, annualSalary)

	rAge, message := CalculateRetirement(currentAge, yearlySavings, savingsGoal)

	retData := RetData{
		Goal:         savingsGoal,
		PercentSaved: percentSaved,
		RetireAge:    rAge,
		CurrentAge:   currentAge,
		Salary:       annualSalary,
		YearlySaved:  yearlySavings,
		Message:      message,
		TimeStamp:    BuildTimeStamp(time.Now()),
	}

	err = WriteRetireData(dbh.Session, retData)

	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, retData)
}

// GetRetirementEntries returns a slice of the retire data entries from the DB
func GetRetirementEntries(c Session) []RetData {
	var re []RetData
	collection := c.DB("swtest").C("retire")
	err := collection.Find(nil).All(&re)
	if err != nil {
		log.Fatalf("Could not parse Retire data: %+v", err)
	}
	return re
}
