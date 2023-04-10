package main

import (
	"fmt"
)
	
func main() {

	//Get the student's midterm and final grade
	var midtermGrade float32
	var finalGrade float32

	    fmt.Println("Please enter the midterm grade:")
	    fmt.Scanln(&midtermGrade)
	    fmt.Println("Please enter the final grade:")
	    fmt.Scanln(&finalGrade)

	//Add up midterm and final grades and take the average
	noteTotal := midtermGrade + finalGrade
	average := noteTotal / 2

	     //Print average grade to screen
	     fmt.Println("Average:", average)

	     //Ask if would like to make another calculation
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
