/*Prepare the code and flowchart that prints the odd numbers from 1 to 100 on the screen.*/

package main

import "fmt"

func main() {
	//Variable i is incremented by 1 in each cycle,
	//starting from 1. Returns all numbers up to 100

	for i := 1; i <= 100; i++ {

		//If i is an odd number, print to screen

		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}
