// Golang Basic Lesson-11
package main

import "fmt"

func main() {

	// LOOPS
	/*Continue => Proceeds to the next step.
	  Break => Get out of the loop.*/
	var firstNumber int = 0
	var numbers [8]int = [8]int{4, 6, 8, 7, 3, 2, 9, 1}

	for i := 0; i < 8; i++ {
		firstNumber = numbers[i]
		firstNumber *= firstNumber
		fmt.Println(firstNumber)
	}

	var randomNumbers [25]int = [25]int{8, 43, 56, 786, 34, 23, 45, 67, 89, 12, 134, 56, 178, 90, 23, 145, 57, 89, 12, 34, 56, 78, 90, 23, 45}
	var oddNumberQuantity int = 0
	var evenNumberQuantity int = 0
	var isİtAvalidSeries bool = true

	for i := 0; i < len(randomNumbers); i++ {

		var nextNumber int = randomNumbers[i]
		fmt.Println(nextNumber)
		if nextNumber == 0 {
			isİtAvalidSeries = false
			break
		}
		if nextNumber > 99 {
			continue
		}

		if nextNumber%2 == 0 {
			evenNumberQuantity++
		} else {
			oddNumberQuantity++
		}
	}

	if isİtAvalidSeries {
		fmt.Println("Odd Number Quantity: ", oddNumberQuantity)
		fmt.Println("Even Number Quantity: ", evenNumberQuantity)
	} else {
		fmt.Println("Array invalid...")
	}

	fmt.Println("Finish.")
}
