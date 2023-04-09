/*Prepare the code and flowchart that calculates the body mass index according to the weight and height information entered 
by the user and prints the ideal weight status of the person according to this result.
Note-1: Body mass index = Weight / Height * Height
Note-2: Weight => kg, Height => m
Note-3: Underweight if body mass index < 18.5
-BMI < 24.9 if 18.5 < Normal
-Overweight if 25 < BMI < 29.9
-30 < BMI < 34.9 if grade I obese
-35 < BMI < 39.9 if grade II obese
-Grade III obese with a body mass index > 40*/

package main

import ( 
	"fmt"
)

func main() {
	//Get weight and height information from user
	var weight, height float64
	fmt.Println("Enter your weight in kg: ")
	fmt.Scanln(&weight)
	fmt.Println("Enter your height in m: ")
	fmt.Scanln(&height)

	//Calculate body mass index
	bmi := weight / (height * height)

	//Print your ideal weight status
	if bmi < 18.5 {
		fmt.Println("Weak")
	} else if bmi < 24.9 {
		fmt.Println("Normal")
	} else if bmi < 29.9 {
		fmt.Println("Overweight")
	} else if bmi < 34.9 {
		fmt.Println("1st degree obese")
	} else if bmi < 39.9 {
		fmt.Println("2st degree obese")
	} else {
		fmt.Println("3st degree obese")
	}
}
