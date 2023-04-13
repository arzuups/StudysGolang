/*You need to run main.go to print to the terminal.*/
package functions

import (
	"fmt"
)

func Collect(num1 int, num2 int) int {
       var sum = num1 + num2
       return sum
}

func Hello(username string) {
       fmt.Println("Hello everyone" , username)
}
