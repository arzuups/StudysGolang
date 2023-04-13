/*You need to run main.go to print to the terminal.*/
package slices

import (
	"fmt"
)

func Demo1() {
	
	names := make([]string, 4)
	fmt.Println(names)
	
	names[0] = "Michale"
	names[1] = "Olivia"
	names[2] = "Tom"
	names[3] = "Jack"
	names = append(names, "Cristiano")
	fmt.Println(names)
	fmt.Println(len(names))
  
}
