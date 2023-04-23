/*You need to run main.go to print to the terminal.*/
package defer_statement

import (
	"fmt"
)
 
func Demo2(number int32) string {
	defer DeferFunc()
	
	if number%2 == 0 {
	    return "This is even number!"
	}
}
