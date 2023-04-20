/*You need to run main.go to print to the terminal.*/
package goroutines

import (
	"fmt"
	"time"
)

func EvenNumbers() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("Even number : ", i)
		time.Sleep(1 * time.Second)
	}

}

func OddNumbers() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Odd number : ", i)
		time.Sleep(1 * time.Second)

	}

}
