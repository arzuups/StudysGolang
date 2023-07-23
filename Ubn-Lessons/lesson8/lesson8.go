// Golang Basic Lesson-8
package main

import "fmt"

func main() {

	//CONDITIONALS
	// && => And => Ve
	// || => Or => Veya
	// true && true = true
	// false || false = false

	englishGrade := 90
	mathGrade := 70
	passingGrade := 50

	if englishGrade > passingGrade && mathGrade > passingGrade {
		fmt.Println("You are successful:)")
	} else {
		fmt.Println("Unfortunately, you have to work harder:(")
	}

	var weight float32 = 108.90
	var training int = 20

	if (weight > 100 && training > 10) ||
		(weight < 100 && training < 8) {
		fmt.Println("Congrats!!")
	} else {
		fmt.Println("Sorry, result failed...")
	}

}

