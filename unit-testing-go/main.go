package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	for {
		fmt.Println("Welcome to the HPC Platform")
		fmt.Println("Enter Your Selection Below (1, 2, 3, 4)")
		fmt.Println("1 - Body Mass Index")
		fmt.Println("2 - Retirement Calculator")
		fmt.Println("3 - Distance Formula")
		fmt.Println("4 - Email Verifier")
		fmt.Println("0 - Exit Program")
		fmt.Print("Enter Selection: ")

		reader := bufio.NewReader(os.Stdin)
		choice, err := reader.ReadByte()
		Check(err)

		switch choice {
		case '1':
			_ = BMIInterface()
		case '2':
			RetirementInterface()
		case '3':
			_ = DistanceInterface()
		case '4':
			EmailInterface()
		case '5':
			TipInterface()
		case '0':
			os.Exit(0)
		default:
			fmt.Println("Invalid input, Exiting...")
			os.Exit(0)
		}
	}
}

// Check panic if error is present
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
