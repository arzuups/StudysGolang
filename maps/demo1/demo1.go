/*You need to run main.go to print to the terminal.*/
package main

import (
	"fmt"
)

func Demo1() {
      //key-value
	Dictionary := make(map[string]string)
	
	Dictionary["Table"] = "Masa"
	Dictionary["Book"] = "Kitap"
	Dictionary["Pencil"] = "Kalem"
	Dictionary["Notebook"] = "Not Defteri"
	
	fmt.Println(Dictionary)
        fmt.Println("Number of elements:", len(Dictionary)) //If we want to print the number of elements
}
