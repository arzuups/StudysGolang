/*You need to run main.go to print to the terminal.*/

/*You need to run main.go to print to the terminal.*/
package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	name := "Arzu"
	fmt.Println(s.HasPrefix(name, "Ar")) // => true
	fmt.Println(s.HasSuffix(name, "el")) // => false
	fmt.Println(s.Index(name, "zu"))     // => 2

	letters := []string{"a", "r", "z", "u", "u", "p", "s"}
	conclusion := s.Join(letters, "/") // => a/r/z/u/u/p/s
	fmt.Println(conclusion)            // =>a/r/z/u/u/p/s

	fmt.Println(s.Replace(conclusion, "*", "+", 3))
	fmt.Println(s.Split(conclusion, "-")) //=> [a/r/z/u/u/p/s]
	fmt.Println(s.Repeat(conclusion, 6))  // => a/r/z/u/u/p/sa/r/z/u/u/p/sa/r/z/u/u/p/sa/r/z/u/u/p/sa/r/z/u/u/p/sa/r/z/u/u/p/s

}
