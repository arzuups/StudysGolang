/*You need to run main.go to print to the terminal.*/
package loops

import (
  "fmt"
)

func Demo1() {

	var text string = "Hello World...Welcome the Go Programming Language..."

	i := 1

	for i <= 100 {
		fmt.Println(text)
		i = i + 1
	}

	fmt.Println("Finish!")

}
