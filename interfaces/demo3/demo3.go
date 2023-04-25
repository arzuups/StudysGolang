/*You need to run main.go to print to the terminal.*/
package interfaces
import (
	"fmt"
)

func verify(i interface{}) {
    number, ok := i.(int)
    fmt.Println(number, ok)
}

func Demo3() {
    var number1 interface{} = "Julia"
    verify(number1)
	
    var number2 interface{} = 20
    verify(number2)
}
