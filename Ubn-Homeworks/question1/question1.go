/*Find the sum of the two defined numbers.
1-)Create two variables n, m and assign values to them.
2) Create a variable called sum.
3) Add the variables n and m and save the result in the variable sum.
4-)Print the result.*/ 

package main

import (
	"fmt"
)
func main() {

	var num1, num2 int
	fmt.Println("Enter the First Number :=  ")
	fmt.Scanln(&num1)

	fmt.Println("Enter the Second Number :=  ")
	fmt.Scanln(&num2)

	fmt.Println("The Sum of num1 and num2  :=  ", num1+num2)

}
