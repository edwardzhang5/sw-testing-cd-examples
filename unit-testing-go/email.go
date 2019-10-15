package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var email string

// EmailInterface verifies email addresses
func EmailInterface() {

	fmt.Println("")
	fmt.Println("*** Email Address Verifier ***")
	fmt.Println("Please enter the email address you want to verify")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Email: ")

	email, err := reader.ReadString('\n')
	Check(err)

	email = strings.Trim(email, "\n")
	if ValidateEmail(email) {
		fmt.Printf("The Email Address '%s' is Valid\n", email)
	} else {
		fmt.Printf("The Email Address '%s' is *NOT* Valid\n", email)
	}
	fmt.Println("")
}

// ValidateEmail returns valid email address
func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,3}$`)
	return Re.MatchString(email)
}
