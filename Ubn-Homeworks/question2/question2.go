/*Prepare the code and flowchart that prints whether a number is odd or even on the screen.*/

package main

import (
	"fmt"
)

func main() {
	//Get number from user
	fmt.Println("Please enter a number:")
	var number int
	fmt.Scanln(&number)

	//Check if the entered number is even
	if number%2 == 0 {
		fmt.Println(number, "even")
	} else {
		fmt.Println(number, "odd")

	}

}
