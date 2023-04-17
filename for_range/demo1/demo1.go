/*You need to run main.go to print to the terminal.*/
package for_range

import (
	"fmt"
)

func Demo1() {

	cities := []string{"Prague", "Istanbul", "New York"}

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])

	}

	for i, city := range cities {
		fmt.Print(i)
		fmt.Println(city)

	}

}

