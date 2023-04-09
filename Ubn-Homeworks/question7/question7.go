/*Prepare the code and flowchart that calculates the average of a student's midterm grade and final grade. 
(Calculate by taking 30% of the midterm grade and 70% of the final grade.) (Grades will be entered by the user.)*/

package main

import (
	"fmt"
)

func main() {

	//Get the student midterm and final grade
	var midterm, final int
	fmt.Println("Please enter midterm note: ")
	fmt.Scanln(&midterm)
	fmt.Println("Please enter final grade: ")
	fmt.Scanln(&final)

	//Calculate 30% of the midterm grade

	midtermContribution := float64(midterm) * 0.3

	// Calculate 70% of final grade
	finalContribution := float64(final) * 0.7

	//Collect 30% of the midterm grade and 70% of the final grade
	noteTotal := midtermContribution + finalContribution
	average := noteTotal

	//Print average grade to screen
	fmt.Println("Average: ", average)

	//Ask if you pwant to do another calculation
	fmt.Println("Do you want to create another account? (Yes/No)")
	var answer string
	fmt.Scanln(&answer)

	//Back to top if user wants to do another calculation
	if answer == "Yes" {

		//Finish if the user doesn't want to do any other calculations

	} else if answer == "No" {
		fmt.Println("Process completed. We thank you.")

	}
}
