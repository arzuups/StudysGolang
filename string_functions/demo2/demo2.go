package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	name := "Arzu"
	fmt.Println(s.HasPrefix(name, "Ar")) // => true
	fmt.Println(s.HasSuffix(name, "el")) // => false
	fmt.Println(s.Index(name, "zu"))  // => 2
	
	letters := []string{"a", "r", "z", "u", "w", "s"}
	conclusion := s.Join(letters, "/") // => a/r/z/u/w/s
} 
