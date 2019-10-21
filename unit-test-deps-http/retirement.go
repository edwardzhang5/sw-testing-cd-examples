package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var aSalary, pSaved, sGoal float64
var cAge int

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
func RetirementInterface(c *mongo.Client) {
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

	retire, rAge := CalculateRetirement(cAge, yearlySavings, sGoal)
	fmt.Println(BuildTimeStamp(time.Now()))

	retData := RetData{
		Goal:         sGoal,
		PercentSaved: pSaved,
		RetireAge:    rAge,
		CurrentAge:   cAge,
		YearlySaved:  yearlySavings,
		TimeStamp:    BuildTimeStamp(time.Now()),
	}

	if retire {
		mes := fmt.Sprintf("Goal of %v reached at age: %v\n", sGoal, rAge)
		retData.Message = mes
		fmt.Println(mes)
		fmt.Println("")
	} else {
		mes := "Sorry, according to our calculations (dead > 100), you'll be dead before you reach your goal"
		fmt.Println(mes)
		retData.Message = mes
		fmt.Println("")
	}
	collection := c.Database("swtest").Collection("retire")
	_, err = collection.InsertOne(context.Background(), retData)
	if err != nil {
		log.Fatalf("Could not write data: %+v\n to DB with err: %v", retData, err)
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

// GetRetirementEntries returns a slice of the retire data entries from the DB
func GetRetirementEntries(c *mongo.Client) []RetData {
	var re []RetData
	collection := c.Database("swtest").Collection("retire")
	ctx := context.Background()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var retRes RetData
		err := cur.Decode(&retRes)
		if err != nil {
			log.Fatalf("Could not get Retirement data: %v", err)
		}

		re = append(re, retRes)
	}
	if err := cur.Err(); err != nil {
		log.Fatalf("Could not get Retirement data: %v", err)
	}

	return re
}
