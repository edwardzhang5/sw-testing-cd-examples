package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var salary, saved, goal float64
var age int64

// RetirementHandler calculates retirement info
func RetirementHandler(c *gin.Context) {
	var err error
	age, err = strconv.ParseInt(c.Query("age"), 10, 64)
	if err != nil {
		c.Error(err)
		return
	}
	salary, err = strconv.ParseFloat(c.Query("salary"), 64)
	if err != nil {
		c.Error(err)
		return
	}
	saved, err = strconv.ParseFloat(c.Query("saved"), 64)
	if err != nil {
		c.Error(err)
		return
	}
	goal, err = strconv.ParseFloat(c.Query("goal"), 64)
	if err != nil {
		c.Error(err)
		return
	}

	yearlySavings := YearlySavings(saved, salary)
	retirement, retAge := CalculateRetirement(age, yearlySavings, goal)

	if retirement {
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Goal of %v reached at age: %v\n", goal, retAge), "reached": true, "retAge": retAge})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sorry, according to our calculations (dead > 100), you'll be dead before you reach your goal", "reached": false})
	return

}

// CalculateRetirement determines age of desired savings goal whether
// it's reached given a persons age, yearly savings and savings goal
func CalculateRetirement(age int64, ySavings, sGoal float64) (bool, int64) {
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
