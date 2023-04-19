/*You need to run main.go to print to the terminal.*/
package structs

import (
	"fmt"
)

func Demo1() {
  fmt.Println(product{"Telephone" , 5.500 , "Iphone"})
  //fmt.Println(product{name: "Telephone", unitPrice: 5.500, brand: "Iphone"}) //If we want to identify them by name

}

type product struct{
  name string
  unitPrice float64
  brand string
  
}
