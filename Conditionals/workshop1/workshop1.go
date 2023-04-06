package conditionals

import (
	"fmt"
)
	
func Workshop1() {
	var num1, num2, num3 int = 12, 45, 23

	var max int = num1

	if num2 > max {
		max = num2
	}

	if num3 > max {
		max = num3
	}
	fmt.Printf("Max number : %d", max)
}
