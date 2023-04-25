/*You need to run main.go to print to the terminal.*/
package error_handling

import (
	"fmt"
	"os"
)

func Demo2() {
	f, err := os.Open("demo2.txt")
	
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("File not found!", pErr.Path)
			return
		} else {
			fmt.Println("File not found!")
			return
		}
	} else {
		fmt.Println(f.Name())
	}

}
