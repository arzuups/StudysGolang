package structs

import (
	"fmt"
)

func Demo1() {
  fmt.Println(product{"Telephone" , 5.500 , "Iphone"})
  //fmt.Println(product{name: "Telephone", unitPrice: 5.500, bland: "Iphone"}) //fmt.Println(product{"Telephone", 5.500, "Iphone"})

}

type product struct{
  name string
  unitPrice float64
  bland string
  
}
