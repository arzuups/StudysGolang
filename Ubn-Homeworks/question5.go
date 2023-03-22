/*Prepare the code and flowchart that prints the smallest element in an array to the screen.*/

package main

import "fmt"

func main() {
	var array [10]int = [10]int{45, 6, 23, 8, 16, 31, 1, 65, 26, 10}
	minimumNumber := array[0]
	for i := 0; i < len(array); i++ {
		if array[i] < minimumNumber {
			minimumNumber = array[i]
		}
	}
	fmt.Println(minimumNumber)

}
