package main

import (
	"fmt"
)

func main() {

	var number1, number2 float64
	var calculator string

	     fmt.Print("Please enter the first number:")
	     fmt.Scanln(&number1)

	     fmt.Print("Please enter the second number:")
	     fmt.Scanln(&number2)

	     fmt.Print("Select the operation you want to perform(+,-,*,/):")
	     fmt.Scanln(&calculator)

	switch calculator {

	     case "+":
		       fmt.Println(number1 + number2)

	     case "-":
		       fmt.Println(number1 - number2)

	     case "*":
		       fmt.Println(number1 * number2)

	     case "/":
		       if number2 == 0 {

			     fmt.Println("404:This number is not divisible by zero")
		     } else {
			     fmt.Println(number1 / number2)

		}

	}
	         fmt.Println("Do you want to make another calculation?(Yes/No)")
	         var answer string
	         fmt.Scanln(&answer)

	//If the user wants to perform another calculation, go back to the beginning

	         if answer == "Yes" {

		//Finish if the user does not want to perform another calculation

	      } else if answer == "No" {
		      fmt.Println("The transaction is complete. Thank you.")
	   }

}
