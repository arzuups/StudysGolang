/*Prepare the ko and flowchart that prints the sum of numbers from 1 to 100 on the screen.*/

package main

import (
	"fmt"
)

func main() {
	//Calculate the sum of wings from 1 to 100

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}

	//Print the sum to the screen
	fmt.Println("Sum of numbers between 1 and 100: ", sum)
}
