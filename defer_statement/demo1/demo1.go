package defer_statement

import (
	"fmt"
)

func A() {
	defer B()
	defer C()
	defer D()
  fmt.Println("A function worked!")
} 

func B() {
  fmt.Println("B function worked!)
}
