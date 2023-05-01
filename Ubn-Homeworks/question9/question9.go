/*Prepare the code and flowchart that sorts the items in a user-generated array in ascending order.*/

package main

import (
	"fmt"
)

func main() {
	arr := []int{29, 8, 5, 15, 32, 18, 1, 40, 63, 34}
	
	for i := 0; i < len(arr); i++ { // A loop is started to traverse the array.

		for j := i; j < len(arr); j++ { // A nested loop is started.

			if arr[j] < arr[i] { // If element j is less than element i...

				arr[i], arr[j] = arr[j], arr[i] // ... i. and j. elements are swapped.
			}
		}
	}
	fmt.Println(arr) // Prints the array to the screen.

}
