package string_functions

import (
	"fmt"
	s "strings"
)

func Demo1() {
	name := "Arzu"
	fmt.Println(s.Count(name, "r")) // => 1
	fmt.Println(s.Count(name, "z")) // => 1
	fmt.Println(s.Count(name, "b")) // => 0
	fmt.Println(s.Contains(name, "c")) // => false
	fmt.Println(s.Count(name, "A"))    // => 1
	fmt.Println(s.Count(name, "u")) // => 1
	
	conclusion := s.Index(name, "A")
	
	if conclusion != -1 {
		fmt.Println("There's the letter `A`")
	} else {
		fmt.Println("No letter `A`")
	}
	fmt.Println(s.ToLower(name)) // => arzu
	fmt.Println(s.ToUpper(name)) // => ARZU
}

