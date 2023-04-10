
package arrays

import "fmt"

func Demo2() {

	var cities [5]string
	cities[0] = "Moscow"
	cities[1] = "London"
	cities[2] = "Sankt-Peterburg"
	cities[3] = "Berlin"
	cities[4] = "Prague"

	fmt.Println(cities) // This is how you print side by side.

	for i := 0; i < 5; i++ {
		fmt.Println(cities[i]) // This is how it is written, one below the other.
	}

}
