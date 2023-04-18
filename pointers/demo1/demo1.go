package pointers

import (
	"fmt"
)

func Demo1(number *int) {
  *number = *number + 1
 fmt.Println("The number from Demo1 :" , number)
  
}
