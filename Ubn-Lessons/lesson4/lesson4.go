//Golang Basic Lesson-4
package main

import "fmt"

func main() {

	//OPERATORS
	//(+ - * / %)

	//(+)
	num1 := 13
	num2 := 34
	sum := num1 + num2
	fmt.Println(sum)

	//(-)
	var number1 float32 = 12.54
	var number2 float32 = 13.43
	var extraction float32 = number1 - number2
	fmt.Println(extraction)

	//(*)
	myNumber1 := 12
	myNumber2 := 43
	Multiplication := myNumber1 * myNumber2
	fmt.Println(Multiplication)

	//(/)
	var queueNumber1 int = 60
	var queueNumber2 int = 100
	var Partition int = queueNumber1 / queueNumber2
	fmt.Println(Partition)
}
