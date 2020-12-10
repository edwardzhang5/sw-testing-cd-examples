package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var x1, x2, y1, y2 float64

// DistanceHandler returns distance for given (x1, x2), (y1, y2)
func DistanceHandler(c *gin.Context) {
	var err error
	x1, err = strconv.ParseFloat(c.Query("x1"), 64)
	if err != nil {
		c.Error(err)
		return
	}
	x2, err = strconv.ParseFloat(c.Query("x2"), 64)
	if err != nil {
		c.Error(err)
		return
	}
	y1, err = strconv.ParseFloat(c.Query("y1"), 64)
	if err != nil {
		c.Error(err)
		return
	}
	y2, err = strconv.ParseFloat(c.Query("y2"), 64)
	if err != nil {
		c.Error(err)
		return
	}
	result := CalculateDistance(x1, x2, y1, y2)

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Distance for (%.2f, %.2f), (%.2f, %.2f) = %.5f", x1, x2, y1, y2, result), "result": result})
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
