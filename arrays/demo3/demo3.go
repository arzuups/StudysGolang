package arrays

import (
	"fmt"

	
func Demo3() {

	numbers := [5]int{13, 54, 23, 64, 87}
	fmt.Println(numbers)

	biggest := numbers[0]

	//LENGTH = UZUNLUK
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > biggest {
			biggest = numbers[i]

		}
	}
	fmt.Println("Biggest :", biggest)
}

