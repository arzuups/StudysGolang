package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	name := "Arzu"
	fmt.Println(s.HasPrefix(name, "Ar")) // => true
	fmt.Println(s.HasSuffix(name, "el"))
}
