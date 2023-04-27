package string_functions

import (
	"fmt"
	"strings"
)

func Demo1() {
	name := "Arzu"
	fmt.Println(strings.Count(name, "r")) // => 1
	fmt.Println(strings.Count(name, "z")) // => 1
	fmt.Println(strings.Count(name, "b"))
}

