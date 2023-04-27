package string_functions

import (
	"fmt"
	"strings"
)

func Demo1() {
	name := "Arzu"
	fmt.Println(strings.Count(name, "r")) // => 1
	fmt.Println(strings.Count(name, "z")) // => 1
	fmt.Println(strings.Count(name, "b")) // => 0
	fmt.Println(strings.Contains(name, "c")) // => false
	fmt.Println(strings.Count(name, "A"))    // => 1
	fmt.Println(strings.Count(name, "u")) // => 1
	
	conclusion := strings.Index(name, "A")
	if conclusion
}

