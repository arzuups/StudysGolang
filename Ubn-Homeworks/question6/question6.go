/*Prepare the code and flowchart that finds the average of an array.*/

package main

import (
	"fmt"
)

func main() {
	//Example series
	numbers := []int{12, 24, 36, 48, 64}

	//Calculate the total
	var sum int
	for _, number := range numbers {
		sum += number
	}

	//Calculate average
	average := float64(sum) / float64(len(numbers))

	//Print the result
	fmt.Println("Average: ", average)
}
