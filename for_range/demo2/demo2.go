package for_range

import (
	"fmt"
)

func Demo2() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0

	for _, number := range numbers {
		if number%2 != 0 {
			sum = sum + number

		}
		fmt.Println("Sum :", sum)
	}

}

