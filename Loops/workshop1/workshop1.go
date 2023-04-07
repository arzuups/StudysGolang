package loops

import (
	"fmt"
)

func Workshop1() {

	myNumber := 30
	estimatedNumber := 0
	forecastNumber := 0

	fmt.Println("Please give me a number:")
	fmt.Scanln(&estimatedNumber) //90
	forecastNumber = forecastNumber + 1

	for estimatedNumber != myNumber {
		if estimatedNumber < myNumber {
			fmt.Println("Enter a larger number:")
			fmt.Scanln(&estimatedNumber)
			forecastNumber = forecastNumber + 1
		}

		if estimatedNumber > myNumber {
			fmt.Println("Enter the smaller number:")
			fmt.Scanln(&estimatedNumber)
			forecastNumber = forecastNumber + 1
		}
	}

	successStatus := ""
	if forecastNumber > 0 && forecastNumber <= 3 {
		successStatus = "Super"
	} else if forecastNumber <= 6 {
		successStatus = "Beautiful"
	} else {
		successStatus = "Not bad"
	}

	fmt.Printf("You got it right!!! %v estimated : %v", forecastNumber, successStatus)

}

