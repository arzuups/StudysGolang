/*You need to run main.go to print to the terminal.*/
package loops

import (
	"fmt"
)

func Workshop3() {
	number1 := 220
	number2 := 284
	sum1 := 0
	sum2 := 0

	for i := 1; i < number1; i++ {
		if number1%i == 0 {
			sum1 = sum1 + i
		}
	}

	for i := 1; i < number2; i++ {
		if number2%i == 0 {
			sum2 = sum2 + i

		}

	}
	if sum1 == number2 && sum2 == number1 {
		fmt.Printf("%v and %v are friend numbers.", number1, number2)
	} else {
		fmt.Printf("%v and %v are not friendly numbers.", number1, number2)
	}

}
