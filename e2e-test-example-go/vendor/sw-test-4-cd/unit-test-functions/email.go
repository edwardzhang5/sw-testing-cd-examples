package main

import (
	"regexp"

	"github.com/gin-gonic/gin"
	"net/http"
)

// EmailHandler HTTP validates email received from post request
func EmailHandler(c *gin.Context) {
	email := c.Query("email")
	if (ValidateEmail(email)) {
		c.JSON(http.StatusOK, gin.H{"email": email, "valid": true, "message": "Email is Validated"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"email": email, "valid": false, "message": "Email address is Invalid"})
}

// ValidateEmail returns valid email address
func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,3}$`)
	return Re.MatchString(email)
}
