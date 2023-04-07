package loops

import "fmt"

func Workshop1() {

	myNumber := 30
	estimatedNumber := 0

	fmt.Println("Please give me a number:")
	fmt.Scanln(&estimatedNumber) //90

	for estimatedNumber != myNumber {
		if estimatedNumber < myNumber {
			fmt.Println("Enter a larger number:")
			fmt.Scanln(&estimatedNumber)
		}

		if estimatedNumber > myNumber {
			fmt.Println("Enter the smaller number:")
			fmt.Scanln(&estimatedNumber) //85
		}
	}

	if estimatedNumber == myNumber {
		fmt.Println("You got it right!!!")
	}
}
