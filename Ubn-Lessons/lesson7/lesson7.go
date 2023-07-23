// Golang Basic Lesson-7
package main

import "fmt"

func main() {

	//CONDITIONALS
	//ELSE-İF == bool
	var isİtComing bool = true

	if isİtComing {
		fmt.Println("Yes,it is coming.")
	} else {
		fmt.Println("No, it is not coming.")
	}

	var myNumber int = 14
	youNumber := 20

	if myNumber <= youNumber {
		fmt.Println("My number is less than youNumber.")
	} else {
		fmt.Println("My number is greater than youNumber.")
	}

	productPrice := 46
	myProduct := 50

	if productPrice >= myProduct {
		fmt.Println("Product price is greater than my product.")
	} else if productPrice == myProduct {
		fmt.Println("Product price is equal than my product.")
	} else {
		fmt.Println("Product price is less than my product.")

	}

	var myMoney int = 100
	var myShoes int = 80

	if myShoes > myMoney {
		fmt.Println("Money exceeded...")
	} else if myShoes == myMoney {
		fmt.Println("The price of money and shoes are equal...")
	} else {
		fmt.Println("The price of the shoe is lower than the money...")
	}

	number1 := 10
	number2 := 20

	var ısItBıg bool = number1 > number2
	fmt.Println(ısItBıg)

	var food string = "Olives"
	amountPayable := 0

	if food == "Cheese" {
		amountPayable = 90
	} else if food == "Egg" {
		amountPayable = 60
	} else if food == "Bread" {
		amountPayable = 6
	} else if food == "Tomato" {
		amountPayable = 14
	} else if food == "Olives" {
		amountPayable = 40
	} else {
		amountPayable = 0
	}

	fmt.Print("Amount payable :")
	fmt.Println(amountPayable)

}
