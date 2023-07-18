// Golang Basic Lesson-7
package main

import "fmt"

func main() {

	//ELSE-İF == bool
	var isİtComing bool = true

	if isİtComing {
		fmt.Println("Yes,it is coming.")
	} else {
		fmt.Println("No, it is not coming.")
	}

	var myNumber int = 14

	if myNumber <= 14 {
		fmt.Println("My number is less than 14.")
	} else if myNumber == 14 {
		fmt.Println("My number is equal at 14.")
	} else {
		fmt.Println("My number is greater than 14.")
	}

	var productPrice int = 46

	if productPrice <= 44 {
		fmt.Println("My number is less than 14.")
	} else if productPrice >= 44 {
		fmt.Println("Product price is greater than 44.")
	} else {
		fmt.Println("Product price is equal at 44.")
	}

}

